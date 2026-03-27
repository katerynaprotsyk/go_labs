/*
	package main

import (

	"fmt"
	"sync"

)

	func main() {
		chEven := make(chan int)
		chOdd := make(chan int)
		var mu sync.Mutex
		counter := 0
		var wg sync.WaitGroup

		go func() {
			for i := 1; i <= 1000; i++ {
				if i%2 == 0 {
					chEven <- i
				} else {
					chOdd <- i
				}
			}
			close(chEven)
			close(chOdd)
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case val, ok := <-chEven:
					if !ok {
						chEven = nil
						break
					}
					if val%3 == 0 {
						mu.Lock()
						counter++
						mu.Unlock()
					}
				case val, ok := <-chOdd:
					if !ok {
						chOdd = nil
						break
					}
					if val%33 == 0 {
						mu.Lock()
						counter--
						mu.Unlock()
					}
				}
				if chEven == nil && chOdd == nil {
					break
				}
			}
		}()

		wg.Wait()
		fmt.Println("Counter:", counter)
	}
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	chEven := make(chan int)
	chOdd := make(chan int)
	var counter int64
	var wg sync.WaitGroup

	go func() {
		for i := 1; i <= 1000; i++ {
			if i%2 == 0 {
				chEven <- i
			} else {
				chOdd <- i
			}
		}
		close(chEven)
		close(chOdd)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case val, ok := <-chEven:
				if !ok {
					chEven = nil
					break
				}
				if val%3 == 0 {
					atomic.AddInt64(&counter, 1)
				}
			case val, ok := <-chOdd:
				if !ok {
					chOdd = nil
					break
				}
				if val%33 == 0 {
					atomic.AddInt64(&counter, -1)
				}
			}
			if chEven == nil && chOdd == nil {
				break
			}
		}
	}()

	wg.Wait()
	fmt.Println("Counter:", atomic.LoadInt64(&counter))
}
