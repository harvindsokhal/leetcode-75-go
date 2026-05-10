package solution

import "testing"

func TestMergeAlternately(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  string
	}{
		{
			name:  "same length",
			word1: "abc",
			word2: "pqr",
			want:  "apbqcr",
		},
		{
			name:  "word1 longer",
			word1: "abcd",
			word2: "pq",
			want:  "apbqcd",
		},
		{
			name:  "word2 longer",
			word1: "ab",
			word2: "pqrs",
			want:  "apbqrs",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeAlternately(tt.word1, tt.word2)

			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}
