package main

import (
	"fmt"
	"testing"
)

func Test_reverseStringArray(t *testing.T) {
	input := "today is the first day of the rest of my life"
	expected := "life my of rest the of day first the is today"
	result := reverseStringArray(input)

	if result != expected {
		t.Errorf("result: %s", result)
	}
	fmt.Println(expected)
	fmt.Println(input)
}

func Test_emptySentence(t *testing.T) {
	input, expected := "", ""
	result := reverseStringArray(input)

	if result != expected {
		t.Errorf("result: %s", result)
	}
}

func Benchmark_reverseStringArray(b *testing.B) {
	b.ReportAllocs()
	input := "today is the first day of the rest of my life"
	expected := "life my of rest the of day first the is today"
	result := reverseStringArray(input)

	if result != expected {
		b.Error("")
	}
}
