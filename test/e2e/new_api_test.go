package e2e_test

import (
	"fmt"
	"os"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type NewCommandTestCase struct {
	title    string
	commands []string
	expects  []string
}

func NewAPITest() {
	const (
		testAPIProjectName = "testapi"
		testAPIModule      = "github.com/user/testapi"
	)
	Context("New With Flag Parameters", func() {
		var outDir string
		BeforeEach(func() {
			var err error
			outDir, err = os.MkdirTemp("", "jaguar-new-api-*")
			Expect(err).NotTo(HaveOccurred())
		})
		AfterEach(func() {
			_ = os.RemoveAll(outDir)
		})
		tests := []NewCommandTestCase{
			genNewCommandTestCase("should create an API project",
				"go-api", testAPIProjectName, testAPIModule,
				true, true, true, true),
			genNewCommandTestCase("should create an API project but disable lint",
				"go-api", testAPIProjectName, testAPIModule,
				false, true, true, true),
			genNewCommandTestCase("should create an API project but disable releaser",
				"go-api", testAPIProjectName, testAPIModule,
				true, false, true, true),
			genNewCommandTestCase("should create an API project but disable semver",
				"go-api", testAPIProjectName, testAPIModule,
				true, true, false, true),
			genNewCommandTestCase("should create an API project but disable actions",
				"go-api", testAPIProjectName, testAPIModule,
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
		It("should create an API project with new command", func() {

		})
	})
}

func genNewCommandTestCase(title, t, n, m string, lint, release, semver, actions bool) NewCommandTestCase {
	expects := []string{
		fmt.Sprintf("Project type:  %s", t),
	}
	framework := "N/A"
	if t == "go-api" || t == "go-embed" {
		framework = "gin"
	}
	expects = append(expects, fmt.Sprintf("Framework:     %s", framework))
	expects = append(expects, fmt.Sprintf("golangci-lint: %s", strconv.FormatBool(lint)))
	expects = append(expects, fmt.Sprintf("GoReleaser:    %s", strconv.FormatBool(release)))
	expects = append(expects, fmt.Sprintf("GSemver:       %s", strconv.FormatBool(semver)))
	expects = append(expects, fmt.Sprintf("GitHub Actions:%s", strconv.FormatBool(actions)))

	commands := []string{"new", n, "-t", t, "-m", m}

	if !lint {
		commands = append(commands, "--use-golangci-lint=false")
	} else {
		commands = append(commands, "--use-golangci-lint")
	}
	if !release {
		commands = append(commands, "--use-goreleaser=false")
	} else {
		commands = append(commands, "--use-goreleaser")
	}
	if !semver {
		commands = append(commands, "--use-gsemver=false")
	} else {
		commands = append(commands, "--use-gsemver")
	}
	if !actions {
		commands = append(commands, "--use-github-actions=false")
	} else {
		commands = append(commands, "--use-github-actions")
	}

	return NewCommandTestCase{
		title:    title,
		commands: commands,
		expects:  expects,
	}
}
