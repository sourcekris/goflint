## goflint

Golang wrapper for [FLINT2](http://www.flintlib.org) functions that I find useful. 

This project is heavily influenced by and in the same pattern as Golang's [GMP wrapper](http://golang.org/misc/cgo/gmp/gmp.go), which seemed appropriate since FLINT uses a similar API to GMP.

## Features

 * `(z *Fmpz) SetUint64(x uint64) *Fmpz` Sets Fpmz z to ulong x and returns z
 * `(z *Fmpz) SetInt64(x int64) *Fmpz` Sets Fpmz z to slong x and returns z
 * `(z *Mpz) SetMpzInt64(x int64) *Mpz` Sets Mpz z to slong x and returns z
 * `NewFmpz(x int64) *Fmpz` allocates a new Fpmz and sets it to x
 * `NewMpz(x int64) *Mpz` allocates a new GMP Mpz and sets it to x
 * `(z *Fmpz) Cmp(y *Fmpz) (r int)` Compares z to y and returns -1, 0, or 1
 * `(z *Fmpz) String() string` Returns a base 10 string representaiton of z
 * `(z *Fmpz) BitLen() int` Returns the length of z in bits
 * `(z *Fmpz) Sign() (r int)` Returns the sign of z returns -1, 0, or 1
 * `(z *Fmpz) Set(x *Fmpz) *Fmpz` Set z to the Fmpz x and return z
 * `(f *Fmpz) GetInt() int` Lowers f to type int
 * `(f *Fmpz) GetUInt() uint` Lowers f to type uint
 * `(z *Fmpz) Int64() (y int64)` Lowers z to type int64
 * `(z *Fmpz) Uint64() (y uint64)` Lower z to type uint64
 * `(z *Fmpz) SetString(s string, base int) (*Fmpz, bool)` Sets z to the value in string s using given base 
 * `(z *Fmpz) SetMpz(x *Mpz)` Set z to the value in Mpz x
 * `(z *Mpz) GetMpz(x *Fmpz)` Set Mpz z to the value of the Fmpz x
 * `(z *Fmpz) SetBytes(buf []byte) *Fmpz` Set z to the value stored in byte array buf and return z
 * `(z *Fmpz) Bytes() []byte` Return the bytes of Fmpz z
 * `(z *Fmpz) Abs(x *Fmpz) *Fmpz` Set z to the absolute value of x and return z
 * `(z *Fmpz) Neg(x *Fmpz) *Fmpz` Set z to the negated value of x and return z
 * `(z *Fmpz) Add(x, y *Fmpz) *Fmpz` Set z to x + y and return z
 * `(z *Fmpz) Sub(x, y *Fmpz) *Fmpz` Set z to x - y and return z
 * `(z *Fmpz) Mul(x, y *Fmpz) *Fmpz` Set z to x * y and return z
 * `(z *Fmpz) Div(x, y *Fmpz) *Fmpz` Set z to x / y and return z
 * `(z *Fmpz) Quo(x, y *Fmpz) *Fmpz`
 * `(z *Fmpz) QuoRem(x, y, r *Fmpz) (*Fmpz, *Fmpz)`
 * `(z *Fmpz) Mod(x, y *Fmpz) *Fmpz` Set z to the value of x % y and return z
 * `(z *Fmpz) DivMod(x, y, m *Fmpz) (*Fmpz, *Fmpz)`
 * `(z *Fmpz) ModInverse(x, y *Fmpz) *Fmpz`
 * `(z *Fmpz) NegMod(x, y *Fmpz) *Fmpz`
 * `(a *Fmpz) Jacobi(p *Fmpz) int`
 * `(z *Fmpz) Exp(x, y, m *Fmpz) *Fmpz` Set z to the value of (x^y)%m and return z
 * `(f *Fmpz) GCD(g, h *Fmpz) *Fmpz` Set z to the value of the greatest common divisor of g and h and return z
 * `(f *Fmpz) Lcm(g, h *Fmpz) *Fmpz` Set z to the value of the lowest common multiple of g and h and return z 
 * `(f *Fmpz) GCDInv(g *Fmpz) (*Fmpz, *Fmpz)`
 * `(z *Fmpz) And(x, y *Fmpz) *Fmpz` Set z to the value of x & y and return z
 * `(z *Fmpz) Sqrt(x *Fmpz) *Fmpz` Set z to the value of the square root of x and return z
 * `(z *Fmpz) Root(x *Fmpz, y int32) *Fmpz` Set z to the value of then yth root of x and return z
 
## Types
```
type Fmpz struct {
  i    C.fmpz_t
  init bool
}

type Mpz struct {
  i    C.mpz_t
  init bool
}
```

## Examples

```
a := NewFmpz(1)

fmt.Println(a.String())
```


## Install

Use go to install the library:
* `go get github.com/sourcekris/goflint`

## License

As this contains a great deal of code copied from the Go source it is licenced identically to the Go source itself - see the LICENSE file for details.

## Note About Prior Art

I found this on the internet, it might help me get going quicker but seems written to do a few tasks only:
 * https://github.com/faahmed/goflint

I am also heavily influenced to do this by nick [at] craig-wood.com due to this repo:
 * https://github.com/ncw/gmp

## Authors

* [The Go team](http://golang.org/AUTHORS)
* Kris Hunt
