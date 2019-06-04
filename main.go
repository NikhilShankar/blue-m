package main

import(
	"os/signal"
	"fmt"
	"os"
)

func main() {

	blue := New(100,3)
	blue.Add([]byte("sachu"))
	blue.Add([]byte("snikhil"))
	blue.Add([]byte("sachushan"))
	fmt.Println("Contains ", blue.Check([]byte("sachu")))
	fmt.Println("Contains ", blue.Check([]byte("nikhil")))
	fmt.Println("Contains ", blue.Check([]byte("sachushan")))


	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

}