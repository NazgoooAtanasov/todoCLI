package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const todoRegex string = "^\\s*/{2} (TODO)(O*): (\\w*)$"

type Todo struct {
    FileName string
	LineNumber int
	LineText string
	Urgency int
}

func CheckTodo(line string) bool {
	check, err := regexp.MatchString(todoRegex, line)

	if err != nil {
		panic(err)
	}

	return check
}

func FormatString(todo *Todo) string {
	var urgencyLenght string = strings.Repeat("O", todo.Urgency)
	return fmt.Sprintf("%s:%d: TODO%s: %s\n",
		todo.FileName,
		todo.LineNumber,
		urgencyLenght,
		todo.LineText,
	)
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

func main() {
	cwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	err = filepath.Walk(cwd, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}

		if !info.IsDir() {
			file, err := os.Open(path)

			if err != nil {
				panic(err)
			}

			defer file.Close()
			scanner := bufio.NewScanner(file)

			var listOfTodos []*Todo
			var lineNumber int = 0
			for scanner.Scan() {
				lineNumber++
				var line string = scanner.Text()
				var isToDo bool = CheckTodo(line)

				if isToDo {
					todo := ExtactTodo(line, path, lineNumber)
					listOfTodos = append(listOfTodos, todo)
				}
			}

			if err := scanner.Err(); err != nil {
				panic(err)
			}

			sort.Slice(listOfTodos, func(i, j int) bool {
				return listOfTodos[i].Urgency > listOfTodos[j].Urgency
			})

			for _, element := range listOfTodos {
				fmt.Print(FormatString(element))
			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}
