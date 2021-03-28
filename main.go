package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
	"math/rand"

	"github.com/AlecAivazis/survey"
)

// Define User Prompts
var getParams = []*survey.Question{
	{
		Name: "length",
		Prompt: &survey.Input{
			Message: "How long do you want your password to be?\n",
		},
		Validate: func(val interface{}) error {
			num, err := strconv.Atoi(val.(string));
			if err != nil || num<8 || num>128 {
				return fmt.Errorf("Must be a number between 8 & 128")
			}
			return nil
		},
		Transform: survey.Title,
	},
	{
		Name: "lowerCase",
		Prompt: &survey.Confirm{Message: "Include Lower Case?\n"},
		Validate: survey.Required,
	},
	{
		Name: "upperCase",
		Prompt: &survey.Confirm{Message: "Include Upper Case?\n"},
		Validate: survey.Required,
	},{
		Name: "numbers",
		Prompt: &survey.Confirm{Message: "Include Numbers?\n"},
		Validate: survey.Required,
	},{
		Name: "specials",
		Prompt: &survey.Confirm{Message: "Include Special Characters?\n"},
		Validate: survey.Required,
	},
}

func main() {
	lowers := "abcdefghijklmnopqrstuvwxyz"
	uppers := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	specials := "!@#$%^&*()_+-=<>,.?:;{}[]|"
	chars := ""
	password := ""
	
	// Store User Input from Prompts 
	answers := struct {
		Length  string `survey:"length"`
		LowerCase  bool `survey:"lowerCase"`
		UpperCase  bool `survey:"upperCase"`
		Numbers  bool `survey:"numbers"`
		Specials  bool `survey:"specials"`
	}{}

	survey.Ask(getParams, &answers)

	// Convert response string to an integer
	length, err := strconv.Atoi(answers.Length)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Build char string based on user input
	if answers.LowerCase { chars += lowers }
	if answers.UpperCase { chars += uppers }
	if answers.Numbers { chars += numbers }
	if answers.Specials { chars += specials }

	// TODO: handle error for 0 char string length

	// build password string to be returned
	for i := 0; i < length; i++ {
		password += string(chars[rand.Intn(utf8.RuneCountInString(chars))])
	}

	fmt.Printf("\nYour new password is:\n\n%s\n\n\n", password)
}
