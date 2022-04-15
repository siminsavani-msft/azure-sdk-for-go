//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfeatures

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AuthorizationProfile.
func (a AuthorizationProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "approvedTime", a.ApprovedTime)
	populate(objectMap, "approver", a.Approver)
	populateTimeRFC3339(objectMap, "requestedTime", a.RequestedTime)
	populate(objectMap, "requester", a.Requester)
	populate(objectMap, "requesterObjectId", a.RequesterObjectID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AuthorizationProfile.
func (a *AuthorizationProfile) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "approvedTime":
			err = unpopulateTimeRFC3339(val, &a.ApprovedTime)
			delete(rawMsg, key)
		case "approver":
			err = unpopulate(val, &a.Approver)
			delete(rawMsg, key)
		case "requestedTime":
			err = unpopulateTimeRFC3339(val, &a.RequestedTime)
			delete(rawMsg, key)
		case "requester":
			err = unpopulate(val, &a.Requester)
			delete(rawMsg, key)
		case "requesterObjectId":
			err = unpopulate(val, &a.RequesterObjectID)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDefinition.
func (e ErrorDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type FeatureOperationsListResult.
func (f FeatureOperationsListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", f.NextLink)
	populate(objectMap, "value", f.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SubscriptionFeatureRegistrationList.
func (s SubscriptionFeatureRegistrationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SubscriptionFeatureRegistrationProperties.
func (s SubscriptionFeatureRegistrationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "approvalType", s.ApprovalType)
	populate(objectMap, "authorizationProfile", s.AuthorizationProfile)
	populate(objectMap, "description", s.Description)
	populate(objectMap, "displayName", s.DisplayName)
	populate(objectMap, "documentationLink", s.DocumentationLink)
	populate(objectMap, "featureName", s.FeatureName)
	populate(objectMap, "metadata", s.Metadata)
	populate(objectMap, "providerNamespace", s.ProviderNamespace)
	populateTimeRFC3339(objectMap, "registrationDate", s.RegistrationDate)
	populateTimeRFC3339(objectMap, "releaseDate", s.ReleaseDate)
	populate(objectMap, "shouldFeatureDisplayInPortal", s.ShouldFeatureDisplayInPortal)
	populate(objectMap, "state", s.State)
	populate(objectMap, "subscriptionId", s.SubscriptionID)
	populate(objectMap, "tenantId", s.TenantID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SubscriptionFeatureRegistrationProperties.
func (s *SubscriptionFeatureRegistrationProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "approvalType":
			err = unpopulate(val, &s.ApprovalType)
			delete(rawMsg, key)
		case "authorizationProfile":
			err = unpopulate(val, &s.AuthorizationProfile)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, &s.Description)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, &s.DisplayName)
			delete(rawMsg, key)
		case "documentationLink":
			err = unpopulate(val, &s.DocumentationLink)
			delete(rawMsg, key)
		case "featureName":
			err = unpopulate(val, &s.FeatureName)
			delete(rawMsg, key)
		case "metadata":
			err = unpopulate(val, &s.Metadata)
			delete(rawMsg, key)
		case "providerNamespace":
			err = unpopulate(val, &s.ProviderNamespace)
			delete(rawMsg, key)
		case "registrationDate":
			err = unpopulateTimeRFC3339(val, &s.RegistrationDate)
			delete(rawMsg, key)
		case "releaseDate":
			err = unpopulateTimeRFC3339(val, &s.ReleaseDate)
			delete(rawMsg, key)
		case "shouldFeatureDisplayInPortal":
			err = unpopulate(val, &s.ShouldFeatureDisplayInPortal)
			delete(rawMsg, key)
		case "state":
			err = unpopulate(val, &s.State)
			delete(rawMsg, key)
		case "subscriptionId":
			err = unpopulate(val, &s.SubscriptionID)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, &s.TenantID)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}