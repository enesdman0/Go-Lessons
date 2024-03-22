package main

func main() {
	// 1
	// topla, fark, carp, bol := dortIslem(4, 5)
	// fmt.Println("Toplam :", topla)
	// fmt.Println("Fark :", fark)
	// fmt.Println("Çarpma :", carp)
	// fmt.Println("Bolme :", bol)

	// 2
	// fmt.Print("Lütfen notunuzu giriniz: ")
	// grade, _ := notHesapla()
	// fmt.Println(grade)

	// 3
	// target := sayiTahmin(1, 100)
	// fmt.Println("1 ila 100 arasındaki sayıyı bulmaya çalışın")
	// reader := bufio.NewReader(os.Stdin)

	// for attemps := 0; attemps < 10; attemps++ {
	// 	fmt.Println(10-attemps, "hakkınız kaldı")
	// 	fmt.Print("Lütfen tahmininizi giriniz: ")
	// 	input, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	input = strings.TrimSpace(input)
	// 	num, err := strconv.Atoi(input)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	if num == target {
	// 		fmt.Println("dogru")
	// 		break
	// 	} else {
	// 		if num < target {
	// 			fmt.Println("Daha büyük tahmin yapınız")
	// 		} else {
	// 			fmt.Println("Daha küçük tahmin yapınız")
	// 		}
	// 	}
	// }
}

// 1
// func dortIslem(x, y int) (topla int, fark int, carp int, bol float64) {
// 	topla = x + y
// 	fark = x - y
// 	carp = x * y
// 	bol = float64(x) / float64(y)
// 	return topla, fark, carp, bol
// }

// 2
// func notHesapla() (int, error) {
// 	reader := bufio.NewReader(os.Stdin)
// 	input, err := reader.ReadString('\n')
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	input = strings.TrimSpace(input)
// 	num, err := strconv.Atoi(input)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	if num < 50 {
// 		fmt.Println("Kaldınız")
// 	} else {
// 		fmt.Println("Geçtiniz")
// 	}
// 	return num, nil
// }

// 3
func sayiTahmin(min, max int) int {
	// rand.Seed(time.Now().Unix())
	// return rand.Intn(max-min) + min
}
