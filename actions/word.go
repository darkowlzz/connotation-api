package actions

import (
	"math/rand"

	"github.com/gobuffalo/buffalo"
)

// WordHandler returns a random word from a collection of words.
func WordHandler(c buffalo.Context) error {
	words := getAllWords()
	word := words[rand.Intn(len(words))]
	return c.Render(200, r.String(word))
}

// AllWordsHandler returns all the words the system known.
func AllWordsHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(getAllWords()))
}
