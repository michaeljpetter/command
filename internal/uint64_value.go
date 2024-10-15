package internal

import (
	"github.com/michaeljpetter/command/value"
	"github.com/michaeljpetter/ptr"
	"strconv"
)

type Uint64Value struct{ Value[uint64] }

func NewUint64Value(defValue *uint64, value *uint64, checks ...value.CheckFunc[uint64]) Uint64Value {
	return Uint64Value{newValue(defValue, value, checks)}
}

func (u Uint64Value) Set(raw string) error {
	parsed, err := strconv.ParseUint(raw, 0, 64)
	*u.value = parsed

	if err != nil {
		return numError(err)
	}

	return u.check(parsed)
}

func (u Uint64Value) String() string {
	return strconv.FormatUint(*ptr.OrZero(u.value), 10)
}
