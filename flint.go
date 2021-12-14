// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file wraps some FLINT (Fast Library for Number Theory) functions

package goflint

/*
#cgo darwin CPPFLAGS: -I/opt/homebrew/include
#cgo darwin LDFLAGS: /opt/homebrew/lib/libflint.a /opt/homebrew/lib/libgmp.a /opt/homebrew/lib/libmpfr.a
#cgo linux LDFLAGS: -lflint -lgmp
#include <flint/arith.h>
#include <flint/flint.h>
#include <flint/fmpz.h>
#include <flint/fmpq.h>
#include <flint/nmod_poly.h>
#include <gmp.h>
#include <stdlib.h>

// Macros

*/
import "C"

import (
	"math/big"
	"runtime"
	"unsafe"
)

var (
	// Zero is an Fmpz of value 0.
	Zero = NewFmpz(0)
)

/*
 * Types
 */

// Fmpz is a arbitrary size integer type.
type Fmpz struct {
	i    C.fmpz_t
	init bool
}

// Mpz is an abitrary size integer type from the Gnu Multiprecision Library.
type Mpz struct {
	i    C.mpz_t
	init bool
}

// NmodPoly type represents elements of Z/nZ[x] for a fixed modulus n.
type NmodPoly struct {
	i    C.nmod_poly_t
	init bool
}

// MpLimb type is a mp_limb_t which is a type alias for ulong which in go is a uint64.
type MpLimb struct {
	i C.mp_limb_t
}

// FlintRandT keeps state for Fmpz random number generation.
type FlintRandT struct {
	i    C.flint_rand_t
	init bool
}

/*
 * Initializers and Finalizers
 */

// fmpzFinalize releases the memory allocated to the Fmpz.
func fmpzFinalize(z *Fmpz) {
	if z.init {
		runtime.SetFinalizer(z, nil)
		C.fmpz_clear(&z.i[0])
		z.init = false
	}
}

// mpzFinalize releases the memory allocated to the Mpz.
func mpzFinalize(z *Mpz) {
	if z.init {
		runtime.SetFinalizer(z, nil)
		C.mpz_clear(&z.i[0])
		z.init = false
	}
}

// nmodPolyFinalize releases the memory allocated to the NmodPoly.
func nmodPolyFinalize(z *NmodPoly) {
	if z.init {
		runtime.SetFinalizer(z, nil)
		C.nmod_poly_clear(&z.i[0])
		z.init = false
	}
}

func flintRandTFinalize(r *FlintRandT) {
	if r.init {
		runtime.SetFinalizer(r, nil)
		C.flint_randclear(&r.i[0])
		r.init = false
	}
}

// doinit initializes an Fmpz type.
func (z *Fmpz) doinit() {
	if z.init {
		return
	}
	z.init = true
	C.fmpz_init(&z.i[0])
	runtime.SetFinalizer(z, fmpzFinalize)
}

// doinitNF initializes an Fmpz type without setting a finalizer.
// This is a hack to help debug double frees in Flint.
func (z *Fmpz) doinitNF() {
	if z.init {
		return
	}
	z.init = true
	C.fmpz_init(&z.i[0])
}

// mpzDoinit initializes an Mpz type.
func (z *Mpz) mpzDoinit() {
	if z.init {
		return
	}
	z.init = true
	C.mpz_init(&z.i[0])
	runtime.SetFinalizer(z, mpzFinalize)
}

// nmodPolyDoinit initializes an NmodPoly type.
func (z *NmodPoly) nmodPolyDoinit(n *MpLimb) {
	if z.init {
		return
	}
	z.init = true
	C.nmod_poly_init(&z.i[0], n.i)
	runtime.SetFinalizer(z, nmodPolyFinalize)
}

func (r *FlintRandT) flintRandTDoinit() {
	if r.init {
		return
	}
	r.init = true
	C.flint_randinit(&r.i[0])
	runtime.SetFinalizer(r, flintRandTFinalize)
}

/*
 * Assignments
 */

// SetUint64 sets z to x and returns z.
func (z *Fmpz) SetUint64(x uint64) *Fmpz {
	z.doinit()
	y := C.ulong(x)
	C.fmpz_set_ui(&z.i[0], y)
	return z
}

// SetInt64 sets z to x and returns z.
func (z *Fmpz) SetInt64(x int64) *Fmpz {
	z.doinit()
	y := C.slong(x)
	C.fmpz_set_si(&z.i[0], y)
	return z
}

// SetInt64NF sets z to x and returns z without setting a Finalizer.
// This is a hack to help debug double frees in Flint.
func (z *Fmpz) SetInt64NF(x int64) *Fmpz {
	z.doinitNF()
	y := C.slong(x)
	C.fmpz_set_si(&z.i[0], y)
	return z
}

