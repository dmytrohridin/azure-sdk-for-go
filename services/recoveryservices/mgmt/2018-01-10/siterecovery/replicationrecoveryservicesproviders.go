package siterecovery

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

// ReplicationRecoveryServicesProvidersClient is the client for the ReplicationRecoveryServicesProviders methods of the
// Siterecovery service.
type ReplicationRecoveryServicesProvidersClient struct {
	BaseClient
}

// NewReplicationRecoveryServicesProvidersClient creates an instance of the ReplicationRecoveryServicesProvidersClient
// client.
func NewReplicationRecoveryServicesProvidersClient(subscriptionID string, resourceGroupName string, resourceName string) ReplicationRecoveryServicesProvidersClient {
	return NewReplicationRecoveryServicesProvidersClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, resourceName)
}

// NewReplicationRecoveryServicesProvidersClientWithBaseURI creates an instance of the
// ReplicationRecoveryServicesProvidersClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewReplicationRecoveryServicesProvidersClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, resourceName string) ReplicationRecoveryServicesProvidersClient {
	return ReplicationRecoveryServicesProvidersClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, resourceName)}
}

// Create the operation to add a recovery services provider.
// Parameters:
// fabricName - fabric name.
// providerName - recovery services provider name.
// addProviderInput - add provider input.
func (client ReplicationRecoveryServicesProvidersClient) Create(ctx context.Context, fabricName string, providerName string, addProviderInput AddRecoveryServicesProviderInput) (result ReplicationRecoveryServicesProvidersCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationRecoveryServicesProvidersClient.Create")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: addProviderInput,
			Constraints: []validation.Constraint{{Target: "addProviderInput.Properties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "addProviderInput.Properties.MachineName", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "addProviderInput.Properties.AuthenticationIdentityInput", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "addProviderInput.Properties.AuthenticationIdentityInput.TenantID", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "addProviderInput.Properties.AuthenticationIdentityInput.ApplicationID", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "addProviderInput.Properties.AuthenticationIdentityInput.ObjectID", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "addProviderInput.Properties.AuthenticationIdentityInput.Audience", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "addProviderInput.Properties.AuthenticationIdentityInput.AadAuthority", Name: validation.Null, Rule: true, Chain: nil},
						}},
					{Target: "addProviderInput.Properties.ResourceAccessIdentityInput", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "addProviderInput.Properties.ResourceAccessIdentityInput.TenantID", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "addProviderInput.Properties.ResourceAccessIdentityInput.ApplicationID", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "addProviderInput.Properties.ResourceAccessIdentityInput.ObjectID", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "addProviderInput.Properties.ResourceAccessIdentityInput.Audience", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "addProviderInput.Properties.ResourceAccessIdentityInput.AadAuthority", Name: validation.Null, Rule: true, Chain: nil},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("siterecovery.ReplicationRecoveryServicesProvidersClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, fabricName, providerName, addProviderInput)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationRecoveryServicesProvidersClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationRecoveryServicesProvidersClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ReplicationRecoveryServicesProvidersClient) CreatePreparer(ctx context.Context, fabricName string, providerName string, addProviderInput AddRecoveryServicesProviderInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"providerName":      autorest.Encode("path", providerName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders/{providerName}", pathParameters),
		autorest.WithJSON(addProviderInput),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationRecoveryServicesProvidersClient) CreateSender(req *http.Request) (future ReplicationRecoveryServicesProvidersCreateFuture, err error) {
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
func (client ReplicationRecoveryServicesProvidersClient) CreateResponder(resp *http.Response) (result RecoveryServicesProvider, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the operation to removes/delete(unregister) a recovery services provider from the vault
// Parameters:
// fabricName - fabric name.
// providerName - recovery services provider name.
func (client ReplicationRecoveryServicesProvidersClient) Delete(ctx context.Context, fabricName string, providerName string) (result ReplicationRecoveryServicesProvidersDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationRecoveryServicesProvidersClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, fabricName, providerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationRecoveryServicesProvidersClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationRecoveryServicesProvidersClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ReplicationRecoveryServicesProvidersClient) DeletePreparer(ctx context.Context, fabricName string, providerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"providerName":      autorest.Encode("path", providerName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders/{providerName}/remove", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationRecoveryServicesProvidersClient) DeleteSender(req *http.Request) (future ReplicationRecoveryServicesProvidersDeleteFuture, err error) {
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
func (client ReplicationRecoveryServicesProvidersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the details of registered recovery services provider.
// Parameters:
// fabricName - fabric name.
// providerName - recovery services provider name
func (client ReplicationRecoveryServicesProvidersClient) Get(ctx context.Context, fabricName string, providerName string) (result RecoveryServicesProvider, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationRecoveryServicesProvidersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, fabricName, providerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationRecoveryServicesProvidersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationRecoveryServicesProvidersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationRecoveryServicesProvidersClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ReplicationRecoveryServicesProvidersClient) GetPreparer(ctx context.Context, fabricName string, providerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"providerName":      autorest.Encode("path", providerName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders/{providerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationRecoveryServicesProvidersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReplicationRecoveryServicesProvidersClient) GetResponder(resp *http.Response) (result RecoveryServicesProvider, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the registered recovery services providers in the vault
func (client ReplicationRecoveryServicesProvidersClient) List(ctx context.Context) (result RecoveryServicesProviderCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationRecoveryServicesProvidersClient.List")
		defer func() {
			sc := -1
			if result.rspc.Response.Response != nil {
				sc = result.rspc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationRecoveryServicesProvidersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rspc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationRecoveryServicesProvidersClient", "List", resp, "Failure sending request")
		return
	}

	result.rspc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationRecoveryServicesProvidersClient", "List", resp, "Failure responding to request")
		return
	}
	if result.rspc.hasNextLink() && result.rspc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ReplicationRecoveryServicesProvidersClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationRecoveryServicesProviders", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationRecoveryServicesProvidersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ReplicationRecoveryServicesProvidersClient) ListResponder(resp *http.Response) (result RecoveryServicesProviderCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ReplicationRecoveryServicesProvidersClient) listNextResults(ctx context.Context, lastResults RecoveryServicesProviderCollection) (result RecoveryServicesProviderCollection, err error) {
	req, err := lastResults.recoveryServicesProviderCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationRecoveryServicesProvidersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationRecoveryServicesProvidersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationRecoveryServicesProvidersClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReplicationRecoveryServicesProvidersClient) ListComplete(ctx context.Context) (result RecoveryServicesProviderCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationRecoveryServicesProvidersClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByReplicationFabrics lists the registered recovery services providers for the specified fabric.
// Parameters:
// fabricName - fabric name
func (client ReplicationRecoveryServicesProvidersClient) ListByReplicationFabrics(ctx context.Context, fabricName string) (result RecoveryServicesProviderCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationRecoveryServicesProvidersClient.ListByReplicationFabrics")
		defer func() {
			sc := -1
			if result.rspc.Response.Response != nil {
				sc = result.rspc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByReplicationFabricsNextResults
	req, err := client.ListByReplicationFabricsPreparer(ctx, fabricName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationRecoveryServicesProvidersClient", "ListByReplicationFabrics", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByReplicationFabricsSender(req)
	if err != nil {
		result.rspc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationRecoveryServicesProvidersClient", "ListByReplicationFabrics", resp, "Failure sending request")
		return
	}

	result.rspc, err = client.ListByReplicationFabricsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationRecoveryServicesProvidersClient", "ListByReplicationFabrics", resp, "Failure responding to request")
		return
	}
	if result.rspc.hasNextLink() && result.rspc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByReplicationFabricsPreparer prepares the ListByReplicationFabrics request.
func (client ReplicationRecoveryServicesProvidersClient) ListByReplicationFabricsPreparer(ctx context.Context, fabricName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByReplicationFabricsSender sends the ListByReplicationFabrics request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationRecoveryServicesProvidersClient) ListByReplicationFabricsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByReplicationFabricsResponder handles the response to the ListByReplicationFabrics request. The method always
// closes the http.Response Body.
func (client ReplicationRecoveryServicesProvidersClient) ListByReplicationFabricsResponder(resp *http.Response) (result RecoveryServicesProviderCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByReplicationFabricsNextResults retrieves the next set of results, if any.
func (client ReplicationRecoveryServicesProvidersClient) listByReplicationFabricsNextResults(ctx context.Context, lastResults RecoveryServicesProviderCollection) (result RecoveryServicesProviderCollection, err error) {
	req, err := lastResults.recoveryServicesProviderCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationRecoveryServicesProvidersClient", "listByReplicationFabricsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByReplicationFabricsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationRecoveryServicesProvidersClient", "listByReplicationFabricsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByReplicationFabricsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationRecoveryServicesProvidersClient", "listByReplicationFabricsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByReplicationFabricsComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReplicationRecoveryServicesProvidersClient) ListByReplicationFabricsComplete(ctx context.Context, fabricName string) (result RecoveryServicesProviderCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationRecoveryServicesProvidersClient.ListByReplicationFabrics")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByReplicationFabrics(ctx, fabricName)
	return
}

// Purge the operation to purge(force delete) a recovery services provider from the vault.
// Parameters:
// fabricName - fabric name.
// providerName - recovery services provider name.
func (client ReplicationRecoveryServicesProvidersClient) Purge(ctx context.Context, fabricName string, providerName string) (result ReplicationRecoveryServicesProvidersPurgeFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationRecoveryServicesProvidersClient.Purge")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PurgePreparer(ctx, fabricName, providerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationRecoveryServicesProvidersClient", "Purge", nil, "Failure preparing request")
		return
	}

	result, err = client.PurgeSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationRecoveryServicesProvidersClient", "Purge", result.Response(), "Failure sending request")
		return
	}

	return
}

// PurgePreparer prepares the Purge request.
func (client ReplicationRecoveryServicesProvidersClient) PurgePreparer(ctx context.Context, fabricName string, providerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"providerName":      autorest.Encode("path", providerName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders/{providerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PurgeSender sends the Purge request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationRecoveryServicesProvidersClient) PurgeSender(req *http.Request) (future ReplicationRecoveryServicesProvidersPurgeFuture, err error) {
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

// PurgeResponder handles the response to the Purge request. The method always
// closes the http.Response Body.
func (client ReplicationRecoveryServicesProvidersClient) PurgeResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// RefreshProvider the operation to refresh the information from the recovery services provider.
// Parameters:
// fabricName - fabric name.
// providerName - recovery services provider name.
func (client ReplicationRecoveryServicesProvidersClient) RefreshProvider(ctx context.Context, fabricName string, providerName string) (result ReplicationRecoveryServicesProvidersRefreshProviderFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationRecoveryServicesProvidersClient.RefreshProvider")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RefreshProviderPreparer(ctx, fabricName, providerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationRecoveryServicesProvidersClient", "RefreshProvider", nil, "Failure preparing request")
		return
	}

	result, err = client.RefreshProviderSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationRecoveryServicesProvidersClient", "RefreshProvider", result.Response(), "Failure sending request")
		return
	}

	return
}

// RefreshProviderPreparer prepares the RefreshProvider request.
func (client ReplicationRecoveryServicesProvidersClient) RefreshProviderPreparer(ctx context.Context, fabricName string, providerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"providerName":      autorest.Encode("path", providerName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders/{providerName}/refreshProvider", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RefreshProviderSender sends the RefreshProvider request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationRecoveryServicesProvidersClient) RefreshProviderSender(req *http.Request) (future ReplicationRecoveryServicesProvidersRefreshProviderFuture, err error) {
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

// RefreshProviderResponder handles the response to the RefreshProvider request. The method always
// closes the http.Response Body.
func (client ReplicationRecoveryServicesProvidersClient) RefreshProviderResponder(resp *http.Response) (result RecoveryServicesProvider, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
