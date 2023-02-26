# Blockchain

A simple blockchain implementation in Go.

## Installation

```bash
go get github.com/eminmuhammadi/blockchain
```

## Usage

```go
package main

import blockchain "github.com/eminmuhammadi/blockchain"

func main() {
	chain := blockchain.CreateChain()

	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")

	fmt.Println(chain.Blocks)

	println(chain.IsValid())
}
```

## Testing

```bash
go test
```