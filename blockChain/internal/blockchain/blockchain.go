package blockchain

import (
	"fmt"
)

type BlockChain struct {
	Chain []*Block
	Genesis Block
	Difficulty int
}

func InitChain(difficulty int) *BlockChain{
	//Create the genesis block at the start of the chain
	genesis := createBlock("genesis Block", 0, "")
	//Create our actual chain struct
	bc := &BlockChain{
		Chain: []*Blocks{genesis},
		Difficulty: difficulty,
	}
	return bc

	fmt.Println("Initalizing Chain")
}

func ValidateChain(){

}
