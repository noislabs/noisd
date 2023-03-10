package types

import (
	"errors"
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys
var (
	KeyDistributionProportions   = []byte("DistributionProportions")
	KeyDeveloperRewardsReceiver  = []byte("DeveloperRewardsReceiver")
	KeyRandomnessRewardsReceiver = []byte("RandomnessRewardsReceiver")
)

// DefaultParams module parameters
func DefaultParams() Params {
	return Params{
		DistributionProportions: DistributionProportions{
			RandomnessRewards: sdk.NewDecWithPrec(4, 2),  // 4%
			ValidatorRewards:  sdk.NewDecWithPrec(6, 2),  // 6%
			DeveloperRewards:  sdk.NewDecWithPrec(20, 2), // 20 %
		},
		WeightedDeveloperRewardsReceivers: []WeightedAddress{},
	}
}

// ParamTable for module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func (p Params) Validate() error {
	return nil
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(
			KeyDistributionProportions,
			&p.DistributionProportions,
			validateDistributionProportions,
		),
		paramtypes.NewParamSetPair(
			KeyDeveloperRewardsReceiver,
			&p.WeightedDeveloperRewardsReceivers,
			validateWeightedDeveloperRewardsReceivers,
		),
		paramtypes.NewParamSetPair(
			KeyRandomnessRewardsReceiver,
			&p.RandomnessRewardsReceiver,
			validateRandomnessRewardsReceiver,
		),
	}
}

func validateRandomnessRewardsReceiver(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	_, err := sdk.AccAddressFromBech32(v)
	if err != nil {
		return fmt.Errorf("invalid address: %s", v)
	}
	return nil
}

func validateDistributionProportions(i interface{}) error {
	v, ok := i.(DistributionProportions)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.ValidatorRewards.IsNegative() {
		return errors.New("validator rewards distribution ratio should not be negative")
	}

	if v.DeveloperRewards.IsNegative() {
		return errors.New("developer rewards distribution ratio should not be negative")
	}

	if v.RandomnessRewards.IsNegative() {
		return errors.New("randomness rewards distribution ratio should not be negative")
	}

	totalProportions := v.ValidatorRewards.Add(v.DeveloperRewards).Add(v.RandomnessRewards)

	if !totalProportions.Equal(sdk.NewDecWithPrec(30, 2)) {
		return errors.New("total distributions ratio should be 30%")
	}

	return nil
}

func validateWeightedDeveloperRewardsReceivers(i interface{}) error {
	v, ok := i.([]WeightedAddress)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// fund community pool when rewards address is empty
	if len(v) == 0 {
		return nil
	}

	weightSum := sdk.NewDec(0)
	for i, w := range v {
		// we allow address to be "" to go to community pool
		if w.Address != "" {
			_, err := sdk.AccAddressFromBech32(w.Address)
			if err != nil {
				return fmt.Errorf("invalid address at %dth", i)
			}
		}
		if !w.Weight.IsPositive() {
			return fmt.Errorf("non-positive weight at %dth", i)
		}
		if w.Weight.GT(sdk.NewDec(1)) {
			return fmt.Errorf("more than 1 weight at %dth", i)
		}
		weightSum = weightSum.Add(w.Weight)
	}

	if !weightSum.Equal(sdk.NewDec(1)) {
		return fmt.Errorf("invalid weight sum: %s", weightSum.String())
	}

	return nil
}
