package main

import (
	"bufio"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/hkdf"

	bip39 "github.com/graemenewlands/go-bip39/pkg/bip39"
)

func NewP256FromSeedHKDF(seed []byte) (*ecdsa.PrivateKey, error) {
	// domain separation string is important; change it if you want a different “context”
	r := hkdf.New(sha256.New, seed, nil, []byte("go-ecdsa-p256-keygen-v1"))
	return ecdsa.GenerateKey(elliptic.P256(), r)
}

func process(line string) {
	seed, err := bip39.MnemonicToByteArray(line)

	if err != nil {
		log.Fatalf("error converting mnemonic %s to byte array: %s", line, err)
	}

	// the seed actually contains an extra array element (checksum), that needs to
	// be removed.
	// the consts declared in the ed25519 package describe the valid seed size settings
	privkey, err := NewP256FromSeedHKDF(seed[:len(seed)-1])
	if err != nil {
		log.Fatalf("Unable to load private key bytes: %v", err)
	}

	privkeyBytes, err := x509.MarshalPKCS8PrivateKey(privkey)
	if err != nil {
		log.Fatalf("Error marshaling private key: %v", err)
	}

	// Manually create a PEM block for the private key
	pemBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privkeyBytes,
	}
	pemData := pem.EncodeToMemory(pemBlock)

	fmt.Printf("%s", pemData)
}

func main() {
	// Create a new scanner to read from standard input (stdin)
	scanner := bufio.NewScanner(os.Stdin)

	// Loop over each line of input
	for scanner.Scan() {
		// Read the line from stdin
		line := scanner.Text()

		// Process or print the line
		process(line)
	}

	// Check for any errors encountered by the scanner
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from input: %v\n", err)
	}

}
