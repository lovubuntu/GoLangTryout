package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	z := complex128(1)
    if (imag(x) != 0) || (real(x) < 0){
        z += 1i
    }
    i:=1
	for ; i < 500000; i++ {
		if v := (2*cmplx.Pow(z, 3) + x) / (3 * cmplx.Pow(z, 2)); v == z {
			return z
		} else {
			z = v
		}
	}
    fmt.Println(i)
	return z
}

func main() {
	fmt.Println(Cbrt(-27))
    fmt.Println(cmplx.Pow(-27,1.0/3.0))
}
