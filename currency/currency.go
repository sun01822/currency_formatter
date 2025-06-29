package currency

import (
	"fmt"
	"github.com/sun01822/currency_formatter/consts"
	"github.com/sun01822/currency_formatter/types"
	"math"
	"strings"
)

func FormatCurrency(formatter types.Formatter) string {
	switch formatter.Currency {
	case consts.USD:
		return USDFormat(formatter.Amount, formatter.Format)
	case consts.MYR:
		return MYRFormat(formatter.Amount, formatter.Format)
	default:
		return fmt.Sprintf("Unsupported currency: %s", formatter.Currency)
	}
}

func USDFormat(amount float64, format string) string {
	return ""
}

func MYRFormat(amount float64, format string) string {
	isNegative := false

	// If value is less than 0, return it directly
	if amount < 0 {
		amount = math.Abs(amount)
		isNegative = true
	}

	value := fmt.Sprintf("%.2f", amount)

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

	return FormatCurrencyWithSymbol(format, isNegative, result.String(), decPart)
}

func FormatCurrencyWithSymbol(format string, isNegative bool, intPart, desPart string) string {
	//RM ###,###,###.## or RM###,###,###.##
	extractSymbol := extractCurrencySymbol(format)

	if !isNegative {
		return extractSymbol + intPart + "." + desPart
	}

	return "-" + intPart + "." + desPart + extractSymbol
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
