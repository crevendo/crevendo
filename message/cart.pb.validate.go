// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/message/cart.proto

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

// Validate checks the field values on CartGetMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CartGetMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CartGetMessage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CartGetMessageMultiError,
// or nil if none found.
func (m *CartGetMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *CartGetMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CartGetMessageMultiError(errors)
	}

	return nil
}

// CartGetMessageMultiError is an error wrapping multiple validation errors
// returned by CartGetMessage.ValidateAll() if the designated constraints
// aren't met.
type CartGetMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CartGetMessageMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CartGetMessageMultiError) AllErrors() []error { return m }

// CartGetMessageValidationError is the validation error returned by
// CartGetMessage.Validate if the designated constraints aren't met.
type CartGetMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CartGetMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CartGetMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CartGetMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CartGetMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CartGetMessageValidationError) ErrorName() string { return "CartGetMessageValidationError" }

// Error satisfies the builtin error interface
func (e CartGetMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCartGetMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CartGetMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CartGetMessageValidationError{}

// Validate checks the field values on CartGetResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CartGetResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CartGetResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CartGetResponseMultiError, or nil if none found.
func (m *CartGetResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CartGetResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCart()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CartGetResponseValidationError{
					field:  "Cart",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CartGetResponseValidationError{
					field:  "Cart",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCart()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CartGetResponseValidationError{
				field:  "Cart",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CartGetResponseMultiError(errors)
	}

	return nil
}

// CartGetResponseMultiError is an error wrapping multiple validation errors
// returned by CartGetResponse.ValidateAll() if the designated constraints
// aren't met.
type CartGetResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CartGetResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CartGetResponseMultiError) AllErrors() []error { return m }

// CartGetResponseValidationError is the validation error returned by
// CartGetResponse.Validate if the designated constraints aren't met.
type CartGetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CartGetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CartGetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CartGetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CartGetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CartGetResponseValidationError) ErrorName() string { return "CartGetResponseValidationError" }

// Error satisfies the builtin error interface
func (e CartGetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCartGetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CartGetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CartGetResponseValidationError{}

// Validate checks the field values on CartCreateMessage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CartCreateMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CartCreateMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CartCreateMessageMultiError, or nil if none found.
func (m *CartCreateMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *CartCreateMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CartCreateMessageMultiError(errors)
	}

	return nil
}

// CartCreateMessageMultiError is an error wrapping multiple validation errors
// returned by CartCreateMessage.ValidateAll() if the designated constraints
// aren't met.
type CartCreateMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CartCreateMessageMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CartCreateMessageMultiError) AllErrors() []error { return m }

// CartCreateMessageValidationError is the validation error returned by
// CartCreateMessage.Validate if the designated constraints aren't met.
type CartCreateMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CartCreateMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CartCreateMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CartCreateMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CartCreateMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CartCreateMessageValidationError) ErrorName() string {
	return "CartCreateMessageValidationError"
}

