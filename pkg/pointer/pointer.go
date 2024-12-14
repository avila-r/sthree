package pointer

import (
	"time"

	"github.com/avila-r/sthree/pkg/constraints"
)

// Of creates a pointer to a given value of any type.
// It is a generic function that works with any type T.
//
// Parameters:
//   - v: The value to be converted to a pointer.
//
// Returns:
//   - A pointer to the value v.
//
// Example usage:
//
//	ptr := pointer.Of(5) // ptr is *int
func Of[T any](v T) *T {
	return &v
}

// Time returns a pointer to a time.Time value if the time is not zero.
// If the time is the zero value (i.e., time.Time{}), it returns nil.
//
// Parameters:
//   - t: A time.Time value.
//
// Returns:
//   - A pointer to t if t is not zero, otherwise nil.
//
// Example usage:
//
//	t := time.Now()
//	ptr := pointer.Time(t) // Returns a pointer to time if t is not zero
func Time(t time.Time) *time.Time {
	if t.IsZero() {
		return nil
	}

	return &t
}

// InSlice takes a variadic argument of any type T and returns a slice of pointers
// to each of the values in the input slice.
//
// Parameters:
//   - t: A variadic list of values of any type T.
//
// Returns:
//   - A slice of pointers to each of the input values.
//
// Example usage:
//
//	a := []int{1, 2, 3}
//	ptrs := pointer.InSlice(a...) // ptrs is []*int
func InSlice[T any](t ...T) []*T {
	r := []*T{}
	for _, v := range t {
		r = append(r, &v)
	}
	return r
}

// NotBlank returns a pointer to a string if the string is not empty.
// If the string is empty, it returns nil.
//
// Parameters:
//   - s: A string value.
//
// Returns:
//   - A pointer to the string s if it is not empty, otherwise nil.
//
// Example usage:
//
//	s := "Hello"
//	ptr := pointer.NotBlank(s) // Returns a pointer to the string if s is not empty
func NotBlank(s string) *string {
	if s == "" {
		return nil
	}

	return &s
}

// NotFalse returns a pointer to a boolean value if the value is true.
// If the boolean is false, it returns nil.
//
// Parameters:
//   - b: A boolean value.
//
// Returns:
//   - A pointer to the boolean b if b is true, otherwise nil.
//
// Example usage:
//
//	b := true
//	ptr := pointer.NotFalse(b) // Returns a pointer to true if b is true
func NotFalse(b bool) *bool {
	if !b {
		return nil
	}

	return &b
}

// NotZero returns a pointer to a numeric value if the value is not zero.
// It works with any type that satisfies the constraints.Arithmetic interface (e.g., int, float64).
// If the numeric value is zero, it returns nil.
//
// Parameters:
//   - n: A numeric value of type T (int, float64, etc.).
//
// Returns:
//   - A pointer to the value n if n is not zero, otherwise nil.
//
// Example usage:
//
//	x := 10
//	ptr := pointer.NotZero(x) // Returns a pointer to x if x is not zero
func NotZero[T constraints.Arithmetic](n T) *T {
	if n == 0 {
		return nil
	}

	return &n
}
