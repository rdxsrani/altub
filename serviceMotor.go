package main

import (
	"fmt"
	"os"
)

const kapasitas int = 1000

type sparePart struct {
	nama        string
	id          string
	stok        int
	harga       float64
	tarif       float64
	jumlahGanti int
}

type pelanggan struct {
	tanggal      string
	bulan        string
	tahun        string
	nama         string
	nomor        string
	tarifService float64
	belanjaan    [kapasitas]buy
	nBuy         int
}

type buy struct {
	barang      string
	totalBarang int
}

type dataPenjualan struct {
	sp    [kapasitas]sparePart
	nSP   int
	cust  [kapasitas]pelanggan
	nCust int
}

func main() {
	var x dataPenjualan
	var pilihan int
	intro()
	for true {
		fmt.Println("======================\t Menu Utama Service Motor \t==================")
		fmt.Println("0. Keluar")
		fmt.Println("1. Menu Spare Part")
		fmt.Println("2. Menu Pelanggan")
		fmt.Println("===========================================================================")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)
		if pilihan == 0 || pilihan == 1 || pilihan == 2 {
			switch pilihan {
			case 0:
				farewell()
				os.Exit(1)
			case 1:
				menuSparePart(&x)
			case 2:
				menuPelanggan(&x)
			}
		}
	}
}

func menuSparePart(data *dataPenjualan) {
	var pilihan int
	for true {
		fmt.Println()
		fmt.Println("-----\tMenu Data Spare Part\t-----")
		fmt.Println("0. Kembali ")
		fmt.Println("1. Input Barang")
		fmt.Println("2. Tampilkan Barang")
		fmt.Println("3. Ubah Data Barang")
		fmt.Println("4. Cari Barang")
		fmt.Println("5. Cek Tarif Service")
		fmt.Println("-------------------------------------")
		fmt.Print("Masukkan pilihan menu: ")
		fmt.Scan(&pilihan)
		fmt.Println()
		if pilihan == 0 || pilihan == 1 || pilihan == 2 || pilihan == 3 || pilihan == 4 || pilihan == 5 {
			switch pilihan {
			case 0:
				return
			case 1:
				entrySparePart(data)
			case 2:
				fmt.Println("\n---------\tData Diurutkan Berdasarkan\t---------")
				fmt.Println("1. Nama Barang")
				fmt.Println("2. Harga Barang")
				fmt.Println("3. ID Barang")
				fmt.Println("4. Stok Barang")
				fmt.Println("5. Tarif Service Barang")
				fmt.Println("6. Penjualan Barang")
				var choose int
				fmt.Print("Masukkan pilihan: ")
				fmt.Scan(&choose)
				if choose == 1 || choose == 2 || choose == 3 || choose == 4 || choose == 5 || choose == 6 {
					fmt.Println("\nUrutkan secara:")
					fmt.Println("1. Ascending")
					fmt.Println("2. Descending")
					fmt.Print("Masukkan Pilihan: ")
					var pil int
					fmt.Scan(&pil)
					if pil == 1 || pil == 2 {
						sortingSP(data, pil, choose)
					}
				}
				tampilkanSP(*data)
			case 3:
				var pil int
				var sp string
				fmt.Println()
				fmt.Println("---------\tMenu Ubah Data\t---------")
				fmt.Println("1. Ubah Harga Barang")
				fmt.Println("2. Tambah Stok Barang")
				fmt.Println("3. Hapus Barang")
				fmt.Print("Masukkan pilihan: ")
				fmt.Scan(&pil)
				fmt.Print("Masukkan nama barang: ")
				fmt.Scan(&sp)
				var idx = searchIdxSP(*data, sp)
				if idx != -1 {
					switch pil {
					case 1:
						updateHargaSP(data, sp)
					case 2:
						addStokSP(data, sp)
					case 3:
						removeSparePart(data, sp)
					}
				} else {
					fmt.Println("Data tidak ditemukan.")
				}
			case 4:
				fmt.Println("\n----\tCari berdasarkan\t----")
				fmt.Println("1. Nama Barang")
				fmt.Println("2. ID Barang")
				fmt.Println("3. Stok Barang")
				fmt.Println("4. Harga Barang")
				fmt.Println("5. Tarif Service Barang")
				fmt.Println("6. Barang Dengan Total Jumlah Ganti")
				var pil int
				fmt.Print("Masukkan pilihan: ")
				fmt.Scan(&pil)
				var hasil dataPenjualan = searchSP(*data, pil)
				if hasil.nSP > 0 {
					tampilkanSP(hasil)
				} else if hasil.nSP <= 0 && pil != 6 {
					fmt.Println("===Data tidak ditemukan===")
				}
			case 5:
				var barang string
				fmt.Print("Masukkan Nama Barang: ")
				fmt.Scan(&barang)
				hitungTarif(*data, barang)
			}
		}
	}
}

