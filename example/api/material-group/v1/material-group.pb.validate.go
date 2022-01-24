// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/material-group/v1/material-group.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on SaveMaterialGroupRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SaveMaterialGroupRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SaveMaterialGroupRequestValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for OrgId

	if v, ok := interface{}(m.GetUserId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SaveMaterialGroupRequestValidationError{
				field:  "UserId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return SaveMaterialGroupRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	if v, ok := interface{}(m.GetType()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SaveMaterialGroupRequestValidationError{
				field:  "Type",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetScope()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SaveMaterialGroupRequestValidationError{
				field:  "Scope",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Order

	return nil
}

// SaveMaterialGroupRequestValidationError is the validation error returned by
// SaveMaterialGroupRequest.Validate if the designated constraints aren't met.
type SaveMaterialGroupRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SaveMaterialGroupRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SaveMaterialGroupRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SaveMaterialGroupRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SaveMaterialGroupRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SaveMaterialGroupRequestValidationError) ErrorName() string {
	return "SaveMaterialGroupRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SaveMaterialGroupRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSaveMaterialGroupRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SaveMaterialGroupRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SaveMaterialGroupRequestValidationError{}

// Validate checks the field values on MaterialGroupModel with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MaterialGroupModel) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Type

	// no validation rules for Order

	return nil
}

// MaterialGroupModelValidationError is the validation error returned by
// MaterialGroupModel.Validate if the designated constraints aren't met.
type MaterialGroupModelValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MaterialGroupModelValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MaterialGroupModelValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MaterialGroupModelValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MaterialGroupModelValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MaterialGroupModelValidationError) ErrorName() string {
	return "MaterialGroupModelValidationError"
}

// Error satisfies the builtin error interface
func (e MaterialGroupModelValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMaterialGroupModel.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MaterialGroupModelValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MaterialGroupModelValidationError{}

// Validate checks the field values on MaterialGroupModelList with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MaterialGroupModelList) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MaterialGroupModelListValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// MaterialGroupModelListValidationError is the validation error returned by
// MaterialGroupModelList.Validate if the designated constraints aren't met.
type MaterialGroupModelListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MaterialGroupModelListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MaterialGroupModelListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MaterialGroupModelListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MaterialGroupModelListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MaterialGroupModelListValidationError) ErrorName() string {
	return "MaterialGroupModelListValidationError"
}

// Error satisfies the builtin error interface
func (e MaterialGroupModelListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMaterialGroupModelList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MaterialGroupModelListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MaterialGroupModelListValidationError{}
