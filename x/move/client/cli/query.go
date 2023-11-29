package cli

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/initia-labs/initia/x/move/types"
	vmapi "github.com/initia-labs/initiavm/api"
)

func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the move module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	queryCmd.AddCommand(
		GetCmdModule(),
		GetCmdModules(),
		GetCmdResource(),
		GetCmdResources(),
		GetCmdTableEntry(),
		GetCmdTableEntries(),
		GetCmdQueryEntryFunction(),
		GetCmdQueryParams(),
	)
	return queryCmd
}

func GetCmdModule() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "module [module owner] [module name]",
		Short:   "Get published move module info",
		Long:    "Get published move module info",
		Aliases: []string{"m"},
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			_, err = types.AccAddressFromString(args[0])
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Module(
				context.Background(),
				&types.QueryModuleRequest{
					Address:    args[0],
					ModuleName: args[1],
				},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdModules() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "modules [module owner]",
		Short:   "Get all published move module infos of an account",
		Long:    "Get all published move module infos of an account",
		Aliases: []string{"ms"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			_, err = types.AccAddressFromString(args[0])
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Modules(
				context.Background(),
				&types.QueryModulesRequest{
					Address: args[0],
				},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "list modules")
	return cmd
}

func GetCmdResource() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "resource [resource owner] [struct tag]",
		Short:   "Get store raw resource data",
		Long:    "Get store raw resource data",
		Aliases: []string{"r"},
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			_, err = types.AccAddressFromString(args[0])
			if err != nil {
				return err
			}

			_, err = vmapi.ParseStructTag(args[1])
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Resource(
				context.Background(),
				&types.QueryResourceRequest{
					Address:   args[0],
					StructTag: args[1],
				},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdResources() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "resources [resource owner]",
		Short:   "Get all raw resource data of an account",
		Long:    "Get all raw resource data of an account",
		Aliases: []string{"rs"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			_, err = types.AccAddressFromString(args[0])
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Resources(
				context.Background(),
				&types.QueryResourcesRequest{
					Address: args[0],
				},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "list modules")
	return cmd
}

func GetCmdTableEntry() *cobra.Command {
	decoder := newArgDecoder(asciiDecodeString)
	cmd := &cobra.Command{
		Use:     "table-entry [table addr] [key_bytes]",
		Short:   "Get store raw table entry",
		Long:    "Get store raw table entry",
		Aliases: []string{"entry"},
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			_, err = types.AccAddressFromString(args[0])
			if err != nil {
				return err
			}

			keyBz, err := decoder.DecodeString(args[1])
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.TableEntry(
				context.Background(),
				&types.QueryTableEntryRequest{
					Address:  args[0],
					KeyBytes: keyBz,
				},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	decoder.RegisterFlags(cmd.PersistentFlags(), "key bytes")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdTableEntries() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "table-entries [table addr]",
		Short:   "Get all table entries",
		Long:    "Get all table entries",
		Aliases: []string{"entries"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			_, err = types.AccAddressFromString(args[0])
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.TableEntries(
				context.Background(),
				&types.QueryTableEntriesRequest{
					Address: args[0],
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "table entries")

	return cmd
}

func GetCmdQueryEntryFunction() *cobra.Command {
	bech32PrefixAccAddr := sdk.GetConfig().GetBech32AccountAddrPrefix()
	cmd := &cobra.Command{
		Use:   "execute [module owner] [module name] [function name]",
		Short: "Get entry function execution result",
		Long: strings.TrimSpace(
			fmt.Sprintf(`
Get an entry function execution result

Supported types : u8, u16, u32, u64, u128, u256, bool, string, address, raw, vector<inner_type>

Example of args: address:0x1 bool:true u8:0 string:hello vector<u32>:a,b,c,d

Example:
$ %s query move execute \
    %s1lwjmdnks33xwnmfayc64ycprww49n33mtm92ne \
	ManagedCoin \
	get_balance \
	--type-args '0x1::native_uinit::Coin 0x1::native_uusdc::Coin' \
 	--args 'u8:0 address:0x1 string:"hello world"'
`, version.AppName, bech32PrefixAccAddr,
			),
		),
		Aliases: []string{"e"},
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			_, err = types.AccAddressFromString(args[0])
			if err != nil {
				return err
			}

			var typeArgs []string
			flagTypeArgs, err := cmd.Flags().GetString(FlagTypeArgs)
			if err != nil {
				return err
			}
			if flagTypeArgs != "" {
				typeArgs = strings.Split(flagTypeArgs, " ")
			}

			flagArgs, err := cmd.Flags().GetString(FlagArgs)
			if err != nil {
				return err
			}

			argTypes, argStrs := parseArguments(flagArgs)
			if len(argTypes) != len(argStrs) {
				return fmt.Errorf("invalid argument format len(types) != len(args)")
			}

			serializer := NewSerializer()
			bcsArgs := [][]byte{}
			for i := range argTypes {
				bcsArg, err := BcsSerializeArg(argTypes[i], argStrs[i], serializer)
				if err != nil {
					return err
				}

				bcsArgs = append(bcsArgs, bcsArg)
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.ViewFunction(
				context.Background(),
				&types.QueryViewFunctionRequest{
					Address:      args[0],
					ModuleName:   args[1],
					FunctionName: args[2],
					TypeArgs:     typeArgs,
					Args:         bcsArgs,
				},
			)

			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	cmd.Flags().AddFlagSet(FlagSetTypeArgs())
	cmd.Flags().AddFlagSet(FlagSetArgs())
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryParams implements the params query command.
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the current move parameters information",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query values set as move parameters.

Example:
$ %s query move params
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(cmd.Context(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(&res.Params)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
