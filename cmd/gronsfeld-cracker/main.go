package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/arturbasinki/gronsfeld-cracker/internal/dictionary"
)

const alphabet = " abcdefghijklmnopqrstuvwxyz"

func getCiphertext() string {
	var cipherText string
	// Prompt user to enter cipher text
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter cipher text: ")
	cipherText, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	cipherText = strings.TrimSpace(cipherText)
	return cipherText
}

func main() {
	ciphertext := getCiphertext()
	// fmt.Println(ciphertext)

	absPath, err := filepath.Abs("slownik.txt")
	if err != nil {
		log.Fatal(err)
	}

	dictionary, err := dictionary.LoadDictionary(absPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s is in dict: %v\n", ciphertext, dictionary.Contains(ciphertext))
	// // Iterate through each word in ciphertext and check if it's in dictionary
	// for _, word := range strings.Fields(ciphertext) {
	// 	if dictionary.contains(word) {
	// 		fmt.Println(word)
	// 	}
	// }

}
