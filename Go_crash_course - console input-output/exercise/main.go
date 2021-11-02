package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type Details struct {
	User_name     string
	Ageing        int
	Favour_number float64
	dogs          bool
}

var details Details

func main() {
	reader = bufio.NewReader(os.Stdin)
	details.User_name = name("what name")
	details.Ageing = age("what age")
	details.Favour_number = favour("Enter fav number")
	Details.dogs = readBool("Do you own a dog (y/n)?")
	fmt.Println(fmt.Sprintf("your name is %s and your age is %d and your favourite number is %.2f", details.User_name, details.Ageing, details.Favour_number))

}
func prompt() {
	fmt.Println("->")
}
func name(s string) string {
	for {
		fmt.Println(s)
		prompt()
		user_input, _ := reader.ReadString('\n')
		user_input = strings.Replace(user_input, "\r\n", "", -1)
		user_input = strings.Replace(user_input, "\n", "", -1)
		return user_input

	}
}
func age(s string) int {
	for {
		fmt.Println(s)
		prompt()
		user_input, _ := reader.ReadString('\n')
		user_input = strings.Replace(user_input, "\r\n", "", -1)
		user_input = strings.Replace(user_input, "\n", "", -1)

		num, err := strconv.Atoi(user_input)
		if err != nil {
			fmt.Println("Enter the whole number")
		} else {
			return num
		}

	}
}

func favour(s string) float64 {
	for {
		fmt.Println(s)
		prompt()
		user_input, _ := reader.ReadString('\n')
		user_input = strings.Replace(user_input, "\r\n", "", -1)
		user_input = strings.Replace(user_input, "\n", "", -1)

		num, err := strconv.ParseFloat(user_input, 64)
		if err != nil {
			fmt.Println("Enter the number")
		} else {
			return num
		}

	}
}
func readBool(s string) bool {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(s)
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if strings.ToLower(string(char)) != "y" && strings.ToLower(string(char)) != "n" {
			fmt.Println("Please type y or n")
		} else if char == 'n' || char == 'N' {
			return false
		} else if char == 'y' || char == 'Y' {
			return true
		}
	}
}
