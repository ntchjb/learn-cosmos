package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ntchjb/learn-cosmos/x/learncosmos/types"
	"github.com/spf13/cobra"
)

func CmdGetGoldPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   types.QueryGoldPool,
		Short: "Get gold pool stats",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			request := &types.QueryGoldPoolRequest{}
			res, err := queryClient.GoldPool(context.Background(), request)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
