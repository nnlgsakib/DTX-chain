package init

import (
	"bytes"
	"fmt"

	"github.com/Dtx-validator/dtx-node/command"

	"github.com/Dtx-validator/dtx-node/command/helper"
	"github.com/Dtx-validator/dtx-node/types"
)

type Results []command.CommandResult

func (r Results) GetOutput() string {
	var buffer bytes.Buffer

	for _, result := range r {
		buffer.WriteString(result.GetOutput())
	}

	return buffer.String()
}

type SecretsInitResult struct {
	Address      types.Address `json:"address"`
	NodeID       string        `json:"node_id"`
	ValidatorKey string        `json:"validator_key"` // Updated field to be exported
}

func (r *SecretsInitResult) GetOutput() string {
	var buffer bytes.Buffer

	vals := make([]string, 0, 3)

	vals = append(
		vals,
		fmt.Sprintf("Public key (address)|%s", r.Address.String()),
	)
	vals = append(
		vals,
		fmt.Sprintf("Private key| 0x%s", r.ValidatorKey),
	)

	// if r.BLSPubkey != "" {
	// 	vals = append(
	// 		vals,
	// 		fmt.Sprintf("BLS Public key|%s", r.BLSPubkey),
	// 	)
	// }

	vals = append(vals, fmt.Sprintf("Node ID|%s", r.NodeID))

	buffer.WriteString("\n[SECRETS INIT]\n")
	buffer.WriteString(helper.FormatKV(vals))
	buffer.WriteString("\n")

	return buffer.String()
}
