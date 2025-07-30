# Spell it out

A tool helping you to phonetically spell a text

## Motivation

At the behest of [Friedrich Mocker](https://github.com/Friedrich-Mocker) I read *[Hypermedia Systems](https://hypermedia.systems/book/contents/).*

When I learn something new, I use it to build something.
Almost mid way through the book I had to spell out my name in German over the phone and missed a letter.

This is where the idea came to write a hypermedia driven application which spells out a string for you using the [NATO phonetic alphabet](https://en.wikipedia.org/wiki/NATO_phonetic_alphabet).

There is already a [tool](https://www.spelltool.com "www.spelltool.com") developed by Anguel Stankov written in PHP, but I wanted to give it a Go as well.

The Go stack was inspired by [go-htmx-tailwind-example](https://github.com/jritsema/go-htmx-tailwind-example).

## Stack

- Go
- [HTMX](https://htmx.org/) (from CDN)
- [Bootstrap](https://getbootstrap.com/) (from CDN)

## Requirements

- [Docker Desktop](https://docs.docker.com/get-docker/)
- [Go](https://go.dev/) ^1.24
- [Air](https://github.com/air-verse/air)

## Run locally

Start and watch project:

```bash
air
```

## Improvements

- Get rid of Make

## Features

- Input a string
- Clean string of whitespaces
- Send request to spell it out
- Convert string to list of phonetic letters
- Display list of phonetic letters
- Pick language between German & English.

## TODOS

1. Highlight a row and move the highlight with the following events
   - click
   - arrowUp
   - arrowDown
2. Clear highlighting with:
   1. esc
   2. click outside the table
3. Improve styling
4. Footer with basic information
