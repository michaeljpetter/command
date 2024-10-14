// Package flag extends the builtin flag package,
// adding support for checked values.
package flag

import (
	"flag"
	"github.com/michaeljpetter/command/internal"
	"github.com/michaeljpetter/command/value"
	"time"
)

// CommandLine wraps [flag.CommandLine] as an extended [FlagSet].
var CommandLine *FlagSet = &FlagSet{flag.CommandLine}

// Aliases for the [flag.ErrorHandling] values.
var (
	ContinueOnError = flag.ContinueOnError
	ExitOnError     = flag.ExitOnError
	PanicOnError    = flag.PanicOnError
)

// ErrorHandling aliases the [flag.ErrorHandling] type.
type ErrorHandling = flag.ErrorHandling

// Value aliases the [flag.Value] type.
type Value = flag.Value

// Flag aliases the [flag.Flag] type.
type Flag = flag.Flag

// FlagSet extends the [flag.FlagSet] type, adding support for checked flag values.
// Its flag definition methods behave in the same way as those of [flag.FlagSet],
// with the addition of a final variadic parameter that can be used to add value checks to the flag.
type FlagSet struct {
	*flag.FlagSet
}

// NewFlagSet creates a new extended [FlagSet].
func NewFlagSet(name string, errorHandling flag.ErrorHandling) *FlagSet {
	return &FlagSet{
		flag.NewFlagSet(name, errorHandling),
	}
}

// IntVar behaves as [flag.FlagSet.IntVar],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func (f *FlagSet) IntVar(p *int, name string, value int, usage string, checks ...value.CheckFunc[int]) {
	f.Var(internal.NewIntValue(&value, p, checks...), name, usage)
}

// IntVar behaves as [flag.IntVar],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func IntVar(p *int, name string, value int, usage string, checks ...value.CheckFunc[int]) {
	CommandLine.IntVar(p, name, value, usage, checks...)
}

// Int behaves as [flag.FlagSet.Int],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func (f *FlagSet) Int(name string, value int, usage string, checks ...value.CheckFunc[int]) *int {
	p := new(int)
	f.IntVar(p, name, value, usage, checks...)
	return p
}

// Int behaves as [flag.Int],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func Int(name string, value int, usage string, checks ...value.CheckFunc[int]) *int {
	return CommandLine.Int(name, value, usage, checks...)
}

// StringVar behaves as [flag.FlagSet.StringVar],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func (f *FlagSet) StringVar(p *string, name string, value string, usage string, checks ...value.CheckFunc[string]) {
	f.Var(internal.NewStringValue(&value, p, checks...), name, usage)
}

// StringVar behaves as [flag.StringVar],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func StringVar(p *string, name string, value string, usage string, checks ...value.CheckFunc[string]) {
	CommandLine.StringVar(p, name, value, usage, checks...)
}

// String behaves as [flag.FlagSet.String],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func (f *FlagSet) String(name string, value string, usage string, checks ...value.CheckFunc[string]) *string {
	p := new(string)
	f.StringVar(p, name, value, usage, checks...)
	return p
}

// String behaves as [flag.String],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func String(name string, value string, usage string, checks ...value.CheckFunc[string]) *string {
	return CommandLine.String(name, value, usage, checks...)
}

// DurationVar behaves as [flag.FlagSet.DurationVar],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func (f *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string, checks ...value.CheckFunc[time.Duration]) {
	f.Var(internal.NewDurationValue(&value, p, checks...), name, usage)
}

// DurationVar behaves as [flag.DurationVar],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func DurationVar(p *time.Duration, name string, value time.Duration, usage string, checks ...value.CheckFunc[time.Duration]) {
	CommandLine.DurationVar(p, name, value, usage, checks...)
}

// Duration behaves as [flag.FlagSet.Duration],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func (f *FlagSet) Duration(name string, value time.Duration, usage string, checks ...value.CheckFunc[time.Duration]) *time.Duration {
	p := new(time.Duration)
	f.DurationVar(p, name, value, usage, checks...)
	return p
}

// Duration behaves as [flag.Duration],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func Duration(name string, value time.Duration, usage string, checks ...value.CheckFunc[time.Duration]) *time.Duration {
	return CommandLine.Duration(name, value, usage, checks...)
}
