package main

import (
	"fmt"
	"os"

	"github.ibm.com/theodore-wilson/whisk-cli/cmd"
)

func main() {

	// client := client.NewClient(*http.Client)
	// cmd := cmd.NewCli(client)
	// err := cmd.Run(); ==> cmd.WhiskCmd.Execute()
	if err := cmd.WhiskCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
