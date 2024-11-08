package assertions

import (
    "encoding/json"
    "errors"
    "fmt"
    "reflect"
    "runtime"
    "strings"
    "path/filepath"

    "github.com/google/go-cmp/cmp"
    "github.com/ninepeach/go-test/interfaces"
)

const depth = 4

func Caller() string {
    _, file, line, ok := runtime.Caller(depth)
    if ok {
        return fmt.Sprintf("%s:%d: ", filepath.Base(file), line)
    }
    return "[???]"
}

// Creates a diff between `a` and `b` using `cmp.Diff`. Falls back to a string comparison if needed.
func diff[A, B any](a A, b B, opts cmp.Options) (result string) {
    defer func() {
        if r := recover(); r != nil {
            result = fmt.Sprintf("↪ Assertion | comparison ↷\na: %#v\nb: %#v\n", a, b)
        }
    }()
    result = "↪ Assertion | differential ↷\n" + cmp.Diff(a, b, opts)
    return
}

// Compares `a` and `b` using `cmp.Equal`, falling back to `reflect.DeepEqual` if necessary.
func equal[A, B any](a A, b B, opts cmp.Options) (isEqual bool) {
    defer func() {
        if r := recover(); r != nil {
            isEqual = reflect.DeepEqual(a, b)
        }
    }()
    return cmp.Equal(a, b, opts)
}

// Checks if `item` exists in `slice`.
func contains[C comparable](slice []C, item C) bool {
    for _, el := range slice {
        if el == item {
            return true
        }
    }
    return false
}

// Checks if `item` exists in `slice` using a custom equality function `eq`.
func containsFunc[A, B any](slice []A, item B, eq func(A, B) bool) bool {
    for _, el := range slice {
        if eq(el, item) {
            return true
        }
    }
    return false
}

// Checks if `items` exists in `slice` using a custom equality function `eq`.
func containsSubsetFunc[A, B any](slice []A, items []B, eq func(A, B) bool) (bool, B) {
    for _, target := range items {
        found := false
        for _, el := range slice {
            if eq(el, target) {
                found = true
                break
            }
        }
        if !found {
            return false, target
        }
    }
    var zero B
    return true, zero
}

// Checks if `val` is nil, including for types with non-nil interface values.
func isNil(val any) bool {
    if val == nil {
        return true
    }
    v := reflect.ValueOf(val)
    switch v.Kind() {
    case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
        return v.IsNil()
    default:
        return false
    }
}

// Asserts `val` is nil, else returns a message.
func Nil(val any) string {
    if !isNil(val) {
        return "expected to be nil; is not nil\n"
    }
    return ""
}

func NotNil(a any) (s string) {
    if isNil(a) {
        s = "expected to not be nil; is nil\n"
    }
    return
}

// Asserts `condition` is true, else returns a message.
func True(condition bool) string {
    if !condition {
        return "expected condition to be true; is false\n"
    }
    return ""
}

func False(condition bool) (s string) {
    if condition {
        s = "expected condition to be false; is true\n"
    }
    return
}

func Zero[N interfaces.Number](value N) (s string) {
    if value != 0 {
        s = "expected value of 0\n"
        s += fmt.Sprintf("↪value: %v\n", value)
    }
    return
}

func NonZero[N interfaces.Number](value N) (s string) {
    if value == 0 {
        s = "expected non-zero value\n"
        s += fmt.Sprintf("↪value: %v\n", value)
    }
    return
}

func Unreachable() (s string) {
    s = "expected not to execute this code path\n"
    return
}

func Error(err error) (s string) {
    if err == nil {
        s = "expected non-nil error; got nil\n"
    }
    return
}

func EqError(err error, msg string) (s string) {
    if err == nil {
        s = "expected non-nil error; got nil\n"
        return
    }
    e := err.Error()
    if e != msg {
        s = "expected matching error strings\n"
        s += fmt.Sprintf("↪msg: %q\n", msg)
        s += fmt.Sprintf("↪err: %q\n", e)
    }
    return
}

func ErrorIs(err error, target error) (s string) {
    if err == nil {
        s = "expected non-nil error; got nil\n"
        return
    }
    if !errors.Is(err, target) {
        s = "expected errors.Is match\n"
        s += fmt.Sprintf("↪target: %v\n", target)
        s += fmt.Sprintf("↪   got: %v\n", err)
    }
    return
}

func ErrorAs[E error, Target *E](err error, target Target) (s string) {
    if err == nil {
        s = "expected non-nil error; got nil\n"
        return
    }
    if target == nil {
        s = "expected non-nil target; got nil\n"
        return
    }
    if !errors.As(err, target) {
        s = "expected errors.As match\n"
        s += fmt.Sprintf("↪target: %v\n", target)
        s += fmt.Sprintf("↪   got: %v\n", err)
    }
    return
}

