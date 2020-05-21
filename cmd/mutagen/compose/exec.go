package compose

import (
	"fmt"

	"github.com/spf13/cobra"
)

func execMain(_ *cobra.Command, arguments []string) {
	// Handle top-level help and version flags.
	handleTopLevelHelp()
	handleTopLevelVersion()

	// TODO: Implement.
	fmt.Println("exec not yet implemented")
}

var execCommand = &cobra.Command{
	Use:          "exec",
	Run:          execMain,
	SilenceUsage: true,
}

var execConfiguration struct {
	// help indicates the presence of the -h/--help flag.
	help bool
}

func init() {
	// Avoid Cobra's built-in help functionality that's triggered when the
	// -h/--help flag is present. We still explicitly register a -h/--help flag
	// below for shell completion support.
	execCommand.SetHelpFunc(commandHelp)

	// Grab a handle for the command line flags.
	flags := execCommand.Flags()

	// Wire up exec command flags.
	flags.BoolVarP(&execConfiguration.help, "help", "h", false, "")
	// TODO: Wire up remaining flags.
}