package node

type DetailBarang struct{
	Production_date string
	Merk string

}

type ProdukCounterHandphone struct{
	Nama_barang string
	Jumlah_barang , Harga_Barang int	
	ID_Barang int
	Detail DetailBarang
}


type Kasir struct{
	Data ProdukCounterHandphone
	Lanjut *Kasir
}