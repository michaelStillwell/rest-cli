package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type App struct {
	url    string
	method string
	body   string
}

func (app App) run() {
	var reqBody io.Reader
	if app.body != "" {
		reqBody = bytes.NewBuffer([]byte(app.body))
	}

	req, err := http.NewRequest(app.method, app.url, reqBody)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("req: %v\n", app.url)

	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: reading body: %s\n", err)
		os.Exit(1)
	}

	ct := strings.Split(res.Header.Get("Content-Type"), ";")[0]
	fmt.Println("body: ", parseBody(body, ct))
}

func (app *App) parseFlags(args []string) {
	url, vals := parseArgs(args)

	app.url = url

	if method, ok := vals["method"]; ok {
		app.method = method
	} else {
		app.method = "GET"
	}
}
