package assignment01bca

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

func calHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

type Block struct {
	TransactionID int
	CurrentHash   string
	Nonce         int
	Data          string
	PreviousHash  string
}

type Blockchain struct {
	Blocks []*Block
}

func NewBlock(TransactionID int, CurrentHash string, Nonce int, Data string, PreviousHash string) *Block {
	block := &Block{
		TransactionID: TransactionID,
		CurrentHash:   CurrentHash,
		Nonce:         Nonce,
		Data:          Data,
		PreviousHash:  PreviousHash,
	}
	return block
}

func (bc *Blockchain) AddToBlockchain(TransactionID int, CurrentHash string, Nonce int, Data string, PreviousHash string) *Block {
	blockData := strconv.Itoa(TransactionID) + CurrentHash + strconv.Itoa(Nonce) + Data + PreviousHash
	calculatedHash := CalculateHash(blockData)
	newBlock := NewBlock(TransactionID, calculatedHash, Nonce, Data, PreviousHash)
	bc.Blocks = append(bc.Blocks, newBlock)
	return newBlock
}

func ListBlocks(bc *Blockchain) {
	for i, block := range bc.Blocks {
		fmt.Printf("\n%s Block %d %s", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Printf("\nTransaction ID: %d", block.TransactionID)
		fmt.Printf("\nNonce: %d", block.Nonce)
		fmt.Printf("\nData: %s", block.Data)
		fmt.Printf("\nPrevious Hash: %s", block.PreviousHash)
		fmt.Printf("\nCurrent Hash: %s", block.CurrentHash)
	}
}

func CalculateHash(input string) string {
	return calHash(input)
}

func ChangeBlock(bc *Blockchain, blockIndex int) {
	bc.Blocks[blockIndex].TransactionID = blockIndex
	block := bc.Blocks[blockIndex]
	blockData := strconv.Itoa(block.TransactionID) + block.CurrentHash + strconv.Itoa(block.Nonce) + block.Data + block.PreviousHash
	newHash := CalculateHash(blockData)
	block.CurrentHash = newHash
}

func VerifyChain(bc *Blockchain) {
	for i := 1; i < len(bc.Blocks); i++ {
		if bc.Blocks[i].PreviousHash != bc.Blocks[i-1].CurrentHash {
			fmt.Printf("\n%s[-] WARNING: Blockchain integrity issue detected at Block %d [-]%s", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
			fmt.Printf("\nBlock %d requires verification or re-mining.", i)
		}
	}
}
