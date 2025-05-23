// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/message/integration.proto

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

// Validate checks the field values on IntegrationListMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IntegrationListMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IntegrationListMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IntegrationListMessageMultiError, or nil if none found.
func (m *IntegrationListMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *IntegrationListMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return IntegrationListMessageMultiError(errors)
	}

	return nil
}

// IntegrationListMessageMultiError is an error wrapping multiple validation
// errors returned by IntegrationListMessage.ValidateAll() if the designated
// constraints aren't met.
type IntegrationListMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IntegrationListMessageMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IntegrationListMessageMultiError) AllErrors() []error { return m }

// IntegrationListMessageValidationError is the validation error returned by
// IntegrationListMessage.Validate if the designated constraints aren't met.
type IntegrationListMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IntegrationListMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IntegrationListMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IntegrationListMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IntegrationListMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IntegrationListMessageValidationError) ErrorName() string {
	return "IntegrationListMessageValidationError"
}

// Error satisfies the builtin error interface
func (e IntegrationListMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIntegrationListMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IntegrationListMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IntegrationListMessageValidationError{}

// Validate checks the field values on IntegrationListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IntegrationListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IntegrationListResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IntegrationListResponseMultiError, or nil if none found.
func (m *IntegrationListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *IntegrationListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetIntegrations() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, IntegrationListResponseValidationError{
						field:  fmt.Sprintf("Integrations[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, IntegrationListResponseValidationError{
						field:  fmt.Sprintf("Integrations[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return IntegrationListResponseValidationError{
					field:  fmt.Sprintf("Integrations[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return IntegrationListResponseMultiError(errors)
	}

	return nil
}

// IntegrationListResponseMultiError is an error wrapping multiple validation
// errors returned by IntegrationListResponse.ValidateAll() if the designated
// constraints aren't met.
type IntegrationListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IntegrationListResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IntegrationListResponseMultiError) AllErrors() []error { return m }

// IntegrationListResponseValidationError is the validation error returned by
// IntegrationListResponse.Validate if the designated constraints aren't met.
type IntegrationListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IntegrationListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IntegrationListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IntegrationListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IntegrationListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IntegrationListResponseValidationError) ErrorName() string {
	return "IntegrationListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e IntegrationListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIntegrationListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IntegrationListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IntegrationListResponseValidationError{}

// Validate checks the field values on IntegrationGetMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IntegrationGetMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IntegrationGetMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IntegrationGetMessageMultiError, or nil if none found.
func (m *IntegrationGetMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *IntegrationGetMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return IntegrationGetMessageMultiError(errors)
	}

	return nil
}

// IntegrationGetMessageMultiError is an error wrapping multiple validation
// errors returned by IntegrationGetMessage.ValidateAll() if the designated
// constraints aren't met.
type IntegrationGetMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IntegrationGetMessageMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IntegrationGetMessageMultiError) AllErrors() []error { return m }

// IntegrationGetMessageValidationError is the validation error returned by
// IntegrationGetMessage.Validate if the designated constraints aren't met.
type IntegrationGetMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IntegrationGetMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IntegrationGetMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IntegrationGetMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IntegrationGetMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IntegrationGetMessageValidationError) ErrorName() string {
	return "IntegrationGetMessageValidationError"
}

// Error satisfies the builtin error interface
func (e IntegrationGetMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIntegrationGetMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IntegrationGetMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IntegrationGetMessageValidationError{}

// Validate checks the field values on IntegrationGetResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IntegrationGetResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IntegrationGetResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IntegrationGetResponseMultiError, or nil if none found.
func (m *IntegrationGetResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *IntegrationGetResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetIntegration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IntegrationGetResponseValidationError{
					field:  "Integration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IntegrationGetResponseValidationError{
					field:  "Integration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIntegration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IntegrationGetResponseValidationError{
				field:  "Integration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return IntegrationGetResponseMultiError(errors)
	}

	return nil
}

// IntegrationGetResponseMultiError is an error wrapping multiple validation
// errors returned by IntegrationGetResponse.ValidateAll() if the designated
// constraints aren't met.
type IntegrationGetResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IntegrationGetResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IntegrationGetResponseMultiError) AllErrors() []error { return m }

// IntegrationGetResponseValidationError is the validation error returned by
// IntegrationGetResponse.Validate if the designated constraints aren't met.
type IntegrationGetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IntegrationGetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IntegrationGetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IntegrationGetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IntegrationGetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IntegrationGetResponseValidationError) ErrorName() string {
	return "IntegrationGetResponseValidationError"
}

