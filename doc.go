// Package currency_formatter provides utilities for formatting currency values
// with custom symbols and formats.
//
// Features:
//   - Format currency amounts with custom symbols and patterns
//   - Supports negative values
//   - Handles USD and MYR out of the box
//   - Easily extendable for other currencies
//
// Example:
//
//	import (
//	    "fmt"
//	    "github.com/sun01822/currency_formatter/currency"
//	    "github.com/sun01822/currency_formatter/types"
//	)
//
//	func main() {
//	    formatted := currency.FormatCurrency(types.Formatter{
//	        Amount:   -1234567.89,
//	        Currency: "MYR",
//	        Format:   "RM ###,###,###.##",
//	    })
//	    fmt.Println(formatted) // Output: -1,234,567.89RM
//	}
package main
