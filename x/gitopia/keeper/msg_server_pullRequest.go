package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gitopia/gitopia/x/gitopia/types"
)

func assignPullRequestIid(repo types.Repository) (uint64, error) {
	var len = uint64(len(repo.Pulls) + 1)
	return len, nil
}

func (k msgServer) CreatePullRequest(goCtx context.Context, msg *types.MsgCreatePullRequest) (*types.MsgCreatePullRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasRepository(ctx, msg.HeadRepoId) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("headRepositoryId %d doesn't exist", msg.HeadRepoId))
	}

	headRepo := k.GetRepository(ctx, msg.HeadRepoId)

	if _, exists := headRepo.Branches[msg.HeadBranch]; !exists {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("headBranch %v doesn't exist", msg.HeadBranch))
	}

	if !k.HasRepository(ctx, msg.BaseRepoId) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("baseRepositoryId %d doesn't exist", msg.BaseRepoId))
	}

	baseRepo := k.GetRepository(ctx, msg.BaseRepoId)

	if _, exists := baseRepo.Branches[msg.BaseBranch]; !exists {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("baseBranch %v doesn't exist", msg.BaseBranch))
	}

	if !((msg.HeadRepoId == msg.BaseRepoId) && (msg.HeadBranch != msg.BaseBranch)) &&
		!(headRepo.Fork && (headRepo.Parent == msg.BaseRepoId)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "operation not permitted")
	}

	iid, err := assignPullRequestIid(baseRepo)
	if err != nil {
		return nil, err
	}
	createdAt := time.Now().Unix()
	zeroTime := time.Time{}.Unix()

	var pullRequest = types.PullRequest{
		Creator:             msg.Creator,
		Iid:                 iid,
		Title:               msg.Title,
		State:               "Open",
		Description:         msg.Description,
		Locked:              false,
		Draft:               false,
		CreatedAt:           createdAt,
		UpdatedAt:           createdAt,
		ClosedAt:            zeroTime,
		MergedAt:            zeroTime,
		MaintainerCanModify: false,
		HeadBranch:          msg.HeadBranch,
		HeadRepoId:          msg.HeadRepoId,
		BaseBranch:          msg.BaseBranch,
		BaseRepoId:          msg.BaseRepoId,
	}

	id := k.AppendPullRequest(
		ctx,
		pullRequest,
	)

	/* Append Pull Request in the respective Repository */
	baseRepo.Pulls = append(baseRepo.Pulls, id)
	k.SetRepository(ctx, baseRepo)

	return &types.MsgCreatePullRequestResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdatePullRequest(goCtx context.Context, msg *types.MsgUpdatePullRequest) (*types.MsgUpdatePullRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	if !k.HasPullRequest(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetPullRequestOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var pullRequest = k.GetPullRequest(ctx, msg.Id)

	pullRequest.Title = msg.Title
	pullRequest.Description = msg.Description

	k.SetPullRequest(ctx, pullRequest)

	return &types.MsgUpdatePullRequestResponse{}, nil
}

func (k msgServer) DeletePullRequest(goCtx context.Context, msg *types.MsgDeletePullRequest) (*types.MsgDeletePullRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasPullRequest(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetPullRequestOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemovePullRequest(ctx, msg.Id)

	return &types.MsgDeletePullRequestResponse{}, nil
}