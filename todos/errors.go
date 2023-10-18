package todos

var ERRORS = map[string]string{
	"DONT_EXIST_TODO": "Warning: Don't exist.",
}

type DONT_EXIST_TODO struct{}

func (e *DONT_EXIST_TODO) Error() string {
	return "El ID proporcionado, no existe."
}

type NEW_TODO_FORMAT_ERROR struct{}

func (e *NEW_TODO_FORMAT_ERROR) Error() string {
	return "Nombre o descripción invalidos."
}
