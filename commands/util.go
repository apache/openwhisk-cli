package commands

import (
	"errors"

	"github.ibm.com/Bluemix/whisk-cli/client"
)

func parseParameters() (client.Parameters, error) {
	parameters := client.Parameters{}
	if len(flags.param)%2 != 0 {
		err := errors.New("--param option must be key-value pairs")
		return parameters, err
	}

	for i := 0; i < len(flags.param); i += 2 {
		keyValue := client.KeyValue{
			Key:   flags.param[i],
			Value: flags.param[i+1],
		}
		parameters = append(parameters, keyValue)

	}
	return parameters, nil
}

func parseAnnotations() (client.Annotations, error) {
	annotations := client.Annotations{}
	if len(flags.param)%2 != 0 {
		err := errors.New("--param option must be key-value pairs")
		return annotations, err
	}

	for i := 0; i < len(flags.param); i += 2 {
		keyValue := client.KeyValue{
			Key:   flags.param[i],
			Value: flags.param[i+1],
		}
		annotations = append(annotations, keyValue)

	}
	return annotations, nil
}
