package internal

import (
	"github.com/michaeljpetter/command/value"
	"github.com/michaeljpetter/ptr"
	"strconv"
)

type Int64Value struct{ Value[int64] }

func NewInt64Value(defValue *int64, value *int64, checks ...value.CheckFunc[int64]) Int64Value {
	return Int64Value{newValue(defValue, value, checks)}
}

func (i Int64Value) Set(raw string) error {
	parsed, err := strconv.ParseInt(raw, 0, 64)
	*i.value = parsed

	if err != nil {
		return numError(err)
	}

	return i.check(parsed)
}

func (i Int64Value) String() string {
	return strconv.FormatInt(*ptr.OrZero(i.value), 10)
}
