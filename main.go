package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput() (string, error) {

}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Create a new bill name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func main() {

	fmt.Println("Hello")
}
