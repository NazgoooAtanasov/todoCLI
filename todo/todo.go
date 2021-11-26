package todo

import (
	"fmt"
	"regexp"
	"strings"
)

const todoRegex string = "^\\s*/{2} (TODO)(O*): (\\w*)$"

type Todo struct {
    FileName string
	LineNumber int
	LineText string
	Urgency int
}

func (todo *Todo) FormatString() string {
	var urgencyLenght string = strings.Repeat("O", todo.Urgency)
	return fmt.Sprintf("%s:%d: TODO%s: %s\n",
		todo.FileName,
		todo.LineNumber,
		urgencyLenght,
		todo.LineText,
	)
}

func CheckTodo(line string) bool {
	check, err := regexp.MatchString(todoRegex, line)

	if err != nil {
		panic(err)
	}

	return check
}

func ExtactTodo(line string, fileName string, lineNumber int) *Todo {
	reg := regexp.MustCompile(todoRegex)

	groups := reg.FindStringSubmatch(line)

	return &Todo {
		FileName: fileName,
		LineNumber: lineNumber,
		LineText: groups[3],
		Urgency: len(groups[2]),
	}
}
