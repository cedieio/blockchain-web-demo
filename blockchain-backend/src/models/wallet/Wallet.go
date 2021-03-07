package wallet

// Wallet model is used to store public and private key data per user
type Wallet struct {
	PublicKey  string
	PrivateKey string
}
