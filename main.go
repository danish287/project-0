package main

import (
	"fmt"
	"github.com/danish287/project-0/shaker"
)

func main() {
	if Zodiac == "EMPTY" {
		fmt.Println("\nPlease enter a zodiac sign using the flag Zodiac. \nExample: -Zodiac=Aquarius")
	}
	zodiacURL := shaker.GetZodiacURL(Zodiac, ReadingFor)
	answer := shaker.GetMyResponse(zodiacURL, ReadingFor)

	fmt.Println("\n\n ", answer)

}
