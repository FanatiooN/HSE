package main

import (
	"fmt"
)

func main() {
	var roomsCnt, a, b int
	fmt.Scan(&roomsCnt, &a, &b)

	ans := 0
	c := lcm(a, b)

	if c == 1 && a*b != 1 {
		fmt.Println(0)
		return
	}

	ans += roomsCnt / a
	ans += roomsCnt / b
	ans -= roomsCnt / c

	fmt.Println(ans)

}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
