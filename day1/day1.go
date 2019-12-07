package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Fuel required to launch a given module is based on its mass.
// Specifically, to find the fuel required for a module, take its mass, divide by three, round down, and subtract 2.

// For a mass of 12, divide by 3 and round down to get 4, then subtract 2 to get 2.
// For a mass of 14, dividing by 3 and rounding down still yields 4, so the fuel required is also 2.
// For a mass of 1969, the fuel required is 654.
// For a mass of 100756, the fuel required is 33583.

func main() {
	lineScanner := bufio.NewScanner(os.Stdin)
	total := 0
	for lineScanner.Scan(); len(lineScanner.Text()) != 0; lineScanner.Scan() {
		mass, _ := strconv.Atoi(lineScanner.Text())
		// if mass == 0 {
		fuel := mass/3 - 2

		// So, for each module mass, calculate its fuel and add it to the total.
		//Then, treat the fuel amount you just calculated as the input mass and repeat the process,
		// continuing until a fuel requirement is zero or negative. For example:
		for additional := fuel/3 - 2; additional >= 0; additional = additional/3 - 2 {
			fuel += additional
		}
		total += fuel
	}
	fmt.Println(total)
}

// func main() {
// 	mass, err := strconv.Atoi(os.Args[1:])
// 	if err != nil {
// 		panic(err)
// 	}
// 	println(mass/3 - 2)
// }
