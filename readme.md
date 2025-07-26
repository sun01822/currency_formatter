# Currency Formatter

A Go library for formatting currency values with custom symbols and formats.

## Features

- Format currency amounts with custom symbols and patterns
- Supports negative values
- Handles USD and MYR out of the box
- Easily extendable for other currencies

## Usage

### Installation

```sh
go get github.com/sun01822/currency_formatter
```

### Example

```go
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
    fmt.Println(value) // Output: -1,234,567.89RM
}
```

## API

### `FormatCurrency(formatter types.Formatter) string`

Formats the currency amount using the provided formatter.

#### `types.Formatter`

- `Amount` (float64): The amount to format
- `Currency` (string): Currency code (e.g., "USD", "MYR")
- `Format` (string): Format string (e.g., "RM ###,###,###.##")

## Supported Currencies

- USD
- MYR

## License

Apache License 2.0