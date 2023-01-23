// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: workspace_messages.proto

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

// define the regex for a UUID once up-front
var _workspace_messages_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Workspace with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Workspace) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Workspace with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in WorkspaceMultiError, or nil
// if none found.
func (m *Workspace) ValidateAll() error {
	return m.validate(true)
}

func (m *Workspace) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OfficeId

	// no validation rules for UserId

	if len(errors) > 0 {
		return WorkspaceMultiError(errors)
	}

	return nil
}

// WorkspaceMultiError is an error wrapping multiple validation errors returned
// by Workspace.ValidateAll() if the designated constraints aren't met.
type WorkspaceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WorkspaceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WorkspaceMultiError) AllErrors() []error { return m }

// WorkspaceValidationError is the validation error returned by
// Workspace.Validate if the designated constraints aren't met.
type WorkspaceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkspaceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkspaceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkspaceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkspaceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkspaceValidationError) ErrorName() string { return "WorkspaceValidationError" }

// Error satisfies the builtin error interface
func (e WorkspaceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkspace.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkspaceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkspaceValidationError{}

// Validate checks the field values on CreateWorkspaceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateWorkspaceRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateWorkspaceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateWorkspaceRequestMultiError, or nil if none found.
func (m *CreateWorkspaceRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateWorkspaceRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetOfficeId() <= 0 {
		err := CreateWorkspaceRequestValidationError{
			field:  "OfficeId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if err := m._validateUuid(m.GetUserId()); err != nil {
		err = CreateWorkspaceRequestValidationError{
			field:  "UserId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateWorkspaceRequestMultiError(errors)
	}

	return nil
}

func (m *CreateWorkspaceRequest) _validateUuid(uuid string) error {
	if matched := _workspace_messages_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// CreateWorkspaceRequestMultiError is an error wrapping multiple validation
// errors returned by CreateWorkspaceRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateWorkspaceRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateWorkspaceRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateWorkspaceRequestMultiError) AllErrors() []error { return m }

// CreateWorkspaceRequestValidationError is the validation error returned by
// CreateWorkspaceRequest.Validate if the designated constraints aren't met.
type CreateWorkspaceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateWorkspaceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateWorkspaceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateWorkspaceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateWorkspaceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateWorkspaceRequestValidationError) ErrorName() string {
	return "CreateWorkspaceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateWorkspaceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateWorkspaceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateWorkspaceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateWorkspaceRequestValidationError{}

// Validate checks the field values on CreateWorkspaceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateWorkspaceResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateWorkspaceResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateWorkspaceResponseMultiError, or nil if none found.
func (m *CreateWorkspaceResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateWorkspaceResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetWorkspace()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateWorkspaceResponseValidationError{
					field:  "Workspace",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateWorkspaceResponseValidationError{
					field:  "Workspace",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWorkspace()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateWorkspaceResponseValidationError{
				field:  "Workspace",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateWorkspaceResponseMultiError(errors)
	}

	return nil
}

// CreateWorkspaceResponseMultiError is an error wrapping multiple validation
// errors returned by CreateWorkspaceResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateWorkspaceResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateWorkspaceResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateWorkspaceResponseMultiError) AllErrors() []error { return m }

// CreateWorkspaceResponseValidationError is the validation error returned by
// CreateWorkspaceResponse.Validate if the designated constraints aren't met.
type CreateWorkspaceResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateWorkspaceResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateWorkspaceResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateWorkspaceResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateWorkspaceResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateWorkspaceResponseValidationError) ErrorName() string {
	return "CreateWorkspaceResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateWorkspaceResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateWorkspaceResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateWorkspaceResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateWorkspaceResponseValidationError{}

// Validate checks the field values on DeleteWorkspaceByIdRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteWorkspaceByIdRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteWorkspaceByIdRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteWorkspaceByIdRequestMultiError, or nil if none found.
func (m *DeleteWorkspaceByIdRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteWorkspaceByIdRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetUserId()); err != nil {
		err = DeleteWorkspaceByIdRequestValidationError{
			field:  "UserId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteWorkspaceByIdRequestMultiError(errors)
	}

	return nil
}

func (m *DeleteWorkspaceByIdRequest) _validateUuid(uuid string) error {
	if matched := _workspace_messages_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// DeleteWorkspaceByIdRequestMultiError is an error wrapping multiple
// validation errors returned by DeleteWorkspaceByIdRequest.ValidateAll() if
// the designated constraints aren't met.
type DeleteWorkspaceByIdRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteWorkspaceByIdRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteWorkspaceByIdRequestMultiError) AllErrors() []error { return m }

// DeleteWorkspaceByIdRequestValidationError is the validation error returned
// by DeleteWorkspaceByIdRequest.Validate if the designated constraints aren't met.
type DeleteWorkspaceByIdRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteWorkspaceByIdRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteWorkspaceByIdRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteWorkspaceByIdRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteWorkspaceByIdRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteWorkspaceByIdRequestValidationError) ErrorName() string {
	return "DeleteWorkspaceByIdRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteWorkspaceByIdRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteWorkspaceByIdRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteWorkspaceByIdRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteWorkspaceByIdRequestValidationError{}

// Validate checks the field values on DeleteWorkspaceByIdResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteWorkspaceByIdResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteWorkspaceByIdResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteWorkspaceByIdResponseMultiError, or nil if none found.
func (m *DeleteWorkspaceByIdResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteWorkspaceByIdResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if len(errors) > 0 {
		return DeleteWorkspaceByIdResponseMultiError(errors)
	}

	return nil
}

// DeleteWorkspaceByIdResponseMultiError is an error wrapping multiple
// validation errors returned by DeleteWorkspaceByIdResponse.ValidateAll() if
// the designated constraints aren't met.
type DeleteWorkspaceByIdResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteWorkspaceByIdResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteWorkspaceByIdResponseMultiError) AllErrors() []error { return m }

// DeleteWorkspaceByIdResponseValidationError is the validation error returned
// by DeleteWorkspaceByIdResponse.Validate if the designated constraints
// aren't met.
type DeleteWorkspaceByIdResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteWorkspaceByIdResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteWorkspaceByIdResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteWorkspaceByIdResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteWorkspaceByIdResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteWorkspaceByIdResponseValidationError) ErrorName() string {
	return "DeleteWorkspaceByIdResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteWorkspaceByIdResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteWorkspaceByIdResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteWorkspaceByIdResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteWorkspaceByIdResponseValidationError{}

// Validate checks the field values on FindWorkspacesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *FindWorkspacesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FindWorkspacesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FindWorkspacesRequestMultiError, or nil if none found.
func (m *FindWorkspacesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *FindWorkspacesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OfficeId

	// no validation rules for Username

	// no validation rules for Email

	// no validation rules for Role

	// no validation rules for Verified

	// no validation rules for Page

	// no validation rules for Limit

	// no validation rules for Sort

	if len(errors) > 0 {
		return FindWorkspacesRequestMultiError(errors)
	}

	return nil
}

// FindWorkspacesRequestMultiError is an error wrapping multiple validation
// errors returned by FindWorkspacesRequest.ValidateAll() if the designated
// constraints aren't met.
type FindWorkspacesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FindWorkspacesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FindWorkspacesRequestMultiError) AllErrors() []error { return m }

// FindWorkspacesRequestValidationError is the validation error returned by
// FindWorkspacesRequest.Validate if the designated constraints aren't met.
type FindWorkspacesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FindWorkspacesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FindWorkspacesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FindWorkspacesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FindWorkspacesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FindWorkspacesRequestValidationError) ErrorName() string {
	return "FindWorkspacesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e FindWorkspacesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFindWorkspacesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FindWorkspacesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FindWorkspacesRequestValidationError{}

// Validate checks the field values on FindWorkspacesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *FindWorkspacesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FindWorkspacesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FindWorkspacesResponseMultiError, or nil if none found.
func (m *FindWorkspacesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *FindWorkspacesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TotalCount

	// no validation rules for TotalPages

	// no validation rules for Page

	// no validation rules for Limit

	// no validation rules for HasMore

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FindWorkspacesResponseValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FindWorkspacesResponseValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FindWorkspacesResponseValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return FindWorkspacesResponseMultiError(errors)
	}

	return nil
}

// FindWorkspacesResponseMultiError is an error wrapping multiple validation
// errors returned by FindWorkspacesResponse.ValidateAll() if the designated
// constraints aren't met.
type FindWorkspacesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FindWorkspacesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FindWorkspacesResponseMultiError) AllErrors() []error { return m }

// FindWorkspacesResponseValidationError is the validation error returned by
// FindWorkspacesResponse.Validate if the designated constraints aren't met.
type FindWorkspacesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FindWorkspacesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FindWorkspacesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FindWorkspacesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FindWorkspacesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FindWorkspacesResponseValidationError) ErrorName() string {
	return "FindWorkspacesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e FindWorkspacesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFindWorkspacesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FindWorkspacesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FindWorkspacesResponseValidationError{}