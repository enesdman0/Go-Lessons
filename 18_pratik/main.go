package main

import "fmt"

func main() {
	// 1 -) 5 string elemandan oluşan bir array ve 5 int elemandan oluşan slice oluşturup index numaralarıyla birlikte gösterelim.
	// myArr := [5]string{"Samsun", "Ankara", "Bolu", "İstanbul", "Ulanbatur"}
	// mySlc := []int{1, 2, 3, 4, 5}

	// for index, value := range myArr {
	// 	fmt.Println(index+1, value)
	// }

	// for _, value := range mySlc {
	// 	fmt.Println(value)
	// }
	// 2 -) [0,1,2,3,4,5,6,7,8] slice dan [0,1,2,3], [4,5,6], [2,3,6,7] slicelarını oluşturunuz.
	// mySlc := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	// mySlc2 := mySlc[:5]
	// mySlc3 := mySlc[4:7]
	// mySlc4 := append(mySlc[2:4], mySlc[6:8]...)

	// fmt.Println(mySlc)
	// fmt.Println(mySlc2)
	// fmt.Println(mySlc3)
	// fmt.Println(mySlc4)
	// 3 -) slicelar için copy metodunu ve assign ( = ) ile farkını açıklayınız.

	// 4 -) map gösterimi ile yazar ve yazarlara ait kitapların isimlerini for range ile gösterin.
	myMap := map[string][]string{
		"Yaşar Kemal":     []string{"İnce Memed", "Yer Demir Gök Bakır"},
		"Sabahattin Ali":  []string{"Kuyucaklı Yusuf", "Kürk Mantolu Madonna", "Değirmen"},
		"Haruki Murakami": []string{"1984", "Dans Dans Dans", "Kumandanı Öldürmek"},
	}
	for key, value := range myMap {
		fmt.Println("Yazarımız: ", key)
		fmt.Println("Bazı kitapları aşağıdadır:")
		for i, v := range value {
			fmt.Println("\t", i+1, v)
		}
	}

}
