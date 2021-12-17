package goflint

/*

#if __FLINT_RELEASE >= 20503
	// Use modern libflint.
	#include <flint/fmpz_mod.h>
#else
	// libflint 2.5.2 or below.
	// Sketch out a skeleton type and constructor.
	#include <flint/fmpz.h>
	typedef struct fmpz_mod_ctx {
		fmpz_t n;
		void (* add_fxn)(fmpz_t, const fmpz_t, const fmpz_t, const struct fmpz_mod_ctx *);
		void (* sub_fxn)(fmpz_t, const fmpz_t, const fmpz_t, const struct fmpz_mod_ctx *);
		void (* mul_fxn)(fmpz_t, const fmpz_t, const fmpz_t, const struct fmpz_mod_ctx *);
		nmod_t mod;
		ulong n_limbs[3];
		ulong ninv_limbs[3];
	} fmpz_mod_ctx_struct;
	typedef fmpz_mod_ctx_struct fmpz_mod_ctx_t[1];

	void fmpz_mod_ctx_init(fmpz_mod_ctx_t ctx, const fmpz_t n)
	{
		fmpz_init_set(ctx->n, n);
	}

	void fmpz_mod_ctx_clear(fmpz_mod_ctx_t ctx)
	{
    	fmpz_clear(ctx->n);
	}
#endif

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
