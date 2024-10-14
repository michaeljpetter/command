package internal

import (
	"github.com/michaeljpetter/command/value"
	"github.com/michaeljpetter/ptr"
	"strconv"
)

type IntValue struct{ Value[int] }

func NewIntValue(defValue *int, value *int, checks ...value.CheckFunc[int]) IntValue {
	return IntValue{newValue(defValue, value, checks)}
}

func (i IntValue) Set(raw string) error {
	parsed, err := strconv.ParseInt(raw, 0, strconv.IntSize)
	*i.value = int(parsed)

	if err != nil {
		return numError(err)
	}

	return i.check(int(parsed))
}

func (i IntValue) String() string {
	return strconv.Itoa(*ptr.OrZero(i.value))
}
