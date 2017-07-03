package fs

import (
	"bufio"
	"os"
	"strings"
)

// GetWordsFromFile reads a given filepath and returns a collection of words
// from the file.
func GetWordsFromFile(filepath string) ([]string, error) {
	var wordsCollection []string

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return nil, err
	}

	if file, err := os.Open(filepath); err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			text := strings.TrimSpace(scanner.Text())
			// skip empty lines
			if text != "" {
				wordsCollection = append(wordsCollection, strings.TrimSpace(scanner.Text()))
			}
		}

		if err := scanner.Err(); err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}

	return wordsCollection, nil
}
