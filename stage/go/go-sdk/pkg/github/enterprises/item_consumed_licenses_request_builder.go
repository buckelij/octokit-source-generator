package enterprises

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemConsumedLicensesRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\consumed-licenses
type ItemConsumedLicensesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemConsumedLicensesRequestBuilderGetQueryParameters lists the license consumption information for all users, including those from connected servers, associated with an enterprise.The authenticated user must be an enterprise admin to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `read:enterprise` scope to use this endpoint.
type ItemConsumedLicensesRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// NewItemConsumedLicensesRequestBuilderInternal instantiates a new ItemConsumedLicensesRequestBuilder and sets the default values.
func NewItemConsumedLicensesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConsumedLicensesRequestBuilder) {
    m := &ItemConsumedLicensesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/consumed-licenses{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemConsumedLicensesRequestBuilder instantiates a new ItemConsumedLicensesRequestBuilder and sets the default values.
func NewItemConsumedLicensesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConsumedLicensesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemConsumedLicensesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the license consumption information for all users, including those from connected servers, associated with an enterprise.The authenticated user must be an enterprise admin to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `read:enterprise` scope to use this endpoint.
// returns a GetConsumedLicensesable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/enterprise-admin/license#list-enterprise-consumed-licenses
func (m *ItemConsumedLicensesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemConsumedLicensesRequestBuilderGetQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.GetConsumedLicensesable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateGetConsumedLicensesFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.GetConsumedLicensesable), nil
}
// ToGetRequestInformation lists the license consumption information for all users, including those from connected servers, associated with an enterprise.The authenticated user must be an enterprise admin to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `read:enterprise` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemConsumedLicensesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemConsumedLicensesRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemConsumedLicensesRequestBuilder when successful
func (m *ItemConsumedLicensesRequestBuilder) WithUrl(rawUrl string)(*ItemConsumedLicensesRequestBuilder) {
    return NewItemConsumedLicensesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
