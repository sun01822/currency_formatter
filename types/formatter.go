package types

type Formatter struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Format   string  `json:"format"`
}
