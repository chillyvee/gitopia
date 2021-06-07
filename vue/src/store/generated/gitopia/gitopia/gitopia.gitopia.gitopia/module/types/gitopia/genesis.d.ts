import { Comment } from "../gitopia/comment";
import { Issue } from "../gitopia/issue";
import { Repository } from "../gitopia/repository";
import { User } from "../gitopia/user";
import { Whois } from "../gitopia/whois";
import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "gitopia.gitopia.gitopia";
/** GenesisState defines the capability module's genesis state. */
export interface GenesisState {
    /** this line is used by starport scaffolding # genesis/proto/state */
    commentList: Comment[];
    /** this line is used by starport scaffolding # genesis/proto/stateField */
    issueList: Issue[];
    /** this line is used by starport scaffolding # genesis/proto/stateField */
    repositoryList: Repository[];
    /** this line is used by starport scaffolding # genesis/proto/stateField */
    userList: User[];
    /** this line is used by starport scaffolding # genesis/proto/stateField */
    whoisList: Whois[];
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
