// Package flag extends the stdlib flag package,
// adding support for checked values.
package flag

import (
	"flag"
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
