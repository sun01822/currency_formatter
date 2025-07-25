package currency

import (
	"fmt"
	"github.com/sun01822/currency_formatter/consts"
	"github.com/sun01822/currency_formatter/types"
	"math"
	"strings"
)

// FormatCurrency formats the amount based on the currency and format provided in the formatter.
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

// FormatCurrencyWithSymbol formats the currency with the given format and symbol
func FormatCurrencyWithSymbol(reqParam types.FormatCurrencyWithSymbol) string {
	switch reqParam.Currency {
	case consts.USD:
		return FormatCurrencyWithSymbolForUsd(reqParam)
	case consts.MYR:
		return FormatCurrencyWithSymbolForRm(reqParam)
	default:
		return "Unsupported currency"
	}
}

// FormatCurrencyWithSymbolForUSD formats the currency with the given format and symbol
func FormatCurrencyWithSymbolForUsd(reqParam types.FormatCurrencyWithSymbol) string {
	//USD $###,###,###.## or $###,###,###.##
	extractSymbol := extractCurrencySymbol(reqParam.Format)

	if !reqParam.IsNegative {
		return extractSymbol + reqParam.IntPart + "." + reqParam.DecimalPart
	}

	return "-" + extractSymbol + reqParam.IntPart + "." + reqParam.DecimalPart
}

// FormatCurrencyWithSymbolForRm formats the currency with the given format and symbol
func FormatCurrencyWithSymbolForRm(reqParam types.FormatCurrencyWithSymbol) string {
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
