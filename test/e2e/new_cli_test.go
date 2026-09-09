package e2e_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func NewCLITest() {
	const (
		testCLIProjectName = "testcli"
		testCLIModule      = "github.com/user/testcli"
	)
	Context("New With Flag Parameters", func() {
		var outDir string
		BeforeEach(func() {
			var err error
			outDir, err = os.MkdirTemp("", "jaguar-new-cli-*")
			Expect(err).NotTo(HaveOccurred())
		})
		AfterEach(func() {
			_ = os.RemoveAll(outDir)
		})
		tests := []NewCommandTestCase{
			genNewCommandTestCase("should create an CLI project",
				"go-cli", testCLIProjectName, testCLIModule,
				true, true, true, true),
			genNewCommandTestCase("should create an CLI project but disable lint",
				"go-cli", testCLIProjectName, testCLIModule,
				false, true, true, true),
			genNewCommandTestCase("should create an CLI project but disable releaser",
				"go-cli", testCLIProjectName, testCLIModule,
				true, false, true, true),
			genNewCommandTestCase("should create an CLI project but disable semver",
				"go-cli", testCLIProjectName, testCLIModule,
				true, true, false, true),
			genNewCommandTestCase("should create an CLI project but disable actions",
				"go-cli", testCLIProjectName, testCLIModule,
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
		It("should create an CLI project with new command", func() {

		})
	})
}
