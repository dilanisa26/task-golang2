package main

import (
	"fmt"
	"os"
	"strconv"
	"project/utils" // Adjust the import path to your directory structure
)

// Define the MagicNumber struct
type MagicNumber struct {
	Number int
}

// Method Multiply for MagicNumber struct
func (mn *MagicNumber) Multiply(n int) {
	mn.Number *= n
}

func main() {
	// Handle command-line arguments for n
	n := 3
	if len(os.Args) > 1 {
		arg, err := strconv.Atoi(os.Args[1])
		if err == nil {
			n = arg
		}
	}

	// Utilize utils package functions
	fmt.Println("MagicSum:", utils.MagicSum(n))
	fmt.Println("MagicPow:", utils.MagicPow(n))
	fmt.Println("MagicOdd:", utils.MagicOdd(n))
	fmt.Println("MagicGrade:", utils.MagicGrade(n))
	fmt.Println("MagicName:", utils.MagicName(n))
	fmt.Println("MagicTria:", utils.MagicTria(n))

	// Demonstrate MagicChange
	number := 5
	fmt.Println("Before MagicChange:", number)
	utils.MagicChange(&number)
	fmt.Println("After MagicChange:", number)

	// Initialize MagicNumber and use the Multiply method
	magicNumber := MagicNumber{Number: n}
	fmt.Println("MagicNumber before Multiply:", magicNumber.Number)
	magicNumber.Multiply(2) // Example: multiply by 2
	fmt.Println("MagicNumber after Multiply:", magicNumber.Number)
}
