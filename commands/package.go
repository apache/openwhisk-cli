package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.ibm.com/Bluemix/whisk-cli/client"

	"github.com/spf13/cobra"
)

var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "work with packages",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("package called")
	},
}

var packageBindCmd = &cobra.Command{
	Use:   "bind <package string> <name string>",
	Short: "bind parameters to the package",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) != 2 {
			err = errors.New("Invalid argument list")
			fmt.Println(err)
			return
		}

		bindingArg := args[0]
		packageName := args[1]

		parameters, err := parseParameters(flags.param)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		annotations, err := parseAnnotations(flags.annotation)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		parsedBindingArg := strings.Split(bindingArg, ":")
		bindingName := parsedBindingArg[0]
		var bindingNamespace string
		if len(parsedBindingArg) == 1 {
			bindingNamespace = whisk.Config.Namespace
		} else if len(parsedBindingArg) == 2 {
			bindingNamespace = parsedBindingArg[1]
		} else {
			err = fmt.Errorf("Invalid binding argument %s", bindingArg)
			fmt.Println(err)
			os.Exit(-1)
		}

		binding := client.Binding{
			Name:      bindingName,
			Namespace: bindingNamespace,
		}

		p := &client.Package{
			Name:        packageName,
			Publish:     flags.shared,
			Annotations: annotations,
			Parameters:  parameters,
			Binding:     binding,
		}
		p, _, err = whisk.Packages.Insert(p, false)
		if err != nil {
			fmt.Println(err)
			return
		}

		printJSON(p)
	},
}

var packageCreateCmd = &cobra.Command{
	Use:   "create <name string>",
	Short: "create a new package",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}

		packageName := args[0]

		parameters, err := parseParameters(flags.param)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		annotations, err := parseAnnotations(flags.annotation)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		p := &client.Package{
			Name:        packageName,
			Publish:     flags.shared,
			Annotations: annotations,
			Parameters:  parameters,
		}
		p, _, err = whisk.Packages.Insert(p, false)
		if err != nil {
			fmt.Println(err)
			return
		}

		printJSON(p)
	},
}

var packageUpdateCmd = &cobra.Command{
	Use:   "update <name string>",
	Short: "update an existing package",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO :: parse annotations
		// TODO ::parse parameters
		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}

		packageName := args[0]

		parameters, err := parseParameters(flags.param)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		annotations, err := parseAnnotations(flags.annotation)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		p := &client.Package{
			Name:        packageName,
			Publish:     flags.shared,
			Annotations: annotations,
			Parameters:  parameters,
		}

		p, _, err = whisk.Packages.Insert(p, true)
		if err != nil {
			fmt.Println(err)
			return
		}

		printJSON(p)
	},
}

var packageGetCmd = &cobra.Command{
	Use:   "get <name string>",
	Short: "get package",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}

		packageName := args[0]

		p, _, err := whisk.Packages.Fetch(packageName)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("ok: got package ", packageName)

		output, _ := json.MarshalIndent(p, "", "    ")
		fmt.Printf("%s", output)
	},
}

var packageDeleteCmd = &cobra.Command{
	Use:   "delete <name string>",
	Short: "delete package",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}

		packageName := args[0]

		_, err = whisk.Packages.Delete(packageName)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("ok: deleted package ", packageName)
	},
}

var packageListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all packages",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		options := &client.PackageListOptions{
			Skip:  flags.skip,
			Limit: flags.limit,
		}

		packages, _, err := whisk.Packages.List(options)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("packages")

		printJSON(packages)
	},
}

func init() {

	packageCreateCmd.Flags().StringSliceVarP(&flags.annotation, "annotation", "a", []string{}, "annotations")
	packageCreateCmd.Flags().StringSliceVarP(&flags.param, "param", "p", []string{}, "default parameters")
	packageCreateCmd.Flags().StringVarP(&flags.serviceGUID, "service_guid", "s", "", "a unique identifier of the service")
	packageCreateCmd.Flags().BoolVar(&flags.shared, "shared", false, "shared action (default: private)")

	packageUpdateCmd.Flags().StringSliceVarP(&flags.annotation, "annotation", "a", []string{}, "annotations")
	packageUpdateCmd.Flags().StringSliceVarP(&flags.param, "param", "p", []string{}, "default parameters")
	packageUpdateCmd.Flags().StringVarP(&flags.serviceGUID, "service_guid", "s", "", "a unique identifier of the service")
	packageUpdateCmd.Flags().BoolVar(&flags.shared, "shared", false, "shared action (default: private)")

	packageBindCmd.Flags().StringSliceVarP(&flags.annotation, "annotation", "a", []string{}, "annotations")
	packageBindCmd.Flags().StringSliceVarP(&flags.param, "param", "p", []string{}, "default parameters")

	packageListCmd.Flags().IntVarP(&flags.skip, "skip", "s", 0, "skip this many entities from the head of the collection")
	packageListCmd.Flags().IntVarP(&flags.limit, "limit", "l", 0, "only return this many entities from the collection")

	packageCmd.AddCommand(
		packageBindCmd,
		packageCreateCmd,
		packageUpdateCmd,
		packageGetCmd,
		packageDeleteCmd,
		packageListCmd,
	)
}
