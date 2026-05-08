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

```
