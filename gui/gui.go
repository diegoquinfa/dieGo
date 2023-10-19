package gui

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/diegoquinfa/dieGo/table"
	"github.com/diegoquinfa/dieGo/todos"
)

const (
	red        = "\x1b[31m"
	green      = "\x1b[32m"
	clearColor = "\x1b[0m"
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
	fmt.Println("\nUrgent Todo")
	urgentTodosRows := []table.Row{}
	urgentTodosRows = append(
		urgentTodosRows,
		*table.NewRow("Id", "Status", "Name", "Create at"),
	)

	for _, todo := range myTodos.UrgentTodos {
		status := "\x1b[31m" + "Uncomplete" + "\x1b[0m"
		if todo.Complete {
			status = "\x1b[32m" + "Complete" + "\x1b[0m"
		}
		id := "U" + fmt.Sprintf("%d", todo.Id)
		urgentTodosRows = append(
			urgentTodosRows,
			*table.NewRow(id, status, todo.Name, todo.CreatedAt),
		)
	}

	table.CreateTable(urgentTodosRows)

	normalTodosRows := []table.Row{}
	normalTodosRows = append(
		normalTodosRows,
		*table.NewRow("Id", "Status", "Name", "Create at"),
	)

	for _, todo := range myTodos.NormalTodos {
		status := red + "Uncomplete" + clearColor
		if todo.Complete {
			status = green + "Complete" + clearColor
		}
		normalTodosRows = append(
			normalTodosRows,
			*table.NewRow(todo.Id, status, todo.Name, todo.CreatedAt),
		)
	}

	fmt.Println("\nNormal Todo")
	table.CreateTable(normalTodosRows)
}

func Complete() {
	defer closeFile(file)

	err := myTodos.CompleteTodo(flag.Arg(1))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	myTodos.SaveTodo(file)
	fmt.Println("Todo was update!")
}

func Details() {

	detailId := flag.Arg(1)

	if detailId[0] == 'U' {
		id, err := strconv.Atoi(detailId[1:])
		if err != nil {
			fmt.Println(err)
		}

		for _, todo := range myTodos.UrgentTodos {
			if todo.Id == id {
				urgentTodosRows := []table.Row{}
				urgentTodosRows = append(
					urgentTodosRows,
					*table.NewRow("Id", "Status", "Name", "Description", "Create at",
						"Update at",
					),
				)

				status := red + "Uncomplete" + clearColor
				if todo.Complete {
					status = red + "Complete" + clearColor
				}
				id := "U" + fmt.Sprintf("%d", todo.Id)
				urgentTodosRows = append(
					urgentTodosRows,
					*table.NewRow(id, status, todo.Name, todo.Description,
						todo.CreatedAt, todo.UpdatedAt,
					),
				)

				table.CreateTable(urgentTodosRows)
				return
			}
		}
	} else {
		id, err := strconv.Atoi(detailId)
		if err != nil {
			fmt.Println(err)
		}

		for _, todo := range myTodos.NormalTodos {
			if todo.Id == id {
				normalTodosRows := []table.Row{}
				normalTodosRows = append(
					normalTodosRows,
					*table.NewRow("Id", "Status", "Name", "Description", "Create at",
						"Update at",
					),
				)

				status := red + "Uncomplete" + clearColor
				if todo.Complete {
					status = green + "Complete" + clearColor
				}
				normalTodosRows = append(
					normalTodosRows,
					*table.NewRow(todo.Id, status, todo.Name, todo.Description,
						todo.CreatedAt, todo.UpdatedAt,
					),
				)

				table.CreateTable(normalTodosRows)
				return
			}
		}
	}
}
