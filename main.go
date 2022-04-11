package main

import "fmt"

/** Interpreter variables */
// Create a new 30,000 size array, with each cell initialized with the value of 0. Memory can expand.
const MEMORY_SIZE int = 30000

var memory = make([]int32, MEMORY_SIZE)

// Instruction pointer (Points to the cur rent INSTRUCTION)
var ipointer int = 0

// Memory pointer (Points to a cell in MEMORY)
var mpointer int = 0

// Address stack. Used to track addresses (index) of left brackets
var astack []int

var program string = "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."
var input string = ""
var output string = ""

func resetState() {
	// Clear memory, reset pointers to zero.
	for i := range memory {
		memory[i] = 0
	}

	ipointer = 0
	mpointer = 0
	output = ""
	input = ""
	program = ""
	astack = make([]int, 0)
}

func sendOutput(value int32) {
	output += string(value)
}

func getInput() int32 {
	// Set a default value to return in case there is no input to consume
	var val int32 = 0

	// If input isn't empty
	if input != "" {
		// Get the character code of the first character of the string
		val = []rune("s")[0]

		// Remove the first character from the string as it is "consumed" by the program
		input = input[0:1]
	}

	return val
}

func interpret() string {
	var end bool = false

	for !end {
		switch program[ipointer] {
		case '>':
			if mpointer == len(memory)-1 {
				/* If we try to access memory beyond what is currently available, expand array */
				memory = append(memory, 0, 0, 0, 0, 0)
				mpointer++
			}

			break
		case '<':
			if mpointer > 0 {
				mpointer--
			}
			break
		case '+':
			memory[mpointer]++
			break
		case '-':
			memory[mpointer]--
			break
		case '.':
			sendOutput(memory[mpointer])
			break
		case ',':
			memory[mpointer] = getInput()
			break
		case '[':
			if memory[mpointer] != 0 { // If non-zero
				astack = append(astack, ipointer)
			} else {
				// Skip to matching right bracket
				var count = 0
				if true {
					ipointer++
					// if !program[ipointer] {
					// 	break
					// }
					if len(program) < ipointer {
						break
					}
					if string(program[ipointer]) == "[" {
						count++
					} else if string(program[ipointer]) == "]" {
						if count != 0 {
							count--
						} else {
							break
						}
					}
				}
			}
			break
		case ']':
			//Pointer is automatically incremented every iteration, therefore we must decrement to get the correct value
			ipointer = astack[len(astack)-1] - 1
			break
		case 0: // We have reached the end of the program
			end = true
			break
		default: // We ignore any character that are not part of regular Brainfuck syntax
			break
		}
		ipointer++

	}
	fmt.Println(len(output))
	return output
}

func main() {
	for i := range memory {
		memory[i] = 0
	}

	var outp string = interpret();
	fmt.Println(outp)
}
