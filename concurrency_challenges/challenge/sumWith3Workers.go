package challenge

import "fmt"

func sum(ch chan int, arr []int) {
	val := 0

	for _, a := range arr {
		val += a
	}
	ch <- val
}

func ThreeSum() {
	ch := make(chan int)

	arr := []int{1, 2, 3, 4, 5, 6, 6, 7, 8, 9, 10, 11, 12}

	for i := range 3 {
		go sum(ch, arr[i*4:i*4+4])
	}
	res := <-ch + <-ch + <-ch
	fmt.Printf("Sum %d:",res)
}
