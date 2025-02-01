package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
    ie70f628569da434b4caf11c9ce2faacf7fdb11804338e3c07c3d6a572da28c6d "github.com/octokit/go-sdk/pkg/github/orgs/item/auditlog"
)

// ItemAuditLogRequestBuilder builds and executes requests for operations under \orgs\{org}\audit-log
type ItemAuditLogRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAuditLogRequestBuilderGetQueryParameters gets the audit log for an organization. For more information, see "[Reviewing the audit log for your organization](https://docs.github.com/enterprise-cloud@latest//github/setting-up-and-managing-organizations-and-teams/reviewing-the-audit-log-for-your-organization)."By default, the response includes up to 30 events from the past three months. Use the `phrase` parameter to filter results and retrieve older events. For example, use the `phrase` parameter with the `created` qualifier to filter events based on when the events occurred. For more information, see "[Reviewing the audit log for your organization](https://docs.github.com/enterprise-cloud@latest//organizations/keeping-your-organization-secure/managing-security-settings-for-your-organization/reviewing-the-audit-log-for-your-organization#searching-the-audit-log)."Use pagination to retrieve fewer or more than 30 events. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/guides/using-pagination-in-the-rest-api)."This endpoint has a rate limit of 1,750 queries per hour per user and IP address.  If your integration receives a rate limit error (typically a 403 or 429 response), it should wait before making another request to the GitHub API. For more information, see "[Rate limits for the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/rate-limits-for-the-rest-api)" and "[Best practices for integrators](https://docs.github.com/enterprise-cloud@latest//rest/guides/best-practices-for-integrators)."The authenticated user must be an organization owner to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `read:audit_log` scope to use this endpoint.
type ItemAuditLogRequestBuilderGetQueryParameters struct {
    // A cursor, as given in the [Link header](https://docs.github.com/enterprise-cloud@latest//rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for events after this cursor.
    After *string `uriparametername:"after"`
    // A cursor, as given in the [Link header](https://docs.github.com/enterprise-cloud@latest//rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for events before this cursor.
    Before *string `uriparametername:"before"`
    // The event types to include:- `web` - returns web (non-Git) events.- `git` - returns Git events.- `all` - returns both web and Git events.The default is `web`.
    Include *ie70f628569da434b4caf11c9ce2faacf7fdb11804338e3c07c3d6a572da28c6d.GetIncludeQueryParameterType `uriparametername:"include"`
    // The order of audit log events. To list newest events first, specify `desc`. To list oldest events first, specify `asc`.The default is `desc`.
    Order *ie70f628569da434b4caf11c9ce2faacf7fdb11804338e3c07c3d6a572da28c6d.GetOrderQueryParameterType `uriparametername:"order"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // A search phrase. For more information, see [Searching the audit log](https://docs.github.com/enterprise-cloud@latest//github/setting-up-and-managing-organizations-and-teams/reviewing-the-audit-log-for-your-organization#searching-the-audit-log).
    Phrase *string `uriparametername:"phrase"`
}
// NewItemAuditLogRequestBuilderInternal instantiates a new ItemAuditLogRequestBuilder and sets the default values.
func NewItemAuditLogRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuditLogRequestBuilder) {
    m := &ItemAuditLogRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/audit-log{?after*,before*,include*,order*,per_page*,phrase*}", pathParameters),
    }
    return m
}
// NewItemAuditLogRequestBuilder instantiates a new ItemAuditLogRequestBuilder and sets the default values.
func NewItemAuditLogRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuditLogRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuditLogRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the audit log for an organization. For more information, see "[Reviewing the audit log for your organization](https://docs.github.com/enterprise-cloud@latest//github/setting-up-and-managing-organizations-and-teams/reviewing-the-audit-log-for-your-organization)."By default, the response includes up to 30 events from the past three months. Use the `phrase` parameter to filter results and retrieve older events. For example, use the `phrase` parameter with the `created` qualifier to filter events based on when the events occurred. For more information, see "[Reviewing the audit log for your organization](https://docs.github.com/enterprise-cloud@latest//organizations/keeping-your-organization-secure/managing-security-settings-for-your-organization/reviewing-the-audit-log-for-your-organization#searching-the-audit-log)."Use pagination to retrieve fewer or more than 30 events. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/guides/using-pagination-in-the-rest-api)."This endpoint has a rate limit of 1,750 queries per hour per user and IP address.  If your integration receives a rate limit error (typically a 403 or 429 response), it should wait before making another request to the GitHub API. For more information, see "[Rate limits for the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/rate-limits-for-the-rest-api)" and "[Best practices for integrators](https://docs.github.com/enterprise-cloud@latest//rest/guides/best-practices-for-integrators)."The authenticated user must be an organization owner to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `read:audit_log` scope to use this endpoint.
// returns a []AuditLogEventable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/orgs/orgs#get-the-audit-log-for-an-organization
func (m *ItemAuditLogRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemAuditLogRequestBuilderGetQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AuditLogEventable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateAuditLogEventFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AuditLogEventable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AuditLogEventable)
        }
    }
    return val, nil
}
// ToGetRequestInformation gets the audit log for an organization. For more information, see "[Reviewing the audit log for your organization](https://docs.github.com/enterprise-cloud@latest//github/setting-up-and-managing-organizations-and-teams/reviewing-the-audit-log-for-your-organization)."By default, the response includes up to 30 events from the past three months. Use the `phrase` parameter to filter results and retrieve older events. For example, use the `phrase` parameter with the `created` qualifier to filter events based on when the events occurred. For more information, see "[Reviewing the audit log for your organization](https://docs.github.com/enterprise-cloud@latest//organizations/keeping-your-organization-secure/managing-security-settings-for-your-organization/reviewing-the-audit-log-for-your-organization#searching-the-audit-log)."Use pagination to retrieve fewer or more than 30 events. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/guides/using-pagination-in-the-rest-api)."This endpoint has a rate limit of 1,750 queries per hour per user and IP address.  If your integration receives a rate limit error (typically a 403 or 429 response), it should wait before making another request to the GitHub API. For more information, see "[Rate limits for the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/rate-limits-for-the-rest-api)" and "[Best practices for integrators](https://docs.github.com/enterprise-cloud@latest//rest/guides/best-practices-for-integrators)."The authenticated user must be an organization owner to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `read:audit_log` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemAuditLogRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemAuditLogRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAuditLogRequestBuilder when successful
func (m *ItemAuditLogRequestBuilder) WithUrl(rawUrl string)(*ItemAuditLogRequestBuilder) {
    return NewItemAuditLogRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
