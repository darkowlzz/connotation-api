package actions

import (
	"log"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/csrf"
	"github.com/gobuffalo/buffalo/middleware/i18n"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/packr"

	asset "github.com/darkowlzz/connotation-api/internal/asset"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var T *i18n.Translator

// A collection of adjectives and adverbs
var words []string

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.Automatic(buffalo.Options{
			Env:         ENV,
			SessionName: "_connotation-api_session",
		})

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}
		if ENV != "test" {
			// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
			// Remove to disable this.
			app.Use(csrf.Middleware)
		}

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.PopTransaction)
		// Remove to disable this.
		// app.Use(middleware.PopTransaction(models.DB))

		// Setup and use translations:
		var err error
		if T, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
			app.Stop(err)
		}
		app.Use(T.Middleware())

		app.GET("/", HomeHandler)

		app.GET("/word", WordHandler)

		app.GET("/allwords", AllWordsHandler)

		app.POST("/train", TrainWordHandler)

		app.GET("/ping", PingHandler)

		app.ServeFiles("/assets", packr.NewBox("../public/assets"))
	}

	return app
}

func appendWordsFromFile(origSlice []string, filepath string) []string {
	newByteWords, err := asset.Asset(filepath)
	if err != nil {
		log.Println(err)
	}
	newWords := strings.Split(string(newByteWords), "\n")
	return append(origSlice, newWords...)
}

func getAllWords() []string {
	const adjectivesFile = "assets/text/adjectives.txt"
	const adverbsFile = "assets/text/adverbs.txt"

	if len(words) < 1 {
		words = appendWordsFromFile(words, adjectivesFile)
		words = appendWordsFromFile(words, adverbsFile)
	}

	return words
}
