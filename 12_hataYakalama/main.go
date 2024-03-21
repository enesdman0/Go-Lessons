package main

import (
	"fmt"
	"os"
)

func main() {
	// Çift sayı kontrol
	// result, err := goster(4)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// }

	// Karekök
	// 	result, err := karekok(56)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println(result)
	// 	}

	// Dosya açma
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Dosya", file)
	}
}

// Çift sayı kontrol
// func goster(num int) (int, error) {
// 	if num%2 != 0 {
// 		return 0, errors.New("Çift sayı girmediniz")
// 	}
// 	return num, nil
// }

// Karekök
// func karekok(x float64) (float64, error) {
// 	if x < 0 {
// 		return 0, errors.New("Negatif sayıların karekökü alınamaz. lütfen pozitif bir sayı giriniz")
// 	}
// 	result := math.Sqrt(x)
// 	return result, nil
// }
