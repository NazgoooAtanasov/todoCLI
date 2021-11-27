package todo

import (
	"fmt"
	"regexp"
	"strings"
)

type Todo struct {
    FileName string
	LineNumber int
	LineText string
	Urgency int
}

func (todo *Todo) FormatString(keyword string, urgencySuffix string) string {
	var urgencyLenght string = strings.Repeat(urgencySuffix, todo.Urgency)
	return fmt.Sprintf("%s:%d: %s%s: %s\n",
		todo.FileName,
		todo.LineNumber,
		keyword,
		urgencyLenght,
		todo.LineText,
	)
}

func CheckTodo(line string, todoRegex string) bool {
	check, err := regexp.MatchString(todoRegex, line)

	if err != nil {
		panic(err)
	}

	return check
}

func ExtactTodo(line string, fileName string, lineNumber int, reg *regexp.Regexp) *Todo {
	groups := reg.FindStringSubmatch(line)

	return &Todo {
		FileName: fileName,
		LineNumber: lineNumber,
		LineText: groups[3],
		Urgency: len(groups[2]),
	}
}
