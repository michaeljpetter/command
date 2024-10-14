package internal

import (
	"errors"
	"github.com/michaeljpetter/command/value"
	"strconv"
)

var (
	errParse = errors.New("parse error")
	errRange = errors.New("value out of range")
)

func numError(err error) error {
	if num, ok := err.(*strconv.NumError); ok {
		switch num.Err {
		case strconv.ErrSyntax:
			return errParse
		case strconv.ErrRange:
			return errRange
		}
	}

	return err
}

type constraint[T any] []value.CheckFunc[T]

func (c constraint[T]) check(value T) error {
	for _, check := range c {
		if err := check(value); err != nil {
			return err
		}
	}

	return nil
}

type Value[T any] struct {
	value    *T
	required bool
	constraint[T]
}

func newValue[T any](defValue *T, value *T, constraint constraint[T]) Value[T] {
	if defValue != nil {
		*value = *defValue
	}
	return Value[T]{value, defValue == nil, constraint}
}

func (v Value[_]) Get() any {
	return *v.value
}

func (v Value[_]) Required() bool {
	return v.required
}
