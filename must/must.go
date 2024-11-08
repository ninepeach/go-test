package must

import (
	"github.com/ninepeach/go-test/assertions"
    "github.com/ninepeach/go-test/interfaces"
)

// ErrorAssertionFunc allows passing Error and NoError in table driven tests
type ErrorAssertionFunc func(t T, err error, settings ...Setting)

// Nil asserts a is nil.
func Nil(t T, a any, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.Nil(a), settings...)
}

// NotNil asserts a is not nil.
func NotNil(t T, a any, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.NotNil(a), settings...)
}

// True asserts that condition is true.
func True(t T, condition bool, settings ...Setting) {
	t.Helper()
	invoke(t, assertions.True(condition), settings...)
}

// False asserts condition is false.
func False(t T, condition bool, settings ...Setting) {
	t.Helper()
	invoke(t, assertions.False(condition), settings...)
}


// Zero asserts n == 0.
func Zero[N interfaces.Number](t T, n N, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.Zero(n), settings...)
}

// NonZero asserts n != 0.
func NonZero[N interfaces.Number](t T, n N, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.NonZero(n), settings...)
}

// Unreachable asserts a code path is not executed.
func Unreachable(t T, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.Unreachable(), settings...)
}

// Error asserts err is a non-nil error.
func Error(t T, err error, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.Error(err), settings...)
}

// Eq asserts exp and val are equal using cmp.Equal.
func Eq[A any](t T, exp, val A, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.Eq(exp, val, options(settings...)...), settings...)
}

// EqOp asserts exp == val.
func EqOp[C comparable](t T, exp, val C, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.EqOp(exp, val), settings...)
}

// EqFunc asserts exp and val are equal using eq.
func EqFunc[A any](t T, exp, val A, eq func(a, b A) bool, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.EqFunc(exp, val, eq), settings...)
}

// NotEq asserts exp and val are not equal using cmp.Equal.
func NotEq[A any](t T, exp, val A, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.NotEq(exp, val, options(settings...)...), settings...)
}

// NotEqOp asserts exp != val.
func NotEqOp[C comparable](t T, exp, val C, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.NotEqOp(exp, val), settings...)
}

// NotEqFunc asserts exp and val are not equal using eq.
func NotEqFunc[A any](t T, exp, val A, eq func(a, b A) bool, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.NotEqFunc(exp, val, eq), settings...)
}

// EqJSON asserts exp and val are equivalent JSON.
func EqJSON(t T, exp, val string, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.EqJSON(exp, val), settings...)
}

// ValidJSON asserts js is valid JSON.
func ValidJSON(t T, js string, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.ValidJSON(js), settings...)
}

// ValidJSONBytes asserts js is valid JSON.
func ValidJSONBytes(t T, js []byte, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.ValidJSONBytes(js))
}

// Equal asserts val.Equal(exp).
func Equal[E interfaces.EqualFunc[E]](t T, exp, val E, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.Equal(exp, val), settings...)
}

// NotEqual asserts !val.Equal(exp).
func NotEqual[E interfaces.EqualFunc[E]](t T, exp, val E, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.NotEqual(exp, val), settings...)
}

// MapEq asserts maps exp and val contain the same key/val pairs, using
// cmp.Equal function to compare vals.
func MapEq[M1, M2 interfaces.Map[K, V], K comparable, V any](t T, exp M1, val M2, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.MapEq(exp, val, options(settings...)), settings...)
}

// MapEqFunc asserts maps exp and val contain the same key/val pairs, using eq to
// compare vals.
func MapEqFunc[M1, M2 interfaces.Map[K, V], K comparable, V any](t T, exp M1, val M2, eq func(V, V) bool, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.MapEqFunc(exp, val, eq), settings...)
}

// MapEqual asserts maps exp and val contain the same key/val pairs, using Equal
// method to compare val
func MapEqual[M interfaces.MapEqualFunc[K, V], K comparable, V interfaces.EqualFunc[V]](t T, exp, val M, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.MapEqual(exp, val), settings...)
}

// MapEqOp asserts maps exp and val contain the same key/val pairs, using == to
// compare vals.
func MapEqOp[M interfaces.Map[K, V], K, V comparable](t T, exp M, val M, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.MapEqOp(exp, val), settings...)
}

