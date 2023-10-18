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

func Add(isUrgent bool) {
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

func Delete(isUrgent bool) {
	defer closeFile(file)

	reader := bufio.NewReader(os.Stdin)

	err := myTodos.DeleteTodo(flag.Arg(1))
	if err != nil {
		res, ok := todos.ERRORS[err.Error()]
		if ok {
			fmt.Println(res)
		} else {
			fmt.Println("jumm", err)
		}
		return
	}

	fmt.Print("Are sure want to delete this todo? (y/n): ")

	option, _ := reader.ReadString('\n')
	option = strings.TrimSpace(option)
	option = strings.ToLower(option)

	if option != "y" {
		fmt.Println("\nNo se elimino la monda esa por subnormal")
		return
	}

	myTodos.SaveTodo(file)

	fmt.Println("Bueno se fue la tarea")

}
