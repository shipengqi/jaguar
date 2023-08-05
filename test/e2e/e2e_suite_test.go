package e2e_test

import (
	"flag"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/shipengqi/jaguar/test/e2e"
)

func init() {
	flag.StringVar(&CliOpts.Cli, "cli", "", "path to the jaguar command to use.")
}

var _ = Describe("Sorted Tests", func() {
	Describe("New API project", NewAPITest)
	Describe("New CLI project", NewCLITest)
	Describe("New gRPC project", NewGRPCTest)
})

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
})
