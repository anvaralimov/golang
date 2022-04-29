package main

import "fmt"

func main() {
	var number int = 34

	if number > 30 {
		fmt.Println(number, " katta 30 dan")
	} else if number < 45 {
		fmt.Println(number, " kichik 45 dan")
	} else {
		fmt.Println("noma'lum qiymat")
	}
}
