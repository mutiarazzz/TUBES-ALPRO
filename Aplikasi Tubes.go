package main

import (
	"fmt"
)

type Pakaian struct {
	Nama       string
	Kategori   string
	Warna      string
	Formalitas int
	Cuaca      string
	Jenis      string

	// TerakhirDipakai time.Time
}

const NMAX int = 1000

type tabPakaian [NMAX]Pakaian

func main() {
	var data tabPakaian
	var nData int
	nData = 1
	var pilihan int
	for {
		fmt.Println("\n--- Menu OOTD Planner ---")
		fmt.Println("1. Tampilkan Semua Pakaian")
		fmt.Println("2. Tambah Outfit")
		fmt.Println("3. Hapus Data Outfit")
		fmt.Println("4. Rekomendasi Outfit")
		fmt.Println("5. Cari Pakaian")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tampilkanData(data, nData)
		case 2:
			tambahData(&data, &nData)
		case 3:
			hapusData(&data, &nData)
		case 4:
			rekomendasiOutfit(data, nData)
		case 5:
			cariPakaian(data, nData)
		case 6:
			fmt.Println("Terima kasih telah menggunakan OOTD Planner!")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

func tambahData(A *tabPakaian, n *int) {
	fmt.Print("Silahkan Masukkan Nama Pakaian: ")
	fmt.Scan(&A[*n].Nama)
	fmt.Print("Silahkan Masukkan Kategori Pakaian: ")
	fmt.Scan(&A[*n].Kategori)
	fmt.Print("Silahkan Masukkan Warna Pakaian yang dimiliki: ")
	fmt.Scan(&A[*n].Warna)
	fmt.Print("Silahkan Masukkan Tingkat Formalitas dari Pakaian: ")
	fmt.Scan(&A[*n].Formalitas)
	fmt.Print("Silahkan Masukkan Cuaca yang Cocok untuk Pakaian Anda(HUjan/Panas/Dingin): ")
	fmt.Scan(&A[*n].Cuaca)
	fmt.Print("Silahkan Masukkan Jenis Pakaian(Atasan/Bawahan/Footwear): ")
	fmt.Scan(&A[*n].Jenis)
	*n++

}

func tampilkanData(A tabPakaian, n int) {
	var i int
	for i = 1; i < n; i++ {
		fmt.Printf("%d. Nama Pakaian: %s, Kategori: %s, Warna: %s, Formalitas: %d, Cuaca: %s, Jenis: %s\n", i, A[i].Nama, A[i].Kategori, A[i].Warna, A[i].Formalitas, A[i].Cuaca, A[i].Jenis)
	}
}

func rekomendasiOutfit(A tabPakaian, n int) {
	var pilihan int
	var cuaca string
	var acara string
	var i int

	fmt.Println("1. Rekomendasi Outfit Berdasarkan Cuaca ")
	fmt.Println("2. Rekomendasi Outfit Berdasarkan Acara yang dihadiri ")
	fmt.Println("Masukkan Pilihan Anda: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		fmt.Print("Masukkan Cuaca Saat ini: ")
		fmt.Scan(&cuaca)
		fmt.Println("Outfit yang Cocok untuk Cuaca", cuaca)
		for i = 1; i < n; i++ {
			if A[i].Cuaca == cuaca {
				fmt.Printf("Nama Pakaian: %s, Kategori: %s, Warna: %s, Formalitas: %d, Cuaca: %s, Jenis: %s\n", A[i].Nama, A[i].Kategori, A[i].Warna, A[i].Formalitas, A[i].Cuaca, A[i].Jenis)
			}
		}
	case 2:
		fmt.Print("Masukkan Acara yang Akan dihadiri(Rapat/Bermain/Pernikahan/Pesta): ")
		fmt.Scan(&acara)
		fmt.Println("Outfit yang Cocok untuk Acara ", acara)
		for i = 1; i < n; i++ {
			if acara == "Rapat" && A[i].Formalitas >= 4 {
				fmt.Printf("Nama Pakaian: %s, Kategori: %s, Warna: %s, Formalitas: %d, Cuaca: %s, Jenis: %s\n", A[i].Nama, A[i].Kategori, A[i].Warna, A[i].Formalitas, A[i].Cuaca, A[i].Jenis)
			} else if acara == "Bermain" && A[i].Formalitas >= 2 {
				fmt.Printf("Nama Pakaian: %s, Kategori: %s, Warna: %s, Formalitas: %d, Cuaca: %s, Jenis: %s\n", A[i].Nama, A[i].Kategori, A[i].Warna, A[i].Formalitas, A[i].Cuaca, A[i].Jenis)
			} else if acara == "Pernikahan" && A[i].Formalitas >= 5 {
				fmt.Printf("Nama Pakaian: %s, Kategori: %s, Warna: %s, Formalitas: %d, Cuaca: %s, Jenis: %s\n", A[i].Nama, A[i].Kategori, A[i].Warna, A[i].Formalitas, A[i].Cuaca, A[i].Jenis)
			} else if acara == "Pesta" && A[i].Formalitas >= 5 {
				fmt.Printf("Nama Pakaian: %s, Kategori: %s, Warna: %s, Formalitas: %d, Cuaca: %s, Jenis: %s\n", A[i].Nama, A[i].Kategori, A[i].Warna, A[i].Formalitas, A[i].Cuaca, A[i].Jenis)
			} else {
				fmt.Println("Tidak Ada Pakaian yang Tersedia")
			}
		}
	}

}

func cariPakaian(A tabPakaian, n int) {
	var pilihan int
	var nama string
	var kategori string
	var warna string
	var cuaca string
	var i int

	fmt.Println("1. Cari Pakaian Bedasarkan Nama")
	fmt.Println("2. Cari Pakaian Bedasarkan Kategori")
	fmt.Println("3. Cari Pakaian Bedasarkan Warna")
	fmt.Println("4. Cari Pakaian Bedasarkan Cuaca")

	switch pilihan {
	case 1:
		fmt.Print("Masukkan Nama Pakaian yang Ingin dicari(Kaos/Jacket/Celana/dll): ")
		fmt.Scan(&nama)
		for i = 1; i < n; i++ {
			if A[i].Nama == nama {
				fmt.Printf("Nama Pakaian: %s, Kategori: %s, Warna: %s, Formalitas: %d, Cuaca: %s, Jenis: %s\n", A[i].Nama, A[i].Kategori, A[i].Warna, A[i].Formalitas, A[i].Cuaca, A[i].Jenis)
			}
		}
	case 2:
		fmt.Print("Masukkan Nama Kategori yang Ingin dicari(Casual/Formal/Sporty): ")
		fmt.Scan(&kategori)
		for i = 1; i < n; i++ {
			if A[i].Kategori == kategori {
				fmt.Printf("Nama Pakaian: %s, Kategori: %s, Warna: %s, Formalitas: %d, Cuaca: %s, Jenis: %s\n", A[i].Nama, A[i].Kategori, A[i].Warna, A[i].Formalitas, A[i].Cuaca, A[i].Jenis)
			}
		}
	case 3:
		fmt.Print("Masukkan Warna Pakaian yang Ingin dicari(Hitam/Putih/Biru/Pink/dll): ")
		fmt.Scan(&warna)
		for i = 1; i < n; i++ {
			if A[i].Warna == warna {
				fmt.Printf("Nama Pakaian: %s, Kategori: %s, Warna: %s, Formalitas: %d, Cuaca: %s, Jenis: %s\n", A[i].Nama, A[i].Kategori, A[i].Warna, A[i].Formalitas, A[i].Cuaca, A[i].Jenis)
			}
		}
	case 4:
		fmt.Print("Masukkan Pakaian dengan Kondisi Cuaca yang Ingin dicari(Panas/Hujan/Dingin): ")
		fmt.Scan(&cuaca)
		for i = 1; i < n; i++ {
			if A[i].Cuaca == cuaca {
				fmt.Printf("Nama Pakaian: %s, Kategori: %s, Warna: %s, Formalitas: %d, Cuaca: %s, Jenis: %s\n", A[i].Nama, A[i].Kategori, A[i].Warna, A[i].Formalitas, A[i].Cuaca, A[i].Jenis)
			}
		}
	}
}

func ubahPakaian(A *tabPakaian, n int) {
	var pilihan int
	var ID int
	var nama string
	var kategori string
	var warna string
	var tingkatFormalitas int
	var cuaca string
	var jenis string

	fmt.Println("Pilih ID Pakaian yang ingin diubah: ")
	fmt.Scan(&ID)
	fmt.Println("1. Edit Pakaian Bedasarkan Nama")
	fmt.Println("2. Edit Pakaian Bedasarkan Kategori")
	fmt.Println("3. Edit Pakaian Bedasarkan Warna")
	fmt.Println("4. Edit Pakaian Bedasarkan Cuaca")

	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		fmt.Print("Masukkan Nama Pakaian yang Baru(Kaos/Jacket/Celana/dll): ")
		fmt.Scan(&nama)
		A[ID].Nama = nama
	case 2:
		fmt.Print("Masukkan Nama Kategori yang Baru(Casual/Formal/Sporty): ")
		fmt.Scan(&kategori)
		A[ID].Kategori = kategori
	case 3:
		fmt.Print("Masukkan Warna Pakaian yang Baru(Hitam/Putih/Biru/Pink/dll): ")
		fmt.Scan(&warna)
		A[ID].Warna = warna
	case 4:
		fmt.Print("Masukkan Pakaian dengan Tingkat Formalitas terbaru(1-5): ")
		fmt.Scan(&tingkatFormalitas)
		A[ID].Formalitas = tingkatFormalitas
	case 5:
		fmt.Print("Masukkan Pakaian dengan Kondisi Cuaca terbaru(Panas/Dingin/Hujan/dll): ")
		fmt.Scan(&cuaca)
		A[ID].Cuaca = cuaca
	case 6:
		fmt.Print("Masukkan Pakaian dengan jenis terbaru: ")
		fmt.Scan(&jenis)
		A[ID].Jenis = jenis
	}
}

func hapusData(A *tabPakaian, n *int) {
	var idx int
	tampilkanData(*A, *n)
	fmt.Print("Pilih data Pakaian ke- berapa yang ingin dihapus: ")
	fmt.Scan(&idx)
	for i := idx; i < *n; i++ {
		A[i].Nama = A[i+1].Nama
		A[i].Kategori = A[i+1].Kategori
		A[i].Warna = A[i+1].Warna
		A[i].Formalitas = A[i+1].Formalitas
		A[i].Jenis = A[i+1].Jenis
	}
	*n -= 1
}
