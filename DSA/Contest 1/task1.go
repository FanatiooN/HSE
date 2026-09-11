package main

import "fmt"

func main() {
	var left, mid, right int
	fmt.Scan(&left, &mid, &right)

	k := min(left, mid/2, right)

	left -= k
	right -= k
	mid -= 2 * k

	ans := 4 * k

	if left > 0 {
		left--
		ans++
		if mid > 0 {
			mid--
			ans++
			if right > 0 {
				right--
				ans++
				if mid > 0 {
					mid--
					ans++
				}
			}
		}
	}

	fmt.Println(ans)
}
