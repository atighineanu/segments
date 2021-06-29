package main

import (
	"fmt"
	"megaproj1/utils"
)

func main() {

	intervals := []utils.Interval{{25, 30}, {2, 19}, {14, 23}, {4, 8}}
	mergedInterval, _ := utils.MERGEI(intervals)
	fmt.Printf("The result: %v\n", mergedInterval)
}
