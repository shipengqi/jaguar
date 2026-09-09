package conv

import "unsafe"

// B2S converts []byte to string without allocation.
func B2S(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return unsafe.String(&b[0], len(b))
}

// S2B converts string to []byte without allocation.
func S2B(s string) []byte {
	if s == "" {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
