package abi

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/Adhara-Tech/ethgo-minimal"
)

func TestEncodingStruct(t *testing.T) {
	typ := MustNewType("tuple(address aa, uint256 b)")

	type Obj struct {
		A ethgo.Address `abi:"aa"`
		B *big.Int
	}
	obj := Obj{
		A: ethgo.Address{0x1},
		B: big.NewInt(1),
	}

	encoded, err := typ.Encode(&obj)
	if err != nil {
		t.Fatal(err)
	}

	var obj2 Obj
	if err := typ.DecodeStruct(encoded, &obj2); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(obj, obj2) {
		t.Fatal("bad")
	}
}

func TestEncodingStruct_camelCase(t *testing.T) {
	typ := MustNewType("tuple(address aA, uint256 b)")

	type Obj struct {
		A ethgo.Address `abi:"aA"`
		B *big.Int
	}
	obj := Obj{
		A: ethgo.Address{0x1},
		B: big.NewInt(1),
	}

	encoded, err := typ.Encode(&obj)
	if err != nil {
		t.Fatal(err)
	}

	var obj2 Obj
	if err := typ.DecodeStruct(encoded, &obj2); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(obj, obj2) {
		t.Fatal("bad")
	}
}
