package spell

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

//go:embed data/*.json
var alphabetsFS embed.FS

type Acrophony struct {
	Letter     string `json:"letter"`
	Word       string `json:"word"`
	IsSelected bool   `json:"-"` // Used for highlighting in the UI
}

type Alphabet struct {
	Language    string      `json:"language"`
	Acrophonies []Acrophony `json:"acrophonies"`
}

func (a *Alphabet) GetAcrophony(letter string) (Acrophony, error) {
	char := strings.ToLower(letter)
	for _, acrophony := range a.Acrophonies {
		if acrophony.Letter == char {
			acrophony.Letter = letter // Maintain original case
			return acrophony, nil
		}
	}
	return Acrophony{}, fmt.Errorf("no acrophony found for letter '%s'", char)
}

func LoadAlphabets() []Alphabet {
	var alphabets []Alphabet

	// Read all files from the data directory
	entries, err := fs.ReadDir(alphabetsFS, "data")
	if err != nil {
		panic(fmt.Errorf("failed to read alphabets directory: %w", err))
	}

	for _, entry := range entries {
		// Read the file content
		filePath := filepath.Join("data", entry.Name())
		content, err := fs.ReadFile(alphabetsFS, filePath)
		if err != nil {
			panic(fmt.Errorf("failed to read alphabet file %s: %w", filePath, err))
		}

		// Parse JSON into Alphabet struct
		var alphabet Alphabet
		if err := json.Unmarshal(content, &alphabet); err != nil {
			panic(fmt.Errorf("failed to parse alphabet file %s: %w", filePath, err))
		}

		alphabets = append(alphabets, alphabet)
	}

	return alphabets
}

var data []Alphabet

func init() {
	data = LoadAlphabets()
}

func getAlphabetByLanguage(language string) (Alphabet, error) {
	for _, alphabet := range data {
		if alphabet.Language == language {
			return alphabet, nil
		}
	}
	// Raise error if not found.
	return Alphabet{}, fmt.Errorf("language '%s' not found", language)
}

func Word(word string, language string) ([]Acrophony, error) {
	// get alphabet
	alphabet, err := getAlphabetByLanguage(language)
	if err != nil {
		return nil, err
	}

	// spell phonetically
	word = strings.TrimSpace(word)
	phoneticWords := make([]Acrophony, 0, len(word))
	for _, letter := range word {
		stringLetter := string(letter)
		acrophony, err := alphabet.GetAcrophony(stringLetter)

		// If the letter is not found in the alphabet, use the letter itself as the word.
		if err != nil {
			acrophony = Acrophony{Letter: stringLetter, Word: stringLetter}
		}
		phoneticWords = append(phoneticWords, acrophony)
	}
	return phoneticWords, nil
}
