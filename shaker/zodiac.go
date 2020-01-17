package shaker

//CONFIGFILE stores the links where we get horoscopes
const CONFIGFILE string = "conf.json"

type Configuration struc {
	string
}

//GetZodiac is a function that takes a string of integers and returns the first thing in tthe listt
func GetZodiac(zodiac string, when string) int {
	linkNum := GetType(when)

	return linkNum
}

//GetType returns an integer referring to the type of reading the user wants
func GetType(when string) int {
	readingLink := []string{"daily", "weekly", "monthly", "yearly"}
	const (
		Daily = iota
		Weekly
		Monthly
		Yearly
	)
	switch {
	case readingLink[Daily] == when:
		return Daily
	case readingLink[Weekly] == when:
		return Weekly
	case readingLink[Monthly] == when:
		return Monthly
	case readingLink[Yearly] == when:
		return Yearly
	}
	return -1

}
