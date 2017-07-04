# Connotation-api

[![Build Status](https://travis-ci.org/darkowlzz/connotation-api.svg?branch=master)](https://travis-ci.org/darkowlzz/connotation-api)

A [Buffalo](http://gobuffalo.io) powered api server.


## Generating binary assets

Some text file assets are at `assets/text/`. When those files are modified,
go source for them have to be regenerated.

Use [go-bindata](https://github.com/jteeuwen/go-bindata) to generate those
files. Run `go generate -v -x ./main.go` to generate source under
`internal/asset/` package. The generated go source should be checked in.
