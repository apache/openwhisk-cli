package commands

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var propertyCmd = &cobra.Command{
	Use:   "property",
	Short: "work with whisk properties",
}

var propertySetCmd = &cobra.Command{
	Use:   "set",
	Short: "set property",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var propertyUnsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "unset property",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var propertyGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get property",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	propertyCmd.AddCommand(
		propertySetCmd,
		propertyUnsetCmd,
		propertyGetCmd,
	)

	// need to set property flags as booleans instead of strings... perhaps with boolApihost...
	propertyGetCmd.Flags().BoolVarP(&flags.property.auth, "auth", "u", false, "authorization key")
	propertyGetCmd.Flags().BoolVar(&flags.property.apihost, "apihost", false, "whisk API host")
	propertyGetCmd.Flags().BoolVar(&flags.property.apiversion, "apiversion", false, "whisk API version")
	propertyGetCmd.Flags().BoolVar(&flags.property.cliversion, "cliversion", false, "whisk CLI version")
	propertyGetCmd.Flags().BoolVar(&flags.property.namespace, "namespace", false, "authorization key")

}

func parseConfigFlags(cmd *cobra.Command, args []string) {

	if flags.global.auth != "" {
		client.Config.AuthToken = flags.global.auth
	}

	if flags.global.namespace != "" {
		client.Config.Namespace = flags.global.namespace
	}

	if flags.global.verbose {
		client.Config.Verbose = flags.global.verbose
	}

	if flags.global.edge != false {
		u, err := url.Parse(edgeHost)
		if err != nil {
			fmt.Println(err)
			return
		}
		client.Config.BaseURL = u
	}

}

// NOTE :: does not return
func readProps(path string) (map[string]string, error) {

	props := map[string]string{}

	// check if props file exists

	file, err := os.Open(path)
	if err != nil {
		// If file does not exist, just return props
		return props, nil
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	props = map[string]string{}
	for _, line := range lines {
		kv := strings.Split(line, "=")
		if len(kv) != 2 {
			// Invalid format; skip
			continue
		}
		props[kv[0]] = kv[1]
	}

	return props, nil

}

func writeProps(path string, props map[string]string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	for key, value := range props {
		line := fmt.Sprintf("%s=%s", strings.ToUpper(key), value)
		_, err = fmt.Fprintln(writer, line)
		if err != nil {
			return err
		}
	}
	return nil
}
