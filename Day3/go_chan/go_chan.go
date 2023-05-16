package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("goroutine")
	fmt.Println("main")

	for i := 0; i < 3; i++ {
		i := i
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(10 * time.Millisecond)

	shadowExample()

	ch := make(chan string)
	go func() {
		ch <- "hi" //send
	}()
	msg := <-ch //receive
	fmt.Println(msg)

	go func() {
		for i := 0; i < 3; i++ {
			msg := fmt.Sprintf("message #%d", i+1)
			ch <- msg
		}
		close(ch)
	}()
	// hits deadlock on a range, it doesn't know when to end - it does with a slice or a map - you need to close the channel close(ch)
	for msg := range ch {
		fmt.Println("got:", msg)
	}

	msg = <-ch
	fmt.Printf("closed: %#v\n", msg)

	msg, ok := <-ch
	fmt.Printf("closed: %#v (ok=%v)\n", msg, ok)

	// ch <- "hi" // ch is closed -> panic

	values := []int{15, 8, 42, 16, 4, 23}
	fmt.Println(sleepSort(values))

}

func sleepSort(values []int) []int {
	ch := make(chan int)
	for _, n := range values {
		n := n
		go func() {
			time.Sleep(time.Duration(n) * time.Millisecond)
			ch <- n
		}()
	}

	var out []int
	//run for the number of values - for i := 0; i < len(values); i++ {}
	for range values {
		n := <-ch
		out = append(out, n)
	}
	return out
}

func shadowExample() {
	n := 7
	{
		n := 2
		fmt.Println("inner", n)
	}
	fmt.Println("outer", n)
}
