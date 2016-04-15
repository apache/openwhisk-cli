package commands

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.ibm.com/BlueMix-Fabric/go-whisk/whisk"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const (
	PollInterval = time.Second * 2
	Delay        = time.Second * 5
)

// activationCmd represents the activation command
var activationCmd = &cobra.Command{
	Use:   "activation",
	Short: "work with activations",
}

var activationListCmd = &cobra.Command{
	Use:   "list",
	Short: "list activations",

	Run: func(cmd *cobra.Command, args []string) {
		options := &whisk.ActivationListOptions{
			Name:  flags.activation.action,
			Limit: flags.common.limit,
			Skip:  flags.common.skip,
			Upto:  flags.activation.upto,
			Since: flags.activation.since,
			Docs:  flags.common.full,
		}
		activations, _, err := client.Activations.List(options)
		if err != nil {
			fmt.Println(err)
			return
		}
		printList(activations)
	},
}

var activationGetCmd = &cobra.Command{
	Use:   "get <id string>",
	Short: "get activation",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			err := errors.New("Invalid ID argument")
			fmt.Println(err)
			return
		}
		id := args[0]
		activation, _, err := client.Activations.Get(id)
		if err != nil {
			fmt.Println(err)
			return
		}

		if flags.common.summary {
			fmt.Printf("activation result for /%s/%s (%s at %s)", activation.Namespace, activation.Name, activation.Response.Status, time.Unix(activation.End/1000, 0))
			printJSON(activation.Response.Result)
		} else {
			fmt.Printf("%s got activation %s\n", color.GreenString("ok:"), boldString(id))
			printJSON(activation)
		}

	},
}

var activationLogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "get the logs of an activation",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			err := errors.New("Invalid ID argument")
			fmt.Println(err)
			return
		}

		id := args[0]
		activation, _, err := client.Activations.Logs(id)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%s got activation %s logs\n", color.GreenString("ok:"), boldString(id))

		printJSON(activation.Logs)
	},
}

var activationResultCmd = &cobra.Command{
	Use:   "result",
	Short: "get the result of an activation",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			err := errors.New("Invalid ID argument")
			fmt.Println(err)
			return
		}

		id := args[0]
		result, _, err := client.Activations.Result(id)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%s got activation %s result\n", color.GreenString("ok:"), boldString(id))
		printJSON(result)
	},
}

var activationPollCmd = &cobra.Command{
	Use:   "poll <namespace string>",
	Short: "poll continuously for log messages from currently running actions",

	Run: func(cmd *cobra.Command, args []string) {
		var name string
		if len(args) == 1 {
			name = args[0]
		}

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		signal.Notify(c, syscall.SIGTERM)
		go func() {
			<-c
			fmt.Println("Poll terminated")
			os.Exit(1)
		}()
		fmt.Println("Enter Ctrl-c to exit.")

		pollSince := time.Now()
		reported := []string{}

		if flags.activation.sinceSeconds+
			flags.activation.sinceMinutes+
			flags.activation.sinceHours+
			flags.activation.sinceDays ==
			0 {
			options := &whisk.ActivationListOptions{
				Limit: 1,
				Docs:  true,
			}
			activationList, _, _ := client.Activations.List(options)
			if len(activationList) > 0 {
				lastActivation := activationList[0]
				pollSince = time.Unix(lastActivation.Start+1, 0).Add(Delay)
			}
		} else {
			t0 := time.Now()

			duration, err := time.ParseDuration(fmt.Sprintf("%ds %dm %dh",
				flags.activation.sinceSeconds,
				flags.activation.sinceMinutes,
				flags.activation.sinceHours+
					flags.activation.sinceDays*24,
			))
			if err == nil {
				pollSince = t0.Add(-duration)
			}
		}

		fmt.Println("Polling for logs")
		localStartTime := time.Now()
		for {
			if flags.activation.exit > 0 {
				localDuration := time.Since(localStartTime)
				if int(localDuration.Seconds()) > flags.activation.exit {
					return
				}
			}

			options := &whisk.ActivationListOptions{
				Name:  name,
				Since: pollSince.Unix(),
				Docs:  true,
			}

			activations, _, err := client.Activations.List(options)
			if err != nil {
				continue
			}

			for _, activation := range activations {
				for _, id := range reported {
					if id == activation.ActivationID {
						continue
					}
				}
				fmt.Printf("\nActivation: %s (%s)\n", activation.Name, activation.ActivationID)
				printJSON(activation.Logs)

				reported = append(reported, activation.ActivationID)
				if activationTime := time.Unix(activation.Start, 0); activationTime.After(pollSince) {
					pollSince = activationTime
				}
			}
			time.Sleep(time.Second * 2)
		}
	},
}

func init() {

	activationListCmd.Flags().StringVarP(&flags.activation.action, "action", "a", "", "retrieve activations for action")
	activationListCmd.Flags().IntVarP(&flags.common.skip, "skip", "s", 0, "skip this many entitites from the head of the collection")
	activationListCmd.Flags().IntVarP(&flags.common.limit, "limit", "l", 30, "only return this many entities from the collection")
	activationListCmd.Flags().BoolVarP(&flags.common.full, "full", "f", false, "include full entity description")
	activationListCmd.Flags().Int64Var(&flags.activation.upto, "upto", 0, "return activations with timestamps earlier than UPTO; measured in miliseconds since Th, 01, Jan 1970")
	activationListCmd.Flags().Int64Var(&flags.activation.since, "since", 0, "return activations with timestamps earlier than UPTO; measured in miliseconds since Th, 01, Jan 1970")

	activationGetCmd.Flags().BoolVarP(&flags.common.summary, "summary", "s", false, "summarize entity details")

	activationPollCmd.Flags().IntVarP(&flags.activation.exit, "exit", "e", 0, "exit after this many seconds")
	activationPollCmd.Flags().IntVar(&flags.activation.sinceSeconds, "since-seconds", 0, "start polling for activations this many seconds ago")
	activationPollCmd.Flags().IntVar(&flags.activation.sinceMinutes, "since-minutes", 0, "start polling for activations this many minutes ago")
	activationPollCmd.Flags().IntVar(&flags.activation.sinceHours, "since-hours", 0, "start polling for activations this many hours ago")
	activationPollCmd.Flags().IntVar(&flags.activation.sinceDays, "since-days", 0, "start polling for activations this many days ago")

	activationCmd.AddCommand(
		activationListCmd,
		activationGetCmd,
		activationLogsCmd,
		activationResultCmd,
		activationPollCmd,
	)
}
