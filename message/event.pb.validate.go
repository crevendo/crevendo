// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/message/event.proto

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

// Validate checks the field values on EventCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *EventCreateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EventCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// EventCreateRequestMultiError, or nil if none found.
func (m *EventCreateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *EventCreateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for UserId

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EventCreateRequestValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EventCreateRequestValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EventCreateRequestValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return EventCreateRequestMultiError(errors)
	}

	return nil
}

// EventCreateRequestMultiError is an error wrapping multiple validation errors
// returned by EventCreateRequest.ValidateAll() if the designated constraints
// aren't met.
type EventCreateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EventCreateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EventCreateRequestMultiError) AllErrors() []error { return m }

// EventCreateRequestValidationError is the validation error returned by
// EventCreateRequest.Validate if the designated constraints aren't met.
type EventCreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventCreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventCreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventCreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventCreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventCreateRequestValidationError) ErrorName() string {
	return "EventCreateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e EventCreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEventCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventCreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventCreateRequestValidationError{}

// Validate checks the field values on EventCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *EventCreateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EventCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// EventCreateResponseMultiError, or nil if none found.
func (m *EventCreateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *EventCreateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return EventCreateResponseMultiError(errors)
	}

	return nil
}

// EventCreateResponseMultiError is an error wrapping multiple validation
// errors returned by EventCreateResponse.ValidateAll() if the designated
// constraints aren't met.
type EventCreateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EventCreateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EventCreateResponseMultiError) AllErrors() []error { return m }

// EventCreateResponseValidationError is the validation error returned by
// EventCreateResponse.Validate if the designated constraints aren't met.
type EventCreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventCreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventCreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventCreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventCreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventCreateResponseValidationError) ErrorName() string {
	return "EventCreateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e EventCreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEventCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventCreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventCreateResponseValidationError{}

// Validate checks the field values on EventListRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *EventListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EventListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// EventListRequestMultiError, or nil if none found.
func (m *EventListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *EventListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return EventListRequestMultiError(errors)
	}

	return nil
}

// EventListRequestMultiError is an error wrapping multiple validation errors
// returned by EventListRequest.ValidateAll() if the designated constraints
// aren't met.
type EventListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EventListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EventListRequestMultiError) AllErrors() []error { return m }

// EventListRequestValidationError is the validation error returned by
// EventListRequest.Validate if the designated constraints aren't met.
type EventListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventListRequestValidationError) ErrorName() string { return "EventListRequestValidationError" }

// Error satisfies the builtin error interface
func (e EventListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEventListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventListRequestValidationError{}

// Validate checks the field values on EventListResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *EventListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EventListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// EventListResponseMultiError, or nil if none found.
func (m *EventListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *EventListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetEvents() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EventListResponseValidationError{
						field:  fmt.Sprintf("Events[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EventListResponseValidationError{
						field:  fmt.Sprintf("Events[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EventListResponseValidationError{
					field:  fmt.Sprintf("Events[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return EventListResponseMultiError(errors)
	}

	return nil
}

// EventListResponseMultiError is an error wrapping multiple validation errors
// returned by EventListResponse.ValidateAll() if the designated constraints
// aren't met.
type EventListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EventListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EventListResponseMultiError) AllErrors() []error { return m }

// EventListResponseValidationError is the validation error returned by
// EventListResponse.Validate if the designated constraints aren't met.
type EventListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventListResponseValidationError) ErrorName() string {
	return "EventListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e EventListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEventListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventListResponseValidationError{}
