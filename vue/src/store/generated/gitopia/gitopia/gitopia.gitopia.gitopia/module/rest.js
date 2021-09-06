/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */
export var OrganizationMemberRole;
(function (OrganizationMemberRole) {
    OrganizationMemberRole["MEMBER"] = "MEMBER";
    OrganizationMemberRole["OWNER"] = "OWNER";
})(OrganizationMemberRole || (OrganizationMemberRole = {}));
export var RepositoryCollaboratorPermission;
(function (RepositoryCollaboratorPermission) {
    RepositoryCollaboratorPermission["READ"] = "READ";
    RepositoryCollaboratorPermission["TRIAGE"] = "TRIAGE";
    RepositoryCollaboratorPermission["WRITE"] = "WRITE";
    RepositoryCollaboratorPermission["MAINTAIN"] = "MAINTAIN";
    RepositoryCollaboratorPermission["ADMIN"] = "ADMIN";
})(RepositoryCollaboratorPermission || (RepositoryCollaboratorPermission = {}));
export var GitopiaCommentType;
(function (GitopiaCommentType) {
    GitopiaCommentType["ISSUE"] = "ISSUE";
    GitopiaCommentType["PULLREQUEST"] = "PULLREQUEST";
})(GitopiaCommentType || (GitopiaCommentType = {}));
export var GitopiaIssueState;
(function (GitopiaIssueState) {
    GitopiaIssueState["OPEN"] = "OPEN";
    GitopiaIssueState["CLOSED"] = "CLOSED";
})(GitopiaIssueState || (GitopiaIssueState = {}));
export var GitopiaPullRequestState;
(function (GitopiaPullRequestState) {
    GitopiaPullRequestState["OPEN"] = "OPEN";
    GitopiaPullRequestState["CLOSED"] = "CLOSED";
    GitopiaPullRequestState["MERGED"] = "MERGED";
})(GitopiaPullRequestState || (GitopiaPullRequestState = {}));
export var GitopiaRepositoryOwnerType;
(function (GitopiaRepositoryOwnerType) {
    GitopiaRepositoryOwnerType["USER"] = "USER";
    GitopiaRepositoryOwnerType["ORGANIZATION"] = "ORGANIZATION";
})(GitopiaRepositoryOwnerType || (GitopiaRepositoryOwnerType = {}));
export var ContentType;
(function (ContentType) {
    ContentType["Json"] = "application/json";
    ContentType["FormData"] = "multipart/form-data";
    ContentType["UrlEncoded"] = "application/x-www-form-urlencoded";
})(ContentType || (ContentType = {}));
export class HttpClient {
    constructor(apiConfig = {}) {
        this.baseUrl = "";
        this.securityData = null;
        this.securityWorker = null;
        this.abortControllers = new Map();
        this.baseApiParams = {
            credentials: "same-origin",
            headers: {},
            redirect: "follow",
            referrerPolicy: "no-referrer",
        };
        this.setSecurityData = (data) => {
            this.securityData = data;
        };
        this.contentFormatters = {
            [ContentType.Json]: (input) => input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
            [ContentType.FormData]: (input) => Object.keys(input || {}).reduce((data, key) => {
                data.append(key, input[key]);
                return data;
            }, new FormData()),
            [ContentType.UrlEncoded]: (input) => this.toQueryString(input),
        };
        this.createAbortSignal = (cancelToken) => {
            if (this.abortControllers.has(cancelToken)) {
                const abortController = this.abortControllers.get(cancelToken);
                if (abortController) {
                    return abortController.signal;
                }
                return void 0;
            }
            const abortController = new AbortController();
            this.abortControllers.set(cancelToken, abortController);
            return abortController.signal;
        };
        this.abortRequest = (cancelToken) => {
            const abortController = this.abortControllers.get(cancelToken);
            if (abortController) {
                abortController.abort();
                this.abortControllers.delete(cancelToken);
            }
        };
        this.request = ({ body, secure, path, type, query, format = "json", baseUrl, cancelToken, ...params }) => {
            const secureParams = (secure && this.securityWorker && this.securityWorker(this.securityData)) || {};
            const requestParams = this.mergeRequestParams(params, secureParams);
            const queryString = query && this.toQueryString(query);
            const payloadFormatter = this.contentFormatters[type || ContentType.Json];
            return fetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
                ...requestParams,
                headers: {
                    ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
                    ...(requestParams.headers || {}),
                },
                signal: cancelToken ? this.createAbortSignal(cancelToken) : void 0,
                body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
            }).then(async (response) => {
                const r = response;
                r.data = null;
                r.error = null;
                const data = await response[format]()
                    .then((data) => {
                    if (r.ok) {
                        r.data = data;
                    }
                    else {
                        r.error = data;
                    }
                    return r;
                })
                    .catch((e) => {
                    r.error = e;
                    return r;
                });
                if (cancelToken) {
                    this.abortControllers.delete(cancelToken);
                }
                if (!response.ok)
                    throw data;
                return data;
            });
        };
        Object.assign(this, apiConfig);
    }
    addQueryParam(query, key) {
        const value = query[key];
        return (encodeURIComponent(key) +
            "=" +
            encodeURIComponent(Array.isArray(value) ? value.join(",") : typeof value === "number" ? value : `${value}`));
    }
    toQueryString(rawQuery) {
        const query = rawQuery || {};
        const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
        return keys
            .map((key) => typeof query[key] === "object" && !Array.isArray(query[key])
            ? this.toQueryString(query[key])
            : this.addQueryParam(query, key))
            .join("&");
    }
    addQueryParams(rawQuery) {
        const queryString = this.toQueryString(rawQuery);
        return queryString ? `?${queryString}` : "";
    }
    mergeRequestParams(params1, params2) {
        return {
            ...this.baseApiParams,
            ...params1,
            ...(params2 || {}),
            headers: {
                ...(this.baseApiParams.headers || {}),
                ...(params1.headers || {}),
                ...((params2 && params2.headers) || {}),
            },
        };
    }
}
/**
 * @title gitopia/comment.proto
 * @version version not set
 */
