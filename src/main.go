package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	args := os.Args[1:]
	app := App{}

	if isInputFromPipe() {
		app.body = readPipe(os.Stdin)
	}

	app.parseFlags(args)
	app.run()
}

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func readPipe(r io.Reader) string {
	scanner := bufio.NewScanner(bufio.NewReader(r))
	var res string
	for scanner.Scan() {
		res = scanner.Text()
	}
	return res
}

// TODO: make better logger
func log(msg string) {
	fmt.Printf("[rest] %s\n", msg)
}
