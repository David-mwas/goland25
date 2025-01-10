package main

import "time"

func calculateAge(birthYear int) int {
	return time.Now().Year() - birthYear
}
