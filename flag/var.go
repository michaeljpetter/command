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

// IntVar behaves as [flag.IntVar],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func IntVar(p *int, name string, value int, usage string, checks ...value.CheckFunc[int]) {
	CommandLine.IntVar(p, name, value, usage, checks...)
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

// String behaves as [flag.FlagSet.String],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func (f *FlagSet) String(name string, value string, usage string, checks ...value.CheckFunc[string]) *string {
	p := new(string)
	f.StringVar(p, name, value, usage, checks...)
	return p
}

// StringVar behaves as [flag.StringVar],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func StringVar(p *string, name string, value string, usage string, checks ...value.CheckFunc[string]) {
	CommandLine.StringVar(p, name, value, usage, checks...)
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

// Duration behaves as [flag.FlagSet.Duration],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func (f *FlagSet) Duration(name string, value time.Duration, usage string, checks ...value.CheckFunc[time.Duration]) *time.Duration {
	p := new(time.Duration)
	f.DurationVar(p, name, value, usage, checks...)
	return p
}

// DurationVar behaves as [flag.DurationVar],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func DurationVar(p *time.Duration, name string, value time.Duration, usage string, checks ...value.CheckFunc[time.Duration]) {
	CommandLine.DurationVar(p, name, value, usage, checks...)
}

// Duration behaves as [flag.Duration],
// with an additional variadic parameter allowing checks to be applied to the parsed value.
func Duration(name string, value time.Duration, usage string, checks ...value.CheckFunc[time.Duration]) *time.Duration {
	return CommandLine.Duration(name, value, usage, checks...)
}
