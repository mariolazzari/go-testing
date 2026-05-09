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
