package main

import (
	"fmt"
)

func main() {
	var parsel, ril_berat, berat, biaya, sisa int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&ril_berat)
	parsel = ril_berat / 1000
	berat = ril_berat % 1000
	fmt.Printf("Detail berat: %d kg + %d gr\n", parsel, berat)
	biaya = parsel * 10000
	if ril_berat > 10 {
		sisa = 0
	} else if berat < 500 {
		sisa = berat * 15
	} else if berat >= 500 {
		sisa = berat * 5
	}
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d \n", biaya, sisa)
	fmt.Printf("Total biaya: Rp. %d", biaya+sisa)
}
