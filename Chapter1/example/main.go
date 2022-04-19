package main

import (
	"github.com/wrbz15/refacetoring-doc/Chapter1/example/types"
	"github.com/wrbz15/refacetoring-doc/Chapter1/example/version1"
)

func main() {
	plays, err := types.LoadPlays("./data/plays.json")
	if err != nil {
		panic(err)
	}
	involices, err := types.LoadInvoice("./data/invoices.json")
	if err != nil {
		panic(err)
	}
	result, err := version1.Statements(involices[0], plays)
	if err != nil {
		panic(err)
	}
	println(result)
}