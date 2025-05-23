// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/option.proto

package data

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Option with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Option) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Option with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in OptionMultiError, or nil if none found.
func (m *Option) ValidateAll() error {
	return m.validate(true)
}

func (m *Option) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Value

	if len(errors) > 0 {
		return OptionMultiError(errors)
	}

	return nil
}

// OptionMultiError is an error wrapping multiple validation errors returned by
// Option.ValidateAll() if the designated constraints aren't met.
type OptionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OptionMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OptionMultiError) AllErrors() []error { return m }

// OptionValidationError is the validation error returned by Option.Validate if
// the designated constraints aren't met.
type OptionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OptionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OptionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OptionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OptionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OptionValidationError) ErrorName() string { return "OptionValidationError" }

// Error satisfies the builtin error interface
func (e OptionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOption.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OptionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OptionValidationError{}
