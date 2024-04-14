package utils

import "math/big"

func GCD(a, b *big.Int) *big.Int {
	for b.Cmp(big.NewInt(0)) != 0 {
		t := b
		b = a.Mod(a, b)
		a = t
	}
	return a
}
