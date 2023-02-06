package main

import "time"

func main() {
	time.AfterFunc(1*time.Second, func() {
		print("world")
	})
	println("hello")
	time.Sleep(2 * time.Second)
}
