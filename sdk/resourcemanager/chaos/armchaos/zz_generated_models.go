//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armchaos

import "time"

// ActionClassification provides polymorphic access to related types.
// Call the interface's GetAction() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *Action, *ContinuousAction, *DelayAction, *DiscreteAction
type ActionClassification interface {
	// GetAction returns the Action content of the underlying type.
	GetAction() *Action
}

// Action - Model that represents the base action model.
type Action struct {
	// REQUIRED; String that represents a Capability URN.
	Name *string `json:"name,omitempty"`

	// REQUIRED; Enum that discriminates between action models.
	Type *string `json:"type,omitempty"`
}

// GetAction implements the ActionClassification interface for type Action.
func (a *Action) GetAction() *Action { return a }

// ActionStatus - Model that represents the an action and its status.
type ActionStatus struct {
	// READ-ONLY; The id of the action status.
	ActionID *string `json:"actionId,omitempty" azure:"ro"`

	// READ-ONLY; The name of the action status.
	ActionName *string `json:"actionName,omitempty" azure:"ro"`

	// READ-ONLY; String that represents the end time of the action.
	EndTime *time.Time `json:"endTime,omitempty" azure:"ro"`

	// READ-ONLY; String that represents the start time of the action.
	StartTime *time.Time `json:"startTime,omitempty" azure:"ro"`

	// READ-ONLY; The status of the action.
	Status *string `json:"status,omitempty" azure:"ro"`

	// READ-ONLY; The array of targets.
	Targets []*ExperimentExecutionActionTargetDetailsProperties `json:"targets,omitempty" azure:"ro"`
}

// Branch - Model that represents a branch in the step.
type Branch struct {
	// REQUIRED; List of actions.
	Actions []ActionClassification `json:"actions,omitempty"`

	// REQUIRED; String of the branch name.
	Name *string `json:"name,omitempty"`
}

// BranchStatus - Model that represents the a list of actions and action statuses.
type BranchStatus struct {
	// READ-ONLY; The array of actions.
	Actions []*ActionStatus `json:"actions,omitempty" azure:"ro"`

	// READ-ONLY; The id of the branch status.
	BranchID *string `json:"branchId,omitempty" azure:"ro"`

	// READ-ONLY; The name of the branch status.
	BranchName *string `json:"branchName,omitempty" azure:"ro"`

	// READ-ONLY; The status of the branch.
	Status *string `json:"status,omitempty" azure:"ro"`
}

// CapabilitiesClientCreateOrUpdateOptions contains the optional parameters for the CapabilitiesClient.CreateOrUpdate method.
type CapabilitiesClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// CapabilitiesClientDeleteOptions contains the optional parameters for the CapabilitiesClient.Delete method.
type CapabilitiesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// CapabilitiesClientGetOptions contains the optional parameters for the CapabilitiesClient.Get method.
type CapabilitiesClientGetOptions struct {
	// placeholder for future optional parameters
}

// CapabilitiesClientListOptions contains the optional parameters for the CapabilitiesClient.List method.
type CapabilitiesClientListOptions struct {
	// String that sets the continuation token.
	ContinuationToken *string
}

// Capability - Model that represents a Capability resource.
type Capability struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The properties of a capability resource.
	Properties *CapabilityProperties `json:"properties,omitempty" azure:"ro"`

	// READ-ONLY; The standard system metadata of a resource type.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// CapabilityListResult - Model that represents a list of Capability resources and a link for pagination.
type CapabilityListResult struct {
	// READ-ONLY; URL to retrieve the next page of Capability resources.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of Capability resources.
	Value []*Capability `json:"value,omitempty" azure:"ro"`
}

// CapabilityProperties - Model that represents the Capability properties model.
type CapabilityProperties struct {
	// READ-ONLY; Localized string of the description.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; URL to retrieve JSON schema of the Capability parameters.
	ParametersSchema *string `json:"parametersSchema,omitempty" azure:"ro"`

	// READ-ONLY; String of the Publisher that this Capability extends.
	Publisher *string `json:"publisher,omitempty" azure:"ro"`

	// READ-ONLY; String of the Target Type that this Capability extends.
	TargetType *string `json:"targetType,omitempty" azure:"ro"`

	// READ-ONLY; String of the URN for this Capability Type.
	Urn *string `json:"urn,omitempty" azure:"ro"`
}

