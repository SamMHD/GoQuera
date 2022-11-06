package main

import (
	"sync"
)

func Solution(f func(uint8) uint8, inp uint32) uint32 {
	var wg sync.WaitGroup
	var res uint32

	wg.Add(4)
	for i := 0; i <= 3; i += 1 {
		go func(i int) {
			defer wg.Done()
			res = res | uint32(f(uint8((inp>>(i<<3)))))<<(i<<3)
		}(i)
	}

	wg.Wait()
	return res
}

// func main() {
// 	fmt.Println(Solution(func(x uint8) uint8 { return x }, 0x1111))
// }
