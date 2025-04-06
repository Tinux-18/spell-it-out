package spell

func Spell(word string) []string {
	phoneticWords := make([]string, len(word))
	phoneticWords[0] = word
	return phoneticWords
}
