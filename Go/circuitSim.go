package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	//"time"
)



// Gate routines
func inverter(in chan bool, out chan bool){

	a := <-in
	out <- !a
}


func andGate(in1 chan bool, in2 chan bool, out chan bool){
	a := <-in1
	b := <-in2
	fmt.Printf("XorGate %v %v\n", a, b)
	out <- (a && b)

}


func orGate(in1 chan bool, in2 chan bool, out chan bool){
	a := <-in1
	b := <-in2

	out <- (a || b)

}

func nandGate(in1 chan bool, in2 chan bool, out chan bool){
	a := <-in1
	b := <-in2

	out <- !(a && b)

}

func norGate(in1 chan bool, in2 chan bool, out chan bool){
	a := <-in1
	b := <-in2

	out <- !(a || b)

}

func xorGate(in1 chan bool, in2 chan bool, out chan bool){
	a := <-in1
	b := <-in2
	fmt.Printf("XorGate %v %v\n", a, b)
	if (a == b){
		out<-false
	} else {
		out<-true
	}
}

// Stores input values
func input(in chan bool, val bool){
	in <-val
}

// Retrieves final output
func output(out chan bool, main chan bool, last bool){
	x := <-out
	main <-x

	if last{
		close(main)
	}
}

// fan out
func fan(target chan bool, n int){

	x := <-target
	for i := 0; i < n+1; i++ {
		target <-x
	}
}

// Simulates clock signal
/*func clkSim(freq int, cycles int, signals chan bool){

	// pulses per millisecond
	pulses := 1000/freq
	sig := false
	time := time.Now().Clock()

	for i := 0; i < cycles; i++ {
		// 1 ms has passed
		if (time.Now().Clock() > time) {
			for j := 0; j < pulses; j++{
				sig = !sig
				signals <-sig
			}
		}
	}

}
*/

func simulate(instr []string, lc int, result chan bool) {

	var channels []chan bool

	for i, val := range instr {
		channel := make(chan bool, 5)

		values := strings.Split(val, ",")

		// Inputs
		if values[0] == "0" {
			go input(channel, false)
		} else if values[0] == "1" {
			go input(channel, true)
		} else if values[0] == "fan"{

			input1, err1 := strconv.Atoi(values[1]);
			input2, err2 := strconv.Atoi(values[2]);

			if  err1 == nil && err2 == nil {
				go fan(channels[input1], input2)
			} else {
				fmt.Println("format error in file")
			}

		} else if values[0] == "AndGate"{

			input1, err1 := strconv.Atoi(values[1]);
			input2, err2 := strconv.Atoi(values[2]);

			if  err1 == nil && err2 == nil {
				go andGate(channels[input1], channels[input2], channel)
			} else {
				fmt.Println("format error in file")
			}

		} else if values[0] == "OrGate" {

			input1, err1 := strconv.Atoi(values[1]);
			input2, err2 := strconv.Atoi(values[2]);

			if  err1 == nil && err2 == nil {
				go orGate(channels[input1], channels[input2], channel)
			} else {
				fmt.Println("format error in file")
			}


		} else if values[0] == "NandGate" {

			input1, err1 := strconv.Atoi(values[1]);
			input2, err2 := strconv.Atoi(values[2]);

			if  err1 == nil && err2 == nil {
				go nandGate(channels[input1], channels[input2], channel)
			} else {
				fmt.Println("format error in file")
			}

		} else if values[0] == "NorGate" {

			input1, err1 := strconv.Atoi(values[1]);
			input2, err2 := strconv.Atoi(values[2]);

			if  err1 == nil && err2 == nil {
				go norGate(channels[input1], channels[input2], channel)
			} else {
				fmt.Println("format error in file")
			}

		} else if values[0] == "XorGate" {

			input1, err1 := strconv.Atoi(values[1]);
			input2, err2 := strconv.Atoi(values[2]);

			if  err1 == nil && err2 == nil {
				go xorGate(channels[input1], channels[input2], channel)
			} else {
				fmt.Println("format error in file")
			}

		} else if values[0] == "Output" {

			if index, err := strconv.Atoi(values[1]); err == nil{


				if i == lc {
					go output(channels[index], result, true)
				} else {
					go output(channels[index], result, true)
				}


			} else {
				fmt.Println("format error in file")
			}

		}

		channels = append(channels, channel)

	}
}

func read(fileName string) ([]string, int) {

	var lines []string
 	lineCount := 0

	if file, err := os.Open(fileName); err == nil {

		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
			lineCount++
		}

		return lines, lineCount

	}

	return lines, lineCount
}


func main(){

	var file string;

	if len(os.Args) > 1 {
		file = os.Args[1]
	} else {
		file = "test"
	}

	var lines, lc = read(file)
	if len(lines) == 0 {
		fmt.Println("Error: could not find file")
	}

	res := make(chan bool)

	simulate(lines, lc, res)
	for x := range res{
		fmt.Println(x)
	}

}
