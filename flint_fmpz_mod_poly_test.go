package goflint

import "testing"

func TestFmpzModPolyString(t *testing.T) {
	for _, tc := range []struct {
		name string
		m    *FmpzModPoly
		want string
	}{
		{
			name: "f(x)=5x^3+2x+1 in (Z/6Z)[x]",
			m:    NewFmpzModPoly(NewFmpzModCtx(NewFmpz(6))).SetCoeffUI(0, 1).SetCoeffUI(1, 2).SetCoeffUI(2, 0).SetCoeffUI(3, 5),
			want: "5*x^3+2*x+1",
		},
	} {
		got := tc.m.String()
		if got != tc.want {
			t.Errorf("String() %s want / got mismatch: %v / %v", tc.name, tc.want, got)
		}
	}
}

func TestFmpzModPolyStringSimple(t *testing.T) {
	for _, tc := range []struct {
		name string
		m    *FmpzModPoly
		want string
	}{
		{
			name: "f(x)=5x^3+2x+1 in (Z/6Z)[x]",
			m:    NewFmpzModPoly(NewFmpzModCtx(NewFmpz(6))).SetCoeffUI(0, 1).SetCoeffUI(1, 2).SetCoeffUI(2, 0).SetCoeffUI(3, 5),
			want: "4 6  1 2 0 5",
		},
	} {

		got := tc.m.StringSimple()
		if got != tc.want {
			t.Errorf("StringSimple() %s want / got mismatch: %v / %v", tc.name, tc.want, got)
		}
	}
}

func TestFmpzModPolySetString(t *testing.T) {
	for _, tc := range []struct {
		name string
		poly string
		want *FmpzModPoly
	}{
		{
			name: "f(x)=5x^3+2x+1 in (Z/6Z)[x]",
			poly: "4 6  1 2 0 5",
			want: NewFmpzModPoly(NewFmpzModCtx(NewFmpz(6))).SetCoeffUI(0, 1).SetCoeffUI(1, 2).SetCoeffUI(2, 0).SetCoeffUI(3, 5),
		},
	} {

		got, err := SetString(tc.poly)
		if err != nil {
			t.Errorf("SetString() got error when not expected: %v", err)
		}

		if !got.Equal(tc.want) {
			t.Errorf("SetString() %s want / got mismatch: %v / %v", tc.name, tc.want, got)
		}
	}
}

func TestFmpzModPolyPow(t *testing.T) {
	for _, tc := range []struct {
		name string
		m    *FmpzModPoly
		e    int
		want string
	}{
		{
			name: "f(x)=5x^3+2x+1 in (Z/6Z)[x]",
			m:    NewFmpzModPoly(NewFmpzModCtx(NewFmpz(6))).SetCoeffUI(0, 1).SetCoeffUI(1, 2).SetCoeffUI(2, 0).SetCoeffUI(3, 5),
			e:    3,
			want: "5*x^9+3*x^6+5*x^3+1",
		},
	} {

		z := tc.m.Pow(tc.m, tc.e)
		got := z.String()
		if got != tc.want {
			t.Errorf("Pow() %s want / got mismatch: %v / %v", tc.name, tc.want, got)
		}
	}
}

func TestFmpzModPolyGCD(t *testing.T) {
	for _, tc := range []struct {
		name string
		m    *FmpzModPoly
		want string
	}{
		{
			name: "f(x)=5x^3+2x+1 in (Z/6Z)[x]",
			m:    NewFmpzModPoly(NewFmpzModCtx(NewFmpz(6))).SetCoeffUI(0, 1).SetCoeffUI(1, 2).SetCoeffUI(2, 0).SetCoeffUI(3, 5),
			want: "x^6+2*x^4+4*x^3+4*x^2+4*x+1",
		},
	} {

		y := NewFmpzModPoly(NewFmpzModCtx(NewFmpz(6))).Mul(tc.m, tc.m)
		z := NewFmpzModPoly(NewFmpzModCtx(NewFmpz(6))).Mul(tc.m, tc.m)
		z.Mul(z, tc.m)

		g := NewFmpzModPoly(NewFmpzModCtx(NewFmpz(6))).GCD(y, z)

		got := g.String()
		if got != tc.want {
			t.Errorf("GCD() %s want / got mismatch: %v / %v", tc.name, tc.want, got)
		}
	}
}

