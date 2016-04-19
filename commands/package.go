package commands

import (
	"errors"
	"fmt"
	"net/http"

	"github.ibm.com/BlueMix-Fabric/go-whisk/whisk"

	"github.com/fatih/color"
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
		fmt.Println("TODO :: this command has been commented out because it is out of date")

		// var err error
		// if len(args) != 2 {
		// 	err = errors.New("Invalid argument list")
		// 	fmt.Println(err)
		// 	return
		// }
		//
		// bindingArg := args[0]
		// packageName := args[1]
		//
		// parameters, err := parseParameters(flags.common.param)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		//
		// annotations, err := parseAnnotations(flags.common.annotation)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		//
		// parsedBindingArg := strings.Split(bindingArg, ":")
		// bindingName := parsedBindingArg[0]
		// var bindingNamespace string
		// if len(parsedBindingArg) == 1 {
		// 	bindingNamespace = client.Config.Namespace
		// } else if len(parsedBindingArg) == 2 {
		// 	bindingNamespace = parsedBindingArg[1]
		// } else {
		// 	err = fmt.Errorf("Invalid binding argument %s", bindingArg)
		// 	fmt.Println(err)
		// 	return
		// }
		//
		// binding := whisk.Binding{
		// 	Name:      bindingName,
		// 	Namespace: bindingNamespace,
		// }
		//
		// p := &whisk.Package{
		// 	Name:        packageName,
		// 	Publish:     flags.common.shared,
		// 	Annotations: annotations,
		// 	Parameters:  parameters,
		// 	Binding:     binding,
		// }
		// p, _, err = client.Packages.Insert(p, false)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		//
		// printJSON(p)
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

		fmt.Printf("%s created package %s\n", color.GreenString("ok:"), boldString(p.Name))
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

		fmt.Printf("%s updated package %s\n", color.GreenString("ok:"), boldString(p.Name))

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

		xPackage, _, err := client.Packages.Get(packageName)
		if err != nil {
			fmt.Println(err)
			return
		}

		if flags.common.summary {
			fmt.Printf("%s /%s/%s\n", boldString("package"), xPackage.Namespace, xPackage.Name)
		} else {
			fmt.Printf("%s got package %s\n", color.GreenString("ok:"), boldString(packageName))
			printJSON(xPackage)
		}
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

		fmt.Printf("%s deleted package %s\n", color.GreenString("ok:"), boldString(packageName))
	},
}

var packageListCmd = &cobra.Command{
	Use:   "list <namespace string>",
	Short: "list all packages",

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		qName := qualifiedName{}
		if len(args) == 1 {
			qName, err = parseQualifiedName(args[0])
			if err != nil {
				fmt.Printf("error: %s", err)
				return
			}
			ns := qName.namespace
			if len(ns) == 0 {
				err = errors.New("No valid namespace detected.  Make sure that namespace argument is preceded by a \"/\"")
				fmt.Printf("error: %s\n", err)
				return
			}

			client.Namespace = ns
		}

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

		printList(packages)
	},
}

var packageRefreshCmd = &cobra.Command{
	Use:   "refresh <namespace string>",
	Short: "refresh package bindings",

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		if len(args) == 1 {
			namespace := args[0]
			currentNamespace := client.Config.Namespace
			client.Config.Namespace = namespace
			defer func() {
				client.Config.Namespace = currentNamespace
			}()
		}

		updates, resp, err := client.Packages.Refresh()
		if err != nil {
			fmt.Println(err)
			return
		}

		switch resp.StatusCode {
		case http.StatusOK:
			fmt.Printf("\n%s refreshed successfully\n", client.Config.Namespace)

			if len(updates.Added) > 0 {
				fmt.Println("created bindings:")
				printJSON(updates.Added)
			} else {
				fmt.Println("no bindings created")
			}

			if len(updates.Updated) > 0 {
				fmt.Println("updated bindings:")
				printJSON(updates.Updated)
			} else {
				fmt.Println("no bindings updated")
			}

			if len(updates.Deleted) > 0 {
				fmt.Println("deleted bindings:")
				printJSON(updates.Deleted)
			} else {
				fmt.Println("no bindings deleted")
			}

		case http.StatusNotImplemented:
			fmt.Println("error: This feature is not implemented in the targeted deployment")
			return
		default:
			fmt.Println("error: ", resp.Status)
			return
		}

	},
}

func init() {

	packageCreateCmd.Flags().StringSliceVarP(&flags.common.annotation, "annotation", "a", []string{}, "annotations")
	packageCreateCmd.Flags().StringSliceVarP(&flags.common.param, "param", "p", []string{}, "default parameters")
	packageCreateCmd.Flags().StringVarP(&flags.xPackage.serviceGUID, "service_guid", "s", "", "a unique identifier of the service")
	packageCreateCmd.Flags().BoolVar(&flags.common.shared, "shared", false, "shared action (default: private)")

	packageUpdateCmd.Flags().StringSliceVarP(&flags.common.annotation, "annotation", "a", []string{}, "annotations")
	packageUpdateCmd.Flags().StringSliceVarP(&flags.common.param, "param", "p", []string{}, "default parameters")
	packageUpdateCmd.Flags().StringVarP(&flags.xPackage.serviceGUID, "service_guid", "s", "", "a unique identifier of the service")
	packageUpdateCmd.Flags().BoolVar(&flags.common.shared, "shared", false, "shared action (default: private)")

	packageGetCmd.Flags().BoolVarP(&flags.common.summary, "summary", "s", false, "summarize entity details")

	packageBindCmd.Flags().StringSliceVarP(&flags.common.annotation, "annotation", "a", []string{}, "annotations")
	packageBindCmd.Flags().StringSliceVarP(&flags.common.param, "param", "p", []string{}, "default parameters")

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
		packageRefreshCmd,
	)
}
