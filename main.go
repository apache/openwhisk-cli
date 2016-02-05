package main

import (
	"fmt"

	"github.ibm.com/BlueMix-Fabric/go-whisk-cli/commands"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Application exited")
		}
	}()

	if err := commands.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}
