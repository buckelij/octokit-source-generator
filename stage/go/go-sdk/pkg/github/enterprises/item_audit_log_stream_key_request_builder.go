package enterprises

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemAuditLogStreamKeyRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\audit-log\stream-key
type ItemAuditLogStreamKeyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemAuditLogStreamKeyRequestBuilderInternal instantiates a new ItemAuditLogStreamKeyRequestBuilder and sets the default values.
func NewItemAuditLogStreamKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuditLogStreamKeyRequestBuilder) {
    m := &ItemAuditLogStreamKeyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/audit-log/stream-key", pathParameters),
    }
    return m
}
// NewItemAuditLogStreamKeyRequestBuilder instantiates a new ItemAuditLogStreamKeyRequestBuilder and sets the default values.
func NewItemAuditLogStreamKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuditLogStreamKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuditLogStreamKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieves the audit log streaming public key for encrypting secrets.When using this endpoint, you must encrypt the credentials following the same encryption steps as outlined in the guide on encrypting secrets. See "[Encrypting secrets for the REST API](/rest/guides/encrypting-secrets-for-the-rest-api)."
// returns a AuditLogStreamKeyable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/enterprise-admin/audit-log#get-the-audit-log-stream-key-for-encrypting-secrets
func (m *ItemAuditLogStreamKeyRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AuditLogStreamKeyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateAuditLogStreamKeyFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AuditLogStreamKeyable), nil
}
// ToGetRequestInformation retrieves the audit log streaming public key for encrypting secrets.When using this endpoint, you must encrypt the credentials following the same encryption steps as outlined in the guide on encrypting secrets. See "[Encrypting secrets for the REST API](/rest/guides/encrypting-secrets-for-the-rest-api)."
// returns a *RequestInformation when successful
func (m *ItemAuditLogStreamKeyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAuditLogStreamKeyRequestBuilder when successful
func (m *ItemAuditLogStreamKeyRequestBuilder) WithUrl(rawUrl string)(*ItemAuditLogStreamKeyRequestBuilder) {
    return NewItemAuditLogStreamKeyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
