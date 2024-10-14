package internal

import (
	"fmt"
	"github.com/michaeljpetter/command/value"
	"github.com/michaeljpetter/ptr"
)

type StringValue struct{ Value[string] }

func NewStringValue(defValue *string, value *string, checks ...value.CheckFunc[string]) StringValue {
	return StringValue{newValue(defValue, value, checks)}
}

func (s StringValue) Set(raw string) error {
	*s.value = raw

	return s.check(raw)
}

func (s StringValue) String() string {
	return fmt.Sprintf("%q", *ptr.OrZero(s.value))
}
