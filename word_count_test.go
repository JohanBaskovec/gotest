package main

import (
	"reflect"
	"testing"
)

type InAndOut struct {
	in  string
	out map[string]int
}

func TestReplaceSpecialCharacters(t *testing.T) {
	in := "   ! ⻔ hello ! world )=+${'&*[⻑ Ỡ⁇≥Ⱘ hello〛"
	actual := ReplaceSpecialCharacters(in, " ")
	expected := "     ⻔ hello   world          ⻑ Ỡ  Ⱘ hello "
	if actual != expected {
		t.Errorf(`Expected: "%v", actual: "%v"`, expected, actual)
	}
}

func TestSplitByWhiteSpace(t *testing.T) {
	in := `

⻔ hello   world          ⻑ Ỡ  Ⱘ hello 

			hey
`
	expected := []string{
		"⻔",
		"hello",
		"world",
		"⻑",
		"Ỡ",
		"Ⱘ",
		"hello",
		"hey",
	}
	actual := SplitByWhiteSpace(in)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf(`Expected: "%#v", actual: "%#v"`, expected, actual)
	}
}

func TestCountWords(t *testing.T) {
	expectedInOut := []InAndOut{
		{
			in: "hello world",
			out: map[string]int{
				"hello": 1,
				"world": 1,
			},
		},
		{
			in: "   ! ⻔ hello ! world )=+${'&*[⻑ Ỡ⁇≥Ⱘ hello〛",
			out: map[string]int{
				"⻔":     1,
				"hello": 2,
				"world": 1,
				"⻑":     1,
				"Ỡ":     1,
				"Ⱘ":     1,
			},
		},
	}
	for _, expected := range expectedInOut {
		wordCount := CountWords(expected.in)
		if !reflect.DeepEqual(wordCount, expected.out) {
			t.Errorf("Maps aren't equal! Expected: %v, actual: %v", expected.out, wordCount)
		}
	}

}
