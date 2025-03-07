package storagesync

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

// ServerEndpointsClient is the microsoft Storage Sync Service API
type ServerEndpointsClient struct {
	BaseClient
}

// NewServerEndpointsClient creates an instance of the ServerEndpointsClient client.
func NewServerEndpointsClient(subscriptionID string) ServerEndpointsClient {
	return NewServerEndpointsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServerEndpointsClientWithBaseURI creates an instance of the ServerEndpointsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewServerEndpointsClientWithBaseURI(baseURI string, subscriptionID string) ServerEndpointsClient {
	return ServerEndpointsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create a new ServerEndpoint.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// storageSyncServiceName - name of Storage Sync Service resource.
// syncGroupName - name of Sync Group resource.
// serverEndpointName - name of Server Endpoint object.
// parameters - body of Server Endpoint object.
func (client ServerEndpointsClient) Create(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string, parameters ServerEndpointCreateParameters) (result ServerEndpointsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServerEndpointsClient.Create")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ServerEndpointCreateParametersProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.ServerEndpointCreateParametersProperties.VolumeFreeSpacePercent", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.ServerEndpointCreateParametersProperties.VolumeFreeSpacePercent", Name: validation.InclusiveMaximum, Rule: int64(100), Chain: nil},
						{Target: "parameters.ServerEndpointCreateParametersProperties.VolumeFreeSpacePercent", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil},
					}},
					{Target: "parameters.ServerEndpointCreateParametersProperties.TierFilesOlderThanDays", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.ServerEndpointCreateParametersProperties.TierFilesOlderThanDays", Name: validation.InclusiveMaximum, Rule: int64(2147483647), Chain: nil},
							{Target: "parameters.ServerEndpointCreateParametersProperties.TierFilesOlderThanDays", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("storagesync.ServerEndpointsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, serverEndpointName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.ServerEndpointsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.ServerEndpointsClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ServerEndpointsClient) CreatePreparer(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string, parameters ServerEndpointCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"serverEndpointName":     autorest.Encode("path", serverEndpointName),
		"storageSyncServiceName": autorest.Encode("path", storageSyncServiceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":          autorest.Encode("path", syncGroupName),
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints/{serverEndpointName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ServerEndpointsClient) CreateSender(req *http.Request) (future ServerEndpointsCreateFuture, err error) {
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

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ServerEndpointsClient) CreateResponder(resp *http.Response) (result ServerEndpoint, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a given ServerEndpoint.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// storageSyncServiceName - name of Storage Sync Service resource.
// syncGroupName - name of Sync Group resource.
// serverEndpointName - name of Server Endpoint object.
func (client ServerEndpointsClient) Delete(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string) (result ServerEndpointsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServerEndpointsClient.Delete")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagesync.ServerEndpointsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, serverEndpointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.ServerEndpointsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.ServerEndpointsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ServerEndpointsClient) DeletePreparer(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"serverEndpointName":     autorest.Encode("path", serverEndpointName),
		"storageSyncServiceName": autorest.Encode("path", storageSyncServiceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":          autorest.Encode("path", syncGroupName),
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints/{serverEndpointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ServerEndpointsClient) DeleteSender(req *http.Request) (future ServerEndpointsDeleteFuture, err error) {
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
func (client ServerEndpointsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a ServerEndpoint.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// storageSyncServiceName - name of Storage Sync Service resource.
// syncGroupName - name of Sync Group resource.
// serverEndpointName - name of Server Endpoint object.
func (client ServerEndpointsClient) Get(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string) (result ServerEndpoint, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServerEndpointsClient.Get")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagesync.ServerEndpointsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, serverEndpointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.ServerEndpointsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storagesync.ServerEndpointsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.ServerEndpointsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ServerEndpointsClient) GetPreparer(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"serverEndpointName":     autorest.Encode("path", serverEndpointName),
		"storageSyncServiceName": autorest.Encode("path", storageSyncServiceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":          autorest.Encode("path", syncGroupName),
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints/{serverEndpointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ServerEndpointsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ServerEndpointsClient) GetResponder(resp *http.Response) (result ServerEndpoint, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySyncGroup get a ServerEndpoint list.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// storageSyncServiceName - name of Storage Sync Service resource.
// syncGroupName - name of Sync Group resource.
func (client ServerEndpointsClient) ListBySyncGroup(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string) (result ServerEndpointArray, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServerEndpointsClient.ListBySyncGroup")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagesync.ServerEndpointsClient", "ListBySyncGroup", err.Error())
	}

	req, err := client.ListBySyncGroupPreparer(ctx, resourceGroupName, storageSyncServiceName, syncGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.ServerEndpointsClient", "ListBySyncGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySyncGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storagesync.ServerEndpointsClient", "ListBySyncGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListBySyncGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.ServerEndpointsClient", "ListBySyncGroup", resp, "Failure responding to request")
		return
	}

	return
}

// ListBySyncGroupPreparer prepares the ListBySyncGroup request.
func (client ServerEndpointsClient) ListBySyncGroupPreparer(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"storageSyncServiceName": autorest.Encode("path", storageSyncServiceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":          autorest.Encode("path", syncGroupName),
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySyncGroupSender sends the ListBySyncGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ServerEndpointsClient) ListBySyncGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySyncGroupResponder handles the response to the ListBySyncGroup request. The method always
// closes the http.Response Body.
func (client ServerEndpointsClient) ListBySyncGroupResponder(resp *http.Response) (result ServerEndpointArray, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RecallAction recall a server endpoint.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// storageSyncServiceName - name of Storage Sync Service resource.
// syncGroupName - name of Sync Group resource.
// serverEndpointName - name of Server Endpoint object.
// parameters - body of Recall Action object.
func (client ServerEndpointsClient) RecallAction(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string, parameters RecallActionParameters) (result ServerEndpointsRecallActionFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServerEndpointsClient.RecallAction")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagesync.ServerEndpointsClient", "RecallAction", err.Error())
	}

	req, err := client.RecallActionPreparer(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, serverEndpointName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.ServerEndpointsClient", "RecallAction", nil, "Failure preparing request")
		return
	}

	result, err = client.RecallActionSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.ServerEndpointsClient", "RecallAction", result.Response(), "Failure sending request")
		return
	}

	return
}

// RecallActionPreparer prepares the RecallAction request.
func (client ServerEndpointsClient) RecallActionPreparer(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string, parameters RecallActionParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"serverEndpointName":     autorest.Encode("path", serverEndpointName),
		"storageSyncServiceName": autorest.Encode("path", storageSyncServiceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":          autorest.Encode("path", syncGroupName),
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints/{serverEndpointName}/recallAction", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RecallActionSender sends the RecallAction request. The method will close the
// http.Response Body if it receives an error.
func (client ServerEndpointsClient) RecallActionSender(req *http.Request) (future ServerEndpointsRecallActionFuture, err error) {
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

// RecallActionResponder handles the response to the RecallAction request. The method always
// closes the http.Response Body.
func (client ServerEndpointsClient) RecallActionResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update patch a given ServerEndpoint.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// storageSyncServiceName - name of Storage Sync Service resource.
// syncGroupName - name of Sync Group resource.
// serverEndpointName - name of Server Endpoint object.
// parameters - any of the properties applicable in PUT request.
func (client ServerEndpointsClient) Update(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string, parameters *ServerEndpointUpdateParameters) (result ServerEndpointsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServerEndpointsClient.Update")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagesync.ServerEndpointsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, serverEndpointName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.ServerEndpointsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.ServerEndpointsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ServerEndpointsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string, parameters *ServerEndpointUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"serverEndpointName":     autorest.Encode("path", serverEndpointName),
		"storageSyncServiceName": autorest.Encode("path", storageSyncServiceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":          autorest.Encode("path", syncGroupName),
	}

	const APIVersion = "2019-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints/{serverEndpointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if parameters != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(parameters))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ServerEndpointsClient) UpdateSender(req *http.Request) (future ServerEndpointsUpdateFuture, err error) {
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
func (client ServerEndpointsClient) UpdateResponder(resp *http.Response) (result ServerEndpoint, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
