package cmd

import (
	"cosmossdk.io/log"
	confixcmd "cosmossdk.io/tools/confix/cmd"
	"errors"
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmcli "github.com/CosmWasm/wasmd/x/wasm/client/cli"
	"github.com/Fairblock/fairyring/app"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/debug"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/pruning"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/client/snapshot"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
)

func initRootCmd(
	rootCmd *cobra.Command,
	txConfig client.TxConfig,
	basicManager module.BasicManager,
) {
	rootCmd.AddCommand(
		genutilcli.InitCmd(basicManager, app.DefaultNodeHome),
		debug.Cmd(),
		confixcmd.ConfigCommand(),
		pruning.Cmd(newApp, app.DefaultNodeHome),
		snapshot.Cmd(newApp),
	)

	server.AddCommands(rootCmd, app.DefaultNodeHome, newApp, appExport, addModuleInitFlags)

	// add keybase, auxiliary RPC, query, genesis, and tx child commands
	rootCmd.AddCommand(
		server.StatusCommand(),
		genesisCommand(txConfig, basicManager),
		queryCommand(),
		txCommand(),
		keys.Commands(),
	)
	wasmcli.ExtendUnsafeResetAllCmd(rootCmd)

}

func addModuleInitFlags(startCmd *cobra.Command) {
	crisis.AddModuleInitFlags(startCmd)
	wasm.AddModuleInitFlags(startCmd)
}

// genesisCommand builds genesis-related `v050x-chain-sample-ignited genesis` command. Users may provide application specific commands as a parameter
func genesisCommand(txConfig client.TxConfig, basicManager module.BasicManager, cmds ...*cobra.Command) *cobra.Command {
	cmd := genutilcli.Commands(txConfig, basicManager, app.DefaultNodeHome)

	for _, subCmd := range cmds {
		cmd.AddCommand(subCmd)
	}
	return cmd
}

func queryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "query",
		Aliases:                    []string{"q"},
		Short:                      "Querying subcommands",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		rpc.QueryEventForTxCmd(),
		rpc.ValidatorCommand(),
		server.QueryBlockCmd(),
		authcmd.QueryTxsByEventsCmd(),
		server.QueryBlocksCmd(),
		authcmd.QueryTxCmd(),
		server.QueryBlockResultsCmd(),
	)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

func txCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "tx",
		Short:                      "Transactions subcommands",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetSignCommand(),
		authcmd.GetSignBatchCommand(),
		authcmd.GetMultiSignCommand(),
		authcmd.GetMultiSignBatchCmd(),
		authcmd.GetValidateSignaturesCommand(),
		flags.LineBreak,
		authcmd.GetBroadcastCommand(),
		authcmd.GetEncodeCommand(),
		authcmd.GetDecodeCommand(),
		authcmd.GetSimulateCmd(),
	)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

func newApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	appOpts servertypes.AppOptions,
) servertypes.Application {
	baseappOptions := server.DefaultBaseappOptions(appOpts)

	app, err := app.New(
		logger, db, traceStore, true,
		appOpts,
		baseappOptions...,
	)
	if err != nil {
		panic(err)
	}
	return app
}

