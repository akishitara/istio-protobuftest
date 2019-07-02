package main

import (
	"testing"
)

func TestAaaBbb(t *testing.T) {
	a := "aaa"
	b := "aaa"
	if a != b {
    t.Errorf("a %v\nb %v", a, b)
  }
}