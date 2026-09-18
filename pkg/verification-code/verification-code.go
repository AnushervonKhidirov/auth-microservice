package verification_code

import (
	"math/rand"
)

func GenerateVerificationCode() string {
	codeLen := 6

	symbols := []rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	}

	generatedCode := make([]rune, codeLen)

	symbolsLen := len(symbols)

	for i := range codeLen {
		randSymIndex := rand.Intn(symbolsLen - 1)
		generatedCode[i] = symbols[randSymIndex]
	}

	return string(generatedCode)
}
