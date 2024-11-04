package main

import (
	"fmt"
	"project/utils" // Adjust the import path if needed based on your directory structure
)

func main() {
	n := 3
	fmt.Println("MagicSum:", utils.MagicSum(n))
	fmt.Println("MagicPow:", utils.MagicPow(n))
	fmt.Println("MagicOdd:", utils.MagicOdd(n))
	fmt.Println("MagicGrade:", utils.MagicGrade(n))
	fmt.Println("MagicName:", utils.MagicName(n))
	fmt.Println("MagicTria:", utils.MagicTria(n))

	number := 5
	utils.MagicChange(&number)
	fmt.Println("MagicChange:", number)
}

