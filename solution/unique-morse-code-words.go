package solution

import (
	"strings"
)

var codes = [26]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func UniqueMorseRepresentations(words []string) int {

	freqMap := make(map[string]int)

	for _, word := range words {
		var combinedCodes strings.Builder

		for _, char := range word {
			combinedCodes.WriteString(string(codes[char-97]))

		}
		freqMap[combinedCodes.String()]++

	}
	return len(freqMap)

}
