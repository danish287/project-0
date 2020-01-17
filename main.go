package main

import (
	"fmt"

	"github.com/danish287/project-0/shaker"
)

func main() {
	if Zodiac == "EMPTY" {
		fmt.Println("\nPlease enter a zodiac sign using the flag Zodiac. \nExample: -Zodiac=Aquarius")
	}

	// if ReadingType == "general" {
	// 	fmt.Println("general reading")
	// } else {
	// 	fmt.Println(ReadingType)
	// }

	if ReadingFor == "yearly" {
		fmt.Println(shaker.GetZodiac(Zodiac, ReadingFor))
	} else {
		fmt.Println(ReadingFor)
	}

}
