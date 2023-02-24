// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/message/auth.proto

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

// Validate checks the field values on LoginMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginMessage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginMessageMultiError, or
// nil if none found.
func (m *LoginMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetUsername()) < 1 {
		err := LoginMessageValidationError{
			field:  "Username",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPassword()) < 1 {
		err := LoginMessageValidationError{
			field:  "Password",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return LoginMessageMultiError(errors)
	}

	return nil
}

// LoginMessageMultiError is an error wrapping multiple validation errors
// returned by LoginMessage.ValidateAll() if the designated constraints aren't met.
type LoginMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginMessageMultiError) AllErrors() []error { return m }

// LoginMessageValidationError is the validation error returned by
// LoginMessage.Validate if the designated constraints aren't met.
type LoginMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginMessageValidationError) ErrorName() string { return "LoginMessageValidationError" }

// Error satisfies the builtin error interface
func (e LoginMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginMessageValidationError{}

// Validate checks the field values on LoginResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginResponseMultiError, or
// nil if none found.
func (m *LoginResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	if len(errors) > 0 {
		return LoginResponseMultiError(errors)
	}

	return nil
}

// LoginResponseMultiError is an error wrapping multiple validation errors
// returned by LoginResponse.ValidateAll() if the designated constraints
// aren't met.
type LoginResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginResponseMultiError) AllErrors() []error { return m }

// LoginResponseValidationError is the validation error returned by
// LoginResponse.Validate if the designated constraints aren't met.
type LoginResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginResponseValidationError) ErrorName() string { return "LoginResponseValidationError" }

// Error satisfies the builtin error interface
func (e LoginResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginResponseValidationError{}

// Validate checks the field values on QuickLoginMessage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *QuickLoginMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuickLoginMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QuickLoginMessageMultiError, or nil if none found.
func (m *QuickLoginMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *QuickLoginMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetUsername()) < 1 {
		err := QuickLoginMessageValidationError{
			field:  "Username",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return QuickLoginMessageMultiError(errors)
	}

	return nil
}

// QuickLoginMessageMultiError is an error wrapping multiple validation errors
// returned by QuickLoginMessage.ValidateAll() if the designated constraints
// aren't met.
type QuickLoginMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuickLoginMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuickLoginMessageMultiError) AllErrors() []error { return m }

// QuickLoginMessageValidationError is the validation error returned by
// QuickLoginMessage.Validate if the designated constraints aren't met.
type QuickLoginMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuickLoginMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuickLoginMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuickLoginMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuickLoginMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuickLoginMessageValidationError) ErrorName() string {
	return "QuickLoginMessageValidationError"
}

// Error satisfies the builtin error interface
func (e QuickLoginMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuickLoginMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuickLoginMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuickLoginMessageValidationError{}

// Validate checks the field values on QuickLoginResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *QuickLoginResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuickLoginResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QuickLoginResponseMultiError, or nil if none found.
func (m *QuickLoginResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *QuickLoginResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	if len(errors) > 0 {
		return QuickLoginResponseMultiError(errors)
	}

	return nil
}

// QuickLoginResponseMultiError is an error wrapping multiple validation errors
// returned by QuickLoginResponse.ValidateAll() if the designated constraints
// aren't met.
type QuickLoginResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuickLoginResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuickLoginResponseMultiError) AllErrors() []error { return m }

// QuickLoginResponseValidationError is the validation error returned by
// QuickLoginResponse.Validate if the designated constraints aren't met.
type QuickLoginResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuickLoginResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuickLoginResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuickLoginResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuickLoginResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuickLoginResponseValidationError) ErrorName() string {
	return "QuickLoginResponseValidationError"
}

// Error satisfies the builtin error interface
func (e QuickLoginResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuickLoginResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuickLoginResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuickLoginResponseValidationError{}

// Validate checks the field values on RegisterMessage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RegisterMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterMessageMultiError, or nil if none found.
func (m *RegisterMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for FirstName

	// no validation rules for LastName

	// no validation rules for Username

	// no validation rules for Email

	// no validation rules for Password

	for idx, item := range m.GetFields() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RegisterMessageValidationError{
						field:  fmt.Sprintf("Fields[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RegisterMessageValidationError{
						field:  fmt.Sprintf("Fields[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RegisterMessageValidationError{
					field:  fmt.Sprintf("Fields[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for IsQuick

	if len(errors) > 0 {
		return RegisterMessageMultiError(errors)
	}

	return nil
}

// RegisterMessageMultiError is an error wrapping multiple validation errors
// returned by RegisterMessage.ValidateAll() if the designated constraints
// aren't met.
type RegisterMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterMessageMultiError) AllErrors() []error { return m }

// RegisterMessageValidationError is the validation error returned by
// RegisterMessage.Validate if the designated constraints aren't met.
type RegisterMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterMessageValidationError) ErrorName() string { return "RegisterMessageValidationError" }

// Error satisfies the builtin error interface
func (e RegisterMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterMessageValidationError{}

// Validate checks the field values on RegisterResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RegisterResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterResponseMultiError, or nil if none found.
func (m *RegisterResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return RegisterResponseMultiError(errors)
	}

	return nil
}

// RegisterResponseMultiError is an error wrapping multiple validation errors
// returned by RegisterResponse.ValidateAll() if the designated constraints
// aren't met.
type RegisterResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterResponseMultiError) AllErrors() []error { return m }

// RegisterResponseValidationError is the validation error returned by
// RegisterResponse.Validate if the designated constraints aren't met.
type RegisterResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterResponseValidationError) ErrorName() string { return "RegisterResponseValidationError" }

// Error satisfies the builtin error interface
func (e RegisterResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterResponseValidationError{}

// Validate checks the field values on VerifyMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *VerifyMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VerifyMessage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in VerifyMessageMultiError, or
// nil if none found.
func (m *VerifyMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *VerifyMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return VerifyMessageMultiError(errors)
	}

	return nil
}

// VerifyMessageMultiError is an error wrapping multiple validation errors
// returned by VerifyMessage.ValidateAll() if the designated constraints
// aren't met.
type VerifyMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VerifyMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VerifyMessageMultiError) AllErrors() []error { return m }

// VerifyMessageValidationError is the validation error returned by
// VerifyMessage.Validate if the designated constraints aren't met.
type VerifyMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyMessageValidationError) ErrorName() string { return "VerifyMessageValidationError" }

// Error satisfies the builtin error interface
func (e VerifyMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyMessageValidationError{}

// Validate checks the field values on VerifyResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *VerifyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VerifyResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in VerifyResponseMultiError,
// or nil if none found.
func (m *VerifyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *VerifyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, VerifyResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, VerifyResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return VerifyResponseValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return VerifyResponseMultiError(errors)
	}

	return nil
}

// VerifyResponseMultiError is an error wrapping multiple validation errors
// returned by VerifyResponse.ValidateAll() if the designated constraints
// aren't met.
type VerifyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VerifyResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VerifyResponseMultiError) AllErrors() []error { return m }

// VerifyResponseValidationError is the validation error returned by
// VerifyResponse.Validate if the designated constraints aren't met.
type VerifyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyResponseValidationError) ErrorName() string { return "VerifyResponseValidationError" }

// Error satisfies the builtin error interface
func (e VerifyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyResponseValidationError{}