func menuPelanggan(data *dataPenjualan) {
	var pilihan int
	for true {
		fmt.Println("\n-----\tMenu Data Pelanggan\t-----")
		fmt.Println("0. Kembali ")
		fmt.Println("1. Input Pelanggan")
		fmt.Println("2. Tampilkan Pelanggan")
		fmt.Println("3. Ubah Data Pelanggan")
		fmt.Println("4. Hapus Data Pelanggan")
		fmt.Println("5. Cari Pelanggan")
		fmt.Println("-------------------------------------")
		fmt.Print("Masukkan pilihan menu: ")
		fmt.Scan(&pilihan)
		fmt.Println()
		if pilihan == 0 || pilihan == 1 || pilihan == 2 || pilihan == 3 || pilihan == 4 || pilihan == 5 {
			switch pilihan {
			case 0:
				return
			case 1:
				entryPelanggan(data)
			case 2:
				fmt.Println("\n---------\tData Diurutkan Berdasarkan\t---------")
				fmt.Println("1. Nama Pelanggan")
				fmt.Println("2. Waktu Transaksi")
				fmt.Println("3. Jumlah Transaksi")
				var choose int
				fmt.Print("Masukkan Pilihan: ")
				fmt.Scan(&choose)
				if choose == 1 || choose == 2 || choose == 3 {
					sortingP(data, choose)
				}
				tampilkanP(*data)
			case 3:
				var p string
				fmt.Print("Masukkan nama pelanggan: ")
				fmt.Scan(&p)
				updateBelanjaP(data, p)
			case 4:
				var p string
				fmt.Print("Masukkan nama pelanggan: ")
				fmt.Scan(&p)
				removePelanggan(data, p)
			case 5:
				fmt.Println("\n----\tCari berdasarkan\t----")
				fmt.Println("1. Nama Pelanggan")
				fmt.Println("2. Waktu Transaksi")
				fmt.Println("3. Total Service")
				fmt.Println("4. Nomor Telepon")
				fmt.Println("5. Pembelian Spare Part")
				var pil int
				var nama string
				fmt.Print("Masukkan pilihan: ")
				fmt.Scan(&pil)
				if pil == 1 {
					fmt.Print("Masukkan Nama Pelanggan: ")
					fmt.Scan(&nama)
					idxByName := searchIdxName(*data, nama)
					if idxByName == -1 {
						fmt.Println("===Data tidak ditemukan===")
					} else {
						fmt.Printf("\nNama\t\t: %s\n", data.cust[idxByName].nama)
						fmt.Printf("Nomor Telepon\t: %v\n", data.cust[idxByName].nomor)
						fmt.Printf("Waktu Transaksi\t: Tanggal %v, Bulan %v, Tahun %v\n", data.cust[idxByName].tanggal, data.cust[idxByName].bulan, data.cust[idxByName].tahun)
						fmt.Println()
						fmt.Println("Daftar Transaksi")
						for j := 0; j < data.cust[idxByName].nBuy; j++ {
							fmt.Printf("%s sebanyak %v buah\n", data.cust[idxByName].belanjaan[j].barang, data.cust[idxByName].belanjaan[j].totalBarang)
						}
					}
				} else {
					var hasil dataPenjualan = searchP(*data, pil)
					if hasil.nCust > 0 {
						tampilkanP(hasil)
					} else {
						fmt.Println("===Data tidak ditemukan===")
					}
				}
			}
		}
	}
}

