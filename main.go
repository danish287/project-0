package main

import (
	"flag"
	//"github.com/danish287/project-0/shaker"
)

//learned init and main etc.
//1. mvp 2. restructuring hello

//get user input from user of zodiac sign
//get user input for love, career, or money
//get user input for weekly / monthly and yearly
// default horoscope is daily general horoscope

var zodiac [12]string
var userZodiac string

func init() {
	// zodiac = [12]string{

	// }

	//flag.String(&userZodiac, "", "get user zodiac")
	temp := flag.Arg(0)
	//flag.Parse()
	println("SUER ZODIAC", temp)

}

func main() {
	//fmt.Println(zodiac[0])
}
