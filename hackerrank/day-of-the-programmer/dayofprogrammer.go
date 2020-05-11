package hackerrank

import "time"

func dayOfProgrammer(year int32) string {
	addDay := 255
	if year < 1917 {
		if int(year)%4 == 0 && int(year)%100 == 0 {
			addDay = 254
		} else if year == 1918 {
			addDay = 268
		}
	}
	startYear := time.Date(int(year), 1, 1, 0, 0, 0, 0, time.UTC)
	dayOfProg := startYear.AddDate(0, 0, addDay)
	return dayOfProg.Format("02.01.2006")
}
