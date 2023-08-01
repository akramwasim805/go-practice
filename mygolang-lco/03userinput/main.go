package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("We are testing third chapter here named as user input")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the ratting of the Pizza: ")

	// commma ok || err ok

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)

}