func NoError(err error) (s string) {
    if err != nil {
        s = "expected nil error\n"
        s += fmt.Sprintf("↪error: %v\n", err)
    }
    return
}

func ErrorContains(err error, sub string) (s string) {
    if err == nil {
        s = "expected non-nil error; got nil\n"
        return
    }
    actual := err.Error()
    if !strings.Contains(actual, sub) {
        s = "expected error to contain substring\n"
        s += fmt.Sprintf("↪substring: %s\n", sub)
        s += fmt.Sprintf("↪      err: %s\n", actual)
    }
    return
}

func Eq[A any](exp, val A, opts ...cmp.Option) (s string) {
    if !equal(exp, val, opts) {
        s = "expected equality via cmp.Equal function\n"
        s += diff(exp, val, opts)
    }
    return
}

func NotEq[A any](exp, val A, opts ...cmp.Option) (s string) {
    if equal(exp, val, opts) {
        s = "expected inequality via cmp.Equal function\n"
    }
    return
}

func EqOp[C comparable](exp, val C) (s string) {
    if exp != val {
        s = "expected equality via ==\n"
        s += diff(exp, val, nil)
    }
    return
}

func EqFunc[A any](exp, val A, eq func(a, b A) bool) (s string) {
    if !eq(exp, val) {
        s = "expected equality via 'eq' function\n"
        s += diff(exp, val, nil)
    }
    return
}

func NotEqOp[C comparable](exp, val C) (s string) {
    if exp == val {
        s = "expected inequality via !=\n"
    }
    return
}

func NotEqFunc[A any](exp, val A, eq func(a, b A) bool) (s string) {
    if eq(exp, val) {
        s = "expected inequality via 'eq' function\n"
    }
    return
}

func EqJSON(exp, val string) (s string) {
    var expA, expB any

    if err := json.Unmarshal([]byte(exp), &expA); err != nil {
        s = fmt.Sprintf("failed to unmarshal first argument as JSON: %v\n", err)
        return
    }

    if err := json.Unmarshal([]byte(val), &expB); err != nil {
        s = fmt.Sprintf("failed to unmarshal second argument as JSON: %v\n", err)
        return
    }

    if !reflect.DeepEqual(expA, expB) {
        jsonA, _ := json.Marshal(expA)
        jsonB, _ := json.Marshal(expB)
        s = "expected equality via JSON marshalling\n"
        s += diff(string(jsonA), string(jsonB), nil)
        return
    }

    return
}

func ValidJSON(input string) (s string) {
    return validJSON([]byte(input))
}

func ValidJSONBytes(input []byte) (s string) {
    return validJSON(input)
}

func validJSON(input []byte) (s string) {
    if !json.Valid([]byte(input)) {
        return "expected input to be valid JSON\n"
    }
    return
}

func EqSliceFunc[A, B any](exp []B, val []A, eq func(a A, b B) bool) (s string) {
    lenA, lenB := len(exp), len(val)

    if lenA != lenB {
        s = "expected slices of same length\n"
        s += fmt.Sprintf("↪len(exp): %d\n", lenA)
        s += fmt.Sprintf("↪len(val): %d\n", lenB)
        s += diff(exp, val, nil)
        return
    }

    miss := false
    for i := 0; i < lenA; i++ {
        if !eq(val[i], exp[i]) {
            miss = true
            break
        }
    }

    if miss {
        s = "expected slice equality via 'eq' function\n"
        s += diff(exp, val, nil)
        return
    }

    return
}

func Equal[E interfaces.EqualFunc[E]](exp, val E) (s string) {
    if !val.Equal(exp) {
        s = "expected equality via .Equal method\n"
        s += diff(exp, val, nil)
    }
    return
}

func NotEqual[E interfaces.EqualFunc[E]](exp, val E) (s string) {
    if val.Equal(exp) {
        s = "expected inequality via .Equal method\n"
    }
    return
}

func SliceEqual[E interfaces.EqualFunc[E]](exp, val []E) (s string) {
    lenA, lenB := len(exp), len(val)

    if lenA != lenB {
        s = "expected slices of same length\n"
        s += fmt.Sprintf("↪len(exp): %d\n", lenA)
        s += fmt.Sprintf("↪len(val): %d\n", lenB)
        s += diff(exp, val, nil)
        return
    }

    for i := 0; i < lenA; i++ {
        if !exp[i].Equal(val[i]) {
            s += "expected slice equality via .Equal method\n"
            s += diff(exp[i], val[i], nil)
            return
        }
    }
    return
}

