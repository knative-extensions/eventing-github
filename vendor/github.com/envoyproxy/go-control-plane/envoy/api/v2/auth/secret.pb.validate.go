// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/api/v2/auth/secret.proto

package envoy_api_v2_auth

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

// Validate checks the field values on GenericSecret with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GenericSecret) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetSecret()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GenericSecretValidationError{
				field:  "Secret",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GenericSecretValidationError is the validation error returned by
// GenericSecret.Validate if the designated constraints aren't met.
type GenericSecretValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GenericSecretValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GenericSecretValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GenericSecretValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GenericSecretValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GenericSecretValidationError) ErrorName() string { return "GenericSecretValidationError" }

// Error satisfies the builtin error interface
func (e GenericSecretValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGenericSecret.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GenericSecretValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GenericSecretValidationError{}

// Validate checks the field values on SdsSecretConfig with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *SdsSecretConfig) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	if v, ok := interface{}(m.GetSdsConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SdsSecretConfigValidationError{
				field:  "SdsConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SdsSecretConfigValidationError is the validation error returned by
// SdsSecretConfig.Validate if the designated constraints aren't met.
type SdsSecretConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SdsSecretConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SdsSecretConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SdsSecretConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SdsSecretConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SdsSecretConfigValidationError) ErrorName() string { return "SdsSecretConfigValidationError" }

// Error satisfies the builtin error interface
func (e SdsSecretConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSdsSecretConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SdsSecretConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SdsSecretConfigValidationError{}

// Validate checks the field values on Secret with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Secret) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	switch m.Type.(type) {

	case *Secret_TlsCertificate:

		if v, ok := interface{}(m.GetTlsCertificate()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SecretValidationError{
					field:  "TlsCertificate",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Secret_SessionTicketKeys:

		if v, ok := interface{}(m.GetSessionTicketKeys()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SecretValidationError{
					field:  "SessionTicketKeys",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Secret_ValidationContext:

		if v, ok := interface{}(m.GetValidationContext()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SecretValidationError{
					field:  "ValidationContext",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Secret_GenericSecret:

		if v, ok := interface{}(m.GetGenericSecret()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SecretValidationError{
					field:  "GenericSecret",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// SecretValidationError is the validation error returned by Secret.Validate if
// the designated constraints aren't met.
type SecretValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SecretValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SecretValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SecretValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SecretValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SecretValidationError) ErrorName() string { return "SecretValidationError" }

// Error satisfies the builtin error interface
func (e SecretValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSecret.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SecretValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SecretValidationError{}
