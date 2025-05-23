// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/hook/category.proto

package hook

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

// Validate checks the field values on CategoryListHookParams with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CategoryListHookParams) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CategoryListHookParams with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CategoryListHookParamsMultiError, or nil if none found.
func (m *CategoryListHookParams) ValidateAll() error {
	return m.validate(true)
}

func (m *CategoryListHookParams) validate(all bool) error {
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
					errors = append(errors, CategoryListHookParamsValidationError{
						field:  fmt.Sprintf("Categories[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CategoryListHookParamsValidationError{
						field:  fmt.Sprintf("Categories[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CategoryListHookParamsValidationError{
					field:  fmt.Sprintf("Categories[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetMessage()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CategoryListHookParamsValidationError{
					field:  "Message",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CategoryListHookParamsValidationError{
					field:  "Message",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMessage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CategoryListHookParamsValidationError{
				field:  "Message",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CategoryListHookParamsMultiError(errors)
	}

	return nil
}

// CategoryListHookParamsMultiError is an error wrapping multiple validation
// errors returned by CategoryListHookParams.ValidateAll() if the designated
// constraints aren't met.
type CategoryListHookParamsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CategoryListHookParamsMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CategoryListHookParamsMultiError) AllErrors() []error { return m }

// CategoryListHookParamsValidationError is the validation error returned by
// CategoryListHookParams.Validate if the designated constraints aren't met.
type CategoryListHookParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CategoryListHookParamsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CategoryListHookParamsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CategoryListHookParamsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CategoryListHookParamsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CategoryListHookParamsValidationError) ErrorName() string {
	return "CategoryListHookParamsValidationError"
}

// Error satisfies the builtin error interface
func (e CategoryListHookParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCategoryListHookParams.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CategoryListHookParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CategoryListHookParamsValidationError{}

// Validate checks the field values on CategoryGetHookParams with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CategoryGetHookParams) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CategoryGetHookParams with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CategoryGetHookParamsMultiError, or nil if none found.
func (m *CategoryGetHookParams) ValidateAll() error {
	return m.validate(true)
}

func (m *CategoryGetHookParams) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCategory()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CategoryGetHookParamsValidationError{
					field:  "Category",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CategoryGetHookParamsValidationError{
					field:  "Category",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCategory()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CategoryGetHookParamsValidationError{
				field:  "Category",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CategoryGetHookParamsMultiError(errors)
	}

	return nil
}

// CategoryGetHookParamsMultiError is an error wrapping multiple validation
// errors returned by CategoryGetHookParams.ValidateAll() if the designated
// constraints aren't met.
type CategoryGetHookParamsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CategoryGetHookParamsMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CategoryGetHookParamsMultiError) AllErrors() []error { return m }

// CategoryGetHookParamsValidationError is the validation error returned by
// CategoryGetHookParams.Validate if the designated constraints aren't met.
type CategoryGetHookParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CategoryGetHookParamsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CategoryGetHookParamsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CategoryGetHookParamsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CategoryGetHookParamsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CategoryGetHookParamsValidationError) ErrorName() string {
	return "CategoryGetHookParamsValidationError"
}

// Error satisfies the builtin error interface
func (e CategoryGetHookParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCategoryGetHookParams.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CategoryGetHookParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CategoryGetHookParamsValidationError{}
