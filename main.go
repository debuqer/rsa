package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	var p, q int
	p = 5
	q = 7

	N := p * q
	T := (p - 1) * (q - 1)

	var e int
	for i := 1; i < T; i++ {
		if big.NewInt(int64(i)).ProbablyPrime(20) {
			if T%i != 0 {
				e = i
			}
		}
	}

	var d int
	for j := 1; j < 100; j++ {
		if e*j%T == 1 {
			d = j
			break
		}
	}

	t := 2

	fmt.Printf("Public key: (%d, %d)\n", e, N)
	fmt.Printf("Private key: (%d, %d)\n", d, N)
	encrypted := math.Mod(math.Pow(float64(t), float64(e)), float64(N))
	fmt.Printf("encrypt: %d %d\n", t, uint(encrypted))

	decrypted := math.Mod(math.Pow(float64(encrypted), float64(d)), float64(N))
	fmt.Printf("decrypted: %d %d\n", uint(encrypted), uint(decrypted))
}
