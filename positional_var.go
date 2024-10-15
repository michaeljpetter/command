package command

import (
	"github.com/michaeljpetter/command/internal"
	"github.com/michaeljpetter/command/value"
	"time"
)

// PositionalIntVar defines a positional int parameter with the given name, default value, usage, and checks.
// The pointer p defines the location to receive the parsed value.
//
// If value is nil, the parameter will have no default and will be treated as required.
func (c *Command) PositionalIntVar(p *int, name string, value *int, usage string, checks ...value.CheckFunc[int]) {
	c.PositionalVar(internal.NewIntValue(value, p, checks...), name, usage)
}

// PositionalInt defines a positional int parameter with the given name, default value, usage, and checks.
// The returned pointer receives the parsed value.
//
// If value is nil, the parameter will have no default and will be treated as required.
func (c *Command) PositionalInt(name string, value *int, usage string, checks ...value.CheckFunc[int]) *int {
	p := new(int)
	c.PositionalIntVar(p, name, value, usage, checks...)
	return p
}

// PositionalInt64Var defines a positional int64 parameter with the given name, default value, usage, and checks.
// The pointer p defines the location to receive the parsed value.
//
// If value is nil, the parameter will have no default and will be treated as required.
func (c *Command) PositionalInt64Var(p *int64, name string, value *int64, usage string, checks ...value.CheckFunc[int64]) {
	c.PositionalVar(internal.NewInt64Value(value, p, checks...), name, usage)
}

// PositionalInt64 defines a positional int64 parameter with the given name, default value, usage, and checks.
// The returned pointer receives the parsed value.
//
// If value is nil, the parameter will have no default and will be treated as required.
func (c *Command) PositionalInt64(name string, value *int64, usage string, checks ...value.CheckFunc[int64]) *int64 {
	p := new(int64)
	c.PositionalInt64Var(p, name, value, usage, checks...)
	return p
}

// PositionalUintVar defines a positional uint parameter with the given name, default value, usage, and checks.
// The pointer p defines the location to receive the parsed value.
//
// If value is nil, the parameter will have no default and will be treated as required.
func (c *Command) PositionalUintVar(p *uint, name string, value *uint, usage string, checks ...value.CheckFunc[uint]) {
	c.PositionalVar(internal.NewUintValue(value, p, checks...), name, usage)
}

// PositionalUint defines a positional uint parameter with the given name, default value, usage, and checks.
// The returned pointer receives the parsed value.
//
// If value is nil, the parameter will have no default and will be treated as required.
func (c *Command) PositionalUint(name string, value *uint, usage string, checks ...value.CheckFunc[uint]) *uint {
	p := new(uint)
	c.PositionalUintVar(p, name, value, usage, checks...)
	return p
}

// PositionalUint64Var defines a positional uint64 parameter with the given name, default value, usage, and checks.
// The pointer p defines the location to receive the parsed value.
//
// If value is nil, the parameter will have no default and will be treated as required.
func (c *Command) PositionalUint64Var(p *uint64, name string, value *uint64, usage string, checks ...value.CheckFunc[uint64]) {
	c.PositionalVar(internal.NewUint64Value(value, p, checks...), name, usage)
}

// PositionalUint64 defines a positional uint64 parameter with the given name, default value, usage, and checks.
// The returned pointer receives the parsed value.
//
// If value is nil, the parameter will have no default and will be treated as required.
func (c *Command) PositionalUint64(name string, value *uint64, usage string, checks ...value.CheckFunc[uint64]) *uint64 {
	p := new(uint64)
	c.PositionalUint64Var(p, name, value, usage, checks...)
	return p
}

// PositionalFloat64Var defines a positional float64 parameter with the given name, default value, usage, and checks.
// The pointer p defines the location to receive the parsed value.
//
// If value is nil, the parameter will have no default and will be treated as required.
func (c *Command) PositionalFloat64Var(p *float64, name string, value *float64, usage string, checks ...value.CheckFunc[float64]) {
	c.PositionalVar(internal.NewFloat64Value(value, p, checks...), name, usage)
}

// PositionalFloat64 defines a positional float64 parameter with the given name, default value, usage, and checks.
// The returned pointer receives the parsed value.
//
// If value is nil, the parameter will have no default and will be treated as required.
func (c *Command) PositionalFloat64(name string, value *float64, usage string, checks ...value.CheckFunc[float64]) *float64 {
	p := new(float64)
	c.PositionalFloat64Var(p, name, value, usage, checks...)
	return p
}

// PositionalStringVar defines a positional string parameter with the given name, default value, usage, and checks.
// The pointer p defines the location to receive the parsed value.
//
// If value is nil, the parameter will have no default and will be treated as required.
func (c *Command) PositionalStringVar(p *string, name string, value *string, usage string, checks ...value.CheckFunc[string]) {
	c.PositionalVar(internal.NewStringValue(value, p, checks...), name, usage)
}

// PositionalString defines a positional string parameter with the given name, default value, usage, and checks.
// The returned pointer receives the parsed value.
//
// If value is nil, the parameter will have no default and will be treated as required.
func (c *Command) PositionalString(name string, value *string, usage string, checks ...value.CheckFunc[string]) *string {
	p := new(string)
	c.PositionalStringVar(p, name, value, usage, checks...)
	return p
}

// PositionalDurationVar defines a positional [time.Duration] parameter with the given name, default value, usage, and checks.
// The pointer p defines the location to receive the parsed value.
//
// If value is nil, the parameter will have no default and will be treated as required.
func (c *Command) PositionalDurationVar(p *time.Duration, name string, value *time.Duration, usage string, checks ...value.CheckFunc[time.Duration]) {
	c.PositionalVar(internal.NewDurationValue(value, p, checks...), name, usage)
}

// PositionalDuration defines a positional [time.Duration] parameter with the given name, default value, usage, and checks.
// The returned pointer receives the parsed value.
//
// If value is nil, the parameter will have no default and will be treated as required.
func (c *Command) PositionalDuration(name string, value *time.Duration, usage string, checks ...value.CheckFunc[time.Duration]) *time.Duration {
	p := new(time.Duration)
	c.PositionalDurationVar(p, name, value, usage, checks...)
	return p
}
