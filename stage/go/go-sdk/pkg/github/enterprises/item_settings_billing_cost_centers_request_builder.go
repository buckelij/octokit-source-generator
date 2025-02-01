package enterprises

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemSettingsBillingCostCentersRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\settings\billing\cost-centers
type ItemSettingsBillingCostCentersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByCost_center_id gets an item from the github.com/octokit/go-sdk/pkg/github.enterprises.item.settings.billing.costCenters.item collection
// returns a *ItemSettingsBillingCostCentersWithCost_center_ItemRequestBuilder when successful
func (m *ItemSettingsBillingCostCentersRequestBuilder) ByCost_center_id(cost_center_id string)(*ItemSettingsBillingCostCentersWithCost_center_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if cost_center_id != "" {
        urlTplParams["cost_center_id"] = cost_center_id
    }
    return NewItemSettingsBillingCostCentersWithCost_center_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSettingsBillingCostCentersRequestBuilderInternal instantiates a new ItemSettingsBillingCostCentersRequestBuilder and sets the default values.
func NewItemSettingsBillingCostCentersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsBillingCostCentersRequestBuilder) {
    m := &ItemSettingsBillingCostCentersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/settings/billing/cost-centers", pathParameters),
    }
    return m
}
// NewItemSettingsBillingCostCentersRequestBuilder instantiates a new ItemSettingsBillingCostCentersRequestBuilder and sets the default values.
func NewItemSettingsBillingCostCentersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsBillingCostCentersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSettingsBillingCostCentersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a list of all the cost centers for an enterprise.
// returns a GetAllCostCentersable when successful
// returns a BasicError error when the service returns a 400 status code
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 500 status code
// returns a GetAllCostCenters503Error error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/enterprise-admin/billing#get-all-cost-centers-for-an-enterprise
func (m *ItemSettingsBillingCostCentersRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.GetAllCostCentersable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "500": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "503": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateGetAllCostCenters503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateGetAllCostCentersFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.GetAllCostCentersable), nil
}
// ToGetRequestInformation gets a list of all the cost centers for an enterprise.
// returns a *RequestInformation when successful
func (m *ItemSettingsBillingCostCentersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSettingsBillingCostCentersRequestBuilder when successful
func (m *ItemSettingsBillingCostCentersRequestBuilder) WithUrl(rawUrl string)(*ItemSettingsBillingCostCentersRequestBuilder) {
    return NewItemSettingsBillingCostCentersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
