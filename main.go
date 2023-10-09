package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Blockchain represents a collection of blocks
type Blockchain struct {
	Blocks []*Block
}

// Define a Block structure
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

// NewBlock creates a new block with the provided transaction, nonce, and previous hash.
func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.Hash = block.CreateHash()
	return block
}

// CreateHash calculates the hash of a block and returns it
func (b *Block) CreateHash() string {
	data := fmt.Sprintf("%s%d%s", b.Transaction, b.Nonce, b.PreviousHash)
	hashBytes := sha256.Sum256([]byte(data)) //sha256 fun to calculate hash
	return hex.EncodeToString(hashBytes[:])
}

// DisplayBlocks prints all blocks in the blockchain
func (bc *Blockchain) DisplayBlocks() {
	for i, block := range bc.Blocks {
		fmt.Printf("Block #%d:\n", i+1)
		fmt.Printf("  Transaction: %s\n", block.Transaction)
		fmt.Printf("  Nonce: %d\n", block.Nonce)
		fmt.Printf("  Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("  Block Hash: %s\n", block.Hash)
		fmt.Println()
	}
}

// Add a method to the Block struct to change the transaction
func (b *Block) ChangeTransaction(newTransaction string) {
	b.Transaction = newTransaction
	b.CreateHash() // Recalculate the hash after changing the transaction
}

// Define a ChangeBlock function to change the transaction of a block
func ChangeBlock(blockToChange *Block, newTransaction string) {
	blockToChange.ChangeTransaction(newTransaction)
}

// VerifyChain checks the integrity of the blockchain
func (bc *Blockchain) VerifyChain() bool {

	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		// Verify the current block's hash
		if currentBlock.Hash != currentBlock.CreateHash() {
			return false
		}

		// Verify that the previous hash in the current block matches the hash of the previous block
		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}

func main() {
	// Create a blockchain with the genesis block
	blockchain := &Blockchain{}

	// Add blocks to the blockchain
	genesisBlock := NewBlock("First block Transaction", 0, "")
	block1 := NewBlock("bob to alice", 12345, genesisBlock.Hash)
	block2 := NewBlock("Alina to Hadia", 45678, block1.Hash)
	blockchain.Blocks = append(blockchain.Blocks, genesisBlock)
	blockchain.Blocks = append(blockchain.Blocks, block1)
	blockchain.Blocks = append(blockchain.Blocks, block2)
	// display the blockchain in a nice formate
	blockchain.DisplayBlocks()
	//changing the block transaction
	ChangeBlock(block1, "Maryum to Maria")
	fmt.Println("Modified Block chain")
	blockchain.DisplayBlocks()

	//Verify the integrity of the blockchain
	if blockchain.VerifyChain() {
		fmt.Println("Blockchain is intact.")
	} else {
		fmt.Println("Blockchain integrity compromised.")
	}

}
