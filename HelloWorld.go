package main

import "fmt"

const breakLine string = "\n"
const timesLoop int = 15

type helloWorld struct {
	say string
	times int
}

func main() {
	var str01 string = "Hello"
	var str02 string = "World"	
	var doLoop bool = true
	var typeDecision int = 1

	if doLoop {

		switch typeDecision {
		case 0:
			var max int = 0
			
			max = timesLoop
	
			for loop := 1 ; loop <= max ; loop++ {
				fmt.Printf( str01 + " " + str02 + breakLine )
	
				if loop == max {
					break
				}
			}
		case 1:
			var loop int = 1
			var xHelloWorld helloWorld

			xHelloWorld.say = "Teste"
			xHelloWorld.times = timesLoop

			for loop <= xHelloWorld.times {
				fmt.Printf( str01 + " " + str02 + breakLine )
				loop++
			}
		default:
			fmt.Printf( getHelloWorld() + breakLine )
		}
	} else {
		fmt.Printf( getHelloWorld() + breakLine )
	}
}

func getHelloWorld() string {
	var str_return string
	str_return = "Hello World"

	return str_return
}
