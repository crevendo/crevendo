// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/gateway.proto

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

// Validate checks the field values on Gateway with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Gateway) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Gateway with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in GatewayMultiError, or nil if none found.
func (m *Gateway) ValidateAll() error {
	return m.validate(true)
}

func (m *Gateway) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CreationType

	// no validation rules for CreationUrl

	if len(errors) > 0 {
		return GatewayMultiError(errors)
	}

	return nil
}

// GatewayMultiError is an error wrapping multiple validation errors returned
// by Gateway.ValidateAll() if the designated constraints aren't met.
type GatewayMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GatewayMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GatewayMultiError) AllErrors() []error { return m }

// GatewayValidationError is the validation error returned by Gateway.Validate
// if the designated constraints aren't met.
type GatewayValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GatewayValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GatewayValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GatewayValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GatewayValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GatewayValidationError) ErrorName() string { return "GatewayValidationError" }

// Error satisfies the builtin error interface
func (e GatewayValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGateway.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GatewayValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GatewayValidationError{}
