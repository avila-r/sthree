package constraints

// Arithmetic is a type constraint for all numeric types.
type Arithmetic interface {
	// Signed integers
	int | int8 | int16 | int32 | int64 |

		// Unsigned integers
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |

		// Floating point types
		float32 | float64 |

		// Complex types
		complex64 | complex128
}
