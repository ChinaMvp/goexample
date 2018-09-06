package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex

	fmt.Println("Lock the lock. (main)")
	mutex.Lock()
	fmt.Println("The lock is locked. (main)")

	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Printf("Lock the lock. (g%d)\n", i)
			mutex.Lock()
			fmt.Printf("The lock is locked. (g%d)\n", i)
		}(i)
	}
	//主例程休眠1秒
	fmt.Println("Sleep one second. (main)")
	time.Sleep(time.Second)
	fmt.Println("Wake. (main)")

	fmt.Println("Unlock the lock. (main)")
	mutex.Unlock()
	fmt.Println("The lock is unlocked. (main)")

	time.Sleep(time.Second)
}
