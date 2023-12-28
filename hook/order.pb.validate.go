// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/hook/order.proto

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

// Validate checks the field values on OrderStatusesHookParams with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OrderStatusesHookParams) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderStatusesHookParams with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OrderStatusesHookParamsMultiError, or nil if none found.
func (m *OrderStatusesHookParams) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderStatusesHookParams) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetOrderStatuses() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, OrderStatusesHookParamsValidationError{
						field:  fmt.Sprintf("OrderStatuses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, OrderStatusesHookParamsValidationError{
						field:  fmt.Sprintf("OrderStatuses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OrderStatusesHookParamsValidationError{
					field:  fmt.Sprintf("OrderStatuses[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return OrderStatusesHookParamsMultiError(errors)
	}

	return nil
}

// OrderStatusesHookParamsMultiError is an error wrapping multiple validation
// errors returned by OrderStatusesHookParams.ValidateAll() if the designated
// constraints aren't met.
type OrderStatusesHookParamsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderStatusesHookParamsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderStatusesHookParamsMultiError) AllErrors() []error { return m }

// OrderStatusesHookParamsValidationError is the validation error returned by
// OrderStatusesHookParams.Validate if the designated constraints aren't met.
type OrderStatusesHookParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderStatusesHookParamsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderStatusesHookParamsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderStatusesHookParamsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderStatusesHookParamsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderStatusesHookParamsValidationError) ErrorName() string {
	return "OrderStatusesHookParamsValidationError"
}

// Error satisfies the builtin error interface
func (e OrderStatusesHookParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderStatusesHookParams.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderStatusesHookParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderStatusesHookParamsValidationError{}

// Validate checks the field values on OrderItemStatusesHookParams with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OrderItemStatusesHookParams) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderItemStatusesHookParams with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OrderItemStatusesHookParamsMultiError, or nil if none found.
func (m *OrderItemStatusesHookParams) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderItemStatusesHookParams) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetOrderItemStatuses() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, OrderItemStatusesHookParamsValidationError{
						field:  fmt.Sprintf("OrderItemStatuses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, OrderItemStatusesHookParamsValidationError{
						field:  fmt.Sprintf("OrderItemStatuses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OrderItemStatusesHookParamsValidationError{
					field:  fmt.Sprintf("OrderItemStatuses[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return OrderItemStatusesHookParamsMultiError(errors)
	}

	return nil
}

// OrderItemStatusesHookParamsMultiError is an error wrapping multiple
// validation errors returned by OrderItemStatusesHookParams.ValidateAll() if
// the designated constraints aren't met.
type OrderItemStatusesHookParamsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderItemStatusesHookParamsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderItemStatusesHookParamsMultiError) AllErrors() []error { return m }

// OrderItemStatusesHookParamsValidationError is the validation error returned
// by OrderItemStatusesHookParams.Validate if the designated constraints
// aren't met.
type OrderItemStatusesHookParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderItemStatusesHookParamsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderItemStatusesHookParamsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderItemStatusesHookParamsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderItemStatusesHookParamsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderItemStatusesHookParamsValidationError) ErrorName() string {
	return "OrderItemStatusesHookParamsValidationError"
}

// Error satisfies the builtin error interface
func (e OrderItemStatusesHookParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderItemStatusesHookParams.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderItemStatusesHookParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderItemStatusesHookParamsValidationError{}

// Validate checks the field values on OrderCreateHookParams with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OrderCreateHookParams) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderCreateHookParams with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OrderCreateHookParamsMultiError, or nil if none found.
func (m *OrderCreateHookParams) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderCreateHookParams) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetOrder()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OrderCreateHookParamsValidationError{
					field:  "Order",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OrderCreateHookParamsValidationError{
					field:  "Order",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOrder()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrderCreateHookParamsValidationError{
				field:  "Order",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OrderCreateHookParamsValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OrderCreateHookParamsValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrderCreateHookParamsValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPayment()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OrderCreateHookParamsValidationError{
					field:  "Payment",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OrderCreateHookParamsValidationError{
					field:  "Payment",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPayment()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrderCreateHookParamsValidationError{
				field:  "Payment",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetAddress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OrderCreateHookParamsValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OrderCreateHookParamsValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrderCreateHookParamsValidationError{
				field:  "Address",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return OrderCreateHookParamsMultiError(errors)
	}

	return nil
}

// OrderCreateHookParamsMultiError is an error wrapping multiple validation
// errors returned by OrderCreateHookParams.ValidateAll() if the designated
// constraints aren't met.
type OrderCreateHookParamsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderCreateHookParamsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderCreateHookParamsMultiError) AllErrors() []error { return m }

// OrderCreateHookParamsValidationError is the validation error returned by
// OrderCreateHookParams.Validate if the designated constraints aren't met.
type OrderCreateHookParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderCreateHookParamsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderCreateHookParamsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderCreateHookParamsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderCreateHookParamsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderCreateHookParamsValidationError) ErrorName() string {
	return "OrderCreateHookParamsValidationError"
}

// Error satisfies the builtin error interface
func (e OrderCreateHookParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderCreateHookParams.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderCreateHookParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderCreateHookParamsValidationError{}

// Validate checks the field values on OrderProcessHookParams with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OrderProcessHookParams) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderProcessHookParams with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OrderProcessHookParamsMultiError, or nil if none found.
func (m *OrderProcessHookParams) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderProcessHookParams) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetOrder()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OrderProcessHookParamsValidationError{
					field:  "Order",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OrderProcessHookParamsValidationError{
					field:  "Order",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOrder()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrderProcessHookParamsValidationError{
				field:  "Order",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OrderProcessHookParamsValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OrderProcessHookParamsValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrderProcessHookParamsValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPayment()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OrderProcessHookParamsValidationError{
					field:  "Payment",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OrderProcessHookParamsValidationError{
					field:  "Payment",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPayment()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrderProcessHookParamsValidationError{
				field:  "Payment",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetAddress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OrderProcessHookParamsValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OrderProcessHookParamsValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrderProcessHookParamsValidationError{
				field:  "Address",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return OrderProcessHookParamsMultiError(errors)
	}

	return nil
}

// OrderProcessHookParamsMultiError is an error wrapping multiple validation
// errors returned by OrderProcessHookParams.ValidateAll() if the designated
// constraints aren't met.
type OrderProcessHookParamsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderProcessHookParamsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderProcessHookParamsMultiError) AllErrors() []error { return m }

// OrderProcessHookParamsValidationError is the validation error returned by
// OrderProcessHookParams.Validate if the designated constraints aren't met.
type OrderProcessHookParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderProcessHookParamsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderProcessHookParamsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderProcessHookParamsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderProcessHookParamsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderProcessHookParamsValidationError) ErrorName() string {
	return "OrderProcessHookParamsValidationError"
}

// Error satisfies the builtin error interface
func (e OrderProcessHookParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderProcessHookParams.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderProcessHookParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderProcessHookParamsValidationError{}
