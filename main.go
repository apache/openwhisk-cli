package main

import (
	"fmt"

	"github.ibm.com/BlueMix-Fabric/go-whisk-cli/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}