// CapabilityType - Model that represents a Capability Type resource.
type CapabilityType struct {
	// Location of the Capability Type resource.
	Location *string `json:"location,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The properties of the capability type resource.
	Properties *CapabilityTypeProperties `json:"properties,omitempty" azure:"ro"`

	// READ-ONLY; The system metadata properties of the capability type resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// CapabilityTypeListResult - Model that represents a list of Capability Type resources and a link for pagination.
type CapabilityTypeListResult struct {
	// READ-ONLY; URL to retrieve the next page of Capability Type resources.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of Capability Type resources.
	Value []*CapabilityType `json:"value,omitempty" azure:"ro"`
}

// CapabilityTypeProperties - Model that represents the Capability Type properties model.
type CapabilityTypeProperties struct {
	// READ-ONLY; Localized string of the description.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; Localized string of the display name.
	DisplayName *string `json:"displayName,omitempty" azure:"ro"`

	// READ-ONLY; URL to retrieve JSON schema of the Capability Type parameters.
	ParametersSchema *string `json:"parametersSchema,omitempty" azure:"ro"`

	// READ-ONLY; String of the Publisher that this Capability Type extends.
	Publisher *string `json:"publisher,omitempty" azure:"ro"`

	// READ-ONLY; String of the Target Type that this Capability Type extends.
	TargetType *string `json:"targetType,omitempty" azure:"ro"`

	// READ-ONLY; String of the URN for this Capability Type.
	Urn *string `json:"urn,omitempty" azure:"ro"`
}

// CapabilityTypesClientGetOptions contains the optional parameters for the CapabilityTypesClient.Get method.
type CapabilityTypesClientGetOptions struct {
	// placeholder for future optional parameters
}

// CapabilityTypesClientListOptions contains the optional parameters for the CapabilityTypesClient.List method.
type CapabilityTypesClientListOptions struct {
	// String that sets the continuation token.
	ContinuationToken *string
}

// ContinuousAction - Model that represents a continuous action.
type ContinuousAction struct {
	// REQUIRED; ISO8601 formatted string that represents a duration.
	Duration *string `json:"duration,omitempty"`

	// REQUIRED; String that represents a Capability URN.
	Name *string `json:"name,omitempty"`

	// REQUIRED; List of key value pairs.
	Parameters []*KeyValuePair `json:"parameters,omitempty"`

	// REQUIRED; String that represents a selector.
	SelectorID *string `json:"selectorId,omitempty"`

	// REQUIRED; Enum that discriminates between action models.
	Type *string `json:"type,omitempty"`
}

// GetAction implements the ActionClassification interface for type ContinuousAction.
func (c *ContinuousAction) GetAction() *Action {
	return &Action{
		Type: c.Type,
		Name: c.Name,
	}
}

// DelayAction - Model that represents a delay action.
type DelayAction struct {
	// REQUIRED; ISO8601 formatted string that represents a duration.
	Duration *string `json:"duration,omitempty"`

	// REQUIRED; String that represents a Capability URN.
	Name *string `json:"name,omitempty"`

	// REQUIRED; Enum that discriminates between action models.
	Type *string `json:"type,omitempty"`
}

// GetAction implements the ActionClassification interface for type DelayAction.
func (d *DelayAction) GetAction() *Action {
	return &Action{
		Type: d.Type,
		Name: d.Name,
	}
}

// DiscreteAction - Model that represents a discrete action.
type DiscreteAction struct {
	// REQUIRED; String that represents a Capability URN.
	Name *string `json:"name,omitempty"`

	// REQUIRED; List of key value pairs.
	Parameters []*KeyValuePair `json:"parameters,omitempty"`

	// REQUIRED; String that represents a selector.
	SelectorID *string `json:"selectorId,omitempty"`

	// REQUIRED; Enum that discriminates between action models.
	Type *string `json:"type,omitempty"`
}

// GetAction implements the ActionClassification interface for type DiscreteAction.
func (d *DiscreteAction) GetAction() *Action {
	return &Action{
		Type: d.Type,
		Name: d.Name,
	}
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info interface{} `json:"info,omitempty" azure:"ro"`

	// READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo `json:"additionalInfo,omitempty" azure:"ro"`

	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error details.
	Details []*ErrorDetail `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The error target.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail `json:"error,omitempty"`
}

