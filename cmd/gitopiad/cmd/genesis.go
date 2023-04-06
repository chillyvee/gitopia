package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1types "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/cosmos/cosmos-sdk/x/group"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibctypes "github.com/cosmos/ibc-go/v5/modules/core/types"
	"github.com/gitopia/gitopia/app/params"
	v2 "github.com/gitopia/gitopia/x/gitopia/migrations/v2"
	v3 "github.com/gitopia/gitopia/x/gitopia/migrations/v3"
	gitopiatypes "github.com/gitopia/gitopia/x/gitopia/types"
	rewardstypes "github.com/gitopia/gitopia/x/rewards/types"
	"github.com/spf13/cobra"
	tmjson "github.com/tendermint/tendermint/libs/json"
	tmtypes "github.com/tendermint/tendermint/types"
)

const (
	flagGenesisTime   = "genesis-time"
	flagInitialHeight = "initial-height"
)

const (
	rewardsServiceAddress = "gitopia1rrad3vleav3svu7tutqp9sqqv9mh4gex62vjvm"
)

var commentMap map[uint64]v2.Comment

// migrateTestnetState performs type migrations from v1.2.0 to v1.4.0. The
// migration includes:
//
// - Removed `issues` and `pullRequests` from Repository.
// - Add `repositoryId` in Comment and Bounty.
// - Modified comment structure - Parent: issue and pull; various comment types like label, assignees etc; reactions; replies; resolved/unresolved
func migrateTestnetState(state v2.GenesisState) (v3.GenesisState, error) {
	var gitopiaV3Genesis v3.GenesisState

	gitopiaV3Genesis.Params = v3.Params{
		NextInflationTime: time.Now().AddDate(2, 0, 0),
		PoolProportions: v3.PoolProportions{
			Ecosystem: &v3.DistributionProportion{Proportion: 30},
			Team:      &v3.DistributionProportion{Proportion: 28},
		},
		TeamProportions: []v3.DistributionProportion{
			{Proportion: 50, Address: "gitopia1k9pvyj845y9a4m4vuxx8sjq5q28yxym520fh2x"},
			{Proportion: 35, Address: "gitopia1njn3grh5ar4ccapyp4uehuq28wpk2sk5heu7ac"},
			{Proportion: 15, Address: "gitopia1d5r0ql0pg5d8xfs5t0pmn7dl72m2zj2wchkfq3"},
		},
	}

	for _, oldTask := range state.TaskList {
		task := v3.Task{
			Id:       oldTask.Id,
			Type:     v3.TaskType(oldTask.Type),
			State:    v3.TaskState(oldTask.State),
			Message:  oldTask.Message,
			Creator:  oldTask.Creator,
			Provider: oldTask.Provider,
		}

		gitopiaV3Genesis.TaskList = append(gitopiaV3Genesis.TaskList, task)
	}

	gitopiaV3Genesis.TaskCount = state.TaskCount

	for _, oldBranch := range state.BranchList {
		branch := v3.Branch{
			Id:             oldBranch.Id,
			RepositoryId:   oldBranch.RepositoryId,
			Name:           oldBranch.Name,
			Sha:            oldBranch.Sha,
			AllowForcePush: oldBranch.AllowForcePush,
			CreatedAt:      oldBranch.CreatedAt,
			UpdatedAt:      oldBranch.UpdatedAt,
		}

		gitopiaV3Genesis.BranchList = append(gitopiaV3Genesis.BranchList, branch)
	}

	gitopiaV3Genesis.BranchCount = state.BranchCount

	for _, oldTag := range state.TagList {
		tag := v3.Tag{
			Id:           oldTag.Id,
			RepositoryId: oldTag.RepositoryId,
			Name:         oldTag.Name,
			Sha:          oldTag.Sha,
			CreatedAt:    oldTag.CreatedAt,
			UpdatedAt:    oldTag.UpdatedAt,
		}

		gitopiaV3Genesis.TagList = append(gitopiaV3Genesis.TagList, tag)
	}

	gitopiaV3Genesis.TagCount = state.TagCount

	for _, oldMember := range state.MemberList {
		member := v3.Member{
			Id:         oldMember.Id,
			Address:    oldMember.Address,
			DaoAddress: oldMember.DaoAddress,
			Role:       v3.MemberRole(oldMember.Role),
		}

		gitopiaV3Genesis.MemberList = append(gitopiaV3Genesis.MemberList, member)
	}

	gitopiaV3Genesis.MemberCount = state.MemberCount

	for _, oldRelease := range state.ReleaseList {
		var attachments []*v3.Attachment
		for _, oldAttachment := range oldRelease.Attachments {
			attachments = append(attachments, &v3.Attachment{
				Name:     oldAttachment.Name,
				Size_:    oldAttachment.Size_,
				Sha:      oldAttachment.Sha,
				Uploader: oldAttachment.Uploader,
			})
		}

		release := v3.Release{
			Creator:      oldRelease.Creator,
			Id:           oldRelease.Id,
			RepositoryId: oldRelease.RepositoryId,
			TagName:      oldRelease.TagName,
			Target:       oldRelease.Target,
			Name:         oldRelease.Name,
			Description:  oldRelease.Description,
			Attachments:  attachments,
			Draft:        oldRelease.Draft,
			PreRelease:   oldRelease.PreRelease,
			IsTag:        oldRelease.IsTag,
			CreatedAt:    oldRelease.CreatedAt,
			UpdatedAt:    oldRelease.UpdatedAt,
			PublishedAt:  oldRelease.PublishedAt,
		}

		gitopiaV3Genesis.ReleaseList = append(gitopiaV3Genesis.ReleaseList, release)
	}

	gitopiaV3Genesis.ReleaseCount = state.ReleaseCount

	for _, oldDao := range state.DaoList {
		dao := v3.Dao{
			Creator:     oldDao.Creator,
			Id:          oldDao.Id,
			Address:     oldDao.Address,
			Name:        oldDao.Name,
			AvatarUrl:   oldDao.AvatarUrl,
			Followers:   oldDao.Followers,
			Following:   oldDao.Following,
			Teams:       oldDao.Teams,
			Location:    oldDao.Location,
			Website:     oldDao.Website,
			Verified:    oldDao.Verified,
			Description: oldDao.Description,
			CreatedAt:   oldDao.CreatedAt,
			UpdatedAt:   oldDao.UpdatedAt,
		}

		gitopiaV3Genesis.DaoList = append(gitopiaV3Genesis.DaoList, dao)
	}

	gitopiaV3Genesis.DaoCount = state.DaoCount

	for _, oldUser := range state.UserList {
		user := v3.User{
			Creator:        oldUser.Creator,
			Id:             oldUser.Id,
			Name:           oldUser.Name,
			Username:       oldUser.Username,
			UsernameGithub: oldUser.UsernameGithub,
			AvatarUrl:      oldUser.AvatarUrl,
			Followers:      oldUser.Followers,
			Following:      oldUser.Following,
			StarredRepos:   oldUser.StarredRepos,
			Subscriptions:  oldUser.Subscriptions,
			Bio:            oldUser.Bio,
			CreatedAt:      oldUser.CreatedAt,
			UpdatedAt:      oldUser.UpdatedAt,
			Verified:       oldUser.Verified,
		}

		gitopiaV3Genesis.UserList = append(gitopiaV3Genesis.UserList, user)
	}

	gitopiaV3Genesis.UserCount = state.UserCount

	for _, oldUserDao := range state.UserDaoList {
		userDao := v3.UserDao{
			UserAddress: oldUserDao.UserAddress,
			DaoAddress:  oldUserDao.DaoAddress,
		}

		gitopiaV3Genesis.UserDaoList = append(gitopiaV3Genesis.UserDaoList, userDao)
	}

	for _, oldWhois := range state.WhoisList {
		whois := v3.Whois{
			Creator:   oldWhois.Creator,
			Id:        oldWhois.Id,
			Name:      oldWhois.Name,
			Address:   oldWhois.Address,
			OwnerType: v3.OwnerType(oldWhois.OwnerType),
		}

		gitopiaV3Genesis.WhoisList = append(gitopiaV3Genesis.WhoisList, whois)
	}

	gitopiaV3Genesis.WhoisCount = state.WhoisCount

	for _, oldbaseRepositoryKey := range state.BaseRepositoryKeyList {
		baseRepositoryKey := v3.BaseRepositoryKey{
			Id:      oldbaseRepositoryKey.Id,
			Address: oldbaseRepositoryKey.Address,
			Name:    oldbaseRepositoryKey.Name,
		}

		gitopiaV3Genesis.BaseRepositoryKeyList = append(gitopiaV3Genesis.BaseRepositoryKeyList, baseRepositoryKey)
	}

	for _, oldRepository := range state.RepositoryList {
		var labels []*v3.RepositoryLabel
		var releases []*v3.RepositoryRelease
		var collaborators []*v3.RepositoryCollaborator
		var backups []*v3.RepositoryBackup

		for _, oldLabel := range oldRepository.Labels {
			labels = append(labels, &v3.RepositoryLabel{
				Id:          oldLabel.Id,
				Name:        oldLabel.Name,
				Color:       oldLabel.Color,
				Description: oldLabel.Description,
			})
		}

		for _, oldRelease := range oldRepository.Releases {
			releases = append(releases, &v3.RepositoryRelease{
				Id:      oldRelease.Id,
				TagName: oldRelease.TagName,
			})
		}

		for _, oldCollaborator := range oldRepository.Collaborators {
			collaborators = append(collaborators, &v3.RepositoryCollaborator{
				Id:         oldCollaborator.Id,
				Permission: v3.RepositoryCollaborator_Permission(oldCollaborator.Permission),
			})
		}

		for _, oldBackup := range oldRepository.Backups {
			backups = append(backups, &v3.RepositoryBackup{
				Store: v3.RepositoryBackup_Store(oldBackup.Store),
				Refs:  oldBackup.Refs,
			})
		}

		repository := v3.Repository{
			Creator: oldRepository.Creator,
			Id:      oldRepository.Id,
			Name:    oldRepository.Name,
			Owner: &v3.RepositoryOwner{
				Id:   oldRepository.Owner.Id,
				Type: v3.OwnerType(oldRepository.Owner.Type),
			},
			Description:         oldRepository.Description,
			Forks:               oldRepository.Forks,
			Subscribers:         oldRepository.Subscribers,
			Commits:             oldRepository.Commits,
			IssuesCount:         oldRepository.IssuesCount,
			PullsCount:          oldRepository.PullsCount,
			Labels:              labels,
			LabelsCount:         oldRepository.LabelsCount,
			Releases:            releases,
			CreatedAt:           oldRepository.CreatedAt,
			UpdatedAt:           oldRepository.UpdatedAt,
			PushedAt:            oldRepository.PushedAt,
			Stargazers:          oldRepository.Stargazers,
			Archived:            oldRepository.Archived,
			License:             oldRepository.License,
			DefaultBranch:       oldRepository.DefaultBranch,
			Parent:              oldRepository.Parent,
			Fork:                oldRepository.Fork,
			Collaborators:       collaborators,
			AllowForking:        oldRepository.AllowForking,
			Backups:             backups,
			EnableArweaveBackup: oldRepository.EnableArweaveBackup,
		}

		gitopiaV3Genesis.RepositoryList = append(gitopiaV3Genesis.RepositoryList, repository)
	}

	gitopiaV3Genesis.RepositoryCount = state.RepositoryCount

	commentMap = make(map[uint64]v2.Comment)
	for _, comment := range state.CommentList {
		commentMap[comment.Id] = comment
	}

	gitopiaV3Genesis.CommentCount = state.CommentCount

	for _, oldIssue := range state.IssueList {
		issue := v3.Issue{
			Creator:       oldIssue.Creator,
			Id:            oldIssue.Id,
			Iid:           oldIssue.Iid,
			Title:         oldIssue.Title,
			State:         v3.Issue_State(oldIssue.State),
			Description:   oldIssue.Description,
			CommentsCount: oldIssue.CommentsCount,
			RepositoryId:  oldIssue.RepositoryId,
			Labels:        oldIssue.Labels,
			Weight:        oldIssue.Weight,
			Assignees:     oldIssue.Assignees,
			CreatedAt:     oldIssue.CreatedAt,
			UpdatedAt:     oldIssue.UpdatedAt,
			ClosedAt:      oldIssue.ClosedAt,
			ClosedBy:      oldIssue.ClosedBy,
		}

		// Migrate comments
		for _, commentId := range oldIssue.Comments {
			oldComment := commentMap[commentId]

			var attachments []*v3.Attachment

			for _, attachmentStr := range oldComment.Attachments {
				var attachment v3.Attachment
				if err := json.Unmarshal([]byte(attachmentStr), &attachment); err != nil {
					return v3.GenesisState{}, err
				}

				attachments = append(attachments, &attachment)
			}

			commentType := v3.CommentTypeNone

			// Set comment type in case of system comment
			if oldComment.System {
				if strings.Contains(oldComment.Body, "assigned to") {
					commentType = v3.CommentTypeAddAssignees
				} else if strings.Contains(oldComment.Body, "unassigned") {
					commentType = v3.CommentTypeRemoveAssignees
				} else if strings.Contains(oldComment.Body, "added") {
					commentType = v3.CommentTypeAddLabels
				} else if strings.Contains(oldComment.Body, "remove") {
					commentType = v3.CommentTypeRemoveLabels
				} else if strings.Contains(oldComment.Body, "changed title from") {
					commentType = v3.CommentTypeModifiedTitle
				} else if strings.Contains(oldComment.Body, "changed the description") {
					commentType = v3.CommentTypeModifiedDescription
				} else if strings.Contains(oldComment.Body, "reopened") {
					commentType = v3.CommentTypeIssueOpened
				} else if strings.Contains(oldComment.Body, "closed") {
					commentType = v3.CommentTypeIssueClosed
				}
			} else {
				commentType = v3.CommentTypeReply
			}

			comment := v3.Comment{
				Creator:           oldComment.Creator,
				Id:                oldComment.Id,
				RepositoryId:      issue.RepositoryId,
				ParentIid:         issue.Iid,
				Parent:            v3.CommentParentIssue,
				CommentIid:        oldComment.CommentIid,
				Body:              oldComment.Body,
				Attachments:       attachments,
				DiffHunk:          oldComment.DiffHunk,
				Path:              oldComment.Path,
				System:            oldComment.System,
				AuthorAssociation: oldComment.AuthorAssociation,
				CreatedAt:         oldComment.CreatedAt,
				UpdatedAt:         oldComment.UpdatedAt,
				CommentType:       commentType,
			}

			gitopiaV3Genesis.CommentList = append(gitopiaV3Genesis.CommentList, comment)
		}

		gitopiaV3Genesis.IssueList = append(gitopiaV3Genesis.IssueList, issue)
	}

	gitopiaV3Genesis.IssueCount = state.IssueCount

	for _, oldPullRequest := range state.GetPullRequestList() {
		pullRequest := v3.PullRequest{
			Creator:             oldPullRequest.Creator,
			Id:                  oldPullRequest.Id,
			Iid:                 oldPullRequest.Iid,
			Title:               oldPullRequest.Title,
			State:               v3.PullRequest_State(oldPullRequest.State),
			Description:         oldPullRequest.Description,
			Locked:              oldPullRequest.Locked,
			CommentsCount:       oldPullRequest.CommentsCount,
			Labels:              oldPullRequest.Labels,
			Assignees:           oldPullRequest.Assignees,
			Reviewers:           oldPullRequest.Reviewers,
			Draft:               oldPullRequest.Draft,
			CreatedAt:           oldPullRequest.CreatedAt,
			UpdatedAt:           oldPullRequest.UpdatedAt,
			ClosedAt:            oldPullRequest.ClosedAt,
			ClosedBy:            oldPullRequest.ClosedBy,
			MergedAt:            oldPullRequest.MergedAt,
			MergedBy:            oldPullRequest.MergedBy,
			MergeCommitSha:      oldPullRequest.MergeCommitSha,
			MaintainerCanModify: oldPullRequest.MaintainerCanModify,
			Head:                (*v3.PullRequestHead)(oldPullRequest.Head),
			Base:                (*v3.PullRequestBase)(oldPullRequest.Base),
		}

		// Migrate comments
		for _, commentId := range oldPullRequest.Comments {
			oldComment := commentMap[commentId]

			var attachments []*v3.Attachment

			for _, attachmentStr := range oldComment.Attachments {
				var attachment v3.Attachment
				if err := json.Unmarshal([]byte(attachmentStr), &attachment); err != nil {
					return v3.GenesisState{}, err
				}

				attachments = append(attachments, &attachment)
			}

			var commentType v3.CommentType

			// Set comment type in case of system comment
			if oldComment.System {
				if strings.Contains(oldComment.Body, "assigned to") {
					commentType = v3.CommentTypeAddAssignees
				} else if strings.Contains(oldComment.Body, "unassigned") {
					commentType = v3.CommentTypeRemoveAssignees
				} else if strings.Contains(oldComment.Body, "added") {
					commentType = v3.CommentTypeAddLabels
				} else if strings.Contains(oldComment.Body, "remove") {
					commentType = v3.CommentTypeRemoveLabels
				} else if strings.Contains(oldComment.Body, "changed title from") {
					commentType = v3.CommentTypeModifiedTitle
				} else if strings.Contains(oldComment.Body, "changed the description") {
					commentType = v3.CommentTypeModifiedDescription
				} else if strings.Contains(oldComment.Body, "reopened") {
					commentType = v3.CommentTypePullRequestOpened
				} else if strings.Contains(oldComment.Body, "closed") {
					commentType = v3.CommentTypePullRequestClosed
				} else if strings.Contains(oldComment.Body, "merged") {
					commentType = v3.CommentTypePullRequestMerged
				} else if strings.Contains(oldComment.Body, "requested review from") {
					commentType = v3.CommentTypeAddReviewers
				} else if strings.Contains(oldComment.Body, "removed review request for") {
					commentType = v3.CommentTypeRemoveReviewers
				}
			}

			comment := v3.Comment{
				Creator:           oldComment.Creator,
				Id:                oldComment.Id,
				RepositoryId:      pullRequest.Base.RepositoryId,
				ParentIid:         pullRequest.Iid,
				Parent:            v3.CommentParentPullRequest,
				CommentIid:        oldComment.CommentIid,
				Body:              oldComment.Body,
				Attachments:       attachments,
				DiffHunk:          oldComment.DiffHunk,
				Path:              oldComment.Path,
				System:            oldComment.System,
				AuthorAssociation: oldComment.AuthorAssociation,
				CreatedAt:         oldComment.CreatedAt,
				UpdatedAt:         oldComment.UpdatedAt,
				CommentType:       commentType,
			}

			gitopiaV3Genesis.CommentList = append(gitopiaV3Genesis.CommentList, comment)
		}

		gitopiaV3Genesis.PullRequestList = append(gitopiaV3Genesis.PullRequestList, pullRequest)
	}

	gitopiaV3Genesis.PullRequestCount = state.PullRequestCount

	return gitopiaV3Genesis, nil
}

func GenerateGenesisCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "generate-genesis [exported-testnet-state]",
		Short: "Generate Mainnet Genesis file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var ctx = client.GetClientContextFromCmd(cmd)

			blob, err := ioutil.ReadFile(args[0])
			if err != nil {
				return err
			}

			chainID, err := cmd.Flags().GetString(flags.FlagChainID)
			if err != nil {
				return err
			}

			genesisTime, err := cmd.Flags().GetString(flagGenesisTime)
			if err != nil {
				return err
			}

			initialHeight, err := cmd.Flags().GetInt64(flagInitialHeight)
			if err != nil {
				return err
			}

			genesis, err := tmtypes.GenesisDocFromJSON(blob)
			if err != nil {
				return err
			}

			var state genutiltypes.AppMap
			if err := json.Unmarshal(genesis.AppState, &state); err != nil {
				return err
			}

			var gitopiaV2Genesis v2.GenesisState
			ctx.Codec.MustUnmarshalJSON(state[gitopiatypes.ModuleName], &gitopiaV2Genesis)
			gitopiaV3Genesis, err := migrateTestnetState(gitopiaV2Genesis)
			if err != nil {
				return err
			}

			var (
				authGenesis         = authtypes.DefaultGenesisState()
				bankGenesis         = banktypes.DefaultGenesisState()
				crisisGenesis       = crisistypes.DefaultGenesisState()
				govGenesis          = govv1types.DefaultGenesisState()
				mintGenesis         = minttypes.DefaultGenesisState()
				ibcGenesis          = ibctypes.DefaultGenesisState()
				distributionGenesis = distributiontypes.DefaultGenesisState()
				slashingGenesis     = slashingtypes.DefaultGenesisState()
				genutilGenesis      = genutiltypes.DefaultGenesisState()
				stakingGenesis      = stakingtypes.DefaultGenesisState()
				authzGenesis        = authz.DefaultGenesisState()
				groupGenesis        = group.NewGenesisState()
				capabilityGenesis   = capabilitytypes.DefaultGenesis()
				evidenceGenesis     = evidencetypes.DefaultGenesisState()
				feegrantGenesis     = feegranttypes.DefaultGenesisState()
				rewardsGenesis      = rewardstypes.DefaultGenesis()
			)

			fourteenDays := 14 * 24 * time.Hour
			depositParams := govv1types.NewDepositParams(
				sdk.NewCoins(sdk.NewCoin(params.BaseCoinUnit, math.NewInt(10000000))),
				fourteenDays,
			)
			govGenesis.DepositParams = &depositParams

			votingParams := govv1types.NewVotingParams(fourteenDays)
			govGenesis.VotingParams = &votingParams

			bankGenesis.DenomMetadata = []banktypes.Metadata{
				{
					Description: "The native staking token of the Gitopia Hub.",
					DenomUnits: []*banktypes.DenomUnit{
						{Denom: params.BaseCoinUnit, Exponent: uint32(0), Aliases: []string{"microlore"}},
						{Denom: params.HumanCoinUnit, Exponent: uint32(6), Aliases: []string{}},
					},
					Base:    params.BaseCoinUnit,
					Display: params.HumanCoinUnit,
					Name:    params.HumanCoinUnit,
					Symbol:  params.HumanCoinUnit,
				},
			}

			crisisGenesis.ConstantFee = sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(1000))

			mintGenesis.Minter.Inflation = sdk.NewDecWithPrec(35, 2)
			mintGenesis.Params = minttypes.Params{
				MintDenom:           params.BaseCoinUnit,
				InflationRateChange: sdk.NewDecWithPrec(20, 2),
				InflationMax:        sdk.NewDecWithPrec(45, 2),
				InflationMin:        sdk.NewDecWithPrec(25, 2),
				GoalBonded:          sdk.NewDecWithPrec(67, 2),
				BlocksPerYear:       uint64(60 * 60 * 8766 / 1.5), // assuming 1.5 second block times
			}

			stakingGenesis.Params = stakingtypes.NewParams(
				stakingtypes.DefaultUnbondingTime,
				stakingtypes.DefaultMaxValidators,
				stakingtypes.DefaultMaxEntries,
				stakingtypes.DefaultHistoricalEntries,
				params.BaseCoinUnit,
				stakingtypes.DefaultMinCommissionRate,
			)

			rewardsGenesis.Params = rewardstypes.NewParams(rewardsServiceAddress, &rewardstypes.RewardSeries{
				SeriesOne: &rewardstypes.RewardPool{
					TotalAmount:   sdk.NewCoin(params.BaseCoinUnit, math.NewInt(2_000_000_000_000)),
					ClaimedAmount: sdk.NewCoin(params.BaseCoinUnit, math.NewInt(0)),
				},
				SeriesTwo: &rewardstypes.RewardPool{
					TotalAmount:   sdk.NewCoin(params.BaseCoinUnit, math.NewInt(1_700_000_000_000)),
					ClaimedAmount: sdk.NewCoin(params.BaseCoinUnit, math.NewInt(0)),
				},
				SeriesThree: &rewardstypes.RewardPool{
					TotalAmount:   sdk.NewCoin(params.BaseCoinUnit, math.NewInt(1_400_000_000_000)),
					ClaimedAmount: sdk.NewCoin(params.BaseCoinUnit, math.NewInt(0)),
				},
				SeriesFour: &rewardstypes.RewardPool{
					TotalAmount:   sdk.NewCoin(params.BaseCoinUnit, math.NewInt(1_100_000_000_000)),
					ClaimedAmount: sdk.NewCoin(params.BaseCoinUnit, math.NewInt(0)),
				},
				SeriesFive: &rewardstypes.RewardPool{
					TotalAmount:   sdk.NewCoin(params.BaseCoinUnit, math.NewInt(800_000_000_000)),
					ClaimedAmount: sdk.NewCoin(params.BaseCoinUnit, math.NewInt(0)),
				},
				SeriesSix: &rewardstypes.RewardPool{
					TotalAmount:   sdk.NewCoin(params.BaseCoinUnit, math.NewInt(500_000_000_000)),
					ClaimedAmount: sdk.NewCoin(params.BaseCoinUnit, math.NewInt(0)),
				},
				SeriesSeven: &rewardstypes.RewardPool{
					TotalAmount:   sdk.NewCoin(params.BaseCoinUnit, math.NewInt(250_000_000_000)),
					ClaimedAmount: sdk.NewCoin(params.BaseCoinUnit, math.NewInt(0)),
				},
			})

			state[authtypes.ModuleName] = ctx.Codec.MustMarshalJSON(authGenesis)
			state[banktypes.ModuleName] = ctx.Codec.MustMarshalJSON(bankGenesis)
			state[crisistypes.ModuleName] = ctx.Codec.MustMarshalJSON(crisisGenesis)
			state[govtypes.ModuleName] = ctx.Codec.MustMarshalJSON(govGenesis)
			state[distributiontypes.ModuleName] = ctx.Codec.MustMarshalJSON(distributionGenesis)
			state[minttypes.ModuleName] = ctx.Codec.MustMarshalJSON(mintGenesis)
			state[slashingtypes.ModuleName] = ctx.Codec.MustMarshalJSON(slashingGenesis)
			state[genutiltypes.ModuleName] = ctx.Codec.MustMarshalJSON(genutilGenesis)
			state[stakingtypes.ModuleName] = ctx.Codec.MustMarshalJSON(stakingGenesis)
			state["ibc"] = ctx.Codec.MustMarshalJSON(ibcGenesis)
			state[authz.ModuleName] = ctx.Codec.MustMarshalJSON(authzGenesis)
			state[gitopiatypes.ModuleName] = ctx.Codec.MustMarshalJSON(&gitopiaV3Genesis)
			state[group.ModuleName] = ctx.Codec.MustMarshalJSON(groupGenesis)
			state[capabilitytypes.ModuleName] = ctx.Codec.MustMarshalJSON(capabilityGenesis)
			state[evidencetypes.ModuleName] = ctx.Codec.MustMarshalJSON(evidenceGenesis)
			state[feegranttypes.ModuleName] = ctx.Codec.MustMarshalJSON(feegrantGenesis)
			state[rewardstypes.ModuleName] = ctx.Codec.MustMarshalJSON(rewardsGenesis)

			genesis.AppState, err = json.Marshal(state)
			if err != nil {
				return err
			}

			if genesisTime != "" {
				var t time.Time
				if err := t.UnmarshalText([]byte(genesisTime)); err != nil {
					return err
				}

				genesis.GenesisTime = t
			}
			if chainID != "" {
				genesis.ChainID = chainID
			}

			genesis.InitialHeight = initialHeight

			genesis.Validators = []tmtypes.GenesisValidator{}

			blob, err = tmjson.Marshal(genesis)
			if err != nil {
				return err
			}

			sortedBlob, err := sdk.SortJSON(blob)
			if err != nil {
				return err
			}

			fmt.Println(string(sortedBlob))
			return nil
		},
	}

	cmd.Flags().String(flags.FlagChainID, "gitopia", "set chain id")
	cmd.Flags().String(flagGenesisTime, time.Now().UTC().Format(time.RFC3339Nano), "set genesis time")
	cmd.Flags().Int64(flagInitialHeight, 1, "set the initial height")

	return &cmd
}