// SetMpzInt64 sets z to x and returns z.
func (z *Mpz) SetMpzInt64(x int64) *Mpz {
	z.mpzDoinit()
	y := C.long(x)
	C.mpz_set_si(&z.i[0], y)
	return z
}

// NewFmpz allocates and returns a new Fmpz set to x.
func NewFmpz(x int64) *Fmpz {
	return new(Fmpz).SetInt64(x)
}

// NewFmpzNF allocates and returns a new Fmpz set to x without setting a finalizer.
func NewFmpzNF(x int64) *Fmpz {
	return new(Fmpz).SetInt64NF(x)
}

// NewMpz allocates and returns a new Fmpz set to x.
func NewMpz(x int64) *Mpz {
	return new(Mpz).SetMpzInt64(x)
}

// NewMpLimb returns a new MpLimb type from a uint64.
func NewMpLimb(x uint64) *MpLimb {
	return &MpLimb{C.mp_limb_t(x)}
}

// Set sets z to x and returns z.
func (z *Fmpz) Set(x *Fmpz) *Fmpz {
	z.doinit()
	C.fmpz_set(&z.i[0], &x.i[0])
	return z
}

/*
 * Comparisons
 */

// Cmp compares z and y and returns:
//   -1 if z <  y
//    0 if z == y
//   +1 if z >  y
func (z *Fmpz) Cmp(y *Fmpz) (r int) {
	z.doinit()
	y.doinit()
	r = int(C.fmpz_cmp(&z.i[0], &y.i[0]))
	if r < 0 {
		r = -1
	} else if r > 0 {
		r = 1
	}
	return
}

// Cmp compares Mpz z and y and returns:
//   -1 if z <  y
//    0 if z == y
//   +1 if z >  y
func (z *Mpz) Cmp(y *Mpz) (r int) {
	z.mpzDoinit()
	y.mpzDoinit()
	r = int(C.mpz_cmp(&z.i[0], &y.i[0]))
	if r < 0 {
		r = -1
	} else if r > 0 {
		r = 1
	}
	return
}

// Equals compares z and y and returns true if they are equal.
func (z *Fmpz) Equals(y *Fmpz) bool {
	z.doinit()
	y.doinit()
	return int(C.fmpz_cmp(&z.i[0], &y.i[0])) == 0
}

// IsZero returns true if z == 0.
func (z *Fmpz) IsZero() bool {
	z.doinit()
	r := int(C.fmpz_cmp(&z.i[0], &Zero.i[0]))
	return r == 0
}

/*
 * Formatting
 */

// string returns z in the base given
func (z *Fmpz) string(base int) string {
	if z == nil {
		return "<nil>"
	}
	z.doinit()
	p := C.fmpz_get_str(nil, C.int(base), &z.i[0])
	s := C.GoString(p)
	C.free(unsafe.Pointer(p))
	return s
}

// string returns z in the base given
func (z *Mpz) string(base int) string {
	if z == nil {
		return "<nil>"
	}
	z.mpzDoinit()
	p := C.mpz_get_str(nil, C.int(base), &z.i[0])
	s := C.GoString(p)
	C.free(unsafe.Pointer(p))
	return s
}

// String returns the decimal representation of z.
func (z *Mpz) String() string {
	return z.string(10)
}

// String returns the decimal representation of z.
func (z *Fmpz) String() string {
	return z.string(10)
}

/*
 * Helpers
 */

// BitLen returns the length of the absolute value of z in bits.
// The bit length of 0 is 0.
func (z *Fmpz) BitLen() int {
	z.doinit()
	if z.Sign() == 0 {
		return 0
	}
	return int(C.fmpz_sizeinbase(&z.i[0], 2))
}

// Lsh left shifts an Fmpz z by an arbitrary number of bits and returns it.
func (z *Fmpz) Lsh(bits int) *Fmpz {
	z.doinit()
	// We proxy this through the std library big.Int type since Flint has no
	// method for bit shifting.
	i := new(big.Int).SetBytes(z.Bytes())
	i = i.Lsh(i, uint(bits))
	return z.SetBytes(i.Bytes())
}

// Rsh right shifts an Fmpz z by an arbitrary number of bits and returns it.
func (z *Fmpz) Rsh(bits int) *Fmpz {
	z.doinit()
	// We proxy this through the std library big.Int type since Flint has no
	// method for bit shifting.
	i := new(big.Int).SetBytes(z.Bytes())
	i = i.Rsh(i, uint(bits))
	return z.SetBytes(i.Bytes())
}

