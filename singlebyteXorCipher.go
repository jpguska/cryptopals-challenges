package main
import (
	"encoding/hex"
	"fmt"
)

func main() {
	hexString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Println("Error decoding hex string:", err)
		return
	}
	fmt.Printf("Decoded bytes: %x\n", bytes)

	// Brute-force single-byte XOR key
	var bestKey byte
	var bestScore int
	var bestDecryption []byte

	for key := 0; key < 256; key++ {
		decryption := make([]byte, len(bytes))
		for i := 0; i < len(bytes); i++ {
			decryption[i] = bytes[i] ^ byte(key)
		}
		score := scoreEnglishText(decryption)
		if score > bestScore {
			bestScore = score
			bestKey = byte(key)
			bestDecryption = decryption
		}
	}

	fmt.Printf("Best key: %x\n", bestKey)
	fmt.Printf("Decrypted text: %s\n", bestDecryption)
}

func scoreEnglishText(text []byte) int { // Retorna int
	var score int // Declarado como int

	for _, b := range text {
		if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || b == ' ' {
			score++
		}
	}

	return score
}