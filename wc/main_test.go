package main

import (
	"bytes"
	"testing"
)

// TestCountWords tests the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	expected_count := 4
	result := count(b)

	if result != expected_count {
		t.Errorf("Expected %d, got %d instead.\n", expected_count, result)
	}
}
