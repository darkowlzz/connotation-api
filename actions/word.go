package actions

import (
	"math/rand"

	"github.com/gobuffalo/buffalo"
)

type wordResponse struct {
	Data  string `json:"data"`
	Error string `json:"error,omitempty"`
}

type wordsResponse struct {
	Data  []string `json:"data"`
	Error string   `json:"error,omitempty"`
}

type trainResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

// WordHandler returns a random word from a collection of words.
func WordHandler(c buffalo.Context) error {
	words := getAllWords()
	word := words[rand.Intn(len(words))]
	return c.Render(200, r.JSON(wordResponse{Data: word}))
}

// AllWordsHandler returns all the words the system known.
func AllWordsHandler(c buffalo.Context) error {
	// return c.Render(200, r.JSON(getAllWords()))
	return c.Render(200, r.JSON(wordsResponse{Data: getAllWords()}))
}

// TrainWordHandler receives training data.
// NOTE: No training engine at the moment.
func TrainWordHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(trainResponse{Success: true}))
}
