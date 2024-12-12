package pointer

import "time"

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
