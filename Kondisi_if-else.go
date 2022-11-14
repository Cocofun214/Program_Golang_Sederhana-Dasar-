package main
import "fmt"
func main() {
    var i int
	fmt.Print("Masukkan Angka : ")
	fmt.Scanf("%d",&i)
	if i == 1 {
	    fmt.Print("Saya Sedang Makan")
	}else if i == 2 {
		fmt.Print("Saya Sedang Belajar")
	}else if i == 3 {
		fmt.Print("Saya Sedang Bersantai")
	}else{
	    fmt.Print("Pilihan Anda Tidak Ada")
	}
}
	
