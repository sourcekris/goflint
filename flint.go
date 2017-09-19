// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file wraps some FLINT (Fast Library for Number Theory) functions

package goflint

/*
#cgo LDFLAGS: -lflint -lgmp
#include <flint/flint.h>
#include <flint/fmpz.h>
#include <gmp.h>
#include <stdlib.h>
*/
import "C"

import (
  "runtime"
  "unsafe"
)

type Fmpz struct {
	i    C.fmpz_t
	init bool
}

// Finalizer - release the memory allocated to the fmpz
func fmpzFinalize(z *Fmpz) {
	if z.init {
		runtime.SetFinalizer(z, nil)
		C.fmpz_clear(&z.i[0])
		z.init = false
	}
}

func (z *Fmpz) doinit() {
	if z.init {
		return
	}
	z.init = true
	C.fmpz_init(&z.i[0])
	runtime.SetFinalizer(z, fmpzFinalize)
}

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

// NewFmpz allocates and returns a new Fmpz set to x.
func NewFmpz(x int64) *Fmpz {
  return new(Fmpz).SetInt64(x)
}

// Cmp compares z and y and returns:
//
//   -1 if z <  y
//    0 if z == y
//   +1 if z >  y
//
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

// String returns the decimal representation of z.
func (z *Fmpz) String() string {
  return z.string(10)
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


// BitLen returns the length of the absolute value of z in bits.
// The bit length of 0 is 0.
func (z *Fmpz) BitLen() int {
  z.doinit()
  if z.Sign() == 0 {
    return 0
  }
  return int(C.fmpz_sizeinbase(&z.i[0], 2))
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

// Set sets z to x and returns z.
func (z *Fmpz) Set(x *Fmpz) *Fmpz {
  z.doinit()
  C.fmpz_set(&z.i[0], &x.i[0])
  return z
}

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

// Sub sets z to the difference x-y and returns z.
func (z *Fmpz) Sub(x, y *Fmpz) *Fmpz {
  x.doinit()
  y.doinit()
  z.doinit()
  C.fmpz_sub(&z.i[0], &x.i[0], &y.i[0])
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

// Mod sets z to the modulus x%y for y != 0 and returns z.
func (z *Fmpz) Mod(x, y *Fmpz) *Fmpz {
  x.doinit()
  y.doinit()
  z.doinit()

  C.fmpz_mod(&z.i[0], &x.i[0], &y.i[0])
  return z
}