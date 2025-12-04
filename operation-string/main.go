package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Hello, Golang! Welcome to the world of Go programming."

	// mengubah menjadi huruf kecil
	fmt.Println("Lowercase:", strings.ToLower(text))

	// mengubah menjadi huruf besar
	fmt.Println("Uppercase:", strings.ToUpper(text))

	// mengecek apakah string mengandung substring tertentu
	fmt.Println("Contains 'Golang':", strings.Contains(text, "Golang"))

	// mengecek apakah string dimulai dengan Hello
	fmt.Println("Starts with 'Hello':", strings.HasPrefix(text, "Hello"))

	// memisahkan string berdasarkan delimiter spasi
	parts := strings.Split(text, " ")
	fmt.Println("Split by space:", parts)

	// memisah string berdasarkan delimiter
	fruits := strings.Split("apel, banana, cerry", ",")
	fmt.Println("Split by comma:", fruits)

	// mengganti bagian dari string
	replacedText := strings.ReplaceAll(text, "Golang", "Go")
	fmt.Println("Replaced Text:", replacedText)

}
