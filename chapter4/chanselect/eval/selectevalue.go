package main

import "fmt"

var intChanA chan int = make(chan int, 1)
var intChanB chan int = make(chan int, 1)
var channelAll = []chan int{intChanA, intChanB}

var numberAll = []int{1, 2, 3, 4, 5}

func main() {
	select {
	case getChannel(0) <- getNumberValue(0):
		fmt.Println("The 1th case is selected.")
	case getChannel(1) <- getNumberValue(1):
		fmt.Println("The 2nd case is selected.")
	default:
		fmt.Println("Default!")
	}
}

func getNumberValue(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numberAll[i]
}

func getChannel(i int) chan int {
	fmt.Printf("channels[%d]\n", i)
	return channelAll[i]
}
