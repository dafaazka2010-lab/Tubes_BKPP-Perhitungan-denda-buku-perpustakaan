package main

import "fmt"

func main() {
	lanjut := true

	for lanjut {
		var nama, judul, kategori string
		var jumlahBuku, hari int
		var inputJumlah, inputHari, jawaban string
		var valid bool
		totalDenda := 0

		fmt.Println()
		fmt.Println("==============================")
		fmt.Println("       DATA PEMINJAMAN BUKU   ")
		fmt.Println("==============================")
		fmt.Println("========= Data Peminjaman =========")
		fmt.Println()

		// INPUT NAMA
		fmt.Print("Masukkan nama peminjam: ")
		fmt.Scan(&nama)

		fmt.Printf("\n============ Data Buku ============\n")

		// ===============================
		// INPUT JUMLAH BUKU
		// ===============================
		valid = false
		for valid == false {
			fmt.Print("Masukkan jumlah buku (1-10): ")
			fmt.Scan(&inputJumlah)

			// CEK TANDA MINUS
			if inputJumlah[0] == '-' {
				fmt.Println("Jumlah buku tidak boleh negatif")
				fmt.Println()
			} else {
				valid = true
				for i := 0; i < len(inputJumlah); i++ {
					if inputJumlah[i] < '0' || inputJumlah[i] > '9' {
						valid = false
					}
				}

				if valid == false {
					fmt.Println("Masukkan jumlah buku dalam bentuk ANGKA (contoh: 2)")
				} else {
					fmt.Sscan(inputJumlah, &jumlahBuku)
					if jumlahBuku < 1 || jumlahBuku > 10 {
						fmt.Println("Jumlah buku harus antara 1 sampai 10")
						valid = false
					}
				}
				fmt.Println()
			}
		}

		// ===============================
		// INPUT HARI KETERLAMBATAN
		// ===============================
		valid = false
		for valid == false {
			fmt.Print("Masukkan hari keterlambatan: ")
			fmt.Scan(&inputHari)

			// CEK TANDA MINUS
			if inputHari[0] == '-' {
				fmt.Println("Hari keterlambatan tidak boleh negatif")
				fmt.Println()
			} else {
				valid = true
				for i := 0; i < len(inputHari); i++ {
					if inputHari[i] < '0' || inputHari[i] > '9' {
						valid = false
					}
				}

				if valid == false {
					fmt.Println("Masukkan hari dalam bentuk ANGKA (contoh: 2)")
				} else {
					fmt.Sscan(inputHari, &hari)
				}
				fmt.Println()
			}
		}

		// ===============================
		// INPUT DATA BUKU
		// ===============================
		for i := 1; i <= jumlahBuku; i++ {
			fmt.Println("\nBuku ke-", i)
			fmt.Println("Catatan: Judul buku TANPA spasi")

			fmt.Print("Judul buku: ")
			fmt.Scan(&judul)

			fmt.Print("Kategori (pelajaran/novel/langka): ")
			fmt.Scan(&kategori)

			for kategori != "pelajaran" &&
				kategori != "novel" &&
				kategori != "langka" {

				fmt.Println("Kategori tidak valid!")
				fmt.Println("Gunakan huruf kecil semua (contoh: novel)")
				fmt.Print("Kategori (pelajaran/novel/langka): ")
				fmt.Scan(&kategori)
			}

			if kategori == "pelajaran" {
				totalDenda += 1000 * hari
			} else if kategori == "novel" {
				totalDenda += 2000 * hari
			} else if kategori == "langka" {
				totalDenda += 5000 * hari
			} else {
				fmt.Println("kategori tidak dikenal")
			}
		}

		// ===============================
		// OUTPUT
		// ===============================
		fmt.Println("\n===== HASIL =====")
		fmt.Println("Nama peminjam   :", nama)
		fmt.Println("Jumlah buku     :", jumlahBuku)
		fmt.Println("Hari terlambat  :", hari)
		fmt.Println("Total denda     : Rp", totalDenda)

		fmt.Print("\nLanjutkan? (true/false): ")
		fmt.Scan(&jawaban)
		if jawaban != "true" {
			lanjut = false
			fmt.Println("Terimakasih")
		}
	}
}
