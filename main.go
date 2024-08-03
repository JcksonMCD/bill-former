package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Chooes option (a - add item, s - save bill, t - add tip)", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("Added item :", name, price)
		promptOptions(b)
	case "s":
		tip, _ := getInput("Tip amount (£): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Println("Tip added: ", tip)
		promptOptions(b)
	case "t":
		fmt.Print("You chose a")
	default:
		fmt.Print("You did not choose a valid option...")
		promptOptions(b)
	}

}

func main() {

	fmt.Println("Hello")
}
