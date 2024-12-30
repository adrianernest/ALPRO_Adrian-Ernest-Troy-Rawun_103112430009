package main

import "fmt"

func main() {
    var temp, count, sum, max, min int
    consecutiveZero := 0

    max = -1 << 31 // Nilai maksimum awal
    min = 1<<31 - 1 // Nilai minimum awal

    for {
        fmt.Scan(&temp)
        if temp == 0 {
            consecutiveZero++
        } else {
            consecutiveZero = 0
        }

        if consecutiveZero == 2 {
            break // Berhenti jika ada dua nol berturut-turut
        }

        if temp != 0 {
            if temp > max {
                max = temp // Memperbarui suhu maksimum
            }
            if temp < min {
                min = temp // Memperbarui suhu minimum
            }
            sum += temp
            count++
        }
    }

    fmt.Println(max)
    fmt.Println(min)
    fmt.Println(float64(sum) / float64(count)) // Rata-rata
}
