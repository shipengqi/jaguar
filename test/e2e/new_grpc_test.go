package e2e_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func NewGRPCTest() {
	const (
		testRPCProjectName = "testrpc"
		testRPCModule      = "github.com/user/testrpc"
	)
	Context("New With Flag Parameters", func() {
		var outDir string
		BeforeEach(func() {
			var err error
			outDir, err = os.MkdirTemp("", "jaguar-new-grpc-*")
			Expect(err).NotTo(HaveOccurred())
		})
		AfterEach(func() {
			_ = os.RemoveAll(outDir)
		})
		tests := []NewCommandTestCase{
			genNewCommandTestCase("should create an gRPC project",
				"go-grpc", testRPCProjectName, testRPCModule,
				true, true, true, true),
			genNewCommandTestCase("should create an gRPC project but disable lint",
				"go-grpc", testRPCProjectName, testRPCModule,
				false, true, true, true),
			genNewCommandTestCase("should create an gRPC project but disable releaser",
				"go-grpc", testRPCProjectName, testRPCModule,
				true, false, true, true),
			genNewCommandTestCase("should create an gRPC project but disable semver",
				"go-grpc", testRPCProjectName, testRPCModule,
				true, true, false, true),
			genNewCommandTestCase("should create an gRPC project but disable actions",
				"go-grpc", testRPCProjectName, testRPCModule,
				true, true, true, false),
		}
		for _, t := range tests {
			testcase := t
			It(testcase.title, func() {
				cmds := append(testcase.commands, "-o", outDir)
				se, err = RunCLITest(cmds...)
				NoError(err)
				for _, v := range testcase.expects {
					ShouldContains(se, v)
				}
			})
		}
	})

	Context("New Without Flag Parameters", func() {
		It("should create an gRPC project with new command", func() {

		})
	})
}
