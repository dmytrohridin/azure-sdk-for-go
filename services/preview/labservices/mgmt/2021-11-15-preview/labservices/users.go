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

// UsersClient is the client for the Users methods of the Labservices service.
type UsersClient struct {
	BaseClient
}

// NewUsersClient creates an instance of the UsersClient client.
func NewUsersClient(subscriptionID string) UsersClient {
	return NewUsersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUsersClientWithBaseURI creates an instance of the UsersClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewUsersClientWithBaseURI(baseURI string, subscriptionID string) UsersClient {
	return UsersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate operation to create or update a lab user.
// Parameters:
// body - the request body.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// labName - the name of the lab that uniquely identifies it within containing lab account. Used in resource
// URIs.
// userName - the name of the user that uniquely identifies it within containing lab. Used in resource URIs.
func (client UsersClient) CreateOrUpdate(ctx context.Context, body User, resourceGroupName string, labName string, userName string) (result UsersCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.CreateOrUpdate")
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
			Constraints: []validation.Constraint{{Target: "body.UserProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "body.UserProperties.Email", Name: validation.Null, Rule: true, Chain: nil}}}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: labName,
			Constraints: []validation.Constraint{{Target: "labName", Name: validation.MaxLength, Rule: 100, Chain: nil},
				{Target: "labName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: userName,
			Constraints: []validation.Constraint{{Target: "userName", Name: validation.MaxLength, Rule: 100, Chain: nil},
				{Target: "userName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "userName", Name: validation.Pattern, Rule: `^[-\w\\._\\(\\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("labservices.UsersClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, body, resourceGroupName, labName, userName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.UsersClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.UsersClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client UsersClient) CreateOrUpdatePreparer(ctx context.Context, body User, resourceGroupName string, labName string, userName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userName":          autorest.Encode("path", userName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/users/{userName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) CreateOrUpdateSender(req *http.Request) (future UsersCreateOrUpdateFuture, err error) {
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
func (client UsersClient) CreateOrUpdateResponder(resp *http.Response) (result User, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete operation to delete a user resource.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// labName - the name of the lab that uniquely identifies it within containing lab account. Used in resource
// URIs.
// userName - the name of the user that uniquely identifies it within containing lab. Used in resource URIs.
func (client UsersClient) Delete(ctx context.Context, resourceGroupName string, labName string, userName string) (result UsersDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.Delete")
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
				{Target: "labName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: userName,
			Constraints: []validation.Constraint{{Target: "userName", Name: validation.MaxLength, Rule: 100, Chain: nil},
				{Target: "userName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "userName", Name: validation.Pattern, Rule: `^[-\w\\._\\(\\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("labservices.UsersClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, labName, userName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.UsersClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.UsersClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client UsersClient) DeletePreparer(ctx context.Context, resourceGroupName string, labName string, userName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userName":          autorest.Encode("path", userName),
	}

	const APIVersion = "2021-11-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/users/{userName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) DeleteSender(req *http.Request) (future UsersDeleteFuture, err error) {
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
func (client UsersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns the properties of a lab user.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// labName - the name of the lab that uniquely identifies it within containing lab account. Used in resource
// URIs.
// userName - the name of the user that uniquely identifies it within containing lab. Used in resource URIs.
func (client UsersClient) Get(ctx context.Context, resourceGroupName string, labName string, userName string) (result User, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.Get")
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
				{Target: "labName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: userName,
			Constraints: []validation.Constraint{{Target: "userName", Name: validation.MaxLength, Rule: 100, Chain: nil},
				{Target: "userName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "userName", Name: validation.Pattern, Rule: `^[-\w\\._\\(\\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("labservices.UsersClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, labName, userName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.UsersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.UsersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.UsersClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client UsersClient) GetPreparer(ctx context.Context, resourceGroupName string, labName string, userName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userName":          autorest.Encode("path", userName),
	}

	const APIVersion = "2021-11-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/users/{userName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client UsersClient) GetResponder(resp *http.Response) (result User, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Invite operation to invite a user to a lab.
// Parameters:
// body - the request body.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// labName - the name of the lab that uniquely identifies it within containing lab account. Used in resource
// URIs.
// userName - the name of the user that uniquely identifies it within containing lab. Used in resource URIs.
func (client UsersClient) Invite(ctx context.Context, body InviteBody, resourceGroupName string, labName string, userName string) (result UsersInviteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.Invite")
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
				{Target: "labName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: userName,
			Constraints: []validation.Constraint{{Target: "userName", Name: validation.MaxLength, Rule: 100, Chain: nil},
				{Target: "userName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "userName", Name: validation.Pattern, Rule: `^[-\w\\._\\(\\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("labservices.UsersClient", "Invite", err.Error())
	}

	req, err := client.InvitePreparer(ctx, body, resourceGroupName, labName, userName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.UsersClient", "Invite", nil, "Failure preparing request")
		return
	}

	result, err = client.InviteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.UsersClient", "Invite", result.Response(), "Failure sending request")
		return
	}

	return
}

// InvitePreparer prepares the Invite request.
func (client UsersClient) InvitePreparer(ctx context.Context, body InviteBody, resourceGroupName string, labName string, userName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userName":          autorest.Encode("path", userName),
	}

	const APIVersion = "2021-11-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/users/{userName}/invite", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// InviteSender sends the Invite request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) InviteSender(req *http.Request) (future UsersInviteFuture, err error) {
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

// InviteResponder handles the response to the Invite request. The method always
// closes the http.Response Body.
func (client UsersClient) InviteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListByLab returns a list of all users for a lab.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// labName - the name of the lab that uniquely identifies it within containing lab account. Used in resource
// URIs.
// filter - the filter to apply to the operation.
func (client UsersClient) ListByLab(ctx context.Context, resourceGroupName string, labName string, filter string) (result PagedUsersPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.ListByLab")
		defer func() {
			sc := -1
			if result.pu.Response.Response != nil {
				sc = result.pu.Response.Response.StatusCode
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
		return result, validation.NewError("labservices.UsersClient", "ListByLab", err.Error())
	}

	result.fn = client.listByLabNextResults
	req, err := client.ListByLabPreparer(ctx, resourceGroupName, labName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.UsersClient", "ListByLab", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByLabSender(req)
	if err != nil {
		result.pu.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.UsersClient", "ListByLab", resp, "Failure sending request")
		return
	}

	result.pu, err = client.ListByLabResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.UsersClient", "ListByLab", resp, "Failure responding to request")
		return
	}
	if result.pu.hasNextLink() && result.pu.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByLabPreparer prepares the ListByLab request.
func (client UsersClient) ListByLabPreparer(ctx context.Context, resourceGroupName string, labName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/users", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByLabSender sends the ListByLab request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) ListByLabSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByLabResponder handles the response to the ListByLab request. The method always
// closes the http.Response Body.
func (client UsersClient) ListByLabResponder(resp *http.Response) (result PagedUsers, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByLabNextResults retrieves the next set of results, if any.
func (client UsersClient) listByLabNextResults(ctx context.Context, lastResults PagedUsers) (result PagedUsers, err error) {
	req, err := lastResults.pagedUsersPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "labservices.UsersClient", "listByLabNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByLabSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "labservices.UsersClient", "listByLabNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByLabResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.UsersClient", "listByLabNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByLabComplete enumerates all values, automatically crossing page boundaries as required.
func (client UsersClient) ListByLabComplete(ctx context.Context, resourceGroupName string, labName string, filter string) (result PagedUsersIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.ListByLab")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByLab(ctx, resourceGroupName, labName, filter)
	return
}

// Update operation to update a lab user.
// Parameters:
// body - the request body.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// labName - the name of the lab that uniquely identifies it within containing lab account. Used in resource
// URIs.
// userName - the name of the user that uniquely identifies it within containing lab. Used in resource URIs.
func (client UsersClient) Update(ctx context.Context, body UserUpdate, resourceGroupName string, labName string, userName string) (result UsersUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.Update")
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
				{Target: "labName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: userName,
			Constraints: []validation.Constraint{{Target: "userName", Name: validation.MaxLength, Rule: 100, Chain: nil},
				{Target: "userName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "userName", Name: validation.Pattern, Rule: `^[-\w\\._\\(\\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("labservices.UsersClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, body, resourceGroupName, labName, userName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.UsersClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.UsersClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client UsersClient) UpdatePreparer(ctx context.Context, body UserUpdate, resourceGroupName string, labName string, userName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userName":          autorest.Encode("path", userName),
	}

	const APIVersion = "2021-11-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/users/{userName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) UpdateSender(req *http.Request) (future UsersUpdateFuture, err error) {
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
func (client UsersClient) UpdateResponder(resp *http.Response) (result User, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
