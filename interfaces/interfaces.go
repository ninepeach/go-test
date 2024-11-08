package interfaces

import (
    "math"
    "github.com/ninepeach/go-test/constraints"
)

// MinFunc represents a type with a Min() method.
type MinFunc[T any] interface {
    Min() T
}

// MaxFunc represents a type with a Max() method.
type MaxFunc[T any] interface {
    Max() T
}

// EqualFunc represents a type with an Equal() method.
type EqualFunc[A any] interface {
    Equal(A) bool
}

// CopyFunc represents a type with a Copy() method.
type CopyFunc[A any] interface {
    Copy() A
}

// CopyEqual represents a type satisfying both EqualFunc and CopyFunc.
type CopyEqual[T any] interface {
    EqualFunc[T]
    CopyFunc[T]
}

// TweakFunc modifies values in tests.
type TweakFunc[E CopyEqual[E]] func(E)

// LessFunc represents a type with a Less() method.
type LessFunc[A any] interface {
    Less(A) bool
}

// Map represents a map with comparable keys.
type Map[K comparable, V any] interface {
    ~map[K]V
}

// MapEqualFunc represents a map where values implement Equal().
type MapEqualFunc[K comparable, V EqualFunc[V]] interface {
    ~map[K]V
}

// Number represents numeric types (integers, floats, or complex).
type Number interface {
    constraints.Ordered
    constraints.Float | constraints.Integer | constraints.Complex
}

// Numeric checks if n is a valid number (not Inf or NaN).
func Numeric[N Number](n N) bool {
    return !math.IsNaN(float64(n)) && !math.IsInf(float64(n), 0)
}

// LengthFunc satisfies types with Len() method.
type LengthFunc interface {
    Len() int
}

// SizeFunc satisfies types with Size() method.
type SizeFunc interface {
    Size() int
}

// EmptyFunc satisfies types with Empty() method.
type EmptyFunc interface {
    Empty() bool
}

// ContainsFunc satisfies types with Contains() method.
type ContainsFunc[T any] interface {
    Contains(T) bool
}