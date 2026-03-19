package main

import "fmt"

func sumPart(arr []int, ch chan int) {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	ch <- sum

}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	numGorutines := 3
	ch := make(chan int, numGorutines)

	partSize := len(arr) / numGorutines

	for i := 0; i < numGorutines; i++ {
		start := 1 * partSize
		end := start + partSize
		go sumPart(arr[start:end], ch)
	}

	totalSum := 0

	for i := 0; i < numGorutines; i++ {
		totalSum += <-ch
	}

	fmt.Println("Total sum: ", totalSum)
}

// type Reader interface {
// Read(p []byte) (n int, err error)

// }

// type Reader2 interface {
// Read(p []byte) (n int, err error)

// }
