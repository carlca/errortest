package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func exampleNewprintf() {
	err := errors.New("whoops")
	fmt.Printf("%+v", err)
	fmt.Printf("Hello")
	fmt.Printf("Hello")

	// Example output:
	// whoops
	// github.com/pkg/errors_test.ExampleNew_printf
	//         /home/dfc/src/github.com/pkg/errors/example_test.go:17
	// testing.runExample
	//         /home/dfc/go/src/testing/example.go:114
	// testing.RunExamples
	//         /home/dfc/go/src/testing/example.go:38
	// testing.(*M).Run
	//         /home/dfc/go/src/testing/testing.go:744
	// main.main
	//         /github.com/pkg/errors/_test/_testmain.go:106
	// runtime.main
	//         /home/dfc/go/src/runtime/proc.go:183
	// runtime.goexit
	//         /home/dfc/go/src/runtime/asm_amd64.s:2059
}

func main() {
	exampleNewprintf()
}
