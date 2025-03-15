package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hi World!")

	os.Getenv("PORT")
}
