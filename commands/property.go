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

var Properties struct {
	Auth       string
	APIHost    string
	APIVersion string
	APIBuild   string
	CLIVersion string
	Namespace  string
	PropsFile  string
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
		props, err := readProps(Properties.PropsFile)
		if err != nil {
			fmt.Println(err)
			return
		}

		// read in each flag, update if necessary

		if auth := flags.global.auth; len(auth) > 0 {
			props["AUTH"] = auth
			fmt.Println("ok: whisk auth set")
		}

		if apiHost := flags.global.apihost; len(apiHost) > 0 {
			props["APIHOST"] = apiHost
			fmt.Println("ok: whisk API host set to ", apiHost)
		}

		if apiVersion := flags.global.apiversion; len(apiVersion) > 0 {
			props["APIVERSION"] = apiVersion
			fmt.Println("ok: whisk API version set to ", apiVersion)
		}

		if namespace := flags.global.namespace; len(namespace) > 0 {

			namespaces, _, err := client.Namespaces.List()
			if err != nil {
				fmt.Println(err)
				return
			}

			var validNamespace bool
			for _, ns := range namespaces {
				if ns.Name == namespace {
					validNamespace = true
				}
			}

			if !validNamespace {
				err = fmt.Errorf("Invalid namespace %s", namespace)
				fmt.Println(err)
				return
			}

			props["NAMESPACE"] = namespace
			fmt.Println("ok: whisk namespace set to ", namespace)
		}

		err = writeProps(Properties.PropsFile, props)
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
		props, err := readProps(Properties.PropsFile)
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

		err = writeProps(Properties.PropsFile, props)
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
			fmt.Println("whisk API host\t\t", Properties.APIHost)
		}

		if flags.property.all || flags.property.apiversion {
			fmt.Println("whisk API version\t", Properties.APIVersion)
		}

		if flags.property.all || flags.property.cliversion {
			fmt.Println("whisk CLI version\t", Properties.CLIVersion)
		}

		if flags.property.all || flags.property.namespace {
			fmt.Println("whisk namespace\t\t", Properties.Namespace)
		}

		if flags.property.all || flags.property.apibuild {
			info, _, err := client.Info.Get()
			if err != nil {
				fmt.Println(err)
			} else {
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

	propertySetCmd.Flags().StringVarP(&flags.global.auth, "auth", "u", "", "authorization key")
	propertySetCmd.Flags().StringVar(&flags.global.apihost, "apihost", "", "whisk API host")
	propertySetCmd.Flags().StringVar(&flags.global.apiversion, "apiversion", "", "whisk API version")
	propertySetCmd.Flags().StringVar(&flags.global.namespace, "namespace", "", "whisk namespace")

	propertyUnsetCmd.Flags().BoolVarP(&flags.property.auth, "auth", "u", false, "authorization key")
	propertyUnsetCmd.Flags().BoolVar(&flags.property.apihost, "apihost", false, "whisk API host")
	propertyUnsetCmd.Flags().BoolVar(&flags.property.apiversion, "apiversion", false, "whisk API version")
	propertyUnsetCmd.Flags().BoolVar(&flags.property.namespace, "namespace", false, "whisk namespace")

}

func setDefaultProperties() {
	Properties.Auth = ""
	Properties.Namespace = "_"
	Properties.APIHost = "https://openwhisk.ng.bluemix.net/api/"
	Properties.APIBuild = ""
	Properties.APIVersion = "v1"
	Properties.CLIVersion = "2016-01-26T06:45:38-06:00"
	Properties.PropsFile = "~/.wskprops"
}

func loadProperties() error {
	var err error

	setDefaultProperties()

	Properties.PropsFile, err = homedir.Expand(Properties.PropsFile)
	if err != nil {
		return err
	}

	props, err := readProps(Properties.PropsFile)
	if err != nil {
		return err
	}

	if authToken, hasProp := props["AUTH"]; hasProp {
		Properties.Auth = authToken
	}

	if authToken := os.Getenv("WHISK_AUTH"); len(authToken) > 0 {
		Properties.Auth = authToken
	}

	if apiVersion, hasProp := props["APIVERSION"]; hasProp {
		Properties.APIVersion = apiVersion
	}

	if apiVersion := os.Getenv("WHISK_APIVERSION"); len(apiVersion) > 0 {
		Properties.APIVersion = apiVersion
	}

	if apiHost, hasProp := props["APIHOST"]; hasProp {
		Properties.APIHost = apiHost
	}

	if apiHost := os.Getenv("WHISK_APIHOST"); len(apiHost) > 0 {
		Properties.APIHost = apiHost
	}

	if namespace, hasProp := props["NAMESPACE"]; hasProp {
		Properties.Namespace = namespace
	}

	if namespace := os.Getenv("WHISK_NAMESPACE"); len(namespace) > 0 {
		Properties.Namespace = namespace
	}

	return nil
}

func parseConfigFlags(cmd *cobra.Command, args []string) {

	if auth := flags.global.auth; len(auth) > 0 {
		Properties.Auth = auth
		client.Config.AuthToken = auth
	}

	if namespace := flags.global.namespace; len(namespace) > 0 {
		Properties.Namespace = namespace
		client.Config.Namespace = namespace
	}

	if apiVersion := flags.global.apiversion; len(apiVersion) > 0 {
		Properties.APIVersion = apiVersion
		client.Config.Version = apiVersion
	}

	if apiHost := flags.global.apihost; len(apiHost) > 0 {
		fmt.Println(apiHost)
		Properties.APIHost = apiHost
		u, err := url.ParseRequestURI(apiHost)
		if err == nil {
			client.Config.BaseURL = u
		} else {
			fmt.Println(err)
			os.Exit(-1)
		}
	}

	if flags.global.verbose {
		client.Config.Verbose = flags.global.verbose
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
