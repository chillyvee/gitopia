import { Reader, Writer } from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { Request } from "../gitopia/request";
import { Release } from "../gitopia/release";
import { PullRequest } from "../gitopia/pullRequest";
import { Organization } from "../gitopia/organization";
import { Comment } from "../gitopia/comment";
import { Issue } from "../gitopia/issue";
import { Repository, RepositoryOwner, RepositoryBranch, RepositoryTag } from "../gitopia/repository";
import { User } from "../gitopia/user";
import { Whois } from "../gitopia/whois";
export declare const protobufPackage = "gitopia.gitopia.gitopia";
export interface QueryAllAddressRequestReceivedRequest {
    address: string;
    option: RequestOptions | undefined;
    pagination: PageRequest | undefined;
}
export interface QueryAllAddressRequestReceivedResponse {
    Request: Request[];
    pagination: PageResponse | undefined;
}
export interface QueryAllRepositoryRequestRequest {
    id: string;
    repositoryName: string;
    option: RequestOptions | undefined;
    pagination: PageRequest | undefined;
}
export interface RequestOptions {
    state: string;
    includeExpired: boolean;
    sort: string;
}
export interface QueryAllRepositoryRequestResponse {
    Request: Request[];
    pagination: PageResponse | undefined;
}
export interface QueryGetRequestRequest {
    id: number;
}
export interface QueryGetRequestResponse {
    Request: Request | undefined;
}
export interface QueryAllRequestRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllRequestResponse {
    Request: Request[];
    pagination: PageResponse | undefined;
}
/** this line is used by starport scaffolding # 3 */
export interface QueryGetReleaseRequest {
    id: number;
}
export interface QueryGetReleaseResponse {
    Release: Release | undefined;
}
export interface QueryAllReleaseRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllReleaseResponse {
    Release: Release[];
    pagination: PageResponse | undefined;
}
export interface QueryGetPullRequestRequest {
    id: number;
}
export interface QueryGetPullRequestResponse {
    PullRequest: PullRequest | undefined;
}
export interface QueryAllPullRequestRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllPullRequestResponse {
    PullRequest: PullRequest[];
    pagination: PageResponse | undefined;
}
export interface QueryGetOrganizationRequest {
    id: string;
}
export interface QueryGetOrganizationResponse {
    Organization: Organization | undefined;
}
export interface QueryAllOrganizationRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllOrganizationResponse {
    Organization: Organization[];
    pagination: PageResponse | undefined;
}
export interface QueryGetCommentRequest {
    id: number;
}
export interface QueryGetCommentResponse {
    Comment: Comment | undefined;
}
export interface QueryAllCommentRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllCommentResponse {
    Comment: Comment[];
    pagination: PageResponse | undefined;
}
export interface QueryGetIssueRequest {
    id: number;
}
export interface QueryGetIssueResponse {
    Issue: Issue | undefined;
}
export interface QueryAllIssueRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllIssueResponse {
    Issue: Issue[];
    pagination: PageResponse | undefined;
}
export interface QueryGetLatestRepositoryReleaseRequest {
    userId: string;
    repositoryName: string;
}
export interface QueryGetLatestRepositoryReleaseResponse {
    Release: Release | undefined;
}
export interface QueryGetRepositoryReleaseRequest {
    userId: string;
    repositoryName: string;
    tagName: string;
}
export interface QueryGetRepositoryReleaseResponse {
    Release: Release | undefined;
}
export interface QueryAllRepositoryReleaseRequest {
    userId: string;
    repositoryName: string;
    pagination: PageRequest | undefined;
}
export interface QueryAllRepositoryReleaseResponse {
    Release: Release[];
    pagination: PageResponse | undefined;
}
export interface QueryGetRepositoryIssueRequest {
    id: string;
    repositoryName: string;
    issueIid: number;
}
export interface QueryGetRepositoryIssueResponse {
    Issue: Issue | undefined;
}
export interface QueryGetRepositoryPullRequestRequest {
    userId: string;
    repositoryName: string;
    pullIid: number;
}
export interface QueryGetRepositoryPullRequestResponse {
    PullRequest: PullRequest | undefined;
}
export interface QueryAllRepositoryIssueRequest {
    id: string;
    repositoryName: string;
    option: IssueOptions | undefined;
    pagination: PageRequest | undefined;
}
export interface IssueOptions {
    createdBy: string;
    state: string;
    labels: string;
    assignee: string;
    labelIds: number[];
    sort: string;
    search: string;
    updatedAfter: number;
    updatedBefore: number;
}
export interface QueryAllRepositoryIssueResponse {
    Issue: Issue[];
    pagination: PageResponse | undefined;
}
export interface QueryAllRepositoryPullRequestRequest {
    userId: string;
    repositoryName: string;
    option: PullRequestOptions | undefined;
    pagination: PageRequest | undefined;
}
export interface PullRequestOptions {
    createdBy: string;
    state: string;
    labels: string;
    assignee: string;
    reviewer: string;
    labelIds: number[];
    sort: string;
    search: string;
    updatedAfter: number;
    updatedBefore: number;
}
export interface QueryAllRepositoryPullRequestResponse {
    PullRequest: PullRequest[];
    pagination: PageResponse | undefined;
}
export interface QueryGetRepositoryRequest {
    id: number;
}
export interface QueryGetRepositoryResponse {
    Repository: Repository | undefined;
}
export interface RepositoryFork {
    creator: string;
    id: number;
    name: string;
    owner: RepositoryOwner | undefined;
    description: string;
    parent: number;
    forksCount: number;
    issuesCount: number;
    pullsCount: number;
}
export interface QueryGetAllForkRequest {
    userId: string;
    repositoryName: string;
    pagination: PageRequest | undefined;
}
export interface QueryGetAllForkResponse {
    forks: RepositoryFork[];
    pagination: PageResponse | undefined;
}
export interface QueryGetAllBranchRequest {
    repositoryId: number;
}
export interface QueryGetAllBranchResponse {
    Branches: RepositoryBranch[];
}
export interface QueryGetBranchShaRequest {
    repositoryId: number;
    branchName: string;
}
export interface QueryGetBranchShaResponse {
    sha: string;
}
export interface QueryGetAllTagRequest {
    repositoryId: number;
}
export interface QueryGetAllTagResponse {
    Tags: RepositoryTag[];
}
export interface QueryGetTagShaRequest {
    repositoryId: number;
    tagName: string;
}
export interface QueryGetTagShaResponse {
    sha: string;
}
export interface QueryAllRepositoryRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllRepositoryResponse {
    Repository: Repository[];
    pagination: PageResponse | undefined;
}
export interface QueryGetUserRequest {
    id: string;
}
export interface QueryGetUserResponse {
    User: User | undefined;
}
export interface QueryAllUserRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllUserResponse {
    User: User[];
    pagination: PageResponse | undefined;
}
export interface QueryAllAddressRepositoryRequest {
    id: string;
    pagination: PageRequest | undefined;
}
export interface QueryAllAddressRepositoryResponse {
    Repository: Repository[];
    pagination: PageResponse | undefined;
}
export interface QueryGetAddressRepositoryRequest {
    id: string;
    repositoryName: string;
}
export interface QueryGetAddressRepositoryResponse {
    Repository: Repository | undefined;
}
export interface QueryAllUserOrganizationRequest {
    id: string;
}
export interface QueryAllUserOrganizationResponse {
    organization: Organization[];
}
export interface QueryGetWhoisRequest {
    name: string;
}
export interface QueryGetWhoisResponse {
    Whois: Whois | undefined;
}
export interface QueryAllWhoisRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllWhoisResponse {
    Whois: Whois[];
    pagination: PageResponse | undefined;
}
export declare const QueryAllAddressRequestReceivedRequest: {
    encode(message: QueryAllAddressRequestReceivedRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllAddressRequestReceivedRequest;
    fromJSON(object: any): QueryAllAddressRequestReceivedRequest;
    toJSON(message: QueryAllAddressRequestReceivedRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllAddressRequestReceivedRequest>): QueryAllAddressRequestReceivedRequest;
};
export declare const QueryAllAddressRequestReceivedResponse: {
    encode(message: QueryAllAddressRequestReceivedResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllAddressRequestReceivedResponse;
    fromJSON(object: any): QueryAllAddressRequestReceivedResponse;
    toJSON(message: QueryAllAddressRequestReceivedResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllAddressRequestReceivedResponse>): QueryAllAddressRequestReceivedResponse;
};
export declare const QueryAllRepositoryRequestRequest: {
    encode(message: QueryAllRepositoryRequestRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllRepositoryRequestRequest;
    fromJSON(object: any): QueryAllRepositoryRequestRequest;
    toJSON(message: QueryAllRepositoryRequestRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllRepositoryRequestRequest>): QueryAllRepositoryRequestRequest;
};
export declare const RequestOptions: {
    encode(message: RequestOptions, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): RequestOptions;
    fromJSON(object: any): RequestOptions;
    toJSON(message: RequestOptions): unknown;
    fromPartial(object: DeepPartial<RequestOptions>): RequestOptions;
};
export declare const QueryAllRepositoryRequestResponse: {
    encode(message: QueryAllRepositoryRequestResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllRepositoryRequestResponse;
    fromJSON(object: any): QueryAllRepositoryRequestResponse;
    toJSON(message: QueryAllRepositoryRequestResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllRepositoryRequestResponse>): QueryAllRepositoryRequestResponse;
};
export declare const QueryGetRequestRequest: {
    encode(message: QueryGetRequestRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetRequestRequest;
    fromJSON(object: any): QueryGetRequestRequest;
    toJSON(message: QueryGetRequestRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetRequestRequest>): QueryGetRequestRequest;
};
export declare const QueryGetRequestResponse: {
    encode(message: QueryGetRequestResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetRequestResponse;
    fromJSON(object: any): QueryGetRequestResponse;
    toJSON(message: QueryGetRequestResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetRequestResponse>): QueryGetRequestResponse;
};
export declare const QueryAllRequestRequest: {
    encode(message: QueryAllRequestRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllRequestRequest;
    fromJSON(object: any): QueryAllRequestRequest;
    toJSON(message: QueryAllRequestRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllRequestRequest>): QueryAllRequestRequest;
};
export declare const QueryAllRequestResponse: {
    encode(message: QueryAllRequestResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllRequestResponse;
    fromJSON(object: any): QueryAllRequestResponse;
    toJSON(message: QueryAllRequestResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllRequestResponse>): QueryAllRequestResponse;
};
export declare const QueryGetReleaseRequest: {
    encode(message: QueryGetReleaseRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetReleaseRequest;
    fromJSON(object: any): QueryGetReleaseRequest;
    toJSON(message: QueryGetReleaseRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetReleaseRequest>): QueryGetReleaseRequest;
};
export declare const QueryGetReleaseResponse: {
    encode(message: QueryGetReleaseResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetReleaseResponse;
    fromJSON(object: any): QueryGetReleaseResponse;
    toJSON(message: QueryGetReleaseResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetReleaseResponse>): QueryGetReleaseResponse;
};
export declare const QueryAllReleaseRequest: {
    encode(message: QueryAllReleaseRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllReleaseRequest;
    fromJSON(object: any): QueryAllReleaseRequest;
    toJSON(message: QueryAllReleaseRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllReleaseRequest>): QueryAllReleaseRequest;
};
export declare const QueryAllReleaseResponse: {
    encode(message: QueryAllReleaseResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllReleaseResponse;
    fromJSON(object: any): QueryAllReleaseResponse;
    toJSON(message: QueryAllReleaseResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllReleaseResponse>): QueryAllReleaseResponse;
};
export declare const QueryGetPullRequestRequest: {
    encode(message: QueryGetPullRequestRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetPullRequestRequest;
    fromJSON(object: any): QueryGetPullRequestRequest;
    toJSON(message: QueryGetPullRequestRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetPullRequestRequest>): QueryGetPullRequestRequest;
};
export declare const QueryGetPullRequestResponse: {
    encode(message: QueryGetPullRequestResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetPullRequestResponse;
    fromJSON(object: any): QueryGetPullRequestResponse;
    toJSON(message: QueryGetPullRequestResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetPullRequestResponse>): QueryGetPullRequestResponse;
};
export declare const QueryAllPullRequestRequest: {
    encode(message: QueryAllPullRequestRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllPullRequestRequest;
    fromJSON(object: any): QueryAllPullRequestRequest;
    toJSON(message: QueryAllPullRequestRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllPullRequestRequest>): QueryAllPullRequestRequest;
};
export declare const QueryAllPullRequestResponse: {
    encode(message: QueryAllPullRequestResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllPullRequestResponse;
    fromJSON(object: any): QueryAllPullRequestResponse;
    toJSON(message: QueryAllPullRequestResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllPullRequestResponse>): QueryAllPullRequestResponse;
};
export declare const QueryGetOrganizationRequest: {
    encode(message: QueryGetOrganizationRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetOrganizationRequest;
    fromJSON(object: any): QueryGetOrganizationRequest;
    toJSON(message: QueryGetOrganizationRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetOrganizationRequest>): QueryGetOrganizationRequest;
};
export declare const QueryGetOrganizationResponse: {
    encode(message: QueryGetOrganizationResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetOrganizationResponse;
    fromJSON(object: any): QueryGetOrganizationResponse;
    toJSON(message: QueryGetOrganizationResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetOrganizationResponse>): QueryGetOrganizationResponse;
};
export declare const QueryAllOrganizationRequest: {
    encode(message: QueryAllOrganizationRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllOrganizationRequest;
    fromJSON(object: any): QueryAllOrganizationRequest;
    toJSON(message: QueryAllOrganizationRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllOrganizationRequest>): QueryAllOrganizationRequest;
};
export declare const QueryAllOrganizationResponse: {
    encode(message: QueryAllOrganizationResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllOrganizationResponse;
    fromJSON(object: any): QueryAllOrganizationResponse;
    toJSON(message: QueryAllOrganizationResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllOrganizationResponse>): QueryAllOrganizationResponse;
};
export declare const QueryGetCommentRequest: {
    encode(message: QueryGetCommentRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetCommentRequest;
    fromJSON(object: any): QueryGetCommentRequest;
    toJSON(message: QueryGetCommentRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetCommentRequest>): QueryGetCommentRequest;
};
export declare const QueryGetCommentResponse: {
    encode(message: QueryGetCommentResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetCommentResponse;
    fromJSON(object: any): QueryGetCommentResponse;
    toJSON(message: QueryGetCommentResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetCommentResponse>): QueryGetCommentResponse;
};
export declare const QueryAllCommentRequest: {
    encode(message: QueryAllCommentRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllCommentRequest;
    fromJSON(object: any): QueryAllCommentRequest;
    toJSON(message: QueryAllCommentRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllCommentRequest>): QueryAllCommentRequest;
};
export declare const QueryAllCommentResponse: {
    encode(message: QueryAllCommentResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllCommentResponse;
    fromJSON(object: any): QueryAllCommentResponse;
    toJSON(message: QueryAllCommentResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllCommentResponse>): QueryAllCommentResponse;
};
export declare const QueryGetIssueRequest: {
    encode(message: QueryGetIssueRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetIssueRequest;
    fromJSON(object: any): QueryGetIssueRequest;
    toJSON(message: QueryGetIssueRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetIssueRequest>): QueryGetIssueRequest;
};
export declare const QueryGetIssueResponse: {
    encode(message: QueryGetIssueResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetIssueResponse;
    fromJSON(object: any): QueryGetIssueResponse;
    toJSON(message: QueryGetIssueResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetIssueResponse>): QueryGetIssueResponse;
};
export declare const QueryAllIssueRequest: {
    encode(message: QueryAllIssueRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllIssueRequest;
    fromJSON(object: any): QueryAllIssueRequest;
    toJSON(message: QueryAllIssueRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllIssueRequest>): QueryAllIssueRequest;
};
export declare const QueryAllIssueResponse: {
    encode(message: QueryAllIssueResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllIssueResponse;
    fromJSON(object: any): QueryAllIssueResponse;
    toJSON(message: QueryAllIssueResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllIssueResponse>): QueryAllIssueResponse;
};
export declare const QueryGetLatestRepositoryReleaseRequest: {
    encode(message: QueryGetLatestRepositoryReleaseRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetLatestRepositoryReleaseRequest;
    fromJSON(object: any): QueryGetLatestRepositoryReleaseRequest;
    toJSON(message: QueryGetLatestRepositoryReleaseRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetLatestRepositoryReleaseRequest>): QueryGetLatestRepositoryReleaseRequest;
};
export declare const QueryGetLatestRepositoryReleaseResponse: {
    encode(message: QueryGetLatestRepositoryReleaseResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetLatestRepositoryReleaseResponse;
    fromJSON(object: any): QueryGetLatestRepositoryReleaseResponse;
    toJSON(message: QueryGetLatestRepositoryReleaseResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetLatestRepositoryReleaseResponse>): QueryGetLatestRepositoryReleaseResponse;
};
export declare const QueryGetRepositoryReleaseRequest: {
    encode(message: QueryGetRepositoryReleaseRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetRepositoryReleaseRequest;
    fromJSON(object: any): QueryGetRepositoryReleaseRequest;
    toJSON(message: QueryGetRepositoryReleaseRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetRepositoryReleaseRequest>): QueryGetRepositoryReleaseRequest;
};
export declare const QueryGetRepositoryReleaseResponse: {
    encode(message: QueryGetRepositoryReleaseResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetRepositoryReleaseResponse;
    fromJSON(object: any): QueryGetRepositoryReleaseResponse;
    toJSON(message: QueryGetRepositoryReleaseResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetRepositoryReleaseResponse>): QueryGetRepositoryReleaseResponse;
};
export declare const QueryAllRepositoryReleaseRequest: {
    encode(message: QueryAllRepositoryReleaseRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllRepositoryReleaseRequest;
    fromJSON(object: any): QueryAllRepositoryReleaseRequest;
    toJSON(message: QueryAllRepositoryReleaseRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllRepositoryReleaseRequest>): QueryAllRepositoryReleaseRequest;
};
export declare const QueryAllRepositoryReleaseResponse: {
    encode(message: QueryAllRepositoryReleaseResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllRepositoryReleaseResponse;
    fromJSON(object: any): QueryAllRepositoryReleaseResponse;
    toJSON(message: QueryAllRepositoryReleaseResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllRepositoryReleaseResponse>): QueryAllRepositoryReleaseResponse;
};
export declare const QueryGetRepositoryIssueRequest: {
    encode(message: QueryGetRepositoryIssueRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetRepositoryIssueRequest;
    fromJSON(object: any): QueryGetRepositoryIssueRequest;
    toJSON(message: QueryGetRepositoryIssueRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetRepositoryIssueRequest>): QueryGetRepositoryIssueRequest;
};
export declare const QueryGetRepositoryIssueResponse: {
    encode(message: QueryGetRepositoryIssueResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetRepositoryIssueResponse;
    fromJSON(object: any): QueryGetRepositoryIssueResponse;
    toJSON(message: QueryGetRepositoryIssueResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetRepositoryIssueResponse>): QueryGetRepositoryIssueResponse;
};
export declare const QueryGetRepositoryPullRequestRequest: {
    encode(message: QueryGetRepositoryPullRequestRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetRepositoryPullRequestRequest;
    fromJSON(object: any): QueryGetRepositoryPullRequestRequest;
    toJSON(message: QueryGetRepositoryPullRequestRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetRepositoryPullRequestRequest>): QueryGetRepositoryPullRequestRequest;
};
export declare const QueryGetRepositoryPullRequestResponse: {
    encode(message: QueryGetRepositoryPullRequestResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetRepositoryPullRequestResponse;
    fromJSON(object: any): QueryGetRepositoryPullRequestResponse;
    toJSON(message: QueryGetRepositoryPullRequestResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetRepositoryPullRequestResponse>): QueryGetRepositoryPullRequestResponse;
};
export declare const QueryAllRepositoryIssueRequest: {
    encode(message: QueryAllRepositoryIssueRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllRepositoryIssueRequest;
    fromJSON(object: any): QueryAllRepositoryIssueRequest;
    toJSON(message: QueryAllRepositoryIssueRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllRepositoryIssueRequest>): QueryAllRepositoryIssueRequest;
};
export declare const IssueOptions: {
    encode(message: IssueOptions, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): IssueOptions;
    fromJSON(object: any): IssueOptions;
    toJSON(message: IssueOptions): unknown;
    fromPartial(object: DeepPartial<IssueOptions>): IssueOptions;
};
export declare const QueryAllRepositoryIssueResponse: {
    encode(message: QueryAllRepositoryIssueResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllRepositoryIssueResponse;
    fromJSON(object: any): QueryAllRepositoryIssueResponse;
    toJSON(message: QueryAllRepositoryIssueResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllRepositoryIssueResponse>): QueryAllRepositoryIssueResponse;
};
export declare const QueryAllRepositoryPullRequestRequest: {
    encode(message: QueryAllRepositoryPullRequestRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllRepositoryPullRequestRequest;
    fromJSON(object: any): QueryAllRepositoryPullRequestRequest;
    toJSON(message: QueryAllRepositoryPullRequestRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllRepositoryPullRequestRequest>): QueryAllRepositoryPullRequestRequest;
};
export declare const PullRequestOptions: {
    encode(message: PullRequestOptions, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): PullRequestOptions;
    fromJSON(object: any): PullRequestOptions;
    toJSON(message: PullRequestOptions): unknown;
    fromPartial(object: DeepPartial<PullRequestOptions>): PullRequestOptions;
};
export declare const QueryAllRepositoryPullRequestResponse: {
    encode(message: QueryAllRepositoryPullRequestResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllRepositoryPullRequestResponse;
    fromJSON(object: any): QueryAllRepositoryPullRequestResponse;
    toJSON(message: QueryAllRepositoryPullRequestResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllRepositoryPullRequestResponse>): QueryAllRepositoryPullRequestResponse;
};
export declare const QueryGetRepositoryRequest: {
    encode(message: QueryGetRepositoryRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetRepositoryRequest;
    fromJSON(object: any): QueryGetRepositoryRequest;
    toJSON(message: QueryGetRepositoryRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetRepositoryRequest>): QueryGetRepositoryRequest;
};
export declare const QueryGetRepositoryResponse: {
    encode(message: QueryGetRepositoryResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetRepositoryResponse;
    fromJSON(object: any): QueryGetRepositoryResponse;
    toJSON(message: QueryGetRepositoryResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetRepositoryResponse>): QueryGetRepositoryResponse;
};
export declare const RepositoryFork: {
    encode(message: RepositoryFork, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): RepositoryFork;
    fromJSON(object: any): RepositoryFork;
    toJSON(message: RepositoryFork): unknown;
    fromPartial(object: DeepPartial<RepositoryFork>): RepositoryFork;
};
export declare const QueryGetAllForkRequest: {
    encode(message: QueryGetAllForkRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetAllForkRequest;
    fromJSON(object: any): QueryGetAllForkRequest;
    toJSON(message: QueryGetAllForkRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetAllForkRequest>): QueryGetAllForkRequest;
};
export declare const QueryGetAllForkResponse: {
    encode(message: QueryGetAllForkResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetAllForkResponse;
    fromJSON(object: any): QueryGetAllForkResponse;
    toJSON(message: QueryGetAllForkResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetAllForkResponse>): QueryGetAllForkResponse;
};
export declare const QueryGetAllBranchRequest: {
    encode(message: QueryGetAllBranchRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetAllBranchRequest;
    fromJSON(object: any): QueryGetAllBranchRequest;
    toJSON(message: QueryGetAllBranchRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetAllBranchRequest>): QueryGetAllBranchRequest;
};
export declare const QueryGetAllBranchResponse: {
    encode(message: QueryGetAllBranchResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetAllBranchResponse;
    fromJSON(object: any): QueryGetAllBranchResponse;
    toJSON(message: QueryGetAllBranchResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetAllBranchResponse>): QueryGetAllBranchResponse;
};
export declare const QueryGetBranchShaRequest: {
    encode(message: QueryGetBranchShaRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetBranchShaRequest;
    fromJSON(object: any): QueryGetBranchShaRequest;
    toJSON(message: QueryGetBranchShaRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetBranchShaRequest>): QueryGetBranchShaRequest;
};
export declare const QueryGetBranchShaResponse: {
    encode(message: QueryGetBranchShaResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetBranchShaResponse;
    fromJSON(object: any): QueryGetBranchShaResponse;
    toJSON(message: QueryGetBranchShaResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetBranchShaResponse>): QueryGetBranchShaResponse;
};
export declare const QueryGetAllTagRequest: {
    encode(message: QueryGetAllTagRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetAllTagRequest;
    fromJSON(object: any): QueryGetAllTagRequest;
    toJSON(message: QueryGetAllTagRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetAllTagRequest>): QueryGetAllTagRequest;
};
export declare const QueryGetAllTagResponse: {
    encode(message: QueryGetAllTagResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetAllTagResponse;
    fromJSON(object: any): QueryGetAllTagResponse;
    toJSON(message: QueryGetAllTagResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetAllTagResponse>): QueryGetAllTagResponse;
};
export declare const QueryGetTagShaRequest: {
    encode(message: QueryGetTagShaRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetTagShaRequest;
    fromJSON(object: any): QueryGetTagShaRequest;
    toJSON(message: QueryGetTagShaRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetTagShaRequest>): QueryGetTagShaRequest;
};
export declare const QueryGetTagShaResponse: {
    encode(message: QueryGetTagShaResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetTagShaResponse;
    fromJSON(object: any): QueryGetTagShaResponse;
    toJSON(message: QueryGetTagShaResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetTagShaResponse>): QueryGetTagShaResponse;
};
export declare const QueryAllRepositoryRequest: {
    encode(message: QueryAllRepositoryRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllRepositoryRequest;
    fromJSON(object: any): QueryAllRepositoryRequest;
    toJSON(message: QueryAllRepositoryRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllRepositoryRequest>): QueryAllRepositoryRequest;
};
export declare const QueryAllRepositoryResponse: {
    encode(message: QueryAllRepositoryResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllRepositoryResponse;
    fromJSON(object: any): QueryAllRepositoryResponse;
    toJSON(message: QueryAllRepositoryResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllRepositoryResponse>): QueryAllRepositoryResponse;
};
export declare const QueryGetUserRequest: {
    encode(message: QueryGetUserRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetUserRequest;
    fromJSON(object: any): QueryGetUserRequest;
    toJSON(message: QueryGetUserRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetUserRequest>): QueryGetUserRequest;
};
export declare const QueryGetUserResponse: {
    encode(message: QueryGetUserResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetUserResponse;
    fromJSON(object: any): QueryGetUserResponse;
    toJSON(message: QueryGetUserResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetUserResponse>): QueryGetUserResponse;
};
export declare const QueryAllUserRequest: {
    encode(message: QueryAllUserRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllUserRequest;
    fromJSON(object: any): QueryAllUserRequest;
    toJSON(message: QueryAllUserRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllUserRequest>): QueryAllUserRequest;
};
export declare const QueryAllUserResponse: {
    encode(message: QueryAllUserResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllUserResponse;
    fromJSON(object: any): QueryAllUserResponse;
    toJSON(message: QueryAllUserResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllUserResponse>): QueryAllUserResponse;
};
export declare const QueryAllAddressRepositoryRequest: {
    encode(message: QueryAllAddressRepositoryRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllAddressRepositoryRequest;
    fromJSON(object: any): QueryAllAddressRepositoryRequest;
    toJSON(message: QueryAllAddressRepositoryRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllAddressRepositoryRequest>): QueryAllAddressRepositoryRequest;
};
export declare const QueryAllAddressRepositoryResponse: {
    encode(message: QueryAllAddressRepositoryResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllAddressRepositoryResponse;
    fromJSON(object: any): QueryAllAddressRepositoryResponse;
    toJSON(message: QueryAllAddressRepositoryResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllAddressRepositoryResponse>): QueryAllAddressRepositoryResponse;
};
export declare const QueryGetAddressRepositoryRequest: {
    encode(message: QueryGetAddressRepositoryRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetAddressRepositoryRequest;
    fromJSON(object: any): QueryGetAddressRepositoryRequest;
    toJSON(message: QueryGetAddressRepositoryRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetAddressRepositoryRequest>): QueryGetAddressRepositoryRequest;
};
export declare const QueryGetAddressRepositoryResponse: {
    encode(message: QueryGetAddressRepositoryResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetAddressRepositoryResponse;
    fromJSON(object: any): QueryGetAddressRepositoryResponse;
    toJSON(message: QueryGetAddressRepositoryResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetAddressRepositoryResponse>): QueryGetAddressRepositoryResponse;
};
export declare const QueryAllUserOrganizationRequest: {
    encode(message: QueryAllUserOrganizationRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllUserOrganizationRequest;
    fromJSON(object: any): QueryAllUserOrganizationRequest;
    toJSON(message: QueryAllUserOrganizationRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllUserOrganizationRequest>): QueryAllUserOrganizationRequest;
};
export declare const QueryAllUserOrganizationResponse: {
    encode(message: QueryAllUserOrganizationResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllUserOrganizationResponse;
    fromJSON(object: any): QueryAllUserOrganizationResponse;
    toJSON(message: QueryAllUserOrganizationResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllUserOrganizationResponse>): QueryAllUserOrganizationResponse;
};
export declare const QueryGetWhoisRequest: {
    encode(message: QueryGetWhoisRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetWhoisRequest;
    fromJSON(object: any): QueryGetWhoisRequest;
    toJSON(message: QueryGetWhoisRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetWhoisRequest>): QueryGetWhoisRequest;
};
export declare const QueryGetWhoisResponse: {
    encode(message: QueryGetWhoisResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetWhoisResponse;
    fromJSON(object: any): QueryGetWhoisResponse;
    toJSON(message: QueryGetWhoisResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetWhoisResponse>): QueryGetWhoisResponse;
};
export declare const QueryAllWhoisRequest: {
    encode(message: QueryAllWhoisRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllWhoisRequest;
    fromJSON(object: any): QueryAllWhoisRequest;
    toJSON(message: QueryAllWhoisRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllWhoisRequest>): QueryAllWhoisRequest;
};
export declare const QueryAllWhoisResponse: {
    encode(message: QueryAllWhoisResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllWhoisResponse;
    fromJSON(object: any): QueryAllWhoisResponse;
    toJSON(message: QueryAllWhoisResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllWhoisResponse>): QueryAllWhoisResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a request by id. */
    Request(request: QueryGetRequestRequest): Promise<QueryGetRequestResponse>;
    /** Queries a list of request items. */
    RequestAll(request: QueryAllRequestRequest): Promise<QueryAllRequestResponse>;
    /** Queries a release by id. */
    Release(request: QueryGetReleaseRequest): Promise<QueryGetReleaseResponse>;
    /** Queries a list of release items. */
    ReleaseAll(request: QueryAllReleaseRequest): Promise<QueryAllReleaseResponse>;
    /** Queries a pullRequest by id. */
    PullRequest(request: QueryGetPullRequestRequest): Promise<QueryGetPullRequestResponse>;
    /** Queries a list of pullRequest items. */
    PullRequestAll(request: QueryAllPullRequestRequest): Promise<QueryAllPullRequestResponse>;
    /** Queries a organization by id. */
    Organization(request: QueryGetOrganizationRequest): Promise<QueryGetOrganizationResponse>;
    /** Queries a list of organization items. */
    OrganizationAll(request: QueryAllOrganizationRequest): Promise<QueryAllOrganizationResponse>;
    /** Queries a comment by id. */
    Comment(request: QueryGetCommentRequest): Promise<QueryGetCommentResponse>;
    /** Queries a list of comment items. */
    CommentAll(request: QueryAllCommentRequest): Promise<QueryAllCommentResponse>;
    /** Queries a issue by id. */
    Issue(request: QueryGetIssueRequest): Promise<QueryGetIssueResponse>;
    /** Queries a list of issue items. */
    IssueAll(request: QueryAllIssueRequest): Promise<QueryAllIssueResponse>;
    RepositoryReleaseLatest(request: QueryGetLatestRepositoryReleaseRequest): Promise<QueryGetLatestRepositoryReleaseResponse>;
    RepositoryRelease(request: QueryGetRepositoryReleaseRequest): Promise<QueryGetRepositoryReleaseResponse>;
    RepositoryReleaseAll(request: QueryAllRepositoryReleaseRequest): Promise<QueryAllRepositoryReleaseResponse>;
    /** Queries a repository by id. */
    RepositoryIssue(request: QueryGetRepositoryIssueRequest): Promise<QueryGetRepositoryIssueResponse>;
    /** Queries a list of repository items. */
    RepositoryIssueAll(request: QueryAllRepositoryIssueRequest): Promise<QueryAllRepositoryIssueResponse>;
    /** Queries a repository pullRequest by id. */
    RepositoryPullRequest(request: QueryGetRepositoryPullRequestRequest): Promise<QueryGetRepositoryPullRequestResponse>;
    RepositoryPullRequestAll(request: QueryAllRepositoryPullRequestRequest): Promise<QueryAllRepositoryPullRequestResponse>;
    RepositoryRequestAll(request: QueryAllRepositoryRequestRequest): Promise<QueryAllRepositoryRequestResponse>;
    /** Queries a repository by id. */
    Repository(request: QueryGetRepositoryRequest): Promise<QueryGetRepositoryResponse>;
    /** Queries a list of repository items. */
    RepositoryAll(request: QueryAllRepositoryRequest): Promise<QueryAllRepositoryResponse>;
    /** Queries a repository forks by id. */
    ForkAll(request: QueryGetAllForkRequest): Promise<QueryGetAllForkResponse>;
    /** Queries a repository by id. */
    BranchAll(request: QueryGetAllBranchRequest): Promise<QueryGetAllBranchResponse>;
    BranchSha(request: QueryGetBranchShaRequest): Promise<QueryGetBranchShaResponse>;
    TagAll(request: QueryGetAllTagRequest): Promise<QueryGetAllTagResponse>;
    TagSha(request: QueryGetTagShaRequest): Promise<QueryGetTagShaResponse>;
    /** Queries a user by id. */
    User(request: QueryGetUserRequest): Promise<QueryGetUserResponse>;
    /** Queries a list of user items. */
    UserAll(request: QueryAllUserRequest): Promise<QueryAllUserResponse>;
    /** Queries a list of user repositories. */
    AddressRepositoryAll(request: QueryAllAddressRepositoryRequest): Promise<QueryAllAddressRepositoryResponse>;
    /** Queries a repository by user id and repository name */
    AddressRepository(request: QueryGetAddressRepositoryRequest): Promise<QueryGetAddressRepositoryResponse>;
    AddressRequestReceivedAll(request: QueryAllAddressRequestReceivedRequest): Promise<QueryAllAddressRequestReceivedResponse>;
    /** Queries a list of user Organizations. */
    UserOrganizationAll(request: QueryAllUserOrganizationRequest): Promise<QueryAllUserOrganizationResponse>;
    /** Queries a whois by id. */
    Whois(request: QueryGetWhoisRequest): Promise<QueryGetWhoisResponse>;
    /** Queries a list of whois items. */
    WhoisAll(request: QueryAllWhoisRequest): Promise<QueryAllWhoisResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Request(request: QueryGetRequestRequest): Promise<QueryGetRequestResponse>;
    RequestAll(request: QueryAllRequestRequest): Promise<QueryAllRequestResponse>;
    Release(request: QueryGetReleaseRequest): Promise<QueryGetReleaseResponse>;
    ReleaseAll(request: QueryAllReleaseRequest): Promise<QueryAllReleaseResponse>;
    PullRequest(request: QueryGetPullRequestRequest): Promise<QueryGetPullRequestResponse>;
    PullRequestAll(request: QueryAllPullRequestRequest): Promise<QueryAllPullRequestResponse>;
    Organization(request: QueryGetOrganizationRequest): Promise<QueryGetOrganizationResponse>;
    OrganizationAll(request: QueryAllOrganizationRequest): Promise<QueryAllOrganizationResponse>;
    Comment(request: QueryGetCommentRequest): Promise<QueryGetCommentResponse>;
    CommentAll(request: QueryAllCommentRequest): Promise<QueryAllCommentResponse>;
    Issue(request: QueryGetIssueRequest): Promise<QueryGetIssueResponse>;
    IssueAll(request: QueryAllIssueRequest): Promise<QueryAllIssueResponse>;
    RepositoryReleaseLatest(request: QueryGetLatestRepositoryReleaseRequest): Promise<QueryGetLatestRepositoryReleaseResponse>;
    RepositoryRelease(request: QueryGetRepositoryReleaseRequest): Promise<QueryGetRepositoryReleaseResponse>;
    RepositoryReleaseAll(request: QueryAllRepositoryReleaseRequest): Promise<QueryAllRepositoryReleaseResponse>;
    RepositoryIssue(request: QueryGetRepositoryIssueRequest): Promise<QueryGetRepositoryIssueResponse>;
    RepositoryIssueAll(request: QueryAllRepositoryIssueRequest): Promise<QueryAllRepositoryIssueResponse>;
    RepositoryPullRequest(request: QueryGetRepositoryPullRequestRequest): Promise<QueryGetRepositoryPullRequestResponse>;
    RepositoryPullRequestAll(request: QueryAllRepositoryPullRequestRequest): Promise<QueryAllRepositoryPullRequestResponse>;
    RepositoryRequestAll(request: QueryAllRepositoryRequestRequest): Promise<QueryAllRepositoryRequestResponse>;
    Repository(request: QueryGetRepositoryRequest): Promise<QueryGetRepositoryResponse>;
    RepositoryAll(request: QueryAllRepositoryRequest): Promise<QueryAllRepositoryResponse>;
    ForkAll(request: QueryGetAllForkRequest): Promise<QueryGetAllForkResponse>;
    BranchAll(request: QueryGetAllBranchRequest): Promise<QueryGetAllBranchResponse>;
    BranchSha(request: QueryGetBranchShaRequest): Promise<QueryGetBranchShaResponse>;
    TagAll(request: QueryGetAllTagRequest): Promise<QueryGetAllTagResponse>;
    TagSha(request: QueryGetTagShaRequest): Promise<QueryGetTagShaResponse>;
    User(request: QueryGetUserRequest): Promise<QueryGetUserResponse>;
    UserAll(request: QueryAllUserRequest): Promise<QueryAllUserResponse>;
    AddressRepositoryAll(request: QueryAllAddressRepositoryRequest): Promise<QueryAllAddressRepositoryResponse>;
    AddressRepository(request: QueryGetAddressRepositoryRequest): Promise<QueryGetAddressRepositoryResponse>;
    AddressRequestReceivedAll(request: QueryAllAddressRequestReceivedRequest): Promise<QueryAllAddressRequestReceivedResponse>;
    UserOrganizationAll(request: QueryAllUserOrganizationRequest): Promise<QueryAllUserOrganizationResponse>;
    Whois(request: QueryGetWhoisRequest): Promise<QueryGetWhoisResponse>;
    WhoisAll(request: QueryAllWhoisRequest): Promise<QueryAllWhoisResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
