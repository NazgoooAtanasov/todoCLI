package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"todoCLI/todo"
	"todoCLI/config"
)

func main() {
	cwd, err := os.Getwd()
	config := config.GetConfig()

	if err != nil {
		panic(err)
	}

	var listOfTodos []*todo.Todo

	fmt.Println("[WARN]: Walking " + cwd)
	err = filepath.Walk(cwd, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}

		if !info.IsDir() && !strings.Contains(path, ".git") {
			file, err := os.Open(path)

			if err != nil {
				panic(err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			var lineNumber int = 0
			for scanner.Scan() {
				lineNumber++
				var line string = scanner.Text()
				var isToDo bool = todo.CheckTodo(line, config.ParsingRegex.String())

				if isToDo {
					todo := todo.ExtactTodo(line, path, lineNumber, config.ParsingRegex)
					listOfTodos = append(listOfTodos, todo)
				}
			}

			if err := scanner.Err(); err != nil {
				panic(err)
			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	sort.Slice(listOfTodos, func(i, j int) bool {
		return listOfTodos[i].Urgency > listOfTodos[j].Urgency
	})

	for _, element := range listOfTodos {
		fmt.Print(element.FormatString(config.Pattern.Keyword, config.Pattern.UrgencySuffix))
	}
}
