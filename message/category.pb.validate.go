// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/message/category.proto

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

// Validate checks the field values on CategoryListMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CategoryListMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CategoryListMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CategoryListMessageMultiError, or nil if none found.
func (m *CategoryListMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *CategoryListMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CategoryListMessageMultiError(errors)
	}

	return nil
}

// CategoryListMessageMultiError is an error wrapping multiple validation
// errors returned by CategoryListMessage.ValidateAll() if the designated
// constraints aren't met.
type CategoryListMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CategoryListMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CategoryListMessageMultiError) AllErrors() []error { return m }

// CategoryListMessageValidationError is the validation error returned by
// CategoryListMessage.Validate if the designated constraints aren't met.
type CategoryListMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CategoryListMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CategoryListMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CategoryListMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CategoryListMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CategoryListMessageValidationError) ErrorName() string {
	return "CategoryListMessageValidationError"
}

// Error satisfies the builtin error interface
func (e CategoryListMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCategoryListMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CategoryListMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CategoryListMessageValidationError{}

// Validate checks the field values on CategoryListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CategoryListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CategoryListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CategoryListResponseMultiError, or nil if none found.
func (m *CategoryListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CategoryListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetCategories() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CategoryListResponseValidationError{
						field:  fmt.Sprintf("Categories[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CategoryListResponseValidationError{
						field:  fmt.Sprintf("Categories[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CategoryListResponseValidationError{
					field:  fmt.Sprintf("Categories[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CategoryListResponseMultiError(errors)
	}

	return nil
}

// CategoryListResponseMultiError is an error wrapping multiple validation
// errors returned by CategoryListResponse.ValidateAll() if the designated
// constraints aren't met.
type CategoryListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CategoryListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CategoryListResponseMultiError) AllErrors() []error { return m }

// CategoryListResponseValidationError is the validation error returned by
// CategoryListResponse.Validate if the designated constraints aren't met.
type CategoryListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CategoryListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CategoryListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CategoryListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CategoryListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CategoryListResponseValidationError) ErrorName() string {
	return "CategoryListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CategoryListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCategoryListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CategoryListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CategoryListResponseValidationError{}

// Validate checks the field values on CategoryGetMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CategoryGetMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CategoryGetMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CategoryGetMessageMultiError, or nil if none found.
func (m *CategoryGetMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *CategoryGetMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CategoryGetMessageMultiError(errors)
	}

	return nil
}

// CategoryGetMessageMultiError is an error wrapping multiple validation errors
// returned by CategoryGetMessage.ValidateAll() if the designated constraints
// aren't met.
type CategoryGetMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CategoryGetMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CategoryGetMessageMultiError) AllErrors() []error { return m }

// CategoryGetMessageValidationError is the validation error returned by
// CategoryGetMessage.Validate if the designated constraints aren't met.
type CategoryGetMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CategoryGetMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CategoryGetMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CategoryGetMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CategoryGetMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CategoryGetMessageValidationError) ErrorName() string {
	return "CategoryGetMessageValidationError"
}

// Error satisfies the builtin error interface
func (e CategoryGetMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCategoryGetMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CategoryGetMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CategoryGetMessageValidationError{}

// Validate checks the field values on CategoryGetResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CategoryGetResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CategoryGetResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CategoryGetResponseMultiError, or nil if none found.
func (m *CategoryGetResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CategoryGetResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCategory()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CategoryGetResponseValidationError{
					field:  "Category",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CategoryGetResponseValidationError{
					field:  "Category",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCategory()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CategoryGetResponseValidationError{
				field:  "Category",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CategoryGetResponseMultiError(errors)
	}

	return nil
}

// CategoryGetResponseMultiError is an error wrapping multiple validation
// errors returned by CategoryGetResponse.ValidateAll() if the designated
// constraints aren't met.
type CategoryGetResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CategoryGetResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CategoryGetResponseMultiError) AllErrors() []error { return m }

// CategoryGetResponseValidationError is the validation error returned by
// CategoryGetResponse.Validate if the designated constraints aren't met.
type CategoryGetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CategoryGetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CategoryGetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CategoryGetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CategoryGetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CategoryGetResponseValidationError) ErrorName() string {
	return "CategoryGetResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CategoryGetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCategoryGetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CategoryGetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CategoryGetResponseValidationError{}

// Validate checks the field values on CategoryCreateMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CategoryCreateMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CategoryCreateMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CategoryCreateMessageMultiError, or nil if none found.
func (m *CategoryCreateMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *CategoryCreateMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for ParentId

	if len(errors) > 0 {
		return CategoryCreateMessageMultiError(errors)
	}

	return nil
}

// CategoryCreateMessageMultiError is an error wrapping multiple validation
// errors returned by CategoryCreateMessage.ValidateAll() if the designated
// constraints aren't met.
type CategoryCreateMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CategoryCreateMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CategoryCreateMessageMultiError) AllErrors() []error { return m }

// CategoryCreateMessageValidationError is the validation error returned by
// CategoryCreateMessage.Validate if the designated constraints aren't met.
type CategoryCreateMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CategoryCreateMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CategoryCreateMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CategoryCreateMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CategoryCreateMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CategoryCreateMessageValidationError) ErrorName() string {
	return "CategoryCreateMessageValidationError"
}

// Error satisfies the builtin error interface
func (e CategoryCreateMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCategoryCreateMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CategoryCreateMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CategoryCreateMessageValidationError{}

// Validate checks the field values on CategoryCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CategoryCreateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CategoryCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CategoryCreateResponseMultiError, or nil if none found.
func (m *CategoryCreateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CategoryCreateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCategory()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CategoryCreateResponseValidationError{
					field:  "Category",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CategoryCreateResponseValidationError{
					field:  "Category",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCategory()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CategoryCreateResponseValidationError{
				field:  "Category",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CategoryCreateResponseMultiError(errors)
	}

	return nil
}

// CategoryCreateResponseMultiError is an error wrapping multiple validation
// errors returned by CategoryCreateResponse.ValidateAll() if the designated
// constraints aren't met.
type CategoryCreateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CategoryCreateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CategoryCreateResponseMultiError) AllErrors() []error { return m }

// CategoryCreateResponseValidationError is the validation error returned by
// CategoryCreateResponse.Validate if the designated constraints aren't met.
type CategoryCreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CategoryCreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CategoryCreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CategoryCreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CategoryCreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CategoryCreateResponseValidationError) ErrorName() string {
	return "CategoryCreateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CategoryCreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCategoryCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CategoryCreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CategoryCreateResponseValidationError{}