// Error satisfies the builtin error interface
func (e CartCreateMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCartCreateMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CartCreateMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CartCreateMessageValidationError{}

// Validate checks the field values on CartCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CartCreateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CartCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CartCreateResponseMultiError, or nil if none found.
func (m *CartCreateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CartCreateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCart()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CartCreateResponseValidationError{
					field:  "Cart",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CartCreateResponseValidationError{
					field:  "Cart",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCart()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CartCreateResponseValidationError{
				field:  "Cart",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CartCreateResponseMultiError(errors)
	}

	return nil
}

// CartCreateResponseMultiError is an error wrapping multiple validation errors
// returned by CartCreateResponse.ValidateAll() if the designated constraints
// aren't met.
type CartCreateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CartCreateResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CartCreateResponseMultiError) AllErrors() []error { return m }

// CartCreateResponseValidationError is the validation error returned by
// CartCreateResponse.Validate if the designated constraints aren't met.
type CartCreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CartCreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CartCreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CartCreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CartCreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CartCreateResponseValidationError) ErrorName() string {
	return "CartCreateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CartCreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCartCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CartCreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CartCreateResponseValidationError{}

// Validate checks the field values on AddItemMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddItemMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddItemMessage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddItemMessageMultiError,
// or nil if none found.
func (m *AddItemMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *AddItemMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CartId

	// no validation rules for Quantity

	for idx, item := range m.GetCustomData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AddItemMessageValidationError{
						field:  fmt.Sprintf("CustomData[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AddItemMessageValidationError{
						field:  fmt.Sprintf("CustomData[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AddItemMessageValidationError{
					field:  fmt.Sprintf("CustomData[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.ProductId != nil {
		// no validation rules for ProductId
	}

	if m.ProductCustomId != nil {
		// no validation rules for ProductCustomId
	}

	if len(errors) > 0 {
		return AddItemMessageMultiError(errors)
	}

	return nil
}

// AddItemMessageMultiError is an error wrapping multiple validation errors
// returned by AddItemMessage.ValidateAll() if the designated constraints
// aren't met.
type AddItemMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddItemMessageMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddItemMessageMultiError) AllErrors() []error { return m }

// AddItemMessageValidationError is the validation error returned by
// AddItemMessage.Validate if the designated constraints aren't met.
type AddItemMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddItemMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddItemMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddItemMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddItemMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddItemMessageValidationError) ErrorName() string { return "AddItemMessageValidationError" }

// Error satisfies the builtin error interface
func (e AddItemMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddItemMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddItemMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddItemMessageValidationError{}

// Validate checks the field values on AddItemResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AddItemResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddItemResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddItemResponseMultiError, or nil if none found.
func (m *AddItemResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AddItemResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCart()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddItemResponseValidationError{
					field:  "Cart",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddItemResponseValidationError{
					field:  "Cart",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCart()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddItemResponseValidationError{
				field:  "Cart",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AddItemResponseMultiError(errors)
	}

	return nil
}

// AddItemResponseMultiError is an error wrapping multiple validation errors
// returned by AddItemResponse.ValidateAll() if the designated constraints
// aren't met.
type AddItemResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddItemResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddItemResponseMultiError) AllErrors() []error { return m }

// AddItemResponseValidationError is the validation error returned by
// AddItemResponse.Validate if the designated constraints aren't met.
type AddItemResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddItemResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddItemResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddItemResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddItemResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddItemResponseValidationError) ErrorName() string { return "AddItemResponseValidationError" }

// Error satisfies the builtin error interface
func (e AddItemResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddItemResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddItemResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddItemResponseValidationError{}

// Validate checks the field values on RemoveItemMessage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RemoveItemMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoveItemMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RemoveItemMessageMultiError, or nil if none found.
func (m *RemoveItemMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoveItemMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CartId

	// no validation rules for ItemId

	if len(errors) > 0 {
		return RemoveItemMessageMultiError(errors)
	}

	return nil
}

// RemoveItemMessageMultiError is an error wrapping multiple validation errors
// returned by RemoveItemMessage.ValidateAll() if the designated constraints
// aren't met.
type RemoveItemMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoveItemMessageMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoveItemMessageMultiError) AllErrors() []error { return m }

// RemoveItemMessageValidationError is the validation error returned by
// RemoveItemMessage.Validate if the designated constraints aren't met.
type RemoveItemMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveItemMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveItemMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveItemMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveItemMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveItemMessageValidationError) ErrorName() string {
	return "RemoveItemMessageValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveItemMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveItemMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveItemMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveItemMessageValidationError{}

// Validate checks the field values on RemoveItemResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RemoveItemResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoveItemResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RemoveItemResponseMultiError, or nil if none found.
func (m *RemoveItemResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoveItemResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCart()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RemoveItemResponseValidationError{
					field:  "Cart",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RemoveItemResponseValidationError{
					field:  "Cart",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCart()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RemoveItemResponseValidationError{
				field:  "Cart",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RemoveItemResponseMultiError(errors)
	}

	return nil
}

// RemoveItemResponseMultiError is an error wrapping multiple validation errors
// returned by RemoveItemResponse.ValidateAll() if the designated constraints
// aren't met.
type RemoveItemResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoveItemResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoveItemResponseMultiError) AllErrors() []error { return m }

// RemoveItemResponseValidationError is the validation error returned by
// RemoveItemResponse.Validate if the designated constraints aren't met.
type RemoveItemResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveItemResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveItemResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveItemResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveItemResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveItemResponseValidationError) ErrorName() string {
	return "RemoveItemResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveItemResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveItemResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveItemResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveItemResponseValidationError{}

// Validate checks the field values on UpdateItemQuantityMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateItemQuantityMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateItemQuantityMessage with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateItemQuantityMessageMultiError, or nil if none found.
func (m *UpdateItemQuantityMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateItemQuantityMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CartId

	// no validation rules for ItemId

	// no validation rules for Quantity

	if len(errors) > 0 {
		return UpdateItemQuantityMessageMultiError(errors)
	}

	return nil
}

// UpdateItemQuantityMessageMultiError is an error wrapping multiple validation
// errors returned by UpdateItemQuantityMessage.ValidateAll() if the
// designated constraints aren't met.
type UpdateItemQuantityMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateItemQuantityMessageMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateItemQuantityMessageMultiError) AllErrors() []error { return m }

// UpdateItemQuantityMessageValidationError is the validation error returned by
// UpdateItemQuantityMessage.Validate if the designated constraints aren't met.
type UpdateItemQuantityMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateItemQuantityMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateItemQuantityMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateItemQuantityMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateItemQuantityMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateItemQuantityMessageValidationError) ErrorName() string {
	return "UpdateItemQuantityMessageValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateItemQuantityMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateItemQuantityMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateItemQuantityMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateItemQuantityMessageValidationError{}

// Validate checks the field values on UpdateItemQuantityResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateItemQuantityResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateItemQuantityResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateItemQuantityResponseMultiError, or nil if none found.
func (m *UpdateItemQuantityResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateItemQuantityResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCart()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateItemQuantityResponseValidationError{
					field:  "Cart",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateItemQuantityResponseValidationError{
					field:  "Cart",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCart()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateItemQuantityResponseValidationError{
				field:  "Cart",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateItemQuantityResponseMultiError(errors)
	}

	return nil
}

// UpdateItemQuantityResponseMultiError is an error wrapping multiple
// validation errors returned by UpdateItemQuantityResponse.ValidateAll() if
// the designated constraints aren't met.
type UpdateItemQuantityResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateItemQuantityResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateItemQuantityResponseMultiError) AllErrors() []error { return m }

// UpdateItemQuantityResponseValidationError is the validation error returned
// by UpdateItemQuantityResponse.Validate if the designated constraints aren't met.
type UpdateItemQuantityResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateItemQuantityResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateItemQuantityResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateItemQuantityResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateItemQuantityResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateItemQuantityResponseValidationError) ErrorName() string {
	return "UpdateItemQuantityResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateItemQuantityResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateItemQuantityResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateItemQuantityResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateItemQuantityResponseValidationError{}

// Validate checks the field values on GetOrderTotalResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetOrderTotalResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetOrderTotalResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetOrderTotalResponseMultiError, or nil if none found.
func (m *GetOrderTotalResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetOrderTotalResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	if len(errors) > 0 {
		return GetOrderTotalResponseMultiError(errors)
	}

	return nil
}

// GetOrderTotalResponseMultiError is an error wrapping multiple validation
// errors returned by GetOrderTotalResponse.ValidateAll() if the designated
// constraints aren't met.
type GetOrderTotalResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetOrderTotalResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetOrderTotalResponseMultiError) AllErrors() []error { return m }

// GetOrderTotalResponseValidationError is the validation error returned by
// GetOrderTotalResponse.Validate if the designated constraints aren't met.
type GetOrderTotalResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOrderTotalResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOrderTotalResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOrderTotalResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOrderTotalResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOrderTotalResponseValidationError) ErrorName() string {
	return "GetOrderTotalResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetOrderTotalResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOrderTotalResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOrderTotalResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOrderTotalResponseValidationError{}
