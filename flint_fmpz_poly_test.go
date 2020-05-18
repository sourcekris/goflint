package goflint

import (
	"fmt"
	"testing"
)

func TestFmpzPolyString(t *testing.T) {
	for _, tc := range []struct {
		name string
		m    *FmpzPoly
		want string
	}{
		{
			name: "f(x)=5x^3+2x+1",
			m:    NewFmpzPoly().SetCoeffUI(0, 1).SetCoeffUI(1, 2).SetCoeffUI(2, 0).SetCoeffUI(3, 5),
			want: "5*x^3+2*x+1",
		},
	} {
		got := tc.m.String()
		if got != tc.want {
			t.Errorf("String() %s want / got mismatch: %v / %v", tc.name, tc.want, got)
		}
	}
}

func TestFmpzPolyStringSimple(t *testing.T) {
	for _, tc := range []struct {
		name string
		m    *FmpzPoly
		want string
	}{
		{
			name: "f(x)=5x^3+2x+1",
			m:    NewFmpzPoly().SetCoeffUI(0, 1).SetCoeffUI(1, 2).SetCoeffUI(2, 0).SetCoeffUI(3, 5),
			want: "4  1 2 0 5",
		},
	} {

		got := tc.m.StringSimple()
		if got != tc.want {
			t.Errorf("StringSimple() %s want / got mismatch: %v / %v", tc.name, tc.want, got)
		}
	}
}

func TestFmpzPolySetString(t *testing.T) {
	for _, tc := range []struct {
		name string
		poly string
		want *FmpzPoly
	}{
		{
			name: "f(x)=5x^3+2x+1",
			poly: "4  1 2 0 5",
			want: NewFmpzPoly().SetCoeffUI(0, 1).SetCoeffUI(1, 2).SetCoeffUI(2, 0).SetCoeffUI(3, 5),
		},
	} {

		got, err := FmpzPolySetString(tc.poly)
		if err != nil {
			t.Errorf("FmpzPolySetString() got error when not expected: %v", err)
		}

		if !got.Equal(tc.want) {
			t.Errorf("FmpzPolySetString() %s want / got mismatch: %v / %v", tc.name, tc.want, got)
		}
	}
}

func TestFmpzPolyPow(t *testing.T) {
	for _, tc := range []struct {
		name string
		m    *FmpzPoly
		e    int
		want string
	}{
		{
			name: "f(x)=5x^3+2x+1",
			m:    NewFmpzPoly().SetCoeffUI(0, 1).SetCoeffUI(1, 2).SetCoeffUI(2, 0).SetCoeffUI(3, 5),
			e:    3,
			want: "125*x^9+150*x^7+75*x^6+60*x^5+60*x^4+23*x^3+12*x^2+6*x+1",
		},
	} {

		z := tc.m.Pow(tc.m, tc.e)
		got := z.String()
		if got != tc.want {
			t.Errorf("Pow() %s want / got mismatch: %v / %v", tc.name, tc.want, got)
		}
	}
}

func TestFmpzPolyGCD(t *testing.T) {
	for _, tc := range []struct {
		name string
		m    *FmpzPoly
		want string
	}{
		{
			name: "f(x)=5x^3+2x+1",
			m:    NewFmpzPoly().SetCoeffUI(0, 1).SetCoeffUI(1, 2).SetCoeffUI(2, 0).SetCoeffUI(3, 5),
			want: "25*x^6+20*x^4+10*x^3+4*x^2+4*x+1",
		},
	} {

		y := NewFmpzPoly().Mul(tc.m, tc.m)
		z := NewFmpzPoly().Mul(tc.m, tc.m)
		z.Mul(z, tc.m)

		g := NewFmpzPoly().GCD(y, z)

		got := g.String()
		if got != tc.want {
			t.Errorf("GCD() %s want / got mismatch: %v / %v", tc.name, tc.want, got)
		}
	}
}

func TestFmpzPolyLen(t *testing.T) {
	for _, tc := range []struct {
		name string
		m    *FmpzPoly
		want int
	}{
		{
			name: "f(x)=5x^3+2x+1 has len 4",
			m:    NewFmpzPoly().SetCoeffUI(0, 1).SetCoeffUI(1, 2).SetCoeffUI(2, 0).SetCoeffUI(3, 5),
			want: 4,
		},
	} {

		got := tc.m.Len()
		if got != tc.want {
			t.Errorf("Len() %s want / got mismatch: %v / %v", tc.name, tc.want, got)
		}
	}
}

func TestFmpzPolyGetCoeff(t *testing.T) {
	for _, tc := range []struct {
		name string
		m    *FmpzPoly
		c    int
		want *Fmpz
	}{
		{
			name: "f(x)=5x^3+2x+1 3rd coefficient of 5",
			m:    NewFmpzPoly().SetCoeffUI(0, 1).SetCoeffUI(1, 2).SetCoeffUI(2, 0).SetCoeffUI(3, 5),
			c:    3,
			want: NewFmpz(5),
		},
	} {

		got := tc.m.GetCoeff(tc.c)
		if got.Cmp(tc.want) != 0 {
			t.Errorf("GetCoeff() %s want / got mismatch: %v / %v", tc.name, tc.want, got)
		}
	}
}

func TestFmpzPolyGetCoeffs(t *testing.T) {
	for _, tc := range []struct {
		name string
		m    *FmpzPoly
		want []*Fmpz
	}{
		{
			name: "f(x)=5x^3+2x+1 coeffs are 1,2,0,5",
			m:    NewFmpzPoly().SetCoeffUI(0, 1).SetCoeffUI(1, 2).SetCoeffUI(2, 0).SetCoeffUI(3, 5),
			want: []*Fmpz{NewFmpz(1), NewFmpz(2), NewFmpz(0), NewFmpz(5)},
		},
	} {

		got := tc.m.GetCoeffs()

		for i, g := range got {
			if g.Cmp(tc.want[i]) != 0 {
				t.Errorf("GetCoeffs() %s want / got mismatch: %v / %v", tc.name, tc.want, got)
			}
		}
	}
}

func TestFmpzPolyAdd(t *testing.T) {
	for _, tc := range []struct {
		name string
		m1   *FmpzPoly
		m2   *FmpzPoly
		want string
	}{
		{
			name: "f(x)=5x^3+2x+1 + x^3+3x^2+2 = 3*x^2+2*x+3",
			m1:   NewFmpzPoly().SetCoeffUI(0, 1).SetCoeffUI(1, 2).SetCoeffUI(2, 0).SetCoeffUI(3, 5),
			m2:   NewFmpzPoly().SetCoeffUI(0, 2).SetCoeffUI(1, 0).SetCoeffUI(2, 3).SetCoeffUI(3, 1),
			want: "6*x^3+3*x^2+2*x+3",
		},
	} {

		got := NewFmpzPoly().Add(tc.m1, tc.m2).String()

		if got != tc.want {
			t.Errorf("Add() %s want / got mismatch: %v / %v", tc.name, tc.want, got)
		}

	}
}

func TestFmpzPolySub(t *testing.T) {
	for _, tc := range []struct {
		name string
		m1   *FmpzPoly
		m2   *FmpzPoly
		want string
	}{
		{
			name: "f(x)=5x^3+2x+1 - x^3+3x^2+2 = 4*x^3+3*x^2+2*x+5",
			m1:   NewFmpzPoly().SetCoeffUI(0, 1).SetCoeffUI(1, 2).SetCoeffUI(2, 0).SetCoeffUI(3, 5),
			m2:   NewFmpzPoly().SetCoeffUI(0, 2).SetCoeffUI(1, 0).SetCoeffUI(2, 3).SetCoeffUI(3, 1),
			want: "4*x^3-3*x^2+2*x-1",
		},
	} {

		got := NewFmpzPoly().Sub(tc.m1, tc.m2).String()

		if got != tc.want {
			t.Errorf("Sub() %s want / got mismatch: %v / %v", tc.name, tc.want, got)
		}

	}
}

func TestFmpzPolyFactor(t *testing.T) {
	for _, tc := range []struct {
		name string
		poly *FmpzPoly
		want string
	}{
		{
			name: "f(x)=10x^3+5x^2+5x+5 factors to 5, (2*x^3+x^2+x+1, 1)",
			poly: NewFmpzPoly().SetCoeffUI(0, 5).SetCoeffUI(1, 5).SetCoeffUI(2, 5).SetCoeffUI(3, 10),
			want: "5, (2*x^3+x^2+x+1, 1)",
		},
	} {

		fac := tc.poly.Factor()
		got := fac.GetCoeff().String()
		for i := 0; i < fac.Len(); i++ {
			p := fac.GetPoly(i)
			e := fac.GetExp(i)
			got = fmt.Sprintf("%s, (%v, %d)", got, p, e)
		}

		if got != tc.want {
			t.Errorf("Factor() %s want / got mismatch: %v / %v", tc.name, tc.want, got)
		}
	}
}
