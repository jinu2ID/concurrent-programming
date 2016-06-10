package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

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


func input(in chan bool, values []string){
	for _, val := range values{
		if val == "0"{
			in <-false
		} else {
			in <-true
		}
	}
}

func output(out chan bool){
	fmt.Println("HERE")
	x := <-out
	fmt.Println(x)
}

func simulate(instr []string) {

	var channels []chan bool

	for _, val := range instr {
		channel := make(chan bool)

		values := strings.Split(val, ",")

		// Inputs
		if values[0] == "0" || values[0] == "1" {
			go input(channel, values)
		} else if values[0] == "AndGate"{
			if input, err := strconv.Atoi(values[1]); err == nil{
				go andGate(channels[input], channel)
			} else {
				fmt.Println("format error in file")
			}
		} else if values[0] == "OrGate" {
			if input, err := strconv.Atoi(values[1]); err == nil{
				go orGate(channels[input], channel)
			} else {
				fmt.Println("format error in file")
			}
		} else if values[0] == "Output" {
		
			if index, err := strconv.Atoi(values[1]); err == nil{
				fmt.Println(index)
				go output(channels[index])

			} else {
				fmt.Println("format error in file")
			}

		}

		channels = append(channels, channel)

	}
	fmt.Println(len(channels))
}

func read(fileName string) []string {

	var lines []string

	if file, err := os.Open(fileName); err == nil {

		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		return lines

	}

	return lines
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

	//fmt.Printf("%t\n", x)

	var lines = read("test")
	if len(lines) == 0 {
		fmt.Println("Error: could not find file")
	}

	simulate(lines)
/*
	for _, val := range lines {
		fmt.Println(val)
	}
*/

}
