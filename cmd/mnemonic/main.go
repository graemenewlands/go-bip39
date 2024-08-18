package main

import (
	"fmt"

	bip39 "github.com/graemenewlands/go-bip39/pkg/bip39"
)

func main() {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	fmt.Printf("%s\n", mnemonic)
}
