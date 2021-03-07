package block

// Block the transaction block of the operation
type Block struct {
	PreviousHash    string `json:"previousHash"`
	TransactionTime int64  `json:"transactionTime"`
}
