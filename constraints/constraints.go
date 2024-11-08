package constraints

// Signed allows signed integer types.
type Signed interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned allows unsigned integer types.
type Unsigned interface {
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer allows both signed and unsigned integer types.
type Integer interface {
    Signed | Unsigned
}

// Float allows floating-point types.
type Float interface {
    ~float32 | ~float64
}

// Complex allows complex number types.
type Complex interface {
    ~complex64 | ~complex128
}

// Ordered allows types that support comparison operators.
type Ordered interface {
    Integer | Float | ~string
}