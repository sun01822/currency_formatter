package main

import (
	"fmt"
	"github.com/sun01822/currency_formatter/currency"
	"github.com/sun01822/currency_formatter/types"
)

func main() {
	value := currency.FormatCurrency(types.Formatter{
		Amount:   -1234567.89,
		Currency: "MYR",
		Format:   "RM ###,###,###.##",
	})

	fmt.Println(value)
}
