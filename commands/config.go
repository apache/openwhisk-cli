package commands

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Hidden: true,
	Use:    "config",
	Short:  "Prints out whisk client configuration",
	Run: func(cmd *cobra.Command, args []string) {
		printJSON(client.Config)
	},
}

var propsCmd = &cobra.Command{
	Hidden: true,
	Use:    "props",
	Short:  "Prints out .wskprops",
	Run: func(cmd *cobra.Command, args []string) {
		props, _ := readProps(PropsFile)
		for key, value := range props {
			fmt.Printf("%s=%s\n", key, value)
		}
	},
}

func parseConfigFlags(cmd *cobra.Command, args []string) {

	if flags.auth != "" {
		client.Config.AuthToken = flags.auth
	}

	if flags.namespace != "" {
		client.Config.Namespace = flags.namespace
	}

	if flags.verbose {
		client.Config.Verbose = flags.verbose
	}

	if flags.edge != false {
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
