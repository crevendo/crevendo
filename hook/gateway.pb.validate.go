// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/hook/gateway.proto

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

// Validate checks the field values on GatewayListHookParams with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GatewayListHookParams) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GatewayListHookParams with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GatewayListHookParamsMultiError, or nil if none found.
func (m *GatewayListHookParams) ValidateAll() error {
	return m.validate(true)
}

func (m *GatewayListHookParams) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserUUID

	for idx, item := range m.GetGateways() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GatewayListHookParamsValidationError{
						field:  fmt.Sprintf("Gateways[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GatewayListHookParamsValidationError{
						field:  fmt.Sprintf("Gateways[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GatewayListHookParamsValidationError{
					field:  fmt.Sprintf("Gateways[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GatewayListHookParamsMultiError(errors)
	}

	return nil
}

// GatewayListHookParamsMultiError is an error wrapping multiple validation
// errors returned by GatewayListHookParams.ValidateAll() if the designated
// constraints aren't met.
type GatewayListHookParamsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GatewayListHookParamsMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GatewayListHookParamsMultiError) AllErrors() []error { return m }

// GatewayListHookParamsValidationError is the validation error returned by
// GatewayListHookParams.Validate if the designated constraints aren't met.
type GatewayListHookParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GatewayListHookParamsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GatewayListHookParamsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GatewayListHookParamsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GatewayListHookParamsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GatewayListHookParamsValidationError) ErrorName() string {
	return "GatewayListHookParamsValidationError"
}

// Error satisfies the builtin error interface
func (e GatewayListHookParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGatewayListHookParams.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GatewayListHookParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GatewayListHookParamsValidationError{}