func TestFmpzModPolyGetMod(t *testing.T) {
	for _, tc := range []struct {
		name string
		m    *FmpzModPoly
		want *Fmpz
	}{
		{
			name: "f(x)=5x^3+2x+1 in (Z/6Z)[x] has modulus of 6",
			m:    NewFmpzModPoly(NewFmpzModCtx(NewFmpz(6))).SetCoeffUI(0, 1).SetCoeffUI(1, 2).SetCoeffUI(2, 0).SetCoeffUI(3, 5),
			want: NewFmpz(6),
		},
	} {

		got := tc.m.GetMod()
		if got.Cmp(tc.want) != 0 {
			t.Errorf("GetMod() %s want / got mismatch: %v / %v", tc.name, tc.want, got)
		}
	}
}

func TestFmpzModPolyLen(t *testing.T) {
	for _, tc := range []struct {
		name string
		m    *FmpzModPoly
		want int
	}{
		{
			name: "f(x)=5x^3+2x+1 in (Z/6Z)[x] has len 4",
			m:    NewFmpzModPoly(NewFmpzModCtx(NewFmpz(6))).SetCoeffUI(0, 1).SetCoeffUI(1, 2).SetCoeffUI(2, 0).SetCoeffUI(3, 5),
			want: 4,
		},
	} {

		got := tc.m.Len()
		if got != tc.want {
			t.Errorf("Len() %s want / got mismatch: %v / %v", tc.name, tc.want, got)
		}
	}
}

func TestFmpzModPolyGetCoeff(t *testing.T) {
	for _, tc := range []struct {
		name string
		m    *FmpzModPoly
		c    int
		want *Fmpz
	}{
		{
			name: "f(x)=5x^3+2x+1 in (Z/6Z)[x] 3rd coefficient of 5",
			m:    NewFmpzModPoly(NewFmpzModCtx(NewFmpz(6))).SetCoeffUI(0, 1).SetCoeffUI(1, 2).SetCoeffUI(2, 0).SetCoeffUI(3, 5),
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

func TestFmpzModPolyGetCoeffs(t *testing.T) {
	for _, tc := range []struct {
		name string
		m    *FmpzModPoly
		want []*Fmpz
	}{
		{
			name: "f(x)=5x^3+2x+1 in (Z/6Z)[x] coeffs are 1,2,0,5",
			m:    NewFmpzModPoly(NewFmpzModCtx(NewFmpz(6))).SetCoeffUI(0, 1).SetCoeffUI(1, 2).SetCoeffUI(2, 0).SetCoeffUI(3, 5),
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

// func TestFmpzModPolyAdd(t *testing.T) {
// 	for _, tc := range []struct {
// 		name string
// 		m1   *FmpzModPoly
// 		m2   *FmpzModPoly
// 		want string
// 	}{
// 		{
// 			name: "f(x)=5x^3+2x+1 in (Z/6Z)[x] + x^3+3x^2+2 = 3*x^2+2*x+3",
// 			m1:   NewFmpzModPoly(NewFmpz(6)).SetCoeffUI(0, 1).SetCoeffUI(1, 2).SetCoeffUI(2, 0).SetCoeffUI(3, 5),
// 			m2:   NewFmpzModPoly(NewFmpz(6)).SetCoeffUI(0, 2).SetCoeffUI(1, 0).SetCoeffUI(2, 3).SetCoeffUI(3, 1),
// 			want: "3*x^2+2*x+3",
// 		},
// 	} {

// 		got := NewFmpzModPoly(tc.m1.GetMod()).Add(tc.m1, tc.m2).String()

// 		if got != tc.want {
// 			t.Errorf("Add() %s want / got mismatch: %v / %v", tc.name, tc.want, got)
// 		}

// 	}
// }

// func TestFmpzModPolySub(t *testing.T) {
// 	for _, tc := range []struct {
// 		name string
// 		m1   *FmpzModPoly
// 		m2   *FmpzModPoly
// 		want string
// 	}{
// 		{
// 			name: "f(x)=5x^3+2x+1 in (Z/6Z)[x] - x^3+3x^2+2 = 4*x^3+3*x^2+2*x+5",
// 			m1:   NewFmpzModPoly(NewFmpz(6)).SetCoeffUI(0, 1).SetCoeffUI(1, 2).SetCoeffUI(2, 0).SetCoeffUI(3, 5),
// 			m2:   NewFmpzModPoly(NewFmpz(6)).SetCoeffUI(0, 2).SetCoeffUI(1, 0).SetCoeffUI(2, 3).SetCoeffUI(3, 1),
// 			want: "4*x^3+3*x^2+2*x+5",
// 		},
// 	} {

// 		got := NewFmpzModPoly(tc.m1.GetMod()).Sub(tc.m1, tc.m2).String()

// 		if got != tc.want {
// 			t.Errorf("Sub() %s want / got mismatch: %v / %v", tc.name, tc.want, got)
// 		}

// 	}
// }
