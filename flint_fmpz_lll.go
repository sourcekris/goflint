package goflint

/*
#cgo LDFLAGS: -lflint -lgmp
#include <flint/fmpz_lll.h>

// Macros

*/
import "C"

// FmpzLLL stores a LLL matrix reduction context including delta, eta, rt and gt values.
type FmpzLLL struct {
	i    C.fmpz_lll_t
	init bool
}

// fmpzLLLDoinit initializes an FmpzLLL type.
func (l *FmpzLLL) fmpzLLLDoinit() {
	if l.init {
		return
	}

	l.init = true
	C.fmpz_lll_context_init_default(&l.i[0])
}

func NewFmpzLLL() *FmpzLLL {
	l := new(FmpzLLL)
	l.fmpzLLLDoinit()
	return l
}

// LLL reduces m in place according to the parameters specified by the default LLL context of
// fl->delta, fl->eta, fl->rt and fl->gt set to 0.99, 0.51, ZBASIS and APPROX respectively.
// u is the matrix used to capture the unimodular transformations if it is not NULL.
func (m *FmpzMat) LLL() *FmpzMat {
	l := NewFmpzLLL()
	C.fmpz_lll(&m.i[0], nil, &l.i[0])
	return m
}
