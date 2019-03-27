// Code generated by protoc-gen-validate
// source: envoy/service/trace/v2/trace_service.proto
// DO NOT EDIT!!!

package v2

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on StreamTracesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StreamTracesResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// StreamTracesResponseValidationError is the validation error returned by
// StreamTracesResponse.Validate if the designated constraints aren't met.
type StreamTracesResponseValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e StreamTracesResponseValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamTracesResponse.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = StreamTracesResponseValidationError{}

// Validate checks the field values on StreamTracesMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StreamTracesMessage) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetIdentifier()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return StreamTracesMessageValidationError{
				Field:  "Identifier",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	for idx, item := range m.GetSpans() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return StreamTracesMessageValidationError{
					Field:  fmt.Sprintf("Spans[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// StreamTracesMessageValidationError is the validation error returned by
// StreamTracesMessage.Validate if the designated constraints aren't met.
type StreamTracesMessageValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e StreamTracesMessageValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamTracesMessage.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = StreamTracesMessageValidationError{}

// Validate checks the field values on StreamTracesMessage_Identifier with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StreamTracesMessage_Identifier) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetNode() == nil {
		return StreamTracesMessage_IdentifierValidationError{
			Field:  "Node",
			Reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetNode()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return StreamTracesMessage_IdentifierValidationError{
				Field:  "Node",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// StreamTracesMessage_IdentifierValidationError is the validation error
// returned by StreamTracesMessage_Identifier.Validate if the designated
// constraints aren't met.
type StreamTracesMessage_IdentifierValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e StreamTracesMessage_IdentifierValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamTracesMessage_Identifier.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = StreamTracesMessage_IdentifierValidationError{}
