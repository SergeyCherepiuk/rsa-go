package rsa

import (
	"math/big"

	"github.com/SergeyCherepiuk/rsa-go/internal/utils"
)

type PublicKey struct {
	E, N *big.Int
}

type PrivateKey struct {
	D, N *big.Int
}

var (
	p = new(big.Int).SetUint64(18446744073709551113)
	q = new(big.Int).SetUint64(18446744073709551191)
)

func GeneratePrivateKey() *PrivateKey {
	n := p.Mul(p, q)

	pMinusOne := p.Sub(p, big.NewInt(1))
	qMinusOne := q.Sub(q, big.NewInt(1))

	f := pMinusOne.Mul(pMinusOne, qMinusOne)

	e := big.NewInt(1)
	for e.Cmp(f) == -1 {
		if utils.GCD(e, f) == big.NewInt(1) {
			return &PrivateKey{e, n}
		}

		e = e.Add(e, big.NewInt(1))
	}

	return nil
}
