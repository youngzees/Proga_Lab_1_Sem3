package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"dbms/dbms"
)

func main() {
	fileFlag := flag.String("file", "file.data", "data file")
	queryFlag := flag.String("query", "", "single query to execute")
	flag.Parse()

	db := dbms.NewDBMS(*fileFlag)
	defer db.Save()

	if *queryFlag != "" {
		db.Execute(*queryFlag)
		return
	}

	fmt.Println("DBMS interactive mode. Type 'exit' to quit.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "exit" || line == "quit" {
			break
		}
		if line == "" {
			continue
		}
		db.Execute(line)
	}
}
