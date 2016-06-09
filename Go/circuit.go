package main

import "fmt"

func inverter(in chan bool, out chan bool){

	a := <-in
	a = !a
	out <- a
}


func andGate(in chan bool, out chan bool){
	a := <-in
	b := <-in

	if a && b {
		out<-true

	} else {
		out<-false
	}

}


func orGate(in chan bool, out chan bool){
	a := <-in
	b := <-in

	if a || b {
		out<-false

	} else {
		out<-true
	}

}

func nandGate(in chan bool, out chan bool){
	a := <-in
	b := <-in

	if a && b {
		out<-false
	} else {
		out<-true
	}
}

func norGate(in chan bool, out chan bool){
	a := <-in
	b := <-in

	if !(a && b) {
		out<-true
	} else {
		out<-false
	}

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

	go xorGate(in, out)
	in <- false
	in <- true

	x := <-out

	//go inverter(in, out)
	//in <- x
	//x = <-out

	fmt.Printf("%t\n", x)


}