// Xor sets z to the bitwise exclusive or of a and b and returns z.
func (z *Fmpz) Xor(a, b *Fmpz) *Fmpz {
	a.doinit()
	b.doinit()
	z.doinit()
	C.fmpz_xor(&z.i[0], &a.i[0], &b.i[0])
	return z
}

// Sign returns:
//
//  -1 if x <  0
//   0 if x == 0
//  +1 if x >  0
//
func (z *Fmpz) Sign() int {
	z.doinit()
	return int(C.fmpz_sgn(&z.i[0]))
}

/*
 * Conversion
 */

// GetInt returns the value of the Fmpz type as an int type if possible.
func (z *Fmpz) GetInt() int {
	z.doinit()
	return int(C.fmpz_get_si(&z.i[0]))
}

// GetUInt returns the value of the Fmpz type as a uint type if possible.
func (z *Fmpz) GetUInt() uint {
	z.doinit()
	return uint(C.fmpz_get_ui(&z.i[0]))
}

// Int64 returns the int64 representation of z.
// If z cannot be represented in an int64, the result is undefined.
func (z *Fmpz) Int64() (y int64) {
	if !z.init {
		return
	}
	if C.fmpz_fits_si(&z.i[0]) != 0 {
		return int64(C.fmpz_get_si(&z.i[0]))
	}
	// Undefined result if > 64 bits
	if z.BitLen() > 64 {
		return
	}
	return
}

// Uint64 returns the uint64 representation of z.
// If z cannot be represented in a uint64, the result is undefined.
func (z *Fmpz) Uint64() (y uint64) {
	if !z.init {
		return
	}
	if z.BitLen() <= 64 {
		return uint64(C.fmpz_get_ui(&z.i[0]))
	}

	return
}

// SetString sets z to the value of s, interpreted in the given base,
// and returns z and a boolean indicating success. If SetString fails,
// the value of z is undefined but the returned value is nil.
//
// The base argument must be 0 or a value from 2 through MaxBase. If the base
// is 0, the string prefix determines the actual conversion base. A prefix of
// ``0x'' or ``0X'' selects base 16; the ``0'' prefix selects base 8, and a
// ``0b'' or ``0B'' prefix selects base 2. Otherwise the selected base is 10.
//
func (z *Fmpz) SetString(s string, base int) (*Fmpz, bool) {
	z.doinit()
	if base != 0 && (base < 2 || base > 36) {
		return nil, false
	}
	// Skip leading + as mpz_set_str doesn't understand them
	if len(s) > 1 && s[0] == '+' {
		s = s[1:]
	}
	// mpz_set_str incorrectly parses "0x" and "0b" as valid
	if base == 0 && len(s) == 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X' || s[1] == 'b' || s[1] == 'B') {
		return nil, false
	}
	p := C.CString(s)
	defer C.free(unsafe.Pointer(p))
	if C.fmpz_set_str(&z.i[0], p, C.int(base)) < 0 {
		return nil, false
	}
	return z, true // err == io.EOF => scan consumed all of s
}

// SetMpz transform x into an Fmpz z.
func (z *Fmpz) SetMpz(x *Mpz) {
	x.mpzDoinit()
	z.doinit()

	C.fmpz_set_mpz(&z.i[0], &x.i[0])
}

// GetMpz transform x into an Mpz z.
func (z *Mpz) GetMpz(x *Fmpz) {
	z.mpzDoinit()
	x.doinit()

	C.fmpz_get_mpz(&z.i[0], &x.i[0])
}

// SetBytes interprets buf as the bytes of a big-endian unsigned
// integer, sets z to that value, and returns z.
func (z *Fmpz) SetBytes(buf []byte) *Fmpz {
	zm := new(Mpz)
	zm.mpzDoinit()
	if len(buf) == 0 {
		z.SetInt64(0)
	} else {
		C.mpz_import(&zm.i[0], C.size_t(len(buf)), 1, 1, 1, 0, unsafe.Pointer(&buf[0]))
	}

	z.SetMpz(zm)

	return z
}

// Bytes returns the absolute value of z as a big-endian byte slice.
func (z *Fmpz) Bytes() []byte {
	zm := new(Mpz)
	zm.GetMpz(z)
	b := make([]byte, 1+(z.BitLen()+7)/8)
	n := C.size_t(len(b))
	C.mpz_export(unsafe.Pointer(&b[0]), &n, 1, 1, 1, 0, &zm.i[0])
	return b[0:n]
}

/*
 * Arithmetic
 */

// Abs sets z to |x| (the absolute value of x) and returns z.
func (z *Fmpz) Abs(x *Fmpz) *Fmpz {
	x.doinit()
	z.doinit()
	C.fmpz_abs(&z.i[0], &x.i[0])
	return z
}

