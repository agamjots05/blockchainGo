package blockchain

import (
	"time"
)
type BlockStatistic struct {
	Numworkers int
	TotalHashesChecked int
	TimeTakenToMine string
}

type Block struct {
	Index int
	PrevHash string
	Timestamp string
	Hash string
	Data string
	Nonce int
	Statistic BlockStatistic
}


func CreateBlock(data string, index int, prevHash string) *Block{
	curTimeStamp := time.Now()
	formatTime := curTimeStamp.Format("2006-01-02 15:04:05")
	block := &Block{
		Index: index,
		PrevHash: prevHash,
		Timestamp: formatTime,
		//We don't set hash. et during proof of work
		Data: data,
		Statistic: BlockStatistic{
			Numworkers: 0,
			TotalHashesChecked: 0,
			TimeTakenToMine: "",
		},
		Nonce: 0,
	}
	return block
}
