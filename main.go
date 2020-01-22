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

	if ReadingFor == "daily" {
		var temp string = shaker.GetZodiacURL(Zodiac, ReadingFor)
		answer := shaker.GetMyResponse(temp)

		//fmt.Printf("THISSS %T\n", temp)
		fmt.Println("MY ANSWER \n\n ", answer)
	} else {
		// var temp string = shaker.GetZodiacURL(Zodiac, "weekly")
		// answer := shaker.GetMyResponse(temp)

		// fmt.Printf("THISSS %T\n", temp)
		// fmt.Printf("SECOND ", answer)
		fmt.Println("YIKES")
	}

}
