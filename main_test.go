package main

import "testing"

func TestBlock(t *testing.T) {
	chain := CreateChain()

	if len(chain.Blocks) != 1 {
		t.Errorf("Expected chain length of 1, got %d", len(chain.Blocks))
	}

	chain.AddBlock("Hello World")

	if len(chain.Blocks) != 2 {
		t.Errorf("Expected chain length of 2, got %d", len(chain.Blocks))
	}

	if chain.Blocks[1].Data != "Hello World" {
		t.Errorf("Expected data of 'Hello World', got %s", chain.Blocks[1].Data)
	}

	if chain.Blocks[1].PrevHash != chain.Blocks[0].Hash {
		t.Errorf("Expected previous hash of %s, got %s", chain.Blocks[0].Hash, chain.Blocks[1].PrevHash)
	}
}

func TestChainValidity(t *testing.T) {
	chain := CreateChain()

	chain.AddBlock("Hello World")
	chain.AddBlock("Hello World 2")

	if !chain.IsValid() {
		t.Errorf("Expected chain to be valid")
	}

	chain.Blocks[1].Data = "Hello World 3"

	if chain.IsValid() {
		t.Errorf("Expected chain to be invalid")
	}
}

func TestChainValidityWithTamperedBlock(t *testing.T) {
	chain := CreateChain()

	chain.AddBlock("Hello World")
	chain.AddBlock("Hello World 2")

	if !chain.IsValid() {
		t.Errorf("Expected chain to be valid")
	}

	chain.Blocks[1].Data = "Hello World 3"
	chain.Blocks[1].GenerateHash()

	if chain.IsValid() {
		t.Errorf("Expected chain to be invalid")
	}
}

func TestChainValidityWithTamperedPreviousHash(t *testing.T) {
	chain := CreateChain()

	chain.AddBlock("Hello World")
	chain.AddBlock("Hello World 2")

	if !chain.IsValid() {
		t.Errorf("Expected chain to be valid")
	}

	chain.Blocks[1].PrevHash = "0"

	if chain.IsValid() {
		t.Errorf("Expected chain to be invalid")
	}
}

func TestChainValidityWithTamperedHash(t *testing.T) {
	chain := CreateChain()

	chain.AddBlock("Hello World")
	chain.AddBlock("Hello World 2")

	if !chain.IsValid() {
		t.Errorf("Expected chain to be valid")
	}

	chain.Blocks[1].Hash = "0"

	if chain.IsValid() {
		t.Errorf("Expected chain to be invalid")
	}
}

func TestChainValidityWithTamperedGenesisBlock(t *testing.T) {
	chain := CreateChain()

	chain.AddBlock("Hello World")
	chain.AddBlock("Hello World 2")

	if !chain.IsValid() {
		t.Errorf("Expected chain to be valid")
	}

	chain.Blocks[0].Data = "Hello World 3"

	if chain.IsValid() {
		t.Errorf("Expected chain to be invalid")
	}
}

func TestChainValidityWithTamperedGenesisBlockHash(t *testing.T) {
	chain := CreateChain()

	chain.AddBlock("Hello World")
	chain.AddBlock("Hello World 2")

	if !chain.IsValid() {
		t.Errorf("Expected chain to be valid")
	}

	chain.Blocks[0].Hash = "0"

	if chain.IsValid() {
		t.Errorf("Expected chain to be invalid")
	}
}
