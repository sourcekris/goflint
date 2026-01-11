## goflint

Golang wrapper for [FLINT](http://www.flintlib.org) functions that I find useful for use upstream
in [goRsaTool](https://github.com/sourcekris/goRsaTool).

This supports FLINT2 and FLINT3 library versions but due to API changes over time has not been
exhaustively integration tested.

## Features

### Assignments
 * `(z *Fmpz) Set(x *Fmpz) *Fmpz` Set z to the Fmpz x and return z
 * `(z *Fmpz) SetUint64(x uint64) *Fmpz` Sets Fpmz z to ulong x and returns z
 * `(z *Fmpz) SetInt64(x int64) *Fmpz` Sets Fpmz z to slong x and returns z
 * `(z *Mpz) SetMpzInt64(x int64) *Mpz` Sets Mpz z to slong x and returns z
 * `NewFmpz(x int64) *Fmpz` allocates a new Fpmz and sets it to x
 * `NewMpz(x int64) *Mpz` allocates a new GMP Mpz and sets it to x
 * `NewFmpq(p, q int64) *Fmpq` allocates and returns a new Fmpq set to p / q
 * `NewFmpqFmpz(p, q *Fmpz) *Fmpq` allocates and returns a new Fmpq set to p / q where p and q are Fmpz types.
 * `(q *Fmpq) SetFmpqFraction(num, den *Fmpz) *Fmpq` sets the value of q to the fraction num / den and returns q.

### Comparisons
 * `(z *Fmpz) Cmp(y *Fmpz) (r int)` Compares z to y and returns -1, 0, or 1
 * `(z *Mpz) Cmp(y *Mpz) (r int)` Compares z to y and returns -1, 0, or 1 where z is an Mpz type.
 * `(z *Fmpq) CmpRational(y *Fmpq) (r int)` Compares rational z to y and returns -1, 0, or 1
 * `(q *Fmpq) Cmp(y *Fmpq) int` Compares rationals z and y and returns -1, 0, or 1
 * `(z *Fmpz) Equals(y *Fmpz) bool` Compares z and y and returns true if they are equal.
 * `(z *Fmpz) IsZero() bool` Returns true if z == 0.

### Formatters
 * `(z *Fmpz) String() string` Returns a base 10 string representaiton of z
 * `(q *Fmpq) String() string` Returns a base 10 string representation of rational q

### Helpers
 * `(z *Fmpz) BitLen() int` Returns the length of z in bits
 * `(z *Fmpz) Bits() int` Returns the number of bits required to store z
 * `(z *Fmpz) Sign() (r int)` Returns the sign of z returns -1, 0, or 1
 * `(z *Fmpz) Lsh(bits int) *Fmpz` Left shifts an Fmpz z by an arbitrary number of bits and returns it.
 * `(z *Fmpz) Rsh(bits int) *Fmpz` Right shifts an Fmpz z by an arbitrary number of bits and returns it.

### Primality Testing and Factorization
 * `(z *Fmpz) IsStrongProbabPrime(a *Fmpz)` returns 1 if z is a strong probable prime to base a, otherwise it returns 0
 * `(z *Fmpz) IsProbabPrimeLucas() int` performs a Lucas probable prime test returns 1 if z is a Lucas probable prime, otherwise return 0
 * `(z *Fmpz) IsProbabPrimeBPSW() int` performs a Baillie-PSW probable prime test returns 1 if z is a probable prime, otherwise return 0
 * `(z *Fmpz) IsProbabPrime() int` returns 1 if z is a probable prime, otherwise return 0
 * `(z *Fmpz) IsProbabPrimePseudosquare() int` returns 0 is z is composite. If z is too large (greater than about 94 bits) the function fails silently and returns −1, otherwise, if z is proven prime by the pseudosquares method, return 1.
 * `(z *Fmpz) LucasChain(v2, a, m, n *Fmpz)` Given V0 = 2, V1 = A compute Vm, Vm+1 (mod n) from the recurrences Vj = AVj−1 − Vj−2 (mod n).

### Random Number Generation
 * `(z *Fmpz) Randm(state *FlintRandT, m *Fmpz) *Fmpz` Sets z to a random number between 0 and m-1 inclusive

### Conversions
 * `(f *Fmpz) GetInt() int` Lowers f to type int
 * `(f *Fmpz) GetUInt() uint` Lowers f to type uint
 * `(z *Fmpz) Int64() (y int64)` Lowers z to type int64
 * `(z *Fmpz) Uint64() (y uint64)` Lower z to type uint64
 * `(q *Fmpq) GetFmpqFraction() (int, int)` gets the numerator and denomenator of the rational q returning them as ints.
 * `(q *Fmpq) NumRef() int` returns the numerator of an Fmpq as an integer.
 * `(q *Fmpq) DenRef() int` returns the denominator of an Fmpq as an integer.
 * `(z *Fmpz) SetString(s string, base int) (*Fmpz, bool)` Sets z to the value in string s using given base 
 * `(z *Fmpz) SetMpz(x *Mpz)` Set z to the value in Mpz x
 * `(z *Mpz) GetMpz(x *Fmpz)` Set Mpz z to the value of the Fmpz x
 * `(z *Fmpz) SetBytes(buf []byte) *Fmpz` Set z to the value stored in byte array buf and return z
 * `(z *Fmpz) Bytes() []byte` Return the bytes of Fmpz z

### Arithmetic
 * `(z *Fmpz) Abs(x *Fmpz) *Fmpz` Set z to the absolute value of x and return z
 * `(z *Fmpz) Neg(x *Fmpz) *Fmpz` Set z to the negated value of x and return z
 * `(z *Fmpz) Add(x, y *Fmpz) *Fmpz` Set z to x + y and return z
 * `(z *Fmpz) AddZ(x *Fmpz) *Fmpz` Set z to z + x and return z
 * `(z *Fmpz) AddI(i int) *Fmpz` Set z to z + i where i is an int type and return z
 * `(z *Fmpz) Sub(x, y *Fmpz) *Fmpz` Set z to x - y and return z
 * `(z *Fmpz) SubZ(x *Fmpz) *Fmpz` Set z to z - x and return z
 * `(z *Fmpz) SubI(i int) *Fmpz` Set z to z - i where i is an int type and return z
 * `(z *Mpz) SubRMpz(y, n *Mpz)` Sets z to z - y in the ring of integers modulo n using Mpz types.
 * `(z *Fmpz) Mul(x, y *Fmpz) *Fmpz` Set z to x * y and return z
 * `(z *Fmpz) MulZ(x *Fmpz) *Fmpz` Set z to z * x and return z
 * `(z *Fmpz) MulI(i int) *Fmpz` Set z to z * i where i is an int type and return z
 * `(z *Mpz) MulRMpz(y, n *Mpz) *Mpz` Sets z to z * y in the integer ring modulo n using Mpz types.
 * `(q *Fmpq) MulRational(o *Fmpq, x *Fmpz) *Fmpq` Sets q to the product of rational o and Fmpz x and returns q.
 * `(z *Fmpz) DivR(y, n *Fmpz) *Fmpz` Sets z to the result of z/y in the ring of integers modulo n. Only works if y fits into the int type
 * `(z *Fmpz) Div(x, y *Fmpz) *Fmpz` Set z to x / y and return z
 * `(z *Fmpz) Quo(x, y *Fmpz) *Fmpz`
 * `(z *Fmpz) QuoRem(x, y, r *Fmpz) (*Fmpz, *Fmpz)`
 * `(z *Fmpz) Mod(x, y *Fmpz) *Fmpz` Set z to the value of x % y and return z
 * `(z *Fmpz) ModZ(y *Fmpz) *Fmpz` Set z to the value of z % y and return z
 * `(z *Fmpz) ModRational(x *Fmpq, n *Fmpz) int` Sets z to the residue of x = n/d (num, den) modulo 
  n and returns 1 if such a modulo exists or 0 if it does not
 * `(z *Fmpz) DivMod(x, y, m *Fmpz) (*Fmpz, *Fmpz)`
 * `(z *Fmpz) ModInverse(x, y *Fmpz) *Fmpz`
 * `(z *Fmpz) NegMod(x, y *Fmpz) *Fmpz`
 * `(a *Fmpz) Jacobi(p *Fmpz) int`
 * `(z *Fmpz) Exp(x, y, m *Fmpz) *Fmpz` Set z to the value of (x^y)%m and return z
 * `(z *Fmpz) ExpZ(x *Fmpz) *Fmpz` Set z to the value of (z^x) and return z
 * `(z *Fmpz) ExpI(x int) *Fmpz` Set z to the value of (z^i) where i is an int type and return z
 * `(z *Fmpz) ExpXY(x, y *Fmpz) *Fmpz` Set z to the value of (x^y) and return z
 * `(z *Fmpz) ExpXI(x *Fmpz, y int) *Fmpz` Set z to the value of (x^y) where y is an int type and return z
 * `(z *Fmpz) ExpXIM(x *Fmpz, i int, m *Fmpz) *Fmpz` Set z to the value of (x^y)%m where y is an int type and return z 
 * `(z *Fmpz) Pow(x, y, m *Fmpz) *Fmpz` Set z to the value of (x^y)%m and return z
 * `(z *Fmpz) Square() *Fmpz` raises z to the power of 2 and returns z.
 * `(z *Fmpz) Cube() *Fmpz` raises z to the power of 3 and returns z.
 * `(f *Fmpz) GCD(g, h *Fmpz) *Fmpz` Set z to the value of the greatest common divisor of g and h and return z
 * `(f *Fmpz) Lcm(g, h *Fmpz) *Fmpz` Set z to the value of the lowest common multiple of g and h and return z 
 * `(f *Fmpz) GCDInv(g *Fmpz) (*Fmpz, *Fmpz)`

### Bitwise Operations
 * `(z *Fmpz) And(x, y *Fmpz) *Fmpz` Set z to the value of x & y and return z
 * `(z *Fmpz) Xor(a, b *Fmpz) *Fmpz` Set z to the bitwise exclusive or of a and b and returns z.
 * `(z *Fmpz) TstBit(i int) int` Returns the value of the bit stored at index i where 0 is the least significant bit.

### Roots
 * `(z *Fmpz) Sqrt(x *Fmpz) *Fmpz` Set z to the value of the square root of x and return z
 * `(z *Fmpz) Root(x *Fmpz, y int32) *Fmpz` Set z to the value of then yth root of x and return z

### Matrices
 * `NewFmpzMat(rows, cols int) *FmpzMat` Creates and allocates a new FmpzMat matrix type of size rows * cols.
 * `(m *FmpzMat) String() string` Returns a pretty printed string version of the matrix as a string.
 * `(m *FmpzMat) Zero() *FmpzMat` Sets all of the values of the matrix to zero and returns the matrix.
 * `(m *FmpzMat) One() *FmpzMat` Sets the diagonal values of the matrix to 1 and returns the matrix.
 * `(m *FmpzMat) NumRows() int ` Returns the number of rows in the matrix as an integer.
 * `(m *FmpzMat) NumCols() int` Returns the number of columns in the matrix.
 * `(m *FmpzMat) Entry(x, y int) *Fmpz` Returns the value at coordinates x, y in the matrix m.
 * `(m *FmpzMat) SetPosVal(val *Fmpz, pos int) *FmpzMat` Sets the value at offset pos in the matrix m and returns m.
 * `(m *FmpzMat) SetVal(val *Fmpz, x, y int) *FmpzMat` Sets the value at coordinates x,y in the matrix and returns m.

### Lattice Basis Reduction
 * `NewFmpzLLL() *FmpzLLL` Creates and allocates a new FmpzLLL context.
 * `(m *FmpzMat) LLL() *FmpzMat` Reduces m in place according to the parameters specified by the default LLL context.

### Univariate Polynomials over the integers.
 * `NewFmpzPoly() *FmpzPoly` NewFmpzPoly allocates a new FmpzPoly and returns it.
 * `NewFmpzPoly2(a int) *FmpzPoly` NewFmpzPoly2 allocates a new FmpzPoly with at least a coefficients and returns it.
 * `NewFmpzPolyFactor() *FmpzPolyFactor` allocates a new FmpzPolyFactor and returns it.
 * `FmpzPolySetString(poly string) (*FmpzPoly, error)` FmpzPolySetString returns a polynomial using the string representation as the definition.
 * `(z *FmpzPoly) Set(poly *FmpzPoly) *FmpzPoly` Set sets z to poly and returns z
 * `(f *FmpzPolyFactor) Set(fac *FmpzPolyFactor) *FmpzPolyFactor` Set sets f to FmpzPolyFactor fac and returns f.
 * `(z *FmpzPoly) String() string` String returns a string representation of the polynomial.
 * `(z *FmpzPoly) StringSimple() string` StringSimple returns a simple string representation of the polynomials length and coefficients. 
 * `(f *FmpzPolyFactor) Print()` Print prints the FmpzPolyFactor to stdout.
 * `(z *FmpzPoly) Zero() *FmpzPoly` Zero sets z to the zero polynomial and returns z.
 * `(z *FmpzPoly) FitLength(l int)` FitLength sets the number of coefficiets in z to l.
 * `(z *FmpzPoly) SetCoeff(c int, x *Fmpz) *FmpzPoly` SetCoeff sets the c'th coefficient of z to x where x is an Fmpz and returns z.
 * `(z *FmpzPoly) SetCoeffUI(c int, x uint) *FmpzPoly` SetCoeffUI sets the c'th coefficient of z to x where x is an uint and returns z.
 * `(z *FmpzPoly) GetCoeff(c int) *Fmpz` GetCoeff gets the c'th coefficient of z and returns an Fmpz.
 * `(z *FmpzPoly) GetCoeffs() []*Fmpz` GetCoeffs gets all of the coefficient of z and returns a slice of Fmpz.
 * `(z *FmpzPoly) Len() int` Len returns the length of the poly z.
 * `(z *FmpzPoly) Neg(p *FmpzPoly) *FmpzPoly` Neg sets z to the negative of p and returns z.
 * `(z *FmpzPoly) GCD(a, b *FmpzPoly) *FmpzPoly` GCD sets z = gcd(a, b) and returns z.
 * `(z *FmpzPoly) Equal(p *FmpzPoly) bool` Equal returns true if z is equal to p otherwise false.
 * `(z *FmpzPoly) Add(a, b *FmpzPoly) *FmpzPoly` Add sets z = a + b and returns z.
 * `(z *FmpzPoly) Sub(a, b *FmpzPoly) *FmpzPoly` Sub sets z = a - b and returns z.
 * `(z *FmpzPoly) Mul(a, b *FmpzPoly) *FmpzPoly` Mul sets z = a * b and returns z.
 * `(z *FmpzPoly) MulScalar(a *FmpzPoly, x *Fmpz) *FmpzPoly` MulScalar sets z = a * x where x is an Fmpz.
 * `(z *FmpzPoly) DivScalar(a *FmpzPoly, x *Fmpz) *FmpzPoly` DivScalar sets z = a / x where x is an Fmpz. Rounding coefficients down toward -infinity.
 * `(z *FmpzPoly) Pow(m *FmpzPoly, e int) *FmpzPoly` Pow sets z to m^e and returns z.
 * `(z *FmpzPoly) DivRem(m *FmpzPoly) (*FmpzPoly, *FmpzPoly)` DivRem computes q, r such that z=mq+r and 0 ≤ len(r) < len(m).
 * `(z *FmpzPoly) Factor() *FmpzPolyFactor` Factor uses the Zassenhaus factoring algorithm.
 * `(f *FmpzPolyFactor) GetPoly(n int) *FmpzPoly` GetPoly gets the nth polynomial factor from a FmpzPolyFactor and returns it.
 * `(f *FmpzPolyFactor) GetExp(n int) int` GetExp gets the exponent of the nth polynomial from the FmpzPolyFactor.
 * `(f *FmpzPolyFactor) GetCoeff() *Fmpz` GetCoeff gets the coefficient from the FmpzPolyFactor.
 * `(f *FmpzPolyFactor) Len() int` Len gets the length of the FmpzPolyFactors list. i.e. the number of factors found.

### Univariate Polynomials over the integers modulo n.
 * `NewFmpzModCtx(n *Fmpz) *FmpzModCtx` NewFmpzModCtx allocates a new FmpzModCtx with modulus n and returns it.
 * `NewFmpzModPoly(n *FmpzModCtx) *FmpzModPoly` NewFmpzModPoly allocates a new FmpzModPoly mod n and returns it.
 * `NewFmpzModPoly2(n *FmpzModCtx, a int) *FmpzModPoly` NewFmpzModPoly2 allocates a new FmpzModPoly mod n with at least a coefficients and returns it.
 * `SetString(poly string) (*FmpzModPoly, error)` SetString returns a polynomial mod n using the string representation as the definition.
 * `(z *FmpzModPoly) Set(poly *FmpzModPoly) *FmpzModPoly` Set sets z to poly and returns z
 * `(z *FmpzModPoly) String() string` String returns a string representation of the polynomial.
 * `(z *FmpzModPoly) StringSimple() string` StringSimple returns a simple string representation of the polynomials length, modulus and coefficients. 
 * `(z *FmpzModPoly) Zero() *FmpzModPoly` Zero sets z to the zero polynomial and returns z.
 * `(z *FmpzModPoly) FitLength(l int)` FitLength sets the number of coefficiets in z to l.
 * `(z *FmpzModPoly) SetCoeff(c int, x *Fmpz) *FmpzModPoly` SetCoeff sets the c'th coefficient of z to x where x is an Fmpz and returns z.
 * `(z *FmpzModPoly) SetCoeffUI(c int, x uint) *FmpzModPoly` SetCoeffUI sets the c'th coefficient of z to x where x is an uint and returns z.
 * `(z *FmpzModPoly) GetCoeff(c int) *Fmpz` GetCoeff gets the c'th coefficient of z and returns an Fmpz.
 * `(z *FmpzModPoly) GetCoeffs() []*Fmpz` GetCoeffs gets all of the coefficient of z and returns a slice of Fmpz.
 * `(z *FmpzModPoly) GetMod() *Fmpz` GetMod gets the modulus of z and returns an Fmpz.
 * `(z *FmpzModPoly) Len() int` Len returns the length of the poly z.
 * `(z *FmpzModPoly) Neg(p *FmpzModPoly) *FmpzModPoly` Neg sets z to the negative of p and returns z.
 * `(z *FmpzModPoly) GCD(a, b *FmpzModPoly) *FmpzModPoly` GCD sets z = gcd(a, b) and returns z.
 * `(z *FmpzModPoly) Equal(p *FmpzModPoly) bool` Equal returns true if z is equal to p otherwise false.
 * `(z *FmpzModPoly) Add(a, b *FmpzModPoly) *FmpzModPoly` Add sets z = a + b and returns z.
 * `(z *FmpzModPoly) Sub(a, b *FmpzModPoly) *FmpzModPoly` Sub sets z = a - b and returns z.
 * `(z *FmpzModPoly) Mul(a, b *FmpzModPoly) *FmpzModPoly` Mul sets z = a * b and returns z.
 * `(z *FmpzModPoly) MulScalar(a *FmpzModPoly, x *Fmpz) *FmpzModPoly` MulScalar sets z = a * x where x is an Fmpz.
 * `(z *FmpzModPoly) DivScalar(a *FmpzModPoly, x *Fmpz) *FmpzModPoly` DivScalar sets z = a / x where x is an Fmpz.
 * `(z *FmpzModPoly) Pow(m *FmpzModPoly, e int) *FmpzModPoly` Pow sets z to m^e and returns z.
 * `(z *FmpzModPoly) DivRem(m *FmpzModPoly) (*FmpzModPoly, *FmpzModPoly)` DivRem computes q, r such that z=mq+r and 0 ≤ len(r) < len(m).

### Chinese Remainder Theorem
 * `(z *Fmpz) CRT(r1, m1, r2, m2 *Fmpz, sign int) *Fmpz` uses the Chinese Remainder Theorem to set out to the unique value.

### Min and Max
 * `(z *Fmpz) Min(a, b *Fmpz) *Fmpz` finds the min(a, b) sets z to it and returns it.
 * `(z *Fmpz) Max(a, b *Fmpz) *Fmpz` finds the max(a, b) sets z to it and returns it.

### Natural Logarithm
 * `(z *Fmpz) DLog() float64` returns log(z) as a float64.

## Types
```
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

// Fmpq is an arbitrary precision rational type.
type Fmpq struct {
  i    C.fmpq_t
  init bool
}

// FmpzPoly type represents a univatiate polynomial over the integers.
type FmpzPoly struct {
	i    C.fmpz_poly_t
	init bool
}

// FmpzPolyFactor type represents the factors univariate polynomial over the integers.
type FmpzPolyFactor struct {
	i    C.fmpz_poly_factor_t
	init bool
}

// FmpzModPoly type represents elements of Z/nZ[x] for a fixed modulus n.
type FmpzModPoly struct {
	i    C.fmpz_mod_poly_t
	ctx  *FmpzModCtx
	init bool
}

// FmpzModCtx type represents a context for modular arithmetic.
type FmpzModCtx struct {
	i    C.fmpz_mod_ctx_t
	n    *Fmpz
	init bool
}

// NmodPoly type represents elements of Z/nZ[x] for a fixed modulus n.
type NmodPoly struct {
  i    C.nmod_poly_t
  init bool
}

// MpLimb type is a uint64.
type MpLimb struct {
  i    C.mp_limb_t
}

// FlintRandT keeps state for Fmpz random number generation.
type FlintRandT struct {
	i    C.flint_rand_t
	init bool
}

// FmpzMat is a matrix of Fmpz.
type FmpzMat struct {
	i    C.fmpz_mat_t
	rows int
	cols int
	init bool
}

// FmpzLLL stores a LLL matrix reduction context.
type FmpzLLL struct {
	i    C.fmpz_lll_t
	init bool
}
```

## Examples

```
a := NewFmpz(1)

fmt.Println(a.String())
```


## Install

Install the golang and the FLINT library. 

 * On Ubuntu (tested on 20.04LTS):
   ```shell
   $ apt install golang libflint-dev`
   ```
 * On MacOS (tested on Monterey 12.1 M1)
   ```shell
   $ brew install flint
   ```

Use go to install the library:
* `go get github.com/sourcekris/goflint`

## License

Apache 2.0. See the LICENSE file for details.

## Note About Prior Art

I found this on the internet, it might help me get going quicker but seems written to do a few tasks only:
 * https://github.com/faahmed/goflint

I am also heavily influenced to do this by nick [at] craig-wood.com due to this repo:
 * https://github.com/ncw/gmp

## Authors

* Kris Hunt

## Contributors

* https://github.com/x9xhack
