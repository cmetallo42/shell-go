package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	messageWelcome = `Welcome to shell-go`
	messageHelp    = `You can type all basic linux shell commands`
	messageExit    = `Exiting...`
)

func getDirectory() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	splitedPath := strings.Split(path, "/")

	return splitedPath[len(splitedPath)-1]
}

var ErrNotFound = errors.New("command not found")

func internal(args []string) (err error) {
	switch args[0] {
	case "cd":
		err = os.Chdir(args[1])
	case "h", "help":
		fmt.Println(messageHelp)
	case "exit", "q", "quit":
		fmt.Println(messageExit)
		os.Exit(0)
	default:
		err = ErrNotFound
	}

	return
}

func main() {
	fmt.Print(messageWelcome + "\n[cmetallo@test " + getDirectory() + "]$ ")

	scanner := bufio.NewScanner(os.Stdin)

scan:
	for scanner.Scan() {
		splited := strings.Split(scanner.Text(), " | ")

		bi := bytes.Buffer{}
		bo := bytes.Buffer{}

		for i := range splited {
			args := strings.Split(splited[i], " ")

			err := internal(args)
			if err != nil && err != ErrNotFound {
				panic(err)
			}

			cmd := exec.Command(args[0], args[1:]...)

			_, err = bi.ReadFrom(&bo)
			if err != nil {
				panic(err)
			}
			bo.Reset()

			cmd.Stdin = &bi

			cmd.Stdout = &bo
			cmd.Stderr = &bo

			err = cmd.Run()
			if err != nil {
				fmt.Print(err.Error() + "\n[cmetallo@test " + getDirectory() + "]$ ")
				continue scan
			}
		}

		fmt.Print(bo.String() + "[cmetallo@test " + getDirectory() + "]$ ")
	}
}
