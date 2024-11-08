package must

import (
    "testing"
    "time"
)

func TestNil(t *testing.T) {
    tc := newCase(t, `expected to be nil; is not nil`)
    t.Cleanup(tc.assert)

    Nil(tc, 42)
    Nil(tc, "hello")
    Nil(tc, time.UTC)
    Nil(tc, []string{"foo"})
    Nil(tc, map[string]int{"foo": 1})
}

func TestNotNil(t *testing.T) {
    tc := newCase(t, `expected to not be nil; is nil`)
    t.Cleanup(tc.assert)

    var s []string
    var m map[string]int

    NotNil(tc, nil)
    NotNil(tc, s)
    NotNil(tc, m)
}

func TestTrue(t *testing.T) {
    tc := newCase(t, `expected condition to be true; is false`)
    t.Cleanup(tc.assert)
    True(tc, false)
}

func TestFalse(t *testing.T) {
    tc := newCase(t, `expected condition to be false; is true`)
    t.Cleanup(tc.assert)
    False(tc, true)
}

func TestUnreachable(t *testing.T) {
    tc := newCase(t, `expected not to execute this code path`)
    t.Cleanup(tc.assert)

    Unreachable(tc)
}

func TestZero(t *testing.T) {
    tc := newCase(t, `expected value of 0`)
    t.Cleanup(tc.assert)

    Zero(tc, 1)
}

func TestNonZero(t *testing.T) {
    tc := newCase(t, `expected non-zero value`)
    t.Cleanup(tc.assert)

    NonZero(tc, 0)
}

func TestError(t *testing.T) {
    tc := newCase(t, `expected non-nil error; got nil`)
    t.Cleanup(tc.assert)

    Error(tc, nil)
}

func TestEq(t *testing.T) {
    t.Run("number", func(t *testing.T) {
        tc := newCase(t, `expected equality via cmp.Equal function`)
        t.Cleanup(tc.assert)

        Eq(tc, 42, 43)
    })

    t.Run("string", func(t *testing.T) {
        tc := newCase(t, `expected equality via cmp.Equal function`)
        t.Cleanup(tc.assert)

        Eq(tc, "foo", "bar")
    })

    t.Run("duration", func(t *testing.T) {
        tc := newCase(t, `expected equality via cmp.Equal function`)
        t.Cleanup(tc.assert)

        a := 2 * time.Second
        b := 3 * time.Minute
        Eq(tc, a, b)
    })

    t.Run("person", func(t *testing.T) {
        tc := newCase(t, `expected equality via cmp.Equal function`)
        t.Cleanup(tc.assert)

        p1 := Person{ID: 100, Name: "Alice"}
        p2 := Person{ID: 101, Name: "Bob"}
        Eq(tc, p1, p2)
    })

    t.Run("slice", func(t *testing.T) {
        tc := newCase(t, `expected equality via cmp.Equal function`)
        t.Cleanup(tc.assert)

        a := []int{1, 2, 3, 4}
        b := []int{1, 2, 9, 4}
        Eq(tc, a, b)
    })
}

func TestEqOp(t *testing.T) {
    t.Run("number", func(t *testing.T) {
        tc := newCase(t, `expected equality via ==`)
        t.Cleanup(tc.assert)
        EqOp(tc, "foo", "bar")
    })
}

func TestEqFunc(t *testing.T) {
    tc := newCase(t, `expected equality via 'eq' function`)
    t.Cleanup(tc.assert)

    a := &Person{ID: 100, Name: "Alice"}
    b := &Person{ID: 101, Name: "Bob"}

    EqFunc(tc, a, b, func(a, b *Person) bool {
        return a.ID == b.ID && a.Name == b.Name
    })
}


func TestNotEq(t *testing.T) {
    tc := newCase(t, `expected inequality via cmp.Equal function`)
    t.Cleanup(tc.assert)

    a := &Person{ID: 100, Name: "Alice"}
    b := &Person{ID: 100, Name: "Alice"}

    NotEq(tc, a, b)
}

func TestNotEqOp(t *testing.T) {
    t.Run("number", func(t *testing.T) {
        tc := newCase(t, `expected inequality via !=`)
        t.Cleanup(tc.assert)
        NotEqOp(tc, 42, 42)
    })

    t.Run("string", func(t *testing.T) {
        tc := newCase(t, `expected inequality via !=`)
        t.Cleanup(tc.assert)
        NotEqOp(tc, "foo", "foo")
    })

    t.Run("duration", func(t *testing.T) {
        tc := newCase(t, `expected inequality via !=`)
        t.Cleanup(tc.assert)
        NotEqOp(tc, 3*time.Second, 3*time.Second)
    })
}

