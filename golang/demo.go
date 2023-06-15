package main

import (
	"fmt"
)

func main() {
	list := []int{1, 2, 3, 4}
	target := 6

	for k1, v1 := range list {
		fmt.Printf("index:%v,value:%v", k1, v1)
		for k2, v2 := range list {
			res := v1 + v2
			if res == target {
				fmt.Printf("output: %v,%v", k1, k2)
			}
		}
	}
}
