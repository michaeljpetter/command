package flag

import (
	"github.com/michaeljpetter/command/internal"
	"github.com/michaeljpetter/command/value"
	"time"
)

// IntVar behaves as [flag.FlagSet.IntVar],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func (f *FlagSet) IntVar(p *int, name string, value int, usage string, checks ...value.CheckFunc[int]) {
	f.Var(internal.NewIntValue(&value, p, checks...), name, usage)
}

// Int behaves as [flag.FlagSet.Int],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func (f *FlagSet) Int(name string, value int, usage string, checks ...value.CheckFunc[int]) *int {
	p := new(int)
	f.IntVar(p, name, value, usage, checks...)
	return p
}

// Int64Var behaves as [flag.FlagSet.Int64Var],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string, checks ...value.CheckFunc[int64]) {
	f.Var(internal.NewInt64Value(&value, p, checks...), name, usage)
}

// Int64 behaves as [flag.FlagSet.Int64],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func (f *FlagSet) Int64(name string, value int64, usage string, checks ...value.CheckFunc[int64]) *int64 {
	p := new(int64)
	f.Int64Var(p, name, value, usage, checks...)
	return p
}

// UintVar behaves as [flag.FlagSet.UintVar],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string, checks ...value.CheckFunc[uint]) {
	f.Var(internal.NewUintValue(&value, p, checks...), name, usage)
}

// Uint behaves as [flag.FlagSet.Uint],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func (f *FlagSet) Uint(name string, value uint, usage string, checks ...value.CheckFunc[uint]) *uint {
	p := new(uint)
	f.UintVar(p, name, value, usage, checks...)
	return p
}

// Uint64Var behaves as [flag.FlagSet.Uint64Var],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string, checks ...value.CheckFunc[uint64]) {
	f.Var(internal.NewUint64Value(&value, p, checks...), name, usage)
}

// Uint64 behaves as [flag.FlagSet.Uint64],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func (f *FlagSet) Uint64(name string, value uint64, usage string, checks ...value.CheckFunc[uint64]) *uint64 {
	p := new(uint64)
	f.Uint64Var(p, name, value, usage, checks...)
	return p
}

// Float64Var behaves as [flag.FlagSet.Float64Var],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func (f *FlagSet) Float64Var(p *float64, name string, value float64, usage string, checks ...value.CheckFunc[float64]) {
	f.Var(internal.NewFloat64Value(&value, p, checks...), name, usage)
}

// Float64 behaves as [flag.FlagSet.Float64],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func (f *FlagSet) Float64(name string, value float64, usage string, checks ...value.CheckFunc[float64]) *float64 {
	p := new(float64)
	f.Float64Var(p, name, value, usage, checks...)
	return p
}

// StringVar behaves as [flag.FlagSet.StringVar],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func (f *FlagSet) StringVar(p *string, name string, value string, usage string, checks ...value.CheckFunc[string]) {
	f.Var(internal.NewStringValue(&value, p, checks...), name, usage)
}

// String behaves as [flag.FlagSet.String],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func (f *FlagSet) String(name string, value string, usage string, checks ...value.CheckFunc[string]) *string {
	p := new(string)
	f.StringVar(p, name, value, usage, checks...)
	return p
}

// DurationVar behaves as [flag.FlagSet.DurationVar],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func (f *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string, checks ...value.CheckFunc[time.Duration]) {
	f.Var(internal.NewDurationValue(&value, p, checks...), name, usage)
}

// Duration behaves as [flag.FlagSet.Duration],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func (f *FlagSet) Duration(name string, value time.Duration, usage string, checks ...value.CheckFunc[time.Duration]) *time.Duration {
	p := new(time.Duration)
	f.DurationVar(p, name, value, usage, checks...)
	return p
}
