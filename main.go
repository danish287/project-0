package main

import "fmt"

//learned init and main etc.
//1. mvp 2. restructuring

//get user input from user to
var zodiac [12]string

func init() {
	zodiac = [12]string{
		"aquarius",
		"pices",
		"aries",
		"taurus",
		"gemini",
		"cancer",
		"leo",
		"virgo",
		"libra",
		"scorpio",
		"sagittarius",
		"capricorn",
	}

}

func main() {
	fmt.Println(zodiac[0])
}
