package main

import (
	"fmt"
	"os"
	)

func main() {

	if len(os.Args) != 2 {
		os.Exit(1)
	}

	fmt.Println("It's over", os.Args[1])

	for i := 0; i < len(os.Args); i++ {
		fmt.Println("i: ",i," , value: ", os.Args[i])
	} 
}