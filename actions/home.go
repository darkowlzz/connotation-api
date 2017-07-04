package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}

// PingHandler returns pong to provide service availability check endpoint.
func PingHandler(c buffalo.Context) error {
	return c.Render(200, r.String("pong!"))
}
