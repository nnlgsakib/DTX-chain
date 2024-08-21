package whitelist

import (
	"github.com/Dtx-validator/dtx-node/command/whitelist/deployment"
	"github.com/Dtx-validator/dtx-node/command/whitelist/show"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	whitelistCmd := &cobra.Command{
		Use:   "whitelist",
		Short: "Top level command for modifying the DTX-CHAIN whitelists within the config. Only accepts subcommands.",
	}

	registerSubcommands(whitelistCmd)

	return whitelistCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		deployment.GetCommand(),
		show.GetCommand(),
	)
}
