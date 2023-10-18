package gui

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/diegoquinfa/dieGo/todos"
)

var file, myTodos, _ = todos.OpenTodosFile()

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		panic(err)
	}
}

func Add(isUrgent bool, Args []string) {
	defer closeFile(file)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Todo description: ")

	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	todoName := flag.Arg(1)
	newTodo, err := todos.NewTodo(todoName, description)
	if err != nil {
		panic(err)
	}

	myTodos.AddTodo(newTodo, isUrgent)
	myTodos.SaveTodo(file)

	fmt.Println("\nAdded todo:", flag.Arg(1))
}
