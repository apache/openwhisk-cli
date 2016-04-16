package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.ibm.com/BlueMix-Fabric/go-whisk/whisk"

	"github.com/fatih/color"
	prettyjson "github.com/hokaccha/go-prettyjson"
)

type qualifiedName struct {
	namespace   string
	packageName string
	entityName  string
}

func (qName qualifiedName) String() string {
	output := []string{}
	if len(qName.namespace) > 0 {
		output = append(output, "/", qName.namespace, "/")
	}
	if len(qName.packageName) > 0 {
		output = append(output, qName.packageName, "/")
	}
	output = append(output, qName.entityName)

	return strings.Join(output, "")
}

func parseQualifiedName(name string) (qName qualifiedName, err error) {
	if len(name) == 0 {
		err = errors.New("Invalid name format")
		return
	}
	if name[:1] == "/" {
		name = name[1:]
		i := strings.Index(name, "/")
		if i == -1 {
			qName.namespace = name
			return
		}
		if i == 0 {
			err = errors.New("Invalid name format")
			return
		}

		qName.namespace = name[:i]
		name = name[i+1:]
	}

	i := strings.Index(name, "/")

	if i > 0 {
		qName.packageName = name[:i]
		name = name[i+1:]
	}

	qName.entityName = name

	return

}

func parseKeyValueArray(args []string) ([]whisk.KeyValue, error) {
	parsed := []whisk.KeyValue{}
	if len(args)%2 != 0 {
		err := errors.New("key|value arguments must be submitted in comma-separated pairs")
		return parsed, err
	}

	for i := 0; i < len(args); i += 2 {
		keyValue := whisk.KeyValue{
			Key:   args[i],
			Value: args[i+1],
		}
		parsed = append(parsed, keyValue)

	}
	return parsed, nil
}

func parseParameters(args []string) (whisk.Parameters, error) {
	parameters := whisk.Parameters{}
	parsedArgs, err := parseKeyValueArray(args)
	if err != nil {
		return parameters, err
	}
	parameters = whisk.Parameters(parsedArgs)
	return parameters, nil
}

func parseAnnotations(args []string) (whisk.Annotations, error) {
	annotations := whisk.Annotations{}
	parsedArgs, err := parseKeyValueArray(args)
	if err != nil {
		return annotations, err
	}
	annotations = whisk.Annotations(parsedArgs)
	return annotations, nil
}

var boldString = color.New(color.Bold).SprintFunc()
var boldPrintf = color.New(color.Bold).PrintfFunc()

func printList(collection interface{}) {
	switch collection := collection.(type) {
	case []whisk.Action:
		printActionList(collection)
	case []whisk.Trigger:
		printTriggerList(collection)
	case []whisk.Package:
		printPackageList(collection)
	case []whisk.Rule:
		printRuleList(collection)
	case []whisk.Namespace:
		printNamespaceList(collection)
	case []whisk.Activation:
		printActivationList(collection)
	}
}

func printActionList(actions []whisk.Action) {
	boldPrintf("actions\n")
	for _, action := range actions {
		publishState := "private"
		if action.Publish {
			publishState = "shared"
		}
		fmt.Printf("%-70s%s\n", fmt.Sprintf("/%s/%s", action.Namespace, action.Name), publishState)
	}
}

func printTriggerList(triggers []whisk.Trigger) {
	boldPrintf("triggers\n")
	for _, trigger := range triggers {
		publishState := "private"
		if trigger.Publish {
			publishState = "shared"
		}
		fmt.Printf("%-70s%s\n", fmt.Sprintf("/%s/%s", trigger.Namespace, trigger.Name), publishState)
	}
}

func printPackageList(packages []whisk.Package) {
	boldPrintf("packages\n")
	for _, xPackage := range packages {
		publishState := "private"
		if xPackage.Publish {
			publishState = "shared"
		}
		fmt.Printf("%-70s%s\n", fmt.Sprintf("/%s/%s", xPackage.Namespace, xPackage.Name), publishState)
	}
}

func printRuleList(rules []whisk.Rule) {
	boldPrintf("rules\n")
	for _, rule := range rules {
		publishState := "private"
		if rule.Publish {
			publishState = "shared"
		}
		fmt.Printf("%-70s%s\n", fmt.Sprintf("/%s/%s", rule.Namespace, rule.Name), publishState)
	}
}

func printNamespaceList(namespaces []whisk.Namespace) {
	boldPrintf("namespaces\n")
	for _, namespace := range namespaces {
		fmt.Printf("%s\n", namespace.Name)
	}
}

func printActivationList(activations []whisk.Activation) {
	boldPrintf("activations\n")
	for _, activation := range activations {
		fmt.Printf("%s%20s\n", activation.ActivationID, activation.Name)
	}
}

//
//
//
// func parseParameters(jsonStr string) (whisk.Parameters, error) {
// 	parameters := whisk.Parameters{}
// 	if len(jsonStr) == 0 {
// 		return parameters, nil
// 	}
// 	reader := strings.NewReader(jsonStr)
// 	err := json.NewDecoder(reader).Decode(&parameters)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return parameters, nil
// }
//
// func parseAnnotations(jsonStr string) (whisk.Annotations, error) {
// 	annotations := whisk.Annotations{}
// 	if len(jsonStr) == 0 {
// 		return annotations, nil
// 	}
// 	reader := strings.NewReader(jsonStr)
// 	err := json.NewDecoder(reader).Decode(&annotations)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return annotations, nil
// }

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
