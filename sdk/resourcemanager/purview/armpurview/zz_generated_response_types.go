//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpurview

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// AccountsAddRootCollectionAdminResponse contains the response from method Accounts.AddRootCollectionAdmin.
type AccountsAddRootCollectionAdminResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsCheckNameAvailabilityResponse contains the response from method Accounts.CheckNameAvailability.
type AccountsCheckNameAvailabilityResponse struct {
	AccountsCheckNameAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsCheckNameAvailabilityResult contains the result from method Accounts.CheckNameAvailability.
type AccountsCheckNameAvailabilityResult struct {
	CheckNameAvailabilityResult
}

// AccountsCreateOrUpdatePollerResponse contains the response from method Accounts.CreateOrUpdate.
type AccountsCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *AccountsCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l AccountsCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (AccountsCreateOrUpdateResponse, error) {
	respType := AccountsCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Account)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a AccountsCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *AccountsCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *AccountsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("AccountsClient.CreateOrUpdate", token, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return err
	}
	poller := &AccountsCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// AccountsCreateOrUpdateResponse contains the response from method Accounts.CreateOrUpdate.
type AccountsCreateOrUpdateResponse struct {
	AccountsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsCreateOrUpdateResult contains the result from method Accounts.CreateOrUpdate.
type AccountsCreateOrUpdateResult struct {
	Account
}

// AccountsDeletePollerResponse contains the response from method Accounts.Delete.
type AccountsDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *AccountsDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l AccountsDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (AccountsDeleteResponse, error) {
	respType := AccountsDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a AccountsDeletePollerResponse from the provided client and resume token.
func (l *AccountsDeletePollerResponse) Resume(ctx context.Context, client *AccountsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("AccountsClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &AccountsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// AccountsDeleteResponse contains the response from method Accounts.Delete.
type AccountsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsGetResponse contains the response from method Accounts.Get.
type AccountsGetResponse struct {
	AccountsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsGetResult contains the result from method Accounts.Get.
type AccountsGetResult struct {
	Account
}

// AccountsListByResourceGroupResponse contains the response from method Accounts.ListByResourceGroup.
type AccountsListByResourceGroupResponse struct {
	AccountsListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsListByResourceGroupResult contains the result from method Accounts.ListByResourceGroup.
type AccountsListByResourceGroupResult struct {
	AccountList
}

// AccountsListBySubscriptionResponse contains the response from method Accounts.ListBySubscription.
type AccountsListBySubscriptionResponse struct {
	AccountsListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsListBySubscriptionResult contains the result from method Accounts.ListBySubscription.
type AccountsListBySubscriptionResult struct {
	AccountList
}

// AccountsListKeysResponse contains the response from method Accounts.ListKeys.
type AccountsListKeysResponse struct {
	AccountsListKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsListKeysResult contains the result from method Accounts.ListKeys.
type AccountsListKeysResult struct {
	AccessKeys
}

// AccountsUpdatePollerResponse contains the response from method Accounts.Update.
type AccountsUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *AccountsUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l AccountsUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (AccountsUpdateResponse, error) {
	respType := AccountsUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Account)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a AccountsUpdatePollerResponse from the provided client and resume token.
func (l *AccountsUpdatePollerResponse) Resume(ctx context.Context, client *AccountsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("AccountsClient.Update", token, client.pl, client.updateHandleError)
	if err != nil {
		return err
	}
	poller := &AccountsUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// AccountsUpdateResponse contains the response from method Accounts.Update.
type AccountsUpdateResponse struct {
	AccountsUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsUpdateResult contains the result from method Accounts.Update.
type AccountsUpdateResult struct {
	Account
}

// DefaultAccountsGetResponse contains the response from method DefaultAccounts.Get.
type DefaultAccountsGetResponse struct {
	DefaultAccountsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DefaultAccountsGetResult contains the result from method DefaultAccounts.Get.
type DefaultAccountsGetResult struct {
	DefaultAccountPayload
}

// DefaultAccountsRemoveResponse contains the response from method DefaultAccounts.Remove.
type DefaultAccountsRemoveResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DefaultAccountsSetResponse contains the response from method DefaultAccounts.Set.
type DefaultAccountsSetResponse struct {
	DefaultAccountsSetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DefaultAccountsSetResult contains the result from method DefaultAccounts.Set.
type DefaultAccountsSetResult struct {
	DefaultAccountPayload
}

// OperationsListResponse contains the response from method Operations.List.
type OperationsListResponse struct {
	OperationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsListResult contains the result from method Operations.List.
type OperationsListResult struct {
	OperationList
}

// PrivateEndpointConnectionsCreateOrUpdatePollerResponse contains the response from method PrivateEndpointConnections.CreateOrUpdate.
type PrivateEndpointConnectionsCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateEndpointConnectionsCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l PrivateEndpointConnectionsCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateEndpointConnectionsCreateOrUpdateResponse, error) {
	respType := PrivateEndpointConnectionsCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.PrivateEndpointConnection)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PrivateEndpointConnectionsCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *PrivateEndpointConnectionsCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *PrivateEndpointConnectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateEndpointConnectionsClient.CreateOrUpdate", token, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return err
	}
	poller := &PrivateEndpointConnectionsCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PrivateEndpointConnectionsCreateOrUpdateResponse contains the response from method PrivateEndpointConnections.CreateOrUpdate.
type PrivateEndpointConnectionsCreateOrUpdateResponse struct {
	PrivateEndpointConnectionsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsCreateOrUpdateResult contains the result from method PrivateEndpointConnections.CreateOrUpdate.
type PrivateEndpointConnectionsCreateOrUpdateResult struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsDeletePollerResponse contains the response from method PrivateEndpointConnections.Delete.
type PrivateEndpointConnectionsDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateEndpointConnectionsDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l PrivateEndpointConnectionsDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateEndpointConnectionsDeleteResponse, error) {
	respType := PrivateEndpointConnectionsDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PrivateEndpointConnectionsDeletePollerResponse from the provided client and resume token.
func (l *PrivateEndpointConnectionsDeletePollerResponse) Resume(ctx context.Context, client *PrivateEndpointConnectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateEndpointConnectionsClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &PrivateEndpointConnectionsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PrivateEndpointConnectionsDeleteResponse contains the response from method PrivateEndpointConnections.Delete.
type PrivateEndpointConnectionsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsGetResponse contains the response from method PrivateEndpointConnections.Get.
type PrivateEndpointConnectionsGetResponse struct {
	PrivateEndpointConnectionsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsGetResult contains the result from method PrivateEndpointConnections.Get.
type PrivateEndpointConnectionsGetResult struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsListByAccountResponse contains the response from method PrivateEndpointConnections.ListByAccount.
type PrivateEndpointConnectionsListByAccountResponse struct {
	PrivateEndpointConnectionsListByAccountResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsListByAccountResult contains the result from method PrivateEndpointConnections.ListByAccount.
type PrivateEndpointConnectionsListByAccountResult struct {
	PrivateEndpointConnectionList
}

// PrivateLinkResourcesGetByGroupIDResponse contains the response from method PrivateLinkResources.GetByGroupID.
type PrivateLinkResourcesGetByGroupIDResponse struct {
	PrivateLinkResourcesGetByGroupIDResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateLinkResourcesGetByGroupIDResult contains the result from method PrivateLinkResources.GetByGroupID.
type PrivateLinkResourcesGetByGroupIDResult struct {
	PrivateLinkResource
}

// PrivateLinkResourcesListByAccountResponse contains the response from method PrivateLinkResources.ListByAccount.
type PrivateLinkResourcesListByAccountResponse struct {
	PrivateLinkResourcesListByAccountResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateLinkResourcesListByAccountResult contains the result from method PrivateLinkResources.ListByAccount.
type PrivateLinkResourcesListByAccountResult struct {
	PrivateLinkResourceList
}