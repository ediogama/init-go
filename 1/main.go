package main

import (
	"errors"
	"fmt"
)

var (
	b bool = true
	c int  = 10
)

type ID int

type Address struct {
	street string
	number int
	city   string
	state  string
}

type Client struct {
	Name     string
	Age      int
	IsActive bool
	Address
}

func (c Client) DisableClient() {
	c.IsActive = false
	fmt.Printf("The client %s has been disable\n", c.Name)
}

func main() {

	firstClient := Client{
		Name:     "Edio",
		Age:      26,
		IsActive: true,
	}

	firstClient.street = "st california"

	fmt.Printf("The name of this client is %s. He has %d years old and he is %v. %s  is the street of him\n", firstClient.Name, firstClient.Age, firstClient.IsActive, firstClient.street)

	firstClient.DisableClient()

	value, err := sum(50, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)

	fmt.Printf("the value of c é %v", c)
	fmt.Printf("the type of c é %T", c)
	println("Hello, World\n")

	fmt.Println(sum_var(1, 2, 3))

}

func sum(a, b int) (int, error) {
	if a+b > 50 {
		return 0, errors.New("the sum is bigger than 50")
	}

	return a + b, nil
}

func sum_var(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
