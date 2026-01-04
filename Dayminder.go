package main

import "fmt"

const NMAX int = 100

type dayMinder struct {
	id, judul, kateg, date, status string
}

type Tabdayminder [NMAX]dayMinder

var datadayminder Tabdayminder
var N int

func main() {
	var pilih int
	for pilih != 6 {
		menu()
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			tambahData(&datadayminder, &N)
		case 2:
			ubahHapusMenu()
		case 3:
			cariMenu()
		case 4:
			urutMenu()
		case 5:
			cetakData(datadayminder, N)
		}
	}
	fmt.Println("Program Selesai.")
}

func menu() {
	fmt.Println("-------------------------------")
	fmt.Println("         Dayminder App    ")
	fmt.Println("-------------------------------")
	fmt.Println("Aplikasi Manajemen Tugas Harian")
	fmt.Println("-------------------------------")
	fmt.Println("1. Tambah Tugas")
	fmt.Println("2. Ubah / Hapus Tugas")
	fmt.Println("3. Cari Tugas")
	fmt.Println("4. Urutkan Tugas")
	fmt.Println("5. Tampilkan Semua")
	fmt.Println("6. Keluar")
	fmt.Println("-----------------------")
	fmt.Print("Pilih [1-6]: ")
}

func tambahData(A *Tabdayminder, n *int) {
	if *n < NMAX {
		fmt.Print("Masukan ID dengan angka yang unik: ")
		fmt.Scan(&A[*n].id)
		fmt.Println("NOTE:Jika ingin menggunakan spasi, ganti dengan underscore (_) atau gunakan (-)")
		fmt.Print("Masukan Judul (tanpa spasi): ")
		fmt.Scan(&A[*n].judul)
		fmt.Print("Masukan Kategori tugas: ")
		fmt.Scan(&A[*n].kateg)
		fmt.Print("Masukan Deadline tugas (YYYY-MM-DD): ")
		fmt.Scan(&A[*n].date)
		A[*n].status = "Belum"
		*n++
		fmt.Println("Berhasil disimpan.")
	} else {
		fmt.Println("Penyimpanan penuh.")
	}
}

func ubahHapusMenu() {
	var pilih int
	var id string

	fmt.Println("1. Ubah Data")
	fmt.Println("2. Hapus Data")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)

	fmt.Print("Masukkan ID Tugas: ")
	fmt.Scan(&id)

	if pilih == 1 {
		ubahData(&datadayminder, N, id)
	} else if pilih == 2 {
		hapusData(&datadayminder, &N, id)
	}
}

func ubahData(A *Tabdayminder, n int, id string) {
	var idx int = seqSearch(*A, n, id)

	if idx != -1 {
		fmt.Println("- Data Ditemukan, Masukkan Data Baru -")
		fmt.Print("Judul Baru: ")
		fmt.Scan(&A[idx].judul)
		fmt.Print("Kategori Baru: ")
		fmt.Scan(&A[idx].kateg)
		fmt.Print("Deadline Baru: ")
		fmt.Scan(&A[idx].date)
		fmt.Print("Status (Selesai/Belum): ")
		fmt.Scan(&A[idx].status)
		fmt.Println("Data berhasil diubah.")
	} else {
		fmt.Println("ID tidak ditemukan.")
	}
}

func hapusData(A *Tabdayminder, n *int, id string) {
	var idx, i int
	idx = seqSearch(*A, *n, id)

	if idx != -1 {
		for i = idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n--
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("ID tidak ditemukan.")
	}
}

func cetakData(A Tabdayminder, n int) {

	for i := 0; i <= n-1; i++ {
		fmt.Printf("ID: %s\n", A[i].id)
		fmt.Printf("JUDUL: %s\n", A[i].judul)
		fmt.Printf("KATEGORI: %s\n", A[i].kateg)
		fmt.Printf("DEADLINE: %s\n", A[i].date)
		fmt.Printf("STATUS: %s\n", A[i].status)
		fmt.Println()
	}
}

func seqSearch(A Tabdayminder, n int, id string) int {
	var i int = 0
	var idx int = -1
	var ketemu bool = false
	for i < n && !ketemu {
		if A[i].id == id {
			idx = i
			ketemu = true
		}
		i++
	}
	return idx
}

func cariMenu() {
	var jenis, key string
	var idx int
	fmt.Println("1. Cari Kategori (Sequential)")
	fmt.Println("2. Cari Deadline (Binary - Auto Sort)")
	fmt.Print("Pilih: ")
	fmt.Scan(&jenis)

	if jenis == "1" {
		fmt.Print("Masukkan Kategori: ")
		fmt.Scan(&key)
		cariKategori(datadayminder, N, key)
	} else {
		fmt.Print("Masukkan Deadline: ")
		fmt.Scan(&key)
		selectionSort(&datadayminder, N)
		idx = binarySearch(datadayminder, N, key)
		if idx != -1 {
			fmt.Printf("Ditemukan: %s - %s\n", datadayminder[idx].judul, datadayminder[idx].status)
		} else {
			fmt.Println("Data tidak ditemukan.")
		}
	}
}

func cariKategori(A Tabdayminder, n int, kateg string) {
	var ketemu bool = false
	var i int
	fmt.Println("Hasil Pencarian:")
	for i = 0; i < n; i++ {
		if A[i].kateg == kateg {
			fmt.Printf("- %s (%s)\n", A[i].judul, A[i].status)
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("Tidak ada data.")
	}
}

func binarySearch(A Tabdayminder, n int, date string) int {
	var kiri, kanan, tengah int
	var idx int = -1
	var ketemu bool = false

	kiri = 0
	kanan = n - 1

	for kiri <= kanan && !ketemu {
		tengah = (kiri + kanan) / 2
		if A[tengah].date == date {
			idx = tengah
			ketemu = true
		} else if date < A[tengah].date {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	return idx
}

func urutMenu() {
	var jenis int
	fmt.Println("1. Urut Deadline (Ascending - Selection Sort)")
	fmt.Println("2. Urut Status (Descending - Insertion Sort)")
	fmt.Print("Pilih: ")
	fmt.Scan(&jenis)

	if jenis == 1 {
		selectionSort(&datadayminder, N)
	} else {
		insertionSort(&datadayminder, N)
	}
	cetakData(datadayminder, N)
}

func selectionSort(A *Tabdayminder, n int) {
	var i, j, min int
	var temp dayMinder
	for i = 0; i < n-1; i++ {
		min = i
		for j = i + 1; j < n; j++ {
			if A[j].date < A[min].date {
				min = j
			}
		}
		temp = A[i]
		A[i] = A[min]
		A[min] = temp
	}
}

func insertionSort(A *Tabdayminder, n int) {
	var i, j int
	var temp dayMinder
	for i = 1; i < n; i++ {
		temp = A[i]
		j = i - 1
		for j >= 0 && A[j].status < temp.status {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = temp
	}
}
