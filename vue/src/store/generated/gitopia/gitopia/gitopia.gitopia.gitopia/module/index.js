// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdateRepositoryCollaborator } from "./types/gitopia/tx";
import { MsgDeleteOrganization } from "./types/gitopia/tx";
import { MsgUpdatePullRequestTitle } from "./types/gitopia/tx";
import { MsgCreateRepositoryLabel } from "./types/gitopia/tx";
import { MsgDeleteBranch } from "./types/gitopia/tx";
import { MsgRemoveIssueLabels } from "./types/gitopia/tx";
import { MsgDeletePullRequest } from "./types/gitopia/tx";
import { MsgCreateIssue } from "./types/gitopia/tx";
import { MsgUpdateRelease } from "./types/gitopia/tx";
import { MsgAddPullRequestReviewers } from "./types/gitopia/tx";
import { MsgRemoveRepositoryCollaborator } from "./types/gitopia/tx";
import { MsgToggleRepositoryForking } from "./types/gitopia/tx";
import { MsgCreateComment } from "./types/gitopia/tx";
import { MsgAddIssueAssignees } from "./types/gitopia/tx";
import { MsgChangeRequestState } from "./types/gitopia/tx";
import { MsgCreateRelease } from "./types/gitopia/tx";
import { MsgUpdateOrganizationMember } from "./types/gitopia/tx";
import { MsgTransferUser } from "./types/gitopia/tx";
import { MsgRemoveIssueAssignees } from "./types/gitopia/tx";
import { MsgUpdateRepositoryLabel } from "./types/gitopia/tx";
import { MsgDeleteIssue } from "./types/gitopia/tx";
import { MsgRemoveOrganizationMember } from "./types/gitopia/tx";
import { MsgDeleteRepositoryLabel } from "./types/gitopia/tx";
import { MsgSetRepositoryTag } from "./types/gitopia/tx";
import { MsgUpdateOrganization } from "./types/gitopia/tx";
import { MsgCreateRepository } from "./types/gitopia/tx";
import { MsgForkRepository } from "./types/gitopia/tx";
import { MsgDeleteUser } from "./types/gitopia/tx";
import { MsgRemovePullRequestReviewers } from "./types/gitopia/tx";
import { MsgRemovePullRequestAssignees } from "./types/gitopia/tx";
import { MsgDeleteRelease } from "./types/gitopia/tx";
import { MsgRenameRepository } from "./types/gitopia/tx";
import { MsgDeleteWhois } from "./types/gitopia/tx";
import { MsgCreatePullRequest } from "./types/gitopia/tx";
import { MsgSetWhois } from "./types/gitopia/tx";
import { MsgRenameOrganization } from "./types/gitopia/tx";
import { MsgAddPullRequestLabels } from "./types/gitopia/tx";
import { MsgUpdateRepository } from "./types/gitopia/tx";
import { MsgUpdateIssue } from "./types/gitopia/tx";
import { MsgCreateUser } from "./types/gitopia/tx";
import { MsgUpdatePullRequest } from "./types/gitopia/tx";
import { MsgSetRepositoryBranch } from "./types/gitopia/tx";
import { MsgDeleteComment } from "./types/gitopia/tx";
import { MsgDeleteRequest } from "./types/gitopia/tx";
import { MsgUpdateIssueTitle } from "./types/gitopia/tx";
import { MsgDeleteRepository } from "./types/gitopia/tx";
import { MsgAddIssueLabels } from "./types/gitopia/tx";
import { MsgUpdateUser } from "./types/gitopia/tx";
import { MsgUpdateRequest } from "./types/gitopia/tx";
import { MsgSetDefaultBranch } from "./types/gitopia/tx";
import { MsgDeleteTag } from "./types/gitopia/tx";
import { MsgUpdateWhois } from "./types/gitopia/tx";
import { MsgRemovePullRequestLabels } from "./types/gitopia/tx";
import { MsgToggleIssueState } from "./types/gitopia/tx";
import { MsgUpdatePullRequestDescription } from "./types/gitopia/tx";
import { MsgUpdateComment } from "./types/gitopia/tx";
import { MsgCreateOrganization } from "./types/gitopia/tx";
import { MsgUpdateIssueDescription } from "./types/gitopia/tx";
import { MsgSetPullRequestState } from "./types/gitopia/tx";
import { MsgAddPullRequestAssignees } from "./types/gitopia/tx";
import { MsgChangeOwner } from "./types/gitopia/tx";
const types = [
    ["/gitopia.gitopia.gitopia.MsgUpdateRepositoryCollaborator", MsgUpdateRepositoryCollaborator],
    ["/gitopia.gitopia.gitopia.MsgDeleteOrganization", MsgDeleteOrganization],
    ["/gitopia.gitopia.gitopia.MsgUpdatePullRequestTitle", MsgUpdatePullRequestTitle],
    ["/gitopia.gitopia.gitopia.MsgCreateRepositoryLabel", MsgCreateRepositoryLabel],
    ["/gitopia.gitopia.gitopia.MsgDeleteBranch", MsgDeleteBranch],
    ["/gitopia.gitopia.gitopia.MsgRemoveIssueLabels", MsgRemoveIssueLabels],
    ["/gitopia.gitopia.gitopia.MsgDeletePullRequest", MsgDeletePullRequest],
    ["/gitopia.gitopia.gitopia.MsgCreateIssue", MsgCreateIssue],
    ["/gitopia.gitopia.gitopia.MsgUpdateRelease", MsgUpdateRelease],
    ["/gitopia.gitopia.gitopia.MsgAddPullRequestReviewers", MsgAddPullRequestReviewers],
    ["/gitopia.gitopia.gitopia.MsgRemoveRepositoryCollaborator", MsgRemoveRepositoryCollaborator],
    ["/gitopia.gitopia.gitopia.MsgToggleRepositoryForking", MsgToggleRepositoryForking],
    ["/gitopia.gitopia.gitopia.MsgCreateComment", MsgCreateComment],
    ["/gitopia.gitopia.gitopia.MsgAddIssueAssignees", MsgAddIssueAssignees],
    ["/gitopia.gitopia.gitopia.MsgChangeRequestState", MsgChangeRequestState],
    ["/gitopia.gitopia.gitopia.MsgCreateRelease", MsgCreateRelease],
    ["/gitopia.gitopia.gitopia.MsgUpdateOrganizationMember", MsgUpdateOrganizationMember],
    ["/gitopia.gitopia.gitopia.MsgTransferUser", MsgTransferUser],
    ["/gitopia.gitopia.gitopia.MsgRemoveIssueAssignees", MsgRemoveIssueAssignees],
    ["/gitopia.gitopia.gitopia.MsgUpdateRepositoryLabel", MsgUpdateRepositoryLabel],
    ["/gitopia.gitopia.gitopia.MsgDeleteIssue", MsgDeleteIssue],
    ["/gitopia.gitopia.gitopia.MsgRemoveOrganizationMember", MsgRemoveOrganizationMember],
    ["/gitopia.gitopia.gitopia.MsgDeleteRepositoryLabel", MsgDeleteRepositoryLabel],
    ["/gitopia.gitopia.gitopia.MsgSetRepositoryTag", MsgSetRepositoryTag],
    ["/gitopia.gitopia.gitopia.MsgUpdateOrganization", MsgUpdateOrganization],
    ["/gitopia.gitopia.gitopia.MsgCreateRepository", MsgCreateRepository],
    ["/gitopia.gitopia.gitopia.MsgForkRepository", MsgForkRepository],
    ["/gitopia.gitopia.gitopia.MsgDeleteUser", MsgDeleteUser],
    ["/gitopia.gitopia.gitopia.MsgRemovePullRequestReviewers", MsgRemovePullRequestReviewers],
    ["/gitopia.gitopia.gitopia.MsgRemovePullRequestAssignees", MsgRemovePullRequestAssignees],
    ["/gitopia.gitopia.gitopia.MsgDeleteRelease", MsgDeleteRelease],
    ["/gitopia.gitopia.gitopia.MsgRenameRepository", MsgRenameRepository],
    ["/gitopia.gitopia.gitopia.MsgDeleteWhois", MsgDeleteWhois],
    ["/gitopia.gitopia.gitopia.MsgCreatePullRequest", MsgCreatePullRequest],
    ["/gitopia.gitopia.gitopia.MsgSetWhois", MsgSetWhois],
    ["/gitopia.gitopia.gitopia.MsgRenameOrganization", MsgRenameOrganization],
    ["/gitopia.gitopia.gitopia.MsgAddPullRequestLabels", MsgAddPullRequestLabels],
    ["/gitopia.gitopia.gitopia.MsgUpdateRepository", MsgUpdateRepository],
    ["/gitopia.gitopia.gitopia.MsgUpdateIssue", MsgUpdateIssue],
    ["/gitopia.gitopia.gitopia.MsgCreateUser", MsgCreateUser],
    ["/gitopia.gitopia.gitopia.MsgUpdatePullRequest", MsgUpdatePullRequest],
    ["/gitopia.gitopia.gitopia.MsgSetRepositoryBranch", MsgSetRepositoryBranch],
    ["/gitopia.gitopia.gitopia.MsgDeleteComment", MsgDeleteComment],
    ["/gitopia.gitopia.gitopia.MsgDeleteRequest", MsgDeleteRequest],
    ["/gitopia.gitopia.gitopia.MsgUpdateIssueTitle", MsgUpdateIssueTitle],
    ["/gitopia.gitopia.gitopia.MsgDeleteRepository", MsgDeleteRepository],
    ["/gitopia.gitopia.gitopia.MsgAddIssueLabels", MsgAddIssueLabels],
    ["/gitopia.gitopia.gitopia.MsgUpdateUser", MsgUpdateUser],
    ["/gitopia.gitopia.gitopia.MsgUpdateRequest", MsgUpdateRequest],
    ["/gitopia.gitopia.gitopia.MsgSetDefaultBranch", MsgSetDefaultBranch],
    ["/gitopia.gitopia.gitopia.MsgDeleteTag", MsgDeleteTag],
    ["/gitopia.gitopia.gitopia.MsgUpdateWhois", MsgUpdateWhois],
    ["/gitopia.gitopia.gitopia.MsgRemovePullRequestLabels", MsgRemovePullRequestLabels],
    ["/gitopia.gitopia.gitopia.MsgToggleIssueState", MsgToggleIssueState],
    ["/gitopia.gitopia.gitopia.MsgUpdatePullRequestDescription", MsgUpdatePullRequestDescription],
    ["/gitopia.gitopia.gitopia.MsgUpdateComment", MsgUpdateComment],
    ["/gitopia.gitopia.gitopia.MsgCreateOrganization", MsgCreateOrganization],
    ["/gitopia.gitopia.gitopia.MsgUpdateIssueDescription", MsgUpdateIssueDescription],
    ["/gitopia.gitopia.gitopia.MsgSetPullRequestState", MsgSetPullRequestState],
    ["/gitopia.gitopia.gitopia.MsgAddPullRequestAssignees", MsgAddPullRequestAssignees],
    ["/gitopia.gitopia.gitopia.MsgChangeOwner", MsgChangeOwner],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgUpdateRepositoryCollaborator: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgUpdateRepositoryCollaborator", value: data }),
        msgDeleteOrganization: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgDeleteOrganization", value: data }),
        msgUpdatePullRequestTitle: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgUpdatePullRequestTitle", value: data }),
        msgCreateRepositoryLabel: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgCreateRepositoryLabel", value: data }),
        msgDeleteBranch: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgDeleteBranch", value: data }),
        msgRemoveIssueLabels: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgRemoveIssueLabels", value: data }),
        msgDeletePullRequest: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgDeletePullRequest", value: data }),
        msgCreateIssue: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgCreateIssue", value: data }),
        msgUpdateRelease: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgUpdateRelease", value: data }),
        msgAddPullRequestReviewers: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgAddPullRequestReviewers", value: data }),
        msgRemoveRepositoryCollaborator: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgRemoveRepositoryCollaborator", value: data }),
        msgToggleRepositoryForking: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgToggleRepositoryForking", value: data }),
        msgCreateComment: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgCreateComment", value: data }),
        msgAddIssueAssignees: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgAddIssueAssignees", value: data }),
        msgChangeRequestState: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgChangeRequestState", value: data }),
        msgCreateRelease: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgCreateRelease", value: data }),
        msgUpdateOrganizationMember: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgUpdateOrganizationMember", value: data }),
        msgTransferUser: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgTransferUser", value: data }),
        msgRemoveIssueAssignees: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgRemoveIssueAssignees", value: data }),
        msgUpdateRepositoryLabel: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgUpdateRepositoryLabel", value: data }),
        msgDeleteIssue: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgDeleteIssue", value: data }),
        msgRemoveOrganizationMember: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgRemoveOrganizationMember", value: data }),
        msgDeleteRepositoryLabel: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgDeleteRepositoryLabel", value: data }),
        msgSetRepositoryTag: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgSetRepositoryTag", value: data }),
        msgUpdateOrganization: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgUpdateOrganization", value: data }),
        msgCreateRepository: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgCreateRepository", value: data }),
        msgForkRepository: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgForkRepository", value: data }),
        msgDeleteUser: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgDeleteUser", value: data }),
        msgRemovePullRequestReviewers: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgRemovePullRequestReviewers", value: data }),
        msgRemovePullRequestAssignees: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgRemovePullRequestAssignees", value: data }),
        msgDeleteRelease: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgDeleteRelease", value: data }),
        msgRenameRepository: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgRenameRepository", value: data }),
        msgDeleteWhois: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgDeleteWhois", value: data }),
        msgCreatePullRequest: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgCreatePullRequest", value: data }),
        msgSetWhois: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgSetWhois", value: data }),
        msgRenameOrganization: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgRenameOrganization", value: data }),
        msgAddPullRequestLabels: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgAddPullRequestLabels", value: data }),
        msgUpdateRepository: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgUpdateRepository", value: data }),
        msgUpdateIssue: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgUpdateIssue", value: data }),
        msgCreateUser: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgCreateUser", value: data }),
        msgUpdatePullRequest: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgUpdatePullRequest", value: data }),
        msgSetRepositoryBranch: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgSetRepositoryBranch", value: data }),
        msgDeleteComment: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgDeleteComment", value: data }),
        msgDeleteRequest: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgDeleteRequest", value: data }),
        msgUpdateIssueTitle: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgUpdateIssueTitle", value: data }),
        msgDeleteRepository: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgDeleteRepository", value: data }),
        msgAddIssueLabels: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgAddIssueLabels", value: data }),
        msgUpdateUser: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgUpdateUser", value: data }),
        msgUpdateRequest: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgUpdateRequest", value: data }),
        msgSetDefaultBranch: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgSetDefaultBranch", value: data }),
        msgDeleteTag: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgDeleteTag", value: data }),
        msgUpdateWhois: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgUpdateWhois", value: data }),
        msgRemovePullRequestLabels: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgRemovePullRequestLabels", value: data }),
        msgToggleIssueState: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgToggleIssueState", value: data }),
        msgUpdatePullRequestDescription: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgUpdatePullRequestDescription", value: data }),
        msgUpdateComment: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgUpdateComment", value: data }),
        msgCreateOrganization: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgCreateOrganization", value: data }),
        msgUpdateIssueDescription: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgUpdateIssueDescription", value: data }),
        msgSetPullRequestState: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgSetPullRequestState", value: data }),
        msgAddPullRequestAssignees: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgAddPullRequestAssignees", value: data }),
        msgChangeOwner: (data) => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgChangeOwner", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
