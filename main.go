package main

import (
	"fmt"
	"time"
)

// function i go
func calculateAge(birthYear int) int {

	// return the age from current year
	return time.Now().Year() - birthYear
}

func main() {
	// variables
	var name string
	var birthYear int

	// infinite loop
	for {
		// get user input
		fmt.Println("Enter your name: ")
		fmt.Scan(&name)

		// validate name
		if name == "" {
			fmt.Println("Invalid name")
			return
		}

		if len(name) < 3 {
			fmt.Println("Name is too short")
			return
		}

		fmt.Println("Enter your birth year: ")

		fmt.Scan(&birthYear)

		if birthYear == 0 || birthYear > time.Now().Year() {
			fmt.Println("Invalid birth year")
			return
		}

		// print the result
		fmt.Printf("Thank you %v your age is %v 🚀 \n", name, calculateAge(birthYear))

		// array of strings
		hobbies := []string{"coding", "reading", "traveling", "swimming", "running"}

		// print the hobbies
		fmt.Println("Ur hobbies are: ")

		// loop through the hobbies
		for i := 0; i < len(hobbies); i++ {
			if hobbies[i] == "coding" {
				fmt.Println("U love coding")
				// continue

			}

			fmt.Println(hobbies[i])
		}

	}
}

// go run main.go
