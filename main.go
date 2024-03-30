package main

import (
	"fmt"
	"numfin/pkg/crr"
)

func main() {
	fmt.Println("Hello CallOption!")

	var call = crr.NewCallOption(100, 100.0)
	var bm = crr.NewBinModel(100.0, 0.1, -0.1, 0.05)

	fmt.Println(crr.Price_by_CRR(bm, call))
}
