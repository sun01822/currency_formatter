package currency

import (
	"fmt"
	"github.com/sun01822/currency_formatter/consts"
	"github.com/sun01822/currency_formatter/types"
	"math"
	"strings"
)

// FormatCurrency formats a monetary amount according to the specified currency and format.
// It handles negative values, applies comma separators to the integer part, and delegates symbol placement.
// Returns the formatted currency string.
func FormatCurrency(formatter types.Formatter) string {
	isNegative := false

	// If value is less than 0, return it directly
	if formatter.Amount < 0 {
		formatter.Amount = math.Abs(formatter.Amount)
		isNegative = true
	}

	value := fmt.Sprintf("%.2f", formatter.Amount)

	// Split integer and decimal
	parts := strings.Split(value, ".")
	intPart := parts[0]
	decPart := parts[1]

	// Format integer part with commas
	var result strings.Builder
	n := len(intPart)
	for i, digit := range intPart {
		if i > 0 && (n-i)%3 == 0 {
			result.WriteRune(',')
		}
		result.WriteRune(digit)
	}

	var reqPayload = types.FormatCurrencyWithSymbol{
		Format:      formatter.Format,
		IsNegative:  isNegative,
		Result:      result.String(),
		IntPart:     result.String(),
		DecimalPart: decPart,
		Currency:    formatter.Currency,
	}

	return FormatCurrencyWithSymbol(reqPayload)
}

// FormatCurrencyWithSymbol determines the currency type from the request payload
// and delegates formatting to the appropriate function (USD or MYR).
func FormatCurrencyWithSymbol(reqParam types.FormatCurrencyWithSymbol) string {
	switch reqParam.Currency {
	case consts.USD:
		return FormatCurrencyWithSymbolForUSD(reqParam)
	case consts.MYR:
		return FormatCurrencyWithSymbolForRM(reqParam)
	default:
		return fmt.Sprintf("Unsupported currency: %s", reqParam.Currency)
	}
}

// FormatCurrencyWithSymbolForUSD returns a formatted USD currency string.
// It places the currency symbol before the amount and handles negative values by prefixing with a minus sign.
func FormatCurrencyWithSymbolForUSD(reqParam types.FormatCurrencyWithSymbol) string {
	//USD $###,###,###.## or $###,###,###.##
	extractSymbol := extractCurrencySymbol(reqParam.Format)

	if !reqParam.IsNegative {
		return extractSymbol + reqParam.IntPart + "." + reqParam.DecimalPart
	}

	return "-" + extractSymbol + reqParam.IntPart + "." + reqParam.DecimalPart
}

// FormatCurrencyWithSymbolForRM formats Malaysian Ringgit currency values according to the specified format.
// For positive values, the currency symbol is prefixed; for negative values, the symbol is suffixed after the amount.
func FormatCurrencyWithSymbolForRM(reqParam types.FormatCurrencyWithSymbol) string {
	//RM ###,###,###.## or RM###,###,###.##
	extractSymbol := extractCurrencySymbol(reqParam.Format)

	if !reqParam.IsNegative {
		return extractSymbol + reqParam.IntPart + "." + reqParam.DecimalPart
	}

	return "-" + reqParam.IntPart + "." + reqParam.DecimalPart + extractSymbol
}

func extractCurrencySymbol(format string) string {
	var symbol strings.Builder
	for _, char := range format {
		if char != '#' && char != '.' && char != ',' {
			symbol.WriteRune(char)
		}
	}
	return symbol.String()
}
