package main

import "fmt"

type atas struct {
	nama   string
	kelas  string
	hoby string
}
type bawah struct {
	atas
	umur    int
	npm     int
}

func main() {
	var array = bawah{}

	fmt.Print("Masukkan Nama : ")
	fmt.Scan(&array.nama)
	fmt.Print("Masukkan Kelas : ")
	fmt.Scan(&array.kelas)
	fmt.Print("Masukkan Hoby  : ")
	fmt.Scan(&array.hoby)
	fmt.Print("Masukkan Umur  : ")
	fmt.Scan(&array.umur)
	fmt.Print("Masukkan NPM  : ")
	fmt.Scan(&array.npm)

	fmt.Println("\n#### #### #### #### ####")
	fmt.Println("Nama :", array.nama)
	fmt.Println("Kelas :", array.kelas)
	fmt.Println("Hoby :", array.hoby)
	fmt.Println("Umur :", array.umur)
	fmt.Println("NPM :", array.npm)
	fmt.Println("#### #### #### #### ####\n")
}
