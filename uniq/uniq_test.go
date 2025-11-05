package uniq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProcessLines(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		opts     Options
		expected []string
	}{
		{
			name:     "empty input",
			input:    []string{},
			opts:     Options{},
			expected: []string{},
		},
		{
			name:     "single line",
			input:    []string{"hello"},
			opts:     Options{},
			expected: []string{"hello"},
		},
		{
			name:     "remove consecutive duplicates",
			input:    []string{"a", "a", "b", "c", "c", "c", "d"},
			opts:     Options{},
			expected: []string{"a", "b", "c", "d"},
		},
		{
			name:     "all duplicates",
			input:    []string{"x", "x", "x", "x"},
			opts:     Options{},
			expected: []string{"x"},
		},
		{
			name:     "no duplicates",
			input:    []string{"a", "b", "c", "d"},
			opts:     Options{},
			expected: []string{"a", "b", "c", "d"},
		},
		{
			name:     "with empty lines",
			input:    []string{"a", "a", "", "b", "b"},
			opts:     Options{},
			expected: []string{"a", "", "b"},
		},
		{
			name:     "count basic",
			input:    []string{"a", "a", "b", "c", "c", "c"},
			opts:     Options{Count: true},
			expected: []string{"2 a", "1 b", "3 c"},
		},
		{
			name:     "count with empty lines",
			input:    []string{"a", "a", "", "b", "b"},
			opts:     Options{Count: true},
			expected: []string{"2 a", "1 ", "2 b"},
		},
		{
			name:     "duplicate only repeated lines",
			input:    []string{"a", "a", "b", "c", "c", "c", "d"},
			opts:     Options{Duplicate: true},
			expected: []string{"a", "c"},
		},
		{
			name:     "duplicate no repeats",
			input:    []string{"a", "b", "c"},
			opts:     Options{Duplicate: true},
			expected: []string{},
		},
		{
			name:     "unique only unique lines",
			input:    []string{"a", "a", "b", "c", "c", "c", "d"},
			opts:     Options{Unique: true},
			expected: []string{"b", "d"},
		},
		{
			name:     "ignore case basic",
			input:    []string{"A", "a", "B", "b", "C"},
			opts:     Options{IgnoreRegister: true},
			expected: []string{"A", "B", "C"},
		},
		{
			name:     "fields ignore first field",
			input:    []string{"1 apple", "2 apple", "1 banana", "3 cherry"},
			opts:     Options{NumFields: 1},
			expected: []string{"1 apple", "1 banana", "3 cherry"},
		},
		{
			name:     "fields more fields than available",
			input:    []string{"short", "long text here"},
			opts:     Options{NumFields: 3},
			expected: []string{"short"},
		},
		{
			name:     "chars ignore first 2 chars",
			input:    []string{"aa_same", "bb_same", "cc_diff", "dd_same"},
			opts:     Options{NumChars: 2},
			expected: []string{"aa_same", "cc_diff", "dd_same"},
		},
		{
			name:     "chars more chars than available",
			input:    []string{"hi", "hello"},
			opts:     Options{NumChars: 5},
			expected: []string{"hi"},
		},
		{
			name:     "ignore case with count",
			input:    []string{"A", "a", "B", "b"},
			opts:     Options{IgnoreRegister: true, Count: true},
			expected: []string{"2 A", "2 B"},
		},
		{
			name:     "fields with duplicate",
			input:    []string{"1 same", "2 same", "3 different", "4 same"},
			opts:     Options{NumFields: 1, Duplicate: true},
			expected: []string{"1 same"},
		},
		{
			name:     "chars with unique",
			input:    []string{"xxA", "yyA", "xxB", "zzC"},
			opts:     Options{NumChars: 2, Unique: true},
			expected: []string{"xxB", "zzC"},
		},
		{
			name: "example from task 1",
			input: []string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
			},
			opts: Options{},
			expected: []string{
				"I love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
			},
		},
		{
			name: "example from task 2",
			input: []string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
			},
			opts: Options{Count: true},
			expected: []string{
				"3 I love music.",
				"1 ",
				"2 I love music of Kartik.",
				"1 Thanks.",
			},
		},
		{
			name: "example from task 3",
			input: []string{
				"I LOVE MUSIC.",
				"I love music.",
				"I LoVe MuSiC.",
				"",
				"I love MuSIC of Kartik.",
				"I love music of kartik.",
				"Thanks.",
			},
			opts: Options{IgnoreRegister: true},
			expected: []string{
				"I LOVE MUSIC.",
				"",
				"I love MuSIC of Kartik.",
				"Thanks.",
			},
		},
		{
			name: "example from task 4",
			input: []string{
				"We love music.",
				"I love music.",
				"They love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
			opts: Options{NumFields: 1},
			expected: []string{
				"We love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := ProcessLines(test.input, test.opts)
			require.NoError(t, err)
			require.Equal(t, test.expected, result)
		})
	}
}
