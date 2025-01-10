package helper

import "time"

func CalculateAge(birthYear int) int {
	return time.Now().Year() - birthYear
}
