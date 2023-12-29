// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/message/query.proto

package message

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

// Validate checks the field values on Query with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Query) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Query with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in QueryMultiError, or nil if none found.
func (m *Query) ValidateAll() error {
	return m.validate(true)
}

func (m *Query) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Page != nil {
		// no validation rules for Page
	}

	if m.OrderBy != nil {
		// no validation rules for OrderBy
	}

	if m.TimeLimit != nil {
		// no validation rules for TimeLimit
	}

	if m.SessionUUID != nil {
		// no validation rules for SessionUUID
	}

	if m.SearchTerm != nil {
		// no validation rules for SearchTerm
	}

	if m.IgnoreCache != nil {
		// no validation rules for IgnoreCache
	}

	if len(errors) > 0 {
		return QueryMultiError(errors)
	}

	return nil
}

// QueryMultiError is an error wrapping multiple validation errors returned by
// Query.ValidateAll() if the designated constraints aren't met.
type QueryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QueryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QueryMultiError) AllErrors() []error { return m }

// QueryValidationError is the validation error returned by Query.Validate if
// the designated constraints aren't met.
type QueryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryValidationError) ErrorName() string { return "QueryValidationError" }

// Error satisfies the builtin error interface
func (e QueryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuery.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryValidationError{}
