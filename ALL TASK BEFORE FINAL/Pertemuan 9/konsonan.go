package main

import (
    "fmt"
    "unicode"
)

func main() {
    var char rune
    fmt.Scanf("%c", &char) // Membaca satu karakter
    if unicode.IsLetter(char) && !isVowel(char) {
        // Jika karakter adalah huruf dan bukan vokal
        fmt.Println("konsonan")
    } else {
        // Selain itu (vokal atau bukan huruf)
        fmt.Println("bukan konsonan")
    }
}

// Fungsi untuk mengecek apakah karakter adalah vokal
func isVowel(c rune) bool {
    vowels := "aeiouAEIOU" // Daftar huruf vokal
    for _, v := range vowels {
        if c == v {
            return true
        }
    }
    return false
}
