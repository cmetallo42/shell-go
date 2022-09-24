package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	splitedPath := strings.Split(path, "/")

	dir := splitedPath[len(splitedPath)-1]

	fmt.Print("[cmetallo@fedora " + dir + "]$ ")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		splited := strings.Split(scanner.Text(), " | ")

		bi := bytes.Buffer{}
		bo := bytes.Buffer{}

		for i := range splited {
			args := strings.Split(splited[i], " ")

			cmd := exec.Command(args[0], args[1:]...)

			bi.ReadFrom(&bo)
			bo.Reset()

			cmd.Stdin = &bi

			cmd.Stdout = &bo
			cmd.Stderr = &bo

			err := cmd.Run()
			if err != nil {
				panic(err)
			}
		}

		fmt.Print(bo.String())

		fmt.Print("[cmetallo@fedora " + dir + "]$ ")
	}
}
