// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/message/product.proto

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

// Validate checks the field values on ProductGetMessage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ProductGetMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProductGetMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProductGetMessageMultiError, or nil if none found.
func (m *ProductGetMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *ProductGetMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Id != nil {

		if m.GetId() <= 0 {
			err := ProductGetMessageValidationError{
				field:  "Id",
				reason: "value must be greater than 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.CustomId != nil {
		// no validation rules for CustomId
	}

	if m.WithVariants != nil {
		// no validation rules for WithVariants
	}

	if m.IgnoreCache != nil {
		// no validation rules for IgnoreCache
	}

	if len(errors) > 0 {
		return ProductGetMessageMultiError(errors)
	}

	return nil
}

// ProductGetMessageMultiError is an error wrapping multiple validation errors
// returned by ProductGetMessage.ValidateAll() if the designated constraints
// aren't met.
type ProductGetMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProductGetMessageMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProductGetMessageMultiError) AllErrors() []error { return m }

// ProductGetMessageValidationError is the validation error returned by
// ProductGetMessage.Validate if the designated constraints aren't met.
type ProductGetMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductGetMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductGetMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductGetMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductGetMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductGetMessageValidationError) ErrorName() string {
	return "ProductGetMessageValidationError"
}

// Error satisfies the builtin error interface
func (e ProductGetMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductGetMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductGetMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductGetMessageValidationError{}

// Validate checks the field values on ProductGetResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProductGetResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProductGetResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProductGetResponseMultiError, or nil if none found.
func (m *ProductGetResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ProductGetResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetProduct()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProductGetResponseValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProductGetResponseValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProduct()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProductGetResponseValidationError{
				field:  "Product",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ProductGetResponseMultiError(errors)
	}

	return nil
}

// ProductGetResponseMultiError is an error wrapping multiple validation errors
// returned by ProductGetResponse.ValidateAll() if the designated constraints
// aren't met.
type ProductGetResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProductGetResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProductGetResponseMultiError) AllErrors() []error { return m }

// ProductGetResponseValidationError is the validation error returned by
// ProductGetResponse.Validate if the designated constraints aren't met.
type ProductGetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductGetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductGetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductGetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductGetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductGetResponseValidationError) ErrorName() string {
	return "ProductGetResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ProductGetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductGetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductGetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductGetResponseValidationError{}

// Validate checks the field values on ProductListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProductListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProductListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProductListResponseMultiError, or nil if none found.
func (m *ProductListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ProductListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetProducts() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ProductListResponseValidationError{
						field:  fmt.Sprintf("Products[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProductListResponseValidationError{
						field:  fmt.Sprintf("Products[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProductListResponseValidationError{
					field:  fmt.Sprintf("Products[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Page != nil {
		// no validation rules for Page
	}

	if m.PageTotal != nil {
		// no validation rules for PageTotal
	}

	if len(errors) > 0 {
		return ProductListResponseMultiError(errors)
	}

	return nil
}

// ProductListResponseMultiError is an error wrapping multiple validation
// errors returned by ProductListResponse.ValidateAll() if the designated
// constraints aren't met.
type ProductListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProductListResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProductListResponseMultiError) AllErrors() []error { return m }

// ProductListResponseValidationError is the validation error returned by
// ProductListResponse.Validate if the designated constraints aren't met.
type ProductListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductListResponseValidationError) ErrorName() string {
	return "ProductListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ProductListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductListResponseValidationError{}

// Validate checks the field values on ProductListMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProductListMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProductListMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProductListMessageMultiError, or nil if none found.
func (m *ProductListMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *ProductListMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.CategoryId != nil {
		// no validation rules for CategoryId
	}

	if m.Type != nil {
		// no validation rules for Type
	}

	if m.SearchTerm != nil {
		// no validation rules for SearchTerm
	}

	if m.Page != nil {
		// no validation rules for Page
	}

	if m.Limit != nil {
		// no validation rules for Limit
	}

	if m.OrderBy != nil {
		// no validation rules for OrderBy
	}

	if m.IgnoreCache != nil {
		// no validation rules for IgnoreCache
	}

	if m.SessionUUID != nil {
		// no validation rules for SessionUUID
	}

	if m.TimeLimit != nil {
		// no validation rules for TimeLimit
	}

	if len(errors) > 0 {
		return ProductListMessageMultiError(errors)
	}

	return nil
}

// ProductListMessageMultiError is an error wrapping multiple validation errors
// returned by ProductListMessage.ValidateAll() if the designated constraints
// aren't met.
type ProductListMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProductListMessageMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProductListMessageMultiError) AllErrors() []error { return m }

// ProductListMessageValidationError is the validation error returned by
// ProductListMessage.Validate if the designated constraints aren't met.
type ProductListMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductListMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductListMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductListMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductListMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductListMessageValidationError) ErrorName() string {
	return "ProductListMessageValidationError"
}

// Error satisfies the builtin error interface
func (e ProductListMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductListMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductListMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductListMessageValidationError{}

// Validate checks the field values on ProductUpdateMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProductUpdateMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProductUpdateMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProductUpdateMessageMultiError, or nil if none found.
func (m *ProductUpdateMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *ProductUpdateMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Price

	// no validation rules for Image

	// no validation rules for Description

	// no validation rules for Excerpt

	// no validation rules for BrandId

	// no validation rules for Stock

	if len(errors) > 0 {
		return ProductUpdateMessageMultiError(errors)
	}

	return nil
}

// ProductUpdateMessageMultiError is an error wrapping multiple validation
// errors returned by ProductUpdateMessage.ValidateAll() if the designated
// constraints aren't met.
type ProductUpdateMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProductUpdateMessageMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProductUpdateMessageMultiError) AllErrors() []error { return m }

// ProductUpdateMessageValidationError is the validation error returned by
// ProductUpdateMessage.Validate if the designated constraints aren't met.
type ProductUpdateMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductUpdateMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductUpdateMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductUpdateMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductUpdateMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductUpdateMessageValidationError) ErrorName() string {
	return "ProductUpdateMessageValidationError"
}

