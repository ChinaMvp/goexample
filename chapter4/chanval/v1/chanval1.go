package main

import (
	"fmt"
	"time"
)

var mapChan = make(chan map[string]int, 1)

func main() {
	syncChan := make(chan struct{}, 2)

	go func() { // 接收操作
		for {
			if elem, ok := <-mapChan; ok {
				elem["count"]++
				fmt.Printf("[receiver] The elem is: %v. Elem address is: %p \n", elem, &elem)
			} else {
				break
			}
		}
		fmt.Println("[receiver] Stopped.")
		syncChan <- struct{}{}
	}()

	go func() { // 发送操作
		countMap := make(map[string]int)
		for i := 0; i < 5; i++ {
			countMap["count"]++
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("[sender] The count map: %v. Count map address: %p \n", countMap, &countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()

	<-syncChan
	<-syncChan
}
