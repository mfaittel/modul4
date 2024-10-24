package main

import (
	"fmt"
	"strings"
)

func hitungSkor(soal []int) (int, int) {
	totalSoal := 0
	totalWaktu := 0
	for _, waktu := range soal {
		if waktu < 301 {
			totalSoal++
			totalWaktu += waktu
		}
	}
	return totalSoal, totalWaktu
}

func main() {

	peserta := []string{
		"Astuti 20 50 301 301 61 71 75 10",
		"Bertha 25 47 301 26 50 60 65 21",
	}

	pemenang := ""
	maxSoal := 0
	minWaktu := 9999999
	totalSoalPemenang := 0

	for _, data := range peserta {
		parts := strings.Fields(data)
		nama := parts[0]
		soal := make([]int, 8)
		for j := 1; j < len(parts); j++ {
			fmt.Sscanf(parts[j], "%d", &soal[j-1])
		}

		totalSoal, totalWaktu := hitungSkor(soal)

		if totalSoal > maxSoal || (totalSoal == maxSoal && totalWaktu < minWaktu) {
			pemenang = nama
			maxSoal = totalSoal
			minWaktu = totalWaktu
			totalSoalPemenang = totalSoal
		}
	}

	fmt.Printf("%s %d %d\n", pemenang, totalSoalPemenang, minWaktu)
}
