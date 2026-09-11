package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	stepsCnt := 0
	var answer string

	for {
		option := (stepsCnt + k) % 3

		if option == 1 {
			stepsCnt += k
			answer = "Yes"
			break
		}

		if option == 2 {
			stepsCnt += k
			answer = "No"
			break
		}

		r := stepsCnt % 3
		stepsCnt += n

		n = (r + n) / 3
		k = (r + k) / 3
	}

	fmt.Printf("%v\n%d", answer, stepsCnt)
}
