package main

import (
	"net/http"

	"github.com/jritsema/gotoolbox/web"
)

func index(r *http.Request) *web.Response {
	return web.HTML(http.StatusOK, html, "index.html", data, nil)
}

// GET /spelling
func getSpelling(r *http.Request) *web.Response {
	queryParams := r.URL.Query()
	word := queryParams.Get("word")
	language := queryParams.Get("lang")
	spelling, err := spell(word, language)
	if err != nil {
		return web.HTML(http.StatusNotFound, html, "spelling.html", spelling, nil)
	}
	return web.HTML(http.StatusOK, html, "spelling.html", spelling, nil)
}
