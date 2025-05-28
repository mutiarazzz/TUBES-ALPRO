package main

import (
	"fmt"
)

type Pakaian struct {
	Nama            string
	Kategori        string
	Warna           string
	Formalitas      int
	Cuaca           string
	Jenis           string
	TerakhirDipakai string
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
		fmt.Println("2. Tambah Data Pakaian")
		fmt.Println("3. Ubah Data Pakaian")
		fmt.Println("4. Hapus Data Pakaian")
		fmt.Println("5. Cari Pakaian")
		fmt.Println("6. Rekomendasi Outfit")
		fmt.Println("7. Urutkan Berdasarkan Formalitas")
		fmt.Println("8. Urutkan Berdasarkan Tanggal Terakhir Dipakai")
		fmt.Println("9. Lihat Semua Kombinasi Outfit")
		fmt.Println("10. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tampilkanData(data, nData)
		case 2:
			tambahData(&data, &nData)
		case 3:
			ubahPakaian(&data, nData)
		case 4:
			hapusData(&data, &nData)
		case 5:
			cariPakaian(data, nData)
		case 6:
			rekomendasiOutfit(data, nData)
		case 7:
			sortFormalitas(data, nData)
		case 8:
			sortByTanggalTerakhirDipakai(data, nData)
		case 9:
			kombinasiOutfit(data, nData)
		case 10:
			fmt.Println("Terima kasih telah menggunakan OOTD Planner!")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

// Tambah Data Pakaian
// Pengguna memasukkan informasi pakaian baru//
// Tujuan: Menyimpan data pakaian ke array
func tambahData(A *tabPakaian, n *int) {
	fmt.Print("Silahkan Masukkan Nama Pakaian: ")
	fmt.Scan(&A[*n].Nama)
	fmt.Print("Silahkan Masukkan Kategori Pakaian: ")
	fmt.Scan(&A[*n].Kategori)
	fmt.Print("Silahkan Masukkan Warna Pakaian yang dimiliki: ")
	fmt.Scan(&A[*n].Warna)
	fmt.Print("Silahkan Masukkan Tingkat Formalitas dari Pakaian: ")
	fmt.Scan(&A[*n].Formalitas)
	fmt.Print("Silahkan Masukkan Cuaca yang Cocok untuk Pakaian Anda(Hujan/Panas/Dingin): ")
	fmt.Scan(&A[*n].Cuaca)
	fmt.Print("Silahkan Masukkan Jenis Pakaian(Atasan/Bawahan/Footwear): ")
	fmt.Scan(&A[*n].Jenis)
	fmt.Print("Tanggal Terakhir Dipakai (YYYY-MM-DD): ")
	fmt.Scan(&A[*n].TerakhirDipakai)
	*n++

}

// Tampilkan Semua Pakaian//
// Menampilkan seluruh data pakaian yang telah dimasukkan//
// Tujuan: Menampilkan seluruh isi array pakaian.
func tampilkanData(A tabPakaian, n int) {
	for i := 1; i < n; i++ {
		fmt.Printf("%d. Nama: %s | Kategori: %s | Warna: %s | Formalitas: %d | Cuaca: %s | Jenis: %s | Terakhir Dipakai: %s\n",
			i, A[i].Nama, A[i].Kategori, A[i].Warna, A[i].Formalitas, A[i].Cuaca, A[i].Jenis, A[i].TerakhirDipakai)
	}
}

// Rekomendasi Outfit//
// Memberi saran outfit berdasarkan cuaca atau jenis acara//
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

// Cari Pakaian
// Mencari pakaian berdasarkan Nama, Kategori, Warna, atau Cuaca
// Sequential Search: Mencari satu per satu.
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

// Edit / Ubah Data Pakaian
// Mengubah atribut pakaian pada indeks tertentu berdasarkan pilihan user
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

// Hapus Data Outfit//
// Menghapus data berdasarkan indeks yang dipilih user//
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

func kombinasiOutfit(A tabPakaian, n int) {
	fmt.Println("\n--- Semua Kombinasi Outfit (Atasan + Bawahan + Footwear) ---")
	var count int = 1
	for i := 1; i < n; i++ {
		if A[i].Jenis == "Atasan" {
			for j := 1; j < n; j++ {
				if A[j].Jenis == "Bawahan" {
					for k := 1; k < n; k++ {
						if A[k].Jenis == "Footwear" {
							fmt.Printf("\nKombinasi %d:\n", count)
							fmt.Printf("  Atasan  : %s (%s)\n", A[i].Nama, A[i].Warna)
							fmt.Printf("  Bawahan : %s (%s)\n", A[j].Nama, A[j].Warna)
							fmt.Printf("  Footwear: %s (%s)\n", A[k].Nama, A[k].Warna)
							count++
						}
					}
				}
			}
		}
	}
	if count == 1 {
		fmt.Println("Belum ada kombinasi yang bisa ditampilkan (cek data jenis pakaian).")
	}
}

// Mengurutkan pakaian berdasarkan tingkat formalitas
// Menampilkan semua pakaian yang ada berdasarkan tingkat formalitas
func sortFormalitas(A tabPakaian, n int) {
	var i, idx, pass int
	var ya Pakaian

	pass = 2
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if A[i].Formalitas > A[idx].Formalitas {
				idx = i
			}
			i += 1
		}
		ya = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = ya
		pass += 1
	}
	fmt.Println("Data terurut dari formalitas terbesar sampai terendah: ")
	tampilkanData(A, n)
}

// Menampilkan semua kombinasi atasan, bawahan, dan footwear
func sortByTanggalTerakhirDipakai(A tabPakaian, n int) {
	var i, j int
	var temp Pakaian
	for i = 1; i < n-1; i++ {
		for j = i + 1; j < n; j++ {
			if A[i].TerakhirDipakai < A[j].TerakhirDipakai { // string comparison
				temp = A[i]
				A[i] = A[j]
				A[j] = temp
			}
		}
	}
	fmt.Println("Data terurut dari yang paling baru dipakai:")
	tampilkanData(A, n)
}