func TestNotEqFunc(t *testing.T) {
    tc := newCase(t, `expected inequality via 'eq' function`)
    t.Cleanup(tc.assert)

    a := &Person{ID: 100, Name: "Alice"}
    b := &Person{ID: 100, Name: "Alice"}

    NotEqFunc(tc, a, b, func(a, b *Person) bool {
        return a.ID == b.ID && a.Name == b.Name
    })
}

func TestEqJSON(t *testing.T) {
    tc := newCase(t, `expected equality via JSON marshalling`)
    t.Cleanup(tc.assert)

    EqJSON(tc, `{"a":1, "b":2}`, `{"b":2, "a":9}`)
}

func TestValidJSON(t *testing.T) {
    tc := newCapture(t)
    t.Cleanup(tc.assert)

    ValidJSON(tc, `{"a":1, "b":}`)
}

func TestValidJSONBytes(t *testing.T) {
    tc := newCapture(t)
    t.Cleanup(tc.assert)

    ValidJSONBytes(tc, []byte(`{"a":1, "b":}`))
}

func TestSliceEqFunc(t *testing.T) {
    t.Run("length", func(t *testing.T) {
        tc := newCase(t, `expected slices of same length`)
        t.Cleanup(tc.assert)

        a := []int{1, 2, 3}
        b := []int{1, 2}
        SliceEqFunc(tc, a, b, func(a, b int) bool {
            return false
        })
    })

    t.Run("elements", func(t *testing.T) {
        tc := newCase(t, `expected slice equality via 'eq' function`)
        t.Cleanup(tc.assert)

        a := []*Person{
            {ID: 100, Name: "Alice"},
            {ID: 101, Name: "Bob"},
            {ID: 102, Name: "Carl"},
        }
        b := []*Person{
            {ID: 100, Name: "Alice"},
            {ID: 101, Name: "Bob"},
            {ID: 103, Name: "Dian"},
        }

        SliceEqFunc(tc, a, b, func(a, b *Person) bool {
            return a.ID == b.ID
        })
    })

    t.Run("translate", func(t *testing.T) {
        tc := newCase(t, `expected slice equality via 'eq' function`)
        t.Cleanup(tc.assert)

        values := []*Person{
            {ID: 100, Name: "Alice"},
            {ID: 101, Name: "Bob"},
        }
        exp := []string{"Alice", "Carl"}
        SliceEqFunc(tc, exp, values, (*Person).NameEquals)
    })
}

// Person implements the Equal and Less functions.
type Person struct {
    ID   int
    Name string
}

func (p *Person) Equal(o *Person) bool {
    return p.ID == o.ID
}

func (p *Person) Less(o *Person) bool {
    return p.ID < o.ID
}

func (p *Person) NameEquals(name string) bool {
    return p.Name == name
}


func TestEqual(t *testing.T) {
    tc := newCase(t, `expected equality via .Equal method`)
    t.Cleanup(tc.assert)

    a := &Person{ID: 100, Name: "Alice"}
    b := &Person{ID: 150, Name: "Alice"}

    Equal(tc, a, b)
}

func TestNotEqual(t *testing.T) {
    tc := newCase(t, `expected inequality via .Equal method`)
    t.Cleanup(tc.assert)

    a := &Person{ID: 100, Name: "Alice"}
    b := &Person{ID: 100, Name: "Alice"}

    NotEqual(tc, a, b)
}

func TestSliceEqual(t *testing.T) {
    t.Run("length", func(t *testing.T) {
        tc := newCase(t, `expected slices of same length`)
        t.Cleanup(tc.assert)

        a := []*Person{
            {ID: 100, Name: "Alice"},
            {ID: 101, Name: "Bob"},
            {ID: 102, Name: "Carl"},
        }
        b := []*Person{
            {ID: 100, Name: "Alice"},
            {ID: 101, Name: "Bob"},
        }
        SliceEqual(tc, a, b)
    })

    t.Run("elements", func(t *testing.T) {
        tc := newCase(t, `expected slice equality via .Equal method`)
        t.Cleanup(tc.assert)

        a := []*Person{
            {ID: 100, Name: "Alice"},
            {ID: 101, Name: "Bob"},
            {ID: 102, Name: "Carl"},
        }
        b := []*Person{
            {ID: 100, Name: "Alice"},
            {ID: 101, Name: "Bob"},
            {ID: 103, Name: "Dian"},
        }

        SliceEqual(tc, a, b)
    })
}

