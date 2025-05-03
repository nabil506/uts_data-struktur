package model

import(
	"uts_data/node"
)

func Create(nama_barang,produksi,merk string, jumlah_barang , id, harga int){
	datanya := node.ProdukCounterHandphone{
		ID_Barang: id,
		Nama_barang: nama_barang,
		Jumlah_barang: jumlah_barang,
		Harga_Barang: harga,
		Detail: node.DetailBarang{
			Production_date: produksi,
			Merk: merk,
		},
	}
	 CreateBarang(datanya)
}