// MapLen asserts map is of size n.
func MapLen[M ~map[K]V, K comparable, V any](t T, n int, m M, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.MapLen(n, m), settings...)
}

// MapEmpty asserts map is empty.
func MapEmpty[M ~map[K]V, K comparable, V any](t T, m M, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.MapEmpty(m), settings...)
}

// MapNotEmpty asserts map is not empty.
func MapNotEmpty[M ~map[K]V, K comparable, V any](t T, m M, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.MapNotEmpty(m), settings...)
}

// MapContainsKey asserts m contains key.
func MapContainsKey[M ~map[K]V, K comparable, V any](t T, m M, key K, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.MapContainsKey(m, key), settings...)
}

// MapNotContainsKey asserts m does not contain key.
func MapNotContainsKey[M ~map[K]V, K comparable, V any](t T, m M, key K, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.MapNotContainsKey(m, key), settings...)
}

// MapContainsKeys asserts m contains each key in keys.
func MapContainsKeys[M ~map[K]V, K comparable, V any](t T, m M, keys []K, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.MapContainsKeys(m, keys), settings...)
}

// MapNotContainsKeys asserts m does not contain any key in keys.
func MapNotContainsKeys[M ~map[K]V, K comparable, V any](t T, m M, keys []K, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.MapNotContainsKeys(m, keys), settings...)
}

// MapContainsValues asserts m contains each val in vals.
func MapContainsValues[M ~map[K]V, K comparable, V any](t T, m M, vals []V, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.MapContainsValues(m, vals, options(settings...)), settings...)
}

// MapNotContainsValues asserts m does not contain any value in vals.
func MapNotContainsValues[M ~map[K]V, K comparable, V any](t T, m M, vals []V, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.MapNotContainsValues(m, vals, options(settings...)), settings...)
}

// MapContainsValuesFunc asserts m contains each val in vals using the eq function.
func MapContainsValuesFunc[M ~map[K]V, K comparable, V any](t T, m M, vals []V, eq func(V, V) bool, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.MapContainsValuesFunc(m, vals, eq), settings...)
}

// MapNotContainsValuesFunc asserts m does not contain any value in vals using the eq function.
func MapNotContainsValuesFunc[M ~map[K]V, K comparable, V any](t T, m M, vals []V, eq func(V, V) bool, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.MapNotContainsValuesFunc(m, vals, eq), settings...)
}

// MapContainsValuesEqual asserts m contains each val in vals using the V.Equal method.
func MapContainsValuesEqual[M ~map[K]V, K comparable, V interfaces.EqualFunc[V]](t T, m M, vals []V, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.MapContainsValuesEqual(m, vals), settings...)
}

// MapNotContainsValuesEqual asserts m does not contain any value in vals using the V.Equal method.
func MapNotContainsValuesEqual[M ~map[K]V, K comparable, V interfaces.EqualFunc[V]](t T, m M, vals []V, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.MapNotContainsValuesEqual(m, vals), settings...)
}

// MapContainsValue asserts m contains val.
func MapContainsValue[M ~map[K]V, K comparable, V any](t T, m M, val V, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.MapContainsValue(m, val, options(settings...)), settings...)
}

// MapNotContainsValue asserts m does not contain val.
func MapNotContainsValue[M ~map[K]V, K comparable, V any](t T, m M, val V, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.MapNotContainsValue(m, val, options(settings...)), settings...)
}

// MapContainsValueFunc asserts m contains val using the eq function.
func MapContainsValueFunc[M ~map[K]V, K comparable, V any](t T, m M, val V, eq func(V, V) bool, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.MapContainsValueFunc(m, val, eq), settings...)
}

// MapNotContainsValueFunc asserts m does not contain val using the eq function.
func MapNotContainsValueFunc[M ~map[K]V, K comparable, V any](t T, m M, val V, eq func(V, V) bool, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.MapNotContainsValueFunc(m, val, eq), settings...)
}

// MapContainsValueEqual asserts m contains val using the V.Equal method.
func MapContainsValueEqual[M ~map[K]V, K comparable, V interfaces.EqualFunc[V]](t T, m M, val V, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.MapContainsValueEqual(m, val), settings...)
}

