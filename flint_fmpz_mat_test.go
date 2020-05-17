package goflint

import "testing"

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
