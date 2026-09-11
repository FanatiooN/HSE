package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	maxSum := (1 + N) * N / 2

	if maxSum < M {
		fmt.Println(0)
		return
	}

	pens := N
	for M > 0 {
		if pens <= M {
			M -= pens
			fmt.Println(pens)
		}
		pens = min(pens-1, M)
	}

	return
}
