package root

import (
	"fmt"
	"os"

	"github.com/Dtx-validator/dtx-node/command/backup"
	"github.com/Dtx-validator/dtx-node/command/genesis"
	"github.com/Dtx-validator/dtx-node/command/helper"
	"github.com/Dtx-validator/dtx-node/command/ibft"
	"github.com/Dtx-validator/dtx-node/command/license"
	"github.com/Dtx-validator/dtx-node/command/loadbot"
	"github.com/Dtx-validator/dtx-node/command/monitor"
	"github.com/Dtx-validator/dtx-node/command/peers"
	"github.com/Dtx-validator/dtx-node/command/secrets"
	"github.com/Dtx-validator/dtx-node/command/server"
	"github.com/Dtx-validator/dtx-node/command/status"
	"github.com/Dtx-validator/dtx-node/command/txpool"
	"github.com/Dtx-validator/dtx-node/command/version"
	"github.com/Dtx-validator/dtx-node/command/whitelist"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "The core node software which powers the DTX blockchain",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		monitor.GetCommand(),
		loadbot.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		whitelist.GetCommand(),
		license.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
