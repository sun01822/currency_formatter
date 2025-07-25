package types

type Formatter struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Format   string  `json:"format"`
}

type FormatCurrencyWithSymbol struct {
	Format      string `json:"format"`
	IsNegative  bool   `json:"is_negative"`
	Result      string `json:"result"`
	IntPart     string `json:"int_part"`
	DecimalPart string `json:"decimal_part"`
	Currency    string `json:"currency"`
}
