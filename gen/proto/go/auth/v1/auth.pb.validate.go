// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: auth/v1/auth.proto

package authv1

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

// Validate checks the field values on CreateRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CreateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CreateRequestMultiError, or
// nil if none found.
func (m *CreateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetAuthUser() == nil {
		err := CreateRequestValidationError{
			field:  "AuthUser",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetAuthUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateRequestValidationError{
					field:  "AuthUser",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateRequestValidationError{
					field:  "AuthUser",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAuthUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateRequestValidationError{
				field:  "AuthUser",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateRequestMultiError(errors)
	}

	return nil
}

// CreateRequestMultiError is an error wrapping multiple validation errors
// returned by CreateRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateRequestMultiError) AllErrors() []error { return m }

// CreateRequestValidationError is the validation error returned by
// CreateRequest.Validate if the designated constraints aren't met.
type CreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRequestValidationError) ErrorName() string { return "CreateRequestValidationError" }

// Error satisfies the builtin error interface
func (e CreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRequestValidationError{}

// Validate checks the field values on CreateResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CreateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CreateResponseMultiError,
// or nil if none found.
func (m *CreateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetAuthUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateResponseValidationError{
					field:  "AuthUser",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateResponseValidationError{
					field:  "AuthUser",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAuthUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateResponseValidationError{
				field:  "AuthUser",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateResponseMultiError(errors)
	}

	return nil
}

// CreateResponseMultiError is an error wrapping multiple validation errors
// returned by CreateResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateResponseMultiError) AllErrors() []error { return m }

// CreateResponseValidationError is the validation error returned by
// CreateResponse.Validate if the designated constraints aren't met.
type CreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateResponseValidationError) ErrorName() string { return "CreateResponseValidationError" }

// Error satisfies the builtin error interface
func (e CreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateResponseValidationError{}

// Validate checks the field values on AuthUser with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AuthUser) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthUser with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AuthUserMultiError, or nil
// if none found.
func (m *AuthUser) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthUser) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for GivenName

	// no validation rules for FamilyName

	// no validation rules for Email

	// no validation rules for Password

	// no validation rules for Connection

	// no validation rules for ClientId

	if all {
		switch v := interface{}(m.GetUserMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AuthUserValidationError{
					field:  "UserMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AuthUserValidationError{
					field:  "UserMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUserMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AuthUserValidationError{
				field:  "UserMetadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AuthUserMultiError(errors)
	}

	return nil
}

// AuthUserMultiError is an error wrapping multiple validation errors returned
// by AuthUser.ValidateAll() if the designated constraints aren't met.
type AuthUserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthUserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthUserMultiError) AllErrors() []error { return m }

// AuthUserValidationError is the validation error returned by
// AuthUser.Validate if the designated constraints aren't met.
type AuthUserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthUserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthUserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthUserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthUserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthUserValidationError) ErrorName() string { return "AuthUserValidationError" }

// Error satisfies the builtin error interface
func (e AuthUserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthUserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthUserValidationError{}

// Validate checks the field values on UserMetadata with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserMetadata) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserMetadata with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserMetadataMultiError, or
// nil if none found.
func (m *UserMetadata) ValidateAll() error {
	return m.validate(true)
}

func (m *UserMetadata) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PatientId

	// no validation rules for Role

	if len(errors) > 0 {
		return UserMetadataMultiError(errors)
	}

	return nil
}

// UserMetadataMultiError is an error wrapping multiple validation errors
// returned by UserMetadata.ValidateAll() if the designated constraints aren't met.
type UserMetadataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMetadataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMetadataMultiError) AllErrors() []error { return m }

// UserMetadataValidationError is the validation error returned by
// UserMetadata.Validate if the designated constraints aren't met.
type UserMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserMetadataValidationError) ErrorName() string { return "UserMetadataValidationError" }

// Error satisfies the builtin error interface
func (e UserMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserMetadataValidationError{}
