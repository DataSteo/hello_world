package main

import (
	"fmt"
	"time"
)

func main() {
	chanRet := make(chan string)
	go func() {
		time.Sleep(time.Second * 5)
		chanRet <- "fin 5"
	}()
	go func() {
		time.Sleep(time.Second * 1)
		chanRet <- "fin 1"
	}()
	for {
		msg1 := <-chanRet
		fmt.Println(msg1)
		msg1 = <-chanRet
		fmt.Println(msg1)
		break
	}
}
