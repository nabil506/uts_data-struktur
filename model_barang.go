package model

import (
	"uts_data/node"

)

var DaftarBarang node.Kasir

func CreateBarang(emp node.ProdukCounterHandphone)bool{
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

	func ReadBarang() []node.ProdukCounterHandphone{
		daftarBarang := []node.ProdukCounterHandphone{}
			Nodebaru := DaftarBarang.Lanjut
			for Nodebaru != nil{
				daftarBarang = append(daftarBarang, Nodebaru.Data)
				Nodebaru = Nodebaru.Lanjut
			}
			return daftarBarang
		}
	
		func UpdateBarang(emp node.ProdukCounterHandphone,id int)bool{
			temp :=DaftarBarang.Lanjut
			for temp != nil{
				if temp.Data.ID_Barang== id{
					temp.Data = emp
					return true
				}
				temp = temp.Lanjut
			}
			return false
		}

		func DeleteBarang( id int)bool{
			if DaftarBarang.Lanjut == nil{
				return false
			}

				temp := &DaftarBarang
				for temp.Lanjut != nil{
					if temp.Lanjut.Data.ID_Barang== id{
						temp.Lanjut = temp.Lanjut.Lanjut
						return true
					}
					temp = temp.Lanjut
				}
				return false
			}


			func SearchBarang(Nama string) *node.ProdukCounterHandphone {
				temp := DaftarBarang.Lanjut
				for temp != nil {
					if temp.Data.Nama_barang == Nama {
						return &temp.Data
					}
					temp = temp.Lanjut
				}
				return nil
			}
		


