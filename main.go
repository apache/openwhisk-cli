package main

import (
	"fmt"
	"os"

	"github.ibm.com/theodore-wilson/whisk-cli/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
