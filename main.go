package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, b *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := b.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func main() {

	fmt.Println("Hello")
}
