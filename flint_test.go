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
		t.Error("Expected Sign of 0 to be 0.")
	}

	if a.Sign() <= 0 {
		t.Error("Expected Sign of 1024 to be positive but it wasnt.")
	}

	if b.Sign() >= 0 {
		t.Error("Expected Sign of -1024 to be negative but it wasnt.")
	}
}

func TestCmp(t *testing.T) {
	a := NewFmpz(1024)
	b := NewFmpz(1024)
	c := NewFmpz(2048)

	if a.Cmp(b) != 0 {
		t.Error("Expected a.Cmp(b) == 0 but it wasnt.")
	}

	if a.Cmp(c) > 0 {
		t.Error("Expected a.Cmp(c) < 0 (i.e. a < c) but it wasnt.")
	}

	if c.Cmp(a) < 0 {
		t.Error("Expected c.Cmp(a) > 0 (i.e. a > c) but it wasnt.")
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

func TestEquals(t *testing.T) {
	a := NewFmpz(1024)
	b := NewFmpz(1024)
	c := NewFmpz(2048)

	if !a.Equals(b) {
		t.Error("Expected a.Equals(b) but it wasnt.")
	}

	if a.Equals(c) {
		t.Error("Expected !a.Equals(c) but it wasnt.")
	}
}

func TestIsZero(t *testing.T) {
	a := NewFmpz(0)
	b := NewFmpz(1024)

	if !a.IsZero() {
		t.Error("Expected a.IsZero() but it wasnt.")
	}

	if b.IsZero() {
		t.Error("Expected !b.IsZero() but it wasnt.")
	}
}

func TestNewFmpz(t *testing.T) {
	a := NewFmpz(1)
	if a.Cmp(a) != 0 {
		t.Error("Expected a == a but got something else. Major bummer.")
	}
}

func TestString(t *testing.T) {
	a := NewFmpz(1024)
	aStr := a.String()

	if aStr != "1024" {
		t.Errorf("Expected aStr == 1024 but got: %s (something else?). Soz pal.", aStr)
	}
}

func TestInt64(t *testing.T) {
	a := NewFmpz(64)
	b := int64(64)
	if a.Int64() != b {
		t.Error("Expected a.Int64() == 64 but got something else.")
	}
}

func TestBitLen(t *testing.T) {
	a := NewFmpz(65536)

	b := a.BitLen()

	if b != 17 {
		t.Errorf("Expected a.BitLen() == 17 but got something else: %v.", b)
	}
}

func TestMod(t *testing.T) {
	a := NewFmpz(64)
	b := NewFmpz(5)
	expected := NewFmpz(4)

	if new(Fmpz).Mod(a, b).Cmp(expected) != 0 {
		t.Errorf("Expected a mod b == 5 but got something else: %v", a)
	}
}

func TestModZ(t *testing.T) {
	a := NewFmpz(64)
	b := NewFmpz(5)
	expected := NewFmpz(4)

	if new(Fmpz).SetInt64(0).AddZ(a).ModZ(b).Cmp(expected) != 0 {
		t.Errorf("Expected a mod b == 5 but got something else: %v", a)
	}
}

func TestAbs(t *testing.T) {
	a := NewFmpz(-64)
	expected := NewFmpz(64)

	if a.Abs(a).Cmp(expected) != 0 {
		t.Errorf("Expected a.Abs(a) == 64 but got something else: %v", a)
	}
}

func TestAdd(t *testing.T) {
	a := NewFmpz(60)
	b := NewFmpz(4)
	expected := NewFmpz(64)

	if a.Add(a, b).Cmp(expected) != 0 {
		t.Errorf("Expected a.Add(a, b) == 64 but got something else: %v", a)
	}
}

func TestSub(t *testing.T) {
	a := NewFmpz(68)
	b := NewFmpz(4)
	expected := NewFmpz(64)

	if a.Sub(a, b).Cmp(expected) != 0 {
		t.Errorf("Expected a.Sub(a, b) == 64 but got something else: %v", a)
	}
}
func TestSubRMpz(t *testing.T) {
	a := NewMpz(7)
	b := NewMpz(8)
	n := NewMpz(10000000000)
	want := NewMpz(9999999999)
	got := a.SubRMpz(b, n)

	if got.Cmp(want) != 0 {
		t.Errorf("a.SubRMpz(b, n) want / got mismatch: %v / %v", want, got)
	}
}

func TestSubZ(t *testing.T) {
	a := NewFmpz(68)
	b := NewFmpz(4)
	expected := NewFmpz(64)

	if a.SubZ(b).Cmp(expected) != 0 {
		t.Errorf("Expected a.SubZ(b) == 64 but got something else: %v", a)
	}
}

func TestSubI(t *testing.T) {
	a := NewFmpz(68)
	b := 4
	expected := NewFmpz(64)

	if a.SubI(b).Cmp(expected) != 0 {
		t.Errorf("Expected a.SubI(b) == 64 but got something else: %v", a)
	}
}

func TestMul(t *testing.T) {
	a := NewFmpz(8)
	b := NewFmpz(8)
	expected := NewFmpz(64)

	if a.Mul(a, b).Cmp(expected) != 0 {
		t.Errorf("Expected a.Mul(a, b) == 64 but got something else: %v", a)
	}
}

func TestMulZ(t *testing.T) {
	a := NewFmpz(8)
	b := NewFmpz(8)
	expected := NewFmpz(64)

	if a.MulZ(b).Cmp(expected) != 0 {
		t.Errorf("Expected a.MulZ(b) == 64 but got something else: %v", a)
	}
}
func TestMulI(t *testing.T) {
	a := NewFmpz(8)
	b := 8
	expected := NewFmpz(64)

	if a.MulI(b).Cmp(expected) != 0 {
		t.Errorf("Expected a.MulI(b) == 64 but got something else: %v", a)
	}
}

func TestMulRMpzI(t *testing.T) {
	a := NewMpz(700000)
	b := NewMpz(800000)
	n := NewMpz(100000000000)
	want := NewMpz(60000000000)
	got := a.MulRMpz(b, n)
	if got.Cmp(want) != 0 {
		t.Errorf("a.MulRMpz(b, n) mismatched want / got: %v / %v", want, got)
	}
}

func TestMulRational(t *testing.T) {
	a := NewFmpq(2, 3)
	b := NewFmpz(8)
	expected := NewFmpq(16, 3)

	if a.MulRational(a, b).CmpRational(expected) != 0 {
		t.Errorf("Expected a.MulRational(a, b) == %s but got something else: %s", expected, a)
	}
}

func TestDiv(t *testing.T) {
	num := NewFmpz(64)
	den := NewFmpz(8)
	expected := NewFmpz(8)

	if expected.Div(num, den).Cmp(expected) != 0 {
		t.Errorf("Expected num / den == 8 but got something else: %v", expected)
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
		t.Errorf("Expected %v and %v but got %v and %v", xE, yE, x, y)
	}

	a.SetInt64(11231231)
	b.SetInt64(-984)
	xE.SetInt64(-11413)
	yE.SetInt64(839)

	x, y = a.DivMod(a, b, n)

	if x.Cmp(xE) != 0 && y.Cmp(yE) != 0 {
		t.Errorf("Expected %v and %v but got %v and %v", xE, yE, x, y)
	}
}

func TestDivR(t *testing.T) {
	tt := []struct {
		name string
		z    *Fmpz
		y    *Fmpz
		n    string
		want string
	}{
		{
			name: "test from the sage documentation 3 / 7 mod 100000000000",
			z:    NewFmpz(3),
			y:    NewFmpz(7),
			n:    "100000000000",
			want: "71428571429",
		},
		{
			name: "j/(1728-j) in ring Integers(n)",
			z:    NewFmpz(-32768),
			y:    NewFmpz(34496),
			n:    "1444329727510154393553799612747635457542181563961160832013134005088873165794135221",
			want: "418024930411102199247481891630113416283080378437738571046472921695480916259527076",
		},
	}

	for _, tc := range tt {
		n, _ := new(Fmpz).SetString(tc.n, 10)
		want, _ := new(Fmpz).SetString(tc.want, 10)
		got := tc.z.DivR(tc.y, n)
		if got.Cmp(want) != 0 {
			t.Errorf("%s: want / got mismatched: %v / %v", tc.name, want, got)
		}
	}
}

func TestQuo(t *testing.T) {
	num := NewFmpz(65)
	den := NewFmpz(8)
	expected := NewFmpz(8)

	if expected.Quo(num, den).Cmp(expected) != 0 {
		t.Errorf("Expected expectected.Quo(num, den) == 8 but got something else: %v", expected)
	}
}

func TestGCD(t *testing.T) {
	a := NewFmpz(15)
	b := NewFmpz(155)
	expected := NewFmpz(5)

	if expected.GCD(a, b).Cmp(expected) != 0 {
		t.Errorf("Expected expectected.GCD(a, b) == 5 but got something else: %v", expected)
	}
}

func TestSetString(t *testing.T) {
	expected := NewFmpz(65293409233)
	num, result := new(Fmpz).SetString("65293409233", 10)
	if !result {
		t.Error("SetString just failed completely. Oh noes.")
	}

	if num.Cmp(expected) != 0 {
		t.Errorf("Expected 65293409233 but got something else: %v", num)
	}
}

func TestBytes(t *testing.T) {
	a := NewFmpz(8746238)
	b := []byte{0x85, 0x74, 0xfe}
	c := a.Bytes()

	if !bytes.Equal(b, c) {
		t.Errorf("Expected %x but got %x.", b, c)
	}
}

func TestSetBytes(t *testing.T) {
	a := []byte{0x85, 0x74, 0xfe}
	b := new(Fmpz).SetBytes(a)
	c := NewFmpz(8746238)

	if b.Cmp(c) != 0 {
		t.Errorf("Expected %v but got %v", c, b)
	}
}

func TestExp(t *testing.T) {
	tt := []struct {
		name string
		a    *Fmpz
		b    *Fmpz
		n    *Fmpz
		want *Fmpz
	}{
		{
			name: "a**b % n",
			a:    NewFmpz(11231231),
			b:    NewFmpz(55),
			n:    NewFmpz(6611116),
			want: NewFmpz(4221059),
		},
		{
			name: "a**b",
			a:    NewFmpz(80),
			b:    NewFmpz(3),
			n:    nil,
			want: NewFmpz(512000),
		},
		{
			name: "a**0",
			a:    NewFmpz(11231231),
			b:    NewFmpz(0),
			n:    nil,
			want: NewFmpz(1),
		},
		{
			name: "a**-1",
			a:    NewFmpz(11231231),
			b:    NewFmpz(-1),
			n:    nil,
			want: NewFmpz(1),
		},
	}
	for _, tc := range tt {
		got := tc.a.Exp(tc.a, tc.b, tc.n)
		if got.Cmp(tc.want) != 0 {
			t.Errorf("TstExp(): %s expected %v got %v", tc.name, tc.want, got)
		}
	}
}

func TestExpZ(t *testing.T) {
	tt := []struct {
		name string
		a    *Fmpz
		b    *Fmpz
		want *Fmpz
	}{
		{
			name: "a**b",
			a:    NewFmpz(80),
			b:    NewFmpz(3),
			want: NewFmpz(512000),
		},
		{
			name: "a**0",
			a:    NewFmpz(11231231),
			b:    NewFmpz(0),
			want: NewFmpz(1),
		},
		{
			name: "a**-1",
			a:    NewFmpz(11231231),
			b:    NewFmpz(-1),
			want: NewFmpz(1),
		},
	}
	for _, tc := range tt {
		got := tc.a.ExpZ(tc.b)
		if got.Cmp(tc.want) != 0 {
			t.Errorf("TstExpZ(): %s expected %v got %v", tc.name, tc.want, got)
		}
	}
}

func TestExpI(t *testing.T) {
	tt := []struct {
		name string
		a    *Fmpz
		b    int
		want *Fmpz
	}{
		{
			name: "a**b",
			a:    NewFmpz(80),
			b:    3,
			want: NewFmpz(512000),
		},
		{
			name: "a**0",
			a:    NewFmpz(11231231),
			b:    0,
			want: NewFmpz(1),
		},
		{
			name: "a**-1",
			a:    NewFmpz(11231231),
			b:    -1,
			want: NewFmpz(1),
		},
	}
	for _, tc := range tt {
		got := tc.a.ExpI(tc.b)
		if got.Cmp(tc.want) != 0 {
			t.Errorf("TstExpI(): %s expected %v got %v", tc.name, tc.want, got)
		}
	}
}

func TestExpXY(t *testing.T) {
	tt := []struct {
		name string
		a    *Fmpz
		b    *Fmpz
		want *Fmpz
	}{
		{
			name: "a**b",
			a:    NewFmpz(80),
			b:    NewFmpz(3),
			want: NewFmpz(512000),
		},
		{
			name: "a**0",
			a:    NewFmpz(11231231),
			b:    NewFmpz(0),
			want: NewFmpz(1),
		},
		{
			name: "a**-1",
			a:    NewFmpz(11231231),
			b:    NewFmpz(-1),
			want: NewFmpz(1),
		},
	}
	for _, tc := range tt {
		got := tc.a.ExpXY(tc.a, tc.b)
		if got.Cmp(tc.want) != 0 {
			t.Errorf("TstExpXY(): %s expected %v got %v", tc.name, tc.want, got)
		}
	}
}

func TestExpXI(t *testing.T) {
	tt := []struct {
		name string
		a    *Fmpz
		b    int
		want *Fmpz
	}{
		{
			name: "a**b",
			a:    NewFmpz(80),
			b:    3,
			want: NewFmpz(512000),
		},
		{
			name: "a**0",
			a:    NewFmpz(11231231),
			b:    0,
			want: NewFmpz(1),
		},
		{
			name: "a**-1",
			a:    NewFmpz(11231231),
			b:    -1,
			want: NewFmpz(1),
		},
	}
	for _, tc := range tt {
		got := tc.a.ExpXI(tc.a, tc.b)
		if got.Cmp(tc.want) != 0 {
			t.Errorf("TstExpXI(): %s expected %v got %v", tc.name, tc.want, got)
		}
	}
}

func TestPow(t *testing.T) {
	a := NewFmpz(11231231)
	b := NewFmpz(55)
	n := NewFmpz(6611116)

	expected := NewFmpz(4221059)

	z := a.Pow(a, b, n)

	if z.Cmp(expected) != 0 {
		t.Errorf("Pow(): expected %v but got %v", expected, z)
	}

	// Pow returns the result but also mutates the receiver.
	if a.Cmp(expected) != 0 {
		t.Errorf("Pow(): expected %v but got %v", expected, z)
	}
}

func TestSquare(t *testing.T) {
	a := NewFmpz(11231231)
	expected := NewFmpz(126140549775361)

	z := a.Square()

	if z.Cmp(expected) != 0 {
		t.Errorf("Square(): expected %v but got %v", expected, z)
	}

	// Square returns the result but also mutates the receiver.
	if a.Cmp(expected) != 0 {
		t.Errorf("Square(): expected %v but got %v", expected, z)
	}
}

func TestCube(t *testing.T) {
	a := NewFmpz(1121)
	expected := NewFmpz(1408694561)

	z := a.Cube()

	if z.Cmp(expected) != 0 {
		t.Errorf("Cube(): expected %v but got %v", expected, z)
	}

	// Cube returns the result but also mutates the receiver.
	if a.Cmp(expected) != 0 {
		t.Errorf("Cube(): expected %v but got %v", expected, z)
	}
}

func TestAnd(t *testing.T) {
	a := NewFmpz(11231231)
	b := NewFmpz(52115)
	expected := NewFmpz(19347)

	z := a.And(a, b)

	if z.Cmp(expected) != 0 {
		t.Errorf("Expected %v but got %v", expected, z)
	}
}

func TestSqrt(t *testing.T) {
	a := NewFmpz(4096)
	expected := NewFmpz(64)
	b := new(Fmpz)
	b.doinit()
	b.Sqrt(a)

	if b.Cmp(expected) != 0 {
		t.Errorf("Expected %v but got %v", expected, b)
	}
}

func TestRoot(t *testing.T) {
	tt := []struct {
		name string
		n    int64
		r    int32
		want int64
	}{
		{
			name: "square root of 4096",
			n:    4096,
			r:    2,
			want: 64,
		},
		{
			name: "cube root of 4096",
			n:    4096,
			r:    3,
			want: 16,
		},
	}

	for _, tc := range tt {
		got := NewFmpz(tc.n)
		got.Root(got, tc.r)

		if got.Cmp(NewFmpz(tc.want)) != 0 {
			t.Errorf("Root() %s expected %v but got %v", tc.name, tc.want, got)
		}
	}
}

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

func TestGetFmpqFraction(t *testing.T) {
	a := 3
	b := 2
	c := new(Fmpq).SetFmpqFraction(NewFmpz(int64(a)), NewFmpz(int64(b)))
	num, den := c.GetFmpqFraction()

	if a != num || b != den {
		t.Errorf("GetFmpqFraction: want / got mismatched %d/%d but got %d/%d", a, b, num, den)
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
			t.Errorf("IsStrongProbabPrime(): %s expected prime got %v", tc.name, got)
		}

		if got != 0 && tc.wantErr {
			t.Errorf("IsStrongProbabPrime(): %s expected not-prime got %v", tc.name, got)
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
			t.Errorf("IsProbabPrimeLucas(): %s expected prime got %v", tc.name, got)
		}

		if got != 0 && tc.wantErr {
			t.Errorf("IsProbabPrimeLucas(): %s expected not-prime got %v", tc.name, got)
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
			t.Errorf("IsProbabPrimeBPSW(): %s expected prime got %v", tc.name, got)
		}

		if got != 0 && tc.wantErr {
			t.Errorf("IsProbabPrimeBPSW(): %s expected not-prime got %v", tc.name, got)
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
			t.Errorf("IsProbabPrime(): %s expected prime got %v", tc.name, got)
		}

		if got != 0 && tc.wantErr {
			t.Errorf("IsProbabPrime(): %s expected not-prime got %v", tc.name, got)
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
			t.Errorf("IsProbabPrimePseudosquare(): %s expected prime got %v", tc.name, got)
		}

		if got < 0 && tc.want != -1 {
			t.Errorf("IsProbabPrimePseudosquare(): %s expected not-prime got %v", tc.name, got)
		}
	}
}

func TestWilliamsPP1(t *testing.T) {
	n := NewFmpz(451889)
	b1 := 10
	b2 := 50
	c := 7
	want := NewFmpz(139)

	z := new(Fmpz)
	got := z.WilliamsPP1(n, b1, b2, c)
	if got != 1 && z.Cmp(want) != 0 {
		t.Errorf("WilliamsPP1 failed - got / want mismatch: %v / %v", z, want)
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
			t.Errorf("Bits(): %s expected %v got %v", tc.name, tc.want, got)
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
			t.Errorf("TstBit(): %s expected %v got %v", tc.name, tc.want, got)
		}
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

func TestNumRows(t *testing.T) {
	for _, tc := range []struct {
		name string
		r    int
		c    int
		want int
	}{
		{
			name: "5 x 5 matrix",
			r:    5,
			c:    5,
			want: 5,
		},
	} {
		m := NewFmpzMat(tc.r, tc.c)
		got := m.NumRows()
		if got != tc.want {
			t.Errorf("NumRows() want / got mismatch: %d / %d", tc.want, got)
		}
	}
}

func TestNumCols(t *testing.T) {
	for _, tc := range []struct {
		name string
		r    int
		c    int
		want int
	}{
		{
			name: "6 x 6 matrix",
			r:    6,
			c:    6,
			want: 6,
		},
	} {
		m := NewFmpzMat(tc.r, tc.c)
		got := m.NumCols()
		if got != tc.want {
			t.Errorf("NumCols() want / got mismatch: %d / %d", tc.want, got)
		}
	}
}

func TestEntry(t *testing.T) {
	for _, tc := range []struct {
		name string
		x    int
		y    int
		want *Fmpz
	}{
		{
			name: "1 at 0,0",
			x:    0,
			y:    0,
			want: NewFmpz(1),
		},
		{
			name: "0 at 1,0",
			x:    0,
			y:    1,
			want: NewFmpz(0),
		},
	} {
		m := NewFmpzMat(4, 4)
		m = m.One()
		got := m.Entry(tc.x, tc.y)

		if got.Cmp(tc.want) != 0 {
			t.Errorf("Entry() %s want / got mismatch: %v / %v", tc.name, tc.want, got)
		}
	}
}

func TestSetPosVal(t *testing.T) {
	for _, tc := range []struct {
		name string
		pos  int
		want *Fmpz
	}{
		{
			name: "666 at 0",
			pos:  0,
			want: NewFmpz(666),
		},
		{
			name: "777 at 1",
			pos:  2,
			want: NewFmpz(777),
		},
	} {
		v := new(Fmpz).Set(tc.want)
		m := NewFmpzMat(4, 4)
		m = m.Zero()
		orig := m.Entry(tc.pos, 0) // Orig should be a zero
		m.SetPosVal(v, tc.pos)
		got := m.Entry(tc.pos, 0)

		if got.Cmp(orig) == 0 {
			t.Errorf("SetPosVal() %s failed to mutate value at pos %d - got %v / want %v", tc.name, tc.pos, got, tc.want)
		}

		if got.Cmp(tc.want) != 0 {
			t.Errorf("SetPosVal() %s want / got mismatch: %v / %v", tc.name, tc.want, got)
		}
	}
}

func TestSetVal(t *testing.T) {
	for _, tc := range []struct {
		name string
		x    int
		y    int
		want *Fmpz
	}{
		{
			name: "1 at 0,0",
			want: NewFmpz(1),
		},
		{
			name: "100 at 1,0",
			x:    1,
			want: NewFmpz(100),
		},
		{
			name: "666 at 3,2",
			x:    3,
			y:    2,
			want: NewFmpz(666),
		},
	} {
		m := new(FmpzMat)
		m.fmpzMatDoinit(4, 4)
		//m = NewFmpzMat(4, 4)
		m = m.Zero()
		orig := m.Entry(tc.x, tc.y) // Orig should be a zero
		m.SetVal(tc.want, tc.x, tc.y)
		got := m.Entry(tc.x, tc.y)

		if got.Cmp(orig) == 0 {
			t.Errorf("SetPosVal() %s failed to mutate value at pos %d,%d - got %v / want %v", tc.name, tc.x, tc.y, got, tc.want)
		}

		if got.Cmp(tc.want) != 0 {
			t.Errorf("SetPosVal() %s want / got mismatch: %v / %v", tc.name, tc.want, got)
		}
	}
}

func TestOne(t *testing.T) {
	m := new(FmpzMat)
	m.fmpzMatDoinit(3, 3)
	m = m.One()

	if m.Entry(0, 0).Cmp(NewFmpz(1)) != 0 {
		t.Errorf("TestOne: Failed setting 1 in top left")
	}
	if m.Entry(1, 1).Cmp(NewFmpz(1)) != 0 {
		t.Errorf("TestOne: Failed setting 1 in center")
	}
	if m.Entry(2, 2).Cmp(NewFmpz(1)) != 0 {
		t.Errorf("TestOne: Failed setting 1 in bottom right")
	}
}

func TestCRT(t *testing.T) {
	for _, tc := range []struct {
		name string
		r1   string
		m1   string
		r2   string
		m2   string
		sign int
		want string
	}{
		{
			name: "large int crt",
			r1:   "112820376318708309511883266356668393396816131447182791445506209031700236878469506355658352414748854472099361508824474365112325602319862842561436679067358900089331778617100580343694334226208753320435002324108477884950933641216044198203776075918323272795752182687450526442079367110656868374931125538339145721573",
			m1:   "129114230505356333697118994510021413915051088225570531043026917550451581564734952308651566723784981323900403426111056537185011193232603296112121643742691356399992969898010827710994350803494919151948423732426591598888439712920344266205641475348312727365971717305409127667391782677854219689063549759596429716629",
			r2:   "45651293556966072304818630107703140982560165499022836594523320391474750266281527820821435052904791681898782443840766880327957385288649094238886877657228597671521358830021677304123300882210216797719566693160533018601632768048713030788957904378243453859832229603157052843135978639276611231634399594108602071349",
			m2:   "109269702205029292120022054633721536134438763741801805368759852020529400112797868566931991813909053016228871499067304592740926931055426540840268677218282537450757063806695831503892336975370479004151114020279110611956433492281834217463178735931660370487895538198474855043942908224106013915984721193047940206159",
			sign: 0,
			want: "17446992834638639179129969961058029457462398677361658450137832328330435503838651797276948890990069700515669656391607670623897280684064423087023742140145529356863469816868212911716782075239982647322703714504545802436551322108638975695013439206776300941300053940942685511792851350404139366581130688518772175108412341696958930756520037",
		},
	} {

		r1, _ := new(Fmpz).SetString(tc.r1, 10)
		m1, _ := new(Fmpz).SetString(tc.m1, 10)
		r2, _ := new(Fmpz).SetString(tc.r2, 10)
		m2, _ := new(Fmpz).SetString(tc.m2, 10)
		want, _ := new(Fmpz).SetString(tc.want, 10)

		got := new(Fmpz)
		got.CRT(r1, m1, r2, m2, tc.sign)

		if got.Cmp(want) != 0 {
			t.Errorf("CRT() %s want / got mismatch: %v / %v", tc.name, want, got)
		}
	}
}
