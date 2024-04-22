package license

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"unicode"

	"github.com/shipengqi/golib/convutil"
	"github.com/shipengqi/log"
)

type copyrightInfo struct {
	Year   string
	Holder string
}

type file struct {
	path string
	mode os.FileMode
}

var heads = []string{
	"#!",                       // shell script
	"<?xml",                    // XML declaration
	"<!doctype",                // HTML doctype
	"# encoding:",              // Ruby encoding
	"# frozen_string_literal:", // Ruby interpreter instruction
	"<?php",                    // PHP opening tag
}

func walk(ch chan<- *file, start string, dirRegs, fileRegs []*regexp.Regexp) {
	_ = filepath.Walk(start, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			log.Debugf("%s error: %v", path, err)
			return nil
		}
		if fi.IsDir() {
			for _, pattern := range dirRegs {
				if pattern.MatchString(fi.Name()) {
					return filepath.SkipDir
				}
			}
			return nil
		}

		for _, pattern := range fileRegs {
			if pattern.MatchString(fi.Name()) {
				return nil
			}
		}

		ch <- &file{path, fi.Mode()}

		return nil
	})
}

func hashBang(b []byte) []byte {
	line := make([]byte, 0, len(b))
	for _, c := range b {
		line = append(line, c)
		if c == '\n' {
			break
		}
	}
	first := bytes.ToLower(line)
	for _, h := range heads {
		if bytes.HasPrefix(first, convutil.S2B(h)) {
			return line
		}
	}

	return nil
}

func addLicense(path string, fmode os.FileMode, tmpl *template.Template, data *copyrightInfo) (bool, error) {
	var lic []byte
	var err error
	lic, err = licenseHeader(path, tmpl, data)
	if err != nil || lic == nil {
		return false, err
	}

	b, err := os.ReadFile(path)
	if err != nil || hasLicense(b) {
		return false, err
	}

	line := hashBang(b)
	if len(line) > 0 {
		b = b[len(line):]
		if line[len(line)-1] != '\n' {
			line = append(line, '\n')
		}
		line = append(line, '\n')
		lic = append(line, lic...)
	}
	b = append(lic, b...)

	return true, os.WriteFile(path, b, fmode)
}

// fileHasLicense reports whether the file at path contains a license header.
func fileHasLicense(path string) (bool, error) {
	b, err := os.ReadFile(path)
	if err != nil || hasLicense(b) {
		return false, err
	}

	return true, nil
}

func hasLicense(b []byte) bool {
	n := 1000
	if len(b) < 1000 {
		n = len(b)
	}

	return bytes.Contains(bytes.ToLower(b[:n]), convutil.S2B("copyright")) ||
		bytes.Contains(bytes.ToLower(b[:n]), convutil.S2B("mozilla public"))
}

func licenseHeader(path string, tmpl *template.Template, data *copyrightInfo) ([]byte, error) {
	var lic []byte
	var err error
	top, mid, bot, unknown := licenseCharsForExt(path)
	if unknown {
		return nil, nil
	}
	lic, err = prefix(tmpl, data, top, mid, bot)
	return lic, err
}

//nolint:nakedret
func licenseCharsForExt(path string) (top, mid, bot string, unknown bool) {
	switch fileExtension(path) {
	default:
		return "", "", "", true
	case ".c", ".h":
		top = "/*"
		mid = " * "
		bot = " */"
	case ".js", ".mjs", ".cjs", ".jsx", ".tsx", ".css", ".tf", ".ts":
		top = "/**"
		mid = " * "
		bot = " */"
	case ".cc", ".cpp", ".cs", ".go", ".hh", ".hpp", ".java", ".m", ".mm",
		".proto", ".rs", ".scala", ".swift", ".dart", ".groovy", ".kt", ".kts":
		top = ""
		mid = "// "
		bot = ""
	case ".py", ".sh", ".yaml", ".yml", ".dockerfile", "dockerfile", ".rb", "gemfile":
		top = ""
		mid = "# "
		bot = ""
	case ".el", ".lisp":
		top = ""
		mid = ";; "
		bot = ""
	case ".erl":
		top = ""
		mid = "% "
		bot = ""
	case ".hs", ".sql":
		top = ""
		mid = "-- "
		bot = ""
	case ".html", ".xml", ".vue":
		top = "<!--"
		mid = " "
		bot = "-->"
	case ".php":
		top = ""
		mid = "// "
		bot = ""
	case ".ml", ".mli", ".mll", ".mly":
		top = "(**"
		mid = "   "
		bot = "*)"
	}
	return
}

// prefix will execute a license template t with data d
// and prefix the result with top, middle and bottom.
func prefix(t *template.Template, d *copyrightInfo, top, mid, bot string) ([]byte, error) {
	var buf bytes.Buffer
	if err := t.Execute(&buf, d); err != nil {
		return nil, err
	}
	var out bytes.Buffer
	if top != "" {
		_, _ = fmt.Fprintln(&out, top)
	}
	s := bufio.NewScanner(&buf)
	for s.Scan() {
		_, _ = fmt.Fprintln(&out, strings.TrimRightFunc(mid+s.Text(), unicode.IsSpace))
	}
	if bot != "" {
		_, _ = fmt.Fprintln(&out, bot)
	}
	_, _ = fmt.Fprintln(&out)

	return out.Bytes(), nil
}

func fileExtension(name string) string {
	if v := filepath.Ext(name); v != "" {
		return strings.ToLower(v)
	}

	return strings.ToLower(filepath.Base(name))
}