// Neg sets z to -x and returns z.
func (z *Fmpz) Neg(x *Fmpz) *Fmpz {
	x.doinit()
	z.doinit()
	C.fmpz_neg(&z.i[0], &x.i[0])
	return z
}

// Add sets z to the sum x+y and returns z.
func (z *Fmpz) Add(x, y *Fmpz) *Fmpz {
	x.doinit()
	y.doinit()
	z.doinit()
	C.fmpz_add(&z.i[0], &x.i[0], &y.i[0])
	return z
}

// AddZ sets z to the sum x+z and returns z.
func (z *Fmpz) AddZ(x *Fmpz) *Fmpz {
	x.doinit()
	z.doinit()
	C.fmpz_add(&z.i[0], &x.i[0], &z.i[0])
	return z
}

// AddI sets z to the sum x+z where x is an int type and returns z.
func (z *Fmpz) AddI(x int) *Fmpz {
	fmpZX := NewFmpz(int64(x))
	z.doinit()
	C.fmpz_add(&z.i[0], &fmpZX.i[0], &z.i[0])
	return z
}

// Sub sets z to the difference x-y and returns z.
func (z *Fmpz) Sub(x, y *Fmpz) *Fmpz {
	x.doinit()
	y.doinit()
	z.doinit()
	C.fmpz_sub(&z.i[0], &x.i[0], &y.i[0])
	return z
}

// SubZ sets z to the difference z-x and returns z.
func (z *Fmpz) SubZ(x *Fmpz) *Fmpz {
	x.doinit()
	z.doinit()
	C.fmpz_sub(&z.i[0], &z.i[0], &x.i[0])
	return z
}

// SubI sets z to the difference z-x where x is an int type and returns z.
func (z *Fmpz) SubI(x int) *Fmpz {
	fmpZX := NewFmpz(int64(x))
	z.doinit()
	C.fmpz_sub(&z.i[0], &z.i[0], &fmpZX.i[0])
	return z
}

// Mul sets z to the product x*y and returns z.
func (z *Fmpz) Mul(x, y *Fmpz) *Fmpz {
	x.doinit()
	y.doinit()
	z.doinit()
	C.fmpz_mul(&z.i[0], &x.i[0], &y.i[0])
	return z
}

// MulZ sets z to the product of z  * x and returns z.
func (z *Fmpz) MulZ(x *Fmpz) *Fmpz {
	x.doinit()
	z.doinit()
	C.fmpz_mul(&z.i[0], &x.i[0], &z.i[0])
	return z
}

// MulI sets z to the product of z  * x where x is an int type
// and returns z.
func (z *Fmpz) MulI(x int) *Fmpz {
	fmpZX := NewFmpz(int64(x))
	z.doinit()
	C.fmpz_mul(&z.i[0], &fmpZX.i[0], &z.i[0])
	return z
}

// MulRMpz sets z to the product of z and y modulo n and returns z using Mpz
// types.
func (z *Mpz) MulRMpz(y, n *Mpz) *Mpz {
	z.mpzDoinit()
	y.mpzDoinit()
	C.mpz_mul(&z.i[0], &z.i[0], &y.i[0])
	C.mpz_fdiv_r(&z.i[0], &z.i[0], &n.i[0])
	return z
}

// DivR sets z to the result of z/y in the ring of integers(n). Currently this
// only works if y fits into the int type supported by the Fmpq type.
func (z *Fmpz) DivR(y, n *Fmpz) *Fmpz {
	z.doinit()
	y.doinit()
	n.doinit()

	// We'll do division of a/b as a * 1/b.

	// Get a denominator.
	d := int64(y.GetInt())
	if d == 0 {
		return NewFmpz(0)
	}

	rat := NewFmpq(1, d)
	rat.MulRational(rat, z)

	res := z.ModRational(rat, n)
	if res == 0 {
		// No residue exists.
		return NewFmpz(0)
	}

	return z
}

// SubRMpz sets z to the z -y modulo n and returns z using Mpz
// types.
func (z *Mpz) SubRMpz(y, n *Mpz) *Mpz {
	z.mpzDoinit()
	y.mpzDoinit()
	C.mpz_sub(&z.i[0], &z.i[0], &y.i[0])
	if z.Cmp(NewMpz(0)) == -1 {
		C.mpz_add(&z.i[0], &z.i[0], &n.i[0])
	}
	return z
}

// Quo sets z to the quotient x/y for y != 0 and returns z.
// If y == 0, a division-by-zero run-time panic occurs.
// Quo implements truncated division (like Go); see QuoRem for more details.
func (z *Fmpz) Quo(x, y *Fmpz) *Fmpz {
	x.doinit()
	y.doinit()
	z.doinit()
	C.fmpz_tdiv_q(&z.i[0], &x.i[0], &y.i[0])
	return z
}

