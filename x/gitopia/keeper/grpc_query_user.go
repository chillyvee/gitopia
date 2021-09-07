package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/gitopia/gitopia/x/gitopia/types"
	"github.com/gitopia/gitopia/x/gitopia/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UserAll(c context.Context, req *types.QueryAllUserRequest) (*types.QueryAllUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var users []*types.User
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	userStore := prefix.NewStore(store, types.KeyPrefix(types.UserKey))

	pageRes, err := query.Paginate(userStore, req.Pagination, func(key []byte, value []byte) error {
		var user types.User
		if err := k.cdc.UnmarshalBinaryBare(value, &user); err != nil {
			return err
		}

		users = append(users, &user)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUserResponse{User: users, Pagination: pageRes}, nil
}

func (k Keeper) User(c context.Context, req *types.QueryGetUserRequest) (*types.QueryGetUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var user types.User
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasUser(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserKey))
	key := []byte(types.UserKey + req.Id)
	k.cdc.MustUnmarshalBinaryBare(store.Get(key), &user)

	return &types.QueryGetUserResponse{User: &user}, nil
}

func (k Keeper) UserRepositoryAll(c context.Context, req *types.QueryAllUserRepositoryRequest) (*types.QueryAllUserRepositoryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var user types.User
	var repositories []*types.Repository
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasUser(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	userStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserKey))
	userKey := []byte(types.UserKey + req.Id)
	k.cdc.MustUnmarshalBinaryBare(userStore.Get(userKey), &user)

	repositoryStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RepositoryKey))

	for _, userRepository := range user.Repositories {
		var repository types.Repository
		k.cdc.MustUnmarshalBinaryBare(repositoryStore.Get(GetRepositoryIDBytes(userRepository.Id)), &repository)
		repositories = append(repositories, &repository)
	}

	return &types.QueryAllUserRepositoryResponse{Repository: repositories}, nil
}

func (k Keeper) UserRepository(c context.Context, req *types.QueryGetUserRepositoryRequest) (*types.QueryGetUserRepositoryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var user types.User
	var organization types.Organization
	var repository types.Repository
	ctx := sdk.UnwrapSDKContext(c)

	if k.HasUser(ctx, req.UserId) {
		userStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserKey))
		userKey := []byte(types.UserKey + req.UserId)
		k.cdc.UnmarshalBinaryBare(userStore.Get(userKey), &user)

		if i, exists := utils.UserRepositoryExists(user.Repositories, req.RepositoryName); exists {
			repositoryStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RepositoryKey))
			k.cdc.MustUnmarshalBinaryBare(repositoryStore.Get(GetRepositoryIDBytes(user.Repositories[i].Id)), &repository)
		}
	} else if k.HasOrganization(ctx, req.UserId) {
		organizationStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrganizationKey))
		organizationKey := []byte(types.OrganizationKey + req.UserId)
		k.cdc.MustUnmarshalBinaryBare(organizationStore.Get(organizationKey), &organization)

		if i, exists := utils.OrganizationRepositoryExists(organization.Repositories, req.RepositoryName); exists {
			repositoryStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RepositoryKey))
			k.cdc.MustUnmarshalBinaryBare(repositoryStore.Get(GetRepositoryIDBytes(organization.Repositories[i].Id)), &repository)
		}
	} else {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetUserRepositoryResponse{Repository: &repository}, nil
}

func (k Keeper) UserOrganizationAll(c context.Context, req *types.QueryAllUserOrganizationRequest) (*types.QueryAllUserOrganizationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var user types.User
	var organizations []*types.Organization
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasUser(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	userStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserKey))
	userKey := []byte(types.UserKey + req.Id)
	k.cdc.MustUnmarshalBinaryBare(userStore.Get(userKey), &user)

	organizationStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrganizationKey))

	for _, userOrganization := range user.Organizations {
		var organization types.Organization
		key := []byte(types.OrganizationKey + userOrganization.Id)
		k.cdc.MustUnmarshalBinaryBare(organizationStore.Get(key), &organization)
		organizations = append(organizations, &organization)
	}

	return &types.QueryAllUserOrganizationResponse{Organization: organizations}, nil
}
