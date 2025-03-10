package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

type Block struct {
	nonce        int
	prevHash     [32]byte
	timestamp    int64
	transactions []string
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) CreateBlock(nonce int, prevHash [32]byte) *Block {
	b := NewBlock(nonce, prevHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp: 		%d\n", b.timestamp)
	fmt.Printf("nonce:         	%d\n", b.nonce)
	fmt.Printf("prev_Hash: 		%x\n", b.prevHash)
	fmt.Printf("transactions: 	%s\n", b.transactions)
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 80))
}

func NewBlock(nonce int, prevHash [32]byte) *Block {
	/*
		b:= new (Block)
		b.timestamp = time.Now().UnixNano()
		b.nonce = nonce
		b.prevHash = prevHash
		return b
	*/
	return &Block{
		timestamp: time.Now().UnixNano(),
		nonce:     nonce,
		prevHash:  prevHash,
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nonce        int      `json:"nonce"`
		PrevHash     [32]byte `json:"prev_hash"`
		Timestamp    int64    `json:"timestamp"`
		Transactions []string `json:"transactions"`
	}{
		Nonce:        b.nonce,
		PrevHash:     b.prevHash,
		Timestamp:    b.timestamp,
		Transactions: b.transactions,
	})
}

func main() {
	//b := NewBlock(0, "init hash")
	//b.Print()

	blocChain := NewBlockchain()
	blocChain.Print()

	preHash := blocChain.LastBlock().Hash()
	blocChain.CreateBlock(1, preHash)
	blocChain.Print()
	preHash = blocChain.LastBlock().Hash()
	blocChain.CreateBlock(2, preHash)
	blocChain.Print()

	//block := &Block{nonce: 1}
	//fmt.Printf("%x\n", block.Hash())
}
