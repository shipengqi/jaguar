package config

import (
	"fmt"
	"os"
	"regexp"
	"text/template"

	"github.com/shipengqi/component-base/json"

	"github.com/shipengqi/jaguar/internal/actions/addlicense/options"
)

type Config struct {
	*options.Options

	SkipDirRegs  []*regexp.Regexp
	SkipFileRegs []*regexp.Regexp
	LicenseTmpl  *template.Template
}

func (o *Config) String() string {
	data, _ := json.Marshal(o)

	return string(data)
}

// CreateConfigFromOptions creates a running configuration instance based
// on a given command line.
func CreateConfigFromOptions(opts *options.Options) (*Config, error) {
	cfg := &Config{Options: opts}

	if len(opts.SkipDirs) > 0 {
		ps, err := getPatterns(opts.SkipDirs)
		if err != nil {
			return nil, err
		}
		cfg.SkipDirRegs = ps
	}
	if len(opts.SkipFiles) > 0 {
		ps, err := getPatterns(opts.SkipFiles)
		if err != nil {
			return nil, err
		}
		cfg.SkipFileRegs = ps
	}

	var t *template.Template
	if opts.LicenseFile != "" {
		d, err := os.ReadFile(opts.LicenseFile)
		if err != nil {
			fmt.Printf("license file: %v\n", err)
			return nil, err
		}
		t, err = template.New("").Parse(string(d))
		if err != nil {
			fmt.Printf("license file: %v\n", err)
			return nil, err
		}
	} else {
		t = licenseTemplates[opts.License]
	}
	cfg.LicenseTmpl = t

	return cfg, nil
}

func getPatterns(patterns []string) ([]*regexp.Regexp, error) {
	patternsRe := make([]*regexp.Regexp, 0, len(patterns))
	for _, p := range patterns {
		patternRe, err := regexp.Compile(p)
		if err != nil {
			fmt.Printf("can't compile regexp %q\n", p)

			return nil, err
		}
		patternsRe = append(patternsRe, patternRe)
	}

	return patternsRe, nil
}
