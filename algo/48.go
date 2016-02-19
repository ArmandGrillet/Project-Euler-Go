// Find the last ten digits of the series, 1**1 + 2**2 + 3**3 + ... + 1000**1000.

package algo

import (
	"fmt"
	"math/big"
)

func E48() {
	sum, power := big.NewInt(0), big.NewInt(0)
	bigI := big.NewInt(1)
	for i := 1; i <= 1000; i++ {
		bigI = big.NewInt(int64(i))
		power.Exp(bigI, bigI, nil)
		sum.Add(sum, power)
	}

	sum.Mod(sum, big.NewInt(10000000000))
	fmt.Println(sum)
}
