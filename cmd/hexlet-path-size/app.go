package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"code"
)

func Run() {
	fmt.Println("Choose mode:")
	fmt.Println("1) Local")
	fmt.Println("2) Server")
	fmt.Print("> ")

	reader := bufio.NewReader(os.Stdin)
	mode, _ := reader.ReadString('\n')
	mode = strings.TrimSpace(mode)

	switch mode {

	case "1":
		runLocal(reader)

	case "2":
		fmt.Println("Starting server...")
		startServer()

	default:
		fmt.Println("Unknown option")
	}
}

func runLocal(reader *bufio.Reader) {
	fmt.Println("Local mode started")
	fmt.Println("Commands:")
	fmt.Println("  path [/path] [-b|-k|-m|-g]")
	fmt.Println("  exit")

	for {
		fmt.Print("> ")

		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "exit" || line == "quit" {
			fmt.Println("bye 👋")
			break
		}

		fields := strings.Fields(line)

		var path string
		format := "auto"

		for _, f := range fields {
			switch f {
			case "-b":
				format = "b"
			case "-k":
				format = "k"
			case "-m":
				format = "m"
			case "-g":
				format = "g"
			default:
				path = f
			}
		}

		if path == "" {
			fmt.Println("error: empty path")
			continue
		}

		result, err := code.Calculate(path, format)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}

		fmt.Println(result)
	}
}