// Experiment - Model that represents a Experiment resource.
type Experiment struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// REQUIRED; The properties of the experiment resource.
	Properties *ExperimentProperties `json:"properties,omitempty"`

	// The identity of the experiment resource.
	Identity *ResourceIdentity `json:"identity,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The system metadata of the experiment resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ExperimentCancelOperationResult - Model that represents the result of a cancel Experiment operation.
type ExperimentCancelOperationResult struct {
	// READ-ONLY; String of the Experiment name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; URL to retrieve the Experiment status.
	StatusURL *string `json:"statusUrl,omitempty" azure:"ro"`
}

// ExperimentExecutionActionTargetDetailsError - Model that represents the Experiment action target details error model.
type ExperimentExecutionActionTargetDetailsError struct {
	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error message
	Message *string `json:"message,omitempty" azure:"ro"`
}

// ExperimentExecutionActionTargetDetailsProperties - Model that represents the Experiment action target details properties
// model.
type ExperimentExecutionActionTargetDetailsProperties struct {
	// READ-ONLY; The error of the action.
	Error *ExperimentExecutionActionTargetDetailsError `json:"error,omitempty" azure:"ro"`

	// READ-ONLY; The status of the execution.
	Status *string `json:"status,omitempty" azure:"ro"`

	// READ-ONLY; The target for the action.
	Target *string `json:"target,omitempty" azure:"ro"`

	// READ-ONLY; String that represents the completed date time.
	TargetCompletedTime *time.Time `json:"targetCompletedTime,omitempty" azure:"ro"`

	// READ-ONLY; String that represents the failed date time.
	TargetFailedTime *time.Time `json:"targetFailedTime,omitempty" azure:"ro"`
}

// ExperimentExecutionDetails - Model that represents the execution details of a Experiment.
type ExperimentExecutionDetails struct {
	// READ-ONLY; String of the fully qualified resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; String of the resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The properties of the experiment execution details.
	Properties *ExperimentExecutionDetailsProperties `json:"properties,omitempty" azure:"ro"`

	// READ-ONLY; String of the resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ExperimentExecutionDetailsListResult - Model that represents a list of Experiment execution details and a link for pagination.
type ExperimentExecutionDetailsListResult struct {
	// READ-ONLY; URL to retrieve the next page of Experiment execution details.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of Experiment execution details.
	Value []*ExperimentExecutionDetails `json:"value,omitempty" azure:"ro"`
}

// ExperimentExecutionDetailsProperties - Model that represents the Experiment execution details properties model.
type ExperimentExecutionDetailsProperties struct {
	// READ-ONLY; String that represents the created date time.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty" azure:"ro"`

	// READ-ONLY; The id of the experiment.
	ExperimentID *string `json:"experimentId,omitempty" azure:"ro"`

	// READ-ONLY; The reason why the execution failed.
	FailureReason *string `json:"failureReason,omitempty" azure:"ro"`

	// READ-ONLY; String that represents the last action date time.
	LastActionDateTime *time.Time `json:"lastActionDateTime,omitempty" azure:"ro"`

	// READ-ONLY; The information of the experiment run.
	RunInformation *ExperimentExecutionDetailsPropertiesRunInformation `json:"runInformation,omitempty" azure:"ro"`

	// READ-ONLY; String that represents the start date time.
	StartDateTime *time.Time `json:"startDateTime,omitempty" azure:"ro"`

	// READ-ONLY; The value of the status of the experiment execution.
	Status *string `json:"status,omitempty" azure:"ro"`

	// READ-ONLY; String that represents the stop date time.
	StopDateTime *time.Time `json:"stopDateTime,omitempty" azure:"ro"`
}

// ExperimentExecutionDetailsPropertiesRunInformation - The information of the experiment run.
type ExperimentExecutionDetailsPropertiesRunInformation struct {
	// READ-ONLY; The steps of the experiment run.
	Steps []*StepStatus `json:"steps,omitempty" azure:"ro"`
}

// ExperimentListResult - Model that represents a list of Experiment resources and a link for pagination.
type ExperimentListResult struct {
	// READ-ONLY; URL to retrieve the next page of Experiment resources.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of Experiment resources.
	Value []*Experiment `json:"value,omitempty" azure:"ro"`
}

// ExperimentProperties - Model that represents the Experiment properties model.
type ExperimentProperties struct {
	// REQUIRED; List of selectors.
	Selectors []*Selector `json:"selectors,omitempty"`

	// REQUIRED; List of steps.
	Steps []*Step `json:"steps,omitempty"`

	// A boolean value that indicates if experiment should be started on creation or not.
	StartOnCreation *bool `json:"startOnCreation,omitempty"`
}

