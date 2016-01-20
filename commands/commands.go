package commands

import (
	"fmt"
	"net/http"

	"github.com/mitchellh/go-homedir"
	"github.ibm.com/Bluemix/go-whisk/whisk"
)

var client *whisk.Client

// PropsFile is the path to the current props file (default ~/.wskprops).
var PropsFile string

func init() {
	var err error
	PropsFile, err = homedir.Expand(defaultPropsFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	clientConfig := &whisk.Config{}

	props, err := readProps(PropsFile)
	if err != nil {
		return
	}

	if namespace, hasProp := props["NAMESPACE"]; hasProp {
		clientConfig.Namespace = namespace
	}

	if authToken, hasProp := props["AUTH"]; hasProp {
		clientConfig.AuthToken = authToken
	}

	// TODO :: set clientConfig based on environment variables
	// Environment variables override prop file variables

	// Setup client
	client, err = whisk.New(http.DefaultClient, clientConfig)
	if err != nil {
		return
	}

}

func Execute() error {
	return WskCmd.Execute()
}
