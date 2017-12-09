package main

import (
	"errors"
	"strings"
)

//StringService provides uppercase and wordcount operations on strings
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

//SimpleStringService an idiotic implementation of StringService
type SimpleStringService struct{}

var _ StringService = &SimpleStringService{}

//Uppercase is pretty intuitive as a name
func (SimpleStringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil

}

//Count counts the words in a string
func (SimpleStringService) Count(s string) int {
	return len(s)
}

//ErrEmpty is a simple Empty String error message
var ErrEmpty = errors.New("Empty String")
