package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 1)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			intChan <- i
		}
		close(intChan)
	}()

	timeout := time.Millisecond * 500
	var timer *time.Timer
	for {
		if timer == nil {
			timer = time.NewTimer(timeout)
		} else {
			timer.Reset(timeout) //重置定时器
		}
		select {
		case e, ok := <-intChan:
			if !ok {
				fmt.Println("End.")
				return
			}
			fmt.Printf("Received: %v\n", e)
		case e := <-timer.C:
			fmt.Printf("Timeout! Now is: %v, format time is: %s!\n", e, time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}
