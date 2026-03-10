package blockchain

import (
	"fmt"
	"strings"
)

func DisplayBlock(curBlock *Block) {
	width := 60
	border := strings.Repeat("=", width)
	side := "║"
	fmt.Println(border)
	if (curBlock.Data == "GENESIS BLOCK"){
		fmt.Printf("%s GENESIS BLOCK \n", side)
		fmt.Println(border)
		fmt.Printf("%s\n", strings.Repeat(" ", width/2) + "║")
		fmt.Printf("%s\n", strings.Repeat(" ", width/2) + "▼")
	}else {
		fmt.Printf("%s %-15s : %d\n", side, "BLOCK INDEX", curBlock.Index)
		fmt.Printf("%s %-15s : %s\n", side, "TIMESTAMP", curBlock.Timestamp)
		
		displayHash := curBlock.Hash[:10] + "..." + curBlock.Hash[len(curBlock.Hash)-10:]
		fmt.Printf("%s %-15s : %s\n", side, "CURRENT HASH", displayHash)
		
		fmt.Printf("%s %-15s : %s\n", side, "DATA", curBlock.Data)
		fmt.Printf("%s %-15s : %d\n", side, "NONCE", curBlock.Nonce)
		
		// Stats
		fmt.Printf("%s %-15s : %s (%d workers)\n", side, "MINING STATS", 
			curBlock.Statistic.TimeTakenToMine, curBlock.Statistic.Numworkers)
		
		fmt.Println(border)
		
		// The "Link" to the next block
		fmt.Printf("%s\n", strings.Repeat(" ", width/2) + "║")
		fmt.Printf("%s\n", strings.Repeat(" ", width/2) + "▼")

	}
}
