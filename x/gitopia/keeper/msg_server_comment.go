package keeper

import (
	"context"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gitopia/gitopia/v2/x/gitopia/types"
)

func (k msgServer) CreateComment(goCtx context.Context, msg *types.MsgCreateComment) (*types.MsgCreateCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, creatorFound := k.GetUser(ctx, msg.Creator)
	if !creatorFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("creator (%v) doesn't exist", msg.Creator))
	}

	var commentIid uint64
	var issue types.Issue
	var pullRequest types.PullRequest
	var found bool
	commentType := types.CommentTypeReply

	if msg.Parent == types.CommentParentIssue {
		issue, found = k.GetRepositoryIssue(ctx, msg.RepositoryId, msg.ParentIid)
		if !found {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("issue (%d) doesn't exist in repository", msg.ParentIid))
		}
		commentIid = issue.CommentsCount + 1
	} else if msg.Parent == types.CommentParentPullRequest {
		pullRequest, found = k.GetRepositoryPullRequest(ctx, msg.RepositoryId, msg.ParentIid)
		if !found {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("pullRequest (%d) doesn't exist in repository", msg.ParentIid))
		}
		commentIid = pullRequest.CommentsCount + 1
		if len(msg.DiffHunk) > 0 {
			commentType = types.CommentTypeReview
		}
	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("invalid parent type %v", msg.Parent))
	}

	var comment = types.Comment{
		Creator:      msg.Creator,
		RepositoryId: msg.RepositoryId,
		ParentIid:    msg.ParentIid,
		Parent:       msg.Parent,
		CommentIid:   commentIid,
		Body:         msg.Body,
		Attachments:  msg.Attachments,
		DiffHunk:     msg.DiffHunk,
		Path:         msg.Path,
		Position:     msg.Position,
		CreatedAt:    ctx.BlockTime().Unix(),
		UpdatedAt:    ctx.BlockTime().Unix(),
		CommentType:  commentType,
	}

	id := k.AppendComment(
		ctx,
		comment,
	)

	/* Increment comment count in the parent issue/pullRequest */
	var parentId uint64
	if comment.Parent == types.CommentParentIssue {
		issue.CommentsCount += 1
		k.SetIssue(ctx, issue)

		parentId = issue.Id
	} else if comment.Parent == types.CommentParentPullRequest {
		pullRequest.CommentsCount += 1
		k.SetPullRequest(ctx, pullRequest)

		parentId = pullRequest.Id
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.CreateCommentEventKey),
			sdk.NewAttribute(types.EventAttributeCreatorKey, msg.Creator),
			sdk.NewAttribute(types.EventAttributeCommentIdKey, strconv.FormatUint(id, 10)),
			sdk.NewAttribute(types.EventAttributeCommentIidKey, strconv.FormatUint(comment.CommentIid, 10)),
			sdk.NewAttribute(types.EventAttributeCommentBodyKey, comment.Body),
			sdk.NewAttribute(types.EventAttributeRepoIdKey, strconv.FormatUint(comment.RepositoryId, 10)),
			sdk.NewAttribute(types.EventAttributeCommentParentIdKey, strconv.FormatUint(parentId, 10)),
			sdk.NewAttribute(types.EventAttributeCommentParentIidKey, strconv.FormatUint(comment.ParentIid, 10)),
			sdk.NewAttribute(types.EventAttributeCommentParentKey, comment.Parent.String()),
			sdk.NewAttribute(types.EventAttributeCommentDiffHunkKey, comment.DiffHunk),
			sdk.NewAttribute(types.EventAttributeCommentPathKey, comment.Path),
			sdk.NewAttribute(types.EventAttributeCommentPositionKey, strconv.FormatUint(comment.Position, 10)),
			sdk.NewAttribute(types.EventAttributeCommentTypeKey, comment.CommentType.String()),
			sdk.NewAttribute(types.EventAttributeCreatedAtKey, strconv.FormatInt(comment.CreatedAt, 10)),
			sdk.NewAttribute(types.EventAttributeUpdatedAtKey, strconv.FormatInt(comment.UpdatedAt, 10)),
		),
	)

	return &types.MsgCreateCommentResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateComment(goCtx context.Context, msg *types.MsgUpdateComment) (*types.MsgUpdateCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetUser(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("creator (%v) doesn't exist", msg.Creator))
	}

	var comment types.Comment
	var parentId uint64

	switch msg.Parent {
	case types.CommentParentIssue:
		comment, found = k.GetIssueComment(ctx, msg.RepositoryId, msg.ParentIid, msg.CommentIid)
		if !found {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("comment  (%d) doesn't exist", msg.CommentIid))
		}

		issue, _ := k.GetRepositoryIssue(ctx, msg.RepositoryId, msg.ParentIid)
		parentId = issue.Id
	case types.CommentParentPullRequest:
		comment, found = k.GetPullRequestComment(ctx, msg.RepositoryId, msg.ParentIid, msg.CommentIid)
		if !found {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("comment  (%d) doesn't exist", msg.CommentIid))
		}

		pullRequest, _ := k.GetRepositoryPullRequest(ctx, msg.RepositoryId, msg.ParentIid)
		parentId = pullRequest.Id
	default:
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("invalid comment parent (%v)", msg.Parent))
	}

	if msg.Creator != comment.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	comment.Body = msg.Body
	comment.Attachments = msg.Attachments
	comment.UpdatedAt = ctx.BlockTime().Unix()

	k.SetComment(ctx, comment)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.UpdateCommentEventKey),
			sdk.NewAttribute(types.EventAttributeCreatorKey, msg.Creator),
			sdk.NewAttribute(types.EventAttributeCommentIdKey, strconv.FormatUint(comment.Id, 10)),
			sdk.NewAttribute(types.EventAttributeCommentIidKey, strconv.FormatUint(comment.CommentIid, 10)),
			sdk.NewAttribute(types.EventAttributeCommentBodyKey, comment.Body),
			sdk.NewAttribute(types.EventAttributeRepoIdKey, strconv.FormatUint(comment.RepositoryId, 10)),
			sdk.NewAttribute(types.EventAttributeCommentParentIdKey, strconv.FormatUint(parentId, 10)),
			sdk.NewAttribute(types.EventAttributeCommentParentIidKey, strconv.FormatUint(comment.ParentIid, 10)),
			sdk.NewAttribute(types.EventAttributeCommentParentKey, comment.Parent.String()),
			sdk.NewAttribute(types.EventAttributeCommentTypeKey, comment.CommentType.String()),
			sdk.NewAttribute(types.EventAttributeUpdatedAtKey, strconv.FormatInt(comment.UpdatedAt, 10)),
		),
	)

	return &types.MsgUpdateCommentResponse{}, nil
}

