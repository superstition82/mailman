package main

import "probemail/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