func SliceEqOp[A comparable, S ~[]A](exp, val S) (s string) {
    lenA, lenB := len(exp), len(val)

    if lenA != lenB {
        s = "expected slices of same length\n"
        s += fmt.Sprintf("↪len(exp): %d\n", lenA)
        s += fmt.Sprintf("↪len(val): %d\n", lenB)
        s += diff(exp, val, nil)
        return
    }

    for i := 0; i < lenA; i++ {
        if exp[i] != val[i] {
            s += "expected slice equality via ==\n"
            s += diff(exp[i], val[i], nil)
            return
        }
    }
    return
}

func Lesser[L interfaces.LessFunc[L]](exp, val L) (s string) {
    if !val.Less(exp) {
        s = "expected val to be less via .Less method\n"
        s += diff(exp, val, nil)
    }
    return
}

func SliceEmpty[A any](slice []A) (s string) {
    if len(slice) != 0 {
        s = "expected slice to be empty\n"
        s += fmt.Sprintf("↪len(slice): %d\n", len(slice))
    }
    return
}

func SliceNotEmpty[A any](slice []A) (s string) {
    if len(slice) == 0 {
        s = "expected slice to not be empty\n"
        s += fmt.Sprintf("↪len(slice): %d\n", len(slice))
    }
    return
}

func SliceLen[A any](n int, slice []A) (s string) {
    if l := len(slice); l != n {
        s = "expected slice to be different length\n"
        s += fmt.Sprintf("↪len(slice): %d, expected: %d\n", l, n)
    }
    return
}

func SliceContainsOp[C comparable](slice []C, item C) (s string) {
    if !contains(slice, item) {
        s = "expected slice to contain missing item via == operator\n"
        s += fmt.Sprintf("↪slice is missing %#v\n", item)
    }
    return
}

func SliceContainsFunc[A, B any](slice []A, item B, eq func(a A, b B) bool) (s string) {
    if !containsFunc(slice, item, eq) {
        s = "expected slice to contain missing item via 'eq' function\n"
        s += fmt.Sprintf("↪slice is missing %#v\n", item)
    }
    return
}

func SliceContainsEqual[E interfaces.EqualFunc[E]](slice []E, item E) (s string) {
    if !containsFunc(slice, item, E.Equal) {
        s = "expected slice to contain missing item via .Equal method\n"
        s += fmt.Sprintf("↪slice is missing %#v\n", item)
    }
    return
}

func SliceContains[A any](slice []A, item A, opts ...cmp.Option) (s string) {
    for _, i := range slice {
        if cmp.Equal(i, item, opts...) {
            return
        }
    }
    s = "expected slice to contain missing item via cmp.Equal method\n"
    s += fmt.Sprintf("↪slice is missing %#v\n", item)
    return
}

func SliceNotContains[A any](slice []A, item A, opts ...cmp.Option) (s string) {
    for _, i := range slice {
        if cmp.Equal(i, item, opts...) {
            s = "expected slice to not contain item but it does\n"
            s += fmt.Sprintf("↪unwanted item %#v\n", item)
            return
        }
    }
    return
}

func MapEq[M1, M2 interfaces.Map[K, V], K comparable, V any](exp M1, val M2, opts cmp.Options) (s string) {
    lenA, lenB := len(exp), len(val)

    if lenA != lenB {
        s = "expected maps of same length\n"
        s += fmt.Sprintf("↪len(exp): %d\n", lenA)
        s += fmt.Sprintf("↪len(val): %d\n", lenB)
        return
    }

    for key, valA := range exp {
        valB, exists := val[key]
        if !exists {
            s = "expected maps of same keys\n"
            s += diff(exp, val, opts)
            return
        }

        if !cmp.Equal(valA, valB, opts) {
            s = "expected maps of same values via cmp.Equal function\n"
            s += diff(exp, val, opts)
            return
        }
    }
    return
}

func MapEqFunc[M1, M2 interfaces.Map[K, V], K comparable, V any](exp M1, val M2, eq func(V, V) bool) (s string) {
    lenA, lenB := len(exp), len(val)

    if lenA != lenB {
        s = "expected maps of same length\n"
        s += fmt.Sprintf("↪len(exp): %d\n", lenA)
        s += fmt.Sprintf("↪len(val): %d\n", lenB)
        return
    }

    for key, valA := range exp {
        valB, exists := val[key]
        if !exists {
            s = "expected maps of same keys\n"
            s += diff(exp, val, nil)
            return
        }

        if !eq(valA, valB) {
            s = "expected maps of same values via 'eq' function\n"
            s += diff(exp, val, nil)
            return
        }
    }
    return
}

