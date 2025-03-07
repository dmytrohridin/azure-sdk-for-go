//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtestbase

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

// CustomerEventsClient contains the methods for the CustomerEvents group.
// Don't use this type directly, use NewCustomerEventsClient() instead.
type CustomerEventsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCustomerEventsClient creates a new instance of CustomerEventsClient with the specified values.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCustomerEventsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CustomerEventsClient, error) {
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
	client := &CustomerEventsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Create or replace a Test Base Customer Event.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-12-16-preview
// resourceGroupName - The name of the resource group that contains the resource.
// testBaseAccountName - The resource name of the Test Base Account.
// customerEventName - The resource name of the Test Base Customer event.
// parameters - Parameters supplied to create a Test Base CustomerEvent.
// options - CustomerEventsClientBeginCreateOptions contains the optional parameters for the CustomerEventsClient.BeginCreate
// method.
func (client *CustomerEventsClient) BeginCreate(ctx context.Context, resourceGroupName string, testBaseAccountName string, customerEventName string, parameters CustomerEventResource, options *CustomerEventsClientBeginCreateOptions) (*runtime.Poller[CustomerEventsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, testBaseAccountName, customerEventName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[CustomerEventsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[CustomerEventsClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Create or replace a Test Base Customer Event.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-12-16-preview
func (client *CustomerEventsClient) create(ctx context.Context, resourceGroupName string, testBaseAccountName string, customerEventName string, parameters CustomerEventResource, options *CustomerEventsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, testBaseAccountName, customerEventName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *CustomerEventsClient) createCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, customerEventName string, parameters CustomerEventResource, options *CustomerEventsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/customerEvents/{customerEventName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	if customerEventName == "" {
		return nil, errors.New("parameter customerEventName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customerEventName}", url.PathEscape(customerEventName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a Test Base Customer Event.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-12-16-preview
// resourceGroupName - The name of the resource group that contains the resource.
// testBaseAccountName - The resource name of the Test Base Account.
// customerEventName - The resource name of the Test Base Customer event.
// options - CustomerEventsClientBeginDeleteOptions contains the optional parameters for the CustomerEventsClient.BeginDelete
// method.
func (client *CustomerEventsClient) BeginDelete(ctx context.Context, resourceGroupName string, testBaseAccountName string, customerEventName string, options *CustomerEventsClientBeginDeleteOptions) (*runtime.Poller[CustomerEventsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, testBaseAccountName, customerEventName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[CustomerEventsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[CustomerEventsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a Test Base Customer Event.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-12-16-preview
func (client *CustomerEventsClient) deleteOperation(ctx context.Context, resourceGroupName string, testBaseAccountName string, customerEventName string, options *CustomerEventsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, testBaseAccountName, customerEventName, options)
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
func (client *CustomerEventsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, customerEventName string, options *CustomerEventsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/customerEvents/{customerEventName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	if customerEventName == "" {
		return nil, errors.New("parameter customerEventName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customerEventName}", url.PathEscape(customerEventName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a Test Base CustomerEvent.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-12-16-preview
// resourceGroupName - The name of the resource group that contains the resource.
// testBaseAccountName - The resource name of the Test Base Account.
// customerEventName - The resource name of the Test Base Customer event.
// options - CustomerEventsClientGetOptions contains the optional parameters for the CustomerEventsClient.Get method.
func (client *CustomerEventsClient) Get(ctx context.Context, resourceGroupName string, testBaseAccountName string, customerEventName string, options *CustomerEventsClientGetOptions) (CustomerEventsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, testBaseAccountName, customerEventName, options)
	if err != nil {
		return CustomerEventsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CustomerEventsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CustomerEventsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CustomerEventsClient) getCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, customerEventName string, options *CustomerEventsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/customerEvents/{customerEventName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	if customerEventName == "" {
		return nil, errors.New("parameter customerEventName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customerEventName}", url.PathEscape(customerEventName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CustomerEventsClient) getHandleResponse(resp *http.Response) (CustomerEventsClientGetResponse, error) {
	result := CustomerEventsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomerEventResource); err != nil {
		return CustomerEventsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByTestBaseAccountPager - Lists all notification events subscribed under a Test Base Account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-12-16-preview
// resourceGroupName - The name of the resource group that contains the resource.
// testBaseAccountName - The resource name of the Test Base Account.
// options - CustomerEventsClientListByTestBaseAccountOptions contains the optional parameters for the CustomerEventsClient.ListByTestBaseAccount
// method.
func (client *CustomerEventsClient) NewListByTestBaseAccountPager(resourceGroupName string, testBaseAccountName string, options *CustomerEventsClientListByTestBaseAccountOptions) *runtime.Pager[CustomerEventsClientListByTestBaseAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[CustomerEventsClientListByTestBaseAccountResponse]{
		More: func(page CustomerEventsClientListByTestBaseAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CustomerEventsClientListByTestBaseAccountResponse) (CustomerEventsClientListByTestBaseAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByTestBaseAccountCreateRequest(ctx, resourceGroupName, testBaseAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CustomerEventsClientListByTestBaseAccountResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return CustomerEventsClientListByTestBaseAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CustomerEventsClientListByTestBaseAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByTestBaseAccountHandleResponse(resp)
		},
	})
}

// listByTestBaseAccountCreateRequest creates the ListByTestBaseAccount request.
func (client *CustomerEventsClient) listByTestBaseAccountCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *CustomerEventsClientListByTestBaseAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/customerEvents"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByTestBaseAccountHandleResponse handles the ListByTestBaseAccount response.
func (client *CustomerEventsClient) listByTestBaseAccountHandleResponse(resp *http.Response) (CustomerEventsClientListByTestBaseAccountResponse, error) {
	result := CustomerEventsClientListByTestBaseAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomerEventListResult); err != nil {
		return CustomerEventsClientListByTestBaseAccountResponse{}, err
	}
	return result, nil
}
