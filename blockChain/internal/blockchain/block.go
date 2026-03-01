package blockchain

import (
	"time"
)

type Block struct {
	Index int
	PrevHash string
	Timestamp string
	Hash string
	Data string
	Nonce int
}

func createBlock(data string, index int, prevHash string) *Block{
	block := *Block{
		Index: index,
		PrevHash: prevHash,
		Timestamp: time.Now(),
		//We don't set hash. We set during proof of work
		Data: data,
		Nonce: 0,
	}
	return block





}
