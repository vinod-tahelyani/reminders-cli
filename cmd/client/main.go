package main

import (
	"flag"
	"fmt"
	"os"
	"reminder-cli.example.com/client"
)
var (
	backendURIFlag = flag.String("backend", "http://localhost:8080", "Backend API URL")
	helpflag = flag.Bool("help", false, "Display a helpful message")
)

func main() {
	flag.Parse()
	s := client.NewSwitch(*backendURIFlag)


	if *helpflag || len(os.Args) == 1 {
		s.Help()
		return
	}

	err := s.Switch()
	if err != nil {
		fmt.Printf("cmd switch error\n")
		os.Exit(2)
	}
}
