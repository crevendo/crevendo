// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/message/option.proto

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

// Validate checks the field values on OptionListMessage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *OptionListMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OptionListMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OptionListMessageMultiError, or nil if none found.
func (m *OptionListMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *OptionListMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return OptionListMessageMultiError(errors)
	}

	return nil
}

// OptionListMessageMultiError is an error wrapping multiple validation errors
// returned by OptionListMessage.ValidateAll() if the designated constraints
// aren't met.
type OptionListMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OptionListMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OptionListMessageMultiError) AllErrors() []error { return m }

// OptionListMessageValidationError is the validation error returned by
// OptionListMessage.Validate if the designated constraints aren't met.
type OptionListMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OptionListMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OptionListMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OptionListMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OptionListMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OptionListMessageValidationError) ErrorName() string {
	return "OptionListMessageValidationError"
}

// Error satisfies the builtin error interface
func (e OptionListMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOptionListMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OptionListMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OptionListMessageValidationError{}

// Validate checks the field values on OptionListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OptionListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OptionListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OptionListResponseMultiError, or nil if none found.
func (m *OptionListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *OptionListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetOptions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, OptionListResponseValidationError{
						field:  fmt.Sprintf("Options[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, OptionListResponseValidationError{
						field:  fmt.Sprintf("Options[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OptionListResponseValidationError{
					field:  fmt.Sprintf("Options[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return OptionListResponseMultiError(errors)
	}

	return nil
}

// OptionListResponseMultiError is an error wrapping multiple validation errors
// returned by OptionListResponse.ValidateAll() if the designated constraints
// aren't met.
type OptionListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OptionListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OptionListResponseMultiError) AllErrors() []error { return m }

// OptionListResponseValidationError is the validation error returned by
// OptionListResponse.Validate if the designated constraints aren't met.
type OptionListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OptionListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OptionListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OptionListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OptionListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OptionListResponseValidationError) ErrorName() string {
	return "OptionListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e OptionListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOptionListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OptionListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OptionListResponseValidationError{}

// Validate checks the field values on OptionUpdateMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OptionUpdateMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OptionUpdateMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OptionUpdateMessageMultiError, or nil if none found.
func (m *OptionUpdateMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *OptionUpdateMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetOptions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, OptionUpdateMessageValidationError{
						field:  fmt.Sprintf("Options[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, OptionUpdateMessageValidationError{
						field:  fmt.Sprintf("Options[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OptionUpdateMessageValidationError{
					field:  fmt.Sprintf("Options[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return OptionUpdateMessageMultiError(errors)
	}

	return nil
}

// OptionUpdateMessageMultiError is an error wrapping multiple validation
// errors returned by OptionUpdateMessage.ValidateAll() if the designated
// constraints aren't met.
type OptionUpdateMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OptionUpdateMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OptionUpdateMessageMultiError) AllErrors() []error { return m }

// OptionUpdateMessageValidationError is the validation error returned by
// OptionUpdateMessage.Validate if the designated constraints aren't met.
type OptionUpdateMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OptionUpdateMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OptionUpdateMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OptionUpdateMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OptionUpdateMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OptionUpdateMessageValidationError) ErrorName() string {
	return "OptionUpdateMessageValidationError"
}

// Error satisfies the builtin error interface
func (e OptionUpdateMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOptionUpdateMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OptionUpdateMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OptionUpdateMessageValidationError{}

// Validate checks the field values on OptionUpdateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OptionUpdateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OptionUpdateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OptionUpdateResponseMultiError, or nil if none found.
func (m *OptionUpdateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *OptionUpdateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return OptionUpdateResponseMultiError(errors)
	}

	return nil
}

// OptionUpdateResponseMultiError is an error wrapping multiple validation
// errors returned by OptionUpdateResponse.ValidateAll() if the designated
// constraints aren't met.
type OptionUpdateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OptionUpdateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OptionUpdateResponseMultiError) AllErrors() []error { return m }

// OptionUpdateResponseValidationError is the validation error returned by
// OptionUpdateResponse.Validate if the designated constraints aren't met.
type OptionUpdateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OptionUpdateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OptionUpdateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OptionUpdateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OptionUpdateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OptionUpdateResponseValidationError) ErrorName() string {
	return "OptionUpdateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e OptionUpdateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOptionUpdateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OptionUpdateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OptionUpdateResponseValidationError{}
