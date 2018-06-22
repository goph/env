package env

import (
	"testing"

	"strings"
)

func TestStringSliceValue_Set(t *testing.T) {
	var p []string

	val := newStringSliceValue([]string{"one", "two"}, &p)

	vals := []string{"three", "four"}

	val.Set(strings.Join(vals, ",")) // nolint: errcheck

	for i, v := range p {
		if vals[i] != v {
			t.Fatalf("expected p[%d] to be %s but got: %s", i, vals[i], v)
		}
	}
}

func TestStringSliceValue_SetEmpty(t *testing.T) {
	var p []string

	val := newStringSliceValue([]string{"one", "two"}, &p)

	val.Set("") // nolint: errcheck

	if len(p) != 0 {
		t.Error("string slice should be clear of default value upon set (even if empty)")
	}
}

func TestStringSliceValue_Comma(t *testing.T) {
	var p []string

	val := newStringSliceValue([]string{}, &p)

	expected := []string{"one,two", "three", "four,five", "six"}

	val.Set(`"one,two","three","four,five",six`) // nolint: errcheck

	if got, want := len(p), len(expected); got != want {
		t.Fatalf("expected length of string slice to be %d but got: %d", want, got)
	}

	for i, v := range p {
		if expected[i] != v {
			t.Fatalf("expected p[%d] to be %s but got: %s", i, expected[i], v)
		}
	}
}
