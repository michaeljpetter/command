// Package flag defines value types used by the command package.
package value

// CheckFunc defines a function that checks a value and
// returns an error when the value fails the check.
type CheckFunc[T any] func(T) error
