package main

import "flag"

//Zodiac stores name of zodiac
var Zodiac string

//ReadingType stores type of reading user is looking get
var ReadingType string

//ReadingFor specifies when the reader wants the reading
var ReadingFor string

func init() {
	flag.StringVar(&Zodiac, "Zodiac", "EMPTY", "name of the zodiac sign")
	flag.StringVar(&ReadingFor, "ReadingFor", "daily", "time period of reading")
	flag.StringVar(&ReadingType, "Type", "general", "type of reading")
	flag.Parse()
}
