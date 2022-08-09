package main

import (
	"fmt"

	"github.com/amiteshrai/golang-apps/src/greetings"
)

func main() {
	message := greetings.Greet("Amitesh")
	fmt.Println(message)
}
