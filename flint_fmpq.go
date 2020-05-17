package goflint

/*
#cgo LDFLAGS: -lflint -lgmp
#include <flint/flint.h>
#include <flint/fmpz.h>
#include <flint/fmpq.h>
#include <gmp.h>
#include <stdlib.h>

// Macros

fmpz *_fmpq_numref(fmpq_t op) {
    return fmpq_numref(op);
}
fmpz *_fmpq_denref(fmpq_t op) {
    return fmpq_denref(op);
}

*/
import "C"
import (
	"runtime"
	"unsafe"
)

// Fmpq is an arbitrary precision rational type.
type Fmpq struct {
	i    C.fmpq_t
	init bool
}

// fmpqFinalize releases the memory allocated to the Fmpq.
func fmpqFinalize(q *Fmpq) {
	if q.init {
		runtime.SetFinalizer(q, nil)
		C.fmpq_clear(&q.i[0])
		q.init = false
	}
}

// fmpqDoinit initializes an Fmpz type.
func (q *Fmpq) fmpqDoinit() {
	if q.init {
		return
	}
	q.init = true
	C.fmpq_init(&q.i[0])
	runtime.SetFinalizer(q, fmpqFinalize)
}

// string returns a string representation of q in the base given
func (q *Fmpq) string(base int) string {
	if q == nil {
		return "<nil>"
	}
	q.fmpqDoinit()
	p := C.fmpq_get_str(nil, C.int(base), &q.i[0])
	s := C.GoString(p)
	C.free(unsafe.Pointer(p))
	return s
}

// String returns the decimal representation of z.
func (q *Fmpq) String() string {
	return q.string(10)
}

// NewFmpq allocates and returns a new Fmpq set to p / q.
func NewFmpq(p, q int64) *Fmpq {
	x := C.slong(p)
	y := C.ulong(q)
	z := new(Fmpq)
	z.fmpqDoinit()
	C.fmpq_set_si(&z.i[0], x, y)
	return z
}

// NewFmpqFmpz allocates and returns a new Fmpq set to p / q where p and q are Fmpz types.
func NewFmpqFmpz(p, q *Fmpz) *Fmpq {
	z := new(Fmpq)
	z.fmpqDoinit()
	C.fmpq_set_fmpz_frac(&z.i[0], &p.i[0], &q.i[0])
	return z
}

// SetFmpqFraction sets the value of q to the canonical form of
// the fraction num / den and returns q.
func (q *Fmpq) SetFmpqFraction(num, den *Fmpz) *Fmpq {
	q.fmpqDoinit()
	C.fmpq_set_fmpz_frac(&q.i[0], &num.i[0], &den.i[0])
	return q
}

// CmpRational compares rationals z and y and returns:
//   -1 if z <  y
//    0 if z == y
//   +1 if z >  y
func (q *Fmpq) CmpRational(y *Fmpq) (r int) {
	q.fmpqDoinit()
	y.fmpqDoinit()
	r = int(C.fmpq_cmp(&q.i[0], &y.i[0]))
	if r < 0 {
		r = -1
	} else if r > 0 {
		r = 1
	}
	return
}

// Cmp wraps CmpRational.
func (q *Fmpq) Cmp(y *Fmpq) int {
	return q.CmpRational(y)
}

// GetFmpqFraction gets the integer numerator and denomenator of the rational Fmpq q.
func (q *Fmpq) GetFmpqFraction() (int, int) {
	q.fmpqDoinit()

	// store the num and den into Mpzs
	// fmpq_get_mpz_frac is not reliably in the flint.h on different FLINT distributions.
	// C.fmpq_get_mpz_frac(&a.i[0], &b.i[0], &q.i[0])

	// store the num and den into ints.
	n := C._fmpq_numref(&q.i[0])
	d := C._fmpq_denref(&q.i[0])

	return int(*n), int(*d)
}

// NumRef returns the numerator of an Fmpq as an integer.
func (q *Fmpq) NumRef() int {
	q.fmpqDoinit()
	z := C._fmpq_numref(&q.i[0])
	return int(*z)
}

// DenRef returns the denominator of an Fmpq as an integer.
func (q *Fmpq) DenRef() int {
	q.fmpqDoinit()
	z := C._fmpq_denref(&q.i[0])
	return int(*z)
}

// MulRational sets q to the product of rational x and integer y and returns q.
func (q *Fmpq) MulRational(o *Fmpq, x *Fmpz) *Fmpq {
	x.doinit()
	o.fmpqDoinit()
	q.fmpqDoinit()
	C.fmpq_mul_fmpz(&q.i[0], &o.i[0], &x.i[0])
	return q
}
