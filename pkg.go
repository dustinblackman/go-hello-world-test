package helloworld

import (
	"fmt"
	"os"
	"strconv"
)

func SayHi() {
	if len(os.Args) > 1 {
		exitCode, err := strconv.Atoi(os.Args[1])
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		os.Exit(exitCode)
	}

	fmt.Println("Hello World")
}
