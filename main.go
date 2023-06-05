package main

import (
	"mails/cmd"

	_ "modernc.org/sqlite"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
