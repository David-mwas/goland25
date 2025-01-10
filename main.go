package main

import (
	"fmt"
	"goland25/helper"
	"time"
)

func getUserInput() (string, int) {
	var name string
	var birthYear int

	fmt.Println("Enter your name: ")
	fmt.Scan(&name)

	if name == "" {
		fmt.Println("Invalid name")
		return "", 0
	}

	if len(name) < 3 {
		fmt.Println("Name is too short")
		return "", 0
	}

	fmt.Println("Enter your birth year: ")
	fmt.Scan(&birthYear)

	if birthYear == 0 || birthYear > time.Now().Year() {
		fmt.Println("Invalid birth year")
		return "", 0
	}

	return name, birthYear
}

func printHobbies() {
	hobbies := []string{"coding", "reading", "traveling", "swimming", "running"}
	fmt.Println("Your hobbies are: ")

	for _, hobby := range hobbies {
		if hobby == "coding" {
			fmt.Println("You love coding")
		}
		fmt.Println(hobby)
	}
}

func main() {
	for {
		name, birthYear := getUserInput()
		if name == "" || birthYear == 0 {
			return
		}

		fmt.Printf("Thank you %v, your age is %v 🚀 \n", name, helper.CalculateAge(birthYear))
		printHobbies()
	}
}
