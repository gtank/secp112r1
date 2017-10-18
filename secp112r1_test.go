package secp112r1

import (
	"crypto/elliptic"
	"encoding/hex"
	"math/big"
	"testing"
)

// Does the curve initialize at all?
func TestCurveInit(t *testing.T) {
	secp112 := Curve()
	params := secp112.Params()

	if params.Name != "secp112r1" {
		t.Errorf("expected secp112r1, got %s", params.Name)
	}

	if params.BitSize != 112 {
		t.Errorf("expected 112-bit curve, got %d", params.BitSize)
	}
}

// Is the base point on the curve?
func TestOnCurve(t *testing.T) {
	secp112 := Curve()
	if !secp112.IsOnCurve(secp112.Params().Gx, secp112.Params().Gy) {
		t.Errorf("FAIL: base point is on curve")
	}
}

// Is an off-curve point detected?
func TestOffCurve(t *testing.T) {
	secp112 := Curve()
	x, y := new(big.Int).SetInt64(0), new(big.Int).SetInt64(1)

	if secp112.IsOnCurve(x, y) {
		t.Errorf("FAIL: (0,1) is not on curve")
	}
}

func TestMarshal(t *testing.T) {
	secp112 := Curve()
	basePoint := "0409487239995A5EE76B55F9C2F098A89CE5AF8724C0A23E0E0FF77500"

	rawG, err := hex.DecodeString(basePoint)
	if err != nil {
		t.Fatalf("Failed to decode base point: %v", err)
	}

	x, y := elliptic.Unmarshal(secp112, rawG)
	if x == nil {
		t.Errorf("FAIL: could not unmarshal base point")
	}

	if x.Cmp(secp112.Params().Gx) != 0 || y.Cmp(secp112.Params().Gy) != 0 {
		t.Errorf("FAIL: unmarshaled base point was wrong, got (0x%X, 0x%X)", x.Int64(), y.Int64())
	}

	// (0,1) in uncompressed SEC1 format
	offCurvePoint := "0400000000000000000000000000000000000000000000000000000001"
	rawPoint, err := hex.DecodeString(offCurvePoint)
	if err != nil {
		t.Fatalf("Failed to decode off-curve point: %v", err)
	}

	x, y = elliptic.Unmarshal(secp112, rawPoint)
	if x != nil {
		t.Errorf("FAIL: successfully unmarshaled an off-curve point")
	}
}
