import { Writer, Reader } from "protobufjs/minimal";
import { Request } from "../gitopia/request";
import { Release } from "../gitopia/release";
import { PullRequest } from "../gitopia/pullRequest";
import { Organization } from "../gitopia/organization";
import { Comment } from "../gitopia/comment";
import { Issue } from "../gitopia/issue";
import { Repository } from "../gitopia/repository";
import { User } from "../gitopia/user";
import { Whois } from "../gitopia/whois";
export declare const protobufPackage = "gitopia.gitopia.gitopia";
/** GenesisState defines the gitopia module's genesis state. */
export interface GenesisState {
    /** this line is used by starport scaffolding # genesis/proto/state */
    requestList: Request[];
    requestCount: number;
    releaseList: Release[];
    releaseCount: number;
    pullRequestList: PullRequest[];
    pullRequestCount: number;
    organizationList: Organization[];
    organizationCount: number;
    commentList: Comment[];
    commentCount: number;
    issueList: Issue[];
    issueCount: number;
    repositoryList: Repository[];
    repositoryCount: number;
    userList: User[];
    userCount: number;
    whoisList: Whois[];
    /** this line is used by starport scaffolding # ibc/genesis/proto */
    whoisCount: number;
}
export declare const GenesisState: {
    encode(message: GenesisState, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): GenesisState;
    fromJSON(object: any): GenesisState;
    toJSON(message: GenesisState): unknown;
    fromPartial(object: DeepPartial<GenesisState>): GenesisState;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
