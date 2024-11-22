package main

import "fmt"

func main() {
	var umur int
	var KK bool
	fmt.Scan(&umur, &KK)
	if umur >= 17 && KK {
		fmt.Println("bisa membuat KTP")
	} else  {
		fmt.Println("belum bisa membuat KTP")
	}
}
