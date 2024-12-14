package pointer

import (
	"time"

	"github.com/avila-r/sthree/pkg/constraints"
)

func Of[T any](v T) *T {
	return &v
}

func Time(t time.Time) *time.Time {
	if t.IsZero() {
		return nil
	}

	return &t
}

func NotBlank(s string) *string {
	if s == "" {
		return nil
	}

	return &s
}

func NotFalse(b bool) *bool {
	if !b {
		return nil
	}

	return &b
}

func NotZero[T constraints.Arithmetic](n T) *T {
	if n == 0 {
		return nil
	}

	return &n
}
