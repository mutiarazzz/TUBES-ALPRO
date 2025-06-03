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
		// Tampilan menu kotak rapi
		fmt.Println("╔════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                        OOTD PLANNER - MENU                         ║")
		fmt.Println("╠════════════════════════════════════════════════════════════════════╣")
		fmt.Println("║ →  1. Tambah Data Pakaian                                          ║")
		fmt.Println("║ →  2. Tampilkan Semua Pakaian                                      ║")
		fmt.Println("║ →  3. Ubah Data Pakaian                                            ║")
		fmt.Println("║ →  4. Hapus Data Pakaian                                           ║")
		fmt.Println("║ →  5. Cari Pakaian                                                 ║")
		fmt.Println("║ →  6. Rekomendasi Outfit                                           ║")
		fmt.Println("║ →  7. Urutkan Berdasarkan Formalitas                               ║")
		fmt.Println("║ →  8. Urutkan Berdasarkan Tanggal Terakhir Dipakai                 ║")
		fmt.Println("║ →  9. Lihat Semua Kombinasi Outfit                                 ║")
		fmt.Println("║ → 10. Keluar                                                       ║")
		fmt.Println("╚════════════════════════════════════════════════════════════════════╝")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahData(&data, &nData)
		case 2:
			tampilkanData(data, nData)
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
			fmt.Println("\n╔════════════════════════════════════════════════════╗")
			fmt.Println("║        💖 Terima kasih telah menggunakan 💖        ║")
			fmt.Println("║                🌟 OOTD Planner 🌟                  ║")
			fmt.Println("║     Semoga harimu penuh gaya dan percaya diri!     ║")
			fmt.Println("╚════════════════════════════════════════════════════╝")

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
	fmt.Println("\n╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║                ✨ MENU TAMBAH DATA PAKAIAN ✨            ║")
	fmt.Println("╠══════════════════════════════════════════════════════════╣")
	fmt.Println("║ Silakan lengkapi informasi berikut tentang pakaian Anda  ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")

	fmt.Print("🧥 Nama Pakaian: ")
	fmt.Scan(&A[*n].Nama)

	fmt.Print("📂 Kategori (misal: Casual, Formal, Sporty): ")
	fmt.Scan(&A[*n].Kategori)

	fmt.Print("🎨 Warna Pakaian: ")
	fmt.Scan(&A[*n].Warna)

	fmt.Print("📈 Tingkat Formalitas (misal: 1-5): ")
	fmt.Scan(&A[*n].Formalitas)

	fmt.Print("🌤️  Cuaca yang Cocok (Hujan/Panas/Dingin): ")
	fmt.Scan(&A[*n].Cuaca)

	fmt.Print("👕 Jenis (Atasan/Bawahan/Footwear): ")
	fmt.Scan(&A[*n].Jenis)

	fmt.Print("📅 Tanggal Terakhir Dipakai (YYYY-MM-DD): ")
	fmt.Scan(&A[*n].TerakhirDipakai)

	*n++
	fmt.Println("\n✅ Data pakaian berhasil ditambahkan!")
}


// Tampilkan Semua Pakaian//
// Menampilkan seluruh data pakaian yang telah dimasukkan//
// Tujuan: Menampilkan seluruh isi array pakaian.
func tampilkanData(A tabPakaian, n int) {
	fmt.Println("\n╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║                   📋 DATA SEMUA PAKAIAN                  ║")
	fmt.Println("╠══════════════════════════════════════════════════════════╣")
	fmt.Println("║ Berikut adalah daftar lengkap pakaian yang tersimpan     ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	for i := 1; i < n; i++ {
		fmt.Printf("%d. Nama: %s | Kategori: %s | Warna: %s | Formalitas: %d | Cuaca: %s | Jenis: %s | Terakhir Dipakai: %s\n",
			i, A[i].Nama, A[i].Kategori, A[i].Warna, A[i].Formalitas, A[i].Cuaca, A[i].Jenis, A[i].TerakhirDipakai)
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
	var cuaca string
	var jenis string

fmt.Println("\n╔═══════════════════════════════════════════════╗")
fmt.Println("║       ✨ PILIH DATA YANG INGIN DIUBAH ✨      ║")
fmt.Println("╚═══════════════════════════════════════════════╝")
fmt.Println() 
fmt.Println("Pilih ID Pakaian yang ingin diubah: ")
fmt.Scan(&ID)
fmt.Println("1. Edit Pakaian Bedasarkan Nama")
fmt.Println("2. Edit Pakaian Bedasarkan Kategori")
fmt.Println("3. Edit Pakaian Bedasarkan Warna")
fmt.Println("4. Edit Pakaian Bedasarkan Jenis")
fmt.Println("5. Edit Pakaian Bedasarkan Cuaca")

	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		fmt.Print("Masukkan Nama Pakaian yang Baru: ")
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
		fmt.Print("Masukkan Jenis Pakaian yang baru: ")
		fmt.Scan(&jenis)
		A[ID].Jenis = jenis
	case 5:
		fmt.Print("Masukkan Pakaian dengan Kondisi Cuaca terbaru(Panas/Dingin/Hujan/dll): ")
		fmt.Scan(&cuaca)
		A[ID].Cuaca = cuaca
	}
}

