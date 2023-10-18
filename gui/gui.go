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
		fmt.Println("Error:", err)
		return
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
		fmt.Println("Error:", err)
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

func List() {
	fmt.Println("                  Urgent todo")
	fmt.Println(" id     nombre    descripción      fecha de creación")
	for _, todo := range myTodos.UrgentTodos {
		status := " "
		if todo.Complete {
			status = "✔️"
		}
		fmt.Printf("[%s] -> %s | %s | %s \n", status, todo.Name, todo.Description, todo.CreatedAt)
	}

	for _, todo := range myTodos.NormalTodos {
		status := " "
		if todo.Complete {
			status = "✔️"
		}
		fmt.Printf("[%s] -> %s | %s | %s", status, todo.Name, todo.Description, todo.CreatedAt)
	}
}
