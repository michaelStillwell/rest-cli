package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	html "github.com/yosssi/gohtml"
)

func parseArgs(args []string) (string, map[string]string) {
	res := make(map[string]string)
	var single string
	for _, arg := range args {
		if !strings.HasPrefix(arg, "--") {
			if single == "" {
				single = arg
			}

			continue
		}

		a := strings.Split(arg, "--")[1]
		if a != "" {
			f := strings.Split(a, "=")

			res[f[0]] = f[1]
		}
	}

	return single, res
}

func parseBody(body []byte, contentType string) string {
	var res string

	switch contentType {
	case "application/json":
		var buffer bytes.Buffer
		err := json.Indent(&buffer, body, "", "  ")
		if err != nil {
			fmt.Printf("error formatting json: %s\n", err)
			os.Exit(1)
		}
		res = string(buffer.Bytes())
	case "text/html":
		htmlBody := html.Format(string(body))
		res = htmlBody
	default:
		res = ""
	}

	return res
}
