package main

import (
	"fmt"
	"sort"
)

// Struct untuk data buku
type Buku struct {
	ID       int
	Judul    string
	Penulis  string
	Penerbit string
	Tahun    int
	Rating   int
}

// Array buku dengan kapasitas maksimum
const Max = 100

var Pustaka [Max]Buku
var n int // jumlah data buku dalam pustaka

// Fungsi untuk memasukkan data buku
func DaftarkanBuku() {
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data buku ke-%d:\n", i+1)
		fmt.Print("ID: ")
		fmt.Scanln(&Pustaka[i].ID)
		fmt.Print("Judul: ")
		fmt.Scanln(&Pustaka[i].Judul)
		fmt.Print("Penulis: ")
		fmt.Scanln(&Pustaka[i].Penulis)
		fmt.Print("Penerbit: ")
		fmt.Scanln(&Pustaka[i].Penerbit)
		fmt.Print("Tahun: ")
		fmt.Scanln(&Pustaka[i].Tahun)
		fmt.Print("Rating: ")
		fmt.Scanln(&Pustaka[i].Rating)
	}
	fmt.Println("Semua buku telah didaftarkan.")
}

// Fungsi untuk mencari buku dengan rating tertinggi
func CetakFavorit() {
	if n == 0 {
		fmt.Println("Tidak ada data buku.")
		return
	}

	// Cari buku dengan rating tertinggi
	tertinggi := Pustaka[0]
	for i := 1; i < n; i++ {
		if Pustaka[i].Rating > tertinggi.Rating {
			tertinggi = Pustaka[i]
		}
	}

	fmt.Println("Buku Favorit:")
	fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
		tertinggi.ID, tertinggi.Judul, tertinggi.Penulis, tertinggi.Penerbit, tertinggi.Tahun, tertinggi.Rating)
}

// Fungsi untuk mengurutkan buku berdasarkan rating
func Urutkan() {
	sort.Slice(Pustaka[:n], func(i, j int) bool {
		return Pustaka[i].Rating > Pustaka[j].Rating
	})
	fmt.Println("Data buku telah diurutkan berdasarkan rating.")
}

// Fungsi untuk mencetak semua buku
func CetakSemuaBuku() {
	if n == 0 {
		fmt.Println("Tidak ada data buku.")
		return
	}

	fmt.Println("Daftar Buku:")
	for i := 0; i < n; i++ {
		fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
			Pustaka[i].ID, Pustaka[i].Judul, Pustaka[i].Penulis, Pustaka[i].Penerbit, Pustaka[i].Tahun, Pustaka[i].Rating)
	}
}

// Fungsi untuk mencari buku berdasarkan rating tertentu
func CariBuku() {
	fmt.Print("Masukkan rating yang ingin dicari: ")
	var ratingCari int
	fmt.Scanln(&ratingCari)

	ditemukan := false
	for i := 0; i < n; i++ {
		if Pustaka[i].Rating == ratingCari {
			fmt.Printf("Ditemukan buku dengan rating %d:\n", ratingCari)
			fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d\n",
				Pustaka[i].ID, Pustaka[i].Judul, Pustaka[i].Penulis, Pustaka[i].Penerbit, Pustaka[i].Tahun)
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Printf("Tidak ada buku dengan rating %d.\n", ratingCari)
	}
}

func main() {
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Daftarkan Buku")
		fmt.Println("2. Cetak Buku Favorit")
		fmt.Println("3. Urutkan Buku")
		fmt.Println("4. Cetak Semua Buku")
		fmt.Println("5. Cari Buku Berdasarkan Rating")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			DaftarkanBuku()
		case 2:
			CetakFavorit()
		case 3:
			Urutkan()
		case 4:
			CetakSemuaBuku()
		case 5:
			CariBuku()
		case 6:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
