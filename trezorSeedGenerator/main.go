package main

import (
	"crypto/rand"
	"fmt"
	"github.com/tyler-smith/go-bip39"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error

	rand.Reader, err = os.Open("/dev/random")
	checkError(err)

	entropy, err := bip39.NewEntropy(256)
	checkError(err)

	mnemonic, err := bip39.NewMnemonic(entropy)
	checkError(err)

	words := strings.Split(mnemonic, " ")
	for idx, word := range words {
		fmt.Printf("%2d: %s\n", idx+1, word)
	}
}
