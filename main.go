package main

import (
	"fmt"

	"github.com/singmyr/sicli/selection"
)

func main() {
	options := []string{
		"Option 1",
		"Option 2",
		"Option 3",
	}
	value := selection.GetValue(options)
	fmt.Println("Selected:", value)
	key := selection.GetKey(options)
	fmt.Println("Selected:", key)
}
