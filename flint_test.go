package goflint

import (
  "testing"
)

func TestSign(t *testing.T) {
  a := NewFmpz(1024)
  b := NewFmpz(-1024)
  c := NewFmpz(0)

  if c.Sign() != 0 {
    t.Error("Expected Sign of 0 to be 0.\n")
  } 

  if a.Sign() <= 0 {
    t.Error("Expected Sign of 1024 to be positive but it wasnt.\n")
  }

  if b.Sign() >= 0 {
    t.Error("Expected Sign of -1024 to be negative but it wasnt.\n")
  }
}

func TestCmp(t *testing.T) {
  a := NewFmpz(1024)
  b := NewFmpz(1024)
  c := NewFmpz(2048)

  if a.Cmp(b) != 0 {
    t.Error("Expected a.Cmp(b) == 0 but it wasnt.\n")
  } 

  if a.Cmp(c) > 0 {
    t.Error("Expected a.Cmp(c) < 0 (i.e. a < c) but it wasnt.\n")
  }

  if c.Cmp(a) < 0 {
    t.Error("Expected c.Cmp(a) > 0 (i.e. a > c) but it wasnt.\n")
  }
}

func TestNewFmpz(t *testing.T) {
  a := NewFmpz(1)
  if a.Cmp(a) != 0 {
    t.Error("Expected a == a but got something else. Major bummer.\n")
  } 
}

func TestString(t *testing.T) {
  a := NewFmpz(1024)
  aStr := a.String()

  if aStr != "1024" {
    t.Errorf("Expected aStr == 1024 but got: %s (something else?). Soz pal.\n", aStr)
  } 
}

func TestInt64(t *testing.T) {
  a := NewFmpz(64)
  b := int64(64)
  if a.Int64() != b {
    t.Error("Expected a.Int64() == 64 but got something else.\n")
  } 
}

func TestBitLen(t *testing.T) {
  a := NewFmpz(65536)

  b := a.BitLen()

  if  b != 17 {
    t.Errorf("Expected a.BitLen() == 17 but got something else: %d.\n", b)
  } 
}
