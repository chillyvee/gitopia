package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gitopia/gitopia/x/gitopia/types"
	"github.com/gitopia/gitopia/x/gitopia/utils"
)

func (k msgServer) CreateComment(goCtx context.Context, msg *types.MsgCreateComment) (*types.MsgCreateCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var commentIid uint64
	var issue types.Issue
	var pullRequest types.PullRequest

	if msg.CommentType == types.Comment_ISSUE.String() {
		if !k.HasIssue(ctx, msg.ParentId) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("issue Id %d doesn't exist", msg.ParentId))
		}

		issue = k.GetIssue(ctx, msg.ParentId)
		commentIid = issue.CommentsCount + 1
	} else if msg.CommentType == types.Comment_PULLREQUEST.String() {
		if !k.HasPullRequest(ctx, msg.ParentId) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("pullRequest Id %d doesn't exist", msg.ParentId))
		}

		pullRequest = k.GetPullRequest(ctx, msg.ParentId)
		commentIid = pullRequest.CommentsCount + 1
	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("invalid comment type %v", msg.CommentType))
	}

	createdAt := ctx.BlockTime().Unix()

	commentType := types.Comment_Type_value[msg.CommentType]

	var comment = types.Comment{
		Creator:           msg.Creator,
		ParentId:          msg.ParentId,
		CommentIid:        commentIid,
		Body:              msg.Body,
		Attachments:       msg.Attachments,
		DiffHunk:          msg.DiffHunk,
		Path:              msg.Path,
		System:            msg.System,
		AuthorAssociation: msg.AuthorAssociation,
		CreatedAt:         createdAt,
		UpdatedAt:         createdAt,
		CommentType:       types.Comment_Type(commentType),
	}

	id := k.AppendComment(
		ctx,
		comment,
	)

	/* Append Comment in the parent issue/pullRequest */
	if msg.CommentType == types.Comment_ISSUE.String() {
		issue.Comments = append(issue.Comments, id)
		issue.CommentsCount += 1
		k.SetIssue(ctx, issue)
	} else if msg.CommentType == types.Comment_PULLREQUEST.String() {
		pullRequest.Comments = append(pullRequest.Comments, id)
		pullRequest.CommentsCount += 1
		k.SetPullRequest(ctx, pullRequest)
	}

	return &types.MsgCreateCommentResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateComment(goCtx context.Context, msg *types.MsgUpdateComment) (*types.MsgUpdateCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var comment = k.GetComment(ctx, msg.Id)

	comment.Body = msg.Body
	comment.Attachments = msg.Attachments
	comment.UpdatedAt = ctx.BlockTime().Unix()

	// Checks that the element exists
	if !k.HasComment(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetCommentOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetComment(ctx, comment)

	return &types.MsgUpdateCommentResponse{}, nil
}

func (k msgServer) DeleteComment(goCtx context.Context, msg *types.MsgDeleteComment) (*types.MsgDeleteCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasComment(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("comment (%d) doesn't exist", msg.Id))
	}

	if msg.Creator != k.GetCommentOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var comment = k.GetComment(ctx, msg.Id)
	var issue types.Issue
	var pullRequest types.PullRequest

	if comment.CommentType == types.Comment_ISSUE {
		if !k.HasIssue(ctx, comment.ParentId) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("issue Id (%d) doesn't exist", comment.ParentId))
		}

		issue = k.GetIssue(ctx, comment.ParentId)

		if i, exists := utils.IssueCommentExists(issue.Comments, msg.Id); exists {
			issue.Comments = append(issue.Comments[:i], issue.Comments[i+1:]...)
		} else {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("can't find commentId (%d) under issue (%d)", msg.Id, comment.ParentId))
		}

		k.SetIssue(ctx, issue)
	} else if comment.CommentType == types.Comment_PULLREQUEST {
		if !k.HasPullRequest(ctx, comment.ParentId) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("pullRequest Id (%d) doesn't exist", comment.ParentId))
		}

		pullRequest = k.GetPullRequest(ctx, comment.ParentId)

		if i, exists := utils.PullRequestCommentExists(pullRequest.Comments, msg.Id); exists {
			pullRequest.Comments = append(pullRequest.Comments[:i], pullRequest.Comments[i+1:]...)
		} else {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("can't find commentId (%d) under pullRequest (%d)", msg.Id, comment.ParentId))
		}

		k.SetPullRequest(ctx, pullRequest)
	}

	k.RemoveComment(ctx, msg.Id)

	return &types.MsgDeleteCommentResponse{}, nil
}
