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

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Block Chain CLI. Choose a number below")
	for {
		fmt.Println("1 --> Add a transaction to a block chain ")
		fmt.Println("2 --> Print a specific blocks statistics")
		fmt.Println("3 --> Print your current block chain ")
		fmt.Println("4 --> Validate Current Chain ")

		if scanner.Scan(){
			input := scanner.Text()
			num, err := strconv.Atoi(input)
			// Invalid User input type
			if err != nil{
	 			fmt.Println("\n Invalid Input Type. Try Again")
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
					fmt.Println("--- Mining for a valid block ---")
					blockchain.Proofofwork(newBlock, curBlockChain.Difficulty)

					//Append valid block to chain
					curBlockChain.Chain = append(curBlockChain.Chain, newBlock)
					fmt.Println("--- Block has been added to blockchain ---")

				}
				continue	
			}
			if (num == 2){
				for {
					fmt.Println("\n Please choose a block from the chain. E.g., (2 -> Corresponds to the second block in the chain)")
					fmt.Println("Enter 0 to go back to main selection")
					if scanner.Scan(){
						input := scanner.Text()
						num, err := strconv.Atoi(input)
						if err != nil{
							fmt.Println("Invalid datatype. Try Again")
							continue
						}
						if (num == 0){
							break
						} else {
							// validate whether block even exists within chain
							if num < 1 || num > len(curBlockChain.Chain) {
								fmt.Printf("Error: Block #%d does not exist. (Max: %d)\n", num, len(curBlockChain.Chain))
								continue
							}

							//Find block via index
							foundBlock := curBlockChain.Chain[num-1]

							fmt.Printf("--- BLOCK %d INFORMATION BELOW --- \n", num)
							fmt.Printf("Block Data: %s\n", foundBlock.Data)
							fmt.Printf("Total of %d miners were used to mine this block \n", foundBlock.Statistic.Numworkers)
							fmt.Printf("Took %s to mine the block \n", foundBlock.Statistic.TimeTakenToMine)
							fmt.Printf("Checked a total of %d hashes to find a valid one \n", foundBlock.Statistic.TotalHashesChecked)
							continue
						}
					}
				}
				continue
			}

		}

	}
}
