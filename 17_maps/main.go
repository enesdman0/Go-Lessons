package main

import "fmt"

func main() {
	// myMap := map[string]int{
	// 	"Ahmet":  70,
	// 	"Mehmet": 30,
	// 	"Enes":   20,
	// }
	// isMarried := map[string]bool{
	// 	"Ahmet":  true,
	// 	"Mehmet": false,
	// 	"Enes":   false,
	// }
	// fmt.Println(myMap)
	// fmt.Println(myMap["Ahmet"])

	// fmt.Println(isMarried)
	// fmt.Println(isMarried["Ahmet"])

	// myMap := make(map[string]int)
	// fmt.Println(myMap)

	studentGrades := map[string]int{
		"Enes":  100,
		"Buse":  100,
		"Merve": 80,
	}
	fmt.Println(studentGrades)
	// // fmt.Println(studentGrades["Merv"])
	// value, ok := studentGrades["Enegs"]
	// if !ok {
	// 	fmt.Println("böyle bir örğenci yok")
	// } else {
	// 	fmt.Println(value)
	// }

	// studentGrades["Sıla"] = 95
	// fmt.Println(studentGrades)
	// delete(studentGrades, "Merve")
	// fmt.Println(studentGrades)
	// fmt.Println(len(studentGrades))

	sg:=studentGrades
	fmt.Println(sg)
	delete(sg,"Enes")
	fmt.Println(studentGrades)
	fmt.Println(sg)
}
