package env

import "testing"

func TestQueryStringValue_Set(t *testing.T) {
	p := new(map[string]string)

	val := newQueryStringValue(map[string]string{"key": "value"}, p)

	val.Set("key2=value2") // nolint: errcheck

	if _, exists := (*val)["key"]; exists {
		t.Error("query string map should be clear of default value upon set")
	}
}

func TestQueryStringValue_SetEmpty(t *testing.T) {
	p := new(map[string]string)

	val := newQueryStringValue(map[string]string{"key": "value"}, p)

	val.Set("") // nolint: errcheck

	if _, exists := (*val)["key"]; exists {
		t.Error("query string map should be clear of default value upon set (even if empty)")
	}
}
