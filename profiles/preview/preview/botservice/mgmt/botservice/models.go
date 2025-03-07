//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package botservice

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/botservice/mgmt/2021-05-01-preview/botservice"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ChannelName = original.ChannelName

const (
	ChannelNameAlexaChannel            ChannelName = original.ChannelNameAlexaChannel
	ChannelNameDirectLineChannel       ChannelName = original.ChannelNameDirectLineChannel
	ChannelNameDirectLineSpeechChannel ChannelName = original.ChannelNameDirectLineSpeechChannel
	ChannelNameEmailChannel            ChannelName = original.ChannelNameEmailChannel
	ChannelNameFacebookChannel         ChannelName = original.ChannelNameFacebookChannel
	ChannelNameKikChannel              ChannelName = original.ChannelNameKikChannel
	ChannelNameLineChannel             ChannelName = original.ChannelNameLineChannel
	ChannelNameMsTeamsChannel          ChannelName = original.ChannelNameMsTeamsChannel
	ChannelNameOutlookChannel          ChannelName = original.ChannelNameOutlookChannel
	ChannelNameSkypeChannel            ChannelName = original.ChannelNameSkypeChannel
	ChannelNameSlackChannel            ChannelName = original.ChannelNameSlackChannel
	ChannelNameSmsChannel              ChannelName = original.ChannelNameSmsChannel
	ChannelNameTelegramChannel         ChannelName = original.ChannelNameTelegramChannel
	ChannelNameWebChatChannel          ChannelName = original.ChannelNameWebChatChannel
)

type ChannelNameBasicChannel = original.ChannelNameBasicChannel

const (
	ChannelNameBasicChannelChannelNameAlexaChannel            ChannelNameBasicChannel = original.ChannelNameBasicChannelChannelNameAlexaChannel
	ChannelNameBasicChannelChannelNameChannel                 ChannelNameBasicChannel = original.ChannelNameBasicChannelChannelNameChannel
	ChannelNameBasicChannelChannelNameDirectLineChannel       ChannelNameBasicChannel = original.ChannelNameBasicChannelChannelNameDirectLineChannel
	ChannelNameBasicChannelChannelNameDirectLineSpeechChannel ChannelNameBasicChannel = original.ChannelNameBasicChannelChannelNameDirectLineSpeechChannel
	ChannelNameBasicChannelChannelNameEmailChannel            ChannelNameBasicChannel = original.ChannelNameBasicChannelChannelNameEmailChannel
	ChannelNameBasicChannelChannelNameFacebookChannel         ChannelNameBasicChannel = original.ChannelNameBasicChannelChannelNameFacebookChannel
	ChannelNameBasicChannelChannelNameKikChannel              ChannelNameBasicChannel = original.ChannelNameBasicChannelChannelNameKikChannel
	ChannelNameBasicChannelChannelNameLineChannel             ChannelNameBasicChannel = original.ChannelNameBasicChannelChannelNameLineChannel
	ChannelNameBasicChannelChannelNameMsTeamsChannel          ChannelNameBasicChannel = original.ChannelNameBasicChannelChannelNameMsTeamsChannel
	ChannelNameBasicChannelChannelNameSkypeChannel            ChannelNameBasicChannel = original.ChannelNameBasicChannelChannelNameSkypeChannel
	ChannelNameBasicChannelChannelNameSlackChannel            ChannelNameBasicChannel = original.ChannelNameBasicChannelChannelNameSlackChannel
	ChannelNameBasicChannelChannelNameSmsChannel              ChannelNameBasicChannel = original.ChannelNameBasicChannelChannelNameSmsChannel
	ChannelNameBasicChannelChannelNameTelegramChannel         ChannelNameBasicChannel = original.ChannelNameBasicChannelChannelNameTelegramChannel
	ChannelNameBasicChannelChannelNameWebChatChannel          ChannelNameBasicChannel = original.ChannelNameBasicChannelChannelNameWebChatChannel
)

type Key = original.Key

const (
	Key1 Key = original.Key1
	Key2 Key = original.Key2
)

type Kind = original.Kind

const (
	KindAzurebot Kind = original.KindAzurebot
	KindBot      Kind = original.KindBot
	KindDesigner Kind = original.KindDesigner
	KindFunction Kind = original.KindFunction
	KindSdk      Kind = original.KindSdk
)

type MsaAppType = original.MsaAppType

const (
	MsaAppTypeMultiTenant     MsaAppType = original.MsaAppTypeMultiTenant
	MsaAppTypeSingleTenant    MsaAppType = original.MsaAppTypeSingleTenant
	MsaAppTypeUserAssignedMSI MsaAppType = original.MsaAppTypeUserAssignedMSI
)

type OperationResultStatus = original.OperationResultStatus

const (
	OperationResultStatusCanceled  OperationResultStatus = original.OperationResultStatusCanceled
	OperationResultStatusFailed    OperationResultStatus = original.OperationResultStatusFailed
	OperationResultStatusRequested OperationResultStatus = original.OperationResultStatusRequested
	OperationResultStatusRunning   OperationResultStatus = original.OperationResultStatusRunning
	OperationResultStatusSucceeded OperationResultStatus = original.OperationResultStatusSucceeded
)

type PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningState

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateCreating
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateDeleting
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateFailed
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateSucceeded
)

type PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatus

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusApproved
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusPending
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusRejected
)

type PublicNetworkAccess = original.PublicNetworkAccess

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = original.PublicNetworkAccessDisabled
	PublicNetworkAccessEnabled  PublicNetworkAccess = original.PublicNetworkAccessEnabled
)

type RegenerateKeysChannelName = original.RegenerateKeysChannelName

const (
	RegenerateKeysChannelNameDirectLineChannel RegenerateKeysChannelName = original.RegenerateKeysChannelNameDirectLineChannel
	RegenerateKeysChannelNameWebChatChannel    RegenerateKeysChannelName = original.RegenerateKeysChannelNameWebChatChannel
)

type SkuName = original.SkuName

const (
	SkuNameF0 SkuName = original.SkuNameF0
	SkuNameS1 SkuName = original.SkuNameS1
)

type SkuTier = original.SkuTier

const (
	SkuTierFree     SkuTier = original.SkuTierFree
	SkuTierStandard SkuTier = original.SkuTierStandard
)

type AlexaChannel = original.AlexaChannel
type AlexaChannelProperties = original.AlexaChannelProperties
type BaseClient = original.BaseClient
type BasicChannel = original.BasicChannel
type Bot = original.Bot
type BotChannel = original.BotChannel
type BotConnectionClient = original.BotConnectionClient
type BotProperties = original.BotProperties
type BotResponseList = original.BotResponseList
type BotResponseListIterator = original.BotResponseListIterator
type BotResponseListPage = original.BotResponseListPage
type BotsClient = original.BotsClient
type Channel = original.Channel
type ChannelResponseList = original.ChannelResponseList
type ChannelResponseListIterator = original.ChannelResponseListIterator
type ChannelResponseListPage = original.ChannelResponseListPage
type ChannelSettings = original.ChannelSettings
type ChannelsClient = original.ChannelsClient
type CheckNameAvailabilityRequestBody = original.CheckNameAvailabilityRequestBody
type CheckNameAvailabilityResponseBody = original.CheckNameAvailabilityResponseBody
type ConnectionItemName = original.ConnectionItemName
type ConnectionSetting = original.ConnectionSetting
type ConnectionSettingParameter = original.ConnectionSettingParameter
type ConnectionSettingProperties = original.ConnectionSettingProperties
type ConnectionSettingResponseList = original.ConnectionSettingResponseList
type ConnectionSettingResponseListIterator = original.ConnectionSettingResponseListIterator
type ConnectionSettingResponseListPage = original.ConnectionSettingResponseListPage
type DirectLineChannel = original.DirectLineChannel
type DirectLineChannelProperties = original.DirectLineChannelProperties
type DirectLineClient = original.DirectLineClient
type DirectLineSite = original.DirectLineSite
type DirectLineSpeechChannel = original.DirectLineSpeechChannel
type DirectLineSpeechChannelProperties = original.DirectLineSpeechChannelProperties
type EmailChannel = original.EmailChannel
type EmailChannelProperties = original.EmailChannelProperties
type Error = original.Error
type ErrorBody = original.ErrorBody
type FacebookChannel = original.FacebookChannel
type FacebookChannelProperties = original.FacebookChannelProperties
type FacebookPage = original.FacebookPage
type HostSettingsClient = original.HostSettingsClient
type HostSettingsResponse = original.HostSettingsResponse
type KikChannel = original.KikChannel
type KikChannelProperties = original.KikChannelProperties
type LineChannel = original.LineChannel
type LineChannelProperties = original.LineChannelProperties
type LineRegistration = original.LineRegistration
type ListChannelWithKeysResponse = original.ListChannelWithKeysResponse
type MsTeamsChannel = original.MsTeamsChannel
type MsTeamsChannelProperties = original.MsTeamsChannelProperties
type OperationDisplayInfo = original.OperationDisplayInfo
type OperationEntity = original.OperationEntity
type OperationEntityListResult = original.OperationEntityListResult
type OperationEntityListResultIterator = original.OperationEntityListResultIterator
type OperationEntityListResultPage = original.OperationEntityListResultPage
type OperationResultsClient = original.OperationResultsClient
type OperationResultsDescription = original.OperationResultsDescription
type OperationResultsGetFuture = original.OperationResultsGetFuture
type OperationsClient = original.OperationsClient
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionListResult = original.PrivateEndpointConnectionListResult
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceBase = original.PrivateLinkResourceBase
type PrivateLinkResourceListResult = original.PrivateLinkResourceListResult
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type Resource = original.Resource
type ServiceProvider = original.ServiceProvider
type ServiceProviderParameter = original.ServiceProviderParameter
type ServiceProviderParameterMetadata = original.ServiceProviderParameterMetadata
type ServiceProviderParameterMetadataConstraints = original.ServiceProviderParameterMetadataConstraints
type ServiceProviderProperties = original.ServiceProviderProperties
type ServiceProviderResponseList = original.ServiceProviderResponseList
type Site = original.Site
type SiteInfo = original.SiteInfo
type Sku = original.Sku
type SkypeChannel = original.SkypeChannel
type SkypeChannelProperties = original.SkypeChannelProperties
type SlackChannel = original.SlackChannel
type SlackChannelProperties = original.SlackChannelProperties
type SmsChannel = original.SmsChannel
type SmsChannelProperties = original.SmsChannelProperties
type TelegramChannel = original.TelegramChannel
type TelegramChannelProperties = original.TelegramChannelProperties
type WebChatChannel = original.WebChatChannel
type WebChatChannelProperties = original.WebChatChannelProperties
type WebChatSite = original.WebChatSite

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewBotConnectionClient(subscriptionID string) BotConnectionClient {
	return original.NewBotConnectionClient(subscriptionID)
}
func NewBotConnectionClientWithBaseURI(baseURI string, subscriptionID string) BotConnectionClient {
	return original.NewBotConnectionClientWithBaseURI(baseURI, subscriptionID)
}
func NewBotResponseListIterator(page BotResponseListPage) BotResponseListIterator {
	return original.NewBotResponseListIterator(page)
}
func NewBotResponseListPage(cur BotResponseList, getNextPage func(context.Context, BotResponseList) (BotResponseList, error)) BotResponseListPage {
	return original.NewBotResponseListPage(cur, getNextPage)
}
func NewBotsClient(subscriptionID string) BotsClient {
	return original.NewBotsClient(subscriptionID)
}
func NewBotsClientWithBaseURI(baseURI string, subscriptionID string) BotsClient {
	return original.NewBotsClientWithBaseURI(baseURI, subscriptionID)
}
func NewChannelResponseListIterator(page ChannelResponseListPage) ChannelResponseListIterator {
	return original.NewChannelResponseListIterator(page)
}
func NewChannelResponseListPage(cur ChannelResponseList, getNextPage func(context.Context, ChannelResponseList) (ChannelResponseList, error)) ChannelResponseListPage {
	return original.NewChannelResponseListPage(cur, getNextPage)
}
func NewChannelsClient(subscriptionID string) ChannelsClient {
	return original.NewChannelsClient(subscriptionID)
}
func NewChannelsClientWithBaseURI(baseURI string, subscriptionID string) ChannelsClient {
	return original.NewChannelsClientWithBaseURI(baseURI, subscriptionID)
}
func NewConnectionSettingResponseListIterator(page ConnectionSettingResponseListPage) ConnectionSettingResponseListIterator {
	return original.NewConnectionSettingResponseListIterator(page)
}
func NewConnectionSettingResponseListPage(cur ConnectionSettingResponseList, getNextPage func(context.Context, ConnectionSettingResponseList) (ConnectionSettingResponseList, error)) ConnectionSettingResponseListPage {
	return original.NewConnectionSettingResponseListPage(cur, getNextPage)
}
func NewDirectLineClient(subscriptionID string) DirectLineClient {
	return original.NewDirectLineClient(subscriptionID)
}
func NewDirectLineClientWithBaseURI(baseURI string, subscriptionID string) DirectLineClient {
	return original.NewDirectLineClientWithBaseURI(baseURI, subscriptionID)
}
func NewHostSettingsClient(subscriptionID string) HostSettingsClient {
	return original.NewHostSettingsClient(subscriptionID)
}
func NewHostSettingsClientWithBaseURI(baseURI string, subscriptionID string) HostSettingsClient {
	return original.NewHostSettingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationEntityListResultIterator(page OperationEntityListResultPage) OperationEntityListResultIterator {
	return original.NewOperationEntityListResultIterator(page)
}
func NewOperationEntityListResultPage(cur OperationEntityListResult, getNextPage func(context.Context, OperationEntityListResult) (OperationEntityListResult, error)) OperationEntityListResultPage {
	return original.NewOperationEntityListResultPage(cur, getNextPage)
}
func NewOperationResultsClient(subscriptionID string) OperationResultsClient {
	return original.NewOperationResultsClient(subscriptionID)
}
func NewOperationResultsClientWithBaseURI(baseURI string, subscriptionID string) OperationResultsClient {
	return original.NewOperationResultsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClient(subscriptionID)
}
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClient(subscriptionID)
}
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleChannelNameBasicChannelValues() []ChannelNameBasicChannel {
	return original.PossibleChannelNameBasicChannelValues()
}
func PossibleChannelNameValues() []ChannelName {
	return original.PossibleChannelNameValues()
}
func PossibleKeyValues() []Key {
	return original.PossibleKeyValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleMsaAppTypeValues() []MsaAppType {
	return original.PossibleMsaAppTypeValues()
}
func PossibleOperationResultStatusValues() []OperationResultStatus {
	return original.PossibleOperationResultStatusValues()
}
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return original.PossiblePrivateEndpointConnectionProvisioningStateValues()
}
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return original.PossiblePrivateEndpointServiceConnectionStatusValues()
}
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return original.PossiblePublicNetworkAccessValues()
}
func PossibleRegenerateKeysChannelNameValues() []RegenerateKeysChannelName {
	return original.PossibleRegenerateKeysChannelNameValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
