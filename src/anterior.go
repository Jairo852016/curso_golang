package main

import (
	"fmt"
)

/*func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}*/
/*func say(text string, c chan<- string) {

}*/
func message(text string, c chan string) {
	c <- text
}

func error() {
	/*var wg sync.WaitGroup
	fmt.Println("Hello")
	wg.Add(1)
	go say("world", &wg)
	wg.Wait()
	go func(text string) {
		fmt.Println(text)
	}("Adios")

	time.Sleep(time.Second * 1)*/

	/*c := make(chan string, 1)
	fmt.Println("Hello")

	go say("bye", c)
	fmt.Println(<-c)?*/
	/*	c := make(chan string, 2)
		c <- "Volumen1"
		c <- "Volumen2"
		fmt.Println(len(c), cap(c))
		//Range y close
		close(c)
		//c <- "Volumen3"
		for message := range c {
			fmt.Println(message)
		}*/
	//select
	email := make(chan string)
	email2 := make(chan string)
	go message("mesaje 1", email)
	go message("mensaje 2", email2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email:
			fmt.Println("Email recibido de email ", m1)
		case m2 := <-email2:
			fmt.Println("Email reibido de emil2", m2)
		}
	}
}
