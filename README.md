# SimpleBlockchainGo
## _Go Blockchain Implementation for Education_

[![](./pngwing.com%20%281%29.png)](https://go.dev/)

## Running the Go Blockchain Project

Follow these steps to run the Go blockchain project on your system:

1. **Install Go**: If you haven't already, you'll need to install the Go programming language on your system. You can download it from the official website: [Go Downloads](https://golang.org/dl/).

2. **Create a Project Directory**: Organize your project files by creating a directory. For example, you can create a directory named `blockchain_project`.

3. **Create a Go File**: Inside the project directory, create a Go file (e.g., `main.go`) and copy the entire code from the provided code snippet into this file.

4. **Open a Terminal/Command Prompt**: Open your terminal or command prompt.

5. **Navigate to the Project Directory**: Use the `cd` command to navigate to your project directory. For example, if you created the `blockchain_project` directory on your desktop, you would navigate to it like this:

   ```bash
   cd ~/Desktop/blockchain_project
6. **Compile and Run** run your Go blockchain project, use the following command:

      ```bash
      go run main.go
      ```

# Implementation:

## Block Structure

Each block is represented by a struct called `Block`. A block contains the following fields:

- `transactions`: A list of transactions stored in the block.
- `prevPointer`: A pointer to the previous block in the blockchain.
- `prevHash`: The hash of the previous block.
- `currentHash`: The hash of the current block.

## CalculateHash Function

The `CalculateHash` function computes the SHA-256 hash of a block. It takes a `Block` as input and returns the hash as a hexadecimal string.

## InsertBlock Function

The `InsertBlock` function adds a new block to the blockchain with the given transactions. It returns the updated head of the blockchain. The new block is linked to the previous block by setting the `prevPointer` and `prevHash` fields. The hash of the new block is calculated and set as `currentHash`.

## ChangeBlock Function

The `ChangeBlock` function searches for a specific transaction within the blockchain and updates it with a new value. It iterates through blocks in reverse order, starting from the head of the blockchain. When a transaction is found and updated, the block's hash is recalculated.

## ListBlocks Function

The `ListBlocks` function displays the contents of the blockchain. It iterates through blocks from the head to the genesis block and prints the block's hash, previous hash, and transactions.

## VerifyChain Function

The `VerifyChain` function checks the integrity of the blockchain by comparing block hashes. It iterates through blocks from the head to the genesis block. If a block's previous hash doesn't match the calculated hash of the previous block, the blockchain is considered compromised.


