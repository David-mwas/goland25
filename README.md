# Go Codebase Basics

This repository contains a simple Go program that calculates the user's age based on their birth year and displays a list of hobbies.

## Getting Started

### Prerequisites

- Go installed on your machine. You can download it from [here](https://golang.org/dl/).

### Running the Program

1. Clone the repository:

   ```sh
   git clone https://github.com/David-mwas/goland25.git
   cd gocodebasebasics
   ```

2. Run the program:
   ```sh
   go run . or go run main.go
   ```

### Program Description

The program performs the following tasks:

1. Prompts the user to enter their name.
2. Validates the entered name.
3. Prompts the user to enter their birth year.
4. Validates the entered birth year.
5. Calculates the user's age based on the current year.
6. Displays the user's age.
7. Displays a list of predefined hobbies and highlights if the user loves coding.

### Code

```go
package main

import (
    "fmt"
    "time"
)

// function to calculate age
func calculateAge(birthYear int) int {
    return time.Now().Year() - birthYear
}

func main() {
    var name string
    var birthYear int

    for {
        fmt.Println("Enter your name: ")
        fmt.Scan(&name)

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

        fmt.Printf("Thank you %v your age is %v 🚀 \n", name, calculateAge(birthYear))

        hobbies := []string{"coding", "reading", "traveling", "swimming", "running"}

        fmt.Println("Your hobbies are: ")

        for i := 0; i < len(hobbies); i++ {
            if hobbies[i] == "coding" {
                fmt.Println("You love coding")
            }

            fmt.Println(hobbies[i])
        }
    }
}
```

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### Acknowledgments

- [Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [The Go Programming Language](https://golang.org/)
