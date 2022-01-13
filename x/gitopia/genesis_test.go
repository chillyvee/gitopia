package gitopia_test

import (
	"testing"

	keepertest "github.com/gitopia/gitopia/testutil/keeper"
	"github.com/gitopia/gitopia/testutil/sample"
	"github.com/gitopia/gitopia/x/gitopia"
	"github.com/gitopia/gitopia/x/gitopia/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		WhoisList: []types.Whois{
			{
				Creator: sample.AccAddress(),
				Name:    "name-1",
				Address: sample.AccAddress(),
			},
			{
				Creator: sample.AccAddress(),
				Name:    "name-2",
				Address: sample.AccAddress(),
			},
		},
		WhoisCount: 2,

		UserList: []types.User{
			{
				Creator: sample.AccAddress(),
				Id:      0,
			},
			{
				Creator: sample.AccAddress(),
				Id:      1,
			},
		},
		UserCount: 2,

		RepositoryList: []types.Repository{
			{
				Creator: sample.AccAddress(),
				Id:      0,
			},
			{
				Creator: sample.AccAddress(),
				Id:      1,
			},
		},
		RepositoryCount: 2,

		OrganizationList: []types.Organization{
			{
				Creator: sample.AccAddress(),
				Id:      0,
				Address: sample.AccAddress(),
			},
			{
				Creator: sample.AccAddress(),
				Id:      1,
				Address: sample.AccAddress(),
			},
		},
		OrganizationCount: 2,

		IssueList: []types.Issue{
			{
				Creator: sample.AccAddress(),
				Id:      0,
			},
			{
				Creator: sample.AccAddress(),
				Id:      1,
			},
		},
		IssueCount: 2,

		PullRequestList: []types.PullRequest{
			{
				Creator: sample.AccAddress(),
				Id:      0,
			},
			{
				Creator: sample.AccAddress(),
				Id:      1,
			},
		},
		PullRequestCount: 2,

		CommentList: []types.Comment{
			{
				Creator: sample.AccAddress(),
				Id:      0,
			},
			{
				Creator: sample.AccAddress(),
				Id:      1,
			},
		},
		CommentCount: 2,

		ReleaseList: []types.Release{
			{
				Creator: sample.AccAddress(),
				Id:      0,
			},
			{
				Creator: sample.AccAddress(),
				Id:      1,
			},
		},
		ReleaseCount: 2,
		RequestList: []types.Request{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		RequestCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GitopiaKeeper(t)
	gitopia.InitGenesis(ctx, *k, genesisState)
	got := gitopia.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.WhoisList, len(genesisState.WhoisList))
	require.Subset(t, genesisState.WhoisList, got.WhoisList)
	require.Equal(t, genesisState.WhoisCount, got.WhoisCount)

	require.Len(t, got.UserList, len(genesisState.UserList))
	require.Subset(t, genesisState.UserList, got.UserList)
	require.Equal(t, genesisState.UserCount, got.UserCount)

	require.Len(t, got.RepositoryList, len(genesisState.RepositoryList))
	require.Subset(t, genesisState.RepositoryList, got.RepositoryList)
	require.Equal(t, genesisState.RepositoryCount, got.RepositoryCount)

	require.Len(t, got.OrganizationList, len(genesisState.OrganizationList))
	require.Subset(t, genesisState.OrganizationList, got.OrganizationList)
	require.Equal(t, genesisState.OrganizationCount, got.OrganizationCount)

	require.Len(t, got.IssueList, len(genesisState.IssueList))
	require.Subset(t, genesisState.IssueList, got.IssueList)
	require.Equal(t, genesisState.IssueCount, got.IssueCount)

	require.Len(t, got.PullRequestList, len(genesisState.PullRequestList))
	require.Subset(t, genesisState.PullRequestList, got.PullRequestList)
	require.Equal(t, genesisState.PullRequestCount, got.PullRequestCount)

	require.Len(t, got.CommentList, len(genesisState.CommentList))
	require.Subset(t, genesisState.CommentList, got.CommentList)
	require.Equal(t, genesisState.CommentCount, got.CommentCount)

	require.Len(t, got.ReleaseList, len(genesisState.ReleaseList))
	require.Subset(t, genesisState.ReleaseList, got.ReleaseList)
	require.Equal(t, genesisState.ReleaseCount, got.ReleaseCount)
	require.Len(t, got.RequestList, len(genesisState.RequestList))
	require.Subset(t, genesisState.RequestList, got.RequestList)
	require.Equal(t, genesisState.RequestCount, got.RequestCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
