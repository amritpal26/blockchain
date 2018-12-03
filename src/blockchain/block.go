package blockchain

import(
	"crypto/sha256"
	"fmt"
	"math"
)

type Block struct {
	PrevHash   []byte
	Generation uint64
	Difficulty uint8
	Data       string
	Proof      uint64
	Hash       []byte
}

// Create new initial (generation 0) block.
func Initial(difficulty uint8) Block {

	block := new(Block)

	block.PrevHash = make([]byte, 32)
	block.Generation = 0
	block.Difficulty = difficulty
	block.Data = ""
	
	return *block
}

// Create new block to follow this block, with provided data.
func (prev_block Block) Next(data string) Block {

	block := new(Block)

	block.PrevHash = prev_block.Hash
	block.Generation = prev_block.Generation + 1
	block.Difficulty = prev_block.Difficulty
	block.Data = data

	return *block
}

// Calculate the block's hash.
func (blk Block) CalcHash() []byte {
	
	str := fmt.Sprintf("%x:%d:%d:%s:%d", blk.PrevHash, blk.Generation, blk.Difficulty, blk.Data, blk.Proof)

	hash := sha256.New()
	hash.Write([]byte(str))

	return hash.Sum(nil)
}

// Is this block's hash valid?
func (blk Block) ValidHash() bool {

	hash := blk.Hash
	difficulty := uint(blk.Difficulty)
	
	i := uint(0)
	for i = 0; i < uint(difficulty / 8); i++{
		testByte := byte(255)
		if(hash[31 - i] & testByte != 0x00){
			return false
		}	
	}
	
	difficulty -= i * 8
	testByte := byte(math.Pow(2, float64(difficulty)) - 1)
	if(hash[31 - i] & testByte != 0x00){
		return false
	}
	
	return true
}

// Set the proof-of-work and calculate the block's "true" hash.
func (blk *Block) SetProof(proof uint64) {
	blk.Proof = proof
	blk.Hash = blk.CalcHash()
}
