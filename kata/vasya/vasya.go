package kata

// Tickets return YES if you can keep selling or No if you dont
func Tickets(peopleInLine []int) string {
	money := 0
	for _, v := range peopleInLine {
		money += v * v
	}
	if money%2 == 25 {
		return "YES"
	}
	return "NO"
}
