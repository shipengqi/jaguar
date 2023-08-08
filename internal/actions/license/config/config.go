package config

import (
	"fmt"
	"os"
	"regexp"
	"text/template"

	"github.com/shipengqi/component-base/json"
	"github.com/shipengqi/log"

	"github.com/shipengqi/jaguar/internal/actions/license/options"
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

	if len(opts.SkipOptions.SkipDirs) > 0 {
		ps, err := getPatterns(opts.SkipOptions.SkipDirs)
		if err != nil {
			return nil, err
		}
		cfg.SkipDirRegs = ps
	}
	if len(opts.SkipOptions.SkipFiles) > 0 {
		ps, err := getPatterns(opts.SkipOptions.SkipFiles)
		if err != nil {
			return nil, err
		}
		cfg.SkipFileRegs = ps
	}

	var t *template.Template
	if opts.HeaderOptions.LicenseFile != "" {
		d, err := os.ReadFile(opts.HeaderOptions.LicenseFile)
		if err != nil {
			log.Errorf("license file: %v\n", err)
			return nil, err
		}
		t, err = template.New("").Parse(string(d))
		if err != nil {
			log.Errorf("license file: %v\n", err)
			return nil, err
		}
	} else {
		t = licenseTemplates[opts.HeaderOptions.License]
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
