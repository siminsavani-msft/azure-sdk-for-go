//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// BlobServicesClient contains the methods for the BlobServices group.
// Don't use this type directly, use NewBlobServicesClient() instead.
type BlobServicesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewBlobServicesClient creates a new instance of BlobServicesClient with the specified values.
func NewBlobServicesClient(con *arm.Connection, subscriptionID string) *BlobServicesClient {
	return &BlobServicesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// GetServiceProperties - Gets the properties of a storage account’s Blob service, including properties for Storage Analytics and CORS (Cross-Origin Resource
// Sharing) rules.
// If the operation fails it returns a generic error.
func (client *BlobServicesClient) GetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, options *BlobServicesGetServicePropertiesOptions) (BlobServicesGetServicePropertiesResponse, error) {
	req, err := client.getServicePropertiesCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return BlobServicesGetServicePropertiesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BlobServicesGetServicePropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BlobServicesGetServicePropertiesResponse{}, client.getServicePropertiesHandleError(resp)
	}
	return client.getServicePropertiesHandleResponse(resp)
}

// getServicePropertiesCreateRequest creates the GetServiceProperties request.
func (client *BlobServicesClient) getServicePropertiesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *BlobServicesGetServicePropertiesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/{BlobServicesName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{BlobServicesName}", url.PathEscape("default"))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getServicePropertiesHandleResponse handles the GetServiceProperties response.
func (client *BlobServicesClient) getServicePropertiesHandleResponse(resp *http.Response) (BlobServicesGetServicePropertiesResponse, error) {
	result := BlobServicesGetServicePropertiesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BlobServiceProperties); err != nil {
		return BlobServicesGetServicePropertiesResponse{}, err
	}
	return result, nil
}

// getServicePropertiesHandleError handles the GetServiceProperties error response.
func (client *BlobServicesClient) getServicePropertiesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// List - List blob services of storage account. It returns a collection of one object named default.
// If the operation fails it returns a generic error.
func (client *BlobServicesClient) List(ctx context.Context, resourceGroupName string, accountName string, options *BlobServicesListOptions) (BlobServicesListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return BlobServicesListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BlobServicesListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BlobServicesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *BlobServicesClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *BlobServicesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BlobServicesClient) listHandleResponse(resp *http.Response) (BlobServicesListResponse, error) {
	result := BlobServicesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BlobServiceItems); err != nil {
		return BlobServicesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *BlobServicesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// SetServiceProperties - Sets the properties of a storage account’s Blob service, including properties for Storage Analytics and CORS (Cross-Origin Resource
// Sharing) rules.
// If the operation fails it returns a generic error.
func (client *BlobServicesClient) SetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, parameters BlobServiceProperties, options *BlobServicesSetServicePropertiesOptions) (BlobServicesSetServicePropertiesResponse, error) {
	req, err := client.setServicePropertiesCreateRequest(ctx, resourceGroupName, accountName, parameters, options)
	if err != nil {
		return BlobServicesSetServicePropertiesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BlobServicesSetServicePropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BlobServicesSetServicePropertiesResponse{}, client.setServicePropertiesHandleError(resp)
	}
	return client.setServicePropertiesHandleResponse(resp)
}

// setServicePropertiesCreateRequest creates the SetServiceProperties request.
func (client *BlobServicesClient) setServicePropertiesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, parameters BlobServiceProperties, options *BlobServicesSetServicePropertiesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/{BlobServicesName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{BlobServicesName}", url.PathEscape("default"))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// setServicePropertiesHandleResponse handles the SetServiceProperties response.
func (client *BlobServicesClient) setServicePropertiesHandleResponse(resp *http.Response) (BlobServicesSetServicePropertiesResponse, error) {
	result := BlobServicesSetServicePropertiesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BlobServiceProperties); err != nil {
		return BlobServicesSetServicePropertiesResponse{}, err
	}
	return result, nil
}

// setServicePropertiesHandleError handles the SetServiceProperties error response.
func (client *BlobServicesClient) setServicePropertiesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}