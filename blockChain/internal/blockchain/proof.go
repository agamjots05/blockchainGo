package blockchain

import (
	"fmt"
	"runtime"
	"sync"
	"encoding/hex"
	"crypto/sha256"
	"strings"
	"strconv"
)

func findHash(curBlock *Block, curNonce int) string {
	data := strconv.Itoa(curBlock.Index) + curBlock.Timestamp + curBlock.Data + curBlock.PrevHash + strconv.Itoa(curNonce)
	hasher := sha256.New()
	hasher.Write([]byte(data))
	hashedData := hasher.Sum(nil)
	strHashedData := hex.EncodeToString(hashedData)
	return strHashedData
}
func validateHash(curHash string, difficulty int) bool{
	//In order to validate our hash and make sure it's difficult enough, we check whether the hash contains a prefix of len(difficulty) 0's
	//e.g. Hash of difficulty 3 --> 000abjhkgjh2jgjn4l...
	targetPrefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(curHash, targetPrefix)
}

func Proofofwork(curBlock *Block, difficulty int){
	/*
	We will be creating go routines based on the number of cores curCPU ha.
	Each worker will be working on their set of nonces each trying to find a valid hash.
	After one worker has found a valid hash, they will populate the foundNonce Channel
	*/
	numWorkers := runtime.NumCPU()
	validNonce := make(chan int, 1)
	validHash := make(chan string, 1)
	stopCond := make(chan struct{})
	var wg sync.WaitGroup

	for i:= 0; i < numWorkers; i++{
		wg.Add(1)
		go func(curNonce int){
			defer wg.Done()
			for {
				select {
				case <-stopCond:
					return 
				default:
					//Determine whether current nonce is a valid one
					/*
					If we find a valid hash -> Close our stopCond chan to 
					signal other workers someone has found a valid hash.

					Then pass our valid hash into our channel to be read later

					*/
					curHash := findHash(curBlock, curNonce)
					if (validateHash(curHash, difficulty)){
						fmt.Println("Miner has found a valid nonce")
						validNonce <- curNonce
						validHash <- curHash
						return
					}
					curNonce += numWorkers
				}
			}
		}(i)
	}

	nonce := <- validNonce
	hash := <- validHash
	close(stopCond)

	wg.Wait()
	curBlock.Hash = hash
	curBlock.Nonce = nonce
	return







}
