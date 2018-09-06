package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Printf("Start. Now is: %s.\n", time.Now().Format("2006-01-02 15:04:05"))

	var rwm sync.RWMutex
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Printf("Try to lock for reading... [g%d]\n", i)
			rwm.RLock()
			fmt.Printf("Locked for reading. [g%d]\n", i)

			fmt.Printf("Sleep 2 seconds. [g%d]\n", i)
			time.Sleep(time.Second * 2)
			fmt.Printf("Wake. [g%d]\n", i)

			fmt.Printf("Try to unlock for reading... [g%d]\n", i)
			rwm.RUnlock()
			fmt.Printf("Unlocked for reading. [g%d]\n", i)
		}(i)
	}

	time.Sleep(time.Millisecond * 100)
	fmt.Println("Try to lock for writing...[main]")
	rwm.Lock()
	fmt.Println("Locked for writing.[main]")

	fmt.Printf("End. Now is: %s.\n", time.Now().Format("2006-01-02 15:04:05"))
}
