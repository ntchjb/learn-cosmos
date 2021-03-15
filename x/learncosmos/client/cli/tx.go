package cli

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ntchjb/learn-cosmos/x/learncosmos/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(CmdBuyGold())
	cmd.AddCommand(CmdSellGold())
	cmd.AddCommand(CmdTransferGold())

	return cmd
}

func CmdBuyGold() *cobra.Command {
	var ibcChannel string
	var oracleScriptID int64

	cmd := &cobra.Command{
		Use:   "buy-gold [gold_amount]",
		Short: "Buy gold from pool",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsGoldAmount := args[0]
			goldAmount, err := strconv.ParseUint(argsGoldAmount, 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBuyGold(clientCtx.GetFromAddress().String(), goldAmount, ibcChannel, oracleScriptID)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().StringVarP(&ibcChannel, "channel", "", types.DefaultGoldPriceIBCChannel, "Name of IBC channel to be used to request price from oracle")
	cmd.Flags().Int64VarP(&oracleScriptID, "oracle-id", "", 1, "ID of oracle script used for querying gold price")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdSellGold() *cobra.Command {
	var ibcChannel string
	var oracleScriptID int64

	cmd := &cobra.Command{
		Use:   "sell-gold [gold_amount]",
		Short: "sell gold to pool",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsGoldAmount := args[0]
			goldAmount, err := strconv.ParseUint(argsGoldAmount, 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSellGold(clientCtx.GetFromAddress().String(), goldAmount, ibcChannel, oracleScriptID)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().StringVarP(&ibcChannel, "channel", "", types.DefaultGoldPriceIBCChannel, "Name of IBC channel to be used to request price from oracle")
	cmd.Flags().Int64VarP(&oracleScriptID, "oracle-id", "", 1, "ID of oracle script used for querying gold price")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdTransferGold() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-gold [receiver_addr] [gold_amount]",
		Short: "transfer gold to target address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsReceiverAddr := args[0]
			receiverAddr, err := sdk.AccAddressFromBech32(argsReceiverAddr)
			if err != nil {
				return err
			}

			argsGoldAmount := args[1]
			goldAmount, err := strconv.ParseUint(argsGoldAmount, 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTransferGold(clientCtx.GetFromAddress().String(), receiverAddr.String(), goldAmount)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
