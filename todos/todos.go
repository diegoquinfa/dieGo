package todos

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"strconv"
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

func NewTodo(name, description string) (*Todo, error) {
	if name == "" || description == "" {
		return nil, &NEW_TODO_FORMAT_ERROR{}
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

func (tj *TodoJson) DeleteTodo(deleteId string) error {
	var (
		todoList *[]Todo
		id       int
		err      error
	)

	if deleteId[0] == 'U' {
		id, err = strconv.Atoi(deleteId[1:])
		todoList = &tj.UrgentTodos

	} else {
		id, err = strconv.Atoi(deleteId)
		todoList = &tj.NormalTodos
	}

	if err != nil {
		return &DONT_EXIST_TODO{}
	}

	todoListCopy := make([]Todo, len(*todoList))
	copy(todoListCopy, *todoList)

	for i, todo := range *todoList {
		if todo.Id == id {
			*todoList = append(todoListCopy[:i], todoListCopy[i+1:]...)
			return nil
		}
	}

	return &DONT_EXIST_TODO{}
}

func (tj *TodoJson) CompleteTodo(updateId string) error {
	var (
		id  int
		err error
	)

	if updateId[0] == 'U' {
		id, err = strconv.Atoi(updateId[1:])
		for i, todo := range tj.UrgentTodos {
			if todo.Id == id {
				tj.UrgentTodos[i].Complete = true
				tj.UrgentTodos[i].UpdatedAt = time.Now().Format("02/01/2006 - 15:04")
				return nil
			}
		}
	} else {
		id, err = strconv.Atoi(updateId)
		for i, todo := range tj.NormalTodos {
			if todo.Id == id {
				tj.NormalTodos[i].Complete = true
				tj.NormalTodos[i].UpdatedAt = time.Now().Format("02/01/2006 - 15:04")
				return nil
			}
		}
	}

	if err != nil {
		return &DONT_EXIST_TODO{}
	}

	return &DONT_EXIST_TODO{}
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

func (tj *TodoJson) getLastId(isUrgent bool) int {
	if isUrgent {
		if len(tj.UrgentTodos) == 0 {
			return 0
		}

		return tj.UrgentTodos[len(tj.UrgentTodos)-1].Id
	}

	if len(tj.NormalTodos) == 0 {
		return 0
	}

	return tj.NormalTodos[len(tj.NormalTodos)-1].Id
}
