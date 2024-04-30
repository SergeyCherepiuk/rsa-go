package rsa

import (
	"crypto/rand"
	"math/big"

	"github.com/SergeyCherepiuk/rsa-go/internal/ascii"
)

type PublicKey struct {
	E, N *big.Int
}

type PrivateKey struct {
	D, N *big.Int
}

var (
	p, _ = rand.Prime(rand.Reader, 2112)
	q, _ = rand.Prime(rand.Reader, 1984)
)

func GeneratePrivateKey() *PrivateKey {
	n := new(big.Int).Mul(p, q)

	pMinusOne := new(big.Int).Sub(p, big.NewInt(1))
	qMinusOne := new(big.Int).Sub(q, big.NewInt(1))

	f := new(big.Int).Mul(pMinusOne, qMinusOne)
	e := big.NewInt(65537)
	d := new(big.Int).Exp(e, big.NewInt(-1), f)

	return &PrivateKey{d, n}
}

func (k *PrivateKey) PublicKey() *PublicKey {
	e := big.NewInt(65537)
	return &PublicKey{e, k.N}
}

func Encode(message []byte, pub *PublicKey) *big.Int {
	codes := string(ascii.Encode(message))
	m, _ := new(big.Int).SetString(codes, 10)
	return new(big.Int).Exp(m, pub.E, pub.N)
}

func Decode(e *big.Int, priv *PrivateKey) []byte {
	m := new(big.Int).Exp(e, priv.D, priv.N).String()
	return ascii.Decode([]byte(m))
}
