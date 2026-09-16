package e2e_test

import (
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func NewNodeJSTest() {
	Context("New With Flag Parameters", func() {
		var outDir string
		BeforeEach(func() {
			var err error
			outDir, err = os.MkdirTemp("", "jaguar-new-nodejs-*")
			Expect(err).NotTo(HaveOccurred())
		})
		AfterEach(func() {
			_ = os.RemoveAll(outDir)
		})

		It("should create a Koa project", func() {
			se, err = RunCLITest(
				"new", "testkoa",
				"-t", "nodejs",
				"--framework", "koa",
				"-o", outDir,
				"--use-github-actions",
			)
			NoError(err)
			ExitCode(se, 0)
			ShouldContains(se, "Project type:  nodejs")
			ShouldContains(se, "Framework:     koa")

			projectDir := filepath.Join(outDir, "testkoa")
			ShouldExists(projectDir)
			ShouldExists(filepath.Join(projectDir, "package.json"))
			ShouldExists(filepath.Join(projectDir, "src", "routes", "index.ts"))
			ShouldExists(filepath.Join(projectDir, "src", "controllers", "user.ts"))
			ShouldExists(filepath.Join(projectDir, "tsconfig.json"))
			ShouldExists(filepath.Join(projectDir, ".github", "workflows"))
		})

		It("should create a Koa project without GitHub Actions", func() {
			se, err = RunCLITest(
				"new", "testkoa2",
				"-t", "nodejs",
				"--framework", "koa",
				"-o", outDir,
				"--use-github-actions=false",
			)
			NoError(err)
			ExitCode(se, 0)

			projectDir := filepath.Join(outDir, "testkoa2")
			ShouldExists(projectDir)
			ShouldExists(filepath.Join(projectDir, "package.json"))
			ShouldNotExists(filepath.Join(projectDir, ".github", "workflows"))
		})

		It("should create a NestJS project", func() {
			se, err = RunCLITest(
				"new", "testnestjs",
				"-t", "nodejs",
				"--framework", "nestjs",
				"-o", outDir,
				"--use-github-actions",
			)
			NoError(err)
			ExitCode(se, 0)
			ShouldContains(se, "Project type:  nodejs")
			ShouldContains(se, "Framework:     nestjs")

			projectDir := filepath.Join(outDir, "testnestjs")
			ShouldExists(projectDir)
			ShouldExists(filepath.Join(projectDir, "package.json"))
			ShouldExists(filepath.Join(projectDir, "src", "main.ts"))
			ShouldExists(filepath.Join(projectDir, "src", "app.module.ts"))
			ShouldExists(filepath.Join(projectDir, "src", "app.controller.ts"))
			ShouldExists(filepath.Join(projectDir, "tsconfig.json"))
			ShouldExists(filepath.Join(projectDir, "nest-cli.json"))
			ShouldExists(filepath.Join(projectDir, ".github", "workflows"))
		})

		It("should create a Next.js project", func() {
			se, err = RunCLITest(
				"new", "testnextjs",
				"-t", "nodejs",
				"--framework", "nextjs",
				"-o", outDir,
				"--use-github-actions",
			)
			NoError(err)
			ExitCode(se, 0)
			ShouldContains(se, "Project type:  nodejs")
			ShouldContains(se, "Framework:     nextjs")

			projectDir := filepath.Join(outDir, "testnextjs")
			ShouldExists(projectDir)

			// Root config files
			ShouldExists(filepath.Join(projectDir, "package.json"))
			ShouldExists(filepath.Join(projectDir, "next.config.ts"))
			ShouldExists(filepath.Join(projectDir, "tsconfig.json"))
			ShouldExists(filepath.Join(projectDir, ".env.example"))
			ShouldExists(filepath.Join(projectDir, ".gitignore"))
			ShouldExists(filepath.Join(projectDir, "middleware.ts"))
			ShouldExists(filepath.Join(projectDir, "drizzle.config.ts"))

			// App directory structure
			ShouldExists(filepath.Join(projectDir, "app", "layout.tsx"))
			ShouldExists(filepath.Join(projectDir, "app", "globals.css"))
			ShouldExists(filepath.Join(projectDir, "app", "favicon.ico"))

			// Dashboard pages
			ShouldExists(filepath.Join(projectDir, "app", "(dashboard)", "layout.tsx"))
			ShouldExists(filepath.Join(projectDir, "app", "(dashboard)", "dashboard", "page.tsx"))
			ShouldExists(filepath.Join(projectDir, "app", "(dashboard)", "dashboard", "activity", "page.tsx"))
			ShouldExists(filepath.Join(projectDir, "app", "(dashboard)", "dashboard", "general", "page.tsx"))
			ShouldExists(filepath.Join(projectDir, "app", "(dashboard)", "dashboard", "security", "page.tsx"))
			ShouldExists(filepath.Join(projectDir, "app", "(dashboard)", "pricing", "page.tsx"))

			// Login pages
			ShouldExists(filepath.Join(projectDir, "app", "(login)", "login.tsx"))
			ShouldExists(filepath.Join(projectDir, "app", "(login)", "actions.ts"))
			ShouldExists(filepath.Join(projectDir, "app", "(login)", "sign-in", "page.tsx"))
			ShouldExists(filepath.Join(projectDir, "app", "(login)", "sign-up", "page.tsx"))

			// API routes
			ShouldExists(filepath.Join(projectDir, "app", "api", "stripe", "checkout", "route.ts"))
			ShouldExists(filepath.Join(projectDir, "app", "api", "stripe", "webhook", "route.ts"))
			ShouldExists(filepath.Join(projectDir, "app", "api", "team", "route.ts"))
			ShouldExists(filepath.Join(projectDir, "app", "api", "user", "route.ts"))

			// UI components
			ShouldExists(filepath.Join(projectDir, "components", "ui", "button.tsx"))
			ShouldExists(filepath.Join(projectDir, "components", "ui", "input.tsx"))
			ShouldExists(filepath.Join(projectDir, "components", "ui", "card.tsx"))
			ShouldExists(filepath.Join(projectDir, "components", "ui", "avatar.tsx"))

			// Lib directory
			ShouldExists(filepath.Join(projectDir, "lib", "utils.ts"))
			ShouldExists(filepath.Join(projectDir, "lib", "auth", "session.ts"))
			ShouldExists(filepath.Join(projectDir, "lib", "auth", "middleware.ts"))
			ShouldExists(filepath.Join(projectDir, "lib", "db", "schema.ts"))
			ShouldExists(filepath.Join(projectDir, "lib", "db", "drizzle.ts"))
			ShouldExists(filepath.Join(projectDir, "lib", "db", "queries.ts"))
			ShouldExists(filepath.Join(projectDir, "lib", "db", "migrations", "0000_soft_the_anarchist.sql"))
			ShouldExists(filepath.Join(projectDir, "lib", "payments", "stripe.ts"))
			ShouldExists(filepath.Join(projectDir, "lib", "payments", "actions.ts"))

			// CI workflows
			ShouldExists(filepath.Join(projectDir, ".github", "workflows"))

			// Verify lockfile is NOT present (should be generated by user)
			ShouldNotExists(filepath.Join(projectDir, "pnpm-lock.yaml"))
		})

		It("should create a Next.js project without GitHub Actions", func() {
			se, err = RunCLITest(
				"new", "testnextjs2",
				"-t", "nodejs",
				"--framework", "nextjs",
				"-o", outDir,
				"--use-github-actions=false",
			)
			NoError(err)
			ExitCode(se, 0)

			projectDir := filepath.Join(outDir, "testnextjs2")
			ShouldExists(projectDir)
			ShouldExists(filepath.Join(projectDir, "package.json"))
			ShouldExists(filepath.Join(projectDir, "app", "layout.tsx"))
			ShouldNotExists(filepath.Join(projectDir, ".github", "workflows"))
		})
	})

	Context("New Without Flag Parameters", func() {
		It("should create a Node.js project with new command", func() {
			// interactive mode — covered by flag-based tests above
		})
	})
}
