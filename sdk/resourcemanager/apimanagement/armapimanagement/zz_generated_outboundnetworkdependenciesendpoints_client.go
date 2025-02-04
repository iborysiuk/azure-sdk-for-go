//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// OutboundNetworkDependenciesEndpointsClient contains the methods for the OutboundNetworkDependenciesEndpoints group.
// Don't use this type directly, use NewOutboundNetworkDependenciesEndpointsClient() instead.
type OutboundNetworkDependenciesEndpointsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewOutboundNetworkDependenciesEndpointsClient creates a new instance of OutboundNetworkDependenciesEndpointsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOutboundNetworkDependenciesEndpointsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OutboundNetworkDependenciesEndpointsClient, error) {
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
	client := &OutboundNetworkDependenciesEndpointsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// ListByService - Gets the network endpoints of all outbound dependencies of a ApiManagement service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// options - OutboundNetworkDependenciesEndpointsClientListByServiceOptions contains the optional parameters for the OutboundNetworkDependenciesEndpointsClient.ListByService
// method.
func (client *OutboundNetworkDependenciesEndpointsClient) ListByService(ctx context.Context, resourceGroupName string, serviceName string, options *OutboundNetworkDependenciesEndpointsClientListByServiceOptions) (OutboundNetworkDependenciesEndpointsClientListByServiceResponse, error) {
	req, err := client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return OutboundNetworkDependenciesEndpointsClientListByServiceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OutboundNetworkDependenciesEndpointsClientListByServiceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OutboundNetworkDependenciesEndpointsClientListByServiceResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByServiceHandleResponse(resp)
}

// listByServiceCreateRequest creates the ListByService request.
func (client *OutboundNetworkDependenciesEndpointsClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *OutboundNetworkDependenciesEndpointsClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/outboundNetworkDependenciesEndpoints"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *OutboundNetworkDependenciesEndpointsClient) listByServiceHandleResponse(resp *http.Response) (OutboundNetworkDependenciesEndpointsClientListByServiceResponse, error) {
	result := OutboundNetworkDependenciesEndpointsClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OutboundEnvironmentEndpointList); err != nil {
		return OutboundNetworkDependenciesEndpointsClientListByServiceResponse{}, err
	}
	return result, nil
}
