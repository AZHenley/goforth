package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type environment struct {
	stack    []int
	auxStack []int
	words    map[string]string
}

func (env *environment) push(i int) {
	env.stack = append(env.stack, i)
}

func (env *environment) pop() int {
	v := env.stack[len(env.stack)-1]
	env.stack = env.stack[:len(env.stack)-1]
	return v
}

func (env *environment) pushAux(i int) {
	env.auxStack = append(env.auxStack, i)
}

func (env *environment) popAux() int {
	v := env.auxStack[len(env.auxStack)-1]
	env.auxStack = env.auxStack[:len(env.auxStack)-1]
	return v
}

func (env *environment) top() int {
	v := env.stack[len(env.stack)-1]
	return v
}

// Gets value X elements from top
func (env *environment) get(i int) int {
	v := env.stack[len(env.stack)-1-i]
	return v
}

func main() {
	fmt.Println("Goforth.")
	reader := bufio.NewReader(os.Stdin)
	var env environment

	// REPL
	for {
		// Read.
		fmt.Print(" > ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		// Evaluate.
		if text == "exit" || text == "quit" {
			return
		}
		eval(&env, text)

		// print
		fmt.Print(env.stack)
	}
}

// TODO: Error handling causes a lot of duplicate code. Unsure how to fix without exceptions.
func eval(env *environment, code string) {
	skipIf := 0            // Keep track of nested If count
	skipElse := 0          // Keep track of nested Else count
	expectingThen := false // Inside an If
	tokens := strings.Fields(code)
	for _, token := range tokens {

		// Based on previous if, skip until...
		// TODO: Refactor this.
		if skipIf > 0 {
			if token == "if" {
				skipIf++
			} else if token == "else" && skipIf == 1 {
				skipIf = 0
			} else if token == "then" {
				skipIf--
			}
			continue
		}
		if skipElse > 0 {
			if token == "if" {
				skipElse++
			} else if token == "then" {
				skipElse--
			}
			continue
		}

		switch token {
		// Arithmetic.
		case "+":
			if len(env.stack) < 2 {
				error("Stack underflow.")
				return
			}
			op1 := env.pop()
			op2 := env.pop()
			env.push(op1 + op2)
		case "-":
			if len(env.stack) < 2 {
				error("Stack underflow.")
				return
			}
			op1 := env.pop()
			op2 := env.pop()
			env.push(op1 - op2)
		case "*":
			if len(env.stack) < 2 {
				error("Stack underflow.")
				return
			}
			op1 := env.pop()
			op2 := env.pop()
			env.push(op1 * op2)
		case "/":
			if len(env.stack) < 2 {
				error("Stack underflow.")
				return
			}
			op1 := env.pop()
			op2 := env.pop()
			if op2 == 0 {
				error("Divide by zero.")
				return
			}
			env.push(op1 / op2)
		case "mod":
			if len(env.stack) < 2 {
				error("Stack underflow.")
				return
			}
			op1 := env.pop()
			op2 := env.pop()
			if op2 == 0 {
				error("Divide by zero.")
				return
			}
			env.push(op1 % op2)
		// Stack manipulation.
		case "dup":
			if len(env.stack) < 1 {
				error("Stack underflow.")
				return
			}
			env.push(env.top())
		case "drop":
			if len(env.stack) < 1 {
				error("Stack underflow.")
				return
			}
			env.pop()
		case "swap":
			if len(env.stack) < 2 {
				error("Stack underflow.")
				return
			}
			op1 := env.pop()
			op2 := env.pop()
			env.push(op1)
			env.push(op2)
		case "rot":
			if len(env.stack) < 3 {
				error("Stack underflow.")
				return
			}
			op1 := env.pop()
			op2 := env.pop()
			op3 := env.pop()
			env.push(op2)
			env.push(op1)
			env.push(op3)
		case "over":
			if len(env.stack) < 2 {
				error("Stack underflow.")
				return
			}
			env.push(env.get(1))
		// Aux stack.
		case "cross": // Moves top of the stack over to a secondary stack.
			if len(env.stack) < 1 {
				error("Stack underflow.")
				return
			}
			env.pushAux(env.pop())
		case "back": // Moves top of the secondary stack over to the stack.
			if len(env.auxStack) < 1 {
				error("Stack underflow.")
				return
			}
			env.push(env.popAux())
		// Comparison.
		case "=":
			if len(env.stack) < 2 {
				error("Stack underflow.")
				return
			}
			op1 := env.get(1)
			op2 := env.top()
			if op1 == op2 {
				env.push(1)
			} else {
				env.push(0)
			}
		case "<":
			if len(env.stack) < 2 {
				error("Stack underflow.")
				return
			}
			op1 := env.get(1)
			op2 := env.top()
			if op1 < op2 {
				env.push(1)
			} else {
				env.push(0)
			}
		case ">":
			if len(env.stack) < 2 {
				error("Stack underflow.")
				return
			}
			op1 := env.get(1)
			op2 := env.top()
			if op1 > op2 {
				env.push(1)
			} else {
				env.push(0)
			}
		// Control flow.
		case "if":
			if len(env.stack) < 1 {
				error("Stack underflow.")
				return
			}
			op1 := env.pop()
			if op1 != 0 {
				//skipElse = 1
			}
			if op1 == 0 {
				skipIf = 1
			}
			expectingThen = true
		case "else":
			// Must have already executed If, so skip Else.
			skipElse = 1
		case "then":
			if !expectingThen {
				error("Unexpected then.")
				return
			}
			expectingThen = false
		case ":":
		case ";":
		case "@": // Starts a label.
		case "goto":

		default:
			i, err := strconv.Atoi(token)
			if err != nil {
				error("Invalid word: " + token)
				return
			}
			env.push(i)
		}
	}
}

func error(msg string) {
	fmt.Println(msg)
}
