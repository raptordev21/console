package helpers

import "testing"

func TestGenerateSpaces(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{"should return 2 spaces", "ab", "  "},
		{"should return 4 spaces", "abcd", "    "},
		{"should return 0 spaces", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := GenerateSpaces(tt.input)
			if out != tt.want {
				t.Errorf("got %s, want %s", out, tt.want)
			}
		})
	}
}
