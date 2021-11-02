package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type Details struct {
	User_name     string
	Ageing        int
	Favour_number float64
}
var details Details

func main() {
	reader = bufio.NewReader(os.Stdin)
	details.User_name = name("what name")
	details.Ageing = age("what age")
	details.Favour_number = favour("Enter fav number")
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
