package number

import (
	"fmt"
	"strconv"
)

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

func ToHex(n uint64) string {
	return fmt.Sprintf("0x%016x", n)
}

func ToInt(n string) string {
	x, err := strconv.ParseInt(n, 16, 64)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("%d", x)
}
