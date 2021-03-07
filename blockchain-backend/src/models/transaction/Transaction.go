package transaction

// Transaction model, contains the amount, payee and payer
type Transaction struct {
	Amount float64 `json:"amount"`
	Payer  string  `json:"payer"`
	Payee  string  `json:"payee"`
}