func TestSliceEqOp(t *testing.T) {
    t.Run("length", func(t *testing.T) {
        tc := newCase(t, `expected slices of same length`)
        t.Cleanup(tc.assert)

        a := []int{1, 2, 3}
        b := []int{1, 2, 3, 4}
        SliceEqOp(tc, a, b)
    })

    t.Run("elements", func(t *testing.T) {
        tc := newCase(t, `expected slice equality via ==`)
        t.Cleanup(tc.assert)

        a := []int{1, 2, 3}
        b := []int{1, 2, 4}
        SliceEqOp(tc, a, b)
    })
}

func TestSliceEmpty(t *testing.T) {
    tc := newCase(t, `expected slice to be empty`)
    t.Cleanup(tc.assert)
    SliceEmpty(tc, []int{1, 2})
}

func TestSliceNotEmpty(t *testing.T) {
    tc := newCase(t, `expected slice to not be empty`)
    t.Cleanup(tc.assert)
    SliceNotEmpty(tc, []int{})
}

func TestSliceLen(t *testing.T) {
    t.Run("strings", func(t *testing.T) {
        tc := newCase(t, `expected slice to be different length`)
        t.Cleanup(tc.assert)
        SliceLen(tc, 2, []string{"a", "b", "c"})
    })

    t.Run("numbers", func(t *testing.T) {
        tc := newCase(t, `expected slice to be different length`)
        t.Cleanup(tc.assert)
        SliceLen(tc, 3, []int{8, 9})
    })
}

func TestLen(t *testing.T) {
    t.Run("strings", func(t *testing.T) {
        tc := newCase(t, `expected slice to be different length`)
        t.Cleanup(tc.assert)
        Len(tc, 2, []string{"a", "b", "c"})
    })

    t.Run("numbers", func(t *testing.T) {
        tc := newCase(t, `expected slice to be different length`)
        t.Cleanup(tc.assert)
        Len(tc, 3, []int{8, 9})
    })
}

func TestSliceContainsOp(t *testing.T) {
    t.Run("numbers", func(t *testing.T) {
        tc := newCase(t, `expected slice to contain missing item via == operator`)
        t.Cleanup(tc.assert)
        SliceContainsOp(tc, []int{3, 4, 5}, 7)
    })

    t.Run("strings", func(t *testing.T) {
        tc := newCase(t, `expected slice to contain missing item via == operator`)
        t.Cleanup(tc.assert)
        SliceContainsOp(tc, []string{"alice", "carl"}, "bob")
    })
}

func TestSliceContainsFunc(t *testing.T) {
    tc := newCase(t, `expected slice to contain missing item via 'eq' function`)
    t.Cleanup(tc.assert)

    s := []*Person{
        {ID: 100, Name: "Alice"},
        {ID: 101, Name: "Bob"},
    }

    SliceContainsFunc(tc, s, "Carl", (*Person).NameEquals)
}

func TestSliceContainsEqual(t *testing.T) {
    tc := newCase(t, `expected slice to contain missing item via .Equal method`)
    t.Cleanup(tc.assert)

    s := []*Person{
        {ID: 100, Name: "Alice"},
        {ID: 101, Name: "Bob"},
    }

    SliceContainsEqual(tc, s, &Person{ID: 102, Name: "Carl"})
}

func TestSliceContains(t *testing.T) {
    tc := newCase(t, `expected slice to contain missing item via cmp.Equal method`)
    t.Cleanup(tc.assert)

    s := []*Person{
        {ID: 100, Name: "Alice"},
        {ID: 101, Name: "Bob"},
    }

    SliceContains(tc, s, &Person{ID: 102, Name: "Carl"})
}

func TestSliceNotContains(t *testing.T) {
    tc := newCase(t, `expected slice to not contain item but it does`)
    t.Cleanup(tc.assert)

    s := []*Person{
        {ID: 100, Name: "Alice"},
        {ID: 101, Name: "Bob"},
        {ID: 102, Name: "Carla"},
    }

    SliceNotContains(tc, s, &Person{ID: 101, Name: "Bob"})
}

