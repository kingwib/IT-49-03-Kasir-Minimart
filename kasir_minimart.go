package main

import "fmt"

type Produk struct {
	kode  string
	nama  string
	harga float64
	stok  int
}

type Member struct {
	id   string
	nama string
	noHP string
	poin int
}

type ItemBeli struct {
	nama     string
	harga    float64
	qty      int
	subtotal float64
}

type Transaksi struct {
	id         string
	idMember   string
	items      [10]ItemBeli
	jumlahItem int
	total      float64
	diskon     float64
	bayar      float64
}

var daftarProduk [10]Produk
var jumlahProduk int
var daftarMember [5]Member
var jumlahMember int
var daftarTransaksi [5]Transaksi
var jumlahTransaksi int

func garis() { fmt.Println("================================================") }

func inputStr(prompt string) string {
	fmt.Print(prompt)
	var s string
	fmt.Scan(&s)
	return s
}

func inputInt(prompt string) int {
	fmt.Print(prompt)
	var n int
	fmt.Scan(&n)
	return n
}

func inputFloat(prompt string) float64 {
	fmt.Print(prompt)
	var f float64
	fmt.Scan(&f)
	return f
}

func seqProduk(kode string) int {
	var i, found int
	found = -1
	i = 0
	for i < jumlahProduk && found == -1 {
		if daftarProduk[i].kode == kode {
			found = i
		}
		i = i + 1
	}
	return found
}

func seqMember(id string) int {
	var i, found int
	found = -1
	i = 0
	for i < jumlahMember && found == -1 {
		if daftarMember[i].id == id {
			found = i
		}
		i = i + 1
	}
	return found
}

func sortProdukKode() {
	var i, j, min int
	i = 0
	for i < jumlahProduk-1 {
		min = i
		j = i + 1
		for j < jumlahProduk {
			if daftarProduk[j].kode < daftarProduk[min].kode {
				min = j
			}
			j = j + 1
		}
		var tmp Produk
		tmp = daftarProduk[i]
		daftarProduk[i] = daftarProduk[min]
		daftarProduk[min] = tmp
		i = i + 1
	}
}

