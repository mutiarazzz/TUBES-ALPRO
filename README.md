👗 OOTD Planner - Lemari Digital Pintar

Aplikasi ini dirancang untuk membantu pengguna dalam mengatur pakaian yang mereka miliki agar lebih rapi, efisien, dan mudah digunakan sehari-hari. Dapat disebut juga sebagai **lemari digital**. Dengan aplikasi ini, pengguna bisa menyimpan data tentang setiap pakaian yang dimiliki, mulai dari nama, jenis, warna, hingga tingkat formalitasnya.

Selain itu, aplikasi ini juga menyediakan fitur untuk membuat kombinasi outfit yang tersedia. Pengguna bisa mendapatkan rekomendasi pakaian berdasarkan **cuaca** atau **acara** yang akan dihadiri. Melalui fitur-fitur seperti **tambah pakaian**, **pencarian**, **rekomendasi outfit**, hingga **pengurutan berdasarkan tanggal terakhir dipakai**, aplikasi ini membantu pengguna membuat keputusan berpakaian yang tepat tanpa membuang waktu.

---

✨ Fitur Utama dan Penjelasannya

1. Menu Utama

* Titik awal program OOTD Planner dan pengelola alur interaksi pengguna.
* Menampilkan daftar menu fitur, memungkinkan pengguna memilih dan menavigasi aplikasi.

**Tampilan Menu mencakup:**

* Daftar pilihan fitur
* Petunjuk input angka
* Akses ke semua fitur aplikasi

**Fungsi terkait:** `main()`

**Proses Input:**

1. Pengguna menjalankan aplikasi
2. Melihat daftar menu
3. Memilih fitur dengan mengetik angka
4. Aplikasi menjalankan fitur tersebut dan kembali ke menu utama
5. Pilih keluar jika ingin mengakhiri aplikasi

---

2. Tambah Pakaian

Memungkinkan pengguna menambahkan pakaian baru ke sistem.

**Input Data:**

* Nama pakaian
* Kategori (Casual, Formal, dll.)
* Warna
* Tingkat formalitas (1–5)
* Cuaca
* Jenis (Atasan, Bawahan, Footwear)
* Tanggal terakhir dipakai (YYYY-MM-DD)

**Fungsi terkait:** `tambahPakaian()`

**Langkah:**

1. Pilih menu Tambah Pakaian
2. Masukkan data pakaian satu per satu

---

3. Tampilkan Semua Pakaian

Menampilkan seluruh daftar pakaian yang tersimpan secara rapi.

**Fitur ini mencakup:**

* Penomoran indeks
* Semua atribut pakaian (nama, kategori, dll.)

**Fungsi terkait:** `tampilPakaian()`

**Langkah:** Pilih opsi Tampilkan Semua Pakaian dari menu

---

4. Ubah Data Pakaian

Memungkinkan pengguna memperbarui atribut pakaian tertentu.

**Langkah-langkah:**

1. Pilih opsi Ubah Data Pakaian
2. Pilih indeks/nama pakaian
3. Pilih atribut yang ingin diubah
4. Masukkan data baru

**Fungsi terkait:** `ubahPakaian()`

---

5. Hapus Pakaian

Menghapus satu data pakaian dari sistem.

**Fitur ini mencakup:**

* Pemilihan indeks/nama
* Penghapusan dari array
* Penyesuaian jumlah data (n)

**Fungsi terkait:** `hapusPakaian()`

**Langkah:** Pilih menu Hapus Pakaian, masukkan indeks/nama

---

6. Cari Pakaian

Pencarian berdasarkan atribut pakaian seperti nama, jenis, warna, dll.

**Pencarian dilakukan dengan:**

* Sequential Search (untuk atribut umum)
* Binary Search (untuk kategori, setelah diurutkan)

**Fungsi terkait:** `cariPakaian()`

**Langkah:**

1. Pilih menu Cari Pakaian
2. Masukkan kategori pencarian
3. Masukkan keyword pencarian

---

7. Rekomendasi Outfit

Memberikan saran outfit berdasarkan:

* **Cuaca saat ini**
* **Jenis acara** (rapat, pesta, bermain, olahraga)

**Fungsi terkait:** `rekomendasiOutfit()`

**Langkah:**

1. Pilih menu Rekomendasi Outfit
2. Input cuaca dan jenis acara
3. Sistem menampilkan outfit yang cocok

---

8. Urut Berdasarkan Formalitas

Mengurutkan daftar pakaian berdasarkan tingkat formalitas (tinggi ke rendah).

**Algoritma:** Selection Sort
**Fungsi terkait:** `sortFormalitas()`

**Langkah:** Pilih menu Urut Berdasarkan Formalitas

---

9. Urut Berdasarkan Tanggal Terakhir Dipakai

Mengurutkan pakaian dari tanggal terakhir digunakan (baru → lama).

**Algoritma:** Insertion Sort
**Fungsi terkait:** `sortByTanggalTerakhirDipakai()`

**Langkah:** Pilih menu Urut Berdasarkan Tanggal Dipakai

---

10. Kombinasi Outfit

Menampilkan semua kemungkinan kombinasi dari:

* Atasan
* Bawahan
* Footwear

**Fungsi terkait:** `kombinasiOutfit()`

**Langkah:** Pilih menu Lihat Semua Kombinasi Outfit

