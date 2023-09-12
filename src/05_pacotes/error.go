package main

import (
	"fmt"
	"time"
)

// MyError is an error implementation that includes a time and message.
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func oops() error {
	return MyError{
		time.Now(),
		//time.Date(2023, 9, 11, 23, 50, 0, 0, time.UTC),
		"o arquivo do sistema desapareceu",
	}
}

func main() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}
}