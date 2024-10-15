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
