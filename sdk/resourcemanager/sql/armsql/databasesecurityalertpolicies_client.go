//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// DatabaseSecurityAlertPoliciesClient contains the methods for the DatabaseSecurityAlertPolicies group.
// Don't use this type directly, use NewDatabaseSecurityAlertPoliciesClient() instead.
type DatabaseSecurityAlertPoliciesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDatabaseSecurityAlertPoliciesClient creates a new instance of DatabaseSecurityAlertPoliciesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDatabaseSecurityAlertPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DatabaseSecurityAlertPoliciesClient, error) {
	cl, err := arm.NewClient(moduleName+".DatabaseSecurityAlertPoliciesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DatabaseSecurityAlertPoliciesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a database's security alert policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database for which the security alert policy is defined.
//   - securityAlertPolicyName - The name of the security alert policy.
//   - parameters - The database security alert policy.
//   - options - DatabaseSecurityAlertPoliciesClientCreateOrUpdateOptions contains the optional parameters for the DatabaseSecurityAlertPoliciesClient.CreateOrUpdate
//     method.
func (client *DatabaseSecurityAlertPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, securityAlertPolicyName SecurityAlertPolicyName, parameters DatabaseSecurityAlertPolicy, options *DatabaseSecurityAlertPoliciesClientCreateOrUpdateOptions) (DatabaseSecurityAlertPoliciesClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, databaseName, securityAlertPolicyName, parameters, options)
	if err != nil {
		return DatabaseSecurityAlertPoliciesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatabaseSecurityAlertPoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return DatabaseSecurityAlertPoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DatabaseSecurityAlertPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, securityAlertPolicyName SecurityAlertPolicyName, parameters DatabaseSecurityAlertPolicy, options *DatabaseSecurityAlertPoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/securityAlertPolicies/{securityAlertPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if securityAlertPolicyName == "" {
		return nil, errors.New("parameter securityAlertPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityAlertPolicyName}", url.PathEscape(string(securityAlertPolicyName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DatabaseSecurityAlertPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (DatabaseSecurityAlertPoliciesClientCreateOrUpdateResponse, error) {
	result := DatabaseSecurityAlertPoliciesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseSecurityAlertPolicy); err != nil {
		return DatabaseSecurityAlertPoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Get - Gets a database's security alert policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database for which the security alert policy is defined.
//   - securityAlertPolicyName - The name of the security alert policy.
//   - options - DatabaseSecurityAlertPoliciesClientGetOptions contains the optional parameters for the DatabaseSecurityAlertPoliciesClient.Get
//     method.
func (client *DatabaseSecurityAlertPoliciesClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, securityAlertPolicyName SecurityAlertPolicyName, options *DatabaseSecurityAlertPoliciesClientGetOptions) (DatabaseSecurityAlertPoliciesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, databaseName, securityAlertPolicyName, options)
	if err != nil {
		return DatabaseSecurityAlertPoliciesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatabaseSecurityAlertPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DatabaseSecurityAlertPoliciesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DatabaseSecurityAlertPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, securityAlertPolicyName SecurityAlertPolicyName, options *DatabaseSecurityAlertPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/securityAlertPolicies/{securityAlertPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if securityAlertPolicyName == "" {
		return nil, errors.New("parameter securityAlertPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityAlertPolicyName}", url.PathEscape(string(securityAlertPolicyName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DatabaseSecurityAlertPoliciesClient) getHandleResponse(resp *http.Response) (DatabaseSecurityAlertPoliciesClientGetResponse, error) {
	result := DatabaseSecurityAlertPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseSecurityAlertPolicy); err != nil {
		return DatabaseSecurityAlertPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDatabasePager - Gets a list of database's security alert policies.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database for which the security alert policy is defined.
//   - options - DatabaseSecurityAlertPoliciesClientListByDatabaseOptions contains the optional parameters for the DatabaseSecurityAlertPoliciesClient.NewListByDatabasePager
//     method.
func (client *DatabaseSecurityAlertPoliciesClient) NewListByDatabasePager(resourceGroupName string, serverName string, databaseName string, options *DatabaseSecurityAlertPoliciesClientListByDatabaseOptions) *runtime.Pager[DatabaseSecurityAlertPoliciesClientListByDatabaseResponse] {
	return runtime.NewPager(runtime.PagingHandler[DatabaseSecurityAlertPoliciesClientListByDatabaseResponse]{
		More: func(page DatabaseSecurityAlertPoliciesClientListByDatabaseResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DatabaseSecurityAlertPoliciesClientListByDatabaseResponse) (DatabaseSecurityAlertPoliciesClientListByDatabaseResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByDatabaseCreateRequest(ctx, resourceGroupName, serverName, databaseName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DatabaseSecurityAlertPoliciesClientListByDatabaseResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DatabaseSecurityAlertPoliciesClientListByDatabaseResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DatabaseSecurityAlertPoliciesClientListByDatabaseResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDatabaseHandleResponse(resp)
		},
	})
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *DatabaseSecurityAlertPoliciesClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, options *DatabaseSecurityAlertPoliciesClientListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/securityAlertPolicies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *DatabaseSecurityAlertPoliciesClient) listByDatabaseHandleResponse(resp *http.Response) (DatabaseSecurityAlertPoliciesClientListByDatabaseResponse, error) {
	result := DatabaseSecurityAlertPoliciesClientListByDatabaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseSecurityAlertListResult); err != nil {
		return DatabaseSecurityAlertPoliciesClientListByDatabaseResponse{}, err
	}
	return result, nil
}