// SPARE PART
func entrySparePart(data *dataPenjualan) {
	var nama string
	fmt.Println("xxx untuk Selesai")
	fmt.Print("Masukkan Spare Part: ")
	fmt.Scan(&nama)
	for nama != "xxx" {
		data.sp[data.nSP].nama = nama
		fmt.Printf("Kode %s: ", data.sp[data.nSP].nama)
		fmt.Scan(&data.sp[data.nSP].id)
		fmt.Printf("Stok %s: ", data.sp[data.nSP].nama)
		fmt.Scan(&data.sp[data.nSP].stok)
		fmt.Printf("Harga %s peritem: ", data.sp[data.nSP].nama)
		fmt.Scan(&data.sp[data.nSP].harga)
		fmt.Printf("Tarif service %s: ", data.sp[data.nSP].nama)
		fmt.Scan(&data.sp[data.nSP].tarif)
		data.nSP++
		fmt.Print("===Data berhasil diinput===\n")
		fmt.Print("\nMasukkan Spare Part: ")
		fmt.Scan(&nama)
	}
}

func sortingSP(data *dataPenjualan, urutan, flag int) {
	var i, j int
	var idx int
	var t sparePart
	for i = 1; i <= data.nSP-1; i++ {
		idx = i - 1
		for j = i; j < data.nSP; j++ {
			switch flag {
			case 1:
				//Berdasarkan Nama
				//1 Ascending dan 2 Descending
				if urutan == 1 && data.sp[idx].nama > data.sp[j].nama {
					idx = j
				} else if urutan == 2 && data.sp[idx].nama < data.sp[j].nama {
					idx = j
				}
			case 2:
				//Berdasarkan Harga
				//1 Ascending dan 2 Descending
				if urutan == 1 && data.sp[idx].harga > data.sp[j].harga {
					idx = j
				} else if urutan == 2 && data.sp[idx].harga < data.sp[j].harga {
					idx = j
				}
			case 3:
				//Berdasarkan ID
				//1 Ascending dan 2 Descending
				if urutan == 1 && data.sp[idx].id > data.sp[j].id {
					idx = j
				} else if urutan == 2 && data.sp[idx].id < data.sp[j].id {
					idx = j
				}
			case 4:
				//Berdasarkan Stok
				//1 Ascending dan 2 Descending
				if urutan == 1 && data.sp[idx].stok > data.sp[j].stok {
					idx = j
				} else if urutan == 2 && data.sp[idx].stok < data.sp[j].stok {
					idx = j
				}
			case 5:
				//Berdasarkan Tarif
				//1 Ascending dan 2 Descending
				if urutan == 1 && data.sp[idx].tarif > data.sp[j].tarif {
					idx = j
				} else if urutan == 2 && data.sp[idx].tarif < data.sp[j].tarif {
					idx = j
				}
			case 6:
				//Berdasarkan banyaknya jumlah ganti
				//1 Ascending dan 2 Descending
				if urutan == 1 && data.sp[idx].jumlahGanti > data.sp[j].jumlahGanti {
					idx = j
				} else if urutan == 2 && data.sp[idx].jumlahGanti < data.sp[j].jumlahGanti {
					idx = j
				}
			}
		}
		t = data.sp[idx]
		data.sp[idx] = data.sp[i-1]
		data.sp[i-1] = t
	}
}

func tampilkanSP(data dataPenjualan) {
	fmt.Println()
	fmt.Println("======================================================")
	for i := 0; i < data.nSP; i++ {
		fmt.Println("Nama\t\t:", data.sp[i].nama)
		fmt.Println("Id\t\t:", data.sp[i].id)
		fmt.Println("Stok\t\t:", data.sp[i].stok)
		fmt.Println("Harga\t\t:", data.sp[i].harga)
		fmt.Println("Tarif Service\t:", data.sp[i].tarif)
		fmt.Println()
	}
	fmt.Println("======================================================")
}

