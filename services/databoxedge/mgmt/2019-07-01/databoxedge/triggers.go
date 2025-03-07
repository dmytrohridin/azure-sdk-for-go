package databoxedge

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// TriggersClient is the client for the Triggers methods of the Databoxedge service.
type TriggersClient struct {
	BaseClient
}

// NewTriggersClient creates an instance of the TriggersClient client.
func NewTriggersClient(subscriptionID string) TriggersClient {
	return NewTriggersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTriggersClientWithBaseURI creates an instance of the TriggersClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewTriggersClientWithBaseURI(baseURI string, subscriptionID string) TriggersClient {
	return TriggersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a trigger.
// Parameters:
// deviceName - creates or updates a trigger
// name - the trigger name.
// trigger - the trigger.
// resourceGroupName - the resource group name.
func (client TriggersClient) CreateOrUpdate(ctx context.Context, deviceName string, name string, trigger BasicTrigger, resourceGroupName string) (result TriggersCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TriggersClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, deviceName, name, trigger, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.TriggersClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.TriggersClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client TriggersClient) CreateOrUpdatePreparer(ctx context.Context, deviceName string, name string, trigger BasicTrigger, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        autorest.Encode("path", deviceName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggers/{name}", pathParameters),
		autorest.WithJSON(trigger),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client TriggersClient) CreateOrUpdateSender(req *http.Request) (future TriggersCreateOrUpdateFuture, err error) {
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
func (client TriggersClient) CreateOrUpdateResponder(resp *http.Response) (result TriggerModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the trigger on the gateway device.
// Parameters:
// deviceName - the device name.
// name - the trigger name.
// resourceGroupName - the resource group name.
func (client TriggersClient) Delete(ctx context.Context, deviceName string, name string, resourceGroupName string) (result TriggersDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TriggersClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, deviceName, name, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.TriggersClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.TriggersClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client TriggersClient) DeletePreparer(ctx context.Context, deviceName string, name string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        autorest.Encode("path", deviceName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggers/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client TriggersClient) DeleteSender(req *http.Request) (future TriggersDeleteFuture, err error) {
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
func (client TriggersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a specific trigger by name.
// Parameters:
// deviceName - the device name.
// name - the trigger name.
// resourceGroupName - the resource group name.
func (client TriggersClient) Get(ctx context.Context, deviceName string, name string, resourceGroupName string) (result TriggerModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TriggersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, deviceName, name, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.TriggersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databoxedge.TriggersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.TriggersClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client TriggersClient) GetPreparer(ctx context.Context, deviceName string, name string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        autorest.Encode("path", deviceName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggers/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client TriggersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client TriggersClient) GetResponder(resp *http.Response) (result TriggerModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDataBoxEdgeDevice lists all the triggers configured in the device.
// Parameters:
// deviceName - the device name.
// resourceGroupName - the resource group name.
// expand - specify $filter='CustomContextTag eq <tag>' to filter on custom context tag property
func (client TriggersClient) ListByDataBoxEdgeDevice(ctx context.Context, deviceName string, resourceGroupName string, expand string) (result TriggerListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TriggersClient.ListByDataBoxEdgeDevice")
		defer func() {
			sc := -1
			if result.tl.Response.Response != nil {
				sc = result.tl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByDataBoxEdgeDeviceNextResults
	req, err := client.ListByDataBoxEdgeDevicePreparer(ctx, deviceName, resourceGroupName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.TriggersClient", "ListByDataBoxEdgeDevice", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDataBoxEdgeDeviceSender(req)
	if err != nil {
		result.tl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databoxedge.TriggersClient", "ListByDataBoxEdgeDevice", resp, "Failure sending request")
		return
	}

	result.tl, err = client.ListByDataBoxEdgeDeviceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.TriggersClient", "ListByDataBoxEdgeDevice", resp, "Failure responding to request")
		return
	}
	if result.tl.hasNextLink() && result.tl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByDataBoxEdgeDevicePreparer prepares the ListByDataBoxEdgeDevice request.
func (client TriggersClient) ListByDataBoxEdgeDevicePreparer(ctx context.Context, deviceName string, resourceGroupName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        autorest.Encode("path", deviceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDataBoxEdgeDeviceSender sends the ListByDataBoxEdgeDevice request. The method will close the
// http.Response Body if it receives an error.
func (client TriggersClient) ListByDataBoxEdgeDeviceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByDataBoxEdgeDeviceResponder handles the response to the ListByDataBoxEdgeDevice request. The method always
// closes the http.Response Body.
func (client TriggersClient) ListByDataBoxEdgeDeviceResponder(resp *http.Response) (result TriggerList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByDataBoxEdgeDeviceNextResults retrieves the next set of results, if any.
func (client TriggersClient) listByDataBoxEdgeDeviceNextResults(ctx context.Context, lastResults TriggerList) (result TriggerList, err error) {
	req, err := lastResults.triggerListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "databoxedge.TriggersClient", "listByDataBoxEdgeDeviceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByDataBoxEdgeDeviceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "databoxedge.TriggersClient", "listByDataBoxEdgeDeviceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByDataBoxEdgeDeviceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.TriggersClient", "listByDataBoxEdgeDeviceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByDataBoxEdgeDeviceComplete enumerates all values, automatically crossing page boundaries as required.
func (client TriggersClient) ListByDataBoxEdgeDeviceComplete(ctx context.Context, deviceName string, resourceGroupName string, expand string) (result TriggerListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TriggersClient.ListByDataBoxEdgeDevice")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByDataBoxEdgeDevice(ctx, deviceName, resourceGroupName, expand)
	return
}
