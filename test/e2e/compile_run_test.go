package e2e_test

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/shipengqi/jaguar/test/e2e"
)

// goCmd creates a go tool invocation with GOPROXY set to bypass corporate proxy restrictions.
func goCmd(dir string, args ...string) *exec.Cmd {
	cmd := exec.Command("go", args...)
	cmd.Dir = dir
	cmd.Stdout = GinkgoWriter
	cmd.Stderr = GinkgoWriter
	cmd.Env = append(os.Environ(), "GOPROXY=https://goproxy.cn,direct")
	return cmd
}

func CompileAndRunTest() {
	Describe("go-api/gin compile and run", Ordered, func() {
		const (
			projectName = "testapi-run"
			moduleName  = "github.com/e2etest/testapi"
			bindPort    = "18080"
		)

		var (
			workDir    string
			projectDir string
		)

		BeforeAll(func() {
			var err error
			workDir, err = os.MkdirTemp("", "jaguar-e2e-*")
			Expect(err).NotTo(HaveOccurred())
			projectDir = filepath.Join(workDir, projectName)
		})

		AfterAll(func() {
			_ = os.RemoveAll(workDir)
		})

		It("generates the project", func() {
			cmd := exec.Command(CliOpts.Cli,
				"new", projectName,
				"-t", "go-api",
				"-m", moduleName,
				"-o", workDir,
				"--use-golangci-lint=false",
				"--use-goreleaser=false",
				"--use-gsemver=false",
				"--use-github-actions=false",
			)
			cmd.Dir = workDir
			cmd.Stdout = GinkgoWriter
			cmd.Stderr = GinkgoWriter
			Expect(cmd.Run()).To(Succeed())
			Expect(projectDir).To(BeADirectory())
		})

		It("compiles the project", func() {
			Expect(goCmd(projectDir, "mod", "tidy").Run()).To(Succeed())
			Expect(goCmd(projectDir, "build", "./...").Run()).To(Succeed())
		})

		It("runs the server and pings /healthz", func() {
			cfg := fmt.Sprintf(`server:
  healthz: true
  mode: debug
  middlewares: recovery,requestid,logger

insecure:
  bind-address: 127.0.0.1
  bind-port: %s

secure:
  bind-port: 0
`, bindPort)

			cfgFile := filepath.Join(projectDir, "configs", "apiserver.yaml")
			Expect(os.WriteFile(cfgFile, []byte(cfg), 0o644)).To(Succeed())

			binPath := filepath.Join(projectDir, "bin", "apiserver")
			Expect(os.MkdirAll(filepath.Dir(binPath), 0o755)).To(Succeed())
			Expect(goCmd(projectDir, "build", "-o", binPath, "./cmd/apiserver/").Run()).To(Succeed())

			srv := exec.Command(binPath, "-c", cfgFile)
			srv.Dir = projectDir
			srv.Stdout = GinkgoWriter
			srv.Stderr = GinkgoWriter
			Expect(srv.Start()).To(Succeed())
			DeferCleanup(func() { _ = srv.Process.Kill() })

			healthURL := fmt.Sprintf("http://127.0.0.1:%s/healthz", bindPort)
			Eventually(func() int {
				resp, err := http.Get(healthURL) //nolint:noctx
				if err != nil {
					return 0
				}
				defer resp.Body.Close()
				return resp.StatusCode
			}, 30*time.Second, time.Second).Should(Equal(http.StatusOK))
		})
	})

	Describe("go-cli compile", Ordered, func() {
		const (
			projectName = "testcli-run"
			moduleName  = "github.com/e2etest/testcli"
		)

		var (
			workDir    string
			projectDir string
		)

		BeforeAll(func() {
			var err error
			workDir, err = os.MkdirTemp("", "jaguar-e2e-*")
			Expect(err).NotTo(HaveOccurred())
			projectDir = filepath.Join(workDir, projectName)
		})

		AfterAll(func() {
			_ = os.RemoveAll(workDir)
		})

		It("generates the project", func() {
			cmd := exec.Command(CliOpts.Cli,
				"new", projectName,
				"-t", "go-cli",
				"-m", moduleName,
				"-o", workDir,
				"--use-golangci-lint=false",
				"--use-goreleaser=false",
				"--use-gsemver=false",
				"--use-github-actions=false",
			)
			cmd.Dir = workDir
			cmd.Stdout = GinkgoWriter
			cmd.Stderr = GinkgoWriter
			Expect(cmd.Run()).To(Succeed())
			Expect(projectDir).To(BeADirectory())
		})

		It("compiles and runs --help", func() {
			Expect(goCmd(projectDir, "mod", "tidy").Run()).To(Succeed())

			binPath := filepath.Join(projectDir, "bin", "examplecli")
			Expect(os.MkdirAll(filepath.Dir(binPath), 0o755)).To(Succeed())
			Expect(goCmd(projectDir, "build", "-o", binPath, "./cmd/examplecli/").Run()).To(Succeed())

			out, err := exec.Command(binPath, "--help").CombinedOutput()
			Expect(err).NotTo(HaveOccurred())
			Expect(string(out)).To(ContainSubstring("hello"))
		})
	})

	Describe("go-embed/gin compile and run", Ordered, func() {
		const (
			projectName = "testembed-run"
			moduleName  = "github.com/e2etest/testembed"
			bindPort    = "18081"
		)

		var (
			workDir    string
			projectDir string
		)

		BeforeAll(func() {
			var err error
			workDir, err = os.MkdirTemp("", "jaguar-e2e-*")
			Expect(err).NotTo(HaveOccurred())
			projectDir = filepath.Join(workDir, projectName)
		})

		AfterAll(func() {
			_ = os.RemoveAll(workDir)
		})

		It("generates the project", func() {
			cmd := exec.Command(CliOpts.Cli,
				"new", projectName,
				"-t", "go-embed",
				"-m", moduleName,
				"-o", workDir,
				"--frontend-framework", "react",
				"--use-golangci-lint=false",
				"--use-goreleaser=false",
				"--use-gsemver=false",
				"--use-github-actions=false",
			)
			cmd.Dir = workDir
			cmd.Stdout = GinkgoWriter
			cmd.Stderr = GinkgoWriter
			Expect(cmd.Run()).To(Succeed())
			Expect(projectDir).To(BeADirectory())
		})

		It("builds the frontend", func() {
			webDir := filepath.Join(projectDir, "web")

			npmInstall := exec.Command("npm", "install")
			npmInstall.Dir = webDir
			npmInstall.Stdout = GinkgoWriter
			npmInstall.Stderr = GinkgoWriter
			Expect(npmInstall.Run()).To(Succeed())

			npmBuild := exec.Command("npm", "run", "build")
			npmBuild.Dir = webDir
			npmBuild.Stdout = GinkgoWriter
			npmBuild.Stderr = GinkgoWriter
			Expect(npmBuild.Run()).To(Succeed())

			// vite outputs to internal/staticfs/web
			Expect(filepath.Join(projectDir, "internal", "staticfs", "web", "index.html")).To(BeAnExistingFile())
		})

		It("compiles the project", func() {
			Expect(goCmd(projectDir, "mod", "tidy").Run()).To(Succeed())
			Expect(goCmd(projectDir, "build", "./...").Run()).To(Succeed())
		})

		It("runs the server and serves the frontend at /", func() {
			cfg := fmt.Sprintf(`server:
  healthz: true
  mode: debug
  middlewares: recovery,requestid,logger

insecure:
  bind-address: 127.0.0.1
  bind-port: %s

secure:
  bind-port: 0
`, bindPort)

			cfgFile := filepath.Join(projectDir, "configs", "apiserver.yaml")
			Expect(os.WriteFile(cfgFile, []byte(cfg), 0o644)).To(Succeed())

			binPath := filepath.Join(projectDir, "bin", "apiserver")
			Expect(os.MkdirAll(filepath.Dir(binPath), 0o755)).To(Succeed())
			Expect(goCmd(projectDir, "build", "-o", binPath, "./cmd/apiserver/").Run()).To(Succeed())

			srv := exec.Command(binPath, "-c", cfgFile)
			srv.Dir = projectDir
			srv.Stdout = GinkgoWriter
			srv.Stderr = GinkgoWriter
			Expect(srv.Start()).To(Succeed())
			DeferCleanup(func() { _ = srv.Process.Kill() })

			// wait for healthz first
			healthURL := fmt.Sprintf("http://127.0.0.1:%s/healthz", bindPort)
			Eventually(func() int {
				resp, err := http.Get(healthURL) //nolint:noctx
				if err != nil {
					return 0
				}
				defer resp.Body.Close()
				return resp.StatusCode
			}, 30*time.Second, time.Second).Should(Equal(http.StatusOK))

			// then verify the frontend is served at /
			resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s/", bindPort)) //nolint:noctx
			Expect(err).NotTo(HaveOccurred())
			defer resp.Body.Close()
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
		})
	})

	Describe("go-embed/gin+vue scaffold", Ordered, func() {
		const (
			projectName = "testembed-vue"
			moduleName  = "github.com/e2etest/testembedvue"
		)

		var (
			workDir    string
			projectDir string
		)

		BeforeAll(func() {
			var err error
			workDir, err = os.MkdirTemp("", "jaguar-e2e-*")
			Expect(err).NotTo(HaveOccurred())
			projectDir = filepath.Join(workDir, projectName)
		})

		AfterAll(func() {
			_ = os.RemoveAll(workDir)
		})

		It("generates the project with vue frontend", func() {
			cmd := exec.Command(CliOpts.Cli,
				"new", projectName,
				"-t", "go-embed",
				"-m", moduleName,
				"-o", workDir,
				"--frontend-framework", "vue",
				"--use-golangci-lint=false",
				"--use-goreleaser=false",
				"--use-gsemver=false",
				"--use-github-actions=false",
			)
			cmd.Dir = workDir
			cmd.Stdout = GinkgoWriter
			cmd.Stderr = GinkgoWriter
			Expect(cmd.Run()).To(Succeed())
			Expect(projectDir).To(BeADirectory())
		})

		It("has the vue frontend files under web/", func() {
			webDir := filepath.Join(projectDir, "web")
			Expect(webDir).To(BeADirectory())
			Expect(filepath.Join(webDir, "package.json")).To(BeAnExistingFile())
			Expect(filepath.Join(webDir, "vite.config.ts")).To(BeAnExistingFile())
			Expect(filepath.Join(webDir, "src", "main.ts")).To(BeAnExistingFile())
			Expect(filepath.Join(webDir, "src", "App.vue")).To(BeAnExistingFile())
		})

		It("compiles the Go project", func() {
			Expect(goCmd(projectDir, "mod", "tidy").Run()).To(Succeed())
			Expect(goCmd(projectDir, "build", "./...").Run()).To(Succeed())
		})
	})

	Describe("go-embed/gin+angular scaffold", Ordered, func() {
		const (
			projectName = "testembed-angular"
			moduleName  = "github.com/e2etest/testembedangular"
		)

		var (
			workDir    string
			projectDir string
		)

		BeforeAll(func() {
			var err error
			workDir, err = os.MkdirTemp("", "jaguar-e2e-*")
			Expect(err).NotTo(HaveOccurred())
			projectDir = filepath.Join(workDir, projectName)
		})

		AfterAll(func() {
			_ = os.RemoveAll(workDir)
		})

		It("generates the project with angular frontend", func() {
			cmd := exec.Command(CliOpts.Cli,
				"new", projectName,
				"-t", "go-embed",
				"-m", moduleName,
				"-o", workDir,
				"--frontend-framework", "angular",
				"--use-golangci-lint=false",
				"--use-goreleaser=false",
				"--use-gsemver=false",
				"--use-github-actions=false",
			)
			cmd.Dir = workDir
			cmd.Stdout = GinkgoWriter
			cmd.Stderr = GinkgoWriter
			Expect(cmd.Run()).To(Succeed())
			Expect(projectDir).To(BeADirectory())
		})

		It("has the angular frontend files under web/", func() {
			webDir := filepath.Join(projectDir, "web")
			Expect(webDir).To(BeADirectory())
			Expect(filepath.Join(webDir, "package.json")).To(BeAnExistingFile())
			Expect(filepath.Join(webDir, "vite.config.ts")).To(BeAnExistingFile())
			Expect(filepath.Join(webDir, "src", "main.ts")).To(BeAnExistingFile())
			Expect(filepath.Join(webDir, "src", "app", "app.config.ts")).To(BeAnExistingFile())
		})

		It("compiles the Go project", func() {
			Expect(goCmd(projectDir, "mod", "tidy").Run()).To(Succeed())
			Expect(goCmd(projectDir, "build", "./...").Run()).To(Succeed())
		})
	})
}
