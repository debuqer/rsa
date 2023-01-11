package main

import (
	"fmt"
	"math/big"
)

func main() {
	p, _ := new(big.Int).SetString("7", 10)
	q, _ := new(big.Int).SetString("23", 10)

	zero, _ := new(big.Int).SetString("0", 10)
	one, _ := new(big.Int).SetString("1", 10)
	max_try, _ := new(big.Int).SetString("100", 10)

	N := new(big.Int).Mul(p, q)
	T := new(big.Int).Mul(new(big.Int).Sub(p, one), new(big.Int).Sub(q, one))

	e := new(big.Int)
	i := new(big.Int).Set(one)

	for ; i.Cmp(T) <= 0; i.Add(i, one) {
		if i.ProbablyPrime(20) {
			if new(big.Int).Mod(T, i).Cmp(zero) == 1 {
				e = i
				break
			}
		}
	}

	d := new(big.Int)
	j := new(big.Int).Set(one)

	for ; j.Cmp(max_try) <= 0; j.Add(j, one) {
		if new(big.Int).Mod(new(big.Int).Mul(e, j), T).Cmp(one) == 0 && j.Cmp(e) != 0 {
			d = j
			break
		}
	}

	fmt.Println(N.Uint64(), T.Uint64(), e.Uint64(), d.Uint64())

	t, _ := new(big.Int).SetString("112", 10)

	encrypted := new(big.Int).Exp(t, e, N)
	decrypted := new(big.Int).Exp(encrypted, d, N)
	fmt.Println(encrypted.Text(10))
	fmt.Println(decrypted.Text(10))
}
