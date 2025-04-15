// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/variant.proto

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

// Validate checks the field values on Variant with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Variant) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Variant with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in VariantMultiError, or nil if none found.
func (m *Variant) ValidateAll() error {
	return m.validate(true)
}

func (m *Variant) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for CustomId

	// no validation rules for Image

	for idx, item := range m.GetAttributes() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, VariantValidationError{
						field:  fmt.Sprintf("Attributes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, VariantValidationError{
						field:  fmt.Sprintf("Attributes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return VariantValidationError{
					field:  fmt.Sprintf("Attributes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return VariantMultiError(errors)
	}

	return nil
}

// VariantMultiError is an error wrapping multiple validation errors returned
// by Variant.ValidateAll() if the designated constraints aren't met.
type VariantMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VariantMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VariantMultiError) AllErrors() []error { return m }

// VariantValidationError is the validation error returned by Variant.Validate
// if the designated constraints aren't met.
type VariantValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VariantValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VariantValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VariantValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VariantValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VariantValidationError) ErrorName() string { return "VariantValidationError" }

// Error satisfies the builtin error interface
func (e VariantValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVariant.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VariantValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VariantValidationError{}
