package main

import (
	"encoding/json"
	"fmt"
	"log"
	"megaproj1/utils"
)

func main() {
	intrv, _ := utils.MERGEI([]utils.Interval{{-20, 0}, {300, 455}, {1, 2}, {5, 10}, {4, 23}, {4, 54}, {55, 100}, {102, 200}})
	a, err := json.MarshalIndent(intrv, "", " ")
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	fmt.Println(string(a))
}
