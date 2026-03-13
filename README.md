# GoLang Block Chain CLI
## This is a fully functional block chain CLI allowing users to create blocks and add it to their own chain!
This project was made to help me understand some of the basics on GoLang but also implement a topic I've been interested in since I first learned it, the blockchain! This is a very simple implementation that can be expanded to what ever domain you seek. This project does the following:
* Creates working blocks that must go through some proof of work in order to be formed
* Allows blocks to be chained together, granting data integrity 
* Allows users to track statistics about each block that has been added to the chain
* Users can see a clean visual representation of the chain to further solidify their understanding of the chain they created
* Has a CLI allowing users to create their own entries within the blockchain

## Here are some images of what it can look like :)
<img width="300" height="500" alt="Image" src="https://github.com/user-attachments/assets/ed69bfba-d0c7-4806-b79b-4b18cb3bc68c" />
<img width="870" height="546" alt="Image" src="https://github.com/user-attachments/assets/d38e96ca-8b23-479d-8d44-69ebc2e92bc1" />


## Installation Guide
1. Clone the repository
2. Cd into blockChain directory
3. run  `go run cmd/bc-cli/main.go`

## How to tweak this project
Since this is a very rudementary setup for a block-chain you can create a more specific version utilizing the current implementation. One example can include creating a specific cryptocurrency that allows transactions. The options are limitless so have fun with it.

