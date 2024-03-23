package main

import (
	"fmt"

	"github.com/golang/packages/getCelcius"
	"github.com/golang/packages/merhaba"
)

func main() {
	merhaba.Hello()

	fmt.Print("Lütfen celcius dereceyi giriniz: ")
	celcius, err := getCelcius.GetCelcius()
	if err != nil {
		fmt.Println(err)
	}
	kelvin := celcius + 273
	fmt.Println(celcius, "celcius derecesi", kelvin, "kelvin derecesine eşittir")
}
