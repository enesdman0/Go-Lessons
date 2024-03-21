package main

import (
	"fmt"
)

func main() {
	// var x int = 10
	// var moment time.Time = time.Now().Add(time.Minute * 5)
	// fmt.Println(x)
	// fmt.Println(moment.Minute())

	// fmt.Print("Lütfen sınav sonucunu gir: ")
	// reader := bufio.NewReader(os.Stdin)
	// value, _ := reader.ReadString('\n')
	// fmt.Println(value)
	
	bolum, kalan := bolme(45, 4)
	fmt.Printf("Bolum %d kalan %d", bolum, kalan)
}

func bolme(x, y int) (bolum int, kalan int) {
	bolum = x / y
	kalan = x % y
	return bolum, kalan
}
