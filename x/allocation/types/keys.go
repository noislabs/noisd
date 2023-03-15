package types

const (
	// ModuleName defines the module name
	ModuleName = "allocation"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for allocation
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	ValidatorRewardsPool = "validator_rewards_pool"
)

var ValidatorRewardsPrefix = []byte{0x01}
