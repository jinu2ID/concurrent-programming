package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sync"
	//"time"
)

var finalOutput []int

// Gate routines
func inverter(in chan bool, out chan bool, wg *sync.WaitGroup){

	a := <-in
	out <- !a
	wg.Done()
}


func andGate(in1 chan bool, in2 chan bool, out chan bool, wg *sync.WaitGroup){
	a := <-in1
	b := <-in2
	in1 <-a
	in2 <-b

	out <- (a && b)
	wg.Done()
	fmt.Println("AND")

}


func orGate(in1 chan bool, in2 chan bool, out chan bool, wg *sync.WaitGroup){
	a := <-in1
	b := <-in2

	in1 <-a
	in2 <-b

	out <- (a || b)
	wg.Done()
	fmt.Println("OR")

}

func nandGate(in1 chan bool, in2 chan bool, out chan bool, wg *sync.WaitGroup){
	fmt.Println("NAND")
	a := <-in1
	b := <-in2

	in1 <-a
	in2 <-b

	out <- !(a && b)
	wg.Done()
}

func norGate(in1 chan bool, in2 chan bool, out chan bool, wg *sync.WaitGroup){
	fmt.Println("NOR")
	a := <-in1
	b := <-in2

	in1 <-a
	in2 <-b

	out <- !(a || b)
	wg.Done()
}

func xorGate(in1 chan bool, in2 chan bool, out chan bool, wg *sync.WaitGroup){
	a := <-in1
	b := <-in2
	in1 <-a
	in2 <-b
	if (a == b){
		out<-false
	} else {
		out<-true
	}
	wg.Done()
	fmt.Println("XOR")

}

// Stores input values
func input(in chan bool, val bool, wg *sync.WaitGroup){
	in <-val
	wg.Done()
	fmt.Println("INPUT")

}

// Stores carry values
func carry(in chan bool, out chan bool, wg *sync.WaitGroup){
	x := <-in
	out <-x
	wg.Done()
	fmt.Println("CARRY")

}

// Retrieves final output
func output(out chan bool, index int, wg *sync.WaitGroup){
	x := <-out
	if x {
//		fmt.Print(1)
		finalOutput[index] = 1
	} else {
//		fmt.Print(0)
		finalOutput[index] = 0
	}
	wg.Done()
	fmt.Println("OUTPUT")

}

// fan out
func fan(target chan bool, n int, wg *sync.WaitGroup){
//	fmt.Println("FAN")
/*	x := <-target
	for i := 0; i < n+1; i++ {
		target <-x
	}
*/
	wg.Done()

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

func simulate(instr []string, lc int) {

	var wg sync.WaitGroup
	//wg.Add(lc)
	var channels []chan bool
	outIndex := lc/14

	for _, val := range instr {
		channel := make(chan bool, 10)

		values := strings.Split(val, ",")

		// Inputs
		if values[0] == "0" {
			wg.Add(1)
			go input(channel, false, &wg)
		} else if values[0] == "1" {
			wg.Add(1)
			go input(channel, true, &wg)
		} else if values[0] == "fan"{

			input1, err1 := strconv.Atoi(values[1]);
			input2, err2 := strconv.Atoi(values[2]);

			if  err1 == nil && err2 == nil {
				wg.Add(1)
				go fan(channels[input1], input2, &wg)
			} else {
				fmt.Println("format error in file")
			}

		} else if values[0] == "AndGate"{

			input1, err1 := strconv.Atoi(values[1]);
			input2, err2 := strconv.Atoi(values[2]);

			if  err1 == nil && err2 == nil {
				wg.Add(1)
				go andGate(channels[input1], channels[input2], channel, &wg)

			} else {
				fmt.Println("format error in file")
			}

		} else if values[0] == "OrGate" {

			input1, err1 := strconv.Atoi(values[1]);
			input2, err2 := strconv.Atoi(values[2]);

			if  err1 == nil && err2 == nil {
				wg.Add(1)
				go orGate(channels[input1], channels[input2], channel, &wg)
			} else {
				fmt.Println("format error in file")
			}


		} else if values[0] == "NandGate" {

			input1, err1 := strconv.Atoi(values[1]);
			input2, err2 := strconv.Atoi(values[2]);

			if  err1 == nil && err2 == nil {
				wg.Add(1)
				go nandGate(channels[input1], channels[input2], channel, &wg)
			} else {
				fmt.Println("format error in file")
			}

		} else if values[0] == "NorGate" {

			input1, err1 := strconv.Atoi(values[1])
			input2, err2 := strconv.Atoi(values[2])

			if  err1 == nil && err2 == nil {
				wg.Add(1)
				go norGate(channels[input1], channels[input2], channel, &wg)
			} else {
				fmt.Println("format error in file")
			}

		} else if values[0] == "XorGate" {

			input1, err1 := strconv.Atoi(values[1]);
			input2, err2 := strconv.Atoi(values[2]);

			if  err1 == nil && err2 == nil {
				wg.Add(1)
				go xorGate(channels[input1], channels[input2], channel, &wg)
			} else {
				fmt.Println("format error in file")
			}

		} else if (values[0] == "CarryOut") || (values[0] == "CarryIn") {
			if index, err := strconv.Atoi(values[1]); err == nil{
				wg.Add(1)
				go carry(channels[index], channel, &wg)
			} else {
				fmt.Println("format error in file")
			}

		} else if values[0] == "Output" {
			if index, err := strconv.Atoi(values[1]); err == nil{

				wg.Add(1)
				go output(channels[index], outIndex, &wg)
				outIndex--

			} else {
				fmt.Println("format error in file")
			}

		}

		channels = append(channels, channel)

	}
	wg.Wait()
	fmt.Println(finalOutput)
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

	finalOutput = make([]int, lc/14+1)	// Magic number needs fix; 14 = num of lines for a full adder

	simulate(lines, lc)
}
