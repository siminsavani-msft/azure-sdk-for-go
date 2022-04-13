//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragepool

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ACL.
func (a ACL) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "initiatorIqn", a.InitiatorIqn)
	populate(objectMap, "mappedLuns", a.MappedLuns)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DiskPool.
func (d DiskPool) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", d.ID)
	populate(objectMap, "location", d.Location)
	populate(objectMap, "managedBy", d.ManagedBy)
	populate(objectMap, "managedByExtended", d.ManagedByExtended)
	populate(objectMap, "name", d.Name)
	populate(objectMap, "properties", d.Properties)
	populate(objectMap, "sku", d.SKU)
	populate(objectMap, "systemData", d.SystemData)
	populate(objectMap, "tags", d.Tags)
	populate(objectMap, "type", d.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DiskPoolCreate.
func (d DiskPoolCreate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", d.ID)
	populate(objectMap, "location", d.Location)
	populate(objectMap, "managedBy", d.ManagedBy)
	populate(objectMap, "managedByExtended", d.ManagedByExtended)
	populate(objectMap, "name", d.Name)
	populate(objectMap, "properties", d.Properties)
	populate(objectMap, "sku", d.SKU)
	populate(objectMap, "tags", d.Tags)
	populate(objectMap, "type", d.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DiskPoolCreateProperties.
func (d DiskPoolCreateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalCapabilities", d.AdditionalCapabilities)
	populate(objectMap, "availabilityZones", d.AvailabilityZones)
	populate(objectMap, "disks", d.Disks)
	populate(objectMap, "subnetId", d.SubnetID)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DiskPoolListResult.
func (d DiskPoolListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", d.NextLink)
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DiskPoolProperties.
func (d DiskPoolProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalCapabilities", d.AdditionalCapabilities)
	populate(objectMap, "availabilityZones", d.AvailabilityZones)
	populate(objectMap, "disks", d.Disks)
	populate(objectMap, "provisioningState", d.ProvisioningState)
	populate(objectMap, "status", d.Status)
	populate(objectMap, "subnetId", d.SubnetID)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DiskPoolUpdate.
func (d DiskPoolUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "managedBy", d.ManagedBy)
	populate(objectMap, "managedByExtended", d.ManagedByExtended)
	populate(objectMap, "properties", d.Properties)
	populate(objectMap, "sku", d.SKU)
	populate(objectMap, "tags", d.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DiskPoolUpdateProperties.
func (d DiskPoolUpdateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "disks", d.Disks)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DiskPoolZoneInfo.
func (d DiskPoolZoneInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalCapabilities", d.AdditionalCapabilities)
	populate(objectMap, "availabilityZones", d.AvailabilityZones)
	populate(objectMap, "sku", d.SKU)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DiskPoolZoneListResult.
func (d DiskPoolZoneListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", d.NextLink)
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type EndpointDependency.
func (e EndpointDependency) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "domainName", e.DomainName)
	populate(objectMap, "endpointDetails", e.EndpointDetails)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponse.
func (e ErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalInfo", e.AdditionalInfo)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type IscsiTarget.
func (i IscsiTarget) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", i.ID)
	populate(objectMap, "managedBy", i.ManagedBy)
	populate(objectMap, "managedByExtended", i.ManagedByExtended)
	populate(objectMap, "name", i.Name)
	populate(objectMap, "properties", i.Properties)
	populate(objectMap, "systemData", i.SystemData)
	populate(objectMap, "type", i.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type IscsiTargetCreate.
func (i IscsiTargetCreate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", i.ID)
	populate(objectMap, "managedBy", i.ManagedBy)
	populate(objectMap, "managedByExtended", i.ManagedByExtended)
	populate(objectMap, "name", i.Name)
	populate(objectMap, "properties", i.Properties)
	populate(objectMap, "type", i.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type IscsiTargetCreateProperties.
func (i IscsiTargetCreateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "aclMode", i.ACLMode)
	populate(objectMap, "luns", i.Luns)
	populate(objectMap, "staticAcls", i.StaticACLs)
	populate(objectMap, "targetIqn", i.TargetIqn)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type IscsiTargetList.
func (i IscsiTargetList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", i.NextLink)
	populate(objectMap, "value", i.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type IscsiTargetProperties.
func (i IscsiTargetProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "aclMode", i.ACLMode)
	populate(objectMap, "endpoints", i.Endpoints)
	populate(objectMap, "luns", i.Luns)
	populate(objectMap, "port", i.Port)
	populate(objectMap, "provisioningState", i.ProvisioningState)
	populate(objectMap, "sessions", i.Sessions)
	populate(objectMap, "staticAcls", i.StaticACLs)
	populate(objectMap, "status", i.Status)
	populate(objectMap, "targetIqn", i.TargetIqn)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type IscsiTargetUpdate.
func (i IscsiTargetUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", i.ID)
	populate(objectMap, "managedBy", i.ManagedBy)
	populate(objectMap, "managedByExtended", i.ManagedByExtended)
	populate(objectMap, "name", i.Name)
	populate(objectMap, "properties", i.Properties)
	populate(objectMap, "type", i.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type IscsiTargetUpdateProperties.
func (i IscsiTargetUpdateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "luns", i.Luns)
	populate(objectMap, "staticAcls", i.StaticACLs)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OutboundEnvironmentEndpoint.
func (o OutboundEnvironmentEndpoint) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "category", o.Category)
	populate(objectMap, "endpoints", o.Endpoints)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OutboundEnvironmentEndpointList.
func (o OutboundEnvironmentEndpointList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceSKUInfo.
func (r ResourceSKUInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "apiVersion", r.APIVersion)
	populate(objectMap, "capabilities", r.Capabilities)
	populate(objectMap, "locationInfo", r.LocationInfo)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "resourceType", r.ResourceType)
	populate(objectMap, "restrictions", r.Restrictions)
	populate(objectMap, "tier", r.Tier)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceSKUListResult.
func (r ResourceSKUListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceSKULocationInfo.
func (r ResourceSKULocationInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "location", r.Location)
	populate(objectMap, "zoneDetails", r.ZoneDetails)
	populate(objectMap, "zones", r.Zones)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceSKURestrictionInfo.
func (r ResourceSKURestrictionInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "locations", r.Locations)
	populate(objectMap, "zones", r.Zones)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceSKURestrictions.
func (r ResourceSKURestrictions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "reasonCode", r.ReasonCode)
	populate(objectMap, "restrictionInfo", r.RestrictionInfo)
	populate(objectMap, "type", r.Type)
	populate(objectMap, "values", r.Values)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceSKUZoneDetails.
func (r ResourceSKUZoneDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "capabilities", r.Capabilities)
	populate(objectMap, "name", r.Name)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SystemMetadata.
func (s SystemMetadata) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemMetadata.
func (s *SystemMetadata) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", t.ID)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
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