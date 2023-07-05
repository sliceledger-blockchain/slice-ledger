package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/sliceledger-blockchain/slice-ledger/op-batcher/batcher"
	"github.com/sliceledger-blockchain/slice-ledger/op-batcher/cmd/doc"
	"github.com/sliceledger-blockchain/slice-ledger/op-batcher/flags"
	oplog "github.com/sliceledger-blockchain/slice-ledger/op-service/log"
	"github.com/ethereum/go-ethereum/log"
)

var (
	Version   = "v0.10.14"
	GitCommit = ""
	GitDate   = ""
)

func main() {
	oplog.SetupDefaults()

	app := cli.NewApp()
	app.Flags = flags.Flags
	app.Version = fmt.Sprintf("%s-%s-%s", Version, GitCommit, GitDate)
	app.Name = "op-batcher"
	app.Usage = "Batch Submitter Service"
	app.Description = "Service for generating and submitting L2 tx batches to L1"
	app.Action = curryMain(Version)
	app.Commands = []*cli.Command{
		{
			Name:        "doc",
			Subcommands: doc.Subcommands,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}

// curryMain transforms the batcher.Main function into an app.Action
// This is done to capture the Version of the batcher.
func curryMain(version string) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		return batcher.Main(version, ctx)
	}
}
