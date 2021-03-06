package main

import (
	"fmt"
	"time"
)

var strChan = make(chan string, 3)

func main() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go receive(strChan, syncChan1, syncChan2) // 接收操作
	go send(strChan, syncChan1, syncChan2)    //发送操作

	<-syncChan2
	<-syncChan2
}

//接收操作
func receive(strChan <-chan string,
	syncChan1 <-chan struct{},
	syncChan2 chan<- struct{}) {
	<-syncChan1
	fmt.Println("Received a sync signal and wait 3 seconds... [receiver]")
	time.Sleep(3*time.Second)

	fmt.Println("Wake. [receiver]")
	for elem := range strChan {
		fmt.Println("Received:", elem, "[receiver]")
	}
	fmt.Println("Stopped. [receiver]")

	syncChan2 <- struct{}{}
}

//发送操作
func send(strChan chan<- string,
	syncChan1 chan<- struct{},
	syncChan2 chan<- struct{}) {
	for _, elem := range []string{"a", "b", "c", "d", "e", "f", "g"} {
		strChan <- elem
		fmt.Println("Sent:", elem, "[sender]")
		if elem == "c" {
			syncChan1 <- struct{}{}
			fmt.Println("Sent a sync signal. [sender]")
		}
	}
	fmt.Println("Wait 2 seconds... [sender]")
	time.Sleep(time.Second * 2)

	close(strChan)

	syncChan2 <- struct{}{}
}
