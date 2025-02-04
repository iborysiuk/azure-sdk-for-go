//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomerinsights

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ProfilesClient contains the methods for the Profiles group.
// Don't use this type directly, use NewProfilesClient() instead.
type ProfilesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewProfilesClient creates a new instance of ProfilesClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewProfilesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProfilesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ProfilesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a profile within a Hub, or updates an existing profile.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// profileName - The name of the profile.
// parameters - Parameters supplied to the create/delete Profile type operation
// options - ProfilesClientBeginCreateOrUpdateOptions contains the optional parameters for the ProfilesClient.BeginCreateOrUpdate
// method.
func (client *ProfilesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, profileName string, parameters ProfileResourceFormat, options *ProfilesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ProfilesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, hubName, profileName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ProfilesClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ProfilesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates a profile within a Hub, or updates an existing profile.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
func (client *ProfilesClient) createOrUpdate(ctx context.Context, resourceGroupName string, hubName string, profileName string, parameters ProfileResourceFormat, options *ProfilesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, hubName, profileName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ProfilesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, hubName string, profileName string, parameters ProfileResourceFormat, options *ProfilesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles/{profileName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a profile within a hub
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// profileName - The name of the profile.
// options - ProfilesClientBeginDeleteOptions contains the optional parameters for the ProfilesClient.BeginDelete method.
func (client *ProfilesClient) BeginDelete(ctx context.Context, resourceGroupName string, hubName string, profileName string, options *ProfilesClientBeginDeleteOptions) (*runtime.Poller[ProfilesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, hubName, profileName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ProfilesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ProfilesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a profile within a hub
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
func (client *ProfilesClient) deleteOperation(ctx context.Context, resourceGroupName string, hubName string, profileName string, options *ProfilesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, hubName, profileName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ProfilesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, hubName string, profileName string, options *ProfilesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles/{profileName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.LocaleCode != nil {
		reqQP.Set("locale-code", *options.LocaleCode)
	}
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets information about the specified profile.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// profileName - The name of the profile.
// options - ProfilesClientGetOptions contains the optional parameters for the ProfilesClient.Get method.
func (client *ProfilesClient) Get(ctx context.Context, resourceGroupName string, hubName string, profileName string, options *ProfilesClientGetOptions) (ProfilesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, hubName, profileName, options)
	if err != nil {
		return ProfilesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProfilesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProfilesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ProfilesClient) getCreateRequest(ctx context.Context, resourceGroupName string, hubName string, profileName string, options *ProfilesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles/{profileName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.LocaleCode != nil {
		reqQP.Set("locale-code", *options.LocaleCode)
	}
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProfilesClient) getHandleResponse(resp *http.Response) (ProfilesClientGetResponse, error) {
	result := ProfilesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProfileResourceFormat); err != nil {
		return ProfilesClientGetResponse{}, err
	}
	return result, nil
}

// GetEnrichingKpis - Gets the KPIs that enrich the profile Type identified by the supplied name. Enrichment happens through
// participants of the Interaction on an Interaction KPI and through Relationships for Profile KPIs.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// profileName - The name of the profile.
// options - ProfilesClientGetEnrichingKpisOptions contains the optional parameters for the ProfilesClient.GetEnrichingKpis
// method.
func (client *ProfilesClient) GetEnrichingKpis(ctx context.Context, resourceGroupName string, hubName string, profileName string, options *ProfilesClientGetEnrichingKpisOptions) (ProfilesClientGetEnrichingKpisResponse, error) {
	req, err := client.getEnrichingKpisCreateRequest(ctx, resourceGroupName, hubName, profileName, options)
	if err != nil {
		return ProfilesClientGetEnrichingKpisResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProfilesClientGetEnrichingKpisResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProfilesClientGetEnrichingKpisResponse{}, runtime.NewResponseError(resp)
	}
	return client.getEnrichingKpisHandleResponse(resp)
}

// getEnrichingKpisCreateRequest creates the GetEnrichingKpis request.
func (client *ProfilesClient) getEnrichingKpisCreateRequest(ctx context.Context, resourceGroupName string, hubName string, profileName string, options *ProfilesClientGetEnrichingKpisOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles/{profileName}/getEnrichingKpis"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEnrichingKpisHandleResponse handles the GetEnrichingKpis response.
func (client *ProfilesClient) getEnrichingKpisHandleResponse(resp *http.Response) (ProfilesClientGetEnrichingKpisResponse, error) {
	result := ProfilesClientGetEnrichingKpisResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KpiDefinitionArray); err != nil {
		return ProfilesClientGetEnrichingKpisResponse{}, err
	}
	return result, nil
}

// NewListByHubPager - Gets all profile in the hub.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// options - ProfilesClientListByHubOptions contains the optional parameters for the ProfilesClient.ListByHub method.
func (client *ProfilesClient) NewListByHubPager(resourceGroupName string, hubName string, options *ProfilesClientListByHubOptions) *runtime.Pager[ProfilesClientListByHubResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProfilesClientListByHubResponse]{
		More: func(page ProfilesClientListByHubResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProfilesClientListByHubResponse) (ProfilesClientListByHubResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByHubCreateRequest(ctx, resourceGroupName, hubName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ProfilesClientListByHubResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ProfilesClientListByHubResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProfilesClientListByHubResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByHubHandleResponse(resp)
		},
	})
}

// listByHubCreateRequest creates the ListByHub request.
func (client *ProfilesClient) listByHubCreateRequest(ctx context.Context, resourceGroupName string, hubName string, options *ProfilesClientListByHubOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.LocaleCode != nil {
		reqQP.Set("locale-code", *options.LocaleCode)
	}
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByHubHandleResponse handles the ListByHub response.
func (client *ProfilesClient) listByHubHandleResponse(resp *http.Response) (ProfilesClientListByHubResponse, error) {
	result := ProfilesClientListByHubResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProfileListResult); err != nil {
		return ProfilesClientListByHubResponse{}, err
	}
	return result, nil
}
