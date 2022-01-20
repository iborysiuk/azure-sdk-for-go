package labservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// LabsClient is the client for the Labs methods of the Labservices service.
type LabsClient struct {
	BaseClient
}

// NewLabsClient creates an instance of the LabsClient client.
func NewLabsClient(subscriptionID string) LabsClient {
	return NewLabsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLabsClientWithBaseURI creates an instance of the LabsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewLabsClientWithBaseURI(baseURI string, subscriptionID string) LabsClient {
	return LabsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate operation to create or update a lab resource.
// Parameters:
// body - the request body.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// labName - the name of the lab that uniquely identifies it within containing lab account. Used in resource
// URIs.
func (client LabsClient) CreateOrUpdate(ctx context.Context, body Lab, resourceGroupName string, labName string) (result LabsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.LabProperties", Name: validation.Null, Rule: true, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: labName,
			Constraints: []validation.Constraint{{Target: "labName", Name: validation.MaxLength, Rule: 100, Chain: nil},
				{Target: "labName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("labservices.LabsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, body, resourceGroupName, labName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client LabsClient) CreateOrUpdatePreparer(ctx context.Context, body Lab, resourceGroupName string, labName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	body.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) CreateOrUpdateSender(req *http.Request) (future LabsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client LabsClient) CreateOrUpdateResponder(resp *http.Response) (result Lab, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete operation to delete a lab resource.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// labName - the name of the lab that uniquely identifies it within containing lab account. Used in resource
// URIs.
func (client LabsClient) Delete(ctx context.Context, resourceGroupName string, labName string) (result LabsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: labName,
			Constraints: []validation.Constraint{{Target: "labName", Name: validation.MaxLength, Rule: 100, Chain: nil},
				{Target: "labName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("labservices.LabsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, labName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client LabsClient) DeletePreparer(ctx context.Context, resourceGroupName string, labName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) DeleteSender(req *http.Request) (future LabsDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client LabsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns the properties of a lab resource.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// labName - the name of the lab that uniquely identifies it within containing lab account. Used in resource
// URIs.
func (client LabsClient) Get(ctx context.Context, resourceGroupName string, labName string) (result Lab, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: labName,
			Constraints: []validation.Constraint{{Target: "labName", Name: validation.MaxLength, Rule: 100, Chain: nil},
				{Target: "labName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("labservices.LabsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, labName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client LabsClient) GetPreparer(ctx context.Context, resourceGroupName string, labName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client LabsClient) GetResponder(resp *http.Response) (result Lab, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup returns a list of all labs in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
func (client LabsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result PagedLabsPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.pl.Response.Response != nil {
				sc = result.pl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("labservices.LabsClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.pl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.pl, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.pl.hasNextLink() && result.pl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client LabsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client LabsClient) ListByResourceGroupResponder(resp *http.Response) (result PagedLabs, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client LabsClient) listByResourceGroupNextResults(ctx context.Context, lastResults PagedLabs) (result PagedLabs, err error) {
	req, err := lastResults.pagedLabsPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "labservices.LabsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "labservices.LabsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client LabsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result PagedLabsIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// ListBySubscription returns a list of all labs for a subscription.
// Parameters:
// filter - the filter to apply to the operation.
func (client LabsClient) ListBySubscription(ctx context.Context, filter string) (result PagedLabsPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.pl.Response.Response != nil {
				sc = result.pl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("labservices.LabsClient", "ListBySubscription", err.Error())
	}

	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.pl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.pl, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.pl.hasNextLink() && result.pl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client LabsClient) ListBySubscriptionPreparer(ctx context.Context, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.LabServices/labs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client LabsClient) ListBySubscriptionResponder(resp *http.Response) (result PagedLabs, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client LabsClient) listBySubscriptionNextResults(ctx context.Context, lastResults PagedLabs) (result PagedLabs, err error) {
	req, err := lastResults.pagedLabsPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "labservices.LabsClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "labservices.LabsClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client LabsClient) ListBySubscriptionComplete(ctx context.Context, filter string) (result PagedLabsIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx, filter)
	return
}

// Publish publish or re-publish a lab. This will create or update all lab resources, such as virtual machines.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// labName - the name of the lab that uniquely identifies it within containing lab account. Used in resource
// URIs.
func (client LabsClient) Publish(ctx context.Context, resourceGroupName string, labName string) (result LabsPublishFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabsClient.Publish")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: labName,
			Constraints: []validation.Constraint{{Target: "labName", Name: validation.MaxLength, Rule: 100, Chain: nil},
				{Target: "labName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("labservices.LabsClient", "Publish", err.Error())
	}

	req, err := client.PublishPreparer(ctx, resourceGroupName, labName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "Publish", nil, "Failure preparing request")
		return
	}

	result, err = client.PublishSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "Publish", result.Response(), "Failure sending request")
		return
	}

	return
}

// PublishPreparer prepares the Publish request.
func (client LabsClient) PublishPreparer(ctx context.Context, resourceGroupName string, labName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/publish", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PublishSender sends the Publish request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) PublishSender(req *http.Request) (future LabsPublishFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// PublishResponder handles the response to the Publish request. The method always
// closes the http.Response Body.
func (client LabsClient) PublishResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// SyncGroup action used to manually kick off an AAD group sync job.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// labName - the name of the lab that uniquely identifies it within containing lab account. Used in resource
// URIs.
func (client LabsClient) SyncGroup(ctx context.Context, resourceGroupName string, labName string) (result LabsSyncGroupFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabsClient.SyncGroup")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: labName,
			Constraints: []validation.Constraint{{Target: "labName", Name: validation.MaxLength, Rule: 100, Chain: nil},
				{Target: "labName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("labservices.LabsClient", "SyncGroup", err.Error())
	}

	req, err := client.SyncGroupPreparer(ctx, resourceGroupName, labName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "SyncGroup", nil, "Failure preparing request")
		return
	}

	result, err = client.SyncGroupSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "SyncGroup", result.Response(), "Failure sending request")
		return
	}

	return
}

// SyncGroupPreparer prepares the SyncGroup request.
func (client LabsClient) SyncGroupPreparer(ctx context.Context, resourceGroupName string, labName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/syncGroup", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SyncGroupSender sends the SyncGroup request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) SyncGroupSender(req *http.Request) (future LabsSyncGroupFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// SyncGroupResponder handles the response to the SyncGroup request. The method always
// closes the http.Response Body.
func (client LabsClient) SyncGroupResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update operation to update a lab resource.
// Parameters:
// body - the request body.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// labName - the name of the lab that uniquely identifies it within containing lab account. Used in resource
// URIs.
func (client LabsClient) Update(ctx context.Context, body LabUpdate, resourceGroupName string, labName string) (result LabsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabsClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: labName,
			Constraints: []validation.Constraint{{Target: "labName", Name: validation.MaxLength, Rule: 100, Chain: nil},
				{Target: "labName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("labservices.LabsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, body, resourceGroupName, labName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client LabsClient) UpdatePreparer(ctx context.Context, body LabUpdate, resourceGroupName string, labName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) UpdateSender(req *http.Request) (future LabsUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client LabsClient) UpdateResponder(resp *http.Response) (result Lab, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}