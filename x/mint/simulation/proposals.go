package simulation

import (
	"math/rand"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/notional-labs/composable/v6/x/mint/types"
)

// Simulation operation weights constants
const (
	DefaultWeightMsgUpdateParams int = 100

	OpWeightMsgUpdateParams = "op_weight_msg_update_params" //nolint:gosec
)

// ProposalMsgs defines the module weighted proposals' contents
func ProposalMsgs() []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			OpWeightMsgUpdateParams,
			DefaultWeightMsgUpdateParams,
			SimulateMsgUpdateParams,
		),
	}
}

// SimulateMsgUpdateParams returns a random MsgUpdateParams
func SimulateMsgUpdateParams(r *rand.Rand, _ sdk.Context, _ []simtypes.Account) sdk.Msg {
	// use the default gov module account address as authority
	var authority sdk.AccAddress = address.Module("gov")

	params := types.DefaultParams()

	// Generate values within safe bounds for uint64
	maxUint32 := uint64(^uint32(0))
	blocksPerYear := uint64(r.Int63n(int64(maxUint32)))
	if blocksPerYear == 0 {
		blocksPerYear = 1
	}
	params.BlocksPerYear = blocksPerYear

	params.GoalBonded = sdkmath.LegacyNewDecWithPrec(int64(simtypes.RandIntBetween(r, 0, 100)), 2)
	params.InflationRateChange = sdkmath.LegacyNewDecWithPrec(int64(simtypes.RandIntBetween(r, 1, 20)), 2)

	// Use safe bounds for token amounts
	maxToken := r.Int63n(100000000000000000)
	if maxToken == 0 {
		maxToken = 1000000000000000
	}
	params.MaxTokenPerYear = sdkmath.NewInt(maxToken)

	minToken := r.Int63n(1000000000000000)
	if minToken == 0 {
		minToken = 1
	}
	params.MinTokenPerYear = sdkmath.NewInt(minToken)
	params.MintDenom = simtypes.RandStringOfLength(r, 10)

	return &types.MsgUpdateParams{
		Authority: authority.String(),
		Params:    params,
	}
}
