package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(time.Second)
	}
}

func mainOld() {
	go task("A") // go routines ou green thread do Go // outra thread 1
	go task("B") // outra thread 2
	task("C") // Thread main 
}

func main() { // Thread 1
	channel := make(chan string) // cria um canal para comunicação entre threads

	go func() { // Thread 2
		channel <- "thread 2 depositou esta string"
	}()

	// Thread 1
	mgs := <- channel
	fmt.Println(mgs)
	// ou
	// fmt.Println(<- channel) // se e ao invés de ou esse dá erro pois já foi lida e a thread 2 não está mais rodando
}
