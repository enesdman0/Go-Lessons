package main

import "fmt"

type myType int
type rectangle struct {
	a int
	b int
}
type family struct {
	name      string
	age       int
	isMarried bool
}

func showFamily() []family {
	father := family{
		name:      "İsmail",
		age:       53,
		isMarried: true,
	}
	mother := family{
		"Ayşe",
		53,
		true,
	}
	var me family
	me.name = "Enes"
	me.age = 20
	me.isMarried = false
	return []family{father, mother, me}

}
func (r rectangle) display() {
	fmt.Println("Kenar uzunlukları ", r.a, r.b)
}
func (r rectangle) area() int {
	return r.a * r.b
}
func (r rectangle) cirrcumference() int {
	return 2 * (r.a + r.b)
}
func main() {
	// 1 -) Underlying Type 'int' olacak şekilde kendi veri tipinizi oluşturunuz veri tipi ve değerini '10' yazdırınız.
	// var x myType = 10
	// fmt.Printf("%T %v", x, x)
	// 2 -) Başlangıç değeri 10 olan bir X değişkeni oluşturun sonrasında bulunduğu adres üzerinden y değişkenini tanımlayıp x değerini 20 yapınız.
	// x := 10
	// y := (&x)
	// *y = 20
	// fmt.Println(x)
	// 3 -) Underlying Type struct olan Rectangle type oluşturunuz. Display, area, circumference metodlarını yazınız.
	// r1 := rectangle{4, 8}
	// r1.display()
	// fmt.Println(r1.area())
	// fmt.Println(r1.cirrcumference())
	// 4-) family aile bireyleri şeklinde veri tipi oluşturalım, underlying struct. Aile bireylerinin hepsini farklı şekilde tanımlayınız. Sonrasında for döngüsünde yazdırınız. field age, name, isMarried field.
	families := showFamily()
	for i := 0; i < len(families); i++ {
		fmt.Println(families[i])
	}
}
