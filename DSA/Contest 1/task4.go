package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := n + 1

	if n >= 1 {
		ans += ((1 + n*2) + (1 + 2)) * n / 2
	}

	fmt.Println(ans)
}
