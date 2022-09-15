package main

import "fmt"

func main() {
	var harga int
	var diskon int
	var Hasil int
	fmt.Println("total belanja: ")
	fmt.Scanf("%d", &harga)

	fmt.Scanf("%d", &diskon)
	if harga <= 50000 {
		diskon = 25
		Hasil = harga - (harga * diskon / 100)
		fmt.Println("diskon :25% ")
		fmt.Println("total pembayaran", Hasil)

	} else if harga >= 70000 {
		diskon = 50
		Hasil = harga - (harga * diskon / 100)
		fmt.Println("diskon :50% ")
		fmt.Println("total pembayaran", Hasil)
	}

}
