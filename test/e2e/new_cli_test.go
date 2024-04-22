package e2e_test

import (
	. "github.com/onsi/ginkgo/v2"
	"os"
)

func NewCLITest() {
	const (
		testCLIProjectName = "testcli"
		testCLIModule      = "github.com/user/testcli"
	)
	Context("New With Flag Parameters", func() {
		AfterEach(func() {
			_ = os.RemoveAll(testCLIProjectName)
		})
		tests := []NewCommandTestCase{
			genNewCommandTestCase("should create an CLI project",
				"api", testCLIProjectName, testCLIModule,
				true, true, true, true),
			genNewCommandTestCase("should create an CLI project but disable lint",
				"api", testCLIProjectName, testCLIModule,
				false, true, true, true),
			genNewCommandTestCase("should create an CLI project but disable releaser",
				"api", testCLIProjectName, testCLIModule,
				true, false, true, true),
			genNewCommandTestCase("should create an CLI project but disable semver",
				"api", testCLIProjectName, testCLIModule,
				true, true, false, true),
			genNewCommandTestCase("should create an CLI project but disable actions",
				"api", testCLIProjectName, testCLIModule,
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
		It("should create an CLI project with new command", func() {

		})
	})

	Context("Run The Sub command", func() {

	})
}
