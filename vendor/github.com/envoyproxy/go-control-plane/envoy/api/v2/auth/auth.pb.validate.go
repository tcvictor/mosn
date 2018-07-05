// Code generated by protoc-gen-validate
// source: envoy/api/v2/auth/auth.proto
// DO NOT EDIT!!!

package auth

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

// Validate checks the field values on AuthAction with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *AuthAction) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ActionType

	for idx, item := range m.GetRules() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AuthActionValidationError{
					Field:  fmt.Sprintf("Rules[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// AuthActionValidationError is the validation error returned by
// AuthAction.Validate if the designated constraints aren't met.
type AuthActionValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e AuthActionValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthAction.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = AuthActionValidationError{}

// Validate checks the field values on AuthAction_AndRule with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AuthAction_AndRule) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRules() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AuthAction_AndRuleValidationError{
					Field:  fmt.Sprintf("Rules[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// AuthAction_AndRuleValidationError is the validation error returned by
// AuthAction_AndRule.Validate if the designated constraints aren't met.
type AuthAction_AndRuleValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e AuthAction_AndRuleValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthAction_AndRule.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = AuthAction_AndRuleValidationError{}

// Validate checks the field values on AuthAction_OrRule with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *AuthAction_OrRule) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRules() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AuthAction_OrRuleValidationError{
					Field:  fmt.Sprintf("Rules[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// AuthAction_OrRuleValidationError is the validation error returned by
// AuthAction_OrRule.Validate if the designated constraints aren't met.
type AuthAction_OrRuleValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e AuthAction_OrRuleValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthAction_OrRule.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = AuthAction_OrRuleValidationError{}

// Validate checks the field values on AuthAction_X509Rule with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AuthAction_X509Rule) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetValidationContext()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AuthAction_X509RuleValidationError{
				Field:  "ValidationContext",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// AuthAction_X509RuleValidationError is the validation error returned by
// AuthAction_X509Rule.Validate if the designated constraints aren't met.
type AuthAction_X509RuleValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e AuthAction_X509RuleValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthAction_X509Rule.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = AuthAction_X509RuleValidationError{}

// Validate checks the field values on AuthAction_Rule with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *AuthAction_Rule) Validate() error {
	if m == nil {
		return nil
	}

	switch m.RuleSpecifier.(type) {

	case *AuthAction_Rule_AndRule:

		if v, ok := interface{}(m.GetAndRule()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AuthAction_RuleValidationError{
					Field:  "AndRule",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *AuthAction_Rule_OrRule:

		if v, ok := interface{}(m.GetOrRule()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AuthAction_RuleValidationError{
					Field:  "OrRule",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *AuthAction_Rule_X509Rule:

		if v, ok := interface{}(m.GetX509Rule()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AuthAction_RuleValidationError{
					Field:  "X509Rule",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// AuthAction_RuleValidationError is the validation error returned by
// AuthAction_Rule.Validate if the designated constraints aren't met.
type AuthAction_RuleValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e AuthAction_RuleValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthAction_Rule.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = AuthAction_RuleValidationError{}