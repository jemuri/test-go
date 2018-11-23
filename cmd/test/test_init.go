package main

import (
	"sync"
	"fmt"
)

func main() {
	WaitGroupTest()
}


func WaitGroupTest() {
	var wg sync.WaitGroup
	wg.Add(2)
	n1, n2 := 1, 1
	go func() {
		n1 = 5
		wg.Done()
	}()
	go func() {
		n2 = 6
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(n1 + n2)
}