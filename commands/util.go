package commands

import (
	"encoding/json"
	"fmt"
	"strings"

	prettyjson "github.com/hokaccha/go-prettyjson"
	"github.ibm.com/BlueMix-Fabric/go-whisk/whisk"
)

func parseParameters(jsonStr string) (whisk.Parameters, error) {
	parameters := whisk.Parameters{}
	if len(jsonStr) == 0 {
		return parameters, nil
	}
	reader := strings.NewReader(jsonStr)
	err := json.NewDecoder(reader).Decode(&parameters)
	if err != nil {
		return nil, err
	}
	return parameters, nil
}

func parseAnnotations(jsonStr string) (whisk.Annotations, error) {
	annotations := whisk.Annotations{}
	if len(jsonStr) == 0 {
		return annotations, nil
	}
	reader := strings.NewReader(jsonStr)
	err := json.NewDecoder(reader).Decode(&annotations)
	if err != nil {
		return nil, err
	}
	return annotations, nil
}

func logoText() string {

	logo := `

__          ___     _     _
\ \        / / |   (_)   | |
 \ \  /\  / /| |__  _ ___| | __
  \ \/  \/ / | '_ \| / __| |/ /
   \  /\  /  | | | | \__ \   <
    \/  \/   |_| |_|_|___/_|\_\

			`

	return logo
}

func printJSON(v interface{}) {
	output, _ := prettyjson.Marshal(v)
	fmt.Println(string(output))
}
