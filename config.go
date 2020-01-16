package main

import "flag"

//Horoscope stores name of zodiac
var Horoscope string

func init() {
	flag.StringVar(&Horoscope, "Horoscope", "EMPTY", "name of the zodiac sign")
	flag.Parse()

}
