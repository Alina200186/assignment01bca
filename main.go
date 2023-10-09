package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Define a Block structure
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

// VerifyChain checks the integrity of the blockchain
func (bc *Blockchain) VerifyChain() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		// Recalculate the hash of the current block
		currentBlock.CreateHash()

		// Compare the recalculated hash with the stored hash
		if currentBlock.Hash != previousBlock.Hash {
			return false // Blockchain integrity compromised
		}
	}
	return true // Blockchain is intact
}

// CreateHash calculates the hash of a block
func (b *Block) CreateHash() {
	data := fmt.Sprintf("%s%d%s", b.Transaction, b.Nonce, b.PreviousHash)
	hashBytes := sha256.Sum256([]byte(data))
	b.Hash = hex.EncodeToString(hashBytes[:])
}

// NewBlock creates a new block with the provided transaction, nonce, and previous block
func NewBlock(transaction string, nonce int, previousBlock *Block) *Block {
	var previousHash string
	if previousBlock != nil {
		previousHash = previousBlock.Hash
	} else {
		// The first block (genesis block) has no previous hash
		previousHash = ""
	}

	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.CreateHash()
	return block
}

// Blockchain represents a collection of blocks
type Blockchain struct {
	Blocks []*Block
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(transaction string, nonce int) {
	// Get the previous block (if any)
	var previousBlock *Block
	if len(bc.Blocks) > 0 {
		previousBlock = bc.Blocks[len(bc.Blocks)-1]
	}

	// Create a new block with the previous block's hash
	newBlock := NewBlock(transaction, nonce, previousBlock)

	// Append the new block to the blockchain
	bc.Blocks = append(bc.Blocks, newBlock)
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
	// Call the Block's ChangeTransaction method to update the transaction
	blockToChange.ChangeTransaction(newTransaction)
}

func main() {
	// Create a blockchain with the genesis block
	blockchain := &Blockchain{}

	// Add the genesis block (first block) to the blockchain
	genesisBlock := NewBlock("First block Transaction", 0, nil)
	blockchain.Blocks = append(blockchain.Blocks, genesisBlock)

	// Add more blocks to the blockchain
	blockchain.AddBlock("bob to alice", 12345)
	blockchain.AddBlock("alice to bob", 67890)

	// Display all blocks in the blockchain
	blockchain.DisplayBlocks()

	// if len(blockchain.Blocks) > 1 {
	// 	blockToChange := blockchain.Blocks[1] // Get a reference to the second block
	// 	ChangeBlock(blockToChange, "Updated Transaction")
	// }
	// fmt.Println("After Modification:")
	// blockchain.DisplayBlocks()

	// Verify the integrity of the blockchain
	if blockchain.VerifyChain() {
		fmt.Println("Blockchain is intact.")
	} else {
		fmt.Println("Blockchain integrity compromised.")
	}

}
