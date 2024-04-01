package main

import "fmt"

func main() {

	a, b := make(chan int), make(chan int)
	go sender(3, a, b)

	sum := Sum(a, b)

	fmt.Println(<-sum) // block until the value is received

}

func sender(num int, a, b chan<- int) {
	for i := 1; i <= num; i++ {
		a <- i * i
		b <- i * i * i
	}
	close(a)
	close(b)

}

func Sum(a, b <-chan int) chan int {
	sum := make(chan int)
	go func() {
		s := 0
		a1Close, b1Close := false, false

		// for {
		// 	select {
		// 	case v1, ok1 := <-a:
		// 		if ok1 {
		// 			s += v1
		// 			//fmt.Println(s)
		// 		} else {
		// 			a1Close = true
		// 		}
		// 	case v2, ok2 := <-b:
		// 		if ok2 {
		// 			s += v2
		// 		} else {
		// 			b1Close = true
		// 		}
		// 	default:
		// 	}
		// 	if a1Close && b1Close {
		// 		//fmt.Println("reachjed?")
		// 		sum <- s
		// 		close(sum)
		// 		break
		// 	}
		// }

		go func() {
			for v := range a {
				s += v
			}
			a1Close = true
		}()
		go func() {
			for v := range b {
				s += v
			}
			b1Close = true
		}()
		go func() {
			for {
				if a1Close && b1Close {
					sum <- s
					close(sum)
					break
				}
			}
		}()
	}()

	return sum
}
