package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/*
Part 1:

-----------------------------------------------------------------------------------------------------------------------------
Part 2:

*/
func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Print(content)
}
