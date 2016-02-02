package commands

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// PropsFile is the path to the current props file (default ~/.wskprops).
var PropsFile string

var Properties struct {
	Auth       string
	APIHost    string
	APIVersion string
	APIBuild   string
	CLIVersion string
	Namespace  string
}

var propertyCmd = &cobra.Command{
	Use:   "property",
	Short: "work with whisk properties",
}

var propertySetCmd = &cobra.Command{
	Use:   "set",
	Short: "set property",
	Run: func(cmd *cobra.Command, args []string) {
		// get current props
		props, err := readProps(PropsFile)
		if err != nil {
			fmt.Println(err)
			return
		}

		// read in each flag, update if necessary

		if flags.global.auth != "" {
			props["AUTH"] = flags.global.auth
		}

		if flags.global.namespace != "" {
			props["NAMESPACE"] = flags.global.namespace
		}

		if flags.global.apihost != "" {
			props["APIHOST"] = flags.global.apihost
		}

		if flags.global.apiversion != "" {
			props["APIVERSION"] = flags.global.apiversion
		}

		err = writeProps(PropsFile, props)
		if err != nil {
			fmt.Println(err)
			return
		}

	},
}

var propertyUnsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "unset property",
	Run: func(cmd *cobra.Command, args []string) {
		props, err := readProps(PropsFile)
		if err != nil {
			fmt.Println(err)
			return
		}

		// read in each flag, update if necessary

		if flags.property.auth {
			delete(props, "AUTH")
		}

		if flags.property.namespace {
			delete(props, "NAMESPACE")
		}

		if flags.property.apihost {
			delete(props, "APIHOST")
		}

		if flags.property.apiversion {
			delete(props, "APIVERSION")
		}

		err = writeProps(PropsFile, props)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

var propertyGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get property",
	Run: func(cmd *cobra.Command, args []string) {

		if flags.property.all || flags.property.auth {
			fmt.Println("whisk auth\t\t", Properties.Auth)
		}

		if flags.property.all || flags.property.apihost {
			fmt.Println("whisk API host")
		}

		if flags.property.all || flags.property.apiversion {
			fmt.Println("whisk API version\t\t", Properties.APIVersion)
		}

		if flags.property.all || flags.property.cliversion {
			fmt.Println("whisk CLI version\t\t", Properties.CLIVersion)
		}

		if flags.property.all || flags.property.namespace {
			fmt.Println("whisk namespace\t\t", Properties.Namespace)
		}

		if flags.property.all || flags.property.apibuild {
			info, _, err := client.Info.Get()
			if err == nil {
				fmt.Println("whisk API build\t\t", info.Build)
			}
		}

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
	propertyGetCmd.Flags().BoolVar(&flags.property.apibuild, "apibuild", false, "whisk API build version")
	propertyGetCmd.Flags().BoolVar(&flags.property.cliversion, "cliversion", false, "whisk CLI version")
	propertyGetCmd.Flags().BoolVar(&flags.property.namespace, "namespace", false, "authorization key")
	propertyGetCmd.Flags().BoolVar(&flags.property.all, "all", false, "all properties")

	propertyUnsetCmd.Flags().BoolVarP(&flags.property.auth, "auth", "u", false, "authorization key")
	propertyUnsetCmd.Flags().BoolVar(&flags.property.apihost, "apihost", false, "whisk API host")
	propertyUnsetCmd.Flags().BoolVar(&flags.property.apiversion, "apiversion", false, "whisk API version")
	propertyUnsetCmd.Flags().BoolVar(&flags.property.namespace, "namespace", false, "authorization key")

	err := loadProperties()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func setDefaultProperties() {
	Properties.Auth = ""
	Properties.Namespace = "_"
	Properties.APIHost = ""
	Properties.APIBuild = "2016-01-26T06:45:38-06:00"
	Properties.APIVersion = "v1"
	Properties.CLIVersion = "2016-01-26T06:45:38-06:00"
}

func loadProperties() error {

	setDefaultProperties()

	PropsFile, err := homedir.Expand(defaultPropsFile)
	if err != nil {
		return err
	}

	props, err := readProps(PropsFile)
	if err != nil {
		return err
	}

	if authToken, hasProp := props["AUTH"]; hasProp {
		Properties.Auth = authToken
		fmt.Println("ok: whisk auth set")
	}

	if apiVersion, hasProp := props["APIVERSION"]; hasProp {
		Properties.APIVersion = apiVersion
		fmt.Println("ok: whisk API version set to ", apiVersion)
	}

	if apiHost, hasProp := props["APIHOST"]; hasProp {
		Properties.APIHost = apiHost
		fmt.Println("ok: whisk API host set to ", apiHost)
	}

	if namespace, hasProp := props["NAMESPACE"]; hasProp {
		Properties.Namespace = namespace
		fmt.Println("ok: whisk namespace set to ", namespace)
	}

	return nil
}

func parseConfigFlags(cmd *cobra.Command, args []string) {

	if flags.global.auth != "" {
		client.Config.AuthToken = flags.global.auth
	}

	if flags.global.namespace != "" {
		client.Config.Namespace = flags.global.namespace
	}

	if flags.global.apiversion != "" {
		client.Config.Version = flags.global.apiversion
	}

	if flags.global.apihost != "" {
		u, err := url.Parse(flags.global.apihost)
		if err == nil {
			client.Config.BaseURL = u
		} else {
			fmt.Println(err)
		}
	}

	if flags.global.verbose {
		client.Config.Verbose = flags.global.verbose
	}

	// TODO :: confirm this is correct
	if flags.global.edge != false {
		u, err := url.Parse(edgeHost)
		if err != nil {
			fmt.Println(err)
			return
		}
		client.Config.BaseURL = u
	}

}

func readProps(path string) (map[string]string, error) {

	props := map[string]string{}

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
