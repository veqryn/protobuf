// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: veqryn/protobuf/ptypes/test/test.proto

package test

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _test_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on TimestampReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *TimestampReq) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetMyTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TimestampReqValidationError{
				field:  "MyTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// TimestampReqValidationError is the validation error returned by
// TimestampReq.Validate if the designated constraints aren't met.
type TimestampReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TimestampReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TimestampReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TimestampReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TimestampReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TimestampReqValidationError) ErrorName() string { return "TimestampReqValidationError" }

// Error satisfies the builtin error interface
func (e TimestampReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTimestampReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TimestampReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TimestampReqValidationError{}

// Validate checks the field values on TimestampResp with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *TimestampResp) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetMyTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TimestampRespValidationError{
				field:  "MyTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// TimestampRespValidationError is the validation error returned by
// TimestampResp.Validate if the designated constraints aren't met.
type TimestampRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TimestampRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TimestampRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TimestampRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TimestampRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TimestampRespValidationError) ErrorName() string { return "TimestampRespValidationError" }

// Error satisfies the builtin error interface
func (e TimestampRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTimestampResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TimestampRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TimestampRespValidationError{}