// Hapus Data Outfit//
// Menghapus data berdasarkan indeks yang dipilih user//
func hapusData(A *tabPakaian, n *int) {
	var idx int
	tampilkanData(*A, *n)
	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║             🗑️ HAPUS DATA PAKAIAN         ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()
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

// Cari Pakaian
// Mencari pakaian berdasarkan Nama, Kategori, Warna, atau Cuaca
// Sequential Search: Mencari satu per satu.
func cariPakaian(A tabPakaian, n int) {
	var pilihan int
	var nama string
	var kategori string
	var warna string
	var formalitas int
	var jenis string
	var i int

fmt.Println("\n╔═══════════════════════════════════════════════╗")
fmt.Println("║               🔍 MENU CARI PAKAIAN            ║")
fmt.Println("╠═══════════════════════════════════════════════╣")
fmt.Println("║ 1. Cari Pakaian Berdasarkan Nama              ║")
fmt.Println("║ 2. Cari Pakaian Berdasarkan Kategori          ║")
fmt.Println("║ 3. Cari Pakaian Berdasarkan Warna             ║")
fmt.Println("║ 4. Cari Pakaian Berdasarkan Tingkat Formalitas║")
fmt.Println("║ 5. Cari Pakaian Berdasarkan Jenis             ║")
fmt.Println("╚═══════════════════════════════════════════════╝")
fmt.Println()
fmt.Print("Apa yang ingin anda cari? ")

	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		fmt.Print("Masukkan Nama Pakaian yang Ingin dicari: ")
		fmt.Scan(&nama)
		for i = 1; i < n; i++ {
			if A[i].Nama == nama {
				fmt.Printf("Nama Pakaian: %s, Kategori: %s, Warna: %s, Formalitas: %d, Cuaca: %s, Jenis: %s\n", A[i].Nama, A[i].Kategori, A[i].Warna, A[i].Formalitas, A[i].Cuaca, A[i].Jenis)
			}
		}
	case 2:
		var pass, idx int
		var ya Pakaian
		var low, high, mid int
		pass = 2
		for pass < n {
			idx = pass - 1
			i = pass
			for i < n {
				if A[i].Kategori > A[idx].Kategori {
					idx = i
				}
				i += 1
			}
			ya = A[pass-1]
			A[pass-1] = A[idx]
			A[idx] = ya
			pass += 1
		}
		fmt.Print("Masukkan Nama Kategori yang Ingin dicari(Casual/Formal/Sporty): ")
		fmt.Scan(&kategori)
		low = 0
		high = n - 1

		for low <= high {
			mid = (low + high) / 2
			if A[mid].Kategori == kategori {
				fmt.Printf("Nama Pakaian: %s, Kategori: %s, Warna: %s, Formalitas: %d, Cuaca: %s, Jenis: %s\n", A[mid].Nama, A[mid].Kategori, A[mid].Warna, A[mid].Formalitas, A[mid].Cuaca, A[mid].Jenis) 
			}
			if A[mid].Kategori > kategori {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		if mid == -1 {
			fmt.Print("Data Tidak dapat ditemukan")
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
		fmt.Print("Masukkan Pakaian dengan Tingkat Formalitas yang anda inginkan(1-5): ")
		fmt.Scan(&formalitas)
		for i = 1; i < n; i++ {
			if A[i].Formalitas == formalitas {
				fmt.Printf("Nama Pakaian: %s, Kategori: %s, Warna: %s, Formalitas: %d, Cuaca: %s, Jenis: %s\n", A[i].Nama, A[i].Kategori, A[i].Warna, A[i].Formalitas, A[i].Cuaca, A[i].Jenis)
			}
		}
	case 5:
		fmt.Print("Masukkan Jenis Pakaian yang Ingin dicari(Atasan/Bawahan/Footwear): ")
		fmt.Scan(&jenis)
		for i = 1; i < n; i++ {
			if A[i].Jenis == jenis {
				fmt.Printf("Nama Pakaian: %s, Kategori: %s, Warna: %s, Formalitas: %d, Cuaca: %s, Jenis: %s\n", A[i].Nama, A[i].Kategori, A[i].Warna, A[i].Formalitas, A[i].Cuaca, A[i].Jenis)
			}
		}
	}
	fmt.Println()
}

// Rekomendasi Outfit//
// Memberi saran outfit berdasarkan cuaca atau jenis acara//
func rekomendasiOutfit(A tabPakaian, n int) {
	var pilihan int
	var cuaca string
	var acara string
	var i int

	fmt.Println("\n╔═══════════════════════════════════════════════════════╗")
	fmt.Println("║              👗 MENU REKOMENDASI OUTFIT               ║")
	fmt.Println("╠═══════════════════════════════════════════════════════╣")
	fmt.Println("║ 1. Rekomendasi Outfit Berdasarkan Cuaca               ║")
	fmt.Println("║ 2. Rekomendasi Outfit Berdasarkan Acara yang dihadiri ║")
	fmt.Println("╚═══════════════════════════════════════════════════════╝")
	fmt.Println()
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
		fmt.Print("Masukkan Acara yang Akan dihadiri(Rapat/Bermain/Olahraga/Pesta): ")
		fmt.Scan(&acara)
		fmt.Println("Outfit yang Cocok untuk Acara ", acara)
		for i = 1; i < n; i++ {
			if acara == "Rapat" && A[i].Formalitas >= 3 {
				fmt.Printf("Nama Pakaian: %s, Kategori: %s, Warna: %s, Formalitas: %d, Cuaca: %s, Jenis: %s\n", A[i].Nama, A[i].Kategori, A[i].Warna, A[i].Formalitas, A[i].Cuaca, A[i].Jenis)
			} else if acara == "Bermain" && A[i].Formalitas >= 2 && A[i].Formalitas <=3 {
				fmt.Printf("Nama Pakaian: %s, Kategori: %s, Warna: %s, Formalitas: %d, Cuaca: %s, Jenis: %s\n", A[i].Nama, A[i].Kategori, A[i].Warna, A[i].Formalitas, A[i].Cuaca, A[i].Jenis)
			} else if acara == "Olahraga" && A[i].Formalitas == 1 {
				fmt.Printf("Nama Pakaian: %s, Kategori: %s, Warna: %s, Formalitas: %d, Cuaca: %s, Jenis: %s\n", A[i].Nama, A[i].Kategori, A[i].Warna, A[i].Formalitas, A[i].Cuaca, A[i].Jenis)
			} else if acara == "Pesta" && A[i].Formalitas >= 4 {
				fmt.Printf("Nama Pakaian: %s, Kategori: %s, Warna: %s, Formalitas: %d, Cuaca: %s, Jenis: %s\n", A[i].Nama, A[i].Kategori, A[i].Warna, A[i].Formalitas, A[i].Cuaca, A[i].Jenis)
			} 
		}
	}
	fmt.Println()
}

// Mengurutkan pakaian berdasarkan tingkat Formalitas secara menaik (dari yang paling besar tingkat formalitasnya)
// Menampilkan semua pakaian yang ada berdasarkan tingkat formalitas
// Menggunakan algoritma Selection Sort
func sortFormalitas(A tabPakaian, n int) {
	var i, idx, pass int
	var ya Pakaian

	fmt.Println()
    fmt.Println("╔════════════════════════════════════════════════════════════╗")
    fmt.Println("║              🔼 Pakaian dari yang Paling Formal 🔽         ║")
    fmt.Println("╠════════════════════════════════════════════════════════════╣")
    fmt.Println("║       Lihat daftar pakaian dengan tingkat formalitas       ║")
    fmt.Println("║          mulai dari yang tertinggi ke yang terendah        ║")
    fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Println()

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
	fmt.Println()
}

// Mengurutkan pakaian berdasarkan kapan pengguna memakai pakaian tersebut
// Menampilkan semua pakaian yang ada berdasarkan terakhir digunakan
// Menggunakan Algoritma : Insertion Sort
func sortByTanggalTerakhirDipakai(A tabPakaian, n int) {
	fmt.Println("\n╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║     📆  Lihat Outfit dari yang Terbaru Kamu Pakai!         ║")
	fmt.Println("║        Pakaian diurutkan berdasarkan tanggal terakhir      ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Println()

	for i := 2; i <= n; i++ {
		temp := A[i]
		j := i - 1
		// Urutkan secara descending: yang lebih baru (lebih besar) di depan
		for j >= 1 && A[j].TerakhirDipakai < temp.TerakhirDipakai {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = temp
	}
	fmt.Println("Data terurut dari pakaian yang terakhir digunakan: ")
	tampilkanData(A, n)
	fmt.Println()
}

// Menampilkan semua kombinasi atasan, bawahan, dan footwear
func kombinasiOutfit(A tabPakaian, n int) {
	var count int = 1

	fmt.Println()
	fmt.Println("╔═══════════════════════════════════════════════════════════╗")
	fmt.Println("║           👗 KOMBINASI OUTFIT YANG COCOK BUAT KAMU        ║")
	fmt.Println("╠═══════════════════════════════════════════════════════════╣")
	fmt.Println("║ Semua Kombinasi Outfit (Atasan + Bawahan + Footwear)      ║")
	fmt.Println("║               dari data pakaian yang tersedia             ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════╝")

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
		fmt.Println("Belum ada kombinasi yang bisa ditampilkan (Tambahkan data pakaian).")
	}
	fmt.Println()
}

