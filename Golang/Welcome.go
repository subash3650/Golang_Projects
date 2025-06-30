package main

import (
	"errors"
	"fmt"
)

func Greeting(str string) (string, error) {
	if str == "" {
		return str, errors.New("empty name")
	}
	msg := fmt.Sprintf("Hi, %v. Welcome!!", str)
	return msg, nil
}
