package internal

import (
	"github.com/michaeljpetter/command/value"
	"github.com/michaeljpetter/ptr"
	"strconv"
)

type UintValue struct{ Value[uint] }

func NewUintValue(defValue *uint, value *uint, checks ...value.CheckFunc[uint]) UintValue {
	return UintValue{newValue(defValue, value, checks)}
}

func (u UintValue) Set(raw string) error {
	parsed, err := strconv.ParseUint(raw, 0, strconv.IntSize)
	*u.value = uint(parsed)

	if err != nil {
		return numError(err)
	}

	return u.check(uint(parsed))
}

func (u UintValue) String() string {
	return strconv.FormatUint(uint64(*ptr.OrZero(u.value)), 10)
}
