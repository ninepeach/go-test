package must

import (
    "reflect"
    "testing"
)

// True asserts that the condition is true
func True(t *testing.T, condition bool) {
    if !condition {
        t.Fatal("expected condition to be true")
    }
}

// False asserts that the condition is false
func False(t *testing.T, condition bool) {
    if condition {
        t.Fatal("expected condition to be false")
    }
}

// Zero asserts that the value is zero
func Zero(t *testing.T, value interface{}) {
    if !reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface()) {
        t.Fatalf("expected zero value, got: %v", value)
    }
}

// Eq asserts equality between expected and actual
func Eq(t *testing.T, expected, actual interface{}) {
    if !reflect.DeepEqual(expected, actual) {
        t.Fatalf("expected: %v, got: %v", expected, actual)
    }
}

// Empty asserts that the container (map, slice, etc.) is empty
func Empty(t *testing.T, actual interface{}) {
    val := reflect.ValueOf(actual)
    if !val.IsValid() || (val.Kind() == reflect.Ptr && val.IsNil()) {
        return // nil or invalid values are considered empty
    }

    // Check slice, map, array, or string length
    switch val.Kind() {
    case reflect.Slice, reflect.Map, reflect.Array, reflect.String:
        if val.Len() != 0 {
            t.Fatalf("Expected empty, but got non-empty: %v", actual)
        }
    default:
        t.Fatalf("Empty check is only valid for slices, maps, arrays, and strings.")
    }
}

// NotEmpty asserts that the container is not empty
func NotEmpty(t *testing.T, container interface{}) {
    v := reflect.ValueOf(container)
    if v.Len() == 0 {
        t.Fatalf("expected not empty, but got: %v", container)
    }
}

// MapEmpty asserts that the map is empty
func MapEmpty(t *testing.T, m interface{}) {
    if reflect.ValueOf(m).Len() != 0 {
        t.Fatalf("expected map to be empty, got: %v", m)
    }
}

// MapContainsKeys asserts that the map contains the specified keys
func MapContainsKeys(t *testing.T, m interface{}, keys []string{}) {
    val := reflect.ValueOf(m)
    for _, key := range keys {
        if !val.MapIndex(reflect.ValueOf(key)).IsValid() {
            t.Fatalf("expected map to contain key: %v", key)
        }
    }
}

// MapContainsValues asserts that the map contains the specified values
func MapContainsValues(t *testing.T, m interface{}, values []interface{}) {
    found := make(map[interface{}]bool)
    val := reflect.ValueOf(m)
    for _, value := range values {
        for _, k := range val.MapKeys() {
            if reflect.DeepEqual(val.MapIndex(k).Interface(), value) {
                found[value] = true
            }
        }
        if !found[value] {
            t.Fatalf("expected map to contain value: %v", value)
        }
    }
}

// SliceEmpty asserts that the slice is empty
func SliceEmpty(t *testing.T, s interface{}) {
    if reflect.ValueOf(s).Len() != 0 {
        t.Fatalf("expected slice to be empty, got: %v", s)
    }
}

// Len asserts that the container has the specified length
func Len(t *testing.T, expectedLen int, container interface{}) {
    if reflect.ValueOf(container).Len() != expectedLen {
        t.Fatalf("expected length: %d, got: %d", expectedLen, reflect.ValueOf(container).Len())
    }
}

// SliceContainsEqual asserts that the slice contains the expected value
func SliceContainsEqual(t *testing.T, s interface{}, expected interface{}) {
    val := reflect.ValueOf(s)
    for i := 0; i < val.Len(); i++ {
        if reflect.DeepEqual(val.Index(i).Interface(), expected) {
            return
        }
    }
    t.Fatalf("expected slice to contain: %v", expected)
}
