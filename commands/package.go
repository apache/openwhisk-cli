package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.ibm.com/BlueMix-Fabric/go-whisk/whisk"

	"github.com/spf13/cobra"
)

var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "work with packages",
}

var packageBindCmd = &cobra.Command{
	Use:   "bind <package string> <name string>",
	Short: "bind parameters to the package",

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) != 2 {
			err = errors.New("Invalid argument list")
			fmt.Println(err)
			return
		}

		bindingArg := args[0]
		packageName := args[1]

		parameters, err := parseParameters(flags.common.param)
		if err != nil {
			fmt.Println(err)
			return
		}

		annotations, err := parseAnnotations(flags.common.annotation)
		if err != nil {
			fmt.Println(err)
			return
		}

		parsedBindingArg := strings.Split(bindingArg, ":")
		bindingName := parsedBindingArg[0]
		var bindingNamespace string
		if len(parsedBindingArg) == 1 {
			bindingNamespace = client.Config.Namespace
		} else if len(parsedBindingArg) == 2 {
			bindingNamespace = parsedBindingArg[1]
		} else {
			err = fmt.Errorf("Invalid binding argument %s", bindingArg)
			fmt.Println(err)
			return
		}

		binding := whisk.Binding{
			Name:      bindingName,
			Namespace: bindingNamespace,
		}

		p := &whisk.Package{
			Name:        packageName,
			Publish:     flags.common.shared,
			Annotations: annotations,
			Parameters:  parameters,
			Binding:     binding,
		}
		p, _, err = client.Packages.Insert(p, false)
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

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}

		packageName := args[0]

		parameters, err := parseParameters(flags.common.param)
		if err != nil {
			fmt.Println(err)
			return
		}

		annotations, err := parseAnnotations(flags.common.annotation)
		if err != nil {
			fmt.Println(err)
			return
		}

		p := &whisk.Package{
			Name:        packageName,
			Publish:     flags.common.shared,
			Annotations: annotations,
			Parameters:  parameters,
		}
		p, _, err = client.Packages.Insert(p, false)
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

		parameters, err := parseParameters(flags.common.param)
		if err != nil {
			fmt.Println(err)
			return
		}

		annotations, err := parseAnnotations(flags.common.annotation)
		if err != nil {
			fmt.Println(err)
			return
		}

		p := &whisk.Package{
			Name:        packageName,
			Publish:     flags.common.shared,
			Annotations: annotations,
			Parameters:  parameters,
		}

		p, _, err = client.Packages.Insert(p, true)
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

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}

		packageName := args[0]

		p, _, err := client.Packages.Get(packageName)
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

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}

		packageName := args[0]

		_, err = client.Packages.Delete(packageName)
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

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		options := &whisk.PackageListOptions{
			Skip:   flags.common.skip,
			Limit:  flags.common.limit,
			Public: flags.common.shared,
			Docs:   flags.common.full,
		}

		packages, _, err := client.Packages.List(options)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("packages")

		printJSON(packages)
	},
}

func init() {

	packageCreateCmd.Flags().StringVarP(&flags.common.annotation, "annotation", "a", "", "annotations")
	packageCreateCmd.Flags().StringVarP(&flags.common.param, "param", "p", "", "default parameters")
	packageCreateCmd.Flags().StringVarP(&flags.xPackage.serviceGUID, "service_guid", "s", "", "a unique identifier of the service")
	packageCreateCmd.Flags().BoolVar(&flags.common.shared, "shared", false, "shared action (default: private)")

	packageUpdateCmd.Flags().StringVarP(&flags.common.annotation, "annotation", "a", "", "annotations")
	packageUpdateCmd.Flags().StringVarP(&flags.common.param, "param", "p", "", "default parameters")
	packageUpdateCmd.Flags().StringVarP(&flags.xPackage.serviceGUID, "service_guid", "s", "", "a unique identifier of the service")
	packageUpdateCmd.Flags().BoolVar(&flags.common.shared, "shared", false, "shared action (default: private)")

	packageBindCmd.Flags().StringVarP(&flags.common.annotation, "annotation", "a", "", "annotations")
	packageBindCmd.Flags().StringVarP(&flags.common.param, "param", "p", "", "default parameters")

	packageListCmd.Flags().BoolVar(&flags.common.shared, "shared", false, "include publicly shared entities in the result")
	packageListCmd.Flags().IntVarP(&flags.common.skip, "skip", "s", 0, "skip this many entities from the head of the collection")
	packageListCmd.Flags().IntVarP(&flags.common.limit, "limit", "l", 0, "only return this many entities from the collection")
	packageListCmd.Flags().BoolVar(&flags.common.full, "full", false, "include full entity description")

	packageCmd.AddCommand(
		packageBindCmd,
		packageCreateCmd,
		packageUpdateCmd,
		packageGetCmd,
		packageDeleteCmd,
		packageListCmd,
	)
}
