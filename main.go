package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if i%7 == 0 {
			fmt.Println(i)
		}
	}
}
