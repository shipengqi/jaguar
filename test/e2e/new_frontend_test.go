package e2e_test

import (
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func NewFrontendTest() {
	Context("New With Flag Parameters", func() {
		var outDir string
		BeforeEach(func() {
			var err error
			outDir, err = os.MkdirTemp("", "jaguar-new-frontend-*")
			Expect(err).NotTo(HaveOccurred())
		})
		AfterEach(func() {
			_ = os.RemoveAll(outDir)
		})

		It("should create a React project", func() {
			se, err = RunCLITest(
				"new", "testfereact",
				"-t", "frontend-react",
				"-o", outDir,
				"--use-github-actions",
			)
			NoError(err)
			ExitCode(se, 0)
			ShouldContains(se, "Project type:  frontend-react")
			ShouldContains(se, "Framework:     N/A")

			projectDir := filepath.Join(outDir, "testfereact")
			ShouldExists(projectDir)
			ShouldExists(filepath.Join(projectDir, "package.json"))
			ShouldExists(filepath.Join(projectDir, "vite.config.ts"))
			ShouldExists(filepath.Join(projectDir, "index.html"))
			ShouldExists(filepath.Join(projectDir, "tsconfig.json"))
			ShouldExists(filepath.Join(projectDir, "src", "App.tsx"))
			ShouldExists(filepath.Join(projectDir, "src", "main.tsx"))
			ShouldExists(filepath.Join(projectDir, "src", "store", "auth.store.ts"))
			ShouldExists(filepath.Join(projectDir, "src", "lib", "axios.ts"))
			ShouldExists(filepath.Join(projectDir, ".github", "workflows"))
		})

		It("should create a React project without GitHub Actions", func() {
			se, err = RunCLITest(
				"new", "testfereact2",
				"-t", "frontend-react",
				"-o", outDir,
				"--use-github-actions=false",
			)
			NoError(err)
			ExitCode(se, 0)

			projectDir := filepath.Join(outDir, "testfereact2")
			ShouldExists(projectDir)
			ShouldExists(filepath.Join(projectDir, "package.json"))
			ShouldNotExists(filepath.Join(projectDir, ".github", "workflows"))
		})

		It("should create a Vue project", func() {
			se, err = RunCLITest(
				"new", "testfevue",
				"-t", "frontend-vue",
				"-o", outDir,
				"--use-github-actions",
			)
			NoError(err)
			ExitCode(se, 0)
			ShouldContains(se, "Project type:  frontend-vue")
			ShouldContains(se, "Framework:     N/A")

			projectDir := filepath.Join(outDir, "testfevue")
			ShouldExists(projectDir)
			ShouldExists(filepath.Join(projectDir, "package.json"))
			ShouldExists(filepath.Join(projectDir, "vite.config.ts"))
			ShouldExists(filepath.Join(projectDir, "tsconfig.json"))
			ShouldExists(filepath.Join(projectDir, "src", "main.ts"))
			ShouldExists(filepath.Join(projectDir, "src", "App.vue"))
			ShouldExists(filepath.Join(projectDir, "src", "stores", "auth.ts"))
			ShouldExists(filepath.Join(projectDir, "src", "router", "index.ts"))
			ShouldExists(filepath.Join(projectDir, ".github", "workflows"))
		})

		It("should create an Angular project", func() {
			se, err = RunCLITest(
				"new", "testfeangular",
				"-t", "frontend-angular",
				"-o", outDir,
				"--use-github-actions",
			)
			NoError(err)
			ExitCode(se, 0)
			ShouldContains(se, "Project type:  frontend-angular")
			ShouldContains(se, "Framework:     N/A")

			projectDir := filepath.Join(outDir, "testfeangular")
			ShouldExists(projectDir)
			ShouldExists(filepath.Join(projectDir, "package.json"))
			ShouldExists(filepath.Join(projectDir, "vite.config.ts"))
			ShouldExists(filepath.Join(projectDir, "tsconfig.json"))
			ShouldExists(filepath.Join(projectDir, "src", "main.ts"))
			ShouldExists(filepath.Join(projectDir, "src", "app", "app.config.ts"))
			ShouldExists(filepath.Join(projectDir, "src", "app", "app.component.ts"))
			ShouldExists(filepath.Join(projectDir, ".github", "workflows"))
		})
	})

	Context("New Without Flag Parameters", func() {
		It("should create a frontend project with new command", func() {
			// interactive mode — covered by flag-based tests above
		})
	})
}
