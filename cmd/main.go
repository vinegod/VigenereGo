package main

import (
	"fmt"
	"os"

	"github.com/vinegod/vigenerecode/vigenere"
)

func main() {

	arguments := os.Args

	if len(arguments) != 4 {
		fmt.Println("Usage: vigenerecode [-e|-d] [path to text] [path to key]")
		return
	}

	text, err := os.ReadFile(arguments[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	key, err := os.ReadFile(arguments[3])
	if err != nil {
		fmt.Println(err)
		return
	}

	vin := vigenere.Vigenere{Key: key}
	switch arguments[1] {
	case "-e":
		_, text = vin.Encode(text)
		fmt.Println(string(text))
	case "-d":
		_, text = vin.Decode(text)
		fmt.Println(string(text))
	}

}