func TestMapEq(t *testing.T) {
    t.Run("different length", func(t *testing.T) {
        tc := newCase(t, `expected maps of same length`)
        t.Cleanup(tc.assert)
        a := map[string]int{"a": 1}
        b := map[string]int{"a": 1, "b": 2}
        MapEq(tc, a, b)
    })

    t.Run("different keys", func(t *testing.T) {
        tc := newCase(t, `expected maps of same keys`)
        t.Cleanup(tc.assert)
        a := map[int]string{1: "a", 2: "b"}
        b := map[int]string{1: "a", 3: "c"}
        MapEq(tc, a, b)
    })

    t.Run("different values", func(t *testing.T) {
        tc := newCase(t, `expected maps of same values via cmp.Equal function`)
        t.Cleanup(tc.assert)
        a := map[string]string{"a": "amp", "b": "bar"}
        b := map[string]string{"a": "amp", "b": "foo"}
        MapEq(tc, a, b)
    })

    t.Run("custom types", func(t *testing.T) {
        tc := newCase(t, `expected maps of same values via cmp.Equal function`)
        t.Cleanup(tc.assert)

        type custom1 map[string]int
        a := custom1{"key": 1}
        type custom2 map[string]int
        b := custom2{"key": 2}
        MapEq(tc, a, b)
    })
}

func TestMapEqFunc(t *testing.T) {
    t.Run("different values", func(t *testing.T) {
        tc := newCase(t, `expected maps of same values via 'eq' function`)
        t.Cleanup(tc.assert)

        a := map[int]Person{
            0: {ID: 100, Name: "Alice"},
            1: {ID: 101, Name: "Bob"},
        }

        b := map[int]Person{
            0: {ID: 100, Name: "Alice"},
            1: {ID: 101, Name: "Bob B."},
        }

        MapEqFunc(tc, a, b, func(p1, p2 Person) bool {
            return p1.ID == p2.ID && p1.Name == p2.Name
        })
    })
}

func TestMapEqual(t *testing.T) {
    t.Run("different values", func(t *testing.T) {
        tc := newCase(t, `expected maps of same values via .Equal method`)
        t.Cleanup(tc.assert)

        a := map[int]*Person{
            0: {ID: 100, Name: "Alice"},
            1: {ID: 101, Name: "Bob"},
        }

        b := map[int]*Person{
            0: {ID: 100, Name: "Alice"},
            1: {ID: 200, Name: "Bob"},
        }

        MapEqual(tc, a, b)
    })
}

func TestMapEqOp(t *testing.T) {
    t.Run("different values", func(t *testing.T) {
        tc := newCase(t, `expected maps of same values via ==`)
        t.Cleanup(tc.assert)

        a := map[int]string{
            0: "zero",
            1: "one",
        }

        b := map[int]string{
            0: "zero",
            1: "eins",
        }

        MapEqOp(tc, a, b)
    })
    t.Run("different lengths", func(t *testing.T) {
        tc := newCase(t, `expected maps of same length`)
        t.Cleanup(tc.assert)

        a := map[int]string{
            0: "zero",
            1: "one",
        }

        b := map[int]string{
            0: "zero",
            1: "one",
            2: "two",
        }

        MapEqOp(tc, a, b)
    })
    t.Run("different keys", func(t *testing.T) {
        tc := newCase(t, `expected maps of same keys`)
        t.Cleanup(tc.assert)

        a := map[int]string{
            0: "zero",
            1: "one",
        }

        b := map[int]string{
            0: "zero",
            2: "one",
        }

        MapEqOp(tc, a, b)
    })
}

func TestMapLen(t *testing.T) {
    tc := newCase(t, `expected map to be different length`)
    t.Cleanup(tc.assert)

    m := map[string]int{"a": 1, "b": 2, "c": 3}
    MapLen(tc, 2, m)
}

func TestMapEmpty(t *testing.T) {
    tc := newCase(t, `expected map to be empty`)
    t.Cleanup(tc.assert)
    m := map[string]int{"a": 1, "b": 2}
    MapEmpty(tc, m)
}

func TestMapEmptyCustom(t *testing.T) {
    tc := newCase(t, `expected map to be empty`)
    t.Cleanup(tc.assert)
    type custom map[string]int
    m := make(custom)
    m["a"] = 1
    m["b"] = 2
    MapEmpty(tc, m)
}

func TestMapNotEmpty(t *testing.T) {
    tc := newCase(t, `expected map to not be empty`)
    t.Cleanup(tc.assert)
    m := make(map[string]string)
    MapNotEmpty(tc, m)
}

func TestMapContainsKey(t *testing.T) {
    tc := newCase(t, `expected map to contain key`)
    t.Cleanup(tc.assert)
    m := map[string]int{"a": 1, "b": 2}
    MapContainsKey(tc, m, "c")
}

func TestMapNotContainsKey(t *testing.T) {
    tc := newCase(t, `expected map to not contain key`)
    t.Cleanup(tc.assert)
    m := map[string]int{"a": 1, "b": 2}
    MapNotContainsKey(tc, m, "b")
}

