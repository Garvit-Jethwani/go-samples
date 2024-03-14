package kmp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKmp_d0a80ef775(t *testing.T) {
	testCases := []struct {
		name         string
		word         string
		text         string
		patternTable []int
		expected     []int
	}{
		{
			name:         "TestKmpWordExists",
			word:         "abc",
			text:         "abcdefabcghi",
			patternTable: []int{-1, 0, 0},
			expected:     []int{0, 6},
		},
		{
			name:         "TestKmpWordNotExists",
			word:         "xyz",
			text:         "abcdefghijkl",
			patternTable: []int{-1, 0, 0},
			expected:     []int{},
		},
		{
			name:         "TestKmpEmptyWord",
			word:         "",
			text:         "abcdefghijkl",
			patternTable: []int{},
			expected:     nil,
		},
		{
			name:         "TestKmpWordLongerThanText",
			word:         "abcdefghijklmn",
			text:         "abcdefg",
			patternTable: []int{-1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			expected:     nil,
		},
		{
			name:         "TestKmpEmptyText",
			word:         "abc",
			text:         "",
			patternTable: []int{-1, 0, 0},
			expected:     nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			matches := Kmp(tc.word, tc.text, tc.patternTable)
			if tc.expected == nil {
				assert.Nil(t, matches, "Expected nil matches")
			} else {
				assert.Equal(t, tc.expected, matches, "Expected matches do not match")
			}
		})
	}
}
