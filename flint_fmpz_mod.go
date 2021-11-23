package goflint

/*
#cgo LDFLAGS: -lflint -lgmp
#include <flint/fmpz_mod.h>

*/
import "C"

import "runtime"

type FmpzModCtx struct {
	i    C.fmpz_mod_ctx_t
	n    *Fmpz
	init bool
}

// fmpzModCtxFinalize releases the memory allocated to the FmpzModCtx.
func fmpzModCtxFinalize(z *FmpzModCtx) {
	if z.init {
		runtime.SetFinalizer(z, nil)
		C.fmpz_mod_ctx_clear(&z.i[0])
		z.init = false
	}
}

// fmpzModCtxDoinit initializes an FmpzModCtx type.
func (z *FmpzModCtx) fmpzModCtxDoinit(n *Fmpz) {
	if z.init {
		return
	}
	z.init = true
	C.fmpz_mod_ctx_init(&z.i[0], &n.i[0])
	z.n = n
	runtime.SetFinalizer(z, fmpzModCtxFinalize)
}

// fmpzModCtxDoinitNF initializes an FmpzModCtx type.
func (z *FmpzModCtx) fmpzModCtxDoinitNF(n *Fmpz) {
	if z.init {
		return
	}
	z.init = true
	C.fmpz_mod_ctx_init(&z.i[0], &n.i[0])
	z.n = n
}

// NewFmpzModCtx allocates a new FmpzModCtx with modulus n and returns it.
func NewFmpzModCtx(n *Fmpz) *FmpzModCtx {
	p := new(FmpzModCtx)
	p.fmpzModCtxDoinit(n)
	return p
}

// NewFmpzModCtxNF allocates a new FmpzModCtx with modulus n and returns it.
func NewFmpzModCtxNF(n *Fmpz) *FmpzModCtx {
	p := new(FmpzModCtx)
	p.fmpzModCtxDoinitNF(n)
	return p
}
