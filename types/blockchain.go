package types

import (
	"fmt"
	"strings"
)

type BlockChain struct {
	transactionPool []*Transaction
	chain           []*Block
}

func NewBlockChain() *BlockChain {
    block := &Block{}
    bchain := &BlockChain{}
    bchain.AddNewBlock(0, block.Hash())
	return bchain
}


func (bc *BlockChain) AddNewBlock(nonce int, previousHash [32]byte) *Block {
    bc.chain = append(bc.chain, NewBlock(nonce, previousHash, bc.transactionPool))
    bc.transactionPool = []*Transaction{}
    return bc.chain[len(bc.chain)-1]
}

func (bc *BlockChain) AddTransaction(sender, recipient string, value float32) {
    t := NewTransaction(sender, recipient, value)
    bc.transactionPool = append(bc.transactionPool, t)
}

func (bc *BlockChain) LastBlock() *Block {
    return bc.chain[len(bc.chain)-1]
}

func (bc *BlockChain) Print() {
    for i, blck := range bc.chain {
        fmt.Printf(
            "%s Chain %d %s\n", 
            strings.Repeat("=", 35), 
            i, strings.Repeat("=", 35),
        )
        blck.Print()
    }
    fmt.Printf("\n%s\n", strings.Repeat("*", 79))
}

