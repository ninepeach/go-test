package must

import (
	"testing"
)

// True asserts that a condition is true; if not, it fails the test.
func True(t *testing.T, condition bool, msg string) {
	if !condition {
		t.Errorf("Assertion failed: %s", msg)
	}
}

// False asserts that a condition is false; if not, it fails the test.
func False(t *testing.T, condition bool, msg string) {
	if condition {
		t.Errorf("Assertion failed: %s", msg)
	}
}

// Equal asserts that two values are equal; if not, it fails the test.
func Equal[T comparable](t *testing.T, expected, actual T, msg string) {
	if expected != actual {
		t.Errorf("Assertion failed: %s - expected %v, got %v", msg, expected, actual)
	}
}

// NotNil asserts that a value is not nil; if it is nil, it fails the test.
func NotNil(t *testing.T, value any, msg string) {
	if value == nil {
		t.Errorf("Assertion failed: %s - value is nil", msg)
	}
}

// Nil asserts that a value is nil; if it's not nil, it fails the test.
func Nil(t *testing.T, value any, msg string) {
	if value != nil {
		t.Errorf("Assertion failed: %s - expected nil, got %v", msg, value)
	}
}
