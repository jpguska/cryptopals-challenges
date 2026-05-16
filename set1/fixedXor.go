package main
import (
	"encoding/hex"
	"fmt"
)

func main() {

	hexString1 := "1c0111001f010100061a024b53535009181c"
	hexString2 := "686974207468652062756c6c277320657965"
	bytes1, err := hex.DecodeString(hexString1)
	if err != nil {
		fmt.Println("Error decoding hex string 1:", err)
		return
	}

	bytes2, err := hex.DecodeString(hexString2)
	if err != nil {
		fmt.Println("Error decoding hex string 2:", err)
		return
	}

	if len(bytes1) != len(bytes2) {
		fmt.Println("Error: Hex strings must be of the same length")
		return
	}
	
	result := make([]byte, len(bytes1))

	for i := 0; i < len(bytes1); i++ {
		result[i] = bytes1[i] ^ bytes2[i]
	}

	fmt.Printf("XOR result: %x\n", result)
}