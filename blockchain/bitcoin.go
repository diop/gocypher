package blockchain

import (
	"fmt"
	"os"

	"github.com/blockcypher/gobcy"
)

func bitcoin() {
	btc := gobcy.API{os.Getenv("TOKEN"), "btc", "main"}
	chain, err := btc.GetChain()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", chain)
}
