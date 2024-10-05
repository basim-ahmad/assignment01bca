package main

import (
	"assignment01bca"
	"fmt"
)

func main() {
	// Initialize the blockchain
	blockchain := &assignment01bca.Blockchain{}

	// Add blocks to the blockchain
	blockchain.AddToBlockchain(1, "", 100, "First Block Data", "0")
	blockchain.AddToBlockchain(2, "", 200, "Second Block Data", blockchain.Blocks[0].CurrentHash)
	blockchain.AddToBlockchain(3, "", 300, "Third Block Data", blockchain.Blocks[1].CurrentHash)

	// List all blocks
	fmt.Println("Blockchain Blocks:")
	assignment01bca.ListBlocks(blockchain)

	// Modify a block
	fmt.Println("\nChanging Block 1...")
	assignment01bca.ChangeBlock(blockchain, 1)

	// List blocks after change
	fmt.Println("\nBlockchain after changing Block 1:")
	assignment01bca.ListBlocks(blockchain)

	// Verify the blockchain integrity
	fmt.Println("\nVerifying Blockchain Integrity...")
	assignment01bca.VerifyChain(blockchain)
}