func MapEqual[M interfaces.MapEqualFunc[K, V], K comparable, V interfaces.EqualFunc[V]](exp, val M) (s string) {
    lenA, lenB := len(exp), len(val)

    if lenA != lenB {
        s = "expected maps of same length\n"
        s += fmt.Sprintf("↪len(exp): %d\n", lenA)
        s += fmt.Sprintf("↪len(val): %d\n", lenB)
        return
    }

    for key, valA := range exp {
        valB, exists := val[key]
        if !exists {
            s = "expected maps of same keys\n"
            s += diff(exp, val, nil)
            return
        }

        if !(valB).Equal(valA) {
            s = "expected maps of same values via .Equal method\n"
            s += diff(exp, val, nil)
            return
        }
    }

    return
}

func MapEqOp[M interfaces.Map[K, V], K, V comparable](exp, val M) (s string) {
    lenA, lenB := len(exp), len(val)

    if lenA != lenB {
        s = "expected maps of same length\n"
        s += fmt.Sprintf("↪len(exp): %d\n", lenA)
        s += fmt.Sprintf("↪len(val): %d\n", lenB)
        return
    }

    for key, valA := range exp {
        valB, exists := val[key]
        if !exists {
            s = "expected maps of same keys\n"
            s += diff(exp, val, nil)
            return
        }

        if valA != valB {
            s = "expected maps of same values via ==\n"
            s += diff(exp, val, nil)
            return
        }
    }

    return
}

func MapLen[M ~map[K]V, K comparable, V any](n int, m M) (s string) {
    if l := len(m); l != n {
        s = "expected map to be different length\n"
        s += fmt.Sprintf("↪len(map): %d, expected: %d\n", l, n)
    }
    return
}

func MapEmpty[M ~map[K]V, K comparable, V any](m M) (s string) {
    if l := len(m); l > 0 {
        s = "expected map to be empty\n"
        s += fmt.Sprintf("↪len(map): %d\n", l)
    }
    return
}

func MapNotEmpty[M ~map[K]V, K comparable, V any](m M) (s string) {
    if l := len(m); l == 0 {
        s = "expected map to not be empty\n"
        s += fmt.Sprintf("↪len(map): %d\n", l)
    }
    return
}

func MapContainsKey[M ~map[K]V, K comparable, V any](m M, key K) (s string) {
    if _, exists := m[key]; !exists {
        s = "expected map to contain key\n"
        s += fmt.Sprintf("↪key: %v\n", key)
    }
    return
}

func MapNotContainsKey[M ~map[K]V, K comparable, V any](m M, key K) (s string) {
    if _, exists := m[key]; exists {
        s = "expected map to not contain key\n"
        s += fmt.Sprintf("↪key: %v\n", key)
    }
    return
}

func MapContainsKeys[M ~map[K]V, K comparable, V any](m M, keys []K) (s string) {
    var missing []K
    for _, key := range keys {
        if _, exists := m[key]; !exists {
            missing = append(missing, key)
        }
    }
    if len(missing) > 0 {
        s = "expected map to contain keys\n"
        for _, key := range missing {
            s += fmt.Sprintf("↪key: %v\n", key)
        }
    }
    return
}

func MapNotContainsKeys[M ~map[K]V, K comparable, V any](m M, keys []K) (s string) {
    var unwanted []K
    for _, key := range keys {
        if _, exists := m[key]; exists {
            unwanted = append(unwanted, key)
        }
    }
    if len(unwanted) > 0 {
        s = "expected map to not contain keys\n"
        for _, key := range unwanted {
            s += fmt.Sprintf("↪key: %v\n", key)
        }
    }
    return
}

func mapContains[M ~map[K]V, K comparable, V any](m M, values []V, eq func(V, V) bool) (s string) {
    var missing []V
    for _, wanted := range values {
        found := false
        for _, v := range m {
            if eq(wanted, v) {
                found = true
                break
            }
        }
        if !found {
            missing = append(missing, wanted)
        }
    }

    if len(missing) > 0 {
        s = "expected map to contain values\n"
        for _, val := range missing {
            s += fmt.Sprintf("↪val: %v\n", val)
        }
    }
    return
}