export class Api extends HttpClient {
    constructor() {
        super(...arguments);
        /**
         * No description
         *
         * @tags Query
         * @name QueryCommentAll
         * @summary Queries a list of comment items.
         * @request GET:/gitopia/gitopia/gitopia/comment
         */
        this.queryCommentAll = (query, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/comment`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryComment
         * @summary Queries a comment by id.
         * @request GET:/gitopia/gitopia/gitopia/comment/{id}
         */
        this.queryComment = (id, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/comment/${id}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryIssueAll
         * @summary Queries a list of issue items.
         * @request GET:/gitopia/gitopia/gitopia/issue
         */
        this.queryIssueAll = (query, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/issue`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryIssue
         * @summary Queries a issue by id.
         * @request GET:/gitopia/gitopia/gitopia/issue/{id}
         */
        this.queryIssue = (id, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/issue/${id}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryOrganizationAll
         * @summary Queries a list of organization items.
         * @request GET:/gitopia/gitopia/gitopia/organization
         */
        this.queryOrganizationAll = (query, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/organization`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryOrganizationByName
         * @summary Queries a organization by name.
         * @request GET:/gitopia/gitopia/gitopia/organization/name/{organizationName}
         */
        this.queryOrganizationByName = (organizationName, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/organization/name/${organizationName}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryOrganization
         * @summary Queries a organization by id.
         * @request GET:/gitopia/gitopia/gitopia/organization/{id}
         */
        this.queryOrganization = (id, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/organization/${id}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryOrganizationRepositoryAll
         * @summary Queries a list of Organization repositories.
         * @request GET:/gitopia/gitopia/gitopia/organization/{organizationName}/repositories
         */
        this.queryOrganizationRepositoryAll = (organizationName, query, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/organization/${organizationName}/repositories`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryOrganizationRepository
         * @summary Queries a repository by Organization name and repository name
         * @request GET:/gitopia/gitopia/gitopia/organization/{organizationName}/repositories/{repositoryName}
         */
        this.queryOrganizationRepository = (organizationName, repositoryName, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/organization/${organizationName}/repositories/${repositoryName}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryPullRequestAll
         * @summary Queries a list of pullRequest items.
         * @request GET:/gitopia/gitopia/gitopia/pullRequest
         */
        this.queryPullRequestAll = (query, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/pullRequest`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryPullRequest
         * @summary Queries a pullRequest by id.
         * @request GET:/gitopia/gitopia/gitopia/pullRequest/{id}
         */
        this.queryPullRequest = (id, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/pullRequest/${id}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryRepositoryAll
         * @summary Queries a list of repository items.
         * @request GET:/gitopia/gitopia/gitopia/repository
         */
        this.queryRepositoryAll = (query, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/repository`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryRepository
         * @summary Queries a repository by id.
         * @request GET:/gitopia/gitopia/gitopia/repository/{id}
         */
        this.queryRepository = (id, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/repository/${id}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryBranchAll
         * @summary Queries a repository by id.
         * @request GET:/gitopia/gitopia/gitopia/repository/{id}/branches
         */
        this.queryBranchAll = (id, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/repository/${id}/branches`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryBranchSha
         * @request GET:/gitopia/gitopia/gitopia/repository/{repositoryId}/branches/{branchName}
         */
        this.queryBranchSha = (repositoryId, branchName, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/repository/${repositoryId}/branches/${branchName}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryTagAll
         * @request GET:/gitopia/gitopia/gitopia/repository/{repositoryId}/tags
         */
        this.queryTagAll = (repositoryId, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/repository/${repositoryId}/tags`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryTagSha
         * @request GET:/gitopia/gitopia/gitopia/repository/{repositoryId}/tags/{tagName}
         */
        this.queryTagSha = (repositoryId, tagName, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/repository/${repositoryId}/tags/${tagName}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryUserRepository
         * @summary Queries a repository by user id and repository name
         * @request GET:/gitopia/gitopia/gitopia/repository/{userId}/{repositoryName}
         */
        this.queryUserRepository = (userId, repositoryName, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/repository/${userId}/${repositoryName}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryUserAll
         * @summary Queries a list of user items.
         * @request GET:/gitopia/gitopia/gitopia/user
         */
        this.queryUserAll = (query, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/user`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryUser
         * @summary Queries a user by id.
         * @request GET:/gitopia/gitopia/gitopia/user/{id}
         */
        this.queryUser = (id, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/user/${id}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryUserOrganizationAll
         * @summary Queries a list of user Organizations.
         * @request GET:/gitopia/gitopia/gitopia/user/{id}/organizations
         */
        this.queryUserOrganizationAll = (id, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/user/${id}/organizations`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryUserRepositoryAll
         * @summary Queries a list of user repositories.
         * @request GET:/gitopia/gitopia/gitopia/user/{id}/repositories
         */
        this.queryUserRepositoryAll = (id, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/user/${id}/repositories`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryWhoisAll
         * @summary Queries a list of whois items.
         * @request GET:/gitopia/gitopia/gitopia/whois
         */
        this.queryWhoisAll = (query, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/whois`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryWhois
         * @summary Queries a whois by id.
         * @request GET:/gitopia/gitopia/gitopia/whois/{name}
         */
        this.queryWhois = (name, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/whois/${name}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryRepositoryIssueAll
         * @summary Queries a list of repository items.
         * @request GET:/gitopia/gitopia/gitopia/{userId}/{repositoryName}/issue
         */
        this.queryRepositoryIssueAll = (userId, repositoryName, query, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/${userId}/${repositoryName}/issue`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryRepositoryIssue
         * @summary Queries a repository by id.
         * @request GET:/gitopia/gitopia/gitopia/{userId}/{repositoryName}/issue/{issueIid}
         */
        this.queryRepositoryIssue = (userId, repositoryName, issueIid, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/${userId}/${repositoryName}/issue/${issueIid}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryRepositoryPullRequestAll
         * @request GET:/gitopia/gitopia/gitopia/{userId}/{repositoryName}/pull
         */
        this.queryRepositoryPullRequestAll = (userId, repositoryName, query, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/${userId}/${repositoryName}/pull`,
            method: "GET",
            query: query,
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryRepositoryPullRequest
         * @summary Queries a repository pullRequest by id.
         * @request GET:/gitopia/gitopia/gitopia/{userId}/{repositoryName}/pull/{pullIid}
         */
        this.queryRepositoryPullRequest = (userId, repositoryName, pullIid, params = {}) => this.request({
            path: `/gitopia/gitopia/gitopia/${userId}/${repositoryName}/pull/${pullIid}`,
            method: "GET",
            format: "json",
            ...params,
        });
    }
}
