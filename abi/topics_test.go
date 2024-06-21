package abi

import (
	"math/big"
	"testing"

	"github.com/Adhara-Tech/ethgo-minimal"
	"github.com/stretchr/testify/assert"
)

func TestTopicEncoding(t *testing.T) {
	cases := []struct {
		Type string
		Val  interface{}
	}{
		{
			Type: "bool",
			Val:  true,
		},
		{
			Type: "bool",
			Val:  false,
		},
		{
			Type: "uint64",
			Val:  uint64(20),
		},
		{
			Type: "uint256",
			Val:  big.NewInt(1000000),
		},
		{
			Type: "address",
			Val:  ethgo.Address{0x1},
		},
	}

	for _, c := range cases {
		tt, err := NewType(c.Type)
		assert.NoError(t, err)

		res, err := EncodeTopic(tt, c.Val)
		assert.NoError(t, err)

		val, err := ParseTopic(tt, res)
		assert.NoError(t, err)

		assert.Equal(t, val, c.Val)
	}
}
