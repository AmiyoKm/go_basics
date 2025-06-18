package channel

import (
	"fmt"
)

func sum(s []int, c chan int) {
	summ := 0

	for _, val := range s {
		summ += val
	}
	c <- summ

}

func ChannelExec() {
	arr := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum(arr[:len(arr)/2], c)
	go sum(arr[len(arr)/2:], c)
	x := <-c
	y := <-c

	fmt.Println(x, y, x+y)

}
