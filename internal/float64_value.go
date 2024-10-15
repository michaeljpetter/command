package internal

import (
	"github.com/michaeljpetter/command/value"
	"github.com/michaeljpetter/ptr"
	"strconv"
)

type Float64Value struct{ Value[float64] }

func NewFloat64Value(defValue *float64, value *float64, checks ...value.CheckFunc[float64]) Float64Value {
	return Float64Value{newValue(defValue, value, checks)}
}

func (f Float64Value) Set(raw string) error {
	parsed, err := strconv.ParseFloat(raw, 64)
	*f.value = parsed

	if err != nil {
		return numError(err)
	}

	return f.check(parsed)
}

func (f Float64Value) String() string {
	return strconv.FormatFloat(*ptr.OrZero(f.value), 'g', -1, 64)
}
