// Copyright (c) 2017 George Tankersley. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// An implementation of the worst curve spec I could find, to demonstrate how
// the generic crypto/elliptic package works. NOT SAFE TO USE for many reasons.

package secp112r1

import (
	"crypto/elliptic"
	"math/big"
	"sync"
)

// A struct that satisfies the Curve interface.
type secp112r1Curve struct {
	*elliptic.CurveParams
}

// Define a set of named curve parameters.
var curveParams = &elliptic.CurveParams{Name: "secp112r1"}

// The concrete curve instance.
var curve = secp112r1Curve{curveParams}

// We use sync.Once to avoid repeatedly handing expensive big.Int operations
// when configuring the curve params.
var once sync.Once

// Configure the parameters of secp112r1, as specified in SEC2 2.2.1.
func initParams() {
	curveParams.P, _ = new(big.Int).SetString("DB7C2ABF62E35E668076BEAD208B", 16)
	curveParams.N, _ = new(big.Int).SetString("DB7C2ABF62E35E7628DFAC6561C5", 16)
	curveParams.B, _ = new(big.Int).SetString("659EF8BA043916EEDE8911702B22", 16)
	curveParams.Gx, _ = new(big.Int).SetString("09487239995A5EE76B55F9C2F098", 16)
	curveParams.Gy, _ = new(big.Int).SetString("A89CE5AF8724C0A23E0E0FF77500", 16)
	curveParams.BitSize = 112
}

// Curve returns a Curve that implements secp112r1.
func Curve() elliptic.Curve {
	once.Do(initParams)
	return curve
}
