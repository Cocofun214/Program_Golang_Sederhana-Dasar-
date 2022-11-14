package main
import "fmt"
func main() {
    var angka int
    fmt.Println(" Berapa Jumlah Elemen Pada Array Elemen Di Dalam Array : ")
    fmt.Scan(&angka)
    fmt.Println("Masukkan Elemen-Elemen  Array : ")
    array := make([]int, angka)
	tambah := 0
    for i := 0; i < angka; i++ {
        fmt.Scan(&array[i])
		tambah = tambah + array[i]
    }
    rata := float64(tambah) / float64(angka)
	fmt.Printf("\nRata-Rata Nya Adalah : %f",rata)
}
