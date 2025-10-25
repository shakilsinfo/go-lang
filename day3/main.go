package main

import (
	"fmt"
	"day3/greetings"
)

func main() {
	message := greetings.SayHello("Alice")
	fmt.Println(message)
}

// If we run this file, it will print: Hello, Alice!
// But if we are trying to go build the project, it will throw an error because there are multiple main functions in the project.
// go: cannot find main module, but found .git/config in /Users/shakelahamed/WorkStations/go-lang
// to create a module there, run:
// cd .. && go mod init
// then try to build again.
// then ./test show the output
