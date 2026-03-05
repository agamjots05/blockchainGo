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

func CreateBlock(data string, index int, prevHash string) *Block{
	curTimeStamp := time.Now()
	formatTime := curTimeStamp.Format("2006-01-02 15:04:05")
	block := &Block{
		Index: index,
		PrevHash: prevHash,
		Timestamp: formatTime,
		//We don't set hash. We set during proof of work
		Data: data,
		Nonce: 0,
	}
	return block
}
