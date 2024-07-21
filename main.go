package main

import (
	"github.com/arturfil/go_chain/types"
)

func main() {
	blckChain := types.NewBlockChain()

    blckChain.AddTransaction("A", "B", 1.0)

    previousHash := blckChain.LastBlock().Hash()
	blckChain.AddNewBlock(5, previousHash)
    blckChain.Print()

    blckChain.AddTransaction("C", "D", 2.0)
    blckChain.AddTransaction("X", "Y", 3.0)
    previousHash = blckChain.LastBlock().Hash()
	blckChain.AddNewBlock(2, previousHash)
    blckChain.Print()
}

