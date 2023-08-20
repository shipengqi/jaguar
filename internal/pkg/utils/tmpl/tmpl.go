package tmpl

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

type Data struct {
	Project *Project
}

type Project struct {
	Name       string
	Module     string
	Bin        string
	BuildPkg   string
	VersionPkg string
}

func ReplaceFile(fpath string, data *Data) {
	var buf bytes.Buffer
	text, _ := os.ReadFile(fpath)
	t, _ := template.New("").Parse(string(text))
	_ = t.Execute(&buf, data)
	fmt.Println(buf.String())
}

func ReplaceModule(fpath, m1, m2 string) {
	text, _ := os.ReadFile(fpath)
	tt := bytes.ReplaceAll(text, []byte(m1), []byte(m2))
	fmt.Println(string(tt))
}
