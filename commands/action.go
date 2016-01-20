package commands

import (
	"archive/tar"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	client "github.ibm.com/Bluemix/go-whisk"

	"github.com/spf13/cobra"
)

//////////////
// Commands //
//////////////

var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "work with actions",
}

var actionCreateCmd = &cobra.Command{
	Use:   "create <name string> <artifact string>",
	Short: "create a new action",

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var actionName, artifact string
		if len(args) < 1 || len(args) > 2 {
			err = errors.New("Invalid argument list")
			fmt.Println(err)
			return
		}

		actionName = args[0]

		if len(args) == 2 {
			artifact = args[1]
		}

		exec := client.Exec{}

		parameters, err := parseParameters(flags.param)
		if err != nil {
			fmt.Println(err)
			return
		}

		annotations, err := parseAnnotations(flags.annotation)
		if err != nil {
			fmt.Println(err)
			return
		}

		limits := client.Limits{
			Timeout: flags.timeout,
			Memory:  flags.memory,
		}

		if flags.docker {
			exec.Image = artifact
		} else if flags.copy {
			existingAction, _, err := whisk.Actions.Get(actionName)
			if err != nil {
				fmt.Println(err)
				return
			}
			exec = existingAction.Exec
		} else if flags.pipe {
			currentNamespace := whisk.Config.Namespace
			whisk.Config.Namespace = "whisk.system"
			pipeAction, _, err := whisk.Actions.Get("common/pipe")
			if err != nil {
				fmt.Println(err)
				return
			}
			exec = pipeAction.Exec
			whisk.Config.Namespace = currentNamespace
		} else if artifact != "" {
			if _, err := os.Stat(artifact); err != nil {
				// file does not exist
				fmt.Println(err)
				return
			}

			file, err := ioutil.ReadFile(artifact)
			if err != nil {
				fmt.Println(err)
				return
			}

			exec.Code = string(file)

		}

		if flags.lib != "" {
			file, err := os.Open(flags.lib)
			if err != nil {
				fmt.Println(err)
				return
			}

			var r io.Reader
			switch ext := filepath.Ext(file.Name()); ext {
			case "tar":
				r = tar.NewReader(file)
			case "gzip":
				r, err = gzip.NewReader(file)
			default:
				err = fmt.Errorf("Unrecognized file compression %s", ext)
			}
			if err != nil {
				fmt.Println(err)
				return
			}
			lib, err := ioutil.ReadAll(r)
			if err != nil {
				fmt.Println(err)
				return
			}

			exec.Init = base64.StdEncoding.EncodeToString(lib)

		}

		action := &client.Action{
			Name:        actionName,
			Publish:     flags.shared,
			Exec:        exec,
			Annotations: annotations,
			Parameters:  parameters,
			Limits:      limits,
		}

		action, resp, err := whisk.Actions.Insert(action, false)
		if err != nil {
			fmt.Println(resp.Status)
			return
		}

		fmt.Println("ok: created action")
		printJSON(action)

	},
}

var actionUpdateCmd = &cobra.Command{
	Use:   "update <name string> <artifact string>",
	Short: "update an existing action",

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var actionName, artifact string
		if len(args) < 1 || len(args) > 2 {
			err = errors.New("Invalid argument list")
			fmt.Println(err)
			return
		}

		actionName = args[0]

		if len(args) == 2 {
			artifact = args[1]
		}

		exec := client.Exec{}

		parameters, err := parseParameters(flags.param)
		if err != nil {
			fmt.Println(err)
			return
		}

		annotations, err := parseAnnotations(flags.annotation)
		if err != nil {
			fmt.Println(err)
			return
		}

		limits := client.Limits{
			Timeout: flags.timeout,
			Memory:  flags.memory,
		}

		if flags.docker {
			exec.Image = artifact
		} else if flags.copy {
			existingAction, _, err := whisk.Actions.Get(actionName)
			if err != nil {
				fmt.Println(err)
				return
			}
			exec = existingAction.Exec
		} else if flags.pipe {
			currentNamespace := whisk.Config.Namespace
			whisk.Config.Namespace = "whisk.system"
			pipeAction, _, err := whisk.Actions.Get("common/pipe")
			if err != nil {
				fmt.Println(err)
				return
			}
			exec = pipeAction.Exec
			whisk.Config.Namespace = currentNamespace
		} else if artifact != "" {
			if _, err := os.Stat(artifact); err != nil {
				// file does not exist
				fmt.Println(err)
				return
			}

			file, err := ioutil.ReadFile(artifact)
			if err != nil {
				fmt.Println(err)
				return
			}

			exec.Code = string(file)

		}

		if flags.lib != "" {
			file, err := os.Open(flags.lib)
			if err != nil {
				fmt.Println(err)
				return
			}

			var r io.Reader
			switch ext := filepath.Ext(file.Name()); ext {
			case "tar":
				r = tar.NewReader(file)
			case "gzip":
				r, err = gzip.NewReader(file)
			default:
				err = fmt.Errorf("Unrecognized file compression %s", ext)
			}
			if err != nil {
				fmt.Println(err)
				return
			}
			lib, err := ioutil.ReadAll(r)
			if err != nil {
				fmt.Println(err)
				return
			}

			exec.Init = base64.StdEncoding.EncodeToString(lib)

		}

		action := &client.Action{
			Name:        actionName,
			Publish:     flags.shared,
			Exec:        exec,
			Annotations: annotations,
			Parameters:  parameters,
			Limits:      limits,
		}

		action, _, err = whisk.Actions.Insert(action, true)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("ok: updated action")
		printJSON(action)

	},
}

