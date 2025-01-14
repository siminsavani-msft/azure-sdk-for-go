//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmanagednetworkfabric

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// IPExtendedCommunitiesClient contains the methods for the IPExtendedCommunities group.
// Don't use this type directly, use NewIPExtendedCommunitiesClient() instead.
type IPExtendedCommunitiesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewIPExtendedCommunitiesClient creates a new instance of IPExtendedCommunitiesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIPExtendedCommunitiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IPExtendedCommunitiesClient, error) {
	cl, err := arm.NewClient(moduleName+".IPExtendedCommunitiesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &IPExtendedCommunitiesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Implements IP Extended Community PUT method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ipExtendedCommunityName - Name of the IP Extended Community.
//   - body - Request payload.
//   - options - IPExtendedCommunitiesClientBeginCreateOptions contains the optional parameters for the IPExtendedCommunitiesClient.BeginCreate
//     method.
func (client *IPExtendedCommunitiesClient) BeginCreate(ctx context.Context, resourceGroupName string, ipExtendedCommunityName string, body IPExtendedCommunity, options *IPExtendedCommunitiesClientBeginCreateOptions) (*runtime.Poller[IPExtendedCommunitiesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, ipExtendedCommunityName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[IPExtendedCommunitiesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[IPExtendedCommunitiesClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Implements IP Extended Community PUT method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *IPExtendedCommunitiesClient) create(ctx context.Context, resourceGroupName string, ipExtendedCommunityName string, body IPExtendedCommunity, options *IPExtendedCommunitiesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, ipExtendedCommunityName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *IPExtendedCommunitiesClient) createCreateRequest(ctx context.Context, resourceGroupName string, ipExtendedCommunityName string, body IPExtendedCommunity, options *IPExtendedCommunitiesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/{ipExtendedCommunityName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipExtendedCommunityName == "" {
		return nil, errors.New("parameter ipExtendedCommunityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipExtendedCommunityName}", url.PathEscape(ipExtendedCommunityName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Implements IP Extended Community DELETE method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ipExtendedCommunityName - Name of the IP Extended Community.
//   - options - IPExtendedCommunitiesClientBeginDeleteOptions contains the optional parameters for the IPExtendedCommunitiesClient.BeginDelete
//     method.
func (client *IPExtendedCommunitiesClient) BeginDelete(ctx context.Context, resourceGroupName string, ipExtendedCommunityName string, options *IPExtendedCommunitiesClientBeginDeleteOptions) (*runtime.Poller[IPExtendedCommunitiesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, ipExtendedCommunityName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[IPExtendedCommunitiesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[IPExtendedCommunitiesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Implements IP Extended Community DELETE method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *IPExtendedCommunitiesClient) deleteOperation(ctx context.Context, resourceGroupName string, ipExtendedCommunityName string, options *IPExtendedCommunitiesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, ipExtendedCommunityName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IPExtendedCommunitiesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, ipExtendedCommunityName string, options *IPExtendedCommunitiesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/{ipExtendedCommunityName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipExtendedCommunityName == "" {
		return nil, errors.New("parameter ipExtendedCommunityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipExtendedCommunityName}", url.PathEscape(ipExtendedCommunityName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Implements IP Extended Community GET method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ipExtendedCommunityName - Name of the IP Extended Community.
//   - options - IPExtendedCommunitiesClientGetOptions contains the optional parameters for the IPExtendedCommunitiesClient.Get
//     method.
func (client *IPExtendedCommunitiesClient) Get(ctx context.Context, resourceGroupName string, ipExtendedCommunityName string, options *IPExtendedCommunitiesClientGetOptions) (IPExtendedCommunitiesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, ipExtendedCommunityName, options)
	if err != nil {
		return IPExtendedCommunitiesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IPExtendedCommunitiesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IPExtendedCommunitiesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IPExtendedCommunitiesClient) getCreateRequest(ctx context.Context, resourceGroupName string, ipExtendedCommunityName string, options *IPExtendedCommunitiesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/{ipExtendedCommunityName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipExtendedCommunityName == "" {
		return nil, errors.New("parameter ipExtendedCommunityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipExtendedCommunityName}", url.PathEscape(ipExtendedCommunityName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IPExtendedCommunitiesClient) getHandleResponse(resp *http.Response) (IPExtendedCommunitiesClientGetResponse, error) {
	result := IPExtendedCommunitiesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPExtendedCommunity); err != nil {
		return IPExtendedCommunitiesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Implements IpExtendedCommunities list by resource group GET method.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - IPExtendedCommunitiesClientListByResourceGroupOptions contains the optional parameters for the IPExtendedCommunitiesClient.NewListByResourceGroupPager
//     method.
func (client *IPExtendedCommunitiesClient) NewListByResourceGroupPager(resourceGroupName string, options *IPExtendedCommunitiesClientListByResourceGroupOptions) *runtime.Pager[IPExtendedCommunitiesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[IPExtendedCommunitiesClientListByResourceGroupResponse]{
		More: func(page IPExtendedCommunitiesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IPExtendedCommunitiesClientListByResourceGroupResponse) (IPExtendedCommunitiesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return IPExtendedCommunitiesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return IPExtendedCommunitiesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IPExtendedCommunitiesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *IPExtendedCommunitiesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *IPExtendedCommunitiesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *IPExtendedCommunitiesClient) listByResourceGroupHandleResponse(resp *http.Response) (IPExtendedCommunitiesClientListByResourceGroupResponse, error) {
	result := IPExtendedCommunitiesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPExtendedCommunityListResult); err != nil {
		return IPExtendedCommunitiesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Implements IpExtendedCommunities list by subscription GET method.
//
// Generated from API version 2023-06-15
//   - options - IPExtendedCommunitiesClientListBySubscriptionOptions contains the optional parameters for the IPExtendedCommunitiesClient.NewListBySubscriptionPager
//     method.
func (client *IPExtendedCommunitiesClient) NewListBySubscriptionPager(options *IPExtendedCommunitiesClientListBySubscriptionOptions) *runtime.Pager[IPExtendedCommunitiesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[IPExtendedCommunitiesClientListBySubscriptionResponse]{
		More: func(page IPExtendedCommunitiesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IPExtendedCommunitiesClientListBySubscriptionResponse) (IPExtendedCommunitiesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return IPExtendedCommunitiesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return IPExtendedCommunitiesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IPExtendedCommunitiesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *IPExtendedCommunitiesClient) listBySubscriptionCreateRequest(ctx context.Context, options *IPExtendedCommunitiesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *IPExtendedCommunitiesClient) listBySubscriptionHandleResponse(resp *http.Response) (IPExtendedCommunitiesClientListBySubscriptionResponse, error) {
	result := IPExtendedCommunitiesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPExtendedCommunityListResult); err != nil {
		return IPExtendedCommunitiesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - API to update certain properties of the IP Extended Community resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ipExtendedCommunityName - Name of the IP Extended Community.
//   - body - IP Extended Community properties to update.
//   - options - IPExtendedCommunitiesClientBeginUpdateOptions contains the optional parameters for the IPExtendedCommunitiesClient.BeginUpdate
//     method.
func (client *IPExtendedCommunitiesClient) BeginUpdate(ctx context.Context, resourceGroupName string, ipExtendedCommunityName string, body IPExtendedCommunityPatch, options *IPExtendedCommunitiesClientBeginUpdateOptions) (*runtime.Poller[IPExtendedCommunitiesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, ipExtendedCommunityName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[IPExtendedCommunitiesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[IPExtendedCommunitiesClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - API to update certain properties of the IP Extended Community resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *IPExtendedCommunitiesClient) update(ctx context.Context, resourceGroupName string, ipExtendedCommunityName string, body IPExtendedCommunityPatch, options *IPExtendedCommunitiesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, ipExtendedCommunityName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *IPExtendedCommunitiesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, ipExtendedCommunityName string, body IPExtendedCommunityPatch, options *IPExtendedCommunitiesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/{ipExtendedCommunityName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipExtendedCommunityName == "" {
		return nil, errors.New("parameter ipExtendedCommunityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipExtendedCommunityName}", url.PathEscape(ipExtendedCommunityName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}
