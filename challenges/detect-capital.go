package challenges

func detectCapitalUse(word string) bool {

	chars := []rune(word)
	return allCaps(chars) || allSmall(chars) || titleCase(chars)

}

func allCaps(chars []rune) bool {
	for i := 0; i < len(chars); i++ {
		if int(chars[i]) < 65 || int(chars[i]) > 90 {
			return false
		}
	}
	return true
}

func allSmall(chars []rune) bool {
	for i := 0; i < len(chars); i++ {
		if int(chars[i]) < 97 || int(chars[i]) > 122 {
			return false
		}
	}
	return true
}

func titleCase(chars []rune) bool {
	for i := 0; i < len(chars); i++ {
		if i == 0 && (int(chars[i]) < 65 || int(chars[i]) > 90) {
			return false
		}
		if i != 0 && (int(chars[i]) < 97 || int(chars[i]) > 122) {
			return false
		}
	}
	return true
}