// ExperimentStartOperationResult - Model that represents the result of a start Experiment operation.
type ExperimentStartOperationResult struct {
	// READ-ONLY; String of the Experiment name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; URL to retrieve the Experiment status.
	StatusURL *string `json:"statusUrl,omitempty" azure:"ro"`
}

// ExperimentStatus - Model that represents the status of a Experiment.
type ExperimentStatus struct {
	// The properties of experiment execution status.
	Properties *ExperimentStatusProperties `json:"properties,omitempty"`

	// READ-ONLY; String of the fully qualified resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; String of the resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; String of the resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ExperimentStatusListResult - Model that represents a list of Experiment statuses and a link for pagination.
type ExperimentStatusListResult struct {
	// READ-ONLY; URL to retrieve the next page of Experiment statuses.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of Experiment statuses.
	Value []*ExperimentStatus `json:"value,omitempty" azure:"ro"`
}

// ExperimentStatusProperties - Model that represents the Experiment status properties model.
type ExperimentStatusProperties struct {
	// READ-ONLY; String that represents the created date time of a Experiment.
	CreatedDateUTC *time.Time `json:"createdDateUtc,omitempty" azure:"ro"`

	// READ-ONLY; String that represents the end date time of a Experiment.
	EndDateUTC *time.Time `json:"endDateUtc,omitempty" azure:"ro"`

	// READ-ONLY; String that represents the status of a Experiment.
	Status *string `json:"status,omitempty" azure:"ro"`
}

// ExperimentsClientBeginCancelOptions contains the optional parameters for the ExperimentsClient.BeginCancel method.
type ExperimentsClientBeginCancelOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ExperimentsClientBeginCreateOrUpdateOptions contains the optional parameters for the ExperimentsClient.BeginCreateOrUpdate
// method.
type ExperimentsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ExperimentsClientDeleteOptions contains the optional parameters for the ExperimentsClient.Delete method.
type ExperimentsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ExperimentsClientGetExecutionDetailsOptions contains the optional parameters for the ExperimentsClient.GetExecutionDetails
// method.
type ExperimentsClientGetExecutionDetailsOptions struct {
	// placeholder for future optional parameters
}

// ExperimentsClientGetOptions contains the optional parameters for the ExperimentsClient.Get method.
type ExperimentsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ExperimentsClientGetStatusOptions contains the optional parameters for the ExperimentsClient.GetStatus method.
type ExperimentsClientGetStatusOptions struct {
	// placeholder for future optional parameters
}

// ExperimentsClientListAllOptions contains the optional parameters for the ExperimentsClient.ListAll method.
type ExperimentsClientListAllOptions struct {
	// String that sets the continuation token.
	ContinuationToken *string
	// Optional value that indicates whether to filter results based on if the Experiment is currently running. If null, then
	// the results will not be filtered.
	Running *bool
}

// ExperimentsClientListAllStatusesOptions contains the optional parameters for the ExperimentsClient.ListAllStatuses method.
type ExperimentsClientListAllStatusesOptions struct {
	// placeholder for future optional parameters
}

// ExperimentsClientListExecutionDetailsOptions contains the optional parameters for the ExperimentsClient.ListExecutionDetails
// method.
type ExperimentsClientListExecutionDetailsOptions struct {
	// placeholder for future optional parameters
}

// ExperimentsClientListOptions contains the optional parameters for the ExperimentsClient.List method.
type ExperimentsClientListOptions struct {
	// String that sets the continuation token.
	ContinuationToken *string
	// Optional value that indicates whether to filter results based on if the Experiment is currently running. If null, then
	// the results will not be filtered.
	Running *bool
}

// ExperimentsClientStartOptions contains the optional parameters for the ExperimentsClient.Start method.
type ExperimentsClientStartOptions struct {
	// placeholder for future optional parameters
}

// KeyValuePair - A map to describe the settings of an action.
type KeyValuePair struct {
	// REQUIRED; The name of the setting for the action.
	Key *string `json:"key,omitempty"`

	// REQUIRED; The value of the setting for the action.
	Value *string `json:"value,omitempty"`
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType `json:"actionType,omitempty" azure:"ro"`

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool `json:"isDataAction,omitempty" azure:"ro"`

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin `json:"origin,omitempty" azure:"ro"`
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// OperationsClientListAllOptions contains the optional parameters for the OperationsClient.ListAll method.
type OperationsClientListAllOptions struct {
	// placeholder for future optional parameters
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ResourceIdentity - The managed identity of a resource.
type ResourceIdentity struct {
	// REQUIRED; String of the resource identity type.
	Type *ResourceIdentityType `json:"type,omitempty"`

	// READ-ONLY; GUID that represents the principal ID of this resource identity.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`

	// READ-ONLY; GUID that represents the tenant ID of this resource identity.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// Selector - Model that represents a selector in the Experiment resource.
type Selector struct {
	// REQUIRED; String of the selector ID.
	ID *string `json:"id,omitempty"`

	// REQUIRED; List of Target references.
	Targets []*TargetReference `json:"targets,omitempty"`

	// REQUIRED; Enum of the selector type.
	Type *SelectorType `json:"type,omitempty"`
}

// Step - Model that represents a step in the Experiment resource.
type Step struct {
	// REQUIRED; List of branches.
	Branches []*Branch `json:"branches,omitempty"`

	// REQUIRED; String of the step name.
	Name *string `json:"name,omitempty"`
}

// StepStatus - Model that represents the a list of branches and branch statuses.
type StepStatus struct {
	// READ-ONLY; The array of branches.
	Branches []*BranchStatus `json:"branches,omitempty" azure:"ro"`

	// READ-ONLY; The value of the status of the step.
	Status *string `json:"status,omitempty" azure:"ro"`

	// READ-ONLY; The id of the step.
	StepID *string `json:"stepId,omitempty" azure:"ro"`

	// READ-ONLY; The name of the step.
	StepName *string `json:"stepName,omitempty" azure:"ro"`
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// Target - Model that represents a Target resource.
type Target struct {
	// REQUIRED; The properties of the target resource.
	Properties map[string]interface{} `json:"properties,omitempty"`

	// Location of the target resource.
	Location *string `json:"location,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The system metadata of the target resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// TargetListResult - Model that represents a list of Target resources and a link for pagination.
type TargetListResult struct {
	// READ-ONLY; URL to retrieve the next page of Target resources.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of Target resources.
	Value []*Target `json:"value,omitempty" azure:"ro"`
}

// TargetReference - Model that represents a reference to a Target in the selector.
type TargetReference struct {
	// REQUIRED; String of the resource ID of a Target resource.
	ID *string `json:"id,omitempty"`

	// REQUIRED; Enum of the Target reference type.
	Type *string `json:"type,omitempty"`
}

// TargetType - Model that represents a Target Type resource.
type TargetType struct {
	// REQUIRED; The properties of the target type resource.
	Properties *TargetTypeProperties `json:"properties,omitempty"`

	// Location of the Target Type resource.
	Location *string `json:"location,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The system metadata properties of the target type resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// TargetTypeListResult - Model that represents a list of Target Type resources and a link for pagination.
type TargetTypeListResult struct {
	// READ-ONLY; URL to retrieve the next page of Target Type resources.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of Target Type resources.
	Value []*TargetType `json:"value,omitempty" azure:"ro"`
}

// TargetTypeProperties - Model that represents the base Target Type properties model.
type TargetTypeProperties struct {
	// READ-ONLY; Localized string of the description.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; Localized string of the display name.
	DisplayName *string `json:"displayName,omitempty" azure:"ro"`

	// READ-ONLY; URL to retrieve JSON schema of the Target Type properties.
	PropertiesSchema *string `json:"propertiesSchema,omitempty" azure:"ro"`

	// READ-ONLY; List of resource types this Target Type can extend.
	ResourceTypes []*string `json:"resourceTypes,omitempty" azure:"ro"`
}

// TargetTypesClientGetOptions contains the optional parameters for the TargetTypesClient.Get method.
type TargetTypesClientGetOptions struct {
	// placeholder for future optional parameters
}

// TargetTypesClientListOptions contains the optional parameters for the TargetTypesClient.List method.
type TargetTypesClientListOptions struct {
	// String that sets the continuation token.
	ContinuationToken *string
}

// TargetsClientCreateOrUpdateOptions contains the optional parameters for the TargetsClient.CreateOrUpdate method.
type TargetsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// TargetsClientDeleteOptions contains the optional parameters for the TargetsClient.Delete method.
type TargetsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// TargetsClientGetOptions contains the optional parameters for the TargetsClient.Get method.
type TargetsClientGetOptions struct {
	// placeholder for future optional parameters
}

// TargetsClientListOptions contains the optional parameters for the TargetsClient.List method.
type TargetsClientListOptions struct {
	// String that sets the continuation token.
	ContinuationToken *string
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}
