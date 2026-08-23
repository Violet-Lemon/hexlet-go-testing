package textstat

import (
	"testing"
)

func TestWordCount(t *testing.T) {
	tests := []struct {
		name, in string
		want     map[string]int
	}{
		{"empty string", "", map[string]int{}},
		{"only separators", " \n\t", map[string]int{}},
		{"basic", "Green green colour COLOUR apple", map[string]int{"green": 2, "colour": 2, "apple": 1}},
		{"punct", "Green,  green colour COLOUR apple", map[string]int{"green": 2, "colour": 2, "apple": 1}},
		{"punct", "авокадо, 5  КОТ аВОКАДО apple", map[string]int{"авокадо": 2, "кот": 1, "apple": 1, "5": 1}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := WordCount(tc.in)
			if len(got) != len(tc.want) {
				t.Fatalf("len: got %d, want %d", len(got), len(tc.want))
			}
			for k, v := range tc.want {
				if got[k] != v {
					t.Fatalf("%s: got %d, want %d", k, got[k], v)
				}
			}
		})
	}
}
