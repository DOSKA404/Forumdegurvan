package main

import (
	"fmt"

	"github.com/GurvanN22/Forum/src/Backend/tools"
)

func main() {
	hash := "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"
	userInput := "test"

	fmt.Println(tools.CheckPasswordValidity(hash, userInput))
}
