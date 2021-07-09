package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"megaproj1/utils"
	"time"
)

var (
	flg = flag.Int("period", 60, "time period between the ticks showing the json")
)

func main() {
	flag.Parse()
	for {
		intrv, _ := utils.MERGEI([]utils.Interval{{-20, 0}, {300, 455}, {1, 2}, {5, 10}, {4, 23}, {4, 54}, {55, 100}, {102, 200}})
		a, err := json.MarshalIndent(intrv, "", " ")
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}
		fmt.Println(string(a))
		time.Sleep(time.Duration(*flg) * time.Second)
	}

}
