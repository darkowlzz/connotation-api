package main

//go:generate go-bindata -pkg assets -o internal/asset/bindata.go assets/text/

import (
	"log"

	"github.com/darkowlzz/connotation-api/actions"
	"github.com/gobuffalo/envy"
)

func main() {
	port := envy.Get("PORT", "3000")
	app := actions.App()
	log.Fatal(app.Start(port))
}
