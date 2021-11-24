package main

import (
	"fmt"
)

//add- Method reads from in channel for count times and adds the values read from the channel. once done will push the
//result to out channel.
func add(in chan int, out chan int, count int) {
	sum := 0
	for i := 0; i < count; i++ {
		sum = sum + <-in
	}
	out <- sum
}

//Creates 10 channels, and all numbers from 1 to 100 will be pushed to corresponding channel based on last digit.
//creates 10 routines to run add() function for each channel.
//create routine to call add with previous results to find total sum
func main() {
	fmt.Println("Task-1 Channels and go Routines")

	noOfChannels := 10

	var sum []chan int
	for i := 0; i < noOfChannels; i++ {
		c := make(chan int)
		sum = append(sum, c)
	}

	total := make(chan int)
	out := make(chan int)

	for i := range sum {
		go add(sum[i], total, 10)
	}

	for i := 1; i <= 100; i++ {
		sum[i%10] <- i
	}

	go add(total, out, 10)
	fmt.Println("Total Sum:", <-out)
}
