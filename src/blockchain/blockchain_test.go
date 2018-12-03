package blockchain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
	"encoding/hex"
)

// TODO: some useful tests of Blocks
func TestInitialBlock(t *testing.T) {
	block := Initial(16)
	block.SetProof(67)
	assert.False(t, block.ValidHash(), "Initial block should should be invalid when proof = 0")

	block.SetProof(56231)
	assert.True(t, block.ValidHash(), "Initial block should should be valid when proof = 56231")
}

func TestNextBlock(t *testing.T) {
	b0 := Initial(19)
	b0.SetProof(87745)
	b1 := b0.Next("hash example 1234")
	b1.SetProof(1407891)

	assert.True(t, b1.ValidHash(), "Block should should be valid when proof = 1407891")

	b1.SetProof(346082)
	assert.False(t, b1.ValidHash(), "Block should should be invalid when proof = 346082")
}

func TestMine1 (t *testing.T) {
	b0 := Initial(7)
	b0.Mine(2)
	assert.Equal(t, b0.Proof, uint64(385), "b0.Proof is invalid")
	assert.Equal(t, hex.EncodeToString(b0.Hash), "379bf2fb1a558872f09442a45e300e72f00f03f2c6f4dd29971f67ea4f3d5300", "b0.Hash is invalid")
	

	b1 := b0.Next("this is an interesting message")
	b1.Mine(2)
	assert.Equal(t, b1.Proof, uint64(20), "b0.Proof is invalid")
	assert.Equal(t, hex.EncodeToString(b1.Hash), "4a1c722d8021346fa2f440d7f0bbaa585e632f68fd20fed812fc944613b92500", "b0.Hash is invalid")

	b2 := b1.Next("this is not interesting")
	b2.Mine(2)
	assert.Equal(t, b2.Proof, uint64(40), "b0.Proof is invalid")
	assert.Equal(t, hex.EncodeToString(b2.Hash), "ba2f9bf0f9ec629db726f1a5fe7312eb76270459e3f5bfdc4e213df9e47cd380", "b0.Hash is invalid")

	fmt.Println(b0.Proof, hex.EncodeToString(b0.Hash))
	fmt.Println(b1.Proof, hex.EncodeToString(b1.Hash))
	fmt.Println(b2.Proof, hex.EncodeToString(b2.Hash))	
}

func TestMine2 (t *testing.T) {
	b0 := Initial(20)
	b0.Mine(1)
	assert.Equal(t, b0.Proof, uint64(1209938), "b0.Proof is invalid")
	assert.Equal(t, hex.EncodeToString(b0.Hash), "19e2d3b3f0e2ebda3891979d76f957a5d51e1ba0b43f4296d8fb37c470600000", "b0.Hash is invalid")
	

	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	assert.Equal(t, b1.Proof, uint64(989099), "b0.Proof is invalid")
	assert.Equal(t, hex.EncodeToString(b1.Hash), "a42b7e319ee2dee845f1eb842c31dac60a94c04432319638ec1b9f989d000000", "b0.Hash is invalid")

	b2 := b1.Next("this is not interesting")
	b2.Mine(1)
	assert.Equal(t, b2.Proof, uint64(1017262), "b0.Proof is invalid")
	assert.Equal(t, hex.EncodeToString(b2.Hash), "6c589f7a3d2df217fdb39cd969006bc8651a0a3251ffb50470cbc9a0e4d00000", "b0.Hash is invalid")

	fmt.Println(b0.Proof, hex.EncodeToString(b0.Hash))
	fmt.Println(b1.Proof, hex.EncodeToString(b1.Hash))
	fmt.Println(b2.Proof, hex.EncodeToString(b2.Hash))	
}

func TestBlockChainAdditionFail(t *testing.T) {
	blockchain := new(Blockchain)

	b0 := Initial(8)
	b0.SetProof(67)

	assert.Panics(t, func(){blockchain.Add(b0)}, "Insert a valid initial block in the chain")
}

func TestBlockChainAdditionFail2(t *testing.T) {
	blockchain := new(Blockchain)

	b0 := Initial(16)
	b0.SetProof(56231)

	assert.NotPanics(t, func(){blockchain.Add(b0)}, "Insert a valid initial block in the chain")

	b1 := b0.Next("This is the second message")
	assert.Panics(t, func(){blockchain.Add(b1)}, "Insert a valid second block in the chain")
}

func TestBlockChainAdditionPass(t *testing.T) {
	blockchain := new(Blockchain)

	b0 := Initial(16)
	b0.SetProof(56231)

	assert.NotPanics(t, func(){blockchain.Add(b0)}, "Insert a valid initial block in the chain")

	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	assert.NotPanics(t, func(){blockchain.Add(b1)}, "Insert a valid second block in the chain")
}

func TestBlockchainValidity(t *testing.T){
	blockchain := new(Blockchain)

	b0 := Initial(16)
	b0.Mine(1)
	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	b2 := b1.Next("this is another interesting message")
	b2.Mine(1)

	blockchain.Add(b0)
	blockchain.Add(b1)
	blockchain.Add(b2)

	assert.True(t, blockchain.IsValid(), "Blockchain isn't valid")
}

