package e2e_test

import (
	"flag"
	"os/exec"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/shipengqi/golib/fsutil"

	. "github.com/shipengqi/jaguar/test/e2e"
)

func init() {
	flag.StringVar(&CliOpts.Cli, "cli", "", "path to the jaguar command to use.")
}

var _ = Describe("Sorted Tests", func() {
	Describe("License Command", LicenseTest)
	Describe("CodeGen Command", CodeGenTest)
	Describe("New API project", NewAPITest)
	Describe("New CLI project", NewCLITest)
	Describe("New gRPC project", NewGRPCTest)
})

var (
	se  *gexec.Session
	err error
)

func TestE2e(t *testing.T) {
	// Skip running E2E tests when running only "short" tests because:
	// 1. E2E tests are long-running tests involving generation of skeletons.
	if testing.Short() {
		t.Skip("Skipping E2E tests")
	}

	RegisterFailHandler(Fail)
	RunSpecs(t, "E2E Suite")
}

var _ = BeforeSuite(func() {
	flag.Parse()

	if CliOpts.Cli == "" {
		CliOpts.Cli = "./jaguar"
	}
})

// ===================================================
// Helpers

func RunCLITest(args ...string) (*gexec.Session, error) {
	return RunCommandTest(CliOpts.Cli, args...)
}

func RunCommandTest(command string, args ...string) (*gexec.Session, error) {
	cmd := exec.Command(command, args...)
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	return session.Wait(), err
}

func NoError(err error) {
	Expect(err).To(BeNil())
}

func ExitCode(session *gexec.Session, expected int) {
	Ω(session.ExitCode()).Should(Equal(expected))
}

func ShouldContains(session *gexec.Session, expected string) {
	Ω(session.Out.Contents()).Should(ContainSubstring(expected))
}

func ShouldNotContains(session *gexec.Session, expected string) {
	Ω(session.Out.Contents()).ShouldNot(ContainSubstring(expected))
}

func ShouldExists(fpath string) {
	exists := fsutil.IsExists(fpath)
	Expect(exists).To(Equal(true))
}

func ShouldNotExists(fpath string) {
	exists := fsutil.IsExists(fpath)
	Expect(exists).To(Equal(false))
}
