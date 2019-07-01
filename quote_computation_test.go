package main

import (
	"reflect"
	"testing"
)

func TestComputeQuote(t *testing.T) {
	in := TranslationQuoteRequest{
		Source:       "hello world world world",
		PricePerWord: 0.54,
	}
	expected := TranslationQuote{
		TotalPrice: in.PricePerWord * 2,
		Words: map[string]int{
			"hello": 1,
			"world": 3,
		},
	}
	actual := ComputeQuote(in)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %#v, actual: %#v", expected, actual)
	}
}
