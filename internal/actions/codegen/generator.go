package codegen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/types"
	"regexp"
	"strings"

	"github.com/shipengqi/log"
	"golang.org/x/tools/go/packages"
)

// Value represents a declared constant.
type Value struct {
	comment      string
	originalName string // The name of the constant.
	name         string // The name with trimmed prefix.
	// The value is stored as a bit pattern alone. The boolean tells us
	// whether to interpret it as an int64 or a uint64; the only place
	// this matters is when sorting.
	// Much of the time the str field is all we need; it is printed
	// by Value.String.
	value  uint64 // Will be converted to int64 when needed.
	signed bool   // Whether the constant is a signed type.
	str    string // The string representation given by the "go/constant" package.
}

func (v *Value) String() string {
	return v.str
}

// ParseComment parse comment to http code and error code description.
func (v *Value) ParseComment() (string, string) {
	reg := regexp.MustCompile(`\w\s*-\s*(\d{3})\s*:\s*([A-Z].*)\s*\.\n*`)
	if !reg.MatchString(v.comment) {
		log.Infof("constant '%s' have wrong comment format, register with 500 as default", v.originalName)

		return "500", "Internal server error"
	}

	groups := reg.FindStringSubmatch(v.comment)
	if len(groups) != 3 {
		return "500", "Internal server error"
	}

	return groups[1], groups[2]
}

// File holds a single parsed file and associated data.
type File struct {
	pkg  *Package  // Package to which this file belongs.
	file *ast.File // Parsed AST.
	// These fields are reset for each type being generated.
	typeName string  // Name of the constant type.
	values   []Value // Accumulator for constant values of that type.

	trimPrefix string
}

// Package defines options for package.
type Package struct {
	name  string
	defs  map[*ast.Ident]types.Object
	files []*File
}

// Generator holds the state of the analysis. Primarily used to buffer
// the output for format.Source.
type Generator struct {
	buf bytes.Buffer // Accumulated output.
	pkg *Package     // Package we are scanning.

	trimPrefix string
}

// Printf like fmt.Printf, but add the string to g.buf.
func (g *Generator) Printf(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(&g.buf, format, args...)
}

// parsePackage analyzes the single package constructed from the patterns and tags.
// parsePackage exits if there is an error.
func (g *Generator) parsePackage(patterns []string, tags []string) {
	cfg := &packages.Config{
		//nolint:staticcheck
		Mode: packages.LoadSyntax,
		// TODO: Need to think about constants in test files. Maybe write type_string_test.go
		// in a separate pass? For later.
		Tests:      false,
		BuildFlags: []string{fmt.Sprintf("-tags=%s", strings.Join(tags, " "))},
	}
	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		// log.Fatal(err)
	}
	if len(pkgs) != 1 {
		log.Fatalf("error: %d packages found", len(pkgs))
	}
	// g.addPackage(pkgs[0])
}