// Error satisfies the builtin error interface
func (e ProductUpdateMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductUpdateMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductUpdateMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductUpdateMessageValidationError{}

// Validate checks the field values on ProductUpdateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProductUpdateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProductUpdateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProductUpdateResponseMultiError, or nil if none found.
func (m *ProductUpdateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ProductUpdateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetProduct()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProductUpdateResponseValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProductUpdateResponseValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProduct()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProductUpdateResponseValidationError{
				field:  "Product",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ProductUpdateResponseMultiError(errors)
	}

	return nil
}

// ProductUpdateResponseMultiError is an error wrapping multiple validation
// errors returned by ProductUpdateResponse.ValidateAll() if the designated
// constraints aren't met.
type ProductUpdateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProductUpdateResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProductUpdateResponseMultiError) AllErrors() []error { return m }

// ProductUpdateResponseValidationError is the validation error returned by
// ProductUpdateResponse.Validate if the designated constraints aren't met.
type ProductUpdateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductUpdateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductUpdateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductUpdateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductUpdateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductUpdateResponseValidationError) ErrorName() string {
	return "ProductUpdateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ProductUpdateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductUpdateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductUpdateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductUpdateResponseValidationError{}

// Validate checks the field values on ProductCreateMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProductCreateMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProductCreateMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProductCreateMessageMultiError, or nil if none found.
func (m *ProductCreateMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *ProductCreateMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Image

	// no validation rules for Price

	// no validation rules for Description

	// no validation rules for Excerpt

	// no validation rules for BrandId

	// no validation rules for Stock

	if len(errors) > 0 {
		return ProductCreateMessageMultiError(errors)
	}

	return nil
}

// ProductCreateMessageMultiError is an error wrapping multiple validation
// errors returned by ProductCreateMessage.ValidateAll() if the designated
// constraints aren't met.
type ProductCreateMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProductCreateMessageMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProductCreateMessageMultiError) AllErrors() []error { return m }

// ProductCreateMessageValidationError is the validation error returned by
// ProductCreateMessage.Validate if the designated constraints aren't met.
type ProductCreateMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductCreateMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductCreateMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductCreateMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductCreateMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductCreateMessageValidationError) ErrorName() string {
	return "ProductCreateMessageValidationError"
}

// Error satisfies the builtin error interface
func (e ProductCreateMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductCreateMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductCreateMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductCreateMessageValidationError{}

// Validate checks the field values on ProductCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProductCreateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProductCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProductCreateResponseMultiError, or nil if none found.
func (m *ProductCreateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ProductCreateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetProduct()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProductCreateResponseValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProductCreateResponseValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProduct()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProductCreateResponseValidationError{
				field:  "Product",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ProductCreateResponseMultiError(errors)
	}

	return nil
}

// ProductCreateResponseMultiError is an error wrapping multiple validation
// errors returned by ProductCreateResponse.ValidateAll() if the designated
// constraints aren't met.
type ProductCreateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProductCreateResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProductCreateResponseMultiError) AllErrors() []error { return m }

// ProductCreateResponseValidationError is the validation error returned by
// ProductCreateResponse.Validate if the designated constraints aren't met.
type ProductCreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductCreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductCreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductCreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductCreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductCreateResponseValidationError) ErrorName() string {
	return "ProductCreateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ProductCreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductCreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductCreateResponseValidationError{}

// Validate checks the field values on ProductDeleteMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProductDeleteMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProductDeleteMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProductDeleteMessageMultiError, or nil if none found.
func (m *ProductDeleteMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *ProductDeleteMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return ProductDeleteMessageMultiError(errors)
	}

	return nil
}

// ProductDeleteMessageMultiError is an error wrapping multiple validation
// errors returned by ProductDeleteMessage.ValidateAll() if the designated
// constraints aren't met.
type ProductDeleteMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProductDeleteMessageMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProductDeleteMessageMultiError) AllErrors() []error { return m }

// ProductDeleteMessageValidationError is the validation error returned by
// ProductDeleteMessage.Validate if the designated constraints aren't met.
type ProductDeleteMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductDeleteMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductDeleteMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductDeleteMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductDeleteMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductDeleteMessageValidationError) ErrorName() string {
	return "ProductDeleteMessageValidationError"
}

// Error satisfies the builtin error interface
func (e ProductDeleteMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductDeleteMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductDeleteMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductDeleteMessageValidationError{}

// Validate checks the field values on ProductDeleteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProductDeleteResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProductDeleteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProductDeleteResponseMultiError, or nil if none found.
func (m *ProductDeleteResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ProductDeleteResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if len(errors) > 0 {
		return ProductDeleteResponseMultiError(errors)
	}

	return nil
}

// ProductDeleteResponseMultiError is an error wrapping multiple validation
// errors returned by ProductDeleteResponse.ValidateAll() if the designated
// constraints aren't met.
type ProductDeleteResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProductDeleteResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProductDeleteResponseMultiError) AllErrors() []error { return m }

// ProductDeleteResponseValidationError is the validation error returned by
// ProductDeleteResponse.Validate if the designated constraints aren't met.
type ProductDeleteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductDeleteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductDeleteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductDeleteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductDeleteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductDeleteResponseValidationError) ErrorName() string {
	return "ProductDeleteResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ProductDeleteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductDeleteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductDeleteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductDeleteResponseValidationError{}
