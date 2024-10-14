package check_test

import (
	"github.com/michaeljpetter/command/check"
	"testing"
)

func TestGreaterThan(t *testing.T) {
	check := check.GreaterThan(3)

	if check(4) != nil {
		t.Error("did not pass with valid value")
	}
	if check(3) == nil {
		t.Error("did not fail with invalid value")
	}
}

func TestLessThan(t *testing.T) {
	check := check.LessThan(3)

	if check(2) != nil {
		t.Error("did not pass with valid value")
	}
	if check(3) == nil {
		t.Error("did not fail with invalid value")
	}
}

func TestAtLeast(t *testing.T) {
	check := check.AtLeast(6.)

	if check(6) != nil {
		t.Error("did not pass with valid value")
	}
	if check(5.9) == nil {
		t.Error("did not fail with invalid value")
	}
}

func TestAtMost(t *testing.T) {
	check := check.AtMost(6.)

	if check(6) != nil {
		t.Error("did not pass with valid value")
	}
	if check(6.1) == nil {
		t.Error("did not fail with invalid value")
	}
}

func TestOneOf(t *testing.T) {
	check := check.OneOf(3, 9, 11)

	if check(9) != nil {
		t.Error("did not pass with valid value")
	}
	if check(5) == nil {
		t.Error("did not fail with invalid value")
	}
}

func TestNotBlank(t *testing.T) {
	check := check.NotBlank

	if check("   ok   ") != nil {
		t.Error("did not pass with valid value")
	}
	if check("  \t   ") == nil {
		t.Error("did not fail with invalid value")
	}
}