func (k msgServer) DeleteComment(goCtx context.Context, msg *types.MsgDeleteComment) (*types.MsgDeleteCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetUser(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("creator (%v) doesn't exist", msg.Creator))
	}

	var comment types.Comment
	var parentId uint64

	switch msg.Parent {
	case types.CommentParentIssue:
		comment, found = k.GetIssueComment(ctx, msg.RepositoryId, msg.ParentIid, msg.CommentIid)
		if !found {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("comment  (%d) doesn't exist", msg.CommentIid))
		}

		issue, _ := k.GetRepositoryIssue(ctx, msg.RepositoryId, msg.ParentIid)
		parentId = issue.Id
	case types.CommentParentPullRequest:
		comment, found = k.GetPullRequestComment(ctx, msg.RepositoryId, msg.ParentIid, msg.CommentIid)
		if !found {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("comment  (%d) doesn't exist", msg.CommentIid))
		}

		pullRequest, _ := k.GetRepositoryPullRequest(ctx, msg.RepositoryId, msg.ParentIid)
		parentId = pullRequest.Id
	default:
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("invalid comment parent (%v)", msg.Parent))
	}

	if msg.Creator != comment.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	switch msg.Parent {
	case types.CommentParentIssue:
		k.RemoveIssueComment(ctx, comment.RepositoryId, comment.ParentIid, comment.CommentIid)
	case types.CommentParentPullRequest:
		k.RemovePullRequestComment(ctx, comment.RepositoryId, comment.ParentIid, comment.CommentIid)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.DeleteCommentEventKey),
			sdk.NewAttribute(types.EventAttributeCreatorKey, msg.Creator),
			sdk.NewAttribute(types.EventAttributeCommentIdKey, strconv.FormatUint(comment.Id, 10)),
			sdk.NewAttribute(types.EventAttributeCommentIidKey, strconv.FormatUint(comment.CommentIid, 10)),
			sdk.NewAttribute(types.EventAttributeRepoIdKey, strconv.FormatUint(comment.RepositoryId, 10)),
			sdk.NewAttribute(types.EventAttributeCommentParentIdKey, strconv.FormatUint(parentId, 10)),
			sdk.NewAttribute(types.EventAttributeCommentParentIidKey, strconv.FormatUint(comment.ParentIid, 10)),
			sdk.NewAttribute(types.EventAttributeCommentParentKey, comment.Parent.String()),
			sdk.NewAttribute(types.EventAttributeCommentTypeKey, comment.CommentType.String()),
		),
	)

	return &types.MsgDeleteCommentResponse{}, nil
}
