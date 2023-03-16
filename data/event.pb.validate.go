// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/event.proto

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

// Validate checks the field values on EventData with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EventData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EventData with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EventDataMultiError, or nil
// if none found.
func (m *EventData) ValidateAll() error {
	return m.validate(true)
}

func (m *EventData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	// no validation rules for Value

	if len(errors) > 0 {
		return EventDataMultiError(errors)
	}

	return nil
}

// EventDataMultiError is an error wrapping multiple validation errors returned
// by EventData.ValidateAll() if the designated constraints aren't met.
type EventDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EventDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EventDataMultiError) AllErrors() []error { return m }

// EventDataValidationError is the validation error returned by
// EventData.Validate if the designated constraints aren't met.
type EventDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventDataValidationError) ErrorName() string { return "EventDataValidationError" }

// Error satisfies the builtin error interface
func (e EventDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEventData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventDataValidationError{}

// Validate checks the field values on Event with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Event) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Event with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in EventMultiError, or nil if none found.
func (m *Event) ValidateAll() error {
	return m.validate(true)
}

func (m *Event) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for UserId

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EventValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EventValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EventValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EventValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EventValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return EventMultiError(errors)
	}

	return nil
}

// EventMultiError is an error wrapping multiple validation errors returned by
// Event.ValidateAll() if the designated constraints aren't met.
type EventMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EventMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EventMultiError) AllErrors() []error { return m }

// EventValidationError is the validation error returned by Event.Validate if
// the designated constraints aren't met.
type EventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventValidationError) ErrorName() string { return "EventValidationError" }

// Error satisfies the builtin error interface
func (e EventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventValidationError{}
