// Даны n каналов типа chan int. Надо написать функцию, которая смерджит все данные из этих каналов в один и вернет его.

package main

import "sync"

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	wg := sync.WaitGroup{}
	go func() {

		for num := range joinChannels(ch1, ch2, ch3) {
			println(num)
		}
	}()

	func() {
		wg.Add(1)
		ch1 <- 1
		close(ch1)
		wg.Done()
	}()

	func() {
		wg.Add(1)
		ch2 <- 1
		ch2 <- 2
		ch2 <- 3
		close(ch2)
		wg.Done()
	}()

	func() {
		wg.Add(1)
		ch3 <- 22
		ch3 <- 33
		ch3 <- 44
		close(ch3)
		wg.Done()
	}()

	wg.Wait()
}

func joinChannels(channels ...chan int) chan int {
	out := make(chan int)
	for _, ch := range channels {
		go func(ch chan int) {
			for {
				v, ok := <-ch
				if !ok {
					break
				}
				out <- v
			}
		}(ch)
	}
	return out
}
