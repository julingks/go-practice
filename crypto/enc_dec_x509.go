package main

import (
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"os"
)

func main() {
	blockType := "RSA PRIVATE KEY"
	password := []byte("password")

	cipherType := x509.PEMCipherAES256

	EncryptPEMBlock, err := x509.EncryptPEMBlock(rand.Reader,
		blockType,
		[]byte("secret message"),
		password,
		cipherType)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if !x509.IsEncryptedPEMBlock(EncryptPEMBlock) {
		fmt.Println("PEM Block is not encrypted!")
		os.Exit(1)
	}

	if EncryptPEMBlock.Type != blockType {
		fmt.Println("Block Type is wrong!")
		os.Exit(1)
	}

	fmt.Printf("Encrypt block \n%v\n", EncryptPEMBlock)

	fmt.Printf("Encrypted Block Headers Info : %v\n", EncryptPEMBlock.Headers)

	DecryptedPEMBlock, err := x509.DecryptPEMBlock(EncryptPEMBlock, password)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Decrypted block message is : \n%s\n", DecryptedPEMBlock)
}
