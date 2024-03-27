package main

import (
	"fmt"
	"numfin/pkg/crr"
)

func main() {
	fmt.Println("Hello CallOption!")

	var call = crr.NewCallOption(100, 100.0)

	fmt.Println(call, call.Payoff(101.9))
}
