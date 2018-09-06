package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 1)

	ticker := time.NewTicker(time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println("[sender] Tick at:cls", t)
			select {
			case intChan <- 1:
				fmt.Println("[sender] send value:", 1)
			case intChan <- 2:
				fmt.Println("[sender] send value:", 2)
			case intChan <- 3:
				fmt.Println("[sender] send value:", 3)
			}
		}
		fmt.Println("[sender] End.") //永远不会执行
	}()

	var sum int
	for e := range intChan {
		fmt.Printf("[receiver] Received: %v\n", e)
		sum += e
		if sum > 10 {
			fmt.Printf("[receiver] Got: %v\n", sum)
			break
		}
	}
	//当sum>10退出时，intChan容量已满，不再接收数据。这时，向intChan发送数据的操作会阻塞，所以[sender] End永远不会被输出。
	fmt.Printf("chan: %v, chan num: %d\n", intChan, cap(intChan))

	fmt.Println("[receiver] Sleep 5 seconds.")
	time.Sleep(5*time.Second)

	fmt.Println("[receiver] End.")
}
