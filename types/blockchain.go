package types

import (
	"fmt"
	"strings"
)

type BlockChain struct {
	TransactionPool []string
	Chain           []*Block
}

func NewBlockChain() *BlockChain {
    block := &Block{}
    bchain := &BlockChain{}
    bchain.AddNewBlock(0, block.Hash())
	return bchain
}


func (bc *BlockChain) AddNewBlock(nonce int, previousHash [32]byte) {
    bc.Chain = append(bc.Chain, NewBlock(nonce, previousHash))
}

func (bc *BlockChain) LastBlock() *Block {
    return bc.Chain[len(bc.Chain)-1]
}

func (bc *BlockChain) Print() {
    for i, blck := range bc.Chain {
        fmt.Printf(
            "%s Chain %d %s\n", 
            strings.Repeat("=", 25), 
            i, strings.Repeat("=", 25),
        )
        blck.Print()
    }
    fmt.Printf("%s\n", strings.Repeat("*", 25))
}

