// Package dotenv is an executable that prints all values in the global env
package main

import (
	"fmt"
	"os"

	"taylz.io/env"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println("dotenv version", env.Version)
	} else if env := env.ParseDefault(); len(env) < 1 {
		fmt.Println("dotenv: env is empty")
	} else {
		for k, v := range env {
			fmt.Printf(k + "=" + v + "\n")
		}
	}
}
