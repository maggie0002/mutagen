package sync

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"google.golang.org/grpc"

	"github.com/mutagen-io/mutagen/cmd"
	"github.com/mutagen-io/mutagen/cmd/mutagen/common"
	"github.com/mutagen-io/mutagen/cmd/mutagen/common/templating"
	"github.com/mutagen-io/mutagen/cmd/mutagen/daemon"

	synchronizationmodels "github.com/mutagen-io/mutagen/pkg/api/models/synchronization"
	"github.com/mutagen-io/mutagen/pkg/grpcutil"
	"github.com/mutagen-io/mutagen/pkg/selection"
	synchronizationsvc "github.com/mutagen-io/mutagen/pkg/service/synchronization"
)

// ListWithSelection is an orchestration convenience method that performs a list
// operation using the provided daemon connection and session selection and then
// prints status information.
func ListWithSelection(
	daemonConnection *grpc.ClientConn,
	selection *selection.Selection,
	long bool,
) error {
	// Load the formatting template (if any has been specified).
	template, err := listConfiguration.TemplateFlags.LoadTemplate()
	if err != nil {
		return fmt.Errorf("unable to load formatting template: %w", err)
	}

	// Determine the listing mode.
	mode := common.SessionDisplayModeList
	if long {
		mode = common.SessionDisplayModeListLong
	}

	// Perform the list operation.
	synchronizationService := synchronizationsvc.NewSynchronizationClient(daemonConnection)
	request := &synchronizationsvc.ListRequest{
		Selection: selection,
	}
	response, err := synchronizationService.List(context.Background(), request)
	if err != nil {
		return grpcutil.PeelAwayRPCErrorLayer(err)
	} else if err = response.EnsureValid(); err != nil {
		return fmt.Errorf("invalid list response received: %w", err)
	}

	// If a template was specified, then use that to format output with public
	// model types, otherwise use custom formatting code.
	if template != nil {
		sessions := synchronizationmodels.ExportSessions(response.SessionStates)
		if err := template.Execute(os.Stdout, sessions); err != nil {
			return fmt.Errorf("unable to execute formatting template: %w", err)
		}
	} else {
		if len(response.SessionStates) > 0 {
			for _, state := range response.SessionStates {
				fmt.Println(cmd.DelimiterLine)
				printSession(state, mode)
			}
			fmt.Println(cmd.DelimiterLine)
		} else {
			fmt.Println(cmd.DelimiterLine)
			fmt.Println("No synchronization sessions found")
			fmt.Println(cmd.DelimiterLine)
		}
	}

	// Success.
	return nil
}

// listMain is the entry point for the list command.
func ListMain(_ *cobra.Command, arguments []string) error {
	// Create session selection specification.
	selection := &selection.Selection{
		All:            len(arguments) == 0 && listConfiguration.labelSelector == "",
		Specifications: arguments,
		LabelSelector:  listConfiguration.labelSelector,
	}
	if err := selection.EnsureValid(); err != nil {
		return fmt.Errorf("invalid session selection specification: %w", err)
	}

	// Connect to the daemon and defer closure of the connection.
	daemonConnection, err := daemon.Connect(true, true)
	if err != nil {
		return fmt.Errorf("unable to connect to daemon: %w", err)
	}
	defer daemonConnection.Close()

	// Perform the list operation and print status information.
	return ListWithSelection(daemonConnection, selection, listConfiguration.long)
}

// listCommand is the list command.
var listCommand = &cobra.Command{
	Use:          "list [<session>...]",
	Short:        "List existing synchronization sessions and their statuses",
	RunE:         ListMain,
	SilenceUsage: true,
}

// listConfiguration stores configuration for the list command.
var listConfiguration struct {
	// help indicates whether or not to show help information and exit.
	help bool
	// long indicates whether or not to use long-format listing.
	long bool
	// labelSelector encodes a label selector to be used in identifying which
	// sessions should be paused.
	labelSelector string
	// TemplateFlags store custom templating behavior.
	templating.TemplateFlags
}

func init() {
	// Grab a handle for the command line flags.
	flags := listCommand.Flags()

	// Disable alphabetical sorting of flags in help output.
	flags.SortFlags = false

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&listConfiguration.help, "help", "h", false, "Show help information")

	// Wire up list flags.
	flags.BoolVarP(&listConfiguration.long, "long", "l", false, "Show detailed session information")
	flags.StringVar(&listConfiguration.labelSelector, "label-selector", "", "List sessions matching the specified label selector")

	// Wire up templating flags.
	listConfiguration.TemplateFlags.Register(flags)
}
