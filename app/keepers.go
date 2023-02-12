package app

// import (
// 	"github.com/CosmWasm/wasmd/x/wasm"

// 	"github.com/cosmos/cosmos-sdk/baseapp"
// 	"github.com/cosmos/cosmos-sdk/codec"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
// 	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
// 	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
// 	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
// 	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
// 	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
// 	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
// 	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
// 	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
// 	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
// 	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
// 	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
// 	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
// 	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
// 	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
// 	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
// 	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
// 	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
// 	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
// 	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
// 	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
// 	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
// 	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
// 	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
// 	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
// 	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
// 	ibcfeetypes "github.com/cosmos/ibc-go/v4/modules/apps/29-fee/types"
// 	ibctransferkeeper "github.com/cosmos/ibc-go/v4/modules/apps/transfer/keeper"
// 	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
// 	ibchost "github.com/cosmos/ibc-go/v4/modules/core/24-host"
// 	ibckeeper "github.com/cosmos/ibc-go/v4/modules/core/keeper"
// )

// type AppKeepers struct {
// 	// base keepers

// 	AccountKeeper    *authkeeper.AccountKeeper
// 	AuthzKeeper      authzkeeper.Keeper
// 	BankKeeper       bankkeeper.Keeper
// 	CapabilityKeeper *capabilitykeeper.Keeper
// 	CrisisKeeper     crisiskeeper.Keeper
// 	DistrKeeper      distrkeeper.Keeper
// 	EvidenceKeeper   evidencekeeper.Keeper
// 	FeeGrantKeeper   feegrantkeeper.Keeper
// 	GovKeeper        govkeeper.Keeper
// 	MintKeeper       mintkeeper.Keeper
// 	ParamsKeeper     paramskeeper.Keeper
// 	SlashingKeeper   slashingkeeper.Keeper
// 	StakingKeeper    stakingkeeper.Keeper
// 	UpgradeKeeper    upgradekeeper.Keeper

// 	// IBC Keepers
// 	IBCKeeper      *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
// 	TransferKeeper ibctransferkeeper.Keeper
// 	// Wasm
// 	WasmKeeper wasm.Keeper

// 	// keys to access the substores
// 	keys    map[string]*sdk.KVStoreKey
// 	tkeys   map[string]*sdk.TransientStoreKey
// 	memKeys map[string]*sdk.MemoryStoreKey

// 	// Nois keepers
// }

// func KVStoreKeys() []string {
// 	return []string{
// 		authtypes.StoreKey,
// 		authzkeeper.StoreKey,
// 		banktypes.StoreKey,
// 		capabilitytypes.StoreKey,
// 		distrtypes.StoreKey,
// 		evidencetypes.StoreKey,
// 		govtypes.StoreKey,
// 		minttypes.StoreKey,
// 		paramstypes.StoreKey,
// 		slashingtypes.StoreKey,
// 		stakingtypes.StoreKey,
// 		upgradetypes.StoreKey,
// 		ibchost.StoreKey,
// 		ibcfeetypes.StoreKey,
// 		ibctransfertypes.StoreKey,
// 		wasm.StoreKey,
// 	}
// }

// func TransietStoreKeys() []string {
// 	return []string{
// 		paramstypes.StoreKey,
// 	}
// }

// func MemStoreKeys() []string {
// 	return []string{
// 		capabilitytypes.MemStoreKey,
// 	}
// }

// // initParamsKeeper init params keeper and its subspaces
// func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey sdk.StoreKey) paramskeeper.Keeper {
// 	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)
// 	paramsKeeper.Subspace(authtypes.ModuleName)
// 	paramsKeeper.Subspace(banktypes.ModuleName)
// 	paramsKeeper.Subspace(stakingtypes.ModuleName)
// 	paramsKeeper.Subspace(minttypes.ModuleName)
// 	paramsKeeper.Subspace(distrtypes.ModuleName)
// 	paramsKeeper.Subspace(slashingtypes.ModuleName)
// 	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypes.ParamKeyTable())
// 	paramsKeeper.Subspace(crisistypes.ModuleName)
// 	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
// 	paramsKeeper.Subspace(ibchost.ModuleName)
// 	paramsKeeper.Subspace(wasm.ModuleName)
// 	return paramsKeeper
// }

// func (appKeepers *AppKeepers) InitBaseKeepers(
// 	appCodec codec.Codec,
// 	bApp *baseapp.BaseApp,
// 	maccPerms map[string][]string,
// 	wasmDir string,
// 	wasmConfig wasm.Config,
// 	wasmEnabledProposals []wasm.ProposalType,
// 	wasmOpts []wasm.Option,
// 	blockedAddress map[string]bool) {

// 	// account keeper
// 	accountKeeper := authkeeper.NewAccountKeeper(
// 		appCodec,
// 		appKeepers.keys[authtypes.StoreKey],
// 		appKeepers.GetSubspace(authtypes.ModuleName),
// 		authtypes.ProtoBaseAccount,
// 		maccPerms,
// 	)

// }

// // GetSubspace gets existing substore from keeper.
// func (appKeepers *AppKeepers) GetSubspace(moduleName string) paramstypes.Subspace {
// 	subspace, _ := appKeepers.ParamsKeeper.GetSubspace(moduleName)
// 	return subspace
// }
