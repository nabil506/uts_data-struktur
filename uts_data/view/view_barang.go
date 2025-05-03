package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"uts_data/model"

)

func MenuUtama() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Tampilkan Barang")
		fmt.Println("3. Update Barang")
		fmt.Println("4. Hapus Barang")
		fmt.Println("5. Search Barang")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) 

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input tidak valid, silakan masukkan angka.")
			continue
		}

		switch choice {
		case 1:
			fmt.Println("Anda memilih: Tambah Barang")
			Insert()
		case 2:
			fmt.Println("Anda memilih: Tampilkan Barang")
			Views()
		case 3:
			fmt.Println("Anda memilih: Update Barang")
			Update()
		case 4:
			fmt.Println("Anda memilih: Hapus Barang")
			Delete()
		case 5:
			fmt.Println("Anda memilih: Search Barang")
			Searching()
		case 6:
			fmt.Println("Terima kasih! Program selesai.")
			return 
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}

func Insert() {
	var nama_barang,merk,produk string
	var id,jumlah,harga int

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan ID Barang: ")
	idStr, _ := reader.ReadString('\n')
 	id, _ = strconv.Atoi(strings.TrimSpace(idStr))

	fmt.Print("Masukkan Nama Barang: ")
	nama_barang, _ = reader.ReadString('\n')
	nama_barang = strings.TrimSpace(nama_barang)

	fmt.Print("Masukkan Jumlah Barang: ")
	jumlahStr, _ := reader.ReadString('\n')
	jumlah, _ = strconv.Atoi(strings.TrimSpace(jumlahStr))

	fmt.Print("Masukkan Harga Barang: ")
	hargaStr, _ := reader.ReadString('\n')
	harga, _ = strconv.Atoi(strings.TrimSpace(hargaStr))

	fmt.Print("Masukkan Tanggal Produksi: ")
	produk , _ = reader.ReadString('\n')
	produk = strings.TrimSpace(produk)


	fmt.Print("Masukkan Merk Barang: ")
	merk, _ = reader.ReadString('\n')
	merk = strings.TrimSpace(merk)


	model.Create(nama_barang, produk, merk, jumlah, id, harga)


}

func Views() {
	fmt.Println("Daftar Barang")
	for i, emp := range model.ReadBarang() {
		fmt.Println("--- Barang ke -", i+1, " ---")
		fmt.Println("ID Barang\t\t  : ", emp.ID_Barang)
		fmt.Println("Nama Barang\t\t  : ", emp.Nama_barang)
		fmt.Println("Jumlah Barang\t\t  : ", emp.Jumlah_barang)
		fmt.Println("Harga Barang\t\t  : ", emp.Harga_Barang)
		fmt.Println("Tanggal Produksi\t  : ", emp.Detail.Production_date)
		fmt.Println("Merk Barang\t\t  : ", emp.Detail.Merk)
		fmt.Print()
	}
}

func Update() {
	var nama_barang,produk,merk string
	var id,jumlah,harga int

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan ID Barang: ")
	idStr, _ := reader.ReadString('\n')
 	id, _ = strconv.Atoi(strings.TrimSpace(idStr))

	fmt.Print("Masukkan Nama Barang: ")
	nama_barang, _ = reader.ReadString('\n')
	nama_barang = strings.TrimSpace(nama_barang)

	fmt.Print("Masukkan Jumlah Barang: ")
	jumlahStr, _ := reader.ReadString('\n')
	jumlah, _ = strconv.Atoi(strings.TrimSpace(jumlahStr))

	fmt.Print("Masukkan Harga Barang: ")
	hargaStr, _ := reader.ReadString('\n')
	harga, _ = strconv.Atoi(strings.TrimSpace(hargaStr))


	fmt.Print("Masukkan Tanggal Produksi: ")
	produk , _ = reader.ReadString('\n')
	produk = strings.TrimSpace(produk)


	fmt.Print("Masukkan Merk Barang: ")
	merk, _ = reader.ReadString('\n')
	merk = strings.TrimSpace(merk)
	
	model.Update(nama_barang, produk, merk, jumlah, id, harga)


}

func Delete() {
	var id int
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan ID Pegawai yang akan dihapus: ")
	fmt.Scan(&id)
	reader.ReadString('\n')

	model.Delete(id)
}


func Searching(){
	var nama string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan Nama Barang yang dicari: ")
	nama, _ = reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	search := model.SearchBarang(nama)
	if search != nil{
		fmt.Println("ID Barang\t\t  : ", search.ID_Barang)	
		fmt.Println("Nama Barang\t\t  : ", search.Nama_barang)
		fmt.Println("Jumlah Barang\t\t  : ", search.Jumlah_barang)
		fmt.Println("Harga Barang\t\t  : ", search.Harga_Barang)
		fmt.Println("Tanggal Produksi\t  : ", search.Detail.Production_date)
		fmt.Println("Merk Barang\t\t  : ", search.Detail.Merk)
}

}