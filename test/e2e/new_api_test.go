package e2e_test

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	"os"
	"strconv"
	"strings"
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
		AfterEach(func() {
			_ = os.RemoveAll(testAPIProjectName)
		})
		tests := []NewCommandTestCase{
			genNewCommandTestCase("should create an API project",
				"api", testAPIProjectName, testAPIModule,
				true, true, true, true),
			genNewCommandTestCase("should create an API project but disable lint",
				"api", testAPIProjectName, testAPIModule,
				false, true, true, true),
			genNewCommandTestCase("should create an API project but disable releaser",
				"api", testAPIProjectName, testAPIModule,
				true, false, true, true),
			genNewCommandTestCase("should create an API project but disable semver",
				"api", testAPIProjectName, testAPIModule,
				true, true, false, true),
			genNewCommandTestCase("should create an API project but disable actions",
				"api", testAPIProjectName, testAPIModule,
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
		It("should create an API project with new command", func() {

		})
	})

	Context("Generate API error codes", func() {

	})

	Context("Open the swagger document", func() {

	})

	Context("Ping", func() {

	})
}

func genNewCommandTestCase(title, t, n, m string, lint, release, semver, actions bool) NewCommandTestCase {
	pt := strings.ToUpper(t)
	if t == "gprc" {
		pt = "gRPC"
	}
	expects := []string{
		fmt.Sprintf("Application type: %s", pt),
	}
	framework := "Gin"
	if t != "api" {
		framework = "N/A"
	}
	expects = append(expects, fmt.Sprintf("Go web framework: %s", framework))
	expects = append(expects, fmt.Sprintf("Is use golangci-lint to lint your Go code? %s", strconv.FormatBool(lint)))
	expects = append(expects, fmt.Sprintf("Is use GoReleaser to deliver your Go binaries? %s", strconv.FormatBool(release)))
	expects = append(expects, fmt.Sprintf("Is use GSemver to generate your next semver version? %s", strconv.FormatBool(semver)))
	expects = append(expects,
		fmt.Sprintf("Is use the GitHub Actions to automate your build, test, and deployment pipeline? %s", strconv.FormatBool(actions)))

	commands := []string{
		"new", "-t",
	}
	commands = append(commands, t, "-n", n, "-m", m)

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
