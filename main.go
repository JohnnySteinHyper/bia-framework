package main

import (
	"os"
	"github.com/anonymousdevhacker/bia/extras"
)

// docker run <container> cmd args
// go run main.go cmd args

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("what??")
	}

}

func run() {
	switch os.Args[2] {
	case "print":
		arg1 := os.Args[3:]
		extras.Print([]string(arg1))
	}
}
