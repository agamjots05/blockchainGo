package main

import(
	"fmt"
	"bufio"
	"strconv"
	"os"
	"time"
	"blockChain/internal/blockchain"
	
	 
)

func main(){
	//create chain before getting user input
	blockchain.InitChain()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Block Chain CLI. Choose a number below")
	fmt.Println("1 --> Add a transaction to a block chain ")
	fmt.Println("2 --> Print your current block chain ")
	fmt.Println("3 --> Validate Current Chain ")
	for {
		if scanner.Scan(){
			input := scanner.Text()
			num, err := strconv.Atoi(input)
			if err != nil{
				fmt.Println("Invalid Input Type. Try Again")
				continue
			}
			if (num == 1){
				curTimeStamp := time.Now()
				formatTime := curTimeStamp.Format("2006-01-02 15:04:05")
				fmt.Println(formatTime)
				break
			}

			

			

		}

	}
}
