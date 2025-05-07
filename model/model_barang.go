package model

import (
	"uts_data/node"
	"strings"

)

var DaftarBarang node.Kasir

func CreateProduk(emp node.Wartawarti)bool{
	Nodebaru := &node.Kasir{
		Data: emp,
		Lanjut: nil,
	}
	if DaftarBarang.Lanjut == nil{
		DaftarBarang.Lanjut = Nodebaru
		return true
	}else{
		temp := &DaftarBarang
			for temp.Lanjut != nil{
				temp = temp.Lanjut
			}	
			temp.Lanjut = Nodebaru
			return true
		}
	}

	func ReadProduk() []node.Wartawarti{
		daftarBarang := []node.Wartawarti{}
			Nodebaru := DaftarBarang.Lanjut
			for Nodebaru != nil{
				daftarBarang = append(daftarBarang, Nodebaru.Data)
				Nodebaru = Nodebaru.Lanjut
			}
			return daftarBarang
		}
	
		func UpdateProduk(emp node.Wartawarti,id int)bool{
			temp :=DaftarBarang.Lanjut
			for temp != nil{
				if temp.Data.ID_Produk== id{
					temp.Data = emp
					return true
				}
				temp = temp.Lanjut
			}
			return false
		}

		func DeleteProduk( id int)bool{
			if DaftarBarang.Lanjut == nil{
				return false
			}

				temp := &DaftarBarang
				for temp.Lanjut != nil{
					if temp.Lanjut.Data.ID_Produk == id{
						temp.Lanjut = temp.Lanjut.Lanjut
						return true
					}
					temp = temp.Lanjut
				}
				return false
			}


			func SearchProduk(kategori string) []node.Wartawarti {
				Produk := []node.Wartawarti{}
				temp := DaftarBarang.Lanjut
				for temp != nil {
					if strings.EqualFold(temp.Data.Kategori, kategori) {
						Produk = append(Produk, temp.Data)
					}
					temp = temp.Lanjut
				}
				return Produk
			}
			


