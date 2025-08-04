package main

import (
	"net/http"

	"github.com/Tinux-18/spell-it-out/internal/spell"
	"github.com/jritsema/gotoolbox/web"
)

func index(r *http.Request) *web.Response {
	return web.HTML(http.StatusOK, html, "index.html", nil, nil)
}

func getSpelling(r *http.Request) *web.Response {
	queryParams := r.URL.Query()
	word := queryParams.Get("word")
	language := queryParams.Get("lang")
	spelling, err := spell.Word(word, language)

	if err != nil {
		return web.HTML(http.StatusNotFound, html, "spelling.html", spelling, nil)
	}

	if r.Header.Get("HX-Request") == "true" {
		// Return just the fragment for HTMX
		return web.HTML(http.StatusOK, html, "spelling.html", spelling, nil)
	} else {
		// Return full page with spelling data for direct browser access
		return web.HTML(http.StatusOK, html, "index.html", spelling, nil)
	}
}
