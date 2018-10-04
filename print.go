// Collection of functions handling printing specific stuff

package main

import (
	"fmt"
	"log"
)

// print_arr prints all the items in the given array
func print_arr(arr []star) {

	log.Printf("[+] Printing array")

	fmt.Println("x\t\ty\t\t | Fx\t\tFy\t\t | mass\n")


	for _, el := range arr{
		fmt.Printf("%.6f\t%.6f\t | %.6f\t%.6f\t | %.6f\n", el.c.x, el.c.y, el.f.x, el.f.y, el.mass)
	}

	fmt.Println("")
}

