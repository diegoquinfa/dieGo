package todos

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type Todo struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type TodoJson struct {
	UrgentTodos []Todo `json:"urgentTodos"`
	NormalTodos []Todo `json:"normalTodos"`
}

func NewTodo(name, description string) (*Todo, error) {
	if name == "" || description == "" {
		return nil, fmt.Errorf("NEW_TODO_FORMAT_ERROR")
	}

	now := time.Now().Format("02/01/2006 - 15:04")

	return &Todo{
		Name:        name,
		Description: description,
		Complete:    false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (tj *TodoJson) AddTodo(newTodo *Todo, isUrgent bool) {
	newTodo.Id = tj.getLastId(isUrgent) + 1

	if isUrgent {
		tj.UrgentTodos = append(tj.UrgentTodos, *newTodo)
	} else {
		tj.NormalTodos = append(tj.NormalTodos, *newTodo)
	}
}

func (tj *TodoJson) SaveTodo(file *os.File) {
	bytes, err := json.Marshal(tj)
	if err != nil {
		panic(err)
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	file.Truncate(0)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	_, err = writer.Write(bytes)

	if err != nil {
		panic(err)
	}

	err = writer.Flush()

	if err != nil {
		panic(err)
	}
}

func OpenTodosFile() (*os.File, *TodoJson, error) {
	file, err := os.OpenFile("todos.json", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return nil, nil, err
	}

	info, err := file.Stat()

	if err != nil {
		return nil, nil, err
	}

	var todos TodoJson
	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)

		if err != nil {
			return nil, nil, err
		}

		err = json.Unmarshal(bytes, &todos)

		if err != nil {
			return nil, nil, err
		}
	} else {
		todos = TodoJson{}
	}

	return file, &todos, err
}

func (tj *TodoJson) getLastId(isUrgent bool) int {
	if isUrgent {
		return len(tj.UrgentTodos)
	}

	return len(tj.NormalTodos)
}
