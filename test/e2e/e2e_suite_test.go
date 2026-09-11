package e2e_test

import (
	"flag"
	"os"
	"os/exec"
	"regexp"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

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
	Describe("New Frontend project", NewFrontendTest)
	Describe("Compile and Run", CompileAndRunTest)
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

	// For more information: https://onsi.github.io/gomega/#adjusting-output
	// format.MaxDepth = 0

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
	cmd.Env = append(os.Environ(), "NO_COLOR=1")
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	return session.Wait(30 * time.Second), err
}

func NoError(err error) {
	Expect(err).To(BeNil())
}

func ExitCode(session *gexec.Session, expected int) {
	Ω(session.ExitCode()).Should(Equal(expected))
}

var ansiEscape = regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)

func stripANSI(b []byte) string {
	return ansiEscape.ReplaceAllString(string(b), "")
}

func ShouldContains(session *gexec.Session, expected string) {
	Ω(stripANSI(session.Out.Contents())).Should(ContainSubstring(expected))
}

func ShouldContainsErr(session *gexec.Session, expected string) {
	Ω(stripANSI(session.Err.Contents())).Should(ContainSubstring(expected))
}

func ShouldNotContains(session *gexec.Session, expected string) {
	Ω(stripANSI(session.Out.Contents())).ShouldNot(ContainSubstring(expected))
}

func ShouldExists(fpath string) {
	_, err := os.Stat(fpath)
	Expect(err).To(BeNil())
}

func ShouldNotExists(fpath string) {
	_, err := os.Stat(fpath)
	Expect(os.IsNotExist(err)).To(Equal(true))
}
