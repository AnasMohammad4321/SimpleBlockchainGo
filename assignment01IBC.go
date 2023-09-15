/*
	Mohammad Anas
	20L-1289
	BDS-7A
	Assignment # 1
	Blockchain & Cryptocurrency
*/

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

// Block represents a single block in the blockchain.
type Block struct {
	transactions []string // List of transactions stored in the block.
	prevPointer  *Block   // Pointer to the previous block in the chain.
	prevHash     string   // Hash of the previous block.
	currentHash  string   // Hash of the current block.
}

// CalculateHash calculates and returns the SHA-256 hash of a block.
func CalculateHash(inputBlock *Block) string {
	if inputBlock == nil {
		return "" // Return an empty string if inputBlock is nil.
	}
	prevHash := inputBlock.prevHash
	transactionsData := strings.Join(inputBlock.transactions, "") // Concatenate all transactions into a single string.
	hashed := sha256.New()                                        // Initialize the SHA-256 hasher.
	hashed.Write([]byte(prevHash + transactionsData))             // Write the data to the hasher.
	hashBytes := hashed.Sum(nil)                                  // Get the hash bytes.
	return hex.EncodeToString(hashBytes)                          // Convert hashBytes to a hexadecimal string.
}

// InsertBlock inserts a new block with the given transactions into the blockchain and returns the updated head of the blockchain.
func InsertBlock(transactionsToInsert []string, chainHead *Block) *Block {
	// Create a new block with the provided transactions.
	newBlock := &Block{
		transactions: transactionsToInsert,
		prevPointer:  chainHead,
		prevHash:     chainHead.currentHash, // Set the previous hash to the current head's hash.
		currentHash:  "",                    // Will be calculated when the block is inserted.
	}
	newBlock.currentHash = CalculateHash(newBlock) // Calculate the hash of the new block.
	return newBlock
}

// ChangeBlock searches for a specific transaction in the blockchain and updates it with a new value.
func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
	currentBlock := chainHead
	for currentBlock != nil {
		transactionUpdated := false
		for i, transaction := range currentBlock.transactions {
			if transaction == oldTrans {
				currentBlock.transactions[i] = newTrans // Update the transaction.
				transactionUpdated = true
			}
		}
		if transactionUpdated {
			currentBlock.currentHash = CalculateHash(currentBlock) // Update the hash.
			return
		}
		currentBlock = currentBlock.prevPointer // Move to the previous block.
	}
}

// ListBlocks displays the contents of the blockchain.
func ListBlocks(chainHead *Block) {
	if chainHead == nil {
		fmt.Println("The blockchain is empty.")
		return
	}
	currentBlock := chainHead
	fmt.Println("List of Blocks")
	for currentBlock != nil {
		fmt.Printf("Block Hash: %s\n", currentBlock.currentHash)
		fmt.Printf("Previous Hash: %s\n", currentBlock.prevHash)
		fmt.Println("Transactions:")
		for _, tx := range currentBlock.transactions {
			fmt.Println("  ", tx)
		}
		fmt.Println()
		currentBlock = currentBlock.prevPointer
	}
}

// VerifyChain checks the integrity of the blockchain by comparing block hashes.
func VerifyChain(chainHead *Block) {
	currentBlock := chainHead

	for currentBlock.prevHash != "" {
		// Skip the comparison for the genesis block (empty previous hash).
		if currentBlock.prevHash != CalculateHash(currentBlock.prevPointer) && currentBlock.prevHash != "" {
			fmt.Println("Block chain is compromised")
			return
		}
		currentBlock = currentBlock.prevPointer
	}

	fmt.Println("Block chain is unchanged")
}

func main() {
	// Initializing the chain with a valid genesis block
	genesisBlock := &Block{
		transactions: []string{"Transaction 0"},
		prevPointer:  nil,
		prevHash:     "",
		currentHash:  "",
	}
	genesisBlock.currentHash = CalculateHash(genesisBlock)

	chainHead := InsertBlock([]string{"Transaction 1", "Transaction 2"}, genesisBlock)
	chainHead = InsertBlock([]string{"Transaction 3"}, chainHead)

	// Uncomment these lines to list the blocks and verify the chain
	// ListBlocks(chainHead)
	// VerifyChain(chainHead)

	transactionToChange := "Transaction 2"
	newTransaction := "Updated Transaction 2"
	ChangeBlock(transactionToChange, newTransaction, chainHead)

	ListBlocks(chainHead)
	VerifyChain(chainHead)
}