func TestMapContainsKeys(t *testing.T) {
    tc := newCase(t, `expected map to contain keys`)
    t.Cleanup(tc.assert)

    m := map[string]int{"a": 1, "b": 2, "c": 3}
    MapContainsKeys(tc, m, []string{"z", "a", "b", "c", "d"})
}

func TestMapNotContainsKeys(t *testing.T) {
    tc := newCase(t, `expected map to not contain keys`)
    t.Cleanup(tc.assert)

    m := map[string]int{"a": 1, "b": 2, "c": 3}
    MapNotContainsKeys(tc, m, []string{"z", "b", "y", "c"})
}

func TestMapContainsValues(t *testing.T) {
    tc := newCase(t, `expected map to contain values`)
    t.Cleanup(tc.assert)

    m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
    MapContainsValues(tc, m, []int{9, 1, 2, 7})
}

func TestMapNotContainsValues(t *testing.T) {
    tc := newCase(t, `expected map to not contain values`)
    t.Cleanup(tc.assert)

    m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
    MapNotContainsValues(tc, m, []int{9, 8, 2, 7})
}

func TestMapContainsValuesFunc(t *testing.T) {
    tc := newCase(t, `expected map to contain values`)
    t.Cleanup(tc.assert)

    m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
    MapContainsValuesFunc(tc, m, []int{9, 1, 2, 7}, func(a, b int) bool {
        return a == b
    })
}

func TestMapNotContainsValuesFunc(t *testing.T) {
    tc := newCase(t, `expected map to not contain values`)
    t.Cleanup(tc.assert)

    m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
    MapNotContainsValuesFunc(tc, m, []int{2, 4, 6, 8}, func(a, b int) bool {
        return a == b
    })
}

func TestMapContainsValuesEqual(t *testing.T) {
    tc := newCase(t, `expected map to contain values`)
    t.Cleanup(tc.assert)

    m := map[int]*Person{
        1: {ID: 100, Name: "Alice"},
        2: {ID: 200, Name: "Bob"},
        3: {ID: 300, Name: "Carl"},
    }
    MapContainsValuesEqual(tc, m, []*Person{
        {ID: 201, Name: "Bob"},
    })
}

func TestMapNotContainsValuesEqual(t *testing.T) {
    tc := newCase(t, `expected map to not contain values`)
    t.Cleanup(tc.assert)

    m := map[int]*Person{
        1: {ID: 100, Name: "Alice"},
        2: {ID: 200, Name: "Bob"},
        3: {ID: 300, Name: "Carl"},
    }
    MapNotContainsValuesEqual(tc, m, []*Person{
        {ID: 201, Name: "Bob"}, {ID: 200, Name: "Daisy"},
    })
}

func TestMapContainsValue(t *testing.T) {
    tc := newCase(t, `expected map to contain value`)
    t.Cleanup(tc.assert)

    m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
    MapContainsValue(tc, m, 5)
}

func TestMapNotContainsValue(t *testing.T) {
    tc := newCase(t, `expected map to not contain value`)
    t.Cleanup(tc.assert)

    m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
    MapNotContainsValue(tc, m, 1)
}

func TestMapContainsValueFunc(t *testing.T) {
    tc := newCase(t, `expected map to contain value`)
    t.Cleanup(tc.assert)

    m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
    MapContainsValueFunc(tc, m, 6, func(a, b int) bool {
        return a == b
    })
}

func TestMapNotContainsValueFunc(t *testing.T) {
    tc := newCase(t, `expected map to not contain value`)
    t.Cleanup(tc.assert)

    m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
    MapNotContainsValueFunc(tc, m, 1, func(a, b int) bool {
        return a == b
    })
}

func TestMapContainsValueEqual(t *testing.T) {
    tc := newCase(t, `expected map to contain value`)
    t.Cleanup(tc.assert)

    m := map[int]*Person{
        1: {ID: 100, Name: "Alice"},
        2: {ID: 200, Name: "Bob"},
        3: {ID: 300, Name: "Carl"},
    }
    MapContainsValueEqual(tc, m, &Person{ID: 201, Name: "Bob"})
}

func TestMapNotContainsValueEqual(t *testing.T) {
    tc := newCase(t, `expected map to not contain value`)
    t.Cleanup(tc.assert)

    m := map[int]*Person{
        1: {ID: 100, Name: "Alice"},
        2: {ID: 200, Name: "Bob"},
        3: {ID: 300, Name: "Carl"},
    }
    MapNotContainsValueEqual(tc, m, &Person{ID: 200, Name: "Daisy"})
}


