# Introduction to Testing in Go (Golang)

## Introduction

[Udemy](https://www.udemy.com/course/introduction-to-testing-in-go-golang/learn/lecture/33522068#overview)

## Setup

[Go](https://go.dev/)

## Simple testing

### Cli application

```go
package main

import "fmt"

func main() {
	n := 2

	_, msg := isPrime(n)
	fmt.Println(msg)
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime by definition
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	}

	// negative numbers are not prime
	if n < 0 {
		return false, "Negative numbers are not prime, by definition!"
	}

	// use the modulus operator repeatedly to see if we have a prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			// not a prime number
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", n)
}
```

### First test

```go
package main

import "testing"

func Test_isPrime(t *testing.T) {
	result, msg := isPrime(0)
	if result {
		t.Errorf("with %d as test parameter: expected false", 0)
	}
	if msg != "0 is not prime, by definition!" {
		t.Error("wrong message returned:", msg)
	}

	result, msg = isPrime(7)
	if !result {
		t.Errorf("with %d as test parameter: expected true", 7)
	}
	if msg != "7 is a prime number!" {
		t.Error("wrong message returned:", msg)
	}
}
```

```sh
go test .
go test -v .
```

### Table testing

```go
package main

import "fmt"

func main() {
	n := 2

	_, msg := isPrime(n)
	fmt.Println(msg)
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime by definition
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	}

	// negative numbers are not prime
	if n < 0 {
		return false, "Negative numbers are not prime, by definition!"
	}

	// use the modulus operator repeatedly to see if we have a prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			// not a prime number
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d!", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", n)
}
```

### Test coverage

```sh
go test -cover
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Cmpleting table tests

```go
package main

import "testing"

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{
			"prime", 7, true, "7 is a prime number!",
		},
		{
			"not prime", 8, false, "8 is not a prime number because it is divisible by 2!",
		},
		{
			"zero", 0, false, "0 is not prime, by definition!",
		},
		{
			"one", 1, false, "1 is not prime, by definition!",
		},
		{
			"one", -1, false, "Negative numbers are not prime, by definition!",
		},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}
		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}
```

### User entered information

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// print welcome message
	intro()

	// channel to indicate user quits
	doneChan := make(chan bool)

	// start goroutine to read user input
	go readUserInput(doneChan)

	// wait for channel value
	<-doneChan

	// close channel
	close(doneChan)

	// say goodbye
	fmt.Println("Ciao!")
}

func checkNumber(s *bufio.Scanner) (string, bool) {
	s.Scan()
	if strings.EqualFold(s.Text(), "q") {
		return "", true
	}
	numToCheck, err := strconv.Atoi(s.Text())
	if err != nil {
		return "Please enter a whole number!", false
	}

	_, msg := isPrime(numToCheck)

	return msg, false
}

func readUserInput(doneChan chan bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		res, done := checkNumber(scanner)
		if done {
			doneChan <- true
			return
		}
		fmt.Println(res)
		prompt()
	}
}

func intro() {
	fmt.Println("Is Prime?")
	fmt.Println("=========")
	fmt.Println("Enter a whole number")
	prompt()
}

func prompt() {
	fmt.Print("-> ")

}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime by definition
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	}

	// negative numbers are not prime
	if n < 0 {
		return false, "Negative numbers are not prime, by definition!"
	}

	// use the modulus operator repeatedly to see if we have a prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			// not a prime number
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d!", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", n)
}
```

### Test prompt
