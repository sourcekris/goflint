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

func TestMod(t *testing.T) {
  a := NewFmpz(64)
  b := NewFmpz(5)
  c := NewFmpz(0)
  expected := NewFmpz(4)
  
  if c.Mod(a,b).Cmp(expected) != 0 {
    t.Errorf("Expected a % b == 5 but got something else: %d\n", a)
  } 
}

func TestAbs(t *testing.T) {
  a := NewFmpz(-64)
  expected := NewFmpz(64)
  
  if a.Abs(a).Cmp(expected) != 0 {
    t.Errorf("Expected a.Abs(a) == 64 but got something else: %d\n", a)
  } 
}

func TestAdd(t *testing.T) {
  a := NewFmpz(60)
  b := NewFmpz(4)
  expected := NewFmpz(64)
  
  if a.Add(a, b).Cmp(expected) != 0 {
    t.Errorf("Expected a.Add(a, b) == 64 but got something else: %d\n", a)
  } 
}

func TestSub(t *testing.T) {
  a := NewFmpz(68)
  b := NewFmpz(4)
  expected := NewFmpz(64)
  
  if a.Sub(a, b).Cmp(expected) != 0 {
    t.Errorf("Expected a.Sub(a, b) == 64 but got something else: %d\n", a)
  } 
}

func TestMul(t *testing.T) {
  a := NewFmpz(8)
  b := NewFmpz(8)
  expected := NewFmpz(64)
  
  if a.Mul(a, b).Cmp(expected) != 0 {
    t.Errorf("Expected a.Mul(a, b) == 64 but got something else: %d\n", a)
  } 
}

func TestDiv(t *testing.T) {
  num := NewFmpz(64)
  den := NewFmpz(8)
  expected := NewFmpz(8)
  
  if expected.Div(num,den).Cmp(expected) != 0 {
    t.Errorf("Expected num / den == 8 but got something else: %d\n", expected)
  } 
}

func TestQuo(t *testing.T) {
  num := NewFmpz(65)
  den := NewFmpz(8)
  expected := NewFmpz(8)
  
  if expected.Quo(num,den).Cmp(expected) != 0 {
    t.Errorf("Expected expectected.Quo(num, den) == 8 but got something else: %d\n", expected)
  } 
}

func TestGCD(t *testing.T) {
  a := NewFmpz(15)
  b := NewFmpz(155)
  expected := NewFmpz(5)
  
  if expected.GCD(a,b).Cmp(expected) != 0 {
    t.Errorf("Expected expectected.GCD(a, b) == 5 but got something else: %d\n", expected)
  } 
}

func TestSetString(t *testing.T) {
  expected := NewFmpz(65293409233)
  num, result := new(Fmpz).SetString("65293409233",10)
  if !result {
    t.Error("SetString just failed completely. Oh noes.")
  }

  if num.Cmp(expected) != 0 {
    t.Errorf("Expected 65293409233 but got something else: %d\n", num)
  } 
}