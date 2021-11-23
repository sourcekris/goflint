package goflint

/*
#cgo LDFLAGS: -lflint -lgmp
#include <flint/flint.h>
#include <flint/fmpz.h>
#include <flint/fmpz_poly.h>
#include <flint/fmpz_poly_factor.h>
#include <gmp.h>
#include <stdlib.h>

// Macros
fmpz_poly_struct fmpz_poly_factor_get_poly(fmpz_poly_factor_t fac, slong i) {
	return fac->p[i];
}

fmpz fmpz_poly_factor_get_coeff(fmpz_poly_factor_t fac) {
	return fac->c;
}

slong fmpz_poly_factor_get_num(fmpz_poly_factor_t fac) {
	return fac->num;
}

slong fmpz_poly_factor_get_exp(fmpz_poly_factor_t fac, slong i) {
	return fac->exp[i];
}

*/
import "C"

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"unsafe"
)

// FmpzPoly type represents a univariate polynomial over the integers.
type FmpzPoly struct {
	i    C.fmpz_poly_t
	init bool
}

// FmpzPolyFactor type represents the factors univariate polynomial over the integers.
type FmpzPolyFactor struct {
	i    C.fmpz_poly_factor_t
	init bool
}

// fmpzPolyFinalize releases the memory allocated to the FmpzPoly.
func fmpzPolyFinalize(z *FmpzPoly) {
	if z.init {
		runtime.SetFinalizer(z, nil)
		C.fmpz_poly_clear(&z.i[0])
		z.init = false
	}
}

// fmpzPolyFactorFinalize releases the memory allocated to the FmpzPoly.
func fmpzPolyFactorFinalize(f *FmpzPolyFactor) {
	if f.init {
		runtime.SetFinalizer(f, nil)
		//C.fmpz_poly_factor_clear(&f.i[0])
		f.init = false
	}
}

// fmpzPolyDoinit initializes an FmpzPoly type.
func (z *FmpzPoly) fmpzPolyDoinit() {
	if z.init {
		return
	}
	z.init = true
	C.fmpz_poly_init(&z.i[0])
	runtime.SetFinalizer(z, fmpzPolyFinalize)
}

// fmpzPolyDoinit2 initializes an FmpzPoly type with at least a coefficients.
func (z *FmpzPoly) fmpzPolyDoinit2(a int) {
	if z.init {
		return
	}
	z.init = true
	C.fmpz_poly_init2(&z.i[0], C.slong(a))
	runtime.SetFinalizer(z, fmpzPolyFinalize)
}

// fmpzPolyFactor initializes an FmpzPolyFactor type.
func (f *FmpzPolyFactor) fmpzPolyFactorDoinit() {
	if f.init {
		return
	}
	f.init = true
	C.fmpz_poly_factor_init(&f.i[0])
	runtime.SetFinalizer(f, fmpzPolyFactorFinalize)
}

// NewFmpzPoly allocates a new FmpzPoly and returns it.
func NewFmpzPoly() *FmpzPoly {
	p := new(FmpzPoly)
	p.fmpzPolyDoinit()
	return p
}

// NewFmpzPoly2 allocates a new FmpzPoly with at least a coefficients and returns it.
func NewFmpzPoly2(a int) *FmpzPoly {
	p := new(FmpzPoly)
	p.fmpzPolyDoinit2(a)
	return p
}

// NewFmpzPolyFactor allocates a new FmpzPolyFactor and returns it.
func NewFmpzPolyFactor() *FmpzPolyFactor {
	f := new(FmpzPolyFactor)
	f.fmpzPolyFactorDoinit()
	return f
}

// Arbitrary precision polynomials over integers.

// Set sets z to poly and returns z.
func (z *FmpzPoly) Set(poly *FmpzPoly) *FmpzPoly {
	C.fmpz_poly_set(&z.i[0], &poly.i[0])
	return z
}

// Set sets f to FmpzPolyFactor fac and returns f.
func (f *FmpzPolyFactor) Set(fac *FmpzPolyFactor) *FmpzPolyFactor {
	C.fmpz_poly_factor_set(&f.i[0], &fac.i[0])
	return f
}

