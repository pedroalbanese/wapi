// Parameters for the OSCCA WAPIP192v1 Elliptic curve
package wapip192v1

import (
	"crypto/elliptic"
	"math/big"
	"sync"
)

var initonce sync.Once
var wapip192v1 *elliptic.CurveParams

func initP192() {
	wapip192v1 = new(elliptic.CurveParams)
	wapip192v1.P, _ = new(big.Int).SetString("bdb6f4fe3e8b1d9e0da8c0d46f4c318cefe4afe3b6b8551f", 16)
	wapip192v1.N, _ = new(big.Int).SetString("bdb6f4fe3e8b1d9e0da8c0d40fc962195dfae76f56564677", 16)
	wapip192v1.B, _ = new(big.Int).SetString("1854bebdc31b21b7aefc80ab0ecd10d5b1b3308e6dbf11c1", 16)
	wapip192v1.Gx, _ = new(big.Int).SetString("4ad5f7048de709ad51236de65e4d4b482c836dc6e4106640", 16)
	wapip192v1.Gy, _ = new(big.Int).SetString("02bb3a02d4aaadacae24817a4ca3a1b014b5270432db27d2", 16)
	wapip192v1.BitSize = 192
}

func P192() elliptic.Curve {
	initonce.Do(initP192)
	return wapip192v1
}
