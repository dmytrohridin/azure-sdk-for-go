//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlabservices

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

// LabsClient contains the methods for the Labs group.
// Don't use this type directly, use NewLabsClient() instead.
type LabsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLabsClient creates a new instance of LabsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLabsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LabsClient, error) {
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
	client := &LabsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Operation to create or update a lab resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-15-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// labName - The name of the lab that uniquely identifies it within containing lab account. Used in resource URIs.
// body - The request body.
// options - LabsClientBeginCreateOrUpdateOptions contains the optional parameters for the LabsClient.BeginCreateOrUpdate
// method.
func (client *LabsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, body Lab, options *LabsClientBeginCreateOrUpdateOptions) (*runtime.Poller[LabsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, labName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[LabsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaOriginalURI,
		})
	} else {
		return runtime.NewPollerFromResumeToken[LabsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Operation to create or update a lab resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-15-preview
func (client *LabsClient) createOrUpdate(ctx context.Context, resourceGroupName string, labName string, body Lab, options *LabsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, labName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LabsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, labName string, body Lab, options *LabsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Operation to delete a lab resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-15-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// labName - The name of the lab that uniquely identifies it within containing lab account. Used in resource URIs.
// options - LabsClientBeginDeleteOptions contains the optional parameters for the LabsClient.BeginDelete method.
func (client *LabsClient) BeginDelete(ctx context.Context, resourceGroupName string, labName string, options *LabsClientBeginDeleteOptions) (*runtime.Poller[LabsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, labName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[LabsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[LabsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Operation to delete a lab resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-15-preview
func (client *LabsClient) deleteOperation(ctx context.Context, resourceGroupName string, labName string, options *LabsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, labName, options)
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
func (client *LabsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, labName string, options *LabsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns the properties of a lab resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-15-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// labName - The name of the lab that uniquely identifies it within containing lab account. Used in resource URIs.
// options - LabsClientGetOptions contains the optional parameters for the LabsClient.Get method.
func (client *LabsClient) Get(ctx context.Context, resourceGroupName string, labName string, options *LabsClientGetOptions) (LabsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, labName, options)
	if err != nil {
		return LabsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LabsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LabsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LabsClient) getCreateRequest(ctx context.Context, resourceGroupName string, labName string, options *LabsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LabsClient) getHandleResponse(resp *http.Response) (LabsClientGetResponse, error) {
	result := LabsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Lab); err != nil {
		return LabsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Returns a list of all labs in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-15-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - LabsClientListByResourceGroupOptions contains the optional parameters for the LabsClient.ListByResourceGroup
// method.
func (client *LabsClient) NewListByResourceGroupPager(resourceGroupName string, options *LabsClientListByResourceGroupOptions) *runtime.Pager[LabsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[LabsClientListByResourceGroupResponse]{
		More: func(page LabsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LabsClientListByResourceGroupResponse) (LabsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LabsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return LabsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LabsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *LabsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *LabsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *LabsClient) listByResourceGroupHandleResponse(resp *http.Response) (LabsClientListByResourceGroupResponse, error) {
	result := LabsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PagedLabs); err != nil {
		return LabsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Returns a list of all labs for a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-15-preview
// options - LabsClientListBySubscriptionOptions contains the optional parameters for the LabsClient.ListBySubscription method.
func (client *LabsClient) NewListBySubscriptionPager(options *LabsClientListBySubscriptionOptions) *runtime.Pager[LabsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[LabsClientListBySubscriptionResponse]{
		More: func(page LabsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LabsClientListBySubscriptionResponse) (LabsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LabsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return LabsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LabsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *LabsClient) listBySubscriptionCreateRequest(ctx context.Context, options *LabsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.LabServices/labs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *LabsClient) listBySubscriptionHandleResponse(resp *http.Response) (LabsClientListBySubscriptionResponse, error) {
	result := LabsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PagedLabs); err != nil {
		return LabsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginPublish - Publish or re-publish a lab. This will create or update all lab resources, such as virtual machines.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-15-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// labName - The name of the lab that uniquely identifies it within containing lab account. Used in resource URIs.
// options - LabsClientBeginPublishOptions contains the optional parameters for the LabsClient.BeginPublish method.
func (client *LabsClient) BeginPublish(ctx context.Context, resourceGroupName string, labName string, options *LabsClientBeginPublishOptions) (*runtime.Poller[LabsClientPublishResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.publish(ctx, resourceGroupName, labName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[LabsClientPublishResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[LabsClientPublishResponse](options.ResumeToken, client.pl, nil)
	}
}

// Publish - Publish or re-publish a lab. This will create or update all lab resources, such as virtual machines.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-15-preview
func (client *LabsClient) publish(ctx context.Context, resourceGroupName string, labName string, options *LabsClientBeginPublishOptions) (*http.Response, error) {
	req, err := client.publishCreateRequest(ctx, resourceGroupName, labName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// publishCreateRequest creates the Publish request.
func (client *LabsClient) publishCreateRequest(ctx context.Context, resourceGroupName string, labName string, options *LabsClientBeginPublishOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/publish"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginSyncGroup - Action used to manually kick off an AAD group sync job.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-15-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// labName - The name of the lab that uniquely identifies it within containing lab account. Used in resource URIs.
// options - LabsClientBeginSyncGroupOptions contains the optional parameters for the LabsClient.BeginSyncGroup method.
func (client *LabsClient) BeginSyncGroup(ctx context.Context, resourceGroupName string, labName string, options *LabsClientBeginSyncGroupOptions) (*runtime.Poller[LabsClientSyncGroupResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.syncGroup(ctx, resourceGroupName, labName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[LabsClientSyncGroupResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[LabsClientSyncGroupResponse](options.ResumeToken, client.pl, nil)
	}
}

// SyncGroup - Action used to manually kick off an AAD group sync job.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-15-preview
func (client *LabsClient) syncGroup(ctx context.Context, resourceGroupName string, labName string, options *LabsClientBeginSyncGroupOptions) (*http.Response, error) {
	req, err := client.syncGroupCreateRequest(ctx, resourceGroupName, labName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// syncGroupCreateRequest creates the SyncGroup request.
func (client *LabsClient) syncGroupCreateRequest(ctx context.Context, resourceGroupName string, labName string, options *LabsClientBeginSyncGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/syncGroup"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginUpdate - Operation to update a lab resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-15-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// labName - The name of the lab that uniquely identifies it within containing lab account. Used in resource URIs.
// body - The request body.
// options - LabsClientBeginUpdateOptions contains the optional parameters for the LabsClient.BeginUpdate method.
func (client *LabsClient) BeginUpdate(ctx context.Context, resourceGroupName string, labName string, body LabUpdate, options *LabsClientBeginUpdateOptions) (*runtime.Poller[LabsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, labName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[LabsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[LabsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Operation to update a lab resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-15-preview
func (client *LabsClient) update(ctx context.Context, resourceGroupName string, labName string, body LabUpdate, options *LabsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, labName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *LabsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, labName string, body LabUpdate, options *LabsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}