func updateHargaSP(data *dataPenjualan, nama string) {
	var idx = searchIdxSP(*data, nama)
	var harga float64
	fmt.Printf("Masukkan harga %s terbaru: ", data.sp[idx].nama)
	fmt.Scan(&harga)
	data.sp[idx].harga = harga
	fmt.Printf("===Harga %s berhasil diubah===\n", data.sp[idx].nama)
}

func addStokSP(data *dataPenjualan, sp string) {
	var idx = searchIdxSP(*data, sp)
	fmt.Println("Nama\t\t:", data.sp[idx].nama)
	fmt.Println("Id\t\t:", data.sp[idx].id)
	fmt.Println("Stok\t\t:", data.sp[idx].stok)
	fmt.Println("Harga\t\t:", data.sp[idx].harga)
	fmt.Println("Tarif Service\t:", data.sp[idx].tarif)
	fmt.Print("Masukkan tambahan stok: ")
	var stok int
	fmt.Scan(&stok)
	data.sp[idx].stok += stok
	fmt.Print("===Stok berhasil ditambah===")
}

func removeSparePart(data *dataPenjualan, x string) {
	var found, i int
	found = searchIdxSP(*data, x)
	if found == -1 {
		fmt.Println("===Data tidak ditemukan===")
	} else {
		i = found
		for i < data.nSP-1 {
			data.sp[i] = data.sp[i+1]
			i++
		}
		data.nSP--
		fmt.Println("===Data berhasil dihapus===")
	}
}

func searchSP(data dataPenjualan, flag int) dataPenjualan {
	var i, stok int
	var harga, tarif float64
	var id, nama string
	var arrSP dataPenjualan
	switch flag {
	case 1:
		fmt.Print("Masukkan nama Spare Part: ")
		fmt.Scan(&nama)
		for i = 0; i < data.nSP; i++ {
			if nama == data.sp[i].nama {
				arrSP.sp[arrSP.nSP] = data.sp[i]
				arrSP.nSP++
			}
		}
	case 2:
		fmt.Print("Masukkan ID Spare Part: ")
		fmt.Scan(&id)
		for i = 0; i < data.nSP; i++ {
			if id == data.sp[i].id {
				arrSP.sp[arrSP.nSP] = data.sp[i]
				arrSP.nSP++
			}
		}
	case 3:
		fmt.Print("Masukkan Stok Spare Part: ")
		fmt.Scan(&stok)
		for i = 0; i < data.nSP; i++ {
			if stok == data.sp[i].stok {
				arrSP.sp[arrSP.nSP] = data.sp[i]
				arrSP.nSP++
			}
		}
	case 4:
		fmt.Print("Masukkan Harga Spare Part: ")
		fmt.Scan(&harga)
		for i = 0; i < data.nSP; i++ {
			if harga == data.sp[i].harga {
				arrSP.sp[arrSP.nSP] = data.sp[i]
				arrSP.nSP++
			}
		}
	case 5:
		fmt.Print("Masukkan Tarif Spare Part: ")
		fmt.Scan(&tarif)
		for i = 0; i < data.nSP; i++ {
			if tarif == data.sp[i].tarif {
				arrSP.sp[arrSP.nSP] = data.sp[i]
				arrSP.nSP++
			}
		}
	case 6:
		var pil int
		fmt.Println("Tampilkan: ")
		fmt.Println("1. Paling Banyak")
		fmt.Println("2. Paling Sedikit")
		fmt.Printf("Masukkan pilihan: ")
		fmt.Scan(&pil)
		var idx = MaxMinJumlahGanti(data, pil)
		fmt.Println("Nama\t\t:", data.sp[idx].nama)
		fmt.Println("Id\t\t:", data.sp[idx].id)
		fmt.Println("Stok\t\t:", data.sp[idx].stok)
		fmt.Println("Jumlah Ganti\t:", data.sp[idx].jumlahGanti)
		fmt.Println("Harga\t\t:", data.sp[idx].harga)
		fmt.Println("Tarif Service\t:", data.sp[idx].tarif)
		fmt.Println()
	}
	return arrSP
}

