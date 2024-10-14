// Package check implements common checks that can be
// used for flag and positional parameter values.
package check

import (
	"cmp"
	"errors"
	"fmt"
	"github.com/michaeljpetter/command/value"
	"slices"
	"strings"
)

// GreaterThan checks that a value is greater than a given minimum.
func GreaterThan[T cmp.Ordered](min T) value.CheckFunc[T] {
	return func(value T) error {
		if min < value {
			return nil
		}
		return fmt.Errorf("must be greater than %v", min)
	}
}

// LessThan checks that a value is less than a given maximum.
func LessThan[T cmp.Ordered](max T) value.CheckFunc[T] {
	return func(value T) error {
		if value < max {
			return nil
		}
		return fmt.Errorf("must be less than %v", max)
	}
}

// AtLeast checks that a value is greater than or equal to a given minimum.
func AtLeast[T cmp.Ordered](min T) value.CheckFunc[T] {
	return func(value T) error {
		if min <= value {
			return nil
		}
		return fmt.Errorf("must be at least %v", min)
	}
}

// AtMost checks that a value is less than or equal to a given maximum.
func AtMost[T cmp.Ordered](max T) value.CheckFunc[T] {
	return func(value T) error {
		if value <= max {
			return nil
		}
		return fmt.Errorf("must be at most %v", max)
	}
}

// OneOf checks that a value is present in a given list of allowed options.
func OneOf[T comparable](options ...T) value.CheckFunc[T] {
	return func(value T) error {
		if slices.Contains(options, value) {
			return nil
		}
		return fmt.Errorf("must be one of %v", options)
	}
}

// NotBlank checks that a string contains at least one non-white-space character.
func NotBlank(value string) error {
	if 0 < len(strings.TrimSpace(value)) {
		return nil
	}
	return errors.New("cannot be blank")
}
