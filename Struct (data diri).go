package main

import "fmt"

type data struct {
    nama string
    npm string
    hoby string
    umur string
    grade string
}
func main() {
    var dataku = data {}
	fmt.Print("Masukkan Nama : ")
	fmt.Scan(&dataku.nama)
	fmt.Print("Masukkan NPM : ")
	fmt.Scan(&dataku.npm)
	fmt.Print("Masukkan Hoby : ")
	fmt.Scan(&dataku.hoby)
	fmt.Print("Masukkan Umur : ")
	fmt.Scan(&dataku.umur)
	fmt.Print("Masukkan Grade : ")
	fmt.Scan(&dataku.grade)
	fmt.Print("======================")
	fmt.Println("Nama aku : ",dataku.nama)
	fmt.Println("NPM : ",dataku.npm)
	fmt.Println("Hobi Aku : ",dataku.hoby)
	fmt.Println("Umur : ",dataku.umur)
	fmt.Println("Grade : ",dataku.grade)
}
