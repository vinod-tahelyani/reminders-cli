package main

import (
	"flag"
	"fmt"
	"os"
	"log"
)

func main()  {
	if len(os.Args) < 2 {
		fmt.Println("no commmand provided")
		os.Exit(2)
	}

	cmd := os.Args[1]

	switch cmd {
	case "greet" :
		greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)
		msgFlag := greetCmd.String("msg", "CLI BASICS - REMINDERS CLI", "the message for greet")
		err := greetCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("hello and welcome: %s", *msgFlag)
	case "help" :
		fmt.Println("some help msg")
	default :
		fmt.Printf("Unknown command: %s\n", cmd)
	}
	fmt.Println(os.Args)
}