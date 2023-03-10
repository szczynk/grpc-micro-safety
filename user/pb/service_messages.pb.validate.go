// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: service_messages.proto

package pb

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

// Validate checks the field values on Service with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Service) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Service with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ServiceMultiError, or nil if none found.
func (m *Service) ValidateAll() error {
	return m.validate(true)
}

func (m *Service) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Service

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServiceValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServiceValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServiceValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServiceValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServiceValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServiceValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ServiceMultiError(errors)
	}

	return nil
}

// ServiceMultiError is an error wrapping multiple validation errors returned
// by Service.ValidateAll() if the designated constraints aren't met.
type ServiceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServiceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServiceMultiError) AllErrors() []error { return m }

// ServiceValidationError is the validation error returned by Service.Validate
// if the designated constraints aren't met.
type ServiceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceValidationError) ErrorName() string { return "ServiceValidationError" }

// Error satisfies the builtin error interface
func (e ServiceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sService.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceValidationError{}

// Validate checks the field values on CreateServiceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateServiceRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateServiceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateServiceRequestMultiError, or nil if none found.
func (m *CreateServiceRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateServiceRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if !strings.HasPrefix(m.GetService(), "pb.") {
		err := CreateServiceRequestValidationError{
			field:  "Service",
			reason: "value does not have prefix \"pb.\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateServiceRequestMultiError(errors)
	}

	return nil
}

// CreateServiceRequestMultiError is an error wrapping multiple validation
// errors returned by CreateServiceRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateServiceRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateServiceRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateServiceRequestMultiError) AllErrors() []error { return m }

// CreateServiceRequestValidationError is the validation error returned by
// CreateServiceRequest.Validate if the designated constraints aren't met.
type CreateServiceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateServiceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateServiceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateServiceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateServiceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateServiceRequestValidationError) ErrorName() string {
	return "CreateServiceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateServiceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateServiceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateServiceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateServiceRequestValidationError{}

// Validate checks the field values on CreateServiceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateServiceResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateServiceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateServiceResponseMultiError, or nil if none found.
func (m *CreateServiceResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateServiceResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetService()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateServiceResponseValidationError{
					field:  "Service",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateServiceResponseValidationError{
					field:  "Service",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateServiceResponseValidationError{
				field:  "Service",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateServiceResponseMultiError(errors)
	}

	return nil
}

// CreateServiceResponseMultiError is an error wrapping multiple validation
// errors returned by CreateServiceResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateServiceResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateServiceResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateServiceResponseMultiError) AllErrors() []error { return m }

// CreateServiceResponseValidationError is the validation error returned by
// CreateServiceResponse.Validate if the designated constraints aren't met.
type CreateServiceResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateServiceResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateServiceResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateServiceResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateServiceResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateServiceResponseValidationError) ErrorName() string {
	return "CreateServiceResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateServiceResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateServiceResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateServiceResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateServiceResponseValidationError{}

// Validate checks the field values on DeleteServiceByIdRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteServiceByIdRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteServiceByIdRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteServiceByIdRequestMultiError, or nil if none found.
func (m *DeleteServiceByIdRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteServiceByIdRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteServiceByIdRequestMultiError(errors)
	}

	return nil
}

// DeleteServiceByIdRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteServiceByIdRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteServiceByIdRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteServiceByIdRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteServiceByIdRequestMultiError) AllErrors() []error { return m }

// DeleteServiceByIdRequestValidationError is the validation error returned by
// DeleteServiceByIdRequest.Validate if the designated constraints aren't met.
type DeleteServiceByIdRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteServiceByIdRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteServiceByIdRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteServiceByIdRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteServiceByIdRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteServiceByIdRequestValidationError) ErrorName() string {
	return "DeleteServiceByIdRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteServiceByIdRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteServiceByIdRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteServiceByIdRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteServiceByIdRequestValidationError{}

// Validate checks the field values on DeleteServiceByIdResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteServiceByIdResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteServiceByIdResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteServiceByIdResponseMultiError, or nil if none found.
func (m *DeleteServiceByIdResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteServiceByIdResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if len(errors) > 0 {
		return DeleteServiceByIdResponseMultiError(errors)
	}

	return nil
}

// DeleteServiceByIdResponseMultiError is an error wrapping multiple validation
// errors returned by DeleteServiceByIdResponse.ValidateAll() if the
// designated constraints aren't met.
type DeleteServiceByIdResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteServiceByIdResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteServiceByIdResponseMultiError) AllErrors() []error { return m }

// DeleteServiceByIdResponseValidationError is the validation error returned by
// DeleteServiceByIdResponse.Validate if the designated constraints aren't met.
type DeleteServiceByIdResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteServiceByIdResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteServiceByIdResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteServiceByIdResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteServiceByIdResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteServiceByIdResponseValidationError) ErrorName() string {
	return "DeleteServiceByIdResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteServiceByIdResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteServiceByIdResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteServiceByIdResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteServiceByIdResponseValidationError{}

// Validate checks the field values on FindServicesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *FindServicesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FindServicesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FindServicesRequestMultiError, or nil if none found.
func (m *FindServicesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *FindServicesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Service

	// no validation rules for Page

	// no validation rules for Limit

	// no validation rules for Sort

	if len(errors) > 0 {
		return FindServicesRequestMultiError(errors)
	}

	return nil
}

// FindServicesRequestMultiError is an error wrapping multiple validation
// errors returned by FindServicesRequest.ValidateAll() if the designated
// constraints aren't met.
type FindServicesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FindServicesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FindServicesRequestMultiError) AllErrors() []error { return m }

// FindServicesRequestValidationError is the validation error returned by
// FindServicesRequest.Validate if the designated constraints aren't met.
type FindServicesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FindServicesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FindServicesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FindServicesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FindServicesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FindServicesRequestValidationError) ErrorName() string {
	return "FindServicesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e FindServicesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFindServicesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FindServicesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FindServicesRequestValidationError{}

// Validate checks the field values on FindServicesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *FindServicesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FindServicesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FindServicesResponseMultiError, or nil if none found.
func (m *FindServicesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *FindServicesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TotalCount

	// no validation rules for TotalPages

	// no validation rules for Page

	// no validation rules for Limit

	// no validation rules for HasMore

	for idx, item := range m.GetServices() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FindServicesResponseValidationError{
						field:  fmt.Sprintf("Services[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FindServicesResponseValidationError{
						field:  fmt.Sprintf("Services[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FindServicesResponseValidationError{
					field:  fmt.Sprintf("Services[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return FindServicesResponseMultiError(errors)
	}

	return nil
}

// FindServicesResponseMultiError is an error wrapping multiple validation
// errors returned by FindServicesResponse.ValidateAll() if the designated
// constraints aren't met.
type FindServicesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FindServicesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FindServicesResponseMultiError) AllErrors() []error { return m }

// FindServicesResponseValidationError is the validation error returned by
// FindServicesResponse.Validate if the designated constraints aren't met.
type FindServicesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FindServicesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FindServicesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FindServicesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FindServicesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FindServicesResponseValidationError) ErrorName() string {
	return "FindServicesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e FindServicesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFindServicesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FindServicesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FindServicesResponseValidationError{}
