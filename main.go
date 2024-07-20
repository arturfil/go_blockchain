package main

import (
	"github.com/arturfil/go_chain/types"
)

func main() {
	blckChain := types.NewBlockChain()

    previousHash := blckChain.LastBlock().Hash()
	blckChain.AddNewBlock(5, previousHash)
    blckChain.Print()

    previousHash = blckChain.LastBlock().Hash()
	blckChain.AddNewBlock(2, previousHash)
    blckChain.Print()
}

