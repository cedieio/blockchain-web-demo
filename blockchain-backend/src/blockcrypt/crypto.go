package blockcrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
)

func generateKeys() (*rsa.PrivateKey, *rsa.PublicKey) {
	keys, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Error : %v", err)
	}
	return keys, &keys.PublicKey
}

func generatePrivateKeyPemString(privateKey *rsa.PrivateKey) string {
	var privateKeyBytes []byte = x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	pemString := pem.EncodeToMemory(privateKeyBlock)
	return string(pemString)
}

func generatePublicKeyPemString(publicKey *rsa.PublicKey) string {
	var publicKeyBytes []byte = x509.MarshalPKCS1PublicKey(publicKey)
	publicKeyBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	pemString := pem.EncodeToMemory(publicKeyBlock)
	return string(pemString)
}

/* Usage
func main() {
	privateKey, publicKey := generateKeys()

	privateKeyString := generatePrivateKeyPemString(privateKey)
	publicKeyString := generatePublicKeyPemString(publicKey)
	cedie := &wallet.Wallet{
		PublicKey:  publicKeyString,
		PrivateKey: privateKeyString,
	}
	fmt.Printf("%+v \n", cedie)
}
*/
