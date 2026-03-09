package blockchain

import (
)

type BlockChain struct {
	Chain []*Block
	Difficulty int
}

func InitChain(difficulty int) *BlockChain{
	//Create the genesis block at the start of the chain
	genesis := CreateBlock("genesis Block", 0, "")
	//Create our actual chain struct
	bc := &BlockChain{
		Chain: []*Block{genesis},
		Difficulty: difficulty,
	}
	return bc

}

func ValidateChain(){

}
