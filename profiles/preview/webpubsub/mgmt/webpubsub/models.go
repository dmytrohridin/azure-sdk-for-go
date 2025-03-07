//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package webpubsub

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/webpubsub/mgmt/2021-10-01/webpubsub"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ACLAction = original.ACLAction

const (
	ACLActionAllow ACLAction = original.ACLActionAllow
	ACLActionDeny  ACLAction = original.ACLActionDeny
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type KeyType = original.KeyType

const (
	KeyTypePrimary   KeyType = original.KeyTypePrimary
	KeyTypeSalt      KeyType = original.KeyTypeSalt
	KeyTypeSecondary KeyType = original.KeyTypeSecondary
)

type ManagedIdentityType = original.ManagedIdentityType

const (
	ManagedIdentityTypeNone           ManagedIdentityType = original.ManagedIdentityTypeNone
	ManagedIdentityTypeSystemAssigned ManagedIdentityType = original.ManagedIdentityTypeSystemAssigned
	ManagedIdentityTypeUserAssigned   ManagedIdentityType = original.ManagedIdentityTypeUserAssigned
)

type PrivateLinkServiceConnectionStatus = original.PrivateLinkServiceConnectionStatus

const (
	PrivateLinkServiceConnectionStatusApproved     PrivateLinkServiceConnectionStatus = original.PrivateLinkServiceConnectionStatusApproved
	PrivateLinkServiceConnectionStatusDisconnected PrivateLinkServiceConnectionStatus = original.PrivateLinkServiceConnectionStatusDisconnected
	PrivateLinkServiceConnectionStatusPending      PrivateLinkServiceConnectionStatus = original.PrivateLinkServiceConnectionStatusPending
	PrivateLinkServiceConnectionStatusRejected     PrivateLinkServiceConnectionStatus = original.PrivateLinkServiceConnectionStatusRejected
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateCanceled  ProvisioningState = original.ProvisioningStateCanceled
	ProvisioningStateCreating  ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleting  ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed    ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateMoving    ProvisioningState = original.ProvisioningStateMoving
	ProvisioningStateRunning   ProvisioningState = original.ProvisioningStateRunning
	ProvisioningStateSucceeded ProvisioningState = original.ProvisioningStateSucceeded
	ProvisioningStateUnknown   ProvisioningState = original.ProvisioningStateUnknown
	ProvisioningStateUpdating  ProvisioningState = original.ProvisioningStateUpdating
)

type RequestType = original.RequestType

const (
	RequestTypeClientConnection RequestType = original.RequestTypeClientConnection
	RequestTypeRESTAPI          RequestType = original.RequestTypeRESTAPI
	RequestTypeServerConnection RequestType = original.RequestTypeServerConnection
	RequestTypeTrace            RequestType = original.RequestTypeTrace
)

type ScaleType = original.ScaleType

const (
	ScaleTypeAutomatic ScaleType = original.ScaleTypeAutomatic
	ScaleTypeManual    ScaleType = original.ScaleTypeManual
	ScaleTypeNone      ScaleType = original.ScaleTypeNone
)

type SharedPrivateLinkResourceStatus = original.SharedPrivateLinkResourceStatus

const (
	SharedPrivateLinkResourceStatusApproved     SharedPrivateLinkResourceStatus = original.SharedPrivateLinkResourceStatusApproved
	SharedPrivateLinkResourceStatusDisconnected SharedPrivateLinkResourceStatus = original.SharedPrivateLinkResourceStatusDisconnected
	SharedPrivateLinkResourceStatusPending      SharedPrivateLinkResourceStatus = original.SharedPrivateLinkResourceStatusPending
	SharedPrivateLinkResourceStatusRejected     SharedPrivateLinkResourceStatus = original.SharedPrivateLinkResourceStatusRejected
	SharedPrivateLinkResourceStatusTimeout      SharedPrivateLinkResourceStatus = original.SharedPrivateLinkResourceStatusTimeout
)

type SkuTier = original.SkuTier

const (
	SkuTierBasic    SkuTier = original.SkuTierBasic
	SkuTierFree     SkuTier = original.SkuTierFree
	SkuTierPremium  SkuTier = original.SkuTierPremium
	SkuTierStandard SkuTier = original.SkuTierStandard
)

type UpstreamAuthType = original.UpstreamAuthType

const (
	UpstreamAuthTypeManagedIdentity UpstreamAuthType = original.UpstreamAuthTypeManagedIdentity
	UpstreamAuthTypeNone            UpstreamAuthType = original.UpstreamAuthTypeNone
)

type BaseClient = original.BaseClient
type Client = original.Client
type CreateOrUpdateFuture = original.CreateOrUpdateFuture
type DeleteFuture = original.DeleteFuture
type Dimension = original.Dimension
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type EventHandler = original.EventHandler
type Hub = original.Hub
type HubList = original.HubList
type HubListIterator = original.HubListIterator
type HubListPage = original.HubListPage
type HubProperties = original.HubProperties
type HubsClient = original.HubsClient
type HubsCreateOrUpdateFuture = original.HubsCreateOrUpdateFuture
type HubsDeleteFuture = original.HubsDeleteFuture
type Keys = original.Keys
type LiveTraceCategory = original.LiveTraceCategory
type LiveTraceConfiguration = original.LiveTraceConfiguration
type LogSpecification = original.LogSpecification
type ManagedIdentity = original.ManagedIdentity
type ManagedIdentitySettings = original.ManagedIdentitySettings
type MetricSpecification = original.MetricSpecification
type NameAvailability = original.NameAvailability
type NameAvailabilityParameters = original.NameAvailabilityParameters
type NetworkACL = original.NetworkACL
type NetworkACLs = original.NetworkACLs
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OperationProperties = original.OperationProperties
type OperationsClient = original.OperationsClient
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointACL = original.PrivateEndpointACL
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionList = original.PrivateEndpointConnectionList
type PrivateEndpointConnectionListIterator = original.PrivateEndpointConnectionListIterator
type PrivateEndpointConnectionListPage = original.PrivateEndpointConnectionListPage
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateEndpointConnectionsDeleteFuture = original.PrivateEndpointConnectionsDeleteFuture
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceList = original.PrivateLinkResourceList
type PrivateLinkResourceListIterator = original.PrivateLinkResourceListIterator
type PrivateLinkResourceListPage = original.PrivateLinkResourceListPage
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type Properties = original.Properties
type ProxyResource = original.ProxyResource
type RegenerateKeyFuture = original.RegenerateKeyFuture
type RegenerateKeyParameters = original.RegenerateKeyParameters
type Resource = original.Resource
type ResourceList = original.ResourceList
type ResourceListIterator = original.ResourceListIterator
type ResourceListPage = original.ResourceListPage
type ResourceLogCategory = original.ResourceLogCategory
type ResourceLogConfiguration = original.ResourceLogConfiguration
type ResourceSku = original.ResourceSku
type ResourceType = original.ResourceType
type RestartFuture = original.RestartFuture
type ServiceSpecification = original.ServiceSpecification
type ShareablePrivateLinkResourceProperties = original.ShareablePrivateLinkResourceProperties
type ShareablePrivateLinkResourceType = original.ShareablePrivateLinkResourceType
type SharedPrivateLinkResource = original.SharedPrivateLinkResource
type SharedPrivateLinkResourceList = original.SharedPrivateLinkResourceList
type SharedPrivateLinkResourceListIterator = original.SharedPrivateLinkResourceListIterator
type SharedPrivateLinkResourceListPage = original.SharedPrivateLinkResourceListPage
type SharedPrivateLinkResourceProperties = original.SharedPrivateLinkResourceProperties
type SharedPrivateLinkResourcesClient = original.SharedPrivateLinkResourcesClient
type SharedPrivateLinkResourcesCreateOrUpdateFuture = original.SharedPrivateLinkResourcesCreateOrUpdateFuture
type SharedPrivateLinkResourcesDeleteFuture = original.SharedPrivateLinkResourcesDeleteFuture
type SignalRServiceUsage = original.SignalRServiceUsage
type SignalRServiceUsageList = original.SignalRServiceUsageList
type SignalRServiceUsageListIterator = original.SignalRServiceUsageListIterator
type SignalRServiceUsageListPage = original.SignalRServiceUsageListPage
type SignalRServiceUsageName = original.SignalRServiceUsageName
type Sku = original.Sku
type SkuCapacity = original.SkuCapacity
type SkuList = original.SkuList
type SystemData = original.SystemData
type TLSSettings = original.TLSSettings
type TrackedResource = original.TrackedResource
type UpdateFuture = original.UpdateFuture
type UpstreamAuthSettings = original.UpstreamAuthSettings
type UsagesClient = original.UsagesClient
type UserAssignedIdentityProperty = original.UserAssignedIdentityProperty

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewClient(subscriptionID string) Client {
	return original.NewClient(subscriptionID)
}
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}
func NewHubListIterator(page HubListPage) HubListIterator {
	return original.NewHubListIterator(page)
}
func NewHubListPage(cur HubList, getNextPage func(context.Context, HubList) (HubList, error)) HubListPage {
	return original.NewHubListPage(cur, getNextPage)
}
func NewHubsClient(subscriptionID string) HubsClient {
	return original.NewHubsClient(subscriptionID)
}
func NewHubsClientWithBaseURI(baseURI string, subscriptionID string) HubsClient {
	return original.NewHubsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return original.NewOperationListIterator(page)
}
func NewOperationListPage(cur OperationList, getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return original.NewOperationListPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionListIterator(page PrivateEndpointConnectionListPage) PrivateEndpointConnectionListIterator {
	return original.NewPrivateEndpointConnectionListIterator(page)
}
func NewPrivateEndpointConnectionListPage(cur PrivateEndpointConnectionList, getNextPage func(context.Context, PrivateEndpointConnectionList) (PrivateEndpointConnectionList, error)) PrivateEndpointConnectionListPage {
	return original.NewPrivateEndpointConnectionListPage(cur, getNextPage)
}
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClient(subscriptionID)
}
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourceListIterator(page PrivateLinkResourceListPage) PrivateLinkResourceListIterator {
	return original.NewPrivateLinkResourceListIterator(page)
}
func NewPrivateLinkResourceListPage(cur PrivateLinkResourceList, getNextPage func(context.Context, PrivateLinkResourceList) (PrivateLinkResourceList, error)) PrivateLinkResourceListPage {
	return original.NewPrivateLinkResourceListPage(cur, getNextPage)
}
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClient(subscriptionID)
}
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceListIterator(page ResourceListPage) ResourceListIterator {
	return original.NewResourceListIterator(page)
}
func NewResourceListPage(cur ResourceList, getNextPage func(context.Context, ResourceList) (ResourceList, error)) ResourceListPage {
	return original.NewResourceListPage(cur, getNextPage)
}
func NewSharedPrivateLinkResourceListIterator(page SharedPrivateLinkResourceListPage) SharedPrivateLinkResourceListIterator {
	return original.NewSharedPrivateLinkResourceListIterator(page)
}
func NewSharedPrivateLinkResourceListPage(cur SharedPrivateLinkResourceList, getNextPage func(context.Context, SharedPrivateLinkResourceList) (SharedPrivateLinkResourceList, error)) SharedPrivateLinkResourceListPage {
	return original.NewSharedPrivateLinkResourceListPage(cur, getNextPage)
}
func NewSharedPrivateLinkResourcesClient(subscriptionID string) SharedPrivateLinkResourcesClient {
	return original.NewSharedPrivateLinkResourcesClient(subscriptionID)
}
func NewSharedPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) SharedPrivateLinkResourcesClient {
	return original.NewSharedPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewSignalRServiceUsageListIterator(page SignalRServiceUsageListPage) SignalRServiceUsageListIterator {
	return original.NewSignalRServiceUsageListIterator(page)
}
func NewSignalRServiceUsageListPage(cur SignalRServiceUsageList, getNextPage func(context.Context, SignalRServiceUsageList) (SignalRServiceUsageList, error)) SignalRServiceUsageListPage {
	return original.NewSignalRServiceUsageListPage(cur, getNextPage)
}
func NewUsagesClient(subscriptionID string) UsagesClient {
	return original.NewUsagesClient(subscriptionID)
}
func NewUsagesClientWithBaseURI(baseURI string, subscriptionID string) UsagesClient {
	return original.NewUsagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleACLActionValues() []ACLAction {
	return original.PossibleACLActionValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleKeyTypeValues() []KeyType {
	return original.PossibleKeyTypeValues()
}
func PossibleManagedIdentityTypeValues() []ManagedIdentityType {
	return original.PossibleManagedIdentityTypeValues()
}
func PossiblePrivateLinkServiceConnectionStatusValues() []PrivateLinkServiceConnectionStatus {
	return original.PossiblePrivateLinkServiceConnectionStatusValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleRequestTypeValues() []RequestType {
	return original.PossibleRequestTypeValues()
}
func PossibleScaleTypeValues() []ScaleType {
	return original.PossibleScaleTypeValues()
}
func PossibleSharedPrivateLinkResourceStatusValues() []SharedPrivateLinkResourceStatus {
	return original.PossibleSharedPrivateLinkResourceStatusValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleUpstreamAuthTypeValues() []UpstreamAuthType {
	return original.PossibleUpstreamAuthTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
