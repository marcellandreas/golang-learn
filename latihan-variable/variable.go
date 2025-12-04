package main

import "fmt"

func main() {
	var nama string = "Marcell"
	var umur int = 24
	city := "Jakarta"
	year := 2025

	fmt.Println("Nama saya", nama)
	fmt.Println("Umur saya", umur, "tahun")
	fmt.Println("Saya tinggal di", city)
	fmt.Println("Tahun sekarang adalah", year)

	// constant (konstanta) itu variable yang nilainya tidak bisa diubah-ubah
	const negara string = "Indonesia"
	fmt.Println("Saya berasal dari", negara)

	// negara = "Malaysia" // ini akan error karena konstanta tidak bisa diubah

}
