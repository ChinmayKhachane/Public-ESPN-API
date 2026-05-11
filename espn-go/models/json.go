package models

import (
	"encoding/json"
	"fmt"
)

// JSON is the native ESPN object shape used by raw domain clients.
type JSON map[string]any

// ESPNError is returned for non-2xx ESPN responses.
type ESPNError struct {
	StatusCode int
	URL        string
	Body       string
}

func (e *ESPNError) Error() string {
	if e.StatusCode == 0 {
		return fmt.Sprintf("espn request failed: %s", e.URL)
	}
	return fmt.Sprintf("espn request failed: status=%d url=%s body=%s", e.StatusCode, e.URL, e.Body)
}

// StringValue converts common JSON scalar values into stable path/query strings.
func StringValue(v any) string {
	switch typed := v.(type) {
	case string:
		return typed
	case json.Number:
		return typed.String()
	case float64:
		return fmt.Sprintf("%.0f", typed)
	case int:
		return fmt.Sprintf("%d", typed)
	case nil:
		return ""
	default:
		return fmt.Sprintf("%v", typed)
	}
}

// AsMap converts a decoded JSON value to JSON when possible.
func AsMap(v any) JSON {
	if m, ok := v.(map[string]any); ok {
		return JSON(m)
	}
	if m, ok := v.(JSON); ok {
		return m
	}
	return JSON{}
}

// AsSlice converts a decoded JSON value to []any when possible.
func AsSlice(v any) []any {
	if s, ok := v.([]any); ok {
		return s
	}
	return nil
}

// Nested reads a nested object path from a JSON object.
func Nested(data JSON, keys ...string) any {
	var current any = data
	for _, key := range keys {
		m := AsMap(current)
		if len(m) == 0 {
			return nil
		}
		current = m[key]
	}
	return current
}

// FirstNonEmpty returns the first value whose string representation is not empty.
func FirstNonEmpty(values ...any) any {
	for _, value := range values {
		if StringValue(value) != "" {
			return value
		}
	}
	return nil
}
