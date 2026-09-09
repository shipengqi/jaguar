package app

import "{{ .App.ModuleName }}/pkg/app/flag"

// CliOptions abstracts configuration options for reading from CLI flags.
type CliOptions interface {
	Flags() flag.NamedFlagSets
	Validate() []error
}

// CompletableOptions abstracts options that can be completed.
type CompletableOptions interface {
	Complete() error
}

// PrintableOptions abstracts options that can be printed.
type PrintableOptions interface {
	String() string
}

// Option is a functional option for App.
type Option func(*App)

// WithRunFunc sets the application run callback.
func WithRunFunc(fn RunFunc) Option {
	return func(a *App) { a.runfunc = fn }
}

// WithBaseName sets the binary base name.
func WithBaseName(name string) Option {
	return func(a *App) { a.basename = name }
}

// WithDesc sets the application description.
func WithDesc(desc string) Option {
	return func(a *App) { a.description = desc }
}

// WithCliOptions sets the CLI options.
func WithCliOptions(opts CliOptions) Option {
	return func(a *App) { a.opts = opts }
}

// WithSilence disables startup log output.
func WithSilence() Option {
	return func(a *App) { a.silence = true }
}

// WithoutVersion disables the --version flag.
func WithoutVersion() Option {
	return func(a *App) { a.noVersion = true }
}

// WithoutConfig disables the --config flag.
func WithoutConfig() Option {
	return func(a *App) { a.noConfig = true }
}
