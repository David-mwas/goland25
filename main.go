package main

import (
	"fmt"
	"goland25/helper"
	"sync"
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

func sendDetails() {
	// simulate blocking operation
	time.Sleep(10 * time.Second)
	fmt.Println("Details sent successfully")

	// decrement the wait group counter
	wg.Done()
}

// wait group to wait for all go routines to finish
var wg = sync.WaitGroup{}

func main() {
	for {
		name, birthYear := getUserInput()
		if name == "" || birthYear == 0 {
			break
		}

		// increment the wait group counter
		wg.Add(1)
		// go routine - non-blocking, concurrency
		go sendDetails()

		fmt.Printf("Thank you %v, your age is %v 🚀 \n", name, helper.CalculateAge(birthYear))

		printHobbies()

	}

	// wait for all go routines to finish
	// wg.Wait()
}
