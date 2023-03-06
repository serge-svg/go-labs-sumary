package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt"}

	fmt.Printf("Print the current date %s problems.\n", time.Now())
	fmt.Printf("Print a random integer %d.\n", rand.Intn(10))
	fmt.Printf("Print the square of 4: %g.\n", math.Sqrt(4))
	fmt.Printf("Print the size of the array of strings %d.\n", len(answers))
	fmt.Printf("The sum of (1 + 1) is: %d \n", add1(1, 1))
	fmt.Printf("The sum of (1 + 1 + 1) is: %d \n", add2(1, 1, 1))
	fmt.Println("Swap the order of these 3 words 'you' 'are' 'How'.")
	fmt.Println(swap3strings("you", "are", "How"))
	fmt.Print("Split the number 17: ")
	fmt.Println(splitNumber(17))

	greeting, _ := fmt.Println("Hi")
	fmt.Println(greeting)
}

func add1(number1 int, number2 int) int {
	return number1 + number2
}

func add2(number1, number2, number3 int) int {
	return number1 + number2 + number3
}

func swap3strings(string1, string2, string3 string) (string, string, string) {
	return string3, string2, string1
}

func splitNumber(number int) (part1, part2 int) {
	part1 = number * 4 / 9
	part2 = number - part1
	return
}
