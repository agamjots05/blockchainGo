package main

import(
	"fmt"
	"bufio"
	"strconv"
	"os"
	"blockChain/internal/blockchain"
)

func main(){
	//create chain before getting user input
	curBlockChain := blockchain.InitChain(3)
	fmt.Println(curBlockChain.Difficulty)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Block Chain CLI. Choose a number below")
	fmt.Println("1 --> Add a transaction to a block chain ")
	fmt.Println("2 --> Print your current block chain ")
	fmt.Println("3 --> Validate Current Chain ")
	for {
		if scanner.Scan(){
			input := scanner.Text()
			num, err := strconv.Atoi(input)
			// Invalid User input type
			if err != nil{
				fmt.Println("Invalid Input Type. Try Again")
				continue
			}
			//First Case:
			if (num == 1){
				fmt.Println("Please Enter Your Transaction Details")
				if scanner.Scan(){
					prevBlock := curBlockChain.Chain[len(curBlockChain.Chain) - 1]
					data := scanner.Text()
					newBlock := blockchain.CreateBlock(data, prevBlock.Index + 1, prevBlock.Hash)
					
					//Now we must generate a valid hash via the proof-of-work algorithm
					blockchain.Proofofwork(newBlock, curBlockChain.Difficulty)

					
				}
				break
			}

		}

	}
}
