package main

import (
	"bufio"
	"fmt"
	"os"
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
	tokens := strings.Fields(code)
	for _, token := range tokens {
		switch token {

		}
	}
}

func error(msg string) {
	fmt.Println(msg)
}
