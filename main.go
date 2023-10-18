package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/diegoquinfa/dieGo/todos"
)

func main() {
	isUrgent := flag.Bool("urgent", false, "Marca la tarea como prioritaria.")
	flag.Parse()

	command := flag.Arg(0)

	if len(flag.Args()) < 2 {
		fmt.Println(`Use: dieGo [--urgent] <command> [arg]`)
		fmt.Println(*isUrgent)
		return
	}
	file, myTodos, _ := todos.OpenTodosFile()

	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	switch command {
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Todo description: ")

		description, _ := reader.ReadString('\n')
		description = strings.TrimSpace(description)

		todoName := flag.Arg(1)
		newTodo, err := todos.NewTodo(todoName, description)
		if err != nil {
			panic(err)
		}

		if *isUrgent {
			myTodos.UrgentTodos = todos.AddTodo(myTodos, newTodo, *isUrgent)
		} else {
			myTodos.NormalTodos = todos.AddTodo(myTodos, newTodo, *isUrgent)
		}

		todos.SaveTodo(file, myTodos)

		fmt.Println("\nAdded todo:", flag.Arg(1))
	}

}