func mapNotContains[M ~map[K]V, K comparable, V any](m M, values []V, eq func(V, V) bool) (s string) {
    var unexpected []V
    for _, target := range values {
        found := false
        for _, v := range m {
            if eq(target, v) {
                found = true
                break
            }
        }
        if found {
            unexpected = append(unexpected, target)
        }
    }
    if len(unexpected) > 0 {
        s = "expected map to not contain values\n"
        for _, val := range unexpected {
            s += fmt.Sprintf("↪val: %v\n", val)
        }
    }
    return
}

func MapContainsValues[M ~map[K]V, K comparable, V any](m M, vals []V, opts cmp.Options) (s string) {
    return mapContains(m, vals, func(a, b V) bool {
        return equal(a, b, opts)
    })
}

func MapNotContainsValues[M ~map[K]V, K comparable, V any](m M, vals []V, opts cmp.Options) (s string) {
    return mapNotContains(m, vals, func(a, b V) bool {
        return equal(a, b, opts)
    })
}

func MapContainsValuesFunc[M ~map[K]V, K comparable, V any](m M, vals []V, eq func(V, V) bool) (s string) {
    return mapContains(m, vals, eq)
}

func MapNotContainsValuesFunc[M ~map[K]V, K comparable, V any](m M, vals []V, eq func(V, V) bool) (s string) {
    return mapNotContains(m, vals, eq)
}

func MapContainsValuesEqual[M ~map[K]V, K comparable, V interfaces.EqualFunc[V]](m M, vals []V) (s string) {
    return mapContains(m, vals, func(a, b V) bool {
        return a.Equal(b)
    })
}

func MapNotContainsValuesEqual[M ~map[K]V, K comparable, V interfaces.EqualFunc[V]](m M, vals []V) (s string) {
    return mapNotContains(m, vals, func(a, b V) bool {
        return a.Equal(b)
    })
}

func MapContainsValue[M ~map[K]V, K comparable, V any](m M, val V, opts cmp.Options) (s string) {
    return mapContains(m, []V{val}, func(a, b V) bool {
        return equal(a, b, opts)
    })
}

func MapNotContainsValue[M ~map[K]V, K comparable, V any](m M, val V, opts cmp.Options) (s string) {
    return mapNotContains(m, []V{val}, func(a, b V) bool {
        return equal(a, b, opts)
    })
}

func MapContainsValueFunc[M ~map[K]V, K comparable, V any](m M, val V, eq func(V, V) bool) (s string) {
    return mapContains(m, []V{val}, eq)
}

func MapNotContainsValueFunc[M ~map[K]V, K comparable, V any](m M, val V, eq func(V, V) bool) (s string) {
    return mapNotContains(m, []V{val}, eq)
}

func MapContainsValueEqual[M ~map[K]V, K comparable, V interfaces.EqualFunc[V]](m M, val V) (s string) {
    return mapContains(m, []V{val}, func(a, b V) bool {
        return a.Equal(b)
    })
}

func MapNotContainsValueEqual[M ~map[K]V, K comparable, V interfaces.EqualFunc[V]](m M, val V) (s string) {
    return mapNotContains(m, []V{val}, func(a, b V) bool {
        return a.Equal(b)
    })
}

func Length(n int, length interfaces.LengthFunc) (s string) {
    if l := length.Len(); l != n {
        s = "expected different length\n"
        s += fmt.Sprintf("↪   length: %d\n", l)
        s += fmt.Sprintf("↪ expected: %d\n", n)
    }
    return
}

func Size(n int, size interfaces.SizeFunc) (s string) {
    if l := size.Size(); l != n {
        s = "expected different size\n"
        s += fmt.Sprintf("↪     size: %d\n", l)
        s += fmt.Sprintf("↪ expected: %d\n", n)
    }
    return
}

func Empty(e interfaces.EmptyFunc) (s string) {
    if !e.Empty() {
        s = "expected to be empty, but was not\n"
    }
    return
}

func NotEmpty(e interfaces.EmptyFunc) (s string) {
    if e.Empty() {
        s = "expected to not be empty, but is\n"
    }
    return
}

func Contains[C any](i C, c interfaces.ContainsFunc[C]) (s string) {
    if !c.Contains(i) {
        s = "expected to contain element, but does not\n"
    }
    return
}

func ContainsSubset[C any](elements []C, container interfaces.ContainsFunc[C]) (s string) {
    for i := 0; i < len(elements); i++ {
        element := elements[i]
        if !container.Contains(element) {
            s = "expected to contain element, but does not\n"
            s += fmt.Sprintf("↪ element: %v\n", element)
            return
        }
    }
    return
}

func NotContains[C any](i C, c interfaces.ContainsFunc[C]) (s string) {
    if c.Contains(i) {
        s = "expected not to contain element, but it does\n"
    }
    return
}