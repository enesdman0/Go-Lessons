package main

import "fmt"

func main() {
	// cities := [4]string{"Samsun", "İstanbul", "Tahran", "Belgrad"}
	// cities := ]string{"Samsun", "İstanbul", "Tahran", "Belgrad"}
	// 	cities := [...]string{"Samsun", "İstanbul", "Tahran", "Belgrad"}

	// 	for i := 0; i < len(cities); i++ {
	// 		fmt.Println(i+1,cities[i])
	// 	}

	myArr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sonuc := kareAl(myArr)
	fmt.Println(sonuc)
}
func kareAl(arr []int) []int {
	// for i := 0; i < len(arr); i++ {
	// 	arr[i] = arr[i] * arr[i]
	// }
	for _, sayi := range arr {
		arr[sayi] = arr[sayi] * arr[sayi]
	}
	return arr
}
