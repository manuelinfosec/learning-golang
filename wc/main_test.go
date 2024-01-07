package main

import (
	"bytes"
	"testing"
)

// TestCountWords tests the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	expected_count := 4
	result := count(b, false, false)

	if result != expected_count {
		t.Errorf("Expected %d, got %d instead.\n", expected_count, result)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1\nword2\nword3\nword4")
	expected_count := 4
	result := count(b, true, false)

	if expected_count != result {
		t.Errorf("Expected %d, got %d instead", expected_count, result)
	}

}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4 word5")
	expected_count := 29
	result := count(b, false, true)

	if expected_count != result {
		t.Errorf("Expected %d, got %d instead", expected_count, result)
	}
}
