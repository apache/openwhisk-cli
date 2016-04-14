package commands

import (
	"errors"
	"fmt"

	"github.ibm.com/BlueMix-Fabric/go-whisk/whisk"

	"github.com/fatih/color"
	prettyjson "github.com/hokaccha/go-prettyjson"
)

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

var bold = color.New(color.Bold).SprintFunc()
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
	}
}

func printActionList(actions []whisk.Action) {
	boldPrintf("actions\n")
	for _, action := range actions {
		publishState := "private"
		if action.Publish {
			publishState = "public"
		}
		fmt.Printf("/%s/%s%30s\n", action.Namespace, action.Name, publishState)
	}
}

func printTriggerList(triggers []whisk.Trigger) {
	boldPrintf("triggers\n")
	for _, trigger := range triggers {
		publishState := "private"
		if trigger.Publish {
			publishState = "public"
		}
		fmt.Printf("/%s/%s%30s\n", trigger.Namespace, trigger.Name, publishState)
	}
}

func printPackageList(packages []whisk.Package) {
	boldPrintf("packages\n")
	for _, xpackage := range packages {
		publishState := "private"
		if xpackage.Publish {
			publishState = "public"
		}
		fmt.Printf("/%s/%s%30s\n", xpackage.Namespace, xpackage.Name, publishState)
	}
}

func printRuleList(rules []whisk.Rule) {
	boldPrintf("rules\n")
	for _, rule := range rules {
		publishState := "private"
		if rule.Publish {
			publishState = "public"
		}
		fmt.Printf("/%s/%s%30s\n", rule.Namespace, rule.Name, publishState)
	}
}

func printNamespaceList(namespaces []whisk.Namespace) {
	boldPrintf("namespaces\n")
	for _, namespace := range namespaces {
		fmt.Printf("%s\n", namespace.Name)
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
