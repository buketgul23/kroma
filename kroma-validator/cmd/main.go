package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"

	"github.com/ethereum-optimism/optimism/op-service/cliapp"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	"github.com/kroma-network/kroma/kroma-validator"
	"github.com/kroma-network/kroma/kroma-validator/cmd/balance"
	"github.com/kroma-network/kroma/kroma-validator/flags"
)

var (
	Version = ""
	Meta    = ""
)

func main() {
	oplog.SetupDefaults()

	app := cli.NewApp()
	app.Flags = cliapp.ProtectFlags(flags.Flags)
	app.Version = fmt.Sprintf("%s-%s", Version, Meta)
	app.Name = "kroma-validator"
	app.Usage = "L2 Output Submitter and Challenger Service"
	app.Description = "Service for generating and submitting L2 output checkpoints to the L2OutputOracle contract as an L2 Output Submitter, " + "detecting and correcting invalid L2 outputs as a Challenger to ensure the integrity of the L2 state."
	app.Action = curryMain(Version)
	app.Commands = cli.Commands{
		{
			Name:  "deposit",
			Usage: "Deposit ETH into ValidatorPool to be used as bond",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "amount",
					Usage:    "Amount to deposit into ValidatorPool (in wei)",
					Required: true,
				},
			},
			Action: balance.Deposit,
		},
		{
			Name:  "withdraw",
			Usage: "Withdraw ETH from ValidatorPool",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "amount",
					Usage:    "Amount to withdraw from ValidatorPool (in wei)",
					Required: true,
				},
			},
			Action: balance.Withdraw,
		},
		{
			Name:   "unbond",
			Usage:  "Attempt to unbond in ValidatorPool",
			Action: balance.Unbond,
		},
		{
			Name:  "approve",
			Usage: "Approve the AssetManager to spend governance tokens",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "amount",
					Usage:    "Amount to approve the AssetManager to spend (in wei)",
					Required: true,
				},
			},
			Action: balance.Approve,
		},
		{
			Name:  "delegate",
			Usage: "Attempt to self-delegate governance tokens",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "amount",
					Usage:    "Amount to self-delegate (in wei)",
					Required: true,
				},
			},
			Action: balance.Delegate,
		},
		{
			Name:  "initUndelegate",
			Usage: "Initiate an undelegation of governance tokens",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "amount",
					Usage:    "Amount to undelegate (in wei)",
					Required: true,
				},
			},
			Action: balance.InitUndelegate,
		},
		{
			Name:   "finalizeUndelegate",
			Usage:  "Finalize an undelegation of KROs",
			Action: balance.FinalizeUndelegate,
		},
		{
			Name:  "initClaimValidatorReward",
			Usage: "Initiate a claim of validator rewards",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "amount",
					Usage:    "The amount of rewards to claim (in wei)",
					Required: true,
				},
			},
			Action: balance.InitClaimValidatorReward,
		},
		{
			Name:   "finalizeClaimValidatorReward",
			Usage:  "Finalize a claim of validator rewards",
			Action: balance.FinalizeClaimValidatorReward,
		},
		{
			Name:  "registerValidator",
			Usage: "Register the validator to ValidatorManager",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "amount",
					Usage:    "The amount of assets to self-delegate (in wei)",
					Required: true,
				},
				&cli.Uint64Flag{
					Name:     "commissionRate",
					Usage:    "The commission rate the validator sets (in percentage). Maximum 100.",
					Required: true,
				},
				&cli.Uint64Flag{
					Name:     "commissionMaxChangeRate",
					Usage:    "The maximum changeable commission rate change (in percentage). Maximum 100.",
					Required: true,
				},
			},
			Action: balance.RegisterValidator,
		},
		{
			Name:   "tryUnjail",
			Usage:  "Attempt to unjail the validator",
			Action: balance.TryUnjail,
		},
		{
			Name:  "changeCommissionRate",
			Usage: "Change the commission rate of the validator",
			Flags: []cli.Flag{
				&cli.Uint64Flag{
					Name:     "commissionRate",
					Usage:    "The new commission rate the validator sets (in percentage). Maximum 100.",
					Required: true,
				},
			},
			Action: balance.ChangeCommissionRate,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}

// curryMain transforms the kroma-validator.Main function into an app.Action
// This is done to capture the Version of the validator.
func curryMain(version string) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		return validator.Main(version, ctx)
	}
}
