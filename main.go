package main

import "fmt"

var chan1=make(chan string)
var chan2=make(chan int)

func f1(){
	chan1<-"Chris"
}
func f2(){
	chan2<-63
}

func main(){
	go f1()
	go f2()
	fmt.Println(<-chan1)
	fmt.Println(<-chan2)
}