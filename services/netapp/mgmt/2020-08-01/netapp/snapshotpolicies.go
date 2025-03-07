package netapp

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

// SnapshotPoliciesClient is the microsoft NetApp Azure Resource Provider specification
type SnapshotPoliciesClient struct {
	BaseClient
}

// NewSnapshotPoliciesClient creates an instance of the SnapshotPoliciesClient client.
func NewSnapshotPoliciesClient(subscriptionID string) SnapshotPoliciesClient {
	return NewSnapshotPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSnapshotPoliciesClientWithBaseURI creates an instance of the SnapshotPoliciesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewSnapshotPoliciesClientWithBaseURI(baseURI string, subscriptionID string) SnapshotPoliciesClient {
	return SnapshotPoliciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create a snapshot policy
// Parameters:
// body - snapshot policy object supplied in the body of the operation.
// resourceGroupName - the name of the resource group.
// accountName - the name of the NetApp account
// snapshotPolicyName - the name of the snapshot policy target
func (client SnapshotPoliciesClient) Create(ctx context.Context, body SnapshotPolicy, resourceGroupName string, accountName string, snapshotPolicyName string) (result SnapshotPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SnapshotPoliciesClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.Location", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "body.SnapshotPolicyProperties", Name: validation.Null, Rule: true, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("netapp.SnapshotPoliciesClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, body, resourceGroupName, accountName, snapshotPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.SnapshotPoliciesClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "netapp.SnapshotPoliciesClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.SnapshotPoliciesClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client SnapshotPoliciesClient) CreatePreparer(ctx context.Context, body SnapshotPolicy, resourceGroupName string, accountName string, snapshotPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":        autorest.Encode("path", accountName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"snapshotPolicyName": autorest.Encode("path", snapshotPolicyName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	body.ID = nil
	body.Name = nil
	body.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/snapshotPolicies/{snapshotPolicyName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client SnapshotPoliciesClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client SnapshotPoliciesClient) CreateResponder(resp *http.Response) (result SnapshotPolicy, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete snapshot policy
// Parameters:
// resourceGroupName - the name of the resource group.
// accountName - the name of the NetApp account
// snapshotPolicyName - the name of the snapshot policy target
func (client SnapshotPoliciesClient) Delete(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string) (result SnapshotPoliciesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SnapshotPoliciesClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("netapp.SnapshotPoliciesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName, snapshotPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.SnapshotPoliciesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.SnapshotPoliciesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SnapshotPoliciesClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":        autorest.Encode("path", accountName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"snapshotPolicyName": autorest.Encode("path", snapshotPolicyName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/snapshotPolicies/{snapshotPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SnapshotPoliciesClient) DeleteSender(req *http.Request) (future SnapshotPoliciesDeleteFuture, err error) {
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
func (client SnapshotPoliciesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a snapshot Policy
// Parameters:
// resourceGroupName - the name of the resource group.
// accountName - the name of the NetApp account
// snapshotPolicyName - the name of the snapshot policy target
func (client SnapshotPoliciesClient) Get(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string) (result SnapshotPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SnapshotPoliciesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("netapp.SnapshotPoliciesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, snapshotPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.SnapshotPoliciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "netapp.SnapshotPoliciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.SnapshotPoliciesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client SnapshotPoliciesClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":        autorest.Encode("path", accountName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"snapshotPolicyName": autorest.Encode("path", snapshotPolicyName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/snapshotPolicies/{snapshotPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SnapshotPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SnapshotPoliciesClient) GetResponder(resp *http.Response) (result SnapshotPolicy, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list snapshot policy
// Parameters:
// resourceGroupName - the name of the resource group.
// accountName - the name of the NetApp account
func (client SnapshotPoliciesClient) List(ctx context.Context, resourceGroupName string, accountName string) (result SnapshotPoliciesList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SnapshotPoliciesClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("netapp.SnapshotPoliciesClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.SnapshotPoliciesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "netapp.SnapshotPoliciesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.SnapshotPoliciesClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client SnapshotPoliciesClient) ListPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/snapshotPolicies", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SnapshotPoliciesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SnapshotPoliciesClient) ListResponder(resp *http.Response) (result SnapshotPoliciesList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListVolumes get volumes associated with snapshot policy
// Parameters:
// resourceGroupName - the name of the resource group.
// accountName - the name of the NetApp account
// snapshotPolicyName - the name of the snapshot policy target
func (client SnapshotPoliciesClient) ListVolumes(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string) (result SnapshotPolicyVolumeList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SnapshotPoliciesClient.ListVolumes")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("netapp.SnapshotPoliciesClient", "ListVolumes", err.Error())
	}

	req, err := client.ListVolumesPreparer(ctx, resourceGroupName, accountName, snapshotPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.SnapshotPoliciesClient", "ListVolumes", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListVolumesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "netapp.SnapshotPoliciesClient", "ListVolumes", resp, "Failure sending request")
		return
	}

	result, err = client.ListVolumesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.SnapshotPoliciesClient", "ListVolumes", resp, "Failure responding to request")
		return
	}

	return
}

// ListVolumesPreparer prepares the ListVolumes request.
func (client SnapshotPoliciesClient) ListVolumesPreparer(ctx context.Context, resourceGroupName string, accountName string, snapshotPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":        autorest.Encode("path", accountName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"snapshotPolicyName": autorest.Encode("path", snapshotPolicyName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/snapshotPolicies/{snapshotPolicyName}/listVolumes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListVolumesSender sends the ListVolumes request. The method will close the
// http.Response Body if it receives an error.
func (client SnapshotPoliciesClient) ListVolumesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListVolumesResponder handles the response to the ListVolumes request. The method always
// closes the http.Response Body.
func (client SnapshotPoliciesClient) ListVolumesResponder(resp *http.Response) (result SnapshotPolicyVolumeList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update patch a snapshot policy
// Parameters:
// body - snapshot policy object supplied in the body of the operation.
// resourceGroupName - the name of the resource group.
// accountName - the name of the NetApp account
// snapshotPolicyName - the name of the snapshot policy target
func (client SnapshotPoliciesClient) Update(ctx context.Context, body SnapshotPolicyPatch, resourceGroupName string, accountName string, snapshotPolicyName string) (result SnapshotPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SnapshotPoliciesClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("netapp.SnapshotPoliciesClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, body, resourceGroupName, accountName, snapshotPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.SnapshotPoliciesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "netapp.SnapshotPoliciesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.SnapshotPoliciesClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client SnapshotPoliciesClient) UpdatePreparer(ctx context.Context, body SnapshotPolicyPatch, resourceGroupName string, accountName string, snapshotPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":        autorest.Encode("path", accountName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"snapshotPolicyName": autorest.Encode("path", snapshotPolicyName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	body.ID = nil
	body.Name = nil
	body.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/snapshotPolicies/{snapshotPolicyName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client SnapshotPoliciesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client SnapshotPoliciesClient) UpdateResponder(resp *http.Response) (result SnapshotPolicy, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
