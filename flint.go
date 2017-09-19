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