func hitungTarif(data dataPenjualan, nama string) {
	var jum int
	var tarif float64
	var toTarif float64
	idx := searchIdxSP(data, nama)
	for nama != "xxx" && idx != -1 {
		fmt.Print("Jumlah barang: ")
		fmt.Scan(&jum)
		tarif = (data.sp[idx].harga * float64(jum))
		toTarif += tarif + data.sp[idx].tarif
		fmt.Print("Masukkan Nama Barang: ")
		fmt.Scan(&nama)
	}
	if nama == "xxx" || idx == -1 {
		if idx == -1 {
			fmt.Println("===Data tidak ditemukan===")
		}
		fmt.Println("Jumlah tarif service: ", toTarif)
	}
}

func searchIdxSP(data dataPenjualan, nama string) int {
	var i int
	for i = 0; i < data.nSP; i++ {
		if nama == data.sp[i].nama {
			return i
		}
	}
	return -1
}

func MaxMinJumlahGanti(data dataPenjualan, flag int) int {
	var idx int = 0
	switch flag {
	case 1:
		//Min
		for i := 1; i < data.nSP; i++ {
			if data.sp[idx].jumlahGanti < data.sp[i].jumlahGanti {
				idx = i
			}
		}
	case 2:
		//Max
		for i := 1; i < data.nSP; i++ {
			if data.sp[idx].jumlahGanti > data.sp[i].jumlahGanti {
				idx = i
			}
		}
	}
	return idx
}

// PELANGGAN
func entryPelanggan(data *dataPenjualan) {
	var tarif float64
	var idxBuy int
	var pelanggan string
	data.cust[data.nCust].tarifService = 0
	fmt.Println("xxx untuk Selesai")
	fmt.Print("Masukkan nama pelanggan: ")
	fmt.Scan(&pelanggan)
	for pelanggan != "xxx" {
		data.cust[data.nCust].nama = pelanggan
		tarif = 0
		idxBuy = 0
		fmt.Print("Nomor telepon: ")
		fmt.Scan(&data.cust[data.nCust].nomor)
		fmt.Print("Waktu service :\n")
		fmt.Print("Tanggal Bulan Tahun\n")
		fmt.Scan(&data.cust[data.nCust].tanggal, &data.cust[data.nCust].bulan, &data.cust[data.nCust].tahun)
		fmt.Print("Item service: ")
		var belanja buy
		fmt.Scan(&belanja.barang)
		for belanja.barang != "xxx" {
			//Belanjaan Pelanggan
			idxSP := searchIdxSP(*data, belanja.barang)
			if idxSP == -1 {
				fmt.Print("===Data tidak ditemukan===\n")
			} else {
				fmt.Print("Banyak item yang digunakan: ")
				fmt.Scan(&belanja.totalBarang)
				if belanja.totalBarang > data.sp[idxSP].stok {
					if data.sp[idxSP].stok == 0 {
						fmt.Println("\nStok Habis")
						fmt.Println("Tambahkan Stok Barang!")
						fmt.Printf("===Pembelian %s gagal===\n", data.sp[idxSP].nama)
					} else {
						fmt.Println("Stok Kurang")
						fmt.Printf("Jumlah yang dapat dibeli adalah %v\n", data.sp[idxSP].stok)
						fmt.Printf("===Pembelian %s gagal===\n", data.sp[idxSP].nama)
					}
				} else {
					data.sp[idxSP].stok -= belanja.totalBarang
					data.sp[idxSP].jumlahGanti += belanja.totalBarang
					tarif = (data.sp[idxSP].harga * float64(belanja.totalBarang))
					data.cust[data.nCust].belanjaan[idxBuy] = belanja
					data.cust[data.nCust].tarifService += tarif + data.sp[idxSP].tarif
					idxBuy++
				}
				data.cust[data.nCust].nBuy = idxBuy
			}
			fmt.Print("Item service: ")
			fmt.Scan(&belanja.barang)
		}
		fmt.Printf("Tarif service pelanggan dengan nama: %s, yaitu: Rp%v\n", data.cust[data.nCust].nama, tarif)
		data.nCust++
		fmt.Println("===Data berhasil diinput===")
		fmt.Print("\nMasukkan nama pelanggan: ")
		fmt.Scan(&pelanggan)
	}
}

func sortingP(data *dataPenjualan, flag int) {
	var pass, i, urutan int
	var temp pelanggan
	for pass = 1; pass < data.nCust; pass++ {
		i = pass
		temp = data.cust[i]
		switch flag {
		case 1:
			// Berdasarkan Nama
			// 1 Ascending dan 2 Descending
			fmt.Println("\nUrutkan secara:")
			fmt.Println("1. Ascending")
			fmt.Println("2. Descending")
			fmt.Print("Masukkan Pilihan: ")
			fmt.Scan(&urutan)
			if urutan == 1 {
				for i > 0 && temp.nama < data.cust[i-1].nama {
					data.cust[i] = data.cust[i-1]
					i--
				}
			} else if urutan == 2 {
				for i > 0 && temp.nama > data.cust[i-1].nama {
					data.cust[i] = data.cust[i-1]
					i--
				}
			}
		case 2:
			//Mengurutkan Waktu Transaksi
			// 1 Terlama dan 2 Terbaru
			fmt.Println("Urutkan secara:")
			fmt.Println("1. Terlama")
			fmt.Println("2. Terbaru")
			fmt.Print("Masukkan Pilihan: ")
			fmt.Scan(&urutan)
			if urutan == 1 {
				for i > 0 && temp.tanggal <= data.cust[i-1].tanggal && temp.bulan <= data.cust[i-1].bulan && temp.tahun <= data.cust[i-1].tahun {
					data.cust[i] = data.cust[i-1]
					i--
				}
			} else if urutan == 2 {
				for i > 0 && temp.tanggal >= data.cust[i-1].tanggal && temp.bulan >= data.cust[i-1].bulan && temp.tahun >= data.cust[i-1].tahun {
					data.cust[i] = data.cust[i-1]
					i--
				}
			}
		case 3:
			// Berdasarkan Belanjaan
			// 1 Ascending dan 2 Descending
			fmt.Println("Urutkan secara:")
			fmt.Println("1. Ascending")
			fmt.Println("2. Descending")
			fmt.Print("Masukkan Pilihan: ")
			fmt.Scan(&urutan)
			if urutan == 1 {
				for i > 0 && temp.tarifService < data.cust[i-1].tarifService {
					data.cust[i] = data.cust[i-1]
					i--
				}
			} else if urutan == 2 {
				for i > 0 && temp.tarifService > data.cust[i-1].tarifService {
					data.cust[i] = data.cust[i-1]
					i--
				}
			}
		}
		data.cust[i] = temp
	}
}

func tampilkanP(data dataPenjualan) {
	fmt.Println("======================================================")
	for i := 0; i < data.nCust; i++ {
		fmt.Printf("\nNama             : %s\n", data.cust[i].nama)
		fmt.Printf("Nomor Telepon    : %v\n", data.cust[i].nomor)
		fmt.Printf("Waktu Transaksi  : Tanggal %v, Bulan %v, Tahun %v\n", data.cust[i].tanggal, data.cust[i].bulan, data.cust[i].tahun)
		tampilkanBelanjaP(data, data.cust[i].nama)
		fmt.Println()
	}
	fmt.Println("======================================================")
}

func updateBelanjaP(data *dataPenjualan, nama string) {
	var idxP = searchIdxP(*data, nama)
	var sp string
	if idxP != -1 {
		fmt.Printf("\nNama : %s\nNomor Telepon: %v\nWaktu Transaksi: \nTanggal %v, Bulan %v, Tahun %v\n", data.cust[idxP].nama, data.cust[idxP].nomor, data.cust[idxP].tanggal, data.cust[idxP].bulan, data.cust[idxP].tahun)
		fmt.Println()
		fmt.Println("Daftar Transaksi")
		for j := 0; j < data.cust[idxP].nBuy; j++ {
			fmt.Printf("%s sebanyak %v buah\n", data.cust[idxP].belanjaan[j].barang, data.cust[idxP].belanjaan[j].totalBarang)
		}
		var pil int
		fmt.Printf("\n1. Hapus data belanja %s", nama)
		fmt.Printf("\n2. Tambah data belanja %s", nama)
		fmt.Print("\nMasukkan pilihan: ")
		fmt.Scan(&pil)
		switch pil {
		case 1:
			fmt.Print("Masukkan nama belanjaan yang ingin dihapus: ")
			fmt.Scan(&sp)
			removeTransaksiP(data, data.cust[idxP].nama, sp)
		case 2:
			fmt.Print("Masukkan nama belanjaan yang ingin ditambah: ")
			fmt.Scan(&sp)
			addBelanjaP(data, nama, sp)
		}
	} else {
		fmt.Println("===Data tidak ditemukan===")
	}
}

func addBelanjaP(data *dataPenjualan, nama, barang string) {
	var nBuy, idxP, idxBuyP, idxSP int
	var tarif float64
	idxBuyP = searchIdxBuyP(*data, nama, barang)
	idxP = searchIdxP(*data, nama)
	idxSP = searchIdxSP(*data, barang)
	nBuy = data.cust[idxP].nBuy
	if idxSP == -1 {
		fmt.Println("===Data tidak ditemukan===")
	} else {
		var jumBarang int
		fmt.Print("Banyak item yang digunakan: ")
		fmt.Scan(&jumBarang)
		if jumBarang > data.sp[idxSP].stok {
			if data.sp[idxSP].stok == 0 {
				fmt.Println("\nStok Habis")
				fmt.Printf("===Penambahan %s gagal===\n", data.sp[idxSP].nama)
			} else {
				fmt.Println("Stok Kurang")
				fmt.Printf("Jumlah yang dapat dibeli adalah %v\n", data.sp[idxSP].stok)
				fmt.Printf("===Penambahan %s gagal===\n", data.sp[idxSP].nama)
			}
		} else {
			data.sp[idxSP].stok -= jumBarang
			data.sp[idxSP].jumlahGanti += jumBarang
			tarif += (data.sp[idxSP].harga * float64(jumBarang))
			if idxBuyP != -1 {
				tarif += (data.sp[idxSP].harga * float64(data.sp[idxSP].stok))
				data.cust[idxP].belanjaan[idxBuyP].totalBarang += jumBarang
				data.cust[idxP].tarifService += tarif + data.sp[idxSP].tarif
				fmt.Println("===Data berhasil ditambah===")
				data.cust[idxP].belanjaan[idxBuyP].barang = barang
			} else {
				data.cust[idxP].belanjaan[nBuy].barang = barang
				data.cust[idxP].belanjaan[nBuy].totalBarang = jumBarang
				data.cust[idxP].tarifService += tarif + data.sp[idxSP].tarif
				fmt.Println("===Data berhasil ditambah===")
				nBuy++
			}
		}
		data.cust[idxP].nBuy = nBuy
	}
}

func removeTransaksiP(data *dataPenjualan, nama, barang string) {
	var idxBuyP, idxP, idxSP, i int
	idxBuyP = searchIdxBuyP(*data, nama, barang)
	idxP = searchIdxP(*data, nama)
	idxSP = searchIdxSP(*data, barang)
	if idxBuyP == -1 {
		fmt.Println("===Data tidak ditemukan===")
	} else {
		data.sp[idxSP].stok += data.cust[idxP].belanjaan[idxBuyP].totalBarang
		i = idxBuyP
		for i < data.cust[idxP].nBuy {
			data.cust[idxP].belanjaan[i] = data.cust[idxP].belanjaan[i+1]
			i++
		}
		data.cust[idxP].nBuy--
		fmt.Println("===Data berhasil dihapus===")
	}
}

func removePelanggan(data *dataPenjualan, x string) {
	var found, i int
	found = searchIdxP(*data, x)
	if found == -1 {
		fmt.Println("===Data tidak ditemukan===")
	} else {
		i = found
		for i < data.nCust-1 {
			data.cust[i] = data.cust[i+1]
			i++
		}
		data.nCust--
		fmt.Println("===Data berhasil dihapus===")
	}
}

func searchP(data dataPenjualan, flag int) dataPenjualan {
	var i int
	var tarif float64
	var belanjaan, nomor, tgl, bln, thn string
	var arrP dataPenjualan
	arrP.nCust = 0
	switch flag {
	case 2:
		fmt.Print("Masukkan waktu Transaksi Pelanggan: ")
		fmt.Scan(&tgl, &bln, &thn)
		for i = 0; i < data.nCust; i++ {
			if tgl == data.cust[i].tanggal && bln == data.cust[i].bulan && thn == data.cust[i].tahun {
				arrP.cust[arrP.nCust] = data.cust[i]
				arrP.nCust++
			}
		}
	case 3:
		fmt.Print("Masukkan Jumlah Pembelian Pelanggan: ")
		fmt.Scan(&tarif)
		for i = 0; i < data.nCust; i++ {
			if tarif == data.cust[i].tarifService {
				arrP.cust[arrP.nCust] = data.cust[i]
				arrP.nCust++
			}
		}
	case 4:
		fmt.Print("Masukkan Nomor Telepon Pelanggan: ")
		fmt.Scan(&nomor)
		for i = 0; i < data.nCust; i++ {
			if nomor == data.cust[i].nomor {
				arrP.cust[arrP.nCust] = data.cust[i]
				arrP.nCust++
			}
		}
	case 5:
		fmt.Print("Masukkan Spare Part: ")
		fmt.Scan(&belanjaan)
		for i = 0; i < data.nCust; i++ {
			for j := 0; j < data.cust[i].nBuy; j++ {
				if belanjaan == data.cust[i].belanjaan[j].barang {
					arrP.cust[arrP.nCust] = data.cust[i]
					arrP.nCust++
				}
			}
		}
	}
	return arrP
}

func searchIdxP(data dataPenjualan, nama string) int {
	for i := 0; i < data.nCust; i++ {
		if nama == data.cust[i].nama {
			return i
		}
	}
	return -1
}

func searchIdxName(data dataPenjualan, nama string) int {
	var mid, left, right int
	right = data.nCust - 1
	for left <= right {
		mid = (left + right) / 2
		if nama == data.cust[mid].nama {
			return mid
		} else if nama > data.cust[mid].nama {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func searchIdxBuyP(data dataPenjualan, nama, belanja string) int {
	var i int
	var idxP = searchIdxP(data, nama)
	for i = 0; i < data.cust[idxP].nBuy; i++ {
		if belanja == data.cust[idxP].belanjaan[i].barang {
			return i
		}
	}
	return -1
}

func tampilkanBelanjaP(data dataPenjualan, nama string) {
	var idxP = searchIdxP(data, nama)
	fmt.Println("======================================================")
	fmt.Println("Daftar Transaksi")
	for i := 0; i < data.cust[idxP].nBuy; i++ {
		fmt.Printf("%s sebanyak %v buah\n", data.cust[idxP].belanjaan[i].barang, data.cust[idxP].belanjaan[i].totalBarang)
	}
	fmt.Println("======================================================")
}

func intro() {
	fmt.Println("Welcome to our app")
}

func farewell() {
	fmt.Println("Bye. See you :)")
}