// QuoRem sets z to the quotient x/y and r to the remainder x%y
// and returns the pair (z, r) for y != 0.
// If y == 0, a division-by-zero run-time panic occurs.
//
// QuoRem implements T-division and modulus (like Go):
//
//  q = x/y      with the result truncated to zero
//  r = x - y*q
//
// (See Daan Leijen, ``Division and Modulus for Computer Scientists''.)
// See DivMod for Euclidean division and modulus (unlike Go).
//
func (z *Fmpz) QuoRem(x, y, r *Fmpz) (*Fmpz, *Fmpz) {
	x.doinit()
	y.doinit()
	r.doinit()
	z.doinit()
	C.fmpz_tdiv_qr(&z.i[0], &r.i[0], &x.i[0], &y.i[0])
	return z, r
}

// Div sets z to the quotient x/y for y != 0 and returns z.
// If y == 0, a division-by-zero run-time panic occurs.
// Div implements Euclidean division (unlike Go); see DivMod for more details.
func (z *Fmpz) Div(x, y *Fmpz) *Fmpz {
	x.doinit()
	y.doinit()
	z.doinit()
	switch y.Sign() {
	case 1:
		C.fmpz_fdiv_q(&z.i[0], &x.i[0], &y.i[0])
	case -1:
		C.fmpz_cdiv_q(&z.i[0], &x.i[0], &y.i[0])
	case 0:
		panic("Division by zero")
	}
	return z
}

// DivMod sets z to the quotient x div y and m to the modulus x mod y
// and returns the pair (z, m) for y != 0.
func (z *Mpz) DivMod(x, y, m *Mpz) (*Mpz, *Mpz) {
	x.mpzDoinit()
	y.mpzDoinit()
	m.mpzDoinit()
	z.mpzDoinit()
	switch y.Cmp(NewMpz(0)) {
	case 1:
		C.mpz_fdiv_qr(&z.i[0], &m.i[0], &x.i[0], &y.i[0])
	case -1:
		C.mpz_cdiv_qr(&z.i[0], &m.i[0], &x.i[0], &y.i[0])
	case 0:
		panic("Division by zero")
	}
	return z, m
}

/*
 * Modular Arithmatic
 */

// Mod sets z to the modulus x%y for y != 0 and returns z.
func (z *Fmpz) Mod(x, y *Fmpz) *Fmpz {
	x.doinit()
	y.doinit()
	z.doinit()

	C.fmpz_mod(&z.i[0], &x.i[0], &y.i[0])
	return z
}

// ModZ sets z to the modulus z%y for y != 0 and returns z.
func (z *Fmpz) ModZ(y *Fmpz) *Fmpz {
	y.doinit()
	z.doinit()

	C.fmpz_mod(&z.i[0], &z.i[0], &y.i[0])
	return z
}

// ModRational sets z to the residue of x = n/d (num, den) modulo n and
// returns 1 if such a residue exists otherwise 0.
func (z *Fmpz) ModRational(x *Fmpq, n *Fmpz) int {
	z.doinit()
	n.doinit()
	x.fmpqDoinit()
	return int(C.fmpq_mod_fmpz(&z.i[0], &x.i[0], &n.i[0]))
}

// DivMod sets z to the quotient x div y and m to the modulus x mod y
// and returns the pair (z, m) for y != 0.
// If y == 0, a division-by-zero run-time panic occurs.
//
// DivMod implements Euclidean division and modulus (unlike Go):
//
//  q = x div y  such that
//  m = x - y*q  with 0 <= m < |q|
//
// (See Raymond T. Boute, ``The Euclidean definition of the functions
// div and mod''. ACM Transactions on Programming Languages and
// Systems (TOPLAS), 14(2):127-144, New York, NY, USA, 4/1992.
// ACM press.)
// See QuoRem for T-division and modulus (like Go).
//
func (z *Fmpz) DivMod(x, y, m *Fmpz) (*Fmpz, *Fmpz) {
	x.doinit()
	y.doinit()
	m.doinit()
	z.doinit()
	switch y.Sign() {
	case 1:
		C.fmpz_fdiv_qr(&z.i[0], &m.i[0], &x.i[0], &y.i[0])
	case -1:
		xm := new(Mpz)
		ym := new(Mpz)
		mm := new(Mpz)
		zm := new(Mpz)

		xm.GetMpz(x)
		ym.GetMpz(y)
		mm.GetMpz(m)
		zm.GetMpz(z)

		xm.mpzDoinit()
		ym.mpzDoinit()
		mm.mpzDoinit()
		zm.mpzDoinit()

		C.mpz_cdiv_qr(&zm.i[0], &mm.i[0], &xm.i[0], &ym.i[0])

		z.SetMpz(zm)
		m.SetMpz(mm)
	case 0:
		panic("Division by zero")
	}
	return z, m
}