// MapNotContainsValueEqual asserts m does not contain val using the V.Equal method.
func MapNotContainsValueEqual[M ~map[K]V, K comparable, V interfaces.EqualFunc[V]](t T, m M, val V, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.MapNotContainsValueEqual(m, val), settings...)
}

// SliceEqFunc asserts elements of val satisfy eq for the corresponding element in exp.
func SliceEqFunc[A, B any](t T, exp []B, val []A, eq func(expectation A, value B) bool, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.EqSliceFunc(exp, val, eq), settings...)
}

// SliceEqual asserts val[n].Equal(exp[n]) for each element n.
func SliceEqual[E interfaces.EqualFunc[E]](t T, exp, val []E, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.SliceEqual(exp, val), settings...)
}

// SliceEqOp asserts exp[n] == val[n] for each element n.
func SliceEqOp[A comparable, S ~[]A](t T, exp, val S, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.SliceEqOp(exp, val), settings...)
}

// SliceEmpty asserts slice is empty.
func SliceEmpty[A any](t T, slice []A, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.SliceEmpty(slice), settings...)
}

// SliceNotEmpty asserts slice is not empty.
func SliceNotEmpty[A any](t T, slice []A, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.SliceNotEmpty(slice), settings...)
}

// SliceLen asserts slice is of length n.
func SliceLen[A any](t T, n int, slice []A, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.SliceLen(n, slice), settings...)
}

// Len asserts slice is of length n.
//
// Shorthand function for SliceLen. For checking Len() of a struct,
// use the Length() assertion.
func Len[A any](t T, n int, slice []A, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.SliceLen(n, slice), settings...)
}

// SliceContainsOp asserts item exists in slice using == operator.
func SliceContainsOp[C comparable](t T, slice []C, item C, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.SliceContainsOp(slice, item), settings...)
}

// SliceContainsFunc asserts item exists in slice, using eq to compare elements.
func SliceContainsFunc[A, B any](t T, slice []A, item B, eq func(a A, b B) bool, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.SliceContainsFunc(slice, item, eq), settings...)
}

// SliceContainsEqual asserts item exists in slice, using Equal to compare elements.
func SliceContainsEqual[E interfaces.EqualFunc[E]](t T, slice []E, item E, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.SliceContainsEqual(slice, item), settings...)
}

// SliceContains asserts item exists in slice, using cmp.Equal to compare elements.
func SliceContains[A any](t T, slice []A, item A, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.SliceContains(slice, item, options(settings...)...), settings...)
}

// SliceNotContains asserts item does not exist in slice, using cmp.Equal to
// compare elements.
func SliceNotContains[A any](t T, slice []A, item A, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.SliceNotContains(slice, item), settings...)
}

// Size asserts s.Size() is equal to exp.
func Size(t T, exp int, s interfaces.SizeFunc, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.Size(exp, s), settings...)
}

// Length asserts l.Len() is equal to exp.
func Length(t T, exp int, l interfaces.LengthFunc, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.Length(exp, l), settings...)
}

// Empty asserts e.Empty() is true.
func Empty(t T, e interfaces.EmptyFunc, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.Empty(e), settings...)
}

// NotEmpty asserts e.Empty() is false.
func NotEmpty(t T, e interfaces.EmptyFunc, settings ...Setting) {
    t.Helper()
    invoke(t, assertions.NotEmpty(e), settings...)
}

// Contains asserts container.ContainsFunc(element) is true.
func Contains[C any](t T, element C, container interfaces.ContainsFunc[C], settings ...Setting) {
    t.Helper()
    invoke(t, assertions.Contains(element, container), settings...)
}

// ContainsSubset asserts each element in elements exists in container, in no particular order.
// There may be elements in container beyond what is present in elements.
func ContainsSubset[C any](t T, elements []C, container interfaces.ContainsFunc[C], settings ...Setting) {
    t.Helper()
    invoke(t, assertions.ContainsSubset(elements, container), settings...)
}

// NotContains asserts container.ContainsFunc(element) is false.
func NotContains[C any](t T, element C, container interfaces.ContainsFunc[C], settings ...Setting) {
    t.Helper()
    invoke(t, assertions.NotContains(element, container), settings...)
}