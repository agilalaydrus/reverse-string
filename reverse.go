package main

import (
	"fmt"
)

// ReverseString membalikkan urutan karakter dalam string
func ReverseString(s string) string {
	runes := []rune(s) // Mengubah string menjadi slice rune untuk menangani karakter Unicode
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // Menukar karakter
	}
	return string(runes) // Mengubah slice rune kembali menjadi string
}

func main() {
	// Contoh penggunaan
	input := "Halo, Dunia!"
	output := ReverseString(input)
	fmt.Println(output) // Output: "!ainuD ,olaH"
}
