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

```go
package main

import (
	"io"
	"os"
	"testing"
)

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

func Test_prompt(t *testing.T) {
	oldOut := os.Stdout

	// create read and write pipe
	r, w, _ := os.Pipe()
	// set stdout to our write pipe
	os.Stdout = w

	prompt()

	// close and reset
	_ = w.Close()
	os.Stdout = oldOut

	// read output
	out, _ := io.ReadAll(r)
	if string(out) != "-> " {
		t.Errorf("incorrect prompt: exptected -> but got %s", string(out))
	}
}
```

### Test intro

```go
package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
)

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

func Test_prompt(t *testing.T) {
	oldOut := os.Stdout

	// create read and write pipe
	r, w, _ := os.Pipe()
	// set stdout to our write pipe
	os.Stdout = w

	prompt()

	// close and reset
	_ = w.Close()
	os.Stdout = oldOut

	// read output
	out, _ := io.ReadAll(r)
	if string(out) != "-> " {
		t.Errorf("incorrect prompt: exptected -> but got %s", string(out))
	}
}

func Test_info(t *testing.T) {
	oldOut := os.Stdout

	// create read and write pipe
	r, w, _ := os.Pipe()
	// set stdout to our write pipe
	os.Stdout = w

	intro()

	// close and reset
	_ = w.Close()
	os.Stdout = oldOut

	// read output
	out, _ := io.ReadAll(r)
	if !strings.Contains(string(out), "Is Prime?") {
		t.Errorf("intro test not correct: got %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a whole number!"},
		{name: "zero", input: "0", expected: "0 is not prime, by definition!"},
		{name: "one", input: "1", expected: "1 is not prime, by definition!"},
		{name: "two", input: "2", expected: "2 is a prime number!"},
		{name: "negative", input: "-2", expected: "Negative numbers are not prime, by definition!"},
		{name: "typed", input: "three", expected: "Please enter a whole number!"},
		{name: "decimal", input: "1.1", expected: "Please enter a whole number!"},
		{name: "quit", input: "q", expected: ""},
	}

	for _, e := range tests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumber(reader)

		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: expected %s, got %s", e.name, e.expected, res)
		}
	}
}
```

### Test user input

```go
package main

import (
	"bufio"
	"fmt"
	"io"
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
	go readUserInput(os.Stdin, doneChan)

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

func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)

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

```go
package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

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

func Test_prompt(t *testing.T) {
	oldOut := os.Stdout

	// create read and write pipe
	r, w, _ := os.Pipe()
	// set stdout to our write pipe
	os.Stdout = w

	prompt()

	// close and reset
	_ = w.Close()
	os.Stdout = oldOut

	// read output
	out, _ := io.ReadAll(r)
	if string(out) != "-> " {
		t.Errorf("incorrect prompt: exptected -> but got %s", string(out))
	}
}

func Test_info(t *testing.T) {
	oldOut := os.Stdout

	// create read and write pipe
	r, w, _ := os.Pipe()
	// set stdout to our write pipe
	os.Stdout = w

	intro()

	// close and reset
	_ = w.Close()
	os.Stdout = oldOut

	// read output
	out, _ := io.ReadAll(r)
	if !strings.Contains(string(out), "Is Prime?") {
		t.Errorf("intro test not correct: got %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a whole number!"},
		{name: "zero", input: "0", expected: "0 is not prime, by definition!"},
		{name: "one", input: "1", expected: "1 is not prime, by definition!"},
		{name: "two", input: "2", expected: "2 is a prime number!"},
		{name: "negative", input: "-2", expected: "Negative numbers are not prime, by definition!"},
		{name: "typed", input: "three", expected: "Please enter a whole number!"},
		{name: "decimal", input: "1.1", expected: "Please enter a whole number!"},
		{name: "quit", input: "q", expected: ""},
	}

	for _, e := range tests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumber(reader)

		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: expected %s, got %s", e.name, e.expected, res)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	// channel
	doneChan := make(chan bool)
	// io reader buffer
	var stdin bytes.Buffer
	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
```

## Test suites

### Single test

```sh
go test .
go test -v .
go test -run Test_isPrime
go test -v -run Test_isPrime
```

### Multiple tests

```sh
go test -run Test_alpha_isPrime
go test -v -run Test_alpha_isPrime
```

## Web applications

### Simple app

```sh
go mod init webapp
go get -u github.com/go-chi/chi/v5
go get -u github.com/go-chi/chi/v5/middleware
go mod tidy
go run ./cmd/web
```
