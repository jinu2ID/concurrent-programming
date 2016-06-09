package main

import "fmt"

func inverter(in chan bool, out chan bool){

	a := <-in
	out <- !a
}


func andGate(in chan bool, out chan bool){
	a := <-in
	b := <-in

	out <- (a && b)

}


func orGate(in chan bool, out chan bool){
	a := <-in
	b := <-in

	out <- (a || b)

}

func nandGate(in chan bool, out chan bool){
	a := <-in
	b := <-in

	out <- !(a && b)

}

func norGate(in chan bool, out chan bool){
	a := <-in
	b := <-in

	out <- !(a || b)

}

func xorGate(in chan bool, out chan bool){
	a := <-in
	b := <-in

	if (a == b){
		out<-false
	} else {
		out<-true
	}
}

func main(){

	in := make(chan bool)
	out := make(chan bool)

	go orGate(in, out)
	in <- true
	in <- true

	x := <-out

	go inverter(in, out)
	in <- x
	x = <-out

	fmt.Printf("%t\n", x)


}
