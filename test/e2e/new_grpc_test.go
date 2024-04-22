package e2e_test

import (
	. "github.com/onsi/ginkgo/v2"
	"os"
)

func NewGRPCTest() {
	const (
		testRPCProjectName = "testrpc"
		testRPCModule      = "github.com/user/testrpc"
	)
	Context("New With Flag Parameters", func() {
		AfterEach(func() {
			_ = os.RemoveAll(testRPCProjectName)
		})
		tests := []NewCommandTestCase{
			genNewCommandTestCase("should create an gRPC project",
				"api", testRPCProjectName, testRPCModule,
				true, true, true, true),
			genNewCommandTestCase("should create an gRPC project but disable lint",
				"api", testRPCProjectName, testRPCModule,
				false, true, true, true),
			genNewCommandTestCase("should create an gRPC project but disable releaser",
				"api", testRPCProjectName, testRPCModule,
				true, false, true, true),
			genNewCommandTestCase("should create an gRPC project but disable semver",
				"api", testRPCProjectName, testRPCModule,
				true, true, false, true),
			genNewCommandTestCase("should create an gRPC project but disable actions",
				"api", testRPCProjectName, testRPCModule,
				true, true, true, false),
		}
		for _, t := range tests {
			testcase := t
			It(testcase.title, func() {
				se, err = RunCLITest(testcase.commands...)
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

	Context("Get the reflection response", func() {

	})

	Context("Ping", func() {

	})
}
