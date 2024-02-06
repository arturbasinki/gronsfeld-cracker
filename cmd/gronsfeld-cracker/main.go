package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const alphabet = " abcdefghijklmnopqrstuvwxyz"

type dictionary map[string]struct{}

// Create a new dictionary
func newDictionary() *dictionary {
	return &dictionary{}
}
func (d *dictionary) add(s string) {
	(*d)[s] = struct{}{}
}
func (d *dictionary) contains(s string) bool {
	_, ok := (*d)[s]
	return ok
}

// loadDictionary function loads text file from path to a dictionary
func loadDictionary(path string) (*dictionary, error) {
	dict := newDictionary()

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		dict.add(word)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return dict, nil
}
func keyGen(limit int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 1; i < limit; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

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

	dictionary, err := loadDictionary("slownik.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through each word in ciphertext and check if it's in dictionary
	for _, word := range strings.Fields(ciphertext) {
		if dictionary.contains(word) {
			fmt.Println(word)
		}
	}

}
