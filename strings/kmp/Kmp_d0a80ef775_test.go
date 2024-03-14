package kmp

import (
	"reflect"
	"testing"
)

func TestKmp_d0a80ef775(t *testing.T) {
	tests := []struct {
		word         string
		text         string
		patternTable []int
		want         []int
	}{
		{
			// Scenario 1: Normal Operation with Valid Inputs
			word:         "abc",
			text:         "abcabcabc",
			patternTable: []int{0, 0, 0},
			want:         []int{0, 3, 6},
		},
		{
			// Scenario 2: Word Length Greater than Text Length
			word:         "abcdef",
			text:         "abc",
			patternTable: []int{0, 0, 0, 0, 0, 0},
			want:         nil,
		},
		{
			// Scenario 3: Word Not Found in Text
			word:         "def",
			text:         "abcabcabc",
			patternTable: []int{0, 0, 0},
			want:         []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			if got := Kmp(tt.word, tt.text, tt.patternTable); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Kmp() = %v, want %v", got, tt.want)
			}
		})
	}
}
