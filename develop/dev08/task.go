package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var currFolder string

func main() {
	wd, err := os.Getwd()
	if err != nil {
		println(err)
		return
	}

	currFolder = wd

	reader := bufio.NewReader(os.Stdin)

	for {
		printReadyForInputString()
		input, err := reader.ReadString('\n')
		if err != nil {
			println(err.Error())
			return
		}

		commands := strings.Split(input, "|")

		var pipeInput, out string

		// Конвеер на пайпах
		for _, commandWithArgs := range commands {
			commandWithArgs = strings.TrimSpace(commandWithArgs)
			split := strings.Split(commandWithArgs, " ")

			if len(split) == 0 {
				continue
			}

			command := split[0]
			args := split[1:]
			args = append(args, pipeInput)

			out, err = handleCommand(command, args...)
			if err != nil {
				println(err.Error())
				break
			}

			pipeInput = out
		}

		println(out)
	}
}

func handleCommand(command string, args ...string) (string, error) {
	switch command {
	case "echo":
		for i, val := range args {
			if strings.HasPrefix(val, "$") {
				env := os.Getenv(strings.TrimLeft(val, "$"))
				if env != "" {
					args[i] = env
				}
			}
		}

		return fmt.Sprintf("%s\n", strings.Join(args, " ")), nil

	case "pwd":
		return unsafeGetWd(), nil

	case "cd":
		if len(args) == 0 {
			return "", errors.New("directory must be provided")
		}

		err := os.Chdir(args[0])
		if err != nil {
			return "", err
		}

	case "ls":
		dirs, err := osReadDir(unsafeGetWd())
		if err != nil {
			return "", err
		}

		return strings.Join(dirs, "\n"), nil

	case "ps":
		out, err := exec.Command("ps").Output()
		if err != nil {
			return "", err
		}
		return string(out), nil

	case "kill":
		out, err := exec.Command("kill", args...).Output()
		if err != nil {
			return "", err
		}
		return string(out), nil

	case "exit":
		os.Exit(1)
	}
	return fmt.Sprintf("command '%s' does not exist", command), nil
}

func printReadyForInputString() {
	fmt.Printf("Shell %s->", unsafeGetWd())
}

func osReadDir(root string) ([]string, error) {
	var files []string
	f, err := os.Open(root)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	err = f.Close()
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}

func unsafeGetWd() (wd string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return
}
