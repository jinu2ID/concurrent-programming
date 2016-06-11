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



// Gate routines
func inverter(in chan bool, out chan bool, wg *sync.WaitGroup){

	a := <-in
	out <- !a
	wg.Done()
}


func andGate(in1 chan bool, in2 chan bool, out chan bool, wg *sync.WaitGroup){
	a := <-in1
	b := <-in2
	out <- (a && b)
	wg.Done()

}


func orGate(in1 chan bool, in2 chan bool, out chan bool, wg *sync.WaitGroup){
	a := <-in1
	b := <-in2

	out <- (a || b)
	wg.Done()
}

func nandGate(in1 chan bool, in2 chan bool, out chan bool, wg *sync.WaitGroup){
	a := <-in1
	b := <-in2

	out <- !(a && b)
	wg.Done()
}

func norGate(in1 chan bool, in2 chan bool, out chan bool, wg *sync.WaitGroup){
	a := <-in1
	b := <-in2

	out <- !(a || b)
	wg.Done()
}

func xorGate(in1 chan bool, in2 chan bool, out chan bool, wg *sync.WaitGroup){
	a := <-in1
	b := <-in2
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

// Retrieves final output
func output(out chan bool, main chan bool, last bool, wg *sync.WaitGroup){
	x := <-out
	if x {
		fmt.Print(1)
	} else {
		fmt.Print(0)
	}
	wg.Done()
}

// fan out
func fan(target chan bool, n int, wg *sync.WaitGroup){

	x := <-target
	for i := 0; i < n+1; i++ {
		target <-x
	}
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

func simulate(instr []string, lc int, result chan bool) {

	var wg sync.WaitGroup
	wg.Add(lc)
	var channels []chan bool

	for i, val := range instr {
		channel := make(chan bool, 2)

		values := strings.Split(val, ",")

		// Inputs
		if values[0] == "0" {
		//	wg.Add(1)
			go input(channel, false, &wg)
		} else if values[0] == "1" {
		//	wg.Add(1)
			go input(channel, true, &wg)
		} else if values[0] == "fan"{

			input1, err1 := strconv.Atoi(values[1]);
			input2, err2 := strconv.Atoi(values[2]);

			if  err1 == nil && err2 == nil {
		//		wg.Add(1)
				go fan(channels[input1], input2, &wg)
			} else {
				fmt.Println("format error in file")
			}

		} else if values[0] == "AndGate"{

			input1, err1 := strconv.Atoi(values[1]);
			input2, err2 := strconv.Atoi(values[2]);

			if  err1 == nil && err2 == nil {
		//		wg.Add(1)
				go andGate(channels[input1], channels[input2], channel, &wg)

			} else {
				fmt.Println("format error in file")
			}

		} else if values[0] == "OrGate" {

			input1, err1 := strconv.Atoi(values[1]);
			input2, err2 := strconv.Atoi(values[2]);

			if  err1 == nil && err2 == nil {
		//		wg.Add(1)
				go orGate(channels[input1], channels[input2], channel, &wg)
			} else {
				fmt.Println("format error in file")
			}


		} else if values[0] == "NandGate" {

			input1, err1 := strconv.Atoi(values[1]);
			input2, err2 := strconv.Atoi(values[2]);

			if  err1 == nil && err2 == nil {
		//		wg.Add(1)
				go nandGate(channels[input1], channels[input2], channel, &wg)
			} else {
				fmt.Println("format error in file")
			}

		} else if values[0] == "NorGate" {

			input1, err1 := strconv.Atoi(values[1]);
			input2, err2 := strconv.Atoi(values[2]);

			if  err1 == nil && err2 == nil {
		//		wg.Add(1)
				go norGate(channels[input1], channels[input2], channel, &wg)
			} else {
				fmt.Println("format error in file")
			}

		} else if values[0] == "XorGate" {

			input1, err1 := strconv.Atoi(values[1]);
			input2, err2 := strconv.Atoi(values[2]);

			if  err1 == nil && err2 == nil {
		//		wg.Add(1)
				go xorGate(channels[input1], channels[input2], channel, &wg)
			} else {
				fmt.Println("format error in file")
			}

		} else if values[0] == "Output" {
				if index, err := strconv.Atoi(values[1]); err == nil{


				if i == lc-1 {
		//			wg.Add(1)
					go output(channels[index], result, true, &wg)

				} else {
		//			wg.Add(1)
					go output(channels[index], result, false, &wg)

				}


			} else {
				fmt.Println("format error in file")
			}

		}

		channels = append(channels, channel)

	}
	wg.Wait()
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
	//for x := range res{
	//	fmt.Println(x)
	//}

}
