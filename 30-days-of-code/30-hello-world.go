// https://www.hackerrank.com/challenges/30-hello-world/problem
package main

import (
	"bufio"
	"fmt"
	"os"
)

func helloWorld() {
	reader := bufio.NewReader(os.Stdin)

	inputString, _ := reader.ReadString('\n')

	fmt.Println("Hello, World.")
	fmt.Println(inputString)
}
