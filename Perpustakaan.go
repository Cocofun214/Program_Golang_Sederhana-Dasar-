package main

import(
    "fmt"
	"os"
	"os/exec"
)
type datapeminjam struct{
	kategoribuku string
	namabuku string
	namapeminjam string
	nim int
	tanggalpeminjaman string
	tanggalpengembalian string
}
type arraypeminjaman [999]datapeminjam

var enter string = ""
var pilihan int

func CLS(){
    c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}
func main(){
    var peminjaman arraypeminjaman
	var n int
	CLS()
	welcome(&peminjaman, &n)
}
func tampilan(){
    var text1, text2 string
	text1 = "* * * * * * * SELAMAT DATANG * * * * * * *"
	text2 = "* * * * * PERPUSTAKAAN SEDERHANA * * * * *"
	fmt.Println(text1)
	fmt.Println(text2, "\n")
}
func welcome(peminjaman *arraypeminjaman, n *int) {
    var namabuku string
	var nama string
	var nim int
	var tanggalpeminjaman string
	var tanggalpengembalian string
	CLS()
	tampilan()
	
	fmt.Println("\nPILIHAN MENU")
	fmt.Println("\n[1] BUKU PELAJARAN")
	fmt.Println("[2] KOMIK")
	fmt.Println("[3] DONGENG")
	fmt.Println("[4] DATA PEMINJAM ")
	fmt.Println("[5] CARI DATA")
	fmt.Println("")
	fmt.Println("\nPILIHAN ANDA : ")
	fmt.Scanln(&pilihan)
	
	if pilihan == 1 {
	    CLS()
		tampilan()
		    fmt.Print("\nApakah Anda Yakin ? (Y/N) : ")
			fmt.Scanln(&enter)
			if enter == "Y"  {
			    insertnamabuku(&namabuku)
				insertnama(&nama)
				insertnim(&nim)
				inserttanggalpeminjaman(&tanggalpeminjaman)
				inserttanggalpengembalian(&tanggalpengembalian)
				peminjaman[*n].kategoribuku = "PELAJARAN"
				peminjaman[*n].namabuku = namabuku
				peminjaman[*n].namapeminjam = nama
				peminjaman[*n].nim = nim
				peminjaman[*n].tanggalpeminjaman = tanggalpeminjaman
				peminjaman[*n].tanggalpengembalian = tanggalpengembalian
				*n++
			}
			CLS()
			tampilan()
			welcome(peminjaman, n)
	} else if pilihan == 2{
	    CLS()
		tampilan()
		    fmt.Print("\nApakah Anda Yakin ? (Y/N) : ")
			fmt.Scanln(&enter)
			if enter == "Y"  {
			    insertnamabuku(&namabuku)
				insertnama(&nama)
				insertnim(&nim)
				inserttanggalpeminjaman(&tanggalpeminjaman)
				inserttanggalpengembalian(&tanggalpengembalian)
				peminjaman[*n].namabuku = namabuku
				peminjaman[*n].namapeminjam = nama
				peminjaman[*n].nim = nim
				peminjaman[*n].tanggalpeminjaman = tanggalpeminjaman
				peminjaman[*n].tanggalpengembalian = tanggalpengembalian
				*n++
			}
			CLS()
			tampilan()
			welcome(peminjaman, n)
	} else if pilihan == 3{
	    CLS()
		tampilan()
		    fmt.Print("\nApakah Anda Yakin ? (Y/N) : ")
			fmt.Scanln(&enter)
			if enter == "Y"  {
			    insertnamabuku(&namabuku)
				insertnama(&nama)
				insertnim(&nim)
				inserttanggalpeminjaman(&tanggalpeminjaman)
				inserttanggalpengembalian(&tanggalpengembalian)
				peminjaman[*n].namabuku = namabuku
				peminjaman[*n].namapeminjam = nama
				peminjaman[*n].nim = nim
				peminjaman[*n].tanggalpeminjaman = tanggalpeminjaman
				peminjaman[*n].tanggalpengembalian = tanggalpengembalian
				*n++
			}
			CLS()
			tampilan()
			welcome(peminjaman, n)
	} else if pilihan == 4{
	    showdata(*peminjaman, *n)
		fmt.Scanln()
		CLS()
		tampilan()
		welcome(peminjaman, n)
	} else if pilihan == 5 {
	    fmt.Println("Masukkan NIM : ")
		fmt.Scanln(&nim)
		searchdata(*peminjaman, *n, nim)
	} else if pilihan == 6 {
	    mSort(peminjaman, *n)
		showdata(*peminjaman, *n)
		CLS()
		tampilan()
		welcome(peminjaman, n)
	}else {
	    CLS()
		tampilan()
		fmt.Println("\nPilihan Yang Anda Masukkan Tidak Ada!")
		fmt.Println("\nTekan Enter Untuk Kembali : ")
		fmt.Scanln(&enter)
		if enter == " "|| enter != " " {
		}
		CLS()
		tampilan()
		welcome(peminjaman, n)
	}
}
func insertnamabuku(namabuku *string) {
    fmt.Print("NAMA BUKU : ")
	fmt.Scanln(&*namabuku)
}
func insertnama(nama *string) {
    fmt.Print("Nama Anda : ")
	fmt.Scanln(&*nama)
}
func insertnim(nim *int) {
    fmt.Print("NIM Anda : ")
	fmt.Scanln(&*nim)
}
func inserttanggalpeminjaman(tanggalpeminjaman *string) {
    fmt.Print("Tanggal Peminjaman : ")
	fmt.Scanln(&*tanggalpeminjaman)
}
func inserttanggalpengembalian(tanggalpengembalian *string) {
    fmt.Print("Tanggal Pengembalian : ")
	fmt.Scanln(&*tanggalpengembalian)
}
func showdata(peminjaman arraypeminjaman, n int) {
    for i := 0; i < n; i++ {
	    fmt.Println("Kategori Buku : ",peminjaman[i].kategoribuku)
		fmt.Println("Nama Buku : ",peminjaman[i].namabuku)
		fmt.Println("Nama Peminjam : ",peminjaman[i].namapeminjam)
		fmt.Println("NIM : ",peminjaman[i].nim)
		fmt.Println("Tanggal Peminjaman : ",peminjaman[i].tanggalpeminjaman)
		fmt.Println("Tanggal Pengembalian : ",peminjaman[i].tanggalpengembalian)
		fmt.Println()
	}
}
func searchdata(peminjaman arraypeminjaman, n int, peminjam int) {
    var found bool
	found = false
	for i := 0; i < n && found == false; i++ {
	    if peminjam == peminjaman[i].nim {
		    found = true
			fmt.Println("Nama Buku : ",peminjaman[i].namabuku)
			fmt.Println("Nama Peminjam : ",peminjaman[i].namapeminjam)
			fmt.Println("NIM : ",peminjaman[i].nim)
			fmt.Println("Tanggal Peminjaman : ",peminjaman[i].tanggalpeminjaman)
			fmt.Println("Tanggal Pengembalian : ",peminjaman[i].tanggalpengembalian)
			fmt.Println()
		}
	}
}
func mSort(peminjaman *arraypeminjaman, n int) {
    for i := 0; i < n-1; i++ {
	    min := i
		for j := i + 1; j < n; j++ {
		    if peminjaman[j].namabuku < peminjaman[min].namabuku {
			    min = j
			}
		}
		temp := peminjaman[i]
		peminjaman[i] = peminjaman[min]
		peminjaman[min] = temp
	}
}
