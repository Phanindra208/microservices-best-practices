package utils

import (
	"time"

	"github.com/go-openapi/strfmt"
)

// StringPtr returns the pointer for the string.
func StringPtr(str string) *string {
	return &str
}

// Float64Ptr returns pointer for the float type.
func Float64Ptr(n float64) *float64 {
	return &n
}

// Int64Ptr returns a pointer for the int64 type.
func Int64Ptr(n int64) *int64 {
	return &n
}

// DateTimePtr returns a pointer.
func DateTimePtr(t time.Time) *strfmt.DateTime {
	return (*strfmt.DateTime)(&t)
}

// Bool returns bool value from pointer.
func Bool(b *bool) bool {
	if b != nil {
		return *b
	}
	return false
}

// BoolPtr returns bool pointer.
func BoolPtr(b bool) *bool {
	return &b
}
