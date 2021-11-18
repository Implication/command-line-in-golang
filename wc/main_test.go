package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	bytes := bytes.NewBufferString("word1 word2 word3 word4\n")

	expectation := 4
	result := count(bytes, false, false)
	if expectation != result {
		t.Errorf("Expected %d, got %d instead. \n", expectation, result)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\nline2\nline word13")

	expectation := 3
	result := count(b, true, false)

	if result != expectation {
		t.Errorf("Expected %d, got %d instead.\n", expectation, result)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3")

	expectation := 17
	result := count(b, true, true)

	if result != expectation {
		t.Errorf("Expected %d, got %d instead.\n", expectation, result)
	}
}
