package main

import "fmt"

/**
2651. 计算列车到站时间
*/

func findDelayedArrivalTime(arrivalTime int, delayedTime int) int {

	return (arrivalTime + delayedTime) % 24
}

func main() {

	time := findDelayedArrivalTime(10, 12)
	fmt.Println(time)
}
