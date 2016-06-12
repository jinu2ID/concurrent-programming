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
}


func orGate(in1 chan bool, in2 chan bool, out chan bool, wg *sync.WaitGroup){
	a := <-in1
	b := <-in2

	in1 <-a
	in2 <-b

	out <- (a || b)
	wg.Done()
}

func nandGate(in1 chan bool, in2 chan bool, out chan bool, wg *sync.WaitGroup){
	a := <-in1
	b := <-in2

	in1 <-a
	in2 <-b

	out <- !(a && b)
	wg.Done()
}

func norGate(in1 chan bool, in2 chan bool, out chan bool, wg *sync.WaitGroup){
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
}


// Stores input values
func input(in chan bool, val bool, wg *sync.WaitGroup){
	in <-val
	wg.Done()
}


// Stores carry values
func carry(in chan bool, out chan bool, wg *sync.WaitGroup){
	x := <-in
	out <-x
	wg.Done()
}


// Retrieves final output
func output(out chan bool, index int, wg *sync.WaitGroup){
	x := <-out
	if x {
		finalOutput[index] = 1
	} else {
		finalOutput[index] = 0
	}
	wg.Done()
}


func simulate(instr []string, outputs int) {

	var wg sync.WaitGroup
	var channels []chan bool
	outIndex := outputs

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

	var file string
	var outputs int

	if len(os.Args) > 2 {
		file = os.Args[1]
		if value,err := strconv.Atoi(os.Args[2]); err == nil {
			outputs = value
		} else {
			fmt.Println("Incorrect arguments, please check README")
			return
		}

	} else {
		fmt.Println("Incorrect arguments, please check README")
		return
	}

	var lines = read(file)
	if len(lines) == 0 {
		fmt.Println("Error: could not find file")
	}

	finalOutput = make([]int, outputs+1)
	simulate(lines, outputs)
}
