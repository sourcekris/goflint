package goflint

import (
	"bytes"
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

func TestCmpRational(t *testing.T) {
	a := NewFmpq(2, 3)
	b := NewFmpq(2, 3)
	c := NewFmpq(5, 3)

	if a.CmpRational(b) != 0 {
		t.Error("Expected a.CmpRational(b) == 0 but it wasnt.\n")
	}

	if a.CmpRational(c) > 0 {
		t.Error("Expected a.CmpRational(c) < 0 (i.e. a < c) but it wasnt.\n")
	}

	if c.CmpRational(a) < 0 {
		t.Error("Expected c.CmpRational(a) > 0 (i.e. a > c) but it wasnt.\n")
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

	if b != 17 {
		t.Errorf("Expected a.BitLen() == 17 but got something else: %v.\n", b)
	}
}

func TestMod(t *testing.T) {
	a := NewFmpz(64)
	b := NewFmpz(5)
	c := NewFmpz(0)
	expected := NewFmpz(4)

	if c.Mod(a, b).Cmp(expected) != 0 {
		t.Errorf("Expected a mod b == 5 but got something else: %v\n", a)
	}
}

func TestAbs(t *testing.T) {
	a := NewFmpz(-64)
	expected := NewFmpz(64)

	if a.Abs(a).Cmp(expected) != 0 {
		t.Errorf("Expected a.Abs(a) == 64 but got something else: %v\n", a)
	}
}

func TestAdd(t *testing.T) {
	a := NewFmpz(60)
	b := NewFmpz(4)
	expected := NewFmpz(64)

	if a.Add(a, b).Cmp(expected) != 0 {
		t.Errorf("Expected a.Add(a, b) == 64 but got something else: %v\n", a)
	}
}

func TestSub(t *testing.T) {
	a := NewFmpz(68)
	b := NewFmpz(4)
	expected := NewFmpz(64)

	if a.Sub(a, b).Cmp(expected) != 0 {
		t.Errorf("Expected a.Sub(a, b) == 64 but got something else: %v\n", a)
	}
}

func TestMul(t *testing.T) {
	a := NewFmpz(8)
	b := NewFmpz(8)
	expected := NewFmpz(64)

	if a.Mul(a, b).Cmp(expected) != 0 {
		t.Errorf("Expected a.Mul(a, b) == 64 but got something else: %v\n", a)
	}
}

func TestMulRational(t *testing.T) {
	a := NewFmpq(2, 3)
	b := NewFmpz(8)
	expected := NewFmpq(16, 3)

	if a.MulRational(a, b).CmpRational(expected) != 0 {
		t.Errorf("Expected a.MulRational(a, b) == %s but got something else: %s\n", expected, a)
	}
}

func TestDiv(t *testing.T) {
	num := NewFmpz(64)
	den := NewFmpz(8)
	expected := NewFmpz(8)

	if expected.Div(num, den).Cmp(expected) != 0 {
		t.Errorf("Expected num / den == 8 but got something else: %v\n", expected)
	}
}

func TestDivMod(t *testing.T) {
	a := NewFmpz(11231231)
	b := NewFmpz(14541)
	n := NewFmpz(666)

	xE := NewFmpz(772)
	yE := NewFmpz(5579)

	x, y := a.DivMod(a, b, n)

	if x.Cmp(xE) != 0 && y.Cmp(yE) != 0 {
		t.Errorf("Expected %v and %v but got %v and %v\n", xE, yE, x, y)
	}

	a.SetInt64(11231231)
	b.SetInt64(-984)
	xE.SetInt64(-11413)
	yE.SetInt64(839)

	x, y = a.DivMod(a, b, n)

	if x.Cmp(xE) != 0 && y.Cmp(yE) != 0 {
		t.Errorf("Expected %v and %v but got %v and %v\n", xE, yE, x, y)
	}
}

func TestQuo(t *testing.T) {
	num := NewFmpz(65)
	den := NewFmpz(8)
	expected := NewFmpz(8)

	if expected.Quo(num, den).Cmp(expected) != 0 {
		t.Errorf("Expected expectected.Quo(num, den) == 8 but got something else: %v\n", expected)
	}
}

func TestGCD(t *testing.T) {
	a := NewFmpz(15)
	b := NewFmpz(155)
	expected := NewFmpz(5)

	if expected.GCD(a, b).Cmp(expected) != 0 {
		t.Errorf("Expected expectected.GCD(a, b) == 5 but got something else: %v\n", expected)
	}
}

func TestSetString(t *testing.T) {
	expected := NewFmpz(65293409233)
	num, result := new(Fmpz).SetString("65293409233", 10)
	if !result {
		t.Error("SetString just failed completely. Oh noes.")
	}

	if num.Cmp(expected) != 0 {
		t.Errorf("Expected 65293409233 but got something else: %v\n", num)
	}
}

func TestBytes(t *testing.T) {
	a := NewFmpz(8746238)
	b := []byte{0x85, 0x74, 0xfe}
	c := a.Bytes()

	if !bytes.Equal(b, c) {
		t.Errorf("Expected %x but got %x.\n", b, c)
	}
}

func TestSetBytes(t *testing.T) {
	a := []byte{0x85, 0x74, 0xfe}
	b := new(Fmpz).SetBytes(a)
	c := NewFmpz(8746238)

	if b.Cmp(c) != 0 {
		t.Errorf("Expected %v but got %v\n", c, b)
	}
}

func TestExp(t *testing.T) {
	a := NewFmpz(11231231)
	b := NewFmpz(55)
	n := NewFmpz(6611116)

	expected := NewFmpz(4221059)

	z := a.Exp(a, b, n)

	if z.Cmp(expected) != 0 {
		t.Errorf("Exp(): expected %v but got %v\n", expected, z)
	}

	// Exp returns the result but also mutates the receiver.
	if a.Cmp(expected) != 0 {
		t.Errorf("Exp(): expected %v but got %v\n", expected, z)
	}
}

func TestPow(t *testing.T) {
	a := NewFmpz(11231231)
	b := NewFmpz(55)
	n := NewFmpz(6611116)

	expected := NewFmpz(4221059)

	z := a.Pow(a, b, n)

	if z.Cmp(expected) != 0 {
		t.Errorf("Pow(): expected %v but got %v\n", expected, z)
	}

	// Pow returns the result but also mutates the receiver.
	if a.Cmp(expected) != 0 {
		t.Errorf("Pow(): expected %v but got %v\n", expected, z)
	}
}

func TestAnd(t *testing.T) {
	a := NewFmpz(11231231)
	b := NewFmpz(52115)
	expected := NewFmpz(19347)

	z := a.And(a, b)

	if z.Cmp(expected) != 0 {
		t.Errorf("Expected %v but got %v\n", expected, z)
	}
}

func TestSqrt(t *testing.T) {
	a := NewFmpz(4096)
	expected := NewFmpz(64)
	b := new(Fmpz)
	b.doinit()
	b.Sqrt(a)

	if b.Cmp(expected) != 0 {
		t.Errorf("Expected %v but got %v\n", expected, b)
	}
}

func TestRoot(t *testing.T) {
	a := NewFmpz(4096)
	expected := NewFmpz(64)

	b := new(Fmpz)
	b.Root(a, 2)

	if b.Cmp(expected) != 0 {
		t.Errorf("Expected %v but got %v\n", expected, b)
	}
}

// TestNewFmpq tests assigning rationals and that the Stringer for the Fmpq type works
func TestNewFmpq(t *testing.T) {
	a := NewFmpq(3, 2)
	b := a.String()
	expected := "3/2"

	if b != expected {
		t.Errorf("Expected %v but got %v\n", expected, b)
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
		t.Errorf("Expected %s but got %s\n", expected, s)
	}
}

func TestGetFmpqFraction(t *testing.T) {
	a := NewFmpz(3)
	b := NewFmpz(2)
	c := new(Fmpq)

	c.SetFmpqFraction(a, b)

	num := new(Fmpz)
	den := new(Fmpz)

	c.GetFmpqFraction(num, den)

	if a.Cmp(num) != 0 || b.Cmp(den) != 0 {
		t.Errorf("Expected %s/%s but got %s/%s\n", a, b, num, den)
	}
}

func TestIsStrongProbabPrime(t *testing.T) {
	tt := []struct {
		name    string
		test    string
		want    int
		wantErr bool
	}{
		{
			name: "large prime is prime",
			test: "863653476616376575308866344984576466644942572246900013156919",
			want: 1,
		},
		{
			name:    "large composite is not prime",
			test:    "833810193564967701912362955539789451139872863794534923259743419423089229206473091408403560311191545764221310666338878019",
			wantErr: true,
		},
	}

	base := NewFmpz(10)

	for _, tc := range tt {
		p, ok := new(Fmpz).SetString(tc.test, 10)
		if !ok && !tc.wantErr {
			t.Errorf("TestIsStrongProbabPrime(): %s failed converting test to Fmpz: %v", tc.name, tc.test)
		}
		got := p.IsStrongProbabPrime(base)
		if got != 1 && !tc.wantErr {
			t.Errorf("IsStrongProbabPrime(): %s expected prime got %v\n", tc.name, got)
		}

		if got != 0 && tc.wantErr {
			t.Errorf("IsStrongProbabPrime(): %s expected not-prime got %v\n", tc.name, got)
		}
	}
}

func TestIsProbabPrimeLucas(t *testing.T) {
	tt := []struct {
		name    string
		test    string
		want    int
		wantErr bool
	}{
		{
			name: "large prime is prime",
			test: "863653476616376575308866344984576466644942572246900013156919",
			want: 1,
		},
		{
			name:    "large composite is not prime",
			test:    "833810193564967701912362955539789451139872863794534923259743419423089229206473091408403560311191545764221310666338878019",
			wantErr: true,
		},
	}

	for _, tc := range tt {
		p, ok := new(Fmpz).SetString(tc.test, 10)
		if !ok && !tc.wantErr {
			t.Errorf("IsProbabPrimeLucas(): %s failed converting test to Fmpz: %v", tc.name, tc.test)
		}
		got := p.IsProbabPrimeLucas()
		if got != 1 && !tc.wantErr {
			t.Errorf("IsProbabPrimeLucas(): %s expected prime got %v\n", tc.name, got)
		}

		if got != 0 && tc.wantErr {
			t.Errorf("IsProbabPrimeLucas(): %s expected not-prime got %v\n", tc.name, got)
		}
	}
}

func TestIsProbabPrimeBPSW(t *testing.T) {
	tt := []struct {
		name    string
		test    string
		want    int
		wantErr bool
	}{
		{
			name: "large prime is prime",
			test: "863653476616376575308866344984576466644942572246900013156919",
			want: 1,
		},
		{
			name:    "large composite is not prime",
			test:    "833810193564967701912362955539789451139872863794534923259743419423089229206473091408403560311191545764221310666338878019",
			wantErr: true,
		},
	}

	for _, tc := range tt {
		p, ok := new(Fmpz).SetString(tc.test, 10)
		if !ok && !tc.wantErr {
			t.Errorf("TestIsProbabPrimeBPSW(): %s failed converting test to Fmpz: %v", tc.name, tc.test)
		}
		got := p.IsProbabPrimeBPSW()
		if got != 1 && !tc.wantErr {
			t.Errorf("IsProbabPrimeBPSW(): %s expected prime got %v\n", tc.name, got)
		}

		if got != 0 && tc.wantErr {
			t.Errorf("IsProbabPrimeBPSW(): %s expected not-prime got %v\n", tc.name, got)
		}
	}
}

func TestIsProbabPrime(t *testing.T) {
	tt := []struct {
		name    string
		test    string
		want    int
		wantErr bool
	}{
		{
			name: "large prime is prime",
			test: "863653476616376575308866344984576466644942572246900013156919",
			want: 1,
		},
		{
			name:    "large composite is not prime",
			test:    "833810193564967701912362955539789451139872863794534923259743419423089229206473091408403560311191545764221310666338878019",
			wantErr: true,
		},
	}

	for _, tc := range tt {
		p, ok := new(Fmpz).SetString(tc.test, 10)
		if !ok && !tc.wantErr {
			t.Errorf("TestIsProbabPrime(): %s failed converting test to Fmpz: %v", tc.name, tc.test)
		}
		got := p.IsProbabPrime()
		if got != 1 && !tc.wantErr {
			t.Errorf("IsProbabPrime(): %s expected prime got %v\n", tc.name, got)
		}

		if got != 0 && tc.wantErr {
			t.Errorf("IsProbabPrime(): %s expected not-prime got %v\n", tc.name, got)
		}
	}
}

func TestIsProbabPrimePseudosquare(t *testing.T) {
	tt := []struct {
		name    string
		test    string
		want    int
		wantErr bool
	}{
		{
			name:    "prime is too large to test",
			test:    "863653476616376575308866344984576466644942572246900013156919",
			want:    -1,
			wantErr: true,
		},
		{
			name: "largish prime is prime",
			test: "18446744073709551557",
			want: 1,
		},

		{
			name:    "composite is not prime",
			test:    "80009000",
			wantErr: true,
		},
	}

	for _, tc := range tt {
		p, ok := new(Fmpz).SetString(tc.test, 10)
		if !ok && !tc.wantErr {
			t.Errorf("TestIsProbabPrimePseudosquare(): %s failed converting test to Fmpz: %v", tc.name, tc.test)
		}
		got := p.IsProbabPrimePseudosquare()
		if got != 1 && !tc.wantErr {
			t.Errorf("IsProbabPrimePseudosquare(): %s expected prime got %v\n", tc.name, got)
		}

		if got < 0 && tc.want != -1 {
			t.Errorf("IsProbabPrimePseudosquare(): %s expected not-prime got %v\n", tc.name, got)
		}
	}
}

func TestBits(t *testing.T) {
	tt := []struct {
		name string
		n    string
		want int
	}{
		{
			name: "zero has zero bits",
			n:    "0",
			want: 0,
		},
		{
			name: "one has one bits",
			n:    "1",
			want: 1,
		},
		{
			name: "this number has 123 bits according to other sources",
			n:    "8237492387492374928472987492874913111",
			want: 123,
		},
	}
	for _, tc := range tt {
		n, ok := new(Fmpz).SetString(tc.n, 10)
		if !ok {
			t.Errorf("TestBits(): %s failed converting n to Fmpz: %v", tc.name, tc.n)
		}
		got := n.Bits()
		if got != tc.want {
			t.Errorf("Bits(): %s expected %v got %v\n", tc.name, tc.want, got)
		}
	}
}
func TestTstBit(t *testing.T) {
	tt := []struct {
		name string
		n    string
		bit  int
		want int
	}{
		{
			name: "zero has zero bits",
			n:    "0",
			bit:  0,
			want: 0,
		},
		{
			name: "one 1 in the zeroth bit",
			n:    "1",
			bit:  0,
			want: 1,
		},
		{
			name: "this number has 1 at bit 6",
			n:    "8237492387492374928472987492874913111",
			bit:  6,
			want: 1,
		},
	}
	for _, tc := range tt {
		n, ok := new(Fmpz).SetString(tc.n, 10)
		if !ok {
			t.Errorf("TestTstBit(): %s failed converting n to Fmpz: %v", tc.name, tc.n)
		}
		got := n.TstBit(tc.bit)
		if got != tc.want {
			t.Errorf("TstBit(): %s expected %v got %v\n", tc.name, tc.want, got)
		}
	}
}
