package commands

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth <key string>",
	Short: "add an authentication key to .wskprops",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			err := errors.New("Invalid auth argument")
			fmt.Println(err)
			return
		}

		authToken := args[0]

		props, err := readProps(propsFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		props["AUTH"] = authToken

		err = writeProps(propsFile, props)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
