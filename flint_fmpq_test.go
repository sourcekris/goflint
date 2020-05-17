package goflint

import "testing"

// TestNewFmpq tests assigning rationals and that the Stringer for the Fmpq type works
func TestNewFmpq(t *testing.T) {
	a := NewFmpq(3, 2)
	b := a.String()
	expected := "3/2"

	if b != expected {
		t.Errorf("Expected %v but got %v", expected, b)
	}
}

// TestSetFmpqFraction tests that assigning rationals with Fmpz numerators and denominators
// works as expected.
func TestSetFmpqFraction(t *testing.T) {
	a := NewFmpz(3)
	b := NewFmpz(2)
	c := new(Fmpq)

	c.SetFmpqFraction(a, b)

	s := c.String()

	expected := "3/2"

	if s != expected {
		t.Errorf("Expected %s but got %s", expected, s)
	}
}

func TestCmpRational(t *testing.T) {
	a := NewFmpq(2, 3)
	b := NewFmpq(2, 3)
	c := NewFmpq(5, 3)

	if a.CmpRational(b) != 0 {
		t.Error("Expected a.CmpRational(b) == 0 but it wasnt.")
	}

	if a.CmpRational(c) > 0 {
		t.Error("Expected a.CmpRational(c) < 0 (i.e. a < c) but it wasnt.")
	}

	if c.CmpRational(a) < 0 {
		t.Error("Expected c.CmpRational(a) > 0 (i.e. a > c) but it wasnt.")
	}
}

func TestGetFmpqFraction(t *testing.T) {
	a := 3
	b := 2
	c := new(Fmpq).SetFmpqFraction(NewFmpz(int64(a)), NewFmpz(int64(b)))
	num, den := c.GetFmpqFraction()

	if a != num || b != den {
		t.Errorf("GetFmpqFraction: want / got mismatched %d/%d but got %d/%d", a, b, num, den)
	}
}

func TestNumRef(t *testing.T) {
	a := NewFmpq(1, 3)
	want := 1

	got := a.NumRef()
	if got != want {
		t.Errorf("NumRef: got %v want %v", got, want)
	}
}

func TestDenRef(t *testing.T) {
	a := NewFmpq(1, 3)
	want := 3

	got := a.DenRef()
	if got != want {
		t.Errorf("DenRef: got %v want %v", got, want)
	}
}
