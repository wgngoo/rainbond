// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/service/discovery/v3/discovery.proto

package envoy_service_discovery_v3

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
var _discovery_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on DiscoveryRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DiscoveryRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for VersionInfo

	if v, ok := interface{}(m.GetNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DiscoveryRequestValidationError{
				field:  "Node",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for TypeUrl

	// no validation rules for ResponseNonce

	if v, ok := interface{}(m.GetErrorDetail()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DiscoveryRequestValidationError{
				field:  "ErrorDetail",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DiscoveryRequestValidationError is the validation error returned by
// DiscoveryRequest.Validate if the designated constraints aren't met.
type DiscoveryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DiscoveryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DiscoveryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DiscoveryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DiscoveryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DiscoveryRequestValidationError) ErrorName() string { return "DiscoveryRequestValidationError" }

// Error satisfies the builtin error interface
func (e DiscoveryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDiscoveryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DiscoveryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DiscoveryRequestValidationError{}

// Validate checks the field values on DiscoveryResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DiscoveryResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for VersionInfo

	for idx, item := range m.GetResources() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DiscoveryResponseValidationError{
					field:  fmt.Sprintf("Resources[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Canary

	// no validation rules for TypeUrl

	// no validation rules for Nonce

	if v, ok := interface{}(m.GetControlPlane()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DiscoveryResponseValidationError{
				field:  "ControlPlane",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DiscoveryResponseValidationError is the validation error returned by
// DiscoveryResponse.Validate if the designated constraints aren't met.
type DiscoveryResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DiscoveryResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DiscoveryResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DiscoveryResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DiscoveryResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DiscoveryResponseValidationError) ErrorName() string {
	return "DiscoveryResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DiscoveryResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDiscoveryResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DiscoveryResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DiscoveryResponseValidationError{}

// Validate checks the field values on DeltaDiscoveryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeltaDiscoveryRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeltaDiscoveryRequestValidationError{
				field:  "Node",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for TypeUrl

	// no validation rules for InitialResourceVersions

	// no validation rules for ResponseNonce

	if v, ok := interface{}(m.GetErrorDetail()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeltaDiscoveryRequestValidationError{
				field:  "ErrorDetail",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DeltaDiscoveryRequestValidationError is the validation error returned by
// DeltaDiscoveryRequest.Validate if the designated constraints aren't met.
type DeltaDiscoveryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeltaDiscoveryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeltaDiscoveryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeltaDiscoveryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeltaDiscoveryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeltaDiscoveryRequestValidationError) ErrorName() string {
	return "DeltaDiscoveryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeltaDiscoveryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeltaDiscoveryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeltaDiscoveryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeltaDiscoveryRequestValidationError{}

// Validate checks the field values on DeltaDiscoveryResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeltaDiscoveryResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for SystemVersionInfo

	for idx, item := range m.GetResources() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DeltaDiscoveryResponseValidationError{
					field:  fmt.Sprintf("Resources[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for TypeUrl

	// no validation rules for Nonce

	return nil
}

// DeltaDiscoveryResponseValidationError is the validation error returned by
// DeltaDiscoveryResponse.Validate if the designated constraints aren't met.
type DeltaDiscoveryResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeltaDiscoveryResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeltaDiscoveryResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeltaDiscoveryResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeltaDiscoveryResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeltaDiscoveryResponseValidationError) ErrorName() string {
	return "DeltaDiscoveryResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeltaDiscoveryResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeltaDiscoveryResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeltaDiscoveryResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeltaDiscoveryResponseValidationError{}

// Validate checks the field values on Resource with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Resource) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Version

	if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ResourceValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ResourceValidationError is the validation error returned by
// Resource.Validate if the designated constraints aren't met.
type ResourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResourceValidationError) ErrorName() string { return "ResourceValidationError" }

// Error satisfies the builtin error interface
func (e ResourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResourceValidationError{}
