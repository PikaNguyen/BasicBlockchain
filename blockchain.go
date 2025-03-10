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
	transactions []*Transaction
}

type Blockchain struct {
	transactionPool []*Transaction
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
	b := NewBlock(nonce, prevHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp: 		%d\n", b.timestamp)
	fmt.Printf("nonce:         	%d\n", b.nonce)
	fmt.Printf("prev_Hash: 		%x\n", b.prevHash)
	//fmt.Printf("transactions: 	%s\n", b.transactions)
	for _, t := range b.transactions {
		t.Print()
	}
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 80))
}

func NewBlock(nonce int, prevHash [32]byte, transaction []*Transaction) *Block {
	/*
		b:= new (Block)
		b.timestamp = time.Now().UnixNano()
		b.nonce = nonce
		b.prevHash = prevHash
		b.transactions = transaction
		return b
	*/
	return &Block{
		timestamp:    time.Now().UnixNano(),
		nonce:        nonce,
		prevHash:     prevHash,
		transactions: transaction,
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nonce        int            `json:"nonce"`
		PrevHash     [32]byte       `json:"prev_hash"`
		Timestamp    int64          `json:"timestamp"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Nonce:        b.nonce,
		PrevHash:     b.prevHash,
		Timestamp:    b.timestamp,
		Transactions: b.transactions,
	})
}

type Transaction struct {
	senderBlockchainAddress    string //địa chỉ người gửi
	recipientBlockchainAddress string //địa chỉ người nhận
	value                      float32
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf("sender_blockchain_address    %s\n", t.senderBlockchainAddress)
	fmt.Printf("recipient_blockchain_address    %s\n", t.recipientBlockchainAddress)
	fmt.Printf("value    %.1f\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json: "sender_blockchain_address`
		Recipient string  `json: "recipient_blockchain_address"`
		Value     float32 `json: "value"`
	}{
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)

}

func main() {
	//b := NewBlock(0, "init hash")
	//b.Print()

	blocChain := NewBlockchain()
	blocChain.Print()

	blocChain.AddTransaction("A", "B", 0.1)
	preHash := blocChain.LastBlock().Hash()
	blocChain.CreateBlock(1, preHash)
	blocChain.Print()

	blocChain.AddTransaction("C", "E", 0.1)
	blocChain.AddTransaction("CDS", "E11", 12)
	preHash = blocChain.LastBlock().Hash()
	blocChain.CreateBlock(2, preHash)
	blocChain.Print()

	//block := &Block{nonce: 1}
	//fmt.Printf("%x\n", block.Hash())
}
