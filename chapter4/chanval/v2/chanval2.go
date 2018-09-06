package main

import (
	"fmt"
	"time"
)

// Counter 代表计数器的类型。
type Counter struct {
	count int
}

var mapChan = make(chan map[string]Counter, 1)

func main() {
	syncChan := make(chan struct{}, 2)

	go func() { // 接收操作
		for {
			if elem, ok := <-mapChan; ok {
				counter := elem["count"]
				counter.count++
				fmt.Printf("[receiver] The elem is: %v. Elem address: %p \n", elem, &elem)
				fmt.Printf("[receiver] The counter: %d.\n", counter.count)
			} else {
				break
			}
		}
		fmt.Println("[receiver] Stopped.")
		syncChan <- struct{}{}
	}()

	go func() { // 发送操作
		countMap := map[string]Counter{
			"count": Counter{},
		}
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("[sender] The count map: %v. Count map address: %p \n", countMap, &countMap)
			fmt.Printf("[sender] The counter: %d.\n", countMap["count"].count)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()

	<-syncChan
	<-syncChan
}
