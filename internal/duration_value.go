package internal

import (
	"github.com/michaeljpetter/command/value"
	"github.com/michaeljpetter/ptr"
	"time"
)

type DurationValue struct{ Value[time.Duration] }

func NewDurationValue(defValue *time.Duration, value *time.Duration, checks ...value.CheckFunc[time.Duration]) DurationValue {
	return DurationValue{newValue(defValue, value, checks)}
}

func (d DurationValue) Set(raw string) error {
	parsed, err := time.ParseDuration(raw)
	*d.value = parsed

	if err != nil {
		return errParse
	}

	return d.check(parsed)
}

func (d DurationValue) String() string {
	return ptr.OrZero(d.value).String()
}
