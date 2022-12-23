package timeConversion

import (
	"strconv"
)

func TimeConversion(s string) string {
	dayTimeSep := len(s) - 2
	hourSeparator := 2
	hour, mins, ending := s[:hourSeparator], s[hourSeparator:dayTimeSep], s[dayTimeSep:]
	numHour, _ := strconv.Atoi(hour)
	if ending == "AM" && numHour == 12 {
		hour = "00"
	}
	if ending == "PM" && numHour != 12 {
		hour = strconv.Itoa(numHour + 12)
	}
	return hour + mins
}
