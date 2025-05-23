// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/attribute.proto

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

// Validate checks the field values on Attribute with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Attribute) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Attribute with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AttributeMultiError, or nil
// if none found.
func (m *Attribute) ValidateAll() error {
	return m.validate(true)
}

func (m *Attribute) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return AttributeMultiError(errors)
	}

	return nil
}

// AttributeMultiError is an error wrapping multiple validation errors returned
// by Attribute.ValidateAll() if the designated constraints aren't met.
type AttributeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttributeMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttributeMultiError) AllErrors() []error { return m }

// AttributeValidationError is the validation error returned by
// Attribute.Validate if the designated constraints aren't met.
type AttributeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttributeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributeValidationError) ErrorName() string { return "AttributeValidationError" }

// Error satisfies the builtin error interface
func (e AttributeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttribute.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeValidationError{}