// ModInverse sets z to the inverse of x modulo y and returns z.
// The value of y may not be 0 otherwise an exception results. If the
// inverse exists the return value will be non-zero, otherwise the return value
// will be 0 and the value of f undefined.
func (z *Fmpz) ModInverse(x, y *Fmpz) *Fmpz {
	x.doinit()
	y.doinit()
	z.doinit()

	C.fmpz_invmod(&z.i[0], &x.i[0], &y.i[0])
	return z
}

// NegMod Sets z to −x (mod y), assuming x is reduced modulo y.
func (z *Fmpz) NegMod(x, y *Fmpz) *Fmpz {
	x.doinit()
	y.doinit()
	z.doinit()

	C.fmpz_negmod(&z.i[0], &x.i[0], &y.i[0])
	return z
}

// Jacobi computes the Jacobi symbol of a modulo p, where p is a prime and a is reduced modulo p
func (z *Fmpz) Jacobi(p *Fmpz) int {
	z.doinit()
	p.doinit()

	return int(C.fmpz_jacobi(&z.i[0], &p.i[0]))
}

// Exp sets z = x**y mod |m| (i.e. the sign of m is ignored), and returns z.
// If y <= 0, the result is 1; if m == nil or m == 0, z = x**y.
// See Knuth, volume 2, section 4.6.3.
func (z *Fmpz) Exp(x, y, m *Fmpz) *Fmpz {
	x.doinit()
	y.doinit()
	z.doinit()
	if y.Sign() <= 0 {
		return z.SetInt64(1)
	}
	if m == nil || m.Sign() == 0 {
		C.fmpz_pow_ui(&z.i[0], &x.i[0], C.fmpz_get_ui(&y.i[0]))
	} else {
		m.doinit()
		C.fmpz_powm(&z.i[0], &x.i[0], &y.i[0], &m.i[0])
	}
	return z
}

// ExpXY sets z = x**y and returns z.
func (z *Fmpz) ExpXY(x, y *Fmpz) *Fmpz {
	x.doinit()
	y.doinit()
	z.doinit()
	if y.Sign() <= 0 {
		return z.SetInt64(1)
	}

	C.fmpz_pow_ui(&z.i[0], &x.i[0], C.fmpz_get_ui(&y.i[0]))

	return z
}

// ExpXI sets z = x**y where u is an int type and returns z.
func (z *Fmpz) ExpXI(x *Fmpz, y int) *Fmpz {
	x.doinit()
	z.doinit()
	if y <= 0 {
		return z.SetInt64(1)
	}

	C.fmpz_pow_ui(&z.i[0], &x.i[0], C.mp_limb_t(y))

	return z
}

// ExpXIM sets z = x**i mod m where u is an int type and returns z.
func (z *Fmpz) ExpXIM(x *Fmpz, i int, m *Fmpz) *Fmpz {
	x.doinit()
	m.doinit()
	z.doinit()
	y := NewFmpz(int64(i))
	if y.Sign() <= 0 {
		return z.SetInt64(1)
	}
	if m == nil || m.Sign() == 0 {
		C.fmpz_pow_ui(&z.i[0], &x.i[0], C.fmpz_get_ui(&y.i[0]))
	} else {
		m.doinit()
		C.fmpz_powm(&z.i[0], &x.i[0], &y.i[0], &m.i[0])
	}
	return z
}

// ExpZ sets z = z**x and returns z.
func (z *Fmpz) ExpZ(x *Fmpz) *Fmpz {
	x.doinit()
	z.doinit()
	if x.Sign() <= 0 {
		return z.SetInt64(1)
	}

	C.fmpz_pow_ui(&z.i[0], &z.i[0], C.fmpz_get_ui(&x.i[0]))

	return z
}

// ExpI sets z = z**x where i is an int type and returns z.
func (z *Fmpz) ExpI(x int) *Fmpz {
	z.doinit()
	if x <= 0 {
		return z.SetInt64(1)
	}

	C.fmpz_pow_ui(&z.i[0], &z.i[0], C.mp_limb_t(x))

	return z
}

// Pow is a wrapper for Exp.
func (z *Fmpz) Pow(x, y, m *Fmpz) *Fmpz {
	return z.Exp(x, y, m)
}

// Square raises z to the power of 2 and returns z.
func (z *Fmpz) Square() *Fmpz {
	return z.ExpXI(z, 2)
}