func binProduk(kode string) int {
	sortProdukKode()
	var lo, hi, mid, found int
	lo = 0
	hi = jumlahProduk - 1
	found = -1
	for lo <= hi && found == -1 {
		mid = (lo + hi) / 2
		if daftarProduk[mid].kode == kode {
			found = mid
		} else if daftarProduk[mid].kode < kode {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return found
}

func selSortProdukHargaAsc() {
	var i, j, min int
	i = 0
	for i < jumlahProduk-1 {
		min = i
		j = i + 1
		for j < jumlahProduk {
			if daftarProduk[j].harga < daftarProduk[min].harga {
				min = j
			}
			j = j + 1
		}
		var tmp Produk
		tmp = daftarProduk[i]
		daftarProduk[i] = daftarProduk[min]
		daftarProduk[min] = tmp
		i = i + 1
	}
}

func selSortProdukHargaDesc() {
	var i, j, max int
	i = 0
	for i < jumlahProduk-1 {
		max = i
		j = i + 1
		for j < jumlahProduk {
			if daftarProduk[j].harga > daftarProduk[max].harga {
				max = j
			}
			j = j + 1
		}
		var tmp Produk
		tmp = daftarProduk[i]
		daftarProduk[i] = daftarProduk[max]
		daftarProduk[max] = tmp
		i = i + 1
	}
}

func insSortMemberPoin() {
	var i, j int
	var key Member
	i = 1
	for i < jumlahMember {
		key = daftarMember[i]
		j = i - 1
		for j >= 0 && daftarMember[j].poin < key.poin {
			daftarMember[j+1] = daftarMember[j]
			j = j - 1
		}
		daftarMember[j+1] = key
		i = i + 1
	}
}

func insSortMemberPoinAsc() {
	var i, j int
	var key Member
	i = 1
	for i < jumlahMember {
		key = daftarMember[i]
		j = i - 1
		for j >= 0 && daftarMember[j].poin > key.poin {
			daftarMember[j+1] = daftarMember[j]
			j = j - 1
		}
		daftarMember[j+1] = key
		i = i + 1
	}
}

func tampilProduk() {
	garis()
	fmt.Println("           DAFTAR PRODUK")
	garis()
	fmt.Printf("%-6s %-20s %10s %6s\n", "KODE", "NAMA", "HARGA", "STOK")
	fmt.Println("------------------------------------------------")
	var i int
	i = 0
	for i < jumlahProduk {
		fmt.Printf("%-6s %-20s %10.0f %6d\n", daftarProduk[i].kode, daftarProduk[i].nama, daftarProduk[i].harga, daftarProduk[i].stok)
		i = i + 1
	}
	garis()
}

func tambahProduk() {
	if jumlahProduk >= 10 {
		fmt.Println("Data produk penuh!")
		return
	}
	var p Produk
	p.kode = inputStr("Kode   : ")
	if seqProduk(p.kode) != -1 {
		fmt.Println("Kode sudah ada!")
		return
	}
	p.nama = inputStr("Nama   : ")
	p.harga = inputFloat("Harga  : ")
	p.stok = inputInt("Stok   : ")
	daftarProduk[jumlahProduk] = p
	jumlahProduk = jumlahProduk + 1
	fmt.Println("Produk ditambahkan!")
}

func editProduk() {
	kode := inputStr("Kode produk: ")
	idx := seqProduk(kode)
	if idx == -1 {
		fmt.Println("Produk tidak ditemukan.")
		return
	}
	fmt.Printf("Edit: %s\n", daftarProduk[idx].nama)
	daftarProduk[idx].nama = inputStr("Nama baru  : ")
	daftarProduk[idx].harga = inputFloat("Harga baru : ")
	daftarProduk[idx].stok = inputInt("Stok baru  : ")
	fmt.Println("Produk diperbarui!")
}

func hapusProduk() {
	kode := inputStr("Kode produk: ")
	idx := seqProduk(kode)
	if idx == -1 {
		fmt.Println("Produk tidak ditemukan.")
		return
	}
	fmt.Printf("Hapus: %s? (y/n): ", daftarProduk[idx].nama)
	var k string
	fmt.Scan(&k)
	if k == "y" || k == "Y" {
		var i int
		i = idx
		for i < jumlahProduk-1 {
			daftarProduk[i] = daftarProduk[i+1]
			i = i + 1
		}
		jumlahProduk = jumlahProduk - 1
		fmt.Println("Produk dihapus!")
	}
}

func menuProduk() {
	var p int
	p = -1
	for p != 0 {
		garis()
		fmt.Println("1.Tampil 2.Tambah 3.Edit 4.Hapus 5.Cari 6.Urut 0.Kembali")
		p = inputInt("Pilih: ")
		if p == 1 {
			tampilProduk()
		} else if p == 2 {
			tambahProduk()
		} else if p == 3 {
			editProduk()
		} else if p == 4 {
			hapusProduk()
		} else if p == 5 {
			kode := inputStr("Kode (Binary Search): ")
			idx := binProduk(kode)
			if idx == -1 {
				fmt.Println("Tidak ditemukan.")
			} else {
				fmt.Printf("Ditemukan: %s | Rp %.0f | Stok: %d\n", daftarProduk[idx].nama, daftarProduk[idx].harga, daftarProduk[idx].stok)
			}
		} else if p == 6 {
			fmt.Println("1.Harga Asc  2.Harga Desc")
			u := inputInt("Pilih: ")
			if u == 1 {
				selSortProdukHargaAsc()
			} else {
				selSortProdukHargaDesc()
			}
			tampilProduk()
		}
	}
}

func tampilMember() {
	garis()
	fmt.Println("            DAFTAR MEMBER")
	garis()
	fmt.Printf("%-8s %-18s %-14s %6s\n", "ID", "NAMA", "NO HP", "POIN")
	fmt.Println("------------------------------------------------")
	var i int
	i = 0
	for i < jumlahMember {
		fmt.Printf("%-8s %-18s %-14s %6d\n", daftarMember[i].id, daftarMember[i].nama, daftarMember[i].noHP, daftarMember[i].poin)
		i = i + 1
	}
	garis()
}

func menuMember() {
	var p int
	p = -1
	for p != 0 {
		garis()
		fmt.Println("1.Tampil 2.Tambah 3.Hapus 4.Cari 5.Urut Poin 6.Edit Poin 0.Kembali")
		p = inputInt("Pilih: ")
		if p == 1 {
			tampilMember()
		} else if p == 2 {
			if jumlahMember >= 5 {
				fmt.Println("Data member penuh!")
			} else {
				var m Member
				m.id = inputStr("ID Member : ")
				if seqMember(m.id) != -1 {
					fmt.Println("ID sudah ada!")
				} else {
					m.nama = inputStr("Nama      : ")
					m.noHP = inputStr("No HP     : ")
					m.poin = 0
					daftarMember[jumlahMember] = m
					jumlahMember = jumlahMember + 1
					fmt.Println("Member ditambahkan!")
				}
			}
		} else if p == 3 {
			id := inputStr("ID Member: ")
			idx := seqMember(id)
			if idx == -1 {
				fmt.Println("Tidak ditemukan.")
			} else {
				fmt.Printf("Hapus: %s? (y/n): ", daftarMember[idx].nama)
				var k string
				fmt.Scan(&k)
				if k == "y" || k == "Y" {
					var i int
					i = idx
					for i < jumlahMember-1 {
						daftarMember[i] = daftarMember[i+1]
						i = i + 1
					}
					jumlahMember = jumlahMember - 1
					fmt.Println("Member dihapus!")
				}
			}
		} else if p == 4 {
			id := inputStr("ID Member (Sequential Search): ")
			idx := seqMember(id)
			if idx == -1 {
				fmt.Println("Tidak ditemukan.")
			} else {
				fmt.Printf("Nama: %s | Poin: %d\n", daftarMember[idx].nama, daftarMember[idx].poin)
			}
		} else if p == 5 {
			fmt.Println("1.Poin Terbanyak  2.Poin Tersedikit")
			u := inputInt("Pilih: ")
			if u == 1 {
				insSortMemberPoin()
			} else {
				insSortMemberPoinAsc()
			}
			tampilMember()
		} else if p == 6 {
			id := inputStr("ID Member: ")
			idx := seqMember(id)
			if idx == -1 {
				fmt.Println("Member tidak ditemukan.")
			} else {
				fmt.Printf("Member: %s | Poin saat ini: %d\n", daftarMember[idx].nama, daftarMember[idx].poin)
				fmt.Println("1.Tambah Poin  2.Kurangi Poin  3.Set Poin")
				opsi := inputInt("Pilih: ")
				jumlah := inputInt("Jumlah poin: ")
				if opsi == 1 {
					daftarMember[idx].poin = daftarMember[idx].poin + jumlah
					fmt.Printf("Poin bertambah. Total poin: %d\n", daftarMember[idx].poin)
				} else if opsi == 2 {
					daftarMember[idx].poin = daftarMember[idx].poin - jumlah
					if daftarMember[idx].poin < 0 {
						daftarMember[idx].poin = 0
					}
					fmt.Printf("Poin berkurang. Total poin: %d\n", daftarMember[idx].poin)
				} else if opsi == 3 {
					daftarMember[idx].poin = jumlah
					fmt.Printf("Poin diset ke: %d\n", daftarMember[idx].poin)
				} else {
					fmt.Println("Pilihan tidak valid.")
				}
			}
		}
	}
}

func buatTransaksi() {
	if jumlahTransaksi >= 5 {
		fmt.Println("Data transaksi penuh!")
		return
	}
	var t Transaksi
	t.id = fmt.Sprintf("TRX%04d", jumlahTransaksi+1)
	t.jumlahItem = 0
	t.total = 0

	var adaMember bool
	adaMember = false
	pakaiMember := inputStr("Pakai member? (y/n): ")
	if pakaiMember == "y" || pakaiMember == "Y" {
		id := inputStr("ID Member: ")
		idx := seqMember(id)
		if idx == -1 {
			fmt.Println("Member tidak ditemukan, lanjut non-member.")
		} else {
			t.idMember = daftarMember[idx].id
			adaMember = true
			fmt.Printf("Member: %s\n", daftarMember[idx].nama)
		}
	}

	lanjut := "y"
	for (lanjut == "y" || lanjut == "Y") && t.jumlahItem < 10 {
		kode := inputStr("Kode produk: ")
		idx := binProduk(kode)
		if idx == -1 {
			fmt.Println("Produk tidak ditemukan.")
		} else if daftarProduk[idx].stok == 0 {
			fmt.Println("Stok habis!")
		} else {
			fmt.Printf("%s | Rp %.0f | Stok: %d\n", daftarProduk[idx].nama, daftarProduk[idx].harga, daftarProduk[idx].stok)
			qty := inputInt("Qty: ")
			if qty > 0 && qty <= daftarProduk[idx].stok {
				var item ItemBeli
				item.nama = daftarProduk[idx].nama
				item.harga = daftarProduk[idx].harga
				item.qty = qty
				item.subtotal = daftarProduk[idx].harga * float64(qty)
				t.items[t.jumlahItem] = item
				t.jumlahItem = t.jumlahItem + 1
				t.total = t.total + item.subtotal
				daftarProduk[idx].stok = daftarProduk[idx].stok - qty
				fmt.Printf("Subtotal: Rp %.0f | Total: Rp %.0f\n", item.subtotal, t.total)
			} else {
				fmt.Println("Qty tidak valid!")
			}
		}
		if t.jumlahItem < 10 {
			lanjut = inputStr("Tambah lagi? (y/n): ")
		}
	}

	if t.jumlahItem == 0 {
		fmt.Println("Transaksi dibatalkan.")
		return
	}

	if adaMember {
		if t.total >= 500000 {
			t.diskon = t.total * 0.10
		} else {
			t.diskon = t.total * 0.05
		}
	}

	totalBayar := t.total - t.diskon
	fmt.Printf("Total: Rp %.0f | Diskon: Rp %.0f | Bayar: Rp %.0f\n", t.total, t.diskon, totalBayar)

	var ok bool
	ok = false
	for !ok {
		t.bayar = inputFloat("Uang bayar: Rp ")
		if t.bayar >= totalBayar {
			ok = true
		} else {
			fmt.Println("Uang kurang!")
		}
	}

	kembalian := t.bayar - totalBayar

	if adaMember {
		idx := seqMember(t.idMember)
		if idx != -1 {
			daftarMember[idx].poin = daftarMember[idx].poin + int(t.total/10000)
		}
	}

	daftarTransaksi[jumlahTransaksi] = t
	jumlahTransaksi = jumlahTransaksi + 1

	garis()
	fmt.Println("         STRUK - MINIMART SERBA ADA")
	garis()
	fmt.Printf("ID: %s\n", t.id)
	var i int
	i = 0
	for i < t.jumlahItem {
		fmt.Printf("%-20s x%d = Rp %.0f\n", t.items[i].nama, t.items[i].qty, t.items[i].subtotal)
		i = i + 1
	}
	fmt.Println("------------------------------------------------")
	fmt.Printf("Total    : Rp %.0f\n", t.total)
	fmt.Printf("Diskon   : Rp %.0f\n", t.diskon)
	fmt.Printf("Bayar    : Rp %.0f\n", t.bayar)
	fmt.Printf("Kembali  : Rp %.0f\n", kembalian)
	garis()
}

func riwayatTransaksi() {
	garis()
	fmt.Println("           RIWAYAT TRANSAKSI")
	garis()
	if jumlahTransaksi == 0 {
		fmt.Println("Belum ada transaksi.")
	} else {
		fmt.Printf("%-10s %-15s %12s %10s %15s\n", "ID", "MEMBER", "TOTAL", "DISKON", "PENDAPATAN")
		fmt.Println("------------------------------------------------")
		var i int
		var totalOmset float64
		var totalDiskon float64
		var totalPendapatan float64
		i = 0
		totalOmset = 0
		totalDiskon = 0
		totalPendapatan = 0
		for i < jumlahTransaksi {
			member := daftarTransaksi[i].idMember
			if member == "" {
				member = "Non-Member"
			}
			pendapatan := daftarTransaksi[i].total - daftarTransaksi[i].diskon
			fmt.Printf("%-10s %-15s %12.0f %10.0f %15.0f\n",
				daftarTransaksi[i].id,
				member,
				daftarTransaksi[i].total,
				daftarTransaksi[i].diskon,
				pendapatan)
			totalOmset = totalOmset + daftarTransaksi[i].total
			totalDiskon = totalDiskon + daftarTransaksi[i].diskon
			totalPendapatan = totalPendapatan + pendapatan
			i = i + 1
		}
		fmt.Println("------------------------------------------------")
		fmt.Printf("%-10s %-15s %12.0f %10.0f %15.0f\n",
			"TOTAL", fmt.Sprintf("(%d trx)", jumlahTransaksi),
			totalOmset, totalDiskon, totalPendapatan)
	}
	garis()
}

func menuTransaksi() {
	var p int
	p = -1
	for p != 0 {
		garis()
		fmt.Println("1.Transaksi Baru  2.Riwayat  0.Kembali")
		p = inputInt("Pilih: ")
		if p == 1 {
			buatTransaksi()
		} else if p == 2 {
			riwayatTransaksi()
		}
	}
}

func main() {
	daftarProduk[0] = Produk{"P001", "Indomie Goreng", 3500, 50}
	daftarProduk[1] = Produk{"P002", "Aqua 600ml", 4000, 80}
	daftarProduk[2] = Produk{"P003", "Teh Botol 350ml", 5000, 60}
	daftarProduk[3] = Produk{"P004", "Roti Tawar", 15000, 20}
	daftarProduk[4] = Produk{"P005", "Susu Ultra 250ml", 6500, 40}
	jumlahProduk = 5
	daftarMember[0] = Member{"M001", "Budi Santoso", "081234567890", 150}
	daftarMember[1] = Member{"M002", "Sari Dewi", "082345678901", 300}
	jumlahMember = 2

	var p int
	p = -1
	for p != 0 {
		garis()
		fmt.Println("     KASIR MINIMART SERBA ADA")
		garis()
		fmt.Println("1. Kelola Produk")
		fmt.Println("2. Kelola Member")
		fmt.Println("3. Transaksi")
		fmt.Println("0. Keluar")
		garis()
		p = inputInt("Pilih: ")
		if p == 1 {
			menuProduk()
		} else if p == 2 {
			menuMember()
		} else if p == 3 {
			menuTransaksi()
		}
	}
	fmt.Println("Terima kasih!")
}
