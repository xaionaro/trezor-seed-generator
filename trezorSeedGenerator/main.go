package main

import (
	"crypto/rand"
	"fmt"
	"github.com/pborman/getopt/v2"
	"math/big"
	"strings"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	//randomSourcePath := getopt.StringLong("random-source", 'r', "/dev/random", "the source of random number (secure) [default: /dev/random]")
	dictionaryPath := getopt.StringLong("dictionary-path", 'd', "", "the source of random number (secure)")

	getopt.Parse()

	if *dictionaryPath == "" {
		panic("dictionary path is not set")
	}

	randomReader, err := os.Open("/dev/random")
	if err != nil {
		panic(err)
	}

	dictionary, err := ioutil.ReadFile(*dictionaryPath)
	if err != nil {
		panic(err)
	}
	dictionaryLines := strings.Split(string(dictionary), "\n")
	dictionaryLinesCountBigInt := big.NewInt(int64(len(dictionaryLines)))

	re := regexp.MustCompile("[^a-z].*")
	wordsCount := 0
	for {
		randValue, err := rand.Int(randomReader, dictionaryLinesCountBigInt)
		if err != nil {
			panic(err)
		}
		word := strings.ToLower(dictionaryLines[randValue.Int64()])
		word = re.ReplaceAllString(word, "")
		if len(word) < 3 {
			continue
		}
		wordsCount++
		fmt.Printf("%2d: %v\n", wordsCount, word)
		if wordsCount >= 24 {
			break
		}
	}
}
