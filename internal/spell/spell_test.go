package spell

import "testing"

func TestWord(t *testing.T) {
	cases := []struct {
		in          string
		want        []Acrophony
		description string
	}{
		{"Ana", []Acrophony{{Letter: "a", Word: "Alpha"}, {Letter: "n", Word: "November"}, {Letter: "a", Word: "Alpha"}}, "Basic test"},
		{"   Ana", []Acrophony{{Letter: "a", Word: "Alpha"}, {Letter: "n", Word: "November"}, {Letter: "a", Word: "Alpha"}}, "Should be trimmed."},
		{"#@", []Acrophony{{Letter: "#", Word: "#"}, {Letter: "@", Word: "@"}}, "Special characters should be handled."},
	}

	for _, c := range cases {
		spelling, error := Word(c.in, "en")
		if !isSameAcrophonyList(spelling, c.want) {
			t.Errorf("Word(%q, \"en\") == %q, want %q. Description: %s", c.in, spelling, c.want, c.description)
		}
		if error != nil {
			t.Errorf("Word(%q, \"en\") returned error: %v. Description: %s", c.in, error, c.description)
		}
	}
}

func isSameAcrophonyList(a, b []Acrophony) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
