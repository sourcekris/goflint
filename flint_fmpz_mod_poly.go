package goflint

/*
#cgo LDFLAGS: -lflint
#include <flint/fmpz_mod_poly.h>

*/
import "C"

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"unsafe"
)

// FmpzModPoly type represents elements of Z/nZ[x] for a fixed modulus n.
type FmpzModPoly struct {
	i    C.fmpz_mod_poly_t
	ctx  *FmpzModCtx
	init bool
}

// fmpzModPolyFinalize releases the memory allocated to the FmpzModPoly.
func fmpzModPolyFinalize(z *FmpzModPoly) {
	if z.init {
		runtime.SetFinalizer(z, nil)
		C.fmpz_mod_poly_clear(&z.i[0], &z.ctx.i[0])
		z.init = false
	}
}

// fmpzModPolyDoinit initializes an FmpzModPoly type.
func (z *FmpzModPoly) fmpzModPolyDoinit(n *FmpzModCtx) {
	if z.init {
		return
	}
	z.init = true
	C.fmpz_mod_poly_init(&z.i[0], &n.i[0])
	runtime.SetFinalizer(z, fmpzModPolyFinalize)
}

// fmpzModPolyDoinit2 initializes an FmpzModPoly type with at least a coefficients.
func (z *FmpzModPoly) fmpzModPolyDoinit2(n *FmpzModCtx, a int) {
	if z.init {
		return
	}
	z.init = true
	C.fmpz_mod_poly_init2(&z.i[0], C.slong(a), &n.i[0])
	runtime.SetFinalizer(z, fmpzModPolyFinalize)
}

// fmpzModPolyDoinitNF initializes an FmpzModPoly type.
func (z *FmpzModPoly) fmpzModPolyDoinitNF(n *FmpzModCtx) {
	if z.init {
		return
	}
	z.init = true
	C.fmpz_mod_poly_init(&z.i[0], &n.i[0])
}

// NewFmpzModPoly allocates a new FmpzModPoly mod n and returns it.
func NewFmpzModPoly(n *FmpzModCtx) *FmpzModPoly {
	p := new(FmpzModPoly)
	p.fmpzModPolyDoinit(n)
	p.ctx = n
	return p
}

// NewFmpzModPolyNF allocates a new FmpzModPoly mod n and returns it.
func NewFmpzModPolyNF(n *FmpzModCtx) *FmpzModPoly {
	p := new(FmpzModPoly)
	p.fmpzModPolyDoinitNF(n)
	p.ctx = n
	return p
}

// NewFmpzModPoly2 allocates a new FmpzModPoly mod n with at least a coefficients and returns it.
func NewFmpzModPoly2(n *FmpzModCtx, a int) *FmpzModPoly {
	p := new(FmpzModPoly)
	p.fmpzModPolyDoinit2(n, a)
	p.ctx = n
	return p
}

// Arbitrary precision polynomials over integers mod n

// Set sets z to poly and returns z.
func (z *FmpzModPoly) Set(poly *FmpzModPoly) *FmpzModPoly {
	C.fmpz_mod_poly_set(&z.i[0], &poly.i[0], &z.ctx.i[0])
	return z
}

// SetString returns a polynomial in mod n using the string representation as the definition.
// e.g. "4 6  1 2 0 5" produces 5x3+2x+1 in (Z/6Z)[x].
func SetString(poly string) (*FmpzModPoly, error) {
	cs := strings.Split(poly, " ")
	if len(cs) < 4 {
		return nil, fmt.Errorf("not enough values in poly - expected len, mod, and at least 1 coefficient")
	}

	if cs[2] != "" {
		return nil, fmt.Errorf("invalid poly format - expected 2 spaces between modulus and coefficients")
	}

	l, err := strconv.Atoi(cs[0])
	if err != nil {
		return nil, fmt.Errorf("failed converting len to int: %v", err)
	}

	n, ok := new(Fmpz).SetString(cs[1], 10)
	if !ok {
		return nil, fmt.Errorf("failed converting modulus to fmpz: %q", cs[1])
	}

	ctx := NewFmpzModCtx(n)

	res := NewFmpzModPoly2(ctx, l)
	for i, c := range cs[3:] {
		cc, ok := new(Fmpz).SetString(c, 10)
		if !ok {
			return nil, fmt.Errorf("failed converting coefficient %d to fmpz: %q", i, cs[i])
		}

		res.SetCoeff(i, cc)
	}

	return res, nil
}

// String returns a string representation of the polynomial.
func (z *FmpzModPoly) String() string {
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
	if pp := C.fmpz_mod_poly_fprint_pretty(ms, &z.i[0], &x, &z.ctx.i[0]); pp <= 0 {
		// Positive value on success.
		return ""
	}

	if rc := C.fflush(ms); rc != 0 {
		return ""
	}

	return C.GoString(buf)
}

// StringSimple returns a simple string representation of the polynomials length, modulus and
// coefficients. e.g. f(x)=5x^3+2x+1  in (Z/6Z)[x] is "4 6  1 2 0 5"
func (z *FmpzModPoly) StringSimple() string {
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

	if pp := C.fmpz_mod_poly_fprint(ms, &z.i[0], &z.ctx.i[0]); pp <= 0 {
		// Positive value on success.
		return ""
	}

	if rc := C.fflush(ms); rc != 0 {
		return ""
	}

	return C.GoString(buf)
}

