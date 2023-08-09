package hw02unpackstring

import (
	"errors"
	"testing"

	test "github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "OchkiN5ada", expected: "OchkiNNNNNada"},
		{input: "di3mo4n", expected: "diiimoooon"},
		{input: "ladas2eda2n", expected: "ladassedaan"},
		{input: "belyash0", expected: "belyas"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			test.NoError(t, err)
			test.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{
		"3abc",
		"45",
		"aaa10b",
		"chukcha11",
		"belyash01",
		"---10",
		"*****!",
		"aaa1**",
	}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			test.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
