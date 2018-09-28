package main

import (
	"fmt"
	"log"
)

// print_arr prints all the items in the given array
func print_arr(arr []star) {

	log.Printf("[+] Printing array")

	// Iterate over all the elements in the array and print their content
	for _, el := range arr{
		fmt.Printf("%f\t%f\t|%f\t%f\t|%f\n", el.c.x, el.c.y, el.f.x, el.f.y, el.mass)
	}

	fmt.Println("---")
}
