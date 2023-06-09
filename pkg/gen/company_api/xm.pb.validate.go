// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: xm.proto

package company_api

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
var _xm_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CreateCompanyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateCompanyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateCompanyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateCompanyRequestMultiError, or nil if none found.
func (m *CreateCompanyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateCompanyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if !_CreateCompanyRequest_Name_Pattern.MatchString(m.GetName()) {
		err := CreateCompanyRequestValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^[\\\\w,\\\\s-]{4,15}\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDescription()) > 3000 {
		err := CreateCompanyRequestValidationError{
			field:  "Description",
			reason: "value length must be at most 3000 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetEmployeesNum() <= 0 {
		err := CreateCompanyRequestValidationError{
			field:  "EmployeesNum",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetRegistered() != true {
		err := CreateCompanyRequestValidationError{
			field:  "Registered",
			reason: "value must equal true",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _CreateCompanyRequest_Type_NotInLookup[m.GetType()]; ok {
		err := CreateCompanyRequestValidationError{
			field:  "Type",
			reason: "value must not be in list [_]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := CompanyType_name[int32(m.GetType())]; !ok {
		err := CreateCompanyRequestValidationError{
			field:  "Type",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateCompanyRequestMultiError(errors)
	}

	return nil
}

// CreateCompanyRequestMultiError is an error wrapping multiple validation
// errors returned by CreateCompanyRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateCompanyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateCompanyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateCompanyRequestMultiError) AllErrors() []error { return m }

// CreateCompanyRequestValidationError is the validation error returned by
// CreateCompanyRequest.Validate if the designated constraints aren't met.
type CreateCompanyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCompanyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCompanyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCompanyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCompanyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCompanyRequestValidationError) ErrorName() string {
	return "CreateCompanyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCompanyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCompanyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCompanyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCompanyRequestValidationError{}

var _CreateCompanyRequest_Name_Pattern = regexp.MustCompile("^[\\w,\\s-]{4,15}")

var _CreateCompanyRequest_Type_NotInLookup = map[CompanyType]struct{}{
	0: {},
}

// Validate checks the field values on UpdateCompanyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateCompanyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateCompanyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateCompanyRequestMultiError, or nil if none found.
func (m *UpdateCompanyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateCompanyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = UpdateCompanyRequestValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.Name != nil {

		if !_UpdateCompanyRequest_Name_Pattern.MatchString(m.GetName()) {
			err := UpdateCompanyRequestValidationError{
				field:  "Name",
				reason: "value does not match regex pattern \"^[\\\\w,\\\\s-]{4,15}\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.Description != nil {

		if utf8.RuneCountInString(m.GetDescription()) > 3000 {
			err := UpdateCompanyRequestValidationError{
				field:  "Description",
				reason: "value length must be at most 3000 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.EmployeesNum != nil {

		if m.GetEmployeesNum() <= 0 {
			err := UpdateCompanyRequestValidationError{
				field:  "EmployeesNum",
				reason: "value must be greater than 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.Registered != nil {
		// no validation rules for Registered
	}

	if m.Type != nil {

		if _, ok := _UpdateCompanyRequest_Type_NotInLookup[m.GetType()]; ok {
			err := UpdateCompanyRequestValidationError{
				field:  "Type",
				reason: "value must not be in list [_]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if _, ok := CompanyType_name[int32(m.GetType())]; !ok {
			err := UpdateCompanyRequestValidationError{
				field:  "Type",
				reason: "value must be one of the defined enum values",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return UpdateCompanyRequestMultiError(errors)
	}

	return nil
}

func (m *UpdateCompanyRequest) _validateUuid(uuid string) error {
	if matched := _xm_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// UpdateCompanyRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateCompanyRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateCompanyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateCompanyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateCompanyRequestMultiError) AllErrors() []error { return m }

// UpdateCompanyRequestValidationError is the validation error returned by
// UpdateCompanyRequest.Validate if the designated constraints aren't met.
type UpdateCompanyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateCompanyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateCompanyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateCompanyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateCompanyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateCompanyRequestValidationError) ErrorName() string {
	return "UpdateCompanyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateCompanyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateCompanyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateCompanyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateCompanyRequestValidationError{}

var _UpdateCompanyRequest_Name_Pattern = regexp.MustCompile("^[\\w,\\s-]{4,15}")

var _UpdateCompanyRequest_Type_NotInLookup = map[CompanyType]struct{}{
	0: {},
}

// Validate checks the field values on RemoveCompanyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RemoveCompanyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoveCompanyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RemoveCompanyRequestMultiError, or nil if none found.
func (m *RemoveCompanyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoveCompanyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = RemoveCompanyRequestValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RemoveCompanyRequestMultiError(errors)
	}

	return nil
}

func (m *RemoveCompanyRequest) _validateUuid(uuid string) error {
	if matched := _xm_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// RemoveCompanyRequestMultiError is an error wrapping multiple validation
// errors returned by RemoveCompanyRequest.ValidateAll() if the designated
// constraints aren't met.
type RemoveCompanyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoveCompanyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoveCompanyRequestMultiError) AllErrors() []error { return m }

// RemoveCompanyRequestValidationError is the validation error returned by
// RemoveCompanyRequest.Validate if the designated constraints aren't met.
type RemoveCompanyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveCompanyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveCompanyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveCompanyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveCompanyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveCompanyRequestValidationError) ErrorName() string {
	return "RemoveCompanyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveCompanyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveCompanyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveCompanyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveCompanyRequestValidationError{}

// Validate checks the field values on GetCompanyRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetCompanyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCompanyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCompanyRequestMultiError, or nil if none found.
func (m *GetCompanyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCompanyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = GetCompanyRequestValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetCompanyRequestMultiError(errors)
	}

	return nil
}

func (m *GetCompanyRequest) _validateUuid(uuid string) error {
	if matched := _xm_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetCompanyRequestMultiError is an error wrapping multiple validation errors
// returned by GetCompanyRequest.ValidateAll() if the designated constraints
// aren't met.
type GetCompanyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCompanyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCompanyRequestMultiError) AllErrors() []error { return m }

// GetCompanyRequestValidationError is the validation error returned by
// GetCompanyRequest.Validate if the designated constraints aren't met.
type GetCompanyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCompanyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCompanyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCompanyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCompanyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCompanyRequestValidationError) ErrorName() string {
	return "GetCompanyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetCompanyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCompanyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCompanyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCompanyRequestValidationError{}

// Validate checks the field values on Company with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Company) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Company with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CompanyMultiError, or nil if none found.
func (m *Company) ValidateAll() error {
	return m.validate(true)
}

func (m *Company) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for EmployeesNum

	// no validation rules for Registered

	// no validation rules for Type

	if m.Description != nil {
		// no validation rules for Description
	}

	if len(errors) > 0 {
		return CompanyMultiError(errors)
	}

	return nil
}

// CompanyMultiError is an error wrapping multiple validation errors returned
// by Company.ValidateAll() if the designated constraints aren't met.
type CompanyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CompanyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CompanyMultiError) AllErrors() []error { return m }

// CompanyValidationError is the validation error returned by Company.Validate
// if the designated constraints aren't met.
type CompanyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CompanyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CompanyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CompanyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CompanyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CompanyValidationError) ErrorName() string { return "CompanyValidationError" }

// Error satisfies the builtin error interface
func (e CompanyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompany.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CompanyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CompanyValidationError{}

// Validate checks the field values on LoginRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginRequestMultiError, or
// nil if none found.
func (m *LoginRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetUsername()) < 8 {
		err := LoginRequestValidationError{
			field:  "Username",
			reason: "value length must be at least 8 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPassword()) < 8 {
		err := LoginRequestValidationError{
			field:  "Password",
			reason: "value length must be at least 8 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return LoginRequestMultiError(errors)
	}

	return nil
}

// LoginRequestMultiError is an error wrapping multiple validation errors
// returned by LoginRequest.ValidateAll() if the designated constraints aren't met.
type LoginRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginRequestMultiError) AllErrors() []error { return m }

// LoginRequestValidationError is the validation error returned by
// LoginRequest.Validate if the designated constraints aren't met.
type LoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginRequestValidationError) ErrorName() string { return "LoginRequestValidationError" }

// Error satisfies the builtin error interface
func (e LoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginRequestValidationError{}

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

	// no validation rules for Id

	// no validation rules for AuthToken

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
