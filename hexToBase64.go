package main
import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Println("Error decoding hex string:", err)
		return
	}

	base64String := base64.StdEncoding.EncodeToString(bytes)
	fmt.Println("Base64 string:", base64String)
}