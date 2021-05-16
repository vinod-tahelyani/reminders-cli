package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	if len(os.Args) < 2 {
		fmt.Println("no commmand provided")
		os.Exit(2)
	}

	cmd := os.Args[1] 

	switch cmd {
	case "greet" :
		msg := "REMINDERS CLI - CLI BASICS"
		f := strings.Split(os.Args[2], "=")
		if len(f) == 2 && f[0] == "--msg" {
			msg = f[1]
		}
		fmt.Printf("hello and welcome %s\n", msg)
	case "help" :
		fmt.Println("some help msg")
	default :
		fmt.Printf("Unknown command: %s\n", cmd)
	}
	fmt.Println(os.Args)
}