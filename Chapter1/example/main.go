package main

import (
	"github.com/wrbz15/refacetoring-doc/Chapter1/example/types"
	statement "github.com/wrbz15/refacetoring-doc/Chapter1/example/version3"
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
	result := statement.Statements(involices[0], plays)
	println(result)
	result = statement.HtmlStatement(involices[0], plays)
	println(result)
}
