package blockchain

import(
	"bytes"
	"encoding/hex"
)

type Blockchain struct {
	Chain []Block
}

func (chain *Blockchain) Add(blk Block) {
	// You can remove the panic() here if you wish.
	if !blk.ValidHash() {
		panic("Trying to insert invalid block.")
	}else {
		chain.Chain = append(chain.Chain, blk)
	}
}

func (chain Blockchain) IsValid() bool {
	ch := chain.Chain
	len := len(ch)

	if (len > 0) {

		initial := ch[0]
		prevHash := initial.Hash
		prevDifficulty := initial.Difficulty
		prevGeneration := initial.Generation

		if hex.EncodeToString(initial.PrevHash) != "0000000000000000000000000000000000000000000000000000000000000000" {
			return false
		}
		if (!initial.ValidHash() || initial.Generation != 0){
			return false
		}
		if (!bytes.Equal(initial.CalcHash(), initial.Hash)){
			return false
		}

		for i := 1; i < len; i++{
			blk := ch[i]
			if (!bytes.Equal(blk.PrevHash, prevHash) || blk.Generation != prevGeneration + 1){
				return false
			}else if (blk.Difficulty != prevDifficulty || !bytes.Equal(blk.CalcHash(), blk.Hash) || !blk.ValidHash() ){
				return false
			}else{
				prevHash = blk.Hash
				prevDifficulty = blk.Difficulty
				prevGeneration = blk.Generation
			}
		}
	}

	return true
}