// Error satisfies the builtin error interface
func (e IntegrationGetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIntegrationGetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IntegrationGetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IntegrationGetResponseValidationError{}

// Validate checks the field values on IntegrationEnableMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IntegrationEnableMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IntegrationEnableMessage with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IntegrationEnableMessageMultiError, or nil if none found.
func (m *IntegrationEnableMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *IntegrationEnableMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return IntegrationEnableMessageMultiError(errors)
	}

	return nil
}

// IntegrationEnableMessageMultiError is an error wrapping multiple validation
// errors returned by IntegrationEnableMessage.ValidateAll() if the designated
// constraints aren't met.
type IntegrationEnableMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IntegrationEnableMessageMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IntegrationEnableMessageMultiError) AllErrors() []error { return m }

// IntegrationEnableMessageValidationError is the validation error returned by
// IntegrationEnableMessage.Validate if the designated constraints aren't met.
type IntegrationEnableMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IntegrationEnableMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IntegrationEnableMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IntegrationEnableMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IntegrationEnableMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IntegrationEnableMessageValidationError) ErrorName() string {
	return "IntegrationEnableMessageValidationError"
}

// Error satisfies the builtin error interface
func (e IntegrationEnableMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIntegrationEnableMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IntegrationEnableMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IntegrationEnableMessageValidationError{}

// Validate checks the field values on IntegrationEnableResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IntegrationEnableResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IntegrationEnableResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IntegrationEnableResponseMultiError, or nil if none found.
func (m *IntegrationEnableResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *IntegrationEnableResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return IntegrationEnableResponseMultiError(errors)
	}

	return nil
}

// IntegrationEnableResponseMultiError is an error wrapping multiple validation
// errors returned by IntegrationEnableResponse.ValidateAll() if the
// designated constraints aren't met.
type IntegrationEnableResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IntegrationEnableResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IntegrationEnableResponseMultiError) AllErrors() []error { return m }

// IntegrationEnableResponseValidationError is the validation error returned by
// IntegrationEnableResponse.Validate if the designated constraints aren't met.
type IntegrationEnableResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IntegrationEnableResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IntegrationEnableResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IntegrationEnableResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IntegrationEnableResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IntegrationEnableResponseValidationError) ErrorName() string {
	return "IntegrationEnableResponseValidationError"
}

// Error satisfies the builtin error interface
func (e IntegrationEnableResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIntegrationEnableResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IntegrationEnableResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IntegrationEnableResponseValidationError{}

// Validate checks the field values on IntegrationDisableMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IntegrationDisableMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IntegrationDisableMessage with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IntegrationDisableMessageMultiError, or nil if none found.
func (m *IntegrationDisableMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *IntegrationDisableMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return IntegrationDisableMessageMultiError(errors)
	}

	return nil
}

// IntegrationDisableMessageMultiError is an error wrapping multiple validation
// errors returned by IntegrationDisableMessage.ValidateAll() if the
// designated constraints aren't met.
type IntegrationDisableMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IntegrationDisableMessageMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IntegrationDisableMessageMultiError) AllErrors() []error { return m }

// IntegrationDisableMessageValidationError is the validation error returned by
// IntegrationDisableMessage.Validate if the designated constraints aren't met.
type IntegrationDisableMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IntegrationDisableMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IntegrationDisableMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IntegrationDisableMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IntegrationDisableMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IntegrationDisableMessageValidationError) ErrorName() string {
	return "IntegrationDisableMessageValidationError"
}

// Error satisfies the builtin error interface
func (e IntegrationDisableMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIntegrationDisableMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IntegrationDisableMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IntegrationDisableMessageValidationError{}

// Validate checks the field values on IntegrationDisableResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IntegrationDisableResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IntegrationDisableResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IntegrationDisableResponseMultiError, or nil if none found.
func (m *IntegrationDisableResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *IntegrationDisableResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return IntegrationDisableResponseMultiError(errors)
	}

	return nil
}

// IntegrationDisableResponseMultiError is an error wrapping multiple
// validation errors returned by IntegrationDisableResponse.ValidateAll() if
// the designated constraints aren't met.
type IntegrationDisableResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IntegrationDisableResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IntegrationDisableResponseMultiError) AllErrors() []error { return m }

// IntegrationDisableResponseValidationError is the validation error returned
// by IntegrationDisableResponse.Validate if the designated constraints aren't met.
type IntegrationDisableResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IntegrationDisableResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IntegrationDisableResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IntegrationDisableResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IntegrationDisableResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IntegrationDisableResponseValidationError) ErrorName() string {
	return "IntegrationDisableResponseValidationError"
}

// Error satisfies the builtin error interface
func (e IntegrationDisableResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIntegrationDisableResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IntegrationDisableResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IntegrationDisableResponseValidationError{}
