package main

import "encoding/base64"

func main() {
	// Base64 encoding is a method of converting binary data into an ASCII string format.
	// It is commonly used to encode data for transmission over text-based protocols, such as email or HTTP.

	value := "Akbar Zulfikar"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	println("Encoded:", encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		println("Error decoding:", err.Error())
	} else {
		println("Decoded:", string(decoded))
	}
}
