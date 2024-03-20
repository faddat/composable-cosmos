package v4_5_1

import (
	"github.com/notional-labs/composable/v6/app/keepers"
	rateLimitKeeper "github.com/notional-labs/composable/v6/x/ratelimit/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RunForkLogic(ctx sdk.Context, keepers *keepers.AppKeepers) {
	ctx.Logger().Info("Applying v5 upgrade" +
		"Remove Rate Limit",
	)

	RemoveRateLimit(ctx, &keepers.RatelimitKeeper)
}

func RemoveRateLimit(ctx sdk.Context, rlKeeper *rateLimitKeeper.Keeper) {
	// Get all current rate limit
	rateLimits := rlKeeper.GetAllRateLimits(ctx)
	// Remove Rate limit
	for _, rateLimit := range rateLimits {
		err := rlKeeper.RemoveRateLimit(ctx, rateLimit.Path.Denom, rateLimit.Path.ChannelID)
		if err != nil {
			panic(err)
		}
	}
}
