package main

import (
	"fmt"
	"os"
)

type Teman struct {
	ID        int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var dataTeman = []Teman{
	{1, "Ibnu", "Jalan ABC No. 123", "Pengembang", "Belajar lebih dalam"},
	{2, "ALfarezi", "Jalan ABC No. 456", "Desainer", "Biar pro"},
	{3, "Ramadhan", "Jalan ABC No. 789", "Manajer Proyek", "cari pengalaman"},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run biodata.go [Nama Teman]")
		return
	}

	// Mengambil argumen nama dari CLI
	namaTeman := os.Args[1]

	// Mencari teman berdasarkan nama
	var temanDitemukan Teman
	for _, teman := range dataTeman {
		if teman.Nama == namaTeman {
			temanDitemukan = teman
			break
		}
	}

	if temanDitemukan.Nama == "" {
		fmt.Println("Teman dengan nama", namaTeman, "tidak ditemukan.")
		return
	}

	// Menampilkan data teman yang ditemukan, termasuk ID
	fmt.Println("ID:", temanDitemukan.ID)
	fmt.Println("Nama:", temanDitemukan.Nama)
	fmt.Println("Alamat:", temanDitemukan.Alamat)
	fmt.Println("Pekerjaan:", temanDitemukan.Pekerjaan)
	fmt.Println("Alasan:", temanDitemukan.Alasan)
}
