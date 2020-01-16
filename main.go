package main

import (
	"fmt"
)

func main() {
	if Horoscope == "EMPTY" {
		fmt.Println("Please enter a zodiac sign using the flag Ho")
	} else {
		fmt.Println(Horoscope)
	}

}
