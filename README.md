## goflint

Golang wrapper for [FLINT2](http://www.flintlib.org) functions that I find useful. 

This project is heavily influenced by and in the same pattern as Golang's [GMP wrapper](http://golang.org/misc/cgo/gmp/gmp.go), which seemed appropriate since FLINT uses a similar API to GMP.

## Features

 * `NewFmpz(x int64) *Fmpz` allocates a new Fpmz and sets it to x
 * `(z *Fmpz) SetInt64(x int64) *Fmpz` Sets Fpmz z to slong x and returns z
 * `(z *Fmpz) SetUint64(x uint64) *Fmpz` Sets Fpmz z to ulong x and returns z
 
## Types
```
type Fmpz struct {
  i    C.fmpz_t
  init bool
}
```

## Examples

Very little is implemented yet. I've only spent 30 minutes on this!

```
a := NewFmpz(1)
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
