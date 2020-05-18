package goflint

/*
#cgo LDFLAGS: -lflint -lgmp
#include <flint/flint.h>
#include <flint/fmpz.h>
#include <flint/fmpz_mat.h>
#include <flint/fmpz_lll.h>
#include <gmp.h>
#include <stdlib.h>

// Macros

fmpz fmpzmat_get_val(fmpz_mat_t mat, slong pos) {
	return mat->entries[pos];
}

void fmpzmat_set_val(fmpz_mat_t mat, fmpz_t val, slong pos) {
	mat->entries[pos] = *val;
}

*/
import "C"

import (
	"errors"
	"runtime"
	"unsafe"
)

// FmpzMat is a matrix of Fmpz.
type FmpzMat struct {
	i    C.fmpz_mat_t
	rows int
	cols int
	init bool
}

// FmpzLLL stores a LLL matrix reduction context including delta, eta, rt and gt values.
type FmpzLLL struct {
	i    C.fmpz_lll_t
	init bool
}

// Matrices.

// fmpzMatFinalize releases the memory allocated to the FmpzMat.
func fmpzMatFinalize(m *FmpzMat) {
	if m.init {
		runtime.SetFinalizer(m, nil)
		C.fmpz_mat_clear(&m.i[0])
		m.init = false
		m.rows = 0
		m.cols = 0
	}
}

// fmpzMatDoinit initializes an FmpzMat type.
func (m *FmpzMat) fmpzMatDoinit(d ...int) error {
	if m.init {
		return nil
	}
	if len(d) == 2 {
		m.rows = d[0]
		m.cols = d[1]
		m.init = true
		C.fmpz_mat_init(&m.i[0], C.slong(m.rows), C.slong(m.cols))
		runtime.SetFinalizer(m, fmpzMatFinalize)

		return nil
	}

	return errors.New("fmpzMatDoinit: pass rows and colums on first init")
}

// fmpzLLLDoinit initializes an FmpzLLL type.
func (l *FmpzLLL) fmpzLLLDoinit() {
	if l.init {
		return
	}

	l.init = true
	C.fmpz_lll_context_init_default(&l.i[0])
}

// NewFmpzMat allocates a rows * cols matrix and returns a new FmpzMat.
func NewFmpzMat(rows, cols int) *FmpzMat {
	m := new(FmpzMat)
	m.fmpzMatDoinit(rows, cols)
	return m
}

func (m *FmpzMat) String() string {
	// Create a FILE * memstream.
	var buf *C.char
	var bufSize C.size_t
	ms := C.open_memstream(&buf, &bufSize)
	if ms == nil {
		return ""
	}
	defer func() {
		C.fclose(ms)
		C.free(unsafe.Pointer(buf))
	}()

	if pp := C.fmpz_mat_fprint_pretty(ms, &m.i[0]); pp <= 0 {
		// Positive value on success.
		return ""
	}

	if rc := C.fflush(ms); rc != 0 {
		// log.Warningf("fflush returned %d", rc)
		return ""
	}

	return C.GoString(buf)
}

// Zero sets all values of matrix m to zero and returns m.
func (m *FmpzMat) Zero() *FmpzMat {
	C.fmpz_mat_zero(&m.i[0])
	return m
}

// One sets diagonal values of matrix m to 1 and returns m.
func (m *FmpzMat) One() *FmpzMat {
	C.fmpz_mat_one(&m.i[0])
	return m
}

// NumRows returns the number of rows in a FmpzMat matrix.
func (m *FmpzMat) NumRows() int {
	return int(C.fmpz_mat_nrows(&m.i[0]))
}

// NumCols returns the number of cols in a FmpzMat matrix.
func (m *FmpzMat) NumCols() int {
	return int(C.fmpz_mat_ncols(&m.i[0]))
}

// Entry returns the value at x, y in the matrix m.
func (m *FmpzMat) Entry(x, y int) *Fmpz {
	z := new(Fmpz)
	z.doinit()
	z.i[0] = *C.fmpz_mat_entry(&m.i[0], C.slong(y), C.slong(x))
	return z
}

// SetPosVal sets position pos in matrix m to val and returns m.
func (m *FmpzMat) SetPosVal(val *Fmpz, pos int) *FmpzMat {
	val.doinit()
	C.fmpzmat_set_val(&m.i[0], &val.i[0], C.slong(pos))
	return m
}

// SetVal sets position x, y in matrix m to val and returns m.
func (m *FmpzMat) SetVal(val *Fmpz, x, y int) *FmpzMat {
	val.doinit()
	pos := x + m.cols*y
	return m.SetPosVal(val, pos)
}

// LLL reduces m in place according to the parameters specified by the default LLL context of
// fl->delta, fl->eta, fl->rt and fl->gt set to 0.99, 0.51, ZBASIS and APPROX respectively.
// u is the matrix used to capture the unimodular transformations if it is not NULL.
func (m *FmpzMat) LLL() *FmpzMat {
	l := new(FmpzLLL)
	l.fmpzLLLDoinit()
	C.fmpz_lll(&m.i[0], nil, &l.i[0])
	return m
}
