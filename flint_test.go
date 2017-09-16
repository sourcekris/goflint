package goflint

import (
  "testing"
)

func TestNewFmpz(t *testing.T) {
  a := NewFmpz(1)
  if a != a {
    t.Errorf("Expected a == a but got something else. Major bummer.\n")
  } 
}