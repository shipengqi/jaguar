package e2e_test

import (
	. "github.com/onsi/ginkgo/v2"
)

func LicenseTest() {
	Context("Check License Header", func() {
		It("should output the files without license", func() {
			se, err = RunCLITest("license", "check", "./testdata/license")
			NoError(err)
			ShouldContains(se, "without-license\\test.go: missing license header")
			ShouldContains(se, "without-license\\test.html: missing license header")
			ShouldContains(se, "without-license\\test.css: missing license header")
			ShouldContains(se, "without-license\\test.sh: missing license header")
			ShouldContains(se, "without-license\\test.ts: missing license header")
			ShouldContains(se, "without-license\\test.yaml: missing license header")
			ShouldContains(se, "skip.css: missing license header")
			ShouldContains(se, "skip-dir\\test.css: missing license header")
			ShouldNotContains(se, "with-license\\")
		})

		It("should skip the given directory and file", func() {
			se, err = RunCLITest("license", "check", "./testdata/license",
				"--skip-files", "skip.css", "--skip-dirs", "skip-dir")
			NoError(err)
			ShouldNotContains(se, "with-license\\")
			ShouldNotContains(se, "skip.css")
			ShouldNotContains(se, "skip-dir")
		})
	})

	Context("Add License Header", func() {
		It("should write the license into files", func() {
			se, err = RunCLITest("license", "add", "./testdata/license/without-license")
			NoError(err)
			ShouldContains(se, "without-license\\test.go: license added")
			ShouldContains(se, "without-license\\test.html: license added")
			ShouldContains(se, "without-license\\test.css: license added")
			ShouldContains(se, "without-license\\test.sh: license added")
			ShouldContains(se, "without-license\\test.ts: license added")
			ShouldContains(se, "without-license\\test.yaml: license added")

			se, err = RunCLITest("license", "check", "./testdata/license",
				"--skip-files", "skip.css", "--skip-dirs", "skip-dir")
			NoError(err)
			ShouldNotContains(se, "without-license\\test.go: license added")
			ShouldNotContains(se, "without-license\\test.html: license added")
			ShouldNotContains(se, "without-license\\test.css: license added")
			ShouldNotContains(se, "without-license\\test.sh: license added")
			ShouldNotContains(se, "without-license\\test.ts: license added")
			ShouldNotContains(se, "without-license\\test.yaml: license added")
		})
	})

	Context("Remove License Header", func() {
		It("should remove the license from files", func() {
			se, err = RunCLITest("license", "remove", "./testdata/license/without-license")
			NoError(err)
			ShouldContains(se, "without-license\\test.go: license removed")
			ShouldContains(se, "without-license\\test.html: license removed")
			ShouldContains(se, "without-license\\test.css: license removed")
			ShouldContains(se, "without-license\\test.sh: license removed")
			ShouldContains(se, "without-license\\test.ts: license removed")
			ShouldContains(se, "without-license\\test.yaml: license removed")

			se, err = RunCLITest("license", "check", "./testdata/license",
				"--skip-files", "skip.css", "--skip-dirs", "skip-dir")
			NoError(err)
			ShouldContains(se, "without-license\\test.go: missing license header")
			ShouldContains(se, "without-license\\test.html: missing license header")
			ShouldContains(se, "without-license\\test.css: missing license header")
			ShouldContains(se, "without-license\\test.sh: missing license header")
			ShouldContains(se, "without-license\\test.ts: missing license header")
			ShouldContains(se, "without-license\\test.yaml: missing license header")
		})
	})
}
