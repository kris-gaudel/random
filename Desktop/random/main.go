package main

import (
	"fmt"
	"random/random"
)

func main() {
	rlist := random.MiddleSquare(678902)
	rlist2 := random.LinearCongruentialGenerator(194835,512839,853932,543214)
	rlist3 := random.Xorshift(912472)

	for i := 0; i <= 19; i++ {
		fmt.Printf("Random Result %d:\n", i+1)
		fmt.Printf("	Middle Square: %d\n", rlist[i])
		fmt.Printf("	LCG: 		   %d\n", rlist2[i])
		fmt.Printf("	Xorshifter:    %d\n", rlist3[i])
	}
}