// newApp creates a new Cosmos SDK app
//func newApp(
//	logger log.Logger,
//	db dbm.DB,
//	traceStore io.Writer,
//	appOpts servertypes.AppOptions,
//) servertypes.Application {
//	var cache storetypes.MultiStorePersistentCache
//
//	if cast.ToBool(appOpts.Get(server.FlagInterBlockCache)) {
//		cache = store.NewCommitKVStoreCacheManager()
//	}
//
//	skipUpgradeHeights := make(map[int64]bool)
//	for _, h := range cast.ToIntSlice(appOpts.Get(server.FlagUnsafeSkipUpgrades)) {
//		skipUpgradeHeights[int64(h)] = true
//	}
//
//	pruningOpts, err := server.GetPruningOptionsFromFlags(appOpts)
//	if err != nil {
//		panic(err)
//	}
//
//	homeDir := cast.ToString(appOpts.Get(flags.FlagHome))
//	chainID := cast.ToString(appOpts.Get(flags.FlagChainID))
//	if chainID == "" {
//		// fallback to genesis chain-id
//		appGenesis, err := tmtypes.GenesisDocFromFile(filepath.Join(homeDir, "config", "genesis.json"))
//		if err != nil {
//			panic(err)
//		}
//
//		chainID = appGenesis.ChainID
//	}
//
//	snapshotDir := filepath.Join(cast.ToString(appOpts.Get(flags.FlagHome)), "data", "snapshots")
//	snapshotDB, err := dbm.NewDB("metadata", dbm.GoLevelDBBackend, snapshotDir)
//	if err != nil {
//		panic(err)
//	}
//	snapshotStore, err := snapshots.NewStore(snapshotDB, snapshotDir)
//	if err != nil {
//		panic(err)
//	}
//
//	snapshotOptions := snapshottypes.NewSnapshotOptions(
//		cast.ToUint64(appOpts.Get(server.FlagStateSyncSnapshotInterval)),
//		cast.ToUint32(appOpts.Get(server.FlagStateSyncSnapshotKeepRecent)),
//	)
//
//	var wasmOpts []wasmkeeper.Option
//	if cast.ToBool(appOpts.Get("telemetry.enabled")) {
//		wasmOpts = append(wasmOpts, wasmkeeper.WithVMCacheMetrics(prometheus.DefaultRegisterer))
//	}
//
//	return app.New(
//		logger,
//		db,
//		traceStore,
//		true,
//		appOpts,
//		wasmOpts,
//		baseapp.SetPruning(pruningOpts),
//		baseapp.SetMinGasPrices(cast.ToString(appOpts.Get(server.FlagMinGasPrices))),
//		baseapp.SetHaltHeight(cast.ToUint64(appOpts.Get(server.FlagHaltHeight))),
//		baseapp.SetHaltTime(cast.ToUint64(appOpts.Get(server.FlagHaltTime))),
//		baseapp.SetMinRetainBlocks(cast.ToUint64(appOpts.Get(server.FlagMinRetainBlocks))),
//		baseapp.SetInterBlockCache(cache),
//		baseapp.SetTrace(cast.ToBool(appOpts.Get(server.FlagTrace))),
//		baseapp.SetIndexEvents(cast.ToStringSlice(appOpts.Get(server.FlagIndexEvents))),
//		baseapp.SetSnapshot(snapshotStore, snapshotOptions),
//		baseapp.SetIAVLCacheSize(cast.ToInt(appOpts.Get(server.FlagIAVLCacheSize))),
//		baseapp.SetIAVLDisableFastNode(cast.ToBool(appOpts.Get(server.FlagDisableIAVLFastNode))),
//		baseapp.SetChainID(chainID),
//	)
//}

func appExport(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	height int64,
	forZeroHeight bool,
	jailAllowedAddrs []string,
	appOpts servertypes.AppOptions,
	modulesToExport []string,
) (servertypes.ExportedApp, error) {
	var (
		bApp *app.App
		err  error
	)

	// this check is necessary as we use the flag in x/upgrade.
	// we can exit more gracefully by checking the flag here.
	homePath, ok := appOpts.Get(flags.FlagHome).(string)
	if !ok || homePath == "" {
		return servertypes.ExportedApp{}, errors.New("application home not set")
	}

	viperAppOpts, ok := appOpts.(*viper.Viper)
	if !ok {
		return servertypes.ExportedApp{}, errors.New("appOpts is not viper.Viper")
	}

	// overwrite the FlagInvCheckPeriod
	viperAppOpts.Set(server.FlagInvCheckPeriod, 1)
	appOpts = viperAppOpts

	if height != -1 {
		bApp, err = app.New(logger, db, traceStore, false, appOpts)
		if err != nil {
			return servertypes.ExportedApp{}, err
		}

		if err := bApp.LoadHeight(height); err != nil {
			return servertypes.ExportedApp{}, err
		}
	} else {
		bApp, err = app.New(logger, db, traceStore, true, appOpts)
		if err != nil {
			return servertypes.ExportedApp{}, err
		}
	}

	return bApp.ExportAppStateAndValidators(forZeroHeight, jailAllowedAddrs, modulesToExport)
}

// appExport creates a new simapp (optionally at a given height)
//func appExport(
//	logger log.Logger,
//	db dbm.DB,
//	traceStore io.Writer,
//	height int64,
//	forZeroHeight bool,
//	jailAllowedAddrs []string,
//	appOpts servertypes.AppOptions,
//	modulesToExport []string,
//) (servertypes.ExportedApp, error) {
//
//	homePath, ok := appOpts.Get(flags.FlagHome).(string)
//	if !ok || homePath == "" {
//		return servertypes.ExportedApp{}, errors.New("application home not set")
//	}
//
//	app := app.New(
//		logger,
//		db,
//		traceStore,
//		height == -1, // -1: no height provided
//		appOpts,
//		app.EmptyWasmOpts,
//	)
//
//	if height != -1 {
//		if err := app.LoadHeight(height); err != nil {
//			return servertypes.ExportedApp{}, err
//		}
//	}
//
//	return app.ExportAppStateAndValidators(forZeroHeight, jailAllowedAddrs, modulesToExport)
//}