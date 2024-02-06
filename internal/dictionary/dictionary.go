package dictionary

import (
	"bufio"
	"os"
)

type dictionary map[string]struct{}

// Create a new dictionary
func newDictionary() *dictionary {
	return &dictionary{}
}
func (d *dictionary) add(s string) {
	(*d)[s] = struct{}{}
}
func (d *dictionary) Contains(s string) bool {
	_, ok := (*d)[s]
	return ok
}

// loadDictionary function loads text file from path to a dictionary
func LoadDictionary(path string) (*dictionary, error) {
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
