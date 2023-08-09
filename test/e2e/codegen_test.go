package e2e_test

import (
	"fmt"
	"os"

	. "github.com/onsi/ginkgo/v2"
)

const (
	_defaultGenGoName  = `codes_generated.go`
	_defaultGenDocName = `codes_generated.md`
	_testDocName       = "codes.md"
)

var (
	_defaultGenGoFilePath  = fmt.Sprintf("./testdata/codes/%s", _defaultGenGoName)
	_defaultGenDocFilePath = fmt.Sprintf("./testdata/codes/%s", _defaultGenDocName)
	_testFilePath          = fmt.Sprintf("./testdata/codes/%s", _testDocName)
)

func CodeGenTest() {
	AfterEach(func() {
		_ = os.Remove(_defaultGenGoFilePath)
		_ = os.Remove(_defaultGenDocFilePath)
		_ = os.Remove(_testFilePath)
	})

	Context("Generate register go file", func() {
		It("should generate go file", func() {
			se, err = RunCLITest("codegen", "--types", "int", "./testdata/codes")
			NoError(err)

			ShouldExists(_defaultGenGoFilePath)
			ShouldNotExists(_defaultGenDocFilePath)
		})

		It("should fail without --types", func() {
			se, err = RunCLITest("codegen", "./testdata/codes")
			NoError(err)
			ExitCode(se, 1)
			ShouldContains(se, "--types is required")
		})

		It("should fail with empty --types", func() {
			se, err = RunCLITest("codegen", "--types", "unknowntype")
			NoError(err)
			ExitCode(se, 1)
			ShouldContains(se, "no values defined for type unknowntype")
		})
	})

	Context("Generate error codes markdown", func() {
		It("should generate markdown with default filename", func() {
			se, err = RunCLITest("codegen", "--types", "int", "--doc", "./testdata/codes")
			NoError(err)

			ShouldExists(_defaultGenDocFilePath)
			ShouldNotExists(_defaultGenGoFilePath)
		})

		It("should generate markdown with given filename", func() {
			se, err = RunCLITest("codegen", "--types", "int", "--doc", "--output", _testFilePath, "./testdata/codes")
			NoError(err)

			ShouldExists(_testFilePath)
			ShouldNotExists(_defaultGenGoFilePath)
		})
	})
}
