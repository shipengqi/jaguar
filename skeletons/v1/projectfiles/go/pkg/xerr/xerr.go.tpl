package xerr

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
)

// Coder defines the interface for error codes.
type Coder interface {
	Code() int
	HTTPStatus() int
	String() string
	Reference() string
}

type defaultCoder struct {
	code   int
	status int
	msg    string
	ref    string
}

func (d defaultCoder) Code() int         { return d.code }
func (d defaultCoder) String() string    { return d.msg }
func (d defaultCoder) Reference() string { return d.ref }
func (d defaultCoder) HTTPStatus() int {
	if d.status == 0 {
		return http.StatusInternalServerError
	}
	return d.status
}

// NewCoder creates a new Coder.
func NewCoder(code, status int, msg, ref string) Coder {
	return defaultCoder{code: code, status: status, msg: msg, ref: ref}
}

var (
	unknownCoder = defaultCoder{code: 1, status: http.StatusInternalServerError, msg: "internal server error"}
	mu           sync.RWMutex
	codes        = map[int]Coder{1: unknownCoder}
)

// Register registers a Coder. Panics on duplicate codes.
func Register(c Coder) {
	if c.Code() == unknownCoder.code {
		panic(fmt.Sprintf("code %d is reserved", c.Code()))
	}
	mu.Lock()
	defer mu.Unlock()
	if _, ok := codes[c.Code()]; ok {
		panic(fmt.Sprintf("code %d already registered", c.Code()))
	}
	codes[c.Code()] = c
}

// ParseCoder returns the Coder for the given error, or the unknown coder.
func ParseCoder(err error) Coder {
	if err == nil {
		return nil
	}
	type icoder interface{ Code() int }
	type causer interface{ Cause() error }

	if v, ok := err.(icoder); ok {
		mu.RLock()
		c, found := codes[v.Code()]
		mu.RUnlock()
		if found {
			return c
		}
	}
	if v, ok := err.(causer); ok {
		return ParseCoder(v.Cause())
	}
	return unknownCoder
}

// IsCode reports whether any error in the chain has the given code.
func IsCode(err error, code int) bool {
	type icoder interface{ Code() int }
	type causer interface{ Cause() error }
	if v, ok := err.(icoder); ok && v.Code() == code {
		return true
	}
	if v, ok := err.(causer); ok {
		return IsCode(v.Cause(), code)
	}
	return false
}

// ----- error types -----

type withCode struct {
	cause error
	code  int
}

func (w *withCode) Error() string  { return fmt.Sprintf("code: %d, %s", w.code, w.cause.Error()) }
func (w *withCode) Cause() error   { return w.cause }
func (w *withCode) Unwrap() error  { return w.cause }
func (w *withCode) Code() int      { return w.code }
func (w *withCode) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = fmt.Fprintf(s, "code: %d, %+v", w.code, w.Cause())
			return
		}
		fallthrough
	case 's', 'q':
		_, _ = io.WriteString(s, w.Error())
	}
}

type withMessage struct {
	cause error
	msg   string
}

func (w *withMessage) Error() string { return w.msg + ": " + w.cause.Error() }
func (w *withMessage) Cause() error  { return w.cause }
func (w *withMessage) Unwrap() error { return w.cause }
func (w *withMessage) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = fmt.Fprintf(s, "%+v\n", w.Cause())
			_, _ = io.WriteString(s, w.msg)
			return
		}
		fallthrough
	case 's', 'q':
		_, _ = io.WriteString(s, w.Error())
	}
}

// New returns a new error with a stack trace.
func New(msg string) error { return fmt.Errorf("%s", msg) }

// Errorf returns a formatted error.
func Errorf(format string, args ...any) error { return fmt.Errorf(format, args...) }

// Wrap annotates err with a message.
func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	return &withMessage{cause: err, msg: msg}
}

// Wrapf annotates err with a formatted message.
func Wrapf(err error, format string, args ...any) error {
	if err == nil {
		return nil
	}
	return &withMessage{cause: err, msg: fmt.Sprintf(format, args...)}
}

// WithCode annotates err with an error code.
func WithCode(err error, code int) error {
	if err == nil {
		return nil
	}
	return &withCode{cause: err, code: code}
}

// Codef creates a new error with the given code and a formatted message.
func Codef(code int, format string, args ...any) error {
	return &withCode{cause: fmt.Errorf(format, args...), code: code}
}

// WrapCode returns an error with a code wrapping err.
func WrapCode(err error, code int) error { return WithCode(err, code) }

// WrapCodef returns an error with a code and formatted message wrapping err.
func WrapCodef(err error, code int, format string, args ...any) error {
	if err == nil {
		return nil
	}
	return &withCode{cause: &withMessage{cause: err, msg: fmt.Sprintf(format, args...)}, code: code}
}

// Cause returns the root cause of the error chain.
func Cause(err error) error {
	type causer interface{ Cause() error }
	for err != nil {
		c, ok := err.(causer)
		if !ok {
			break
		}
		err = c.Cause()
	}
	return err
}

// Is reports whether any error in the chain matches target.
func Is(err, target error) bool { return errors.Is(err, target) }
