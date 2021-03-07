package array

import "testing"

func TestIndexOf(t *testing.T) {
	tests := []struct {
		element string
		data    []string
		want    int
	}{
		{"a", []string{}, -1},
		{"a", []string{"a", "b", "c"}, 0},
		{"b", []string{"a", "b", "c"}, 1},
		{"c", []string{"a", "b", "c"}, 2},
		{"d", []string{"a", "b", "c"}, -1},
	}

	for _, tt := range tests {
		got := IndexOf(tt.element, tt.data)
		if got != tt.want {
			t.Errorf("IndexOf(%v, %v) = %v; want %v", tt.element, tt.data, got, tt.want)
		}
	}
}