// FmpzPolySetString returns a polynomial using the string representation as the definition.
// e.g. "4  1 2 0 5" produces 5x3+2x+1.
func FmpzPolySetString(poly string) (*FmpzPoly, error) {
	cs := strings.Split(poly, " ")
	if len(cs) < 3 {
		return nil, fmt.Errorf("not enough values in poly - expected len, and at least 1 coefficient")
	}

	if cs[1] != "" {
		return nil, fmt.Errorf("invalid poly format - expected 2 spaces between len and coefficients")
	}

	l, err := strconv.Atoi(cs[0])
	if err != nil {
		return nil, fmt.Errorf("failed converting len to int: %v", err)
	}

	res := NewFmpzPoly2(l)
	for i, c := range cs[2:] {
		cc, ok := new(Fmpz).SetString(c, 10)
		if !ok {
			return nil, fmt.Errorf("failed converting coefficient %d to fmpz: %q", i, cs[i])
		}

		res.SetCoeff(i, cc)
	}

	return res, nil
}

// String returns a string representation of the polynomial.
func (z *FmpzPoly) String() string {
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

	var x C.char = 'x'
	if pp := C.fmpz_poly_fprint_pretty(ms, &z.i[0], &x); pp <= 0 {
		// Positive value on success.
		return ""
	}

	if rc := C.fflush(ms); rc != 0 {
		return ""
	}

	return C.GoString(buf)
}

// StringSimple returns a simple string representation of the polynomials length and
// coefficients. e.g. f(x)=5x^3+2x+1  is "4  1 2 0 5"
func (z *FmpzPoly) StringSimple() string {
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

	if pp := C.fmpz_poly_fprint(ms, &z.i[0]); pp <= 0 {
		// Positive value on success.
		return ""
	}

	if rc := C.fflush(ms); rc != 0 {
		return ""
	}

	return C.GoString(buf)
}

// Print prints the FmpzPolyFactor to stdout.
func (f *FmpzPolyFactor) Print() {
	C.fmpz_poly_factor_print(&f.i[0])
}

// Zero sets z to the zero polynomial and returns z.
func (z *FmpzPoly) Zero() *FmpzPoly {
	C.fmpz_poly_zero(&z.i[0])
	return z
}

// FitLength sets the number of coefficiets in z to l.
func (z *FmpzPoly) FitLength(l int) {
	C.fmpz_poly_fit_length(&z.i[0], C.slong(l))
}

// SetCoeff sets the c'th coefficient of z to x where x is an Fmpz and returns z.
func (z *FmpzPoly) SetCoeff(c int, x *Fmpz) *FmpzPoly {
	C.fmpz_poly_set_coeff_fmpz(&z.i[0], C.slong(c), &x.i[0])
	return z
}

// Len returns the length of the poly z.
func (z *FmpzPoly) Len() int {
	return int(C.fmpz_poly_length(&z.i[0]))
}

// GetCoeff gets the c'th coefficient of z and returns an Fmpz.
func (z *FmpzPoly) GetCoeff(c int) *Fmpz {
	r := new(Fmpz)
	r.doinit()
	C.fmpz_poly_get_coeff_fmpz(&r.i[0], &z.i[0], C.slong(c))
	return r
}

// GetCoeffs gets all of the coefficient of z and returns a slice of Fmpz.
func (z *FmpzPoly) GetCoeffs() []*Fmpz {
	var coefficients []*Fmpz
	for i := 0; i < z.Len(); i++ {
		r := new(Fmpz)
		r.doinit()
		C.fmpz_poly_get_coeff_fmpz(&r.i[0], &z.i[0], C.slong(i))
		coefficients = append(coefficients, r)
	}
	return coefficients
}

// SetCoeffUI sets the c'th coefficient of z to x where x is an uint and returns z.
func (z *FmpzPoly) SetCoeffUI(c int, x uint) *FmpzPoly {
	z.fmpzPolyDoinit()
	C.fmpz_poly_set_coeff_ui(&z.i[0], C.slong(c), C.ulong(x))
	return z
}

// Neg sets z to the negative of p and returns z.
func (z *FmpzPoly) Neg(p *FmpzPoly) *FmpzPoly {
	z.fmpzPolyDoinit()
	C.fmpz_poly_neg(&z.i[0], &p.i[0])
	return z
}