// Cube raises z to the power of 3 and returns z.
func (z *Fmpz) Cube() *Fmpz {
	return z.ExpXI(z, 3)
}

/*
 * Greatest Common Divisor
 */

// GCD sets f to the greatest common divisor of g and h. The result is always positive, even if
// one of g and h is negative
func (z *Fmpz) GCD(g, h *Fmpz) *Fmpz {
	g.doinit()
	h.doinit()
	z.doinit()

	C.fmpz_gcd(&z.i[0], &g.i[0], &h.i[0])
	return z
}

// Lcm sets f to the least common multiple of g and h. The result is always nonnegative, even
// if one of g and h is negative.
func (z *Fmpz) Lcm(g, h *Fmpz) *Fmpz {
	g.doinit()
	h.doinit()
	z.doinit()

	C.fmpz_lcm(&z.i[0], &g.i[0], &h.i[0])
	return z
}

// GCDInv given integers f, g with 0 ≤ f < g, computes the greatest common divisor d = gcd(f, g)
// and the modular inverse a = f^-1 (mod g), whenever f != 0
// void fmpz_gcdinv (fmpz_t d , fmpz_t a , const fmpz_t f , const fmpz_t g )
func (z *Fmpz) GCDInv(g *Fmpz) (*Fmpz, *Fmpz) {

	d := new(Fmpz)
	a := new(Fmpz)
	z.doinit()
	g.doinit()
	d.doinit()
	a.doinit()
	C.fmpz_gcdinv(&d.i[0], &a.i[0], &z.i[0], &g.i[0])
	return d, a
}

// And sets z = x & y and returns z.
func (z *Fmpz) And(x, y *Fmpz) *Fmpz {
	x.doinit()
	y.doinit()
	z.doinit()
	C.fmpz_and(&z.i[0], &x.i[0], &y.i[0])
	return z
}

// Sqrt sets x to the truncated integer part of the square root of x
func (z *Fmpz) Sqrt(x *Fmpz) *Fmpz {
	x.doinit()
	z.doinit()
	C.fmpz_sqrt(&z.i[0], &x.i[0])
	return z
}

// Root sets x to the truncated integer part of the yth root of x
func (z *Fmpz) Root(x *Fmpz, y int32) *Fmpz {
	x.doinit()
	z.doinit()
	C.fmpz_root(&z.i[0], &x.i[0], C.slong(y))
	return z
}

// Arbitrary precision primality testing and factorization.

// IsStrongProbabPrime returns 1 if z is a strong probable prime to base a, otherwise it returns 0
func (z *Fmpz) IsStrongProbabPrime(a *Fmpz) int {
	a.doinit()
	z.doinit()
	return int(C.fmpz_is_strong_probabprime(&z.i[0], &a.i[0]))
}

// IsProbabPrimeLucas performs a Lucas probable prime test with parameters chosen by Selfridge's
// method A as per [4]. Return 1 if z is a Lucas probable prime, otherwise return 0. This function
// declares some composites probably prime, but no primes composite.
func (z *Fmpz) IsProbabPrimeLucas() int {
	z.doinit()
	return int(C.fmpz_is_probabprime_lucas(&z.i[0]))
}

// IsProbabPrimeBPSW performs a Baillie-PSW probable prime test with parameters chosen by
// Selfridge's method A as per [4]. Return 1 if z is a Lucas probable prime, otherwise return
// 0. There are no known composites passed as prime by this test, though infinitely many
// probably exist. The test will declare no primes composite.
func (z *Fmpz) IsProbabPrimeBPSW() int {
	z.doinit()
	return int(C.fmpz_is_probabprime_BPSW(&z.i[0]))
}

// IsProbabPrime performs some trial division and then some probabilistic primality tests.
// If z is definitely composite, the function returns 0, otherwise it is declared probably
// prime, i.e. prime for most practical purposes, and the function returns 1. The chance
// of declaring a composite prime is very small.
func (z *Fmpz) IsProbabPrime() int {
	z.doinit()
	return int(C.fmpz_is_probabprime(&z.i[0]))
}

