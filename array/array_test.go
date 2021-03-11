package array

import (
	"testing"
)

func TestIndexOf(t *testing.T) {
	tests := map[string]struct {
		element string
		data    []string
		want    int
	}{
		"empty array":        {"a", []string{}, -1},
		"element at index 0": {"a", []string{"a", "b", "c"}, 0},
		"element at index 1": {"b", []string{"a", "b", "c"}, 1},
		"element at index 2": {"c", []string{"a", "b", "c"}, 2},
		"element not found":  {"d", []string{"a", "b", "c"}, -1},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := IndexOf(tt.element, tt.data)
			if got != tt.want {
				t.Errorf("IndexOf(%v, %v) = %v; want %v", tt.element, tt.data, got, tt.want)
			}
		})
	}
}
