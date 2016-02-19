package algo

import (
	"fmt"
	"time"
)

func E1() {
	max := 1000

	sum := 0

	start := time.Now()
	for i := 0; i < max; i += 3 {
		sum += i
	}
	for i := 0; i < max; i += 5 {
		if i%3 != 0 {
			sum += i
		}
	}
	elapsed := time.Since(start)

	fmt.Println(sum)
	fmt.Println(elapsed)
}