// GCD sets z = gcd(a, b) and returns
func (z *FmpzPoly) GCD(a, b *FmpzPoly) *FmpzPoly {
	z.fmpzPolyDoinit()
	C.fmpz_poly_gcd(&z.i[0], &a.i[0], &b.i[0])
	return z
}

// Equal returns true if z is equal to p otherwise false.
func (z *FmpzPoly) Equal(p *FmpzPoly) bool {
	r := int(C.fmpz_poly_equal(&z.i[0], &p.i[0]))
	return r != 0
}

// Add sets z = a + b and returns z.
func (z *FmpzPoly) Add(a, b *FmpzPoly) *FmpzPoly {
	C.fmpz_poly_add(&z.i[0], &a.i[0], &b.i[0])
	return z
}

// Sub sets z = a - b and returns z.
func (z *FmpzPoly) Sub(a, b *FmpzPoly) *FmpzPoly {
	C.fmpz_poly_sub(&z.i[0], &a.i[0], &b.i[0])
	return z
}

// Mul sets z = a * b and returns z.
func (z *FmpzPoly) Mul(a, b *FmpzPoly) *FmpzPoly {
	C.fmpz_poly_mul(&z.i[0], &a.i[0], &b.i[0])
	return z
}

// MulScalar sets z = a * x where x is an Fmpz.
func (z *FmpzPoly) MulScalar(a *FmpzPoly, x *Fmpz) *FmpzPoly {
	C.fmpz_poly_scalar_mul_fmpz(&z.i[0], &a.i[0], &x.i[0])
	return z
}

// DivScalar sets z = a / x where x is an Fmpz. Rounding coefficients down toward -infinity.
func (z *FmpzPoly) DivScalar(a *FmpzPoly, x *Fmpz) *FmpzPoly {
	C.fmpz_poly_scalar_fdiv_fmpz(&z.i[0], &a.i[0], &x.i[0])
	return z
}

// Pow sets z to m^e and returns z.
func (z *FmpzPoly) Pow(m *FmpzPoly, e int) *FmpzPoly {
	C.fmpz_poly_pow(&z.i[0], &m.i[0], C.ulong(e))
	return z
}

// DivRem computes q, r such that z=mq+r and 0 ≤ len(r) < len(m).
func (z *FmpzPoly) DivRem(m *FmpzPoly) (*FmpzPoly, *FmpzPoly) {
	q := NewFmpzPoly()
	r := NewFmpzPoly()
	C.fmpz_poly_divrem(&q.i[0], &r.i[0], &z.i[0], &m.i[0])
	return q, r
}

// Factor uses the Zassenhaus factoring algorithm, which takes as input any FmpzPoly z, and
// returns a factorization in an FmpzPolyFac type.
func (z *FmpzPoly) Factor() *FmpzPolyFactor {
	fac := NewFmpzPolyFactor()
	// In FLINT 2.5.3 this should be just fmpz_poly_factor but most Linux distros are still on
	// FLINT 2.5.2.
	C.fmpz_poly_factor_zassenhaus(&fac.i[0], &z.i[0])
	return fac
}

// GetPoly gets the nth polynomial factor from a FmpzPolyFactor and returns it.
func (f *FmpzPolyFactor) GetPoly(n int) *FmpzPoly {
	p := NewFmpzPoly()
	p.i[0] = C.fmpz_poly_factor_get_poly(&f.i[0], C.slong(n))
	return p
}

// GetExp gets the exponent of the nth polynomial from the FmpzPolyFactor.
func (f *FmpzPolyFactor) GetExp(n int) int {
	return int(C.fmpz_poly_factor_get_exp(&f.i[0], C.slong(n)))
}

// GetCoeff gets the coefficient from the FmpzPolyFactor.
func (f *FmpzPolyFactor) GetCoeff() *Fmpz {
	z := new(Fmpz)
	z.doinit()
	z.i[0] = C.fmpz_poly_factor_get_coeff(&f.i[0])
	return z
}

// Len gets the length of the FmpzPolyFactors list. i.e. the number of factors found.
func (f *FmpzPolyFactor) Len() int {
	return int(C.fmpz_poly_factor_get_num(&f.i[0]))
}
