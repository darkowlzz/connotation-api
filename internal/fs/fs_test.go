package fs

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

const tempFile = "/tmp/connotation-words-file.txt"

func TestGetWordsFromFile(t *testing.T) {
	words, err := GetWordsFromFile(tempFile)
	if err != nil {
		t.Fatal(err)
	}
	expectedWords := []string{"hello", "world", "it", "is", "a", "test"}
	for _, word := range expectedWords {
		if !contains(words, word) {
			t.Fatalf("expected %s to be in %s", word, words)
		}
	}
}

func createFileWithWords(filepath string) error {
	data := []byte("hello\nworld\nit\nis\na\n\n\ntest")
	err := ioutil.WriteFile(filepath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func TestMain(m *testing.M) {
	if err := createFileWithWords(tempFile); err != nil {
		log.Println("Failed to create file:", err)
		os.Exit(1)
	}

	exitVal := m.Run()

	if err := os.Remove(tempFile); err != nil {
		log.Println("Failed to delete file:", err)
	}

	os.Exit(exitVal)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
