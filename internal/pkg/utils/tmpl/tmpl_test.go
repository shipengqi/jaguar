package tmpl

import (
	"os"
	"testing"
)

func TestReplaceFile(t *testing.T) {
	p := &Project{
		Name:       "testname",
		Module:     "github.com/test/module",
		Bin:        "testbin",
		BuildPkg:   "testbuildpkg",
		VersionPkg: "testversionpkg",
	}

	data := &Data{Project: p}

	dirs, _ := os.ReadDir("./testdata")
	for _, v := range dirs {
		t.Run(v.Name(), func(t *testing.T) {
			ReplaceFile("./testdata/"+v.Name(), data)
		})
	}
}

func TestReplaceModule(t *testing.T) {
	ReplaceModule("./testdata/tmpl_test.go", "github.com/shipengqi/jaguar", "github.com/test/module")
}