// Zero sets z to the zero polynomial and returns z.
func (z *FmpzModPoly) Zero() *FmpzModPoly {
	C.fmpz_mod_poly_zero(&z.i[0], &z.ctx.i[0])
	return z
}

// FitLength sets the number of coefficiets in z to l.
func (z *FmpzModPoly) FitLength(l int) {
	C.fmpz_mod_poly_fit_length(&z.i[0], C.slong(l), &z.ctx.i[0])
}

// SetCoeff sets the c'th coefficient of z to x where x is an Fmpz and returns z.
func (z *FmpzModPoly) SetCoeff(c int, x *Fmpz) *FmpzModPoly {
	C.fmpz_mod_poly_set_coeff_fmpz(&z.i[0], C.slong(c), &x.i[0], &z.ctx.i[0])
	return z
}

// GetMod gets the modulus of z and returns an Fmpz.
func (z *FmpzModPoly) GetMod() *Fmpz {
	return z.ctx.n
}

// Len returns the length of the poly z.
func (z *FmpzModPoly) Len() int {
	return int(C.fmpz_mod_poly_length(&z.i[0], &z.ctx.i[0]))
}

// GetCoeff gets the c'th coefficient of z and returns an Fmpz.
func (z *FmpzModPoly) GetCoeff(c int) *Fmpz {
	r := new(Fmpz)
	r.doinit()
	C.fmpz_mod_poly_get_coeff_fmpz(&r.i[0], &z.i[0], C.slong(c), &z.ctx.i[0])
	return r
}

// GetCoeffs gets all of the coefficient of z and returns a slice of Fmpz.
func (z *FmpzModPoly) GetCoeffs() []*Fmpz {
	var coefficients []*Fmpz
	for i := 0; i < z.Len(); i++ {
		r := new(Fmpz)
		r.doinit()
		C.fmpz_mod_poly_get_coeff_fmpz(&r.i[0], &z.i[0], C.slong(i), &z.ctx.i[0])
		coefficients = append(coefficients, r)
	}
	return coefficients
}

// SetCoeffUI sets the c'th coefficient of z to x where x is an uint and returns z.
func (z *FmpzModPoly) SetCoeffUI(c int, x uint) *FmpzModPoly {
	C.fmpz_mod_poly_set_coeff_ui(&z.i[0], C.slong(c), C.ulong(x), &z.ctx.i[0])
	return z
}

// Neg sets z to the negative of p and returns z.
func (z *FmpzModPoly) Neg(p *FmpzModPoly) *FmpzModPoly {
	C.fmpz_mod_poly_neg(&z.i[0], &p.i[0], &z.ctx.i[0])
	return z
}

// GCD sets z = gcd(a, b) and returns
func (z *FmpzModPoly) GCD(a, b *FmpzModPoly) *FmpzModPoly {
	C.fmpz_mod_poly_gcd(&z.i[0], &a.i[0], &b.i[0], &z.ctx.i[0])
	return z
}

// Equal returns true if z is equal to p otherwise false.
func (z *FmpzModPoly) Equal(p *FmpzModPoly) bool {
	r := int(C.fmpz_mod_poly_equal(&z.i[0], &p.i[0], &z.ctx.i[0]))
	return r != 0
}

// Add sets z = a + b and returns z.
func (z *FmpzModPoly) Add(a, b *FmpzModPoly) *FmpzModPoly {
	C.fmpz_mod_poly_add(&z.i[0], &a.i[0], &b.i[0], &z.ctx.i[0])
	return z
}

// Sub sets z = a - b and returns z.
func (z *FmpzModPoly) Sub(a, b *FmpzModPoly) *FmpzModPoly {
	C.fmpz_mod_poly_sub(&z.i[0], &a.i[0], &b.i[0], &z.ctx.i[0])
	return z
}

// Mul sets z = a * b and returns z.
func (z *FmpzModPoly) Mul(a, b *FmpzModPoly) *FmpzModPoly {
	C.fmpz_mod_poly_mul(&z.i[0], &a.i[0], &b.i[0], &z.ctx.i[0])
	return z
}

// MulScalar sets z = a * x where x is an Fmpz.
func (z *FmpzModPoly) MulScalar(a *FmpzModPoly, x *Fmpz) *FmpzModPoly {
	C.fmpz_mod_poly_scalar_mul_fmpz(&z.i[0], &a.i[0], &x.i[0], &z.ctx.i[0])
	return z
}

// DivScalar sets z = a / x where x is an Fmpz.
func (z *FmpzModPoly) DivScalar(a *FmpzModPoly, x *Fmpz) *FmpzModPoly {
	C.fmpz_mod_poly_scalar_div_fmpz(&z.i[0], &a.i[0], &x.i[0], &z.ctx.i[0])
	return z
}

// Pow sets z to m^e and returns z.
func (z *FmpzModPoly) Pow(m *FmpzModPoly, e int) *FmpzModPoly {
	C.fmpz_mod_poly_pow(&z.i[0], &m.i[0], C.ulong(e), &z.ctx.i[0])
	return z
}

// DivRem computes q, r such that z=mq+r and 0 ≤ len(r) < len(m).
func (z *FmpzModPoly) DivRem(m *FmpzModPoly) (*FmpzModPoly, *FmpzModPoly) {
	q := NewFmpzModPoly(z.ctx)
	r := NewFmpzModPoly(z.ctx)
	C.fmpz_mod_poly_divrem(&q.i[0], &r.i[0], &z.i[0], &m.i[0], &z.ctx.i[0])
	return q, r
}
