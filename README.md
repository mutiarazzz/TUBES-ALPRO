👕 OOTD Planner: Lemari Digital Anda!
Aplikasi OOTD Planner dirancang untuk membantu Anda mengatur koleksi pakaian agar lebih rapi, efisien, dan mudah digunakan sehari-hari. Anggap saja ini adalah lemari digital pribadi Anda! Dengan aplikasi ini, Anda bisa menyimpan data lengkap tentang setiap pakaian, mulai dari nama, jenis, warna, hingga tingkat formalitasnya.

Tidak hanya itu, OOTD Planner juga menyediakan fitur canggih untuk membuat kombinasi outfit yang keren. Anda bisa mendapatkan rekomendasi pakaian berdasarkan cuaca atau acara yang akan dihadiri. Melalui fitur-fitur seperti tambah pakaian, pencarian, rekomendasi outfit, hingga pengurutan berdasarkan tanggal terakhir dipakai, aplikasi ini akan membantu Anda membuat keputusan berpakaian yang tepat dan menghemat waktu setiap harinya.

✨ Fitur - fitur yang ada pada aplikasi ini
Berikut adalah penjabaran mendetail dari setiap fitur beserta cara kerjanya:

1. Menu Utama
Sebagai titik awal eksekusi program, Fitur Menu mengelola seluruh alur interaksi pengguna dengan aplikasi. Sebuah perulangan akan terus menampilkan pilihan fitur kepada pengguna hingga Anda memutuskan untuk keluar. Setiap iterasi menampilkan daftar menu dengan berbagai fitur, mulai dari penambahan, penampilan, pengubahan, penghapusan data pakaian, pencarian, rekomendasi outfit, pengurutan, hingga melihat semua kombinasi outfit.

Menampilkan: Daftar pilihan menu utama dan petunjuk untuk memilih menu dengan mengetik angka.
Akses ke: Berbagai fitur aplikasi.
Fungsi Terkait: main()
Proses Input:
Jalankan aplikasi.
Lihat daftar menu yang muncul.
Pilih satu fitur dengan memasukkan angka.
Aplikasi menjalankan fitur yang dipilih, lalu kembali ke menu utama.
Pilih opsi keluar jika sudah selesai.

2. Tambah Pakaian
Fungsi tambahPakaian() memungkinkan Anda memasukkan data pakaian baru ke dalam sistem. Prosesnya dilakukan melalui input bertahap untuk atribut seperti nama, kategori (atasan, bawahan, footwear), warna, cuaca yang sesuai, jenis pakaian, tingkat formalitas, serta tanggal terakhir digunakan. Semua data akan disimpan dalam struktur array dan jumlah data akan bertambah otomatis.

Memungkinkan: Menambahkan pakaian baru ke sistem.
Mencakup: Input nama pakaian, kategori, warna, tingkat formalitas (1-5), cuaca yang sesuai, jenis pakaian (Atasan, Bawahan, Footwear), dan tanggal terakhir dipakai (format: YYYY-MM-DD).
Fungsi Terkait: tambahPakaian()
Proses Input:
Akses dari menu utama, pilih opsi satu (Tambah Pakaian).
Masukkan seluruh data pakaian sesuai urutan.

3. Tampil Pakaian
Melalui fungsi tampilPakaian(), program menampilkan seluruh daftar pakaian yang tersimpan dalam array. Data ditampilkan dalam bentuk tabel atau list yang rapi, menunjukkan semua atribut masing-masing pakaian. Fitur ini bertujuan untuk membantu Anda mengetahui isi lemari virtual Anda dengan lebih mudah.

Menampilkan: Seluruh koleksi pakaian yang telah dimasukkan.
Mencakup: Semua atribut pakaian dan penomoran indeks untuk identifikasi item.
Fungsi Terkait: tampilkanPakaian()
Proses Input: Akses dari menu utama, pilih opsi dua (Tampilkan Semua Pakaian).

4. Ubah Pakaian
Fungsi ubahPakaian() memberikan fleksibilitas untuk memperbarui informasi pakaian yang telah tersimpan. Anda cukup memasukkan nama pakaian yang ingin diubah, sistem akan mencari dan menampilkan data lama. Setelah itu, Anda dapat memilih atribut mana yang ingin diperbarui dan memasukkan data baru, menjaga informasi pakaian tetap relevan dan terkini.

Memungkinkan: Mengubah data dari pakaian tertentu.
Mencakup: Pemilihan pakaian berdasarkan indeks, pemilihan atribut yang ingin diubah (nama, warna, jenis, dll.), dan memasukkan data baru.
Fungsi Terkait: ubahPakaian()
Proses Input:
Akses dari menu utama, pilih opsi tiga (Ubah Data Pakaian).
Pilih indeks pakaian.
Pilih atribut yang ingin diubah.
Masukkan data baru.

