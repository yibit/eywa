package number

import "fmt"

func OnesCount(n uint64) uint8 {
	var count uint8 = 0
	for n > 0 {
		count++
		n &= (n - 1)
	}

	return count
}

func Bits(n uint64) string {
	return fmt.Sprintf("%b", n)
}
