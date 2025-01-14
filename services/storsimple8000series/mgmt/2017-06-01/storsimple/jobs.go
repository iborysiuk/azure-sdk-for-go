package storsimple

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

// JobsClient is the client for the Jobs methods of the Storsimple service.
type JobsClient struct {
	BaseClient
}

// NewJobsClient creates an instance of the JobsClient client.
func NewJobsClient(subscriptionID string) JobsClient {
	return NewJobsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewJobsClientWithBaseURI creates an instance of the JobsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewJobsClientWithBaseURI(baseURI string, subscriptionID string) JobsClient {
	return JobsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Cancel cancels a job on the device.
// Parameters:
// deviceName - the device name
// jobName - the jobName.
// resourceGroupName - the resource group name
// managerName - the manager name
func (client JobsClient) Cancel(ctx context.Context, deviceName string, jobName string, resourceGroupName string, managerName string) (result JobsCancelFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.Cancel")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.JobsClient", "Cancel", err.Error())
	}

	req, err := client.CancelPreparer(ctx, deviceName, jobName, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.JobsClient", "Cancel", nil, "Failure preparing request")
		return
	}

	result, err = client.CancelSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.JobsClient", "Cancel", result.Response(), "Failure sending request")
		return
	}

	return
}

// CancelPreparer prepares the Cancel request.
func (client JobsClient) CancelPreparer(ctx context.Context, deviceName string, jobName string, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        deviceName,
		"jobName":           jobName,
		"managerName":       managerName,
		"resourceGroupName": resourceGroupName,
		"subscriptionId":    client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/jobs/{jobName}/cancel", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CancelSender sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) CancelSender(req *http.Request) (future JobsCancelFuture, err error) {
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

// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client JobsClient) CancelResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the details of the specified job name.
// Parameters:
// deviceName - the device name
// jobName - the job Name.
// resourceGroupName - the resource group name
// managerName - the manager name
func (client JobsClient) Get(ctx context.Context, deviceName string, jobName string, resourceGroupName string, managerName string) (result Job, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.JobsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, deviceName, jobName, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.JobsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple.JobsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.JobsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client JobsClient) GetPreparer(ctx context.Context, deviceName string, jobName string, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        deviceName,
		"jobName":           jobName,
		"managerName":       managerName,
		"resourceGroupName": resourceGroupName,
		"subscriptionId":    client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/jobs/{jobName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client JobsClient) GetResponder(resp *http.Response) (result Job, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDevice gets all the jobs for specified device. With optional OData query parameters, a filtered set of jobs is
// returned.
// Parameters:
// deviceName - the device name
// resourceGroupName - the resource group name
// managerName - the manager name
// filter - oData Filter options
func (client JobsClient) ListByDevice(ctx context.Context, deviceName string, resourceGroupName string, managerName string, filter string) (result JobListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.ListByDevice")
		defer func() {
			sc := -1
			if result.jl.Response.Response != nil {
				sc = result.jl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.JobsClient", "ListByDevice", err.Error())
	}

	result.fn = client.listByDeviceNextResults
	req, err := client.ListByDevicePreparer(ctx, deviceName, resourceGroupName, managerName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.JobsClient", "ListByDevice", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDeviceSender(req)
	if err != nil {
		result.jl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple.JobsClient", "ListByDevice", resp, "Failure sending request")
		return
	}

	result.jl, err = client.ListByDeviceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.JobsClient", "ListByDevice", resp, "Failure responding to request")
		return
	}
	if result.jl.hasNextLink() && result.jl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByDevicePreparer prepares the ListByDevice request.
func (client JobsClient) ListByDevicePreparer(ctx context.Context, deviceName string, resourceGroupName string, managerName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        deviceName,
		"managerName":       managerName,
		"resourceGroupName": resourceGroupName,
		"subscriptionId":    client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/jobs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDeviceSender sends the ListByDevice request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) ListByDeviceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByDeviceResponder handles the response to the ListByDevice request. The method always
// closes the http.Response Body.
func (client JobsClient) ListByDeviceResponder(resp *http.Response) (result JobList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByDeviceNextResults retrieves the next set of results, if any.
func (client JobsClient) listByDeviceNextResults(ctx context.Context, lastResults JobList) (result JobList, err error) {
	req, err := lastResults.jobListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "storsimple.JobsClient", "listByDeviceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByDeviceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "storsimple.JobsClient", "listByDeviceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByDeviceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.JobsClient", "listByDeviceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByDeviceComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobsClient) ListByDeviceComplete(ctx context.Context, deviceName string, resourceGroupName string, managerName string, filter string) (result JobListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.ListByDevice")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByDevice(ctx, deviceName, resourceGroupName, managerName, filter)
	return
}

// ListByManager gets all the jobs for the specified manager. With optional OData query parameters, a filtered set of
// jobs is returned.
// Parameters:
// resourceGroupName - the resource group name
// managerName - the manager name
// filter - oData Filter options
func (client JobsClient) ListByManager(ctx context.Context, resourceGroupName string, managerName string, filter string) (result JobListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.ListByManager")
		defer func() {
			sc := -1
			if result.jl.Response.Response != nil {
				sc = result.jl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.JobsClient", "ListByManager", err.Error())
	}

	result.fn = client.listByManagerNextResults
	req, err := client.ListByManagerPreparer(ctx, resourceGroupName, managerName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.JobsClient", "ListByManager", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByManagerSender(req)
	if err != nil {
		result.jl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple.JobsClient", "ListByManager", resp, "Failure sending request")
		return
	}

	result.jl, err = client.ListByManagerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.JobsClient", "ListByManager", resp, "Failure responding to request")
		return
	}
	if result.jl.hasNextLink() && result.jl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByManagerPreparer prepares the ListByManager request.
func (client JobsClient) ListByManagerPreparer(ctx context.Context, resourceGroupName string, managerName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managerName":       managerName,
		"resourceGroupName": resourceGroupName,
		"subscriptionId":    client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/jobs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByManagerSender sends the ListByManager request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) ListByManagerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByManagerResponder handles the response to the ListByManager request. The method always
// closes the http.Response Body.
func (client JobsClient) ListByManagerResponder(resp *http.Response) (result JobList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByManagerNextResults retrieves the next set of results, if any.
func (client JobsClient) listByManagerNextResults(ctx context.Context, lastResults JobList) (result JobList, err error) {
	req, err := lastResults.jobListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "storsimple.JobsClient", "listByManagerNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByManagerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "storsimple.JobsClient", "listByManagerNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByManagerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.JobsClient", "listByManagerNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByManagerComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobsClient) ListByManagerComplete(ctx context.Context, resourceGroupName string, managerName string, filter string) (result JobListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.ListByManager")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByManager(ctx, resourceGroupName, managerName, filter)
	return
}
