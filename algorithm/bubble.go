package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	s1 := make([]int, 20)
	rand.Seed(time.Now().Unix())

	for k := range s1 {
		s1[k] = rand.Intn(1000)
	}

	fmt.Println(s1)
	Bubble(s1)
	fmt.Println(s1)
}

func Bubble(s []int) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}
