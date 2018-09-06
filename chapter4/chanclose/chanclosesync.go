package main

import (
	"fmt"
	"time"
)

func main() {
	dataChan := make(chan int, 5)
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go func() { // 接收操作。
		<-syncChan1
		fmt.Println("Received a sync signal and wait 3 seconds... [receiver]")
		time.Sleep(3*time.Second)
		for {
			if elem, ok := <-dataChan; ok {
				fmt.Printf("Received: %d [receiver]\n", elem)
			} else {
				break
			}
		}
		fmt.Println("Done. [receiver]")
		syncChan2 <- struct{}{}
	}()

	go func() { // 发送操作。
		for i := 0; i < 9; i++ {
			dataChan <- i
			if i == 4 {
				syncChan1 <- struct{}{}
			}
			fmt.Printf("Sent: %d [sender]\n", i)
		}

		close(dataChan) //关闭传输通道
		fmt.Println("Done. [sender]")
		syncChan2 <- struct{}{}
	}()

	<-syncChan2
	<-syncChan2
}
