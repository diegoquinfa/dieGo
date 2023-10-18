package main

import (
	"flag"
	"fmt"

	"github.com/diegoquinfa/dieGo/gui"
)

func main() {
	isUrgent := flag.Bool("u", false, "Marca la tarea como prioritaria.")
	flag.Parse()

	command := flag.Arg(0)

	if len(flag.Args()) < 2 {
		fmt.Println(`Use: dieGo [-u] <command> [arg]`)
		fmt.Println(*isUrgent)
		return
	}

	switch command {
	case "add":
		gui.Add(*isUrgent)
	case "delete":
		gui.Delete(*isUrgent)
	}

}
