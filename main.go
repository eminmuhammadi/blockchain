package main

import (
	"crypto/sha256"
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Chain struct {
	Blocks []Block
}

type Block struct {
	Data     string
	Hash     string
	PrevHash string
}

func CreateChain() *Chain {
	return &Chain{
		Blocks: []Block{CreateGenesisBlock()},
	}
}

func RandInt() int {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	return r.Intn(math.MaxInt64)
}

func (block *Block) GenerateHash() {
	block.Hash = CalculateHash(block)
}

func CalculateHash(block *Block) string {
	hasher := sha256.New()

	hasher.Write([]byte(block.Data + block.PrevHash))

	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func CreateGenesisBlock() Block {
	block := Block{
		Data:     fmt.Sprintf("GENESIS_BLOCK_%d", RandInt()),
		PrevHash: "0",
	}

	block.GenerateHash()

	return block
}

func (chain *Chain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]

	newBlock := Block{
		Data:     data,
		PrevHash: prevBlock.Hash,
	}

	newBlock.GenerateHash()

	chain.Blocks = append(chain.Blocks, newBlock)
}

func (chain *Chain) IsValid() bool {
	for i := 0; i < len(chain.Blocks); i++ {
		// Genesis block
		if i == 0 {
			if chain.Blocks[i].Hash != CalculateHash(&chain.Blocks[i]) {
				return false
			}

			continue
		}

		prevBlock := chain.Blocks[i-1]
		currentBlock := chain.Blocks[i]

		if currentBlock.PrevHash != prevBlock.Hash {
			return false
		}

		if currentBlock.Hash != CalculateHash(&currentBlock) {
			return false
		}
	}

	return true
}
