package cosmos

import (
	"github.com/aaely/cosmos/x/cosmos/keeper"
	"github.com/aaely/cosmos/x/cosmos/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func handleMsgCreatePost(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreatePost) (*sdk.Result, error) {
	k.CreatePost(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
