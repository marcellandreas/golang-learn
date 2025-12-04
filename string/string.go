package main

import "fmt"

func main() {
	// Dekrasi string dengan kutip ganda

	nama := "Marcell"
	pesan := "Selamat datang di Golang"

	// Dekrasi string dengan backtick
	paragraf := `Golang adalah bahasa pemrograman yang
				dikembangkan oleh Google.
				Bahasa ini dikenal karena performanya yang cepat
				dan sintaksisnya yang sederhana.`

	// bedanya dengan kutip ganda, backtick dapat menampung string multi-baris tanpa perlu karakter khusus seperti \n

	// Menampilkan nilai variabel string
	fmt.Println("Nama:", nama)
	fmt.Println("Pesan:", pesan)
	fmt.Println("Paragraf:\n", paragraf)

}
