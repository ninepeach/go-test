package must

import "testing"

func TestTrue(t *testing.T) {
    True(t, true)    // Should pass
    True(t, 1 == 1)  // Should pass
}

func TestFalse(t *testing.T) {
    False(t, false)    // Should pass
    False(t, 1 == 2)   // Should pass
}

func TestEq(t *testing.T) {
    Eq(t, 5, 5)       // Should pass
    Eq(t, "test", "test") // Should pass
}

func TestZero(t *testing.T) {
    Zero(t, 0)        // Should pass
    Zero(t, "")       // Should pass
}

func TestNotEmpty(t *testing.T) {
    NotEmpty(t, []int{1, 2, 3}) // Should pass
    NotEmpty(t, "hello")        // Should pass
}

func TestEmpty(t *testing.T) {
    Empty(t, "")      // Should pass
    Empty(t, nil)     // Should pass
}

func TestMapContainsKeys(t *testing.T) {
    data := map[string]int{"one": 1, "two": 2}
    MapContainsKeys(t, data, []interface{}{"one", "two"}) // Corrected
}

func TestSliceContainsEqual(t *testing.T) {
    slice := []int{1, 2, 3}
    SliceContainsEqual(t, slice, 2) // Should pass
}

func TestLen(t *testing.T) {
    slice := []int{1, 2, 3}
    Len(t, 3, slice) // Should pass
}
