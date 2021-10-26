//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// OriginGroupsClient contains the methods for the OriginGroups group.
// Don't use this type directly, use NewOriginGroupsClient() instead.
type OriginGroupsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewOriginGroupsClient creates a new instance of OriginGroupsClient with the specified values.
func NewOriginGroupsClient(con *arm.Connection, subscriptionID string) *OriginGroupsClient {
	return &OriginGroupsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreate - Creates a new origin group within the specified endpoint.
// If the operation fails it returns the *ErrorResponse error type.
func (client *OriginGroupsClient) BeginCreate(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string, originGroup OriginGroup, options *OriginGroupsBeginCreateOptions) (OriginGroupsCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, profileName, endpointName, originGroupName, originGroup, options)
	if err != nil {
		return OriginGroupsCreatePollerResponse{}, err
	}
	result := OriginGroupsCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("OriginGroupsClient.Create", "", resp, client.pl, client.createHandleError)
	if err != nil {
		return OriginGroupsCreatePollerResponse{}, err
	}
	result.Poller = &OriginGroupsCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Creates a new origin group within the specified endpoint.
// If the operation fails it returns the *ErrorResponse error type.
func (client *OriginGroupsClient) create(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string, originGroup OriginGroup, options *OriginGroupsBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, profileName, endpointName, originGroupName, originGroup, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *OriginGroupsClient) createCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string, originGroup OriginGroup, options *OriginGroupsBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/originGroups/{originGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if originGroupName == "" {
		return nil, errors.New("parameter originGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originGroupName}", url.PathEscape(originGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, originGroup)
}

// createHandleError handles the Create error response.
func (client *OriginGroupsClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes an existing origin group within an endpoint.
// If the operation fails it returns the *ErrorResponse error type.
func (client *OriginGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string, options *OriginGroupsBeginDeleteOptions) (OriginGroupsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, profileName, endpointName, originGroupName, options)
	if err != nil {
		return OriginGroupsDeletePollerResponse{}, err
	}
	result := OriginGroupsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("OriginGroupsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return OriginGroupsDeletePollerResponse{}, err
	}
	result.Poller = &OriginGroupsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes an existing origin group within an endpoint.
// If the operation fails it returns the *ErrorResponse error type.
func (client *OriginGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string, options *OriginGroupsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, profileName, endpointName, originGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *OriginGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string, options *OriginGroupsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/originGroups/{originGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if originGroupName == "" {
		return nil, errors.New("parameter originGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originGroupName}", url.PathEscape(originGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *OriginGroupsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets an existing origin group within an endpoint.
// If the operation fails it returns the *ErrorResponse error type.
func (client *OriginGroupsClient) Get(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string, options *OriginGroupsGetOptions) (OriginGroupsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, profileName, endpointName, originGroupName, options)
	if err != nil {
		return OriginGroupsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OriginGroupsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OriginGroupsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OriginGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string, options *OriginGroupsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/originGroups/{originGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if originGroupName == "" {
		return nil, errors.New("parameter originGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originGroupName}", url.PathEscape(originGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OriginGroupsClient) getHandleResponse(resp *http.Response) (OriginGroupsGetResponse, error) {
	result := OriginGroupsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.OriginGroup); err != nil {
		return OriginGroupsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *OriginGroupsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByEndpoint - Lists all of the existing origin groups within an endpoint.
// If the operation fails it returns the *ErrorResponse error type.
func (client *OriginGroupsClient) ListByEndpoint(resourceGroupName string, profileName string, endpointName string, options *OriginGroupsListByEndpointOptions) *OriginGroupsListByEndpointPager {
	return &OriginGroupsListByEndpointPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByEndpointCreateRequest(ctx, resourceGroupName, profileName, endpointName, options)
		},
		advancer: func(ctx context.Context, resp OriginGroupsListByEndpointResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.OriginGroupListResult.NextLink)
		},
	}
}

// listByEndpointCreateRequest creates the ListByEndpoint request.
func (client *OriginGroupsClient) listByEndpointCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, options *OriginGroupsListByEndpointOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/originGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByEndpointHandleResponse handles the ListByEndpoint response.
func (client *OriginGroupsClient) listByEndpointHandleResponse(resp *http.Response) (OriginGroupsListByEndpointResponse, error) {
	result := OriginGroupsListByEndpointResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.OriginGroupListResult); err != nil {
		return OriginGroupsListByEndpointResponse{}, err
	}
	return result, nil
}

// listByEndpointHandleError handles the ListByEndpoint error response.
func (client *OriginGroupsClient) listByEndpointHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - Updates an existing origin group within an endpoint.
// If the operation fails it returns the *ErrorResponse error type.
func (client *OriginGroupsClient) BeginUpdate(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string, originGroupUpdateProperties OriginGroupUpdateParameters, options *OriginGroupsBeginUpdateOptions) (OriginGroupsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, profileName, endpointName, originGroupName, originGroupUpdateProperties, options)
	if err != nil {
		return OriginGroupsUpdatePollerResponse{}, err
	}
	result := OriginGroupsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("OriginGroupsClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return OriginGroupsUpdatePollerResponse{}, err
	}
	result.Poller = &OriginGroupsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates an existing origin group within an endpoint.
// If the operation fails it returns the *ErrorResponse error type.
func (client *OriginGroupsClient) update(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string, originGroupUpdateProperties OriginGroupUpdateParameters, options *OriginGroupsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, profileName, endpointName, originGroupName, originGroupUpdateProperties, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *OriginGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string, originGroupUpdateProperties OriginGroupUpdateParameters, options *OriginGroupsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/originGroups/{originGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if originGroupName == "" {
		return nil, errors.New("parameter originGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originGroupName}", url.PathEscape(originGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, originGroupUpdateProperties)
}

// updateHandleError handles the Update error response.
func (client *OriginGroupsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}