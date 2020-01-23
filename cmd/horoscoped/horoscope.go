package main

import (
	"fmt"

	"github.com/danish287/project-0/shaker"
)

func main() {
	if Zodiac == "EMPTY" {
		fmt.Println("\nPlease enter a zodiac sign using the flag Zodiac. \nExample: -Zodiac=Aquarius")
	}
	zodiacURL := shaker.GetZodiacURL(Zodiac, ReadingFor, ReadingType)
	if zodiacURL == "Please try again using valid arguments." {
		fmt.Println("\nPlease try again using valid arguments.\n")
	} else {
		answer := shaker.GetMyResponse(zodiacURL, ReadingFor)

		fmt.Println("\n\n ", answer)
	}

}
