package main

import (
	"bufio"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"

	bip39 "github.com/graemenewlands/go-bip39/pkg/bip39"
)

func process(line string) {
	seed, err := bip39.MnemonicToByteArray(line)

	if err != nil {
		log.Fatalf("error converting mnemonic %s to byte array: %s", line, err)
	}

	// the seed actually contains an extra array element (checksum), that needs to
	// be removed.
	// the consts declared in the ed25519 package describe the valid seed size settings
	privkey := ed25519.NewKeyFromSeed(seed[:len(seed)-1])

	privkeyBytes, err := x509.MarshalPKCS8PrivateKey(privkey)
	if err != nil {
		log.Fatalf("Error marshaling private key: %v", err)
	}

	// Manually create a PEM block for the private key
	pemBlock := &pem.Block{
		Type:  "EC PRIVATE KEY",
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
