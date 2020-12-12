package main

import (
	"./core"
	"flag"
	"log"
)

func main() {

	arg := flag.String("ouput", "web", "specify where to you the data to be displayed")
	flag.Parse()

	if (*arg)=="console" {
		log.Print("Starting CONSOLE")
		core.HandleConsole()
	} else if (*arg)=="web" {
		log.Print("Starting WEB")
		core.HandleWeb()
	} else {
		log.Print("This is not a valid argument for -output flag : ", arg)
	}


}

