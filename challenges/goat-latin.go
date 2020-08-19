package challenges

import (
	"strings"
)

func toGoatLatin(S string) string {
	words := strings.Split(S, " ")
	for i := 0; i < len(words); i++ {
		words[i] = appendAs(processToGL(words[i]), i)

	}

	return strings.Join(words, " ")
}

func appendAs(word string, i int) string {
	for j := 0; j < i+1; j++ {
		word += "a"
	}
	return word
}
func processToGL(word string) string {
	vowels := []string{"A", "E", "I", "O", "U", "a", "e", "i", "o", "u"}

	if inList([]rune(word)[0], vowels) {
		return word + "ma"
	}

	return word[1:] + string(word[0]) + "ma"
}

func inList(ele rune, vowels []string) bool {
	for _, v := range vowels {
		if ele == []rune(v)[0] {
			return true
		}
	}
	return false
}