5. Hapus Pakaian
Fitur hapusPakaian() berfungsi untuk menghapus salah satu pakaian dari data. Anda cukup memasukkan nama pakaian yang ingin dihapus. Jika ditemukan, sistem akan menghapusnya dengan menggeser elemen array setelahnya ke kiri, dan jumlah total data pakaian akan berkurang secara otomatis.

Memungkinkan: Menghapus pakaian dari daftar.
Mencakup: Pemilihan indeks pakaian, penghapusan item dari array, dan penyesuaian jumlah data.
Fungsi Terkait: hapusPakaian()
Proses Input:
Akses dari menu utama, pilih opsi empat (Hapus Data Pakaian).
Pilih indeks pakaian yang ingin dihapus.

6. Cari Pakaian
Fungsi cariPakaian() memungkinkan Anda mencari pakaian berdasarkan atribut tertentu, seperti nama, kategori, warna, jenis, atau tingkat formalitas. Pencarian menggunakan sequential search untuk data tidak terurut dan binary search untuk kategori (setelah diurutkan), membuat pencarian lebih cepat dan efisien.

Memfasilitasi: Pencarian pakaian berdasarkan kata kunci dan kategori.
Mencakup: Input kategori pencarian (nama, warna, jenis, cuaca, dll.) dan kata kunci pencarian.
Fungsi Terkait: cariPakaian()
Proses Input:
Akses dari menu utama, pilih opsi lima (Cari Pakaian).
Masukkan kategori pencarian.
Masukkan kata kunci pencarian.

7. Rekomendasi Outfit
Fitur rekomendasiOutfit() berfungsi untuk memberikan saran outfit berdasarkan dua parameter utama: cuaca dan jenis acara. Sistem akan memfilter pakaian yang sesuai dengan cuaca, lalu mencocokkan tingkat formalitas pakaian dengan standar minimal formalitas dari acara tertentu (misalnya, rapat membutuhkan formalitas ≥4, bermain ≥1, dan pesta/pernikahan ≥5).

Memberikan: Saran outfit berdasarkan cuaca dan jenis acara.
Mencakup: Input cuaca saat ini dan jenis acara (rapat, pesta, main, dll.).
Fungsi Terkait: rekomendasiOOTD()
Proses Input:
Akses dari menu utama, pilih opsi enam (Rekomendasi Outfit).
Input cuaca dan jenis acara.

8. Urut Berdasarkan Formalitas
Fungsi sortFormalitas() mengurutkan seluruh data pakaian berdasarkan nilai formalitas dari terendah hingga tertinggi menggunakan algoritma selection sort. Dengan fitur ini, Anda dapat melihat daftar pakaian secara berurutan berdasarkan tingkat formalitas.

Menampilkan: Hasil setelah diurutkan.
Fungsi Terkait: selectionSortFormalitas()
Proses Input: Akses dari menu utama, pilih opsi tujuh (Urutkan Berdasarkan Formalitas).

9. Urut Berdasarkan Tanggal Terakhir Dipakai
Fitur sortTanggal() digunakan untuk menyusun pakaian berdasarkan tanggal terakhir digunakan, dari yang paling baru hingga yang paling lama, menggunakan algoritma insertion sort. Ini cocok untuk data yang sebagian besar sudah terurut. Dengan melihat tanggal terakhir dipakai, Anda bisa lebih bijak dalam menggunakan pakaian secara bergiliran.

Mengurutkan: Pakaian berdasarkan tanggal terakhir dipakai dari yang terbaru ke terlama.
Menampilkan: Hasil setelah diurutkan.
Fungsi Terkait: insertionSortTanggal()
Proses Input: Akses dari menu utama, pilih opsi delapan (Urutkan Berdasarkan Tanggal Dipakai).

10. Kombinasi Outfit
Fungsi kombinasiOutfit() menyajikan semua kemungkinan kombinasi outfit yang terdiri dari tiga kategori utama: atasan, bawahan, dan alas kaki (footwear). Program akan melakukan iterasi bertingkat untuk setiap kategori dan menghasilkan seluruh kombinasi yang mungkin, sangat berguna bagi Anda yang ingin melihat variasi gaya dari koleksi pakaian Anda.

Mencakup: Menggabungkan setiap atasan, bawahan, dan footwear yang ada, lalu menampilkan seluruh kombinasi yang tersedia.
Fungsi Terkait: kombinasiOutfit()
Proses Input: Akses dari menu utama, pilih opsi sembilan (Lihat Semua Kombinasi Outfit).
