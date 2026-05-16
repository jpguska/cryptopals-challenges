package main
import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("4.txt")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo:", err)
		return
	}
	defer file.Close()

	// Variáveis para guardar o "campeão global" de todo o arquivo
	var bestScore int
	var bestText []byte

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linhaHex := scanner.Text()

		linhaBytes, err := hex.DecodeString(linhaHex)
		if err != nil {
			fmt.Println("Erro ao decodificar linha:", err)
			continue
		}

		for key := 0; key < 256; key++ {
			decrypted := make([]byte, len(linhaBytes))
			for i, b := range linhaBytes {
				decrypted[i] = b ^ byte(key)
			}

			score := scoreEnglishText(decrypted)
			
			// Se a pontuação desta chave/linha for a maior que já vimos, atualizamos o campeão
			if score > bestScore {
				bestScore = score
				// Fazemos uma cópia limpa dos bytes na memória usando append
				bestText = append([]byte(nil), decrypted...)
			}
		}
	}

	// Fora de todos os laços, imprimimos o grande vencedor!
	fmt.Printf("A mensagem oculta é: %s\n", bestText)
}

func scoreEnglishText(text []byte) int {
	var score int

	for _, b := range text {
		// Dá pontos para letras maiúsculas, minúsculas e espaços
		if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || b == ' ' {
			score++
		}
	}

	return score
}