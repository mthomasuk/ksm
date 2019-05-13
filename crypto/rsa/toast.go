package rsa

import (
	"log"

	"github.com/mthomasuk/toast/crypto"
	"github.com/mthomasuk/toast/rsa"
)

//RSA/ECB/OAEPPadding
func OAEPPDecrypt(pub, pri string, ciphertext []byte) ([]byte, error) {
	key, err := rsa.LoadKeyFromPEMByte(
		[]byte(pub),
		[]byte(pri),
		rsa.ParsePKCS1KeyByCert)

	if err != nil {
		log.Printf("Loading key from PEM byte(s) failed :: %v", err)
		panic(err)
	}

	padding := rsa.NewOAEPPadding(key.Modulus())
	cipherMode := rsa.NewOAEPCipher()
	signMode := rsa.NewPKCS1v15Sign()

	cipher, err := crypto.NewRSAWith(key, padding, cipherMode, signMode)
	if err != nil {
		log.Printf("Creating cipher from RSA key failed :: %v", err)
		panic(err)
	}

	deT, err := cipher.Decrypt(ciphertext)
	if err != nil {
		log.Printf("Decrypting ciphertext failed :: %v", err)
		panic(err)
	}
	return deT, err
}
