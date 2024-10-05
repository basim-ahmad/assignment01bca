
# Blockchain Implementation in Go

This repository contains a simple blockchain implementation in Go. It includes two main parts: the blockchain package (`assignment01bca`) and an executable that demonstrates the usage of this package (`main.go`).

## Structure

### Package: `assignment01bca`

This package defines the core components of the blockchain, including the structure of a block, the blockchain itself, and functions to manipulate and verify the blockchain.

#### Core Components

- **Block**: Defines the structure for individual blocks in the blockchain.
- **Blockchain**: Manages a series of blocks linked together to form the blockchain.
- **Functions**:
  - `AddToBlockchain`: Adds a new block to the blockchain.
  - `ListBlocks`: Lists all blocks in the blockchain.
  - `ChangeBlock`: Modifies a block at a given index.
  - `VerifyChain`: Checks the integrity of the blockchain.

### Executable: `main.go`

Demonstrates how to:
- Initialize a new blockchain.
- Add blocks with transaction data.
- List all blocks.
- Modify a block to see the impact on the blockchain integrity.
- Verify the integrity of the entire blockchain.

## Usage

To run this project, ensure you have Go installed on your machine. Clone this repository, navigate to the directory containing `main.go`, and run the following command:

```bash
go run main.go
```

This will execute the blockchain demonstration, showing how blocks are added, listed, modified, and verified.

## Development

This project is structured to allow easy enhancements and modifications to the blockchain protocol. You can modify the block structure, add new features to the blockchain management system, or improve the existing cryptographic functions.

## Contribution

Contributions are welcome. Please fork the repository, make your changes, and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