// IsProbabPrimePseudosquare returns 0 is z is composite. If z is too large (greater
// than about 94 bits) the function fails silently and returns −1, otherwise, if z
// is proven prime by the pseudosquares method, return 1.
// Tests if z is a prime according to [28, Theorem 2.7]. We first factor N using trial
// division up to some limit B. In fact, the number of primes used in the trial factoring
// is at most FLINT_PSEUDOSQUARES_CUTOFF.
// Next we compute N/B and find the next pseudosquare Lp above this value, using a
// static table as per http://research.att.com/~njas/sequences/b002189.txt.
// As noted in the text, if p is prime then Step 3 will pass. This test rejects many
// composites, and so by this time we suspect that p is prime. If N is 3 or 7 modulo 8,
// we are done, and N is prime.
// We now run a probable prime test, for which no known counterexamples are known, to
// reject any composites. We then proceed to prove N prime by executing Step 4. In the
// case that N is 1 modulo 8, if Step 4 fails, we extend the number of primes pi at Step 3
// and hope to find one which passes Step 4. We take the test one past the largest p for
// which we have pseudosquares Lp tabulated, as this already corresponds to the next Lp
// which is bigger than 264 and hence larger than any prime we might be testing.
// As explained in the text, Condition 4 cannot fail if N is prime.
// The possibility exists that the probable prime test declares a composite prime. However
// in that case an error is printed, as that would be of independent interest.
func (z *Fmpz) IsProbabPrimePseudosquare() int {
	z.doinit()
	return int(C.fmpz_is_prime_pseudosquare(&z.i[0]))
}

// WilliamsPP1 uses Use Williams' p+1 method to factor n, using a prime bound in stage 1 of B1 and a
// prime limit in stage 2 of at least the square of B2_sqrt. If a factor is found, the function
// returns 1 and factor is set to the factor that is found. Otherwise, the function returns 0.
// c should be a random value greater than 2. Successive calls to the function
// with different values of c give additional chances to factor n with roughly exponentially
// decaying probability of finding a factor which has been missed (if p+1 or p−1 is not smooth for
// any prime factors p of n then the function will not ever succeed).
func (z *Fmpz) WilliamsPP1(n *Fmpz, b1, b2, c int) int {
	return int(C.fmpz_factor_pp1(&z.i[0], &n.i[0], C.ulong(b1), C.ulong(b2), C.ulong(c)))
}

// LucasChain Given V0 = 2, V1 = A compute Vm, Vm+1 (mod n) from the recurrences Vj = AVj−1 −
// Vj−2 (mod n).
func (z *Fmpz) LucasChain(v2, a, m, n *Fmpz) {
	z.doinit() // v1
	v2.doinit()
	a.doinit()
	m.doinit()
	n.doinit()
	C.fmpz_lucas_chain(&z.i[0], &v2.i[0], &a.i[0], &m.i[0], &n.i[0])
}

// Bits returns the number of bits required to store the absolute value of z. If z is 0 then 0 is
// returned.
func (z *Fmpz) Bits() int {
	z.doinit()
	return int(C.fmpz_bits(&z.i[0]))
}

// TstBit tests bit index i of z and return 0 or 1, accordingly.
func (z *Fmpz) TstBit(i int) int {
	z.doinit()
	return int(C.fmpz_tstbit(&z.i[0], C.mp_limb_t(i)))
}

// Random number generation.

// Randm sets z to a random integer between 0 and m-1 inclusive.
func (z *Fmpz) Randm(state *FlintRandT, m *Fmpz) *Fmpz {
	z.doinit()
	m.doinit()
	state.flintRandTDoinit()
	C.fmpz_randm(&z.i[0], &state.i[0], &m.i[0])

	return z
}

// Chinese Remainder Theorem.

// CRT uses the Chinese Remainder Theorem to set out to the unique value 0≤x<M (if sign = 0) or
// −M/2<x≤M/2 (if sign = 1) congruent to r1 modulo m1 and r2 modulo m2, where where M=m1×m2. It is
// assumed that m1 and m2 are positive integers greater than 1 and coprime.  If sign = 0, it is
// assumed that 0≤r1<m1 and 0≤r2<m2. Otherwise, it is assumed that −m1≤r1<m1 and 0≤r2<m2.
func (z *Fmpz) CRT(r1, m1, r2, m2 *Fmpz, sign int) *Fmpz {
	z.doinit()
	C.fmpz_CRT(&z.i[0], &r1.i[0], &m1.i[0], &r2.i[0], &m2.i[0], C.int(sign))
	return z
}

// Min and Max.

// Min finds the min(a, b) sets z to it and returns it.
func (z *Fmpz) Min(a, b *Fmpz) *Fmpz {
	if a.Cmp(z) < 0 {
		return z.Set(a)
	}

	return z.Set(b)
}

// Max finds the max(a, b) sets z to it and returns it.
func (z *Fmpz) Max(a, b *Fmpz) *Fmpz {
	if a.Cmp(z) > 0 {
		return z.Set(a)
	}

	return z.Set(b)
}

// Natural logarithm.

// DLog returns log(z) as a float64.
func (z *Fmpz) DLog() float64 {
	return float64(C.fmpz_dlog(&z.i[0]))
}
