package config

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

//ZodiacSign stores name of zodiac
var ZodiacSign string

//ReadingType stores type of reading user is looking get
var ReadingType string

//ReadingFor specifies when the reader wants the reading
var ReadingFor string

//CONFIGFILE stores the links where we get horoscopes
const CONFIGFILE string = "conf.json"

//const CONFIGFILE string = "../../conf.json"

//MyURL stores the URL we are trying to get
var MyURL string

//LinkCluster derives Links from json file
type LinkCluster struct {
	General []string
	Love    []string
	Career  []string
	Money   []string
}

//ZodiacCluster derives links from json file
type ZodiacCluster struct {
	Zodiac   string
	Readings []LinkCluster
}

//Configuration derives links from json file
type Configuration struct {
	Links []ZodiacCluster
}

var ConfigMe Configuration

func init() {

	myLink, err := os.Open(CONFIGFILE)
	json.NewDecoder(myLink).Decode(&ConfigMe)

	if err != nil {

		log.Fatal(err)
	}
	flag.StringVar(&ZodiacSign, "ZodiacSign", "EMPTY", "name of the zodiac sign")
	flag.StringVar(&ReadingFor, "ChooseReading", "daily", "time period of reading")
	flag.StringVar(&ReadingType, "ChooseType", "general", "type of reading")
	flag.Parse()

}