var actionInvokeCmd = &cobra.Command{
	Use:   "invoke <name string> <payload string>",
	Short: "invoke action",
	Run: func(cmd *cobra.Command, args []string) {

		var err error
		var actionName, payloadArg string
		if len(args) < 1 || len(args) > 2 {
			err = errors.New("Invalid argument list")
			fmt.Println(err)
			return
		}

		actionName = args[0]

		payload := map[string]interface{}{}

		if len(flags.param) > 0 {
			parameters, err := parseParameters(flags.param)
			if err != nil {
				fmt.Printf("error: %s", err)
				return
			}

			for key, value := range parameters {
				payload[key] = value
			}
		}

		if len(args) == 2 {
			payloadArg = args[1]
			reader := strings.NewReader(payloadArg)
			err = json.NewDecoder(reader).Decode(&payload)
			if err != nil {
				payload["payload"] = payloadArg
			}
		}

		printJSON(payload)

		activation, _, err := whisk.Actions.Invoke(actionName, payload, flags.blocking)
		if err != nil {
			fmt.Printf("error: %s", err)
			return
		}

		fmt.Printf("ok: invoked %s with id %s\n", actionName, activation.ActivationID)
		printJSON(activation)
	},
}

var actionGetCmd = &cobra.Command{
	Use:   "get <name string>",
	Short: "get action",

	Run: func(cmd *cobra.Command, args []string) {

		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}

		actionName := args[0]
		action, _, err := whisk.Actions.Get(actionName)
		if err != nil {
			fmt.Printf("error: %s", err)
			return
		}
		// print out response
		fmt.Printf("ok: got action %s\n", actionName)
		printJSON(action)
	},
}

var actionDeleteCmd = &cobra.Command{
	Use:   "delete <name string>",
	Short: "delete action",

	Run: func(cmd *cobra.Command, args []string) {
		actionName := args[0]
		_, err := whisk.Actions.Delete(actionName)
		if err != nil {
			fmt.Printf("error: %s", err)
			return
		}
		// print out response
		fmt.Printf("ok: deleted action %s\n", actionName)
	},
}

var actionListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all actions",

	Run: func(cmd *cobra.Command, args []string) {
		options := &client.ActionListOptions{
			Skip:  flags.skip,
			Limit: flags.limit,
		}

		actions, _, err := whisk.Actions.List(options)
		if err != nil {
			fmt.Printf("error: %s", err)
			return
		}
		fmt.Printf("actions\n")
		for _, action := range actions {
			var publishState string
			if action.Publish {
				publishState = "public"
			} else {
				publishState = "private"
			}

			fmt.Printf("%s\t\t\t\t%s\n", action.Name, publishState)
		}

	},
}

///////////
// Flags //
///////////

func init() {

	actionCreateCmd.Flags().BoolVar(&flags.docker, "docker", false, "treat artifact as docker image path on dockerhub")
	actionCreateCmd.Flags().BoolVar(&flags.copy, "copy", false, "treat artifact as the name of an existing action")
	actionCreateCmd.Flags().BoolVar(&flags.pipe, "pipe", false, "pipe treat artifact as comma separated sequence of actions to invoke")
	actionCreateCmd.Flags().BoolVar(&flags.shared, "shared", false, "add library to artifact (must be a gzipped tar file)")
	actionCreateCmd.Flags().StringVar(&flags.lib, "lib", "", "add library to artifact (must be a gzipped tar file)")
	actionCreateCmd.Flags().StringVar(&flags.xPackage, "package", "", "package")

	actionUpdateCmd.Flags().BoolVar(&flags.docker, "docker", false, "treat artifact as docker image path on dockerhub")
	actionUpdateCmd.Flags().BoolVar(&flags.copy, "copy", false, "treat artifact as the name of an existing action")
	actionUpdateCmd.Flags().BoolVar(&flags.pipe, "pipe", false, "pipe treat artifact as comma separated sequence of actions to invoke")
	actionUpdateCmd.Flags().BoolVar(&flags.shared, "shared", false, "add library to artifact (must be a gzipped tar file)")
	actionUpdateCmd.Flags().StringVar(&flags.lib, "lib", "", "add library to artifact (must be a gzipped tar file)")
	actionUpdateCmd.Flags().StringVar(&flags.xPackage, "package", "", "package")

	actionInvokeCmd.Flags().StringVarP(&flags.param, "param", "p", "", "parameters")
	actionInvokeCmd.Flags().BoolVarP(&flags.blocking, "blocking", "b", false, "blocking invoke")

	actionListCmd.Flags().IntVarP(&flags.skip, "skip", "s", 0, "skip this many entitites from the head of the collection")
	actionListCmd.Flags().IntVarP(&flags.limit, "limit", "l", 30, "only return this many entities from the collection")
	actionListCmd.Flags().BoolVar(&flags.full, "full", false, "include full entity description")

	actionCmd.AddCommand(
		actionCreateCmd,
		actionUpdateCmd,
		actionInvokeCmd,
		actionGetCmd,
		actionDeleteCmd,
		actionListCmd,
	)
}
