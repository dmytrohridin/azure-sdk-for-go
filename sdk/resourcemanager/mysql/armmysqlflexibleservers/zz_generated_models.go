//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysqlflexibleservers

import "time"

// Backup - Storage Profile properties of a server
type Backup struct {
	// Backup retention days for the server.
	BackupRetentionDays *int32 `json:"backupRetentionDays,omitempty"`

	// Whether or not geo redundant backup is enabled.
	GeoRedundantBackup *EnableStatusEnum `json:"geoRedundantBackup,omitempty"`

	// READ-ONLY; Earliest restore point creation time (ISO8601 format)
	EarliestRestoreDate *time.Time `json:"earliestRestoreDate,omitempty" azure:"ro"`
}

// BackupsClientGetOptions contains the optional parameters for the BackupsClient.Get method.
type BackupsClientGetOptions struct {
	// placeholder for future optional parameters
}

// BackupsClientListByServerOptions contains the optional parameters for the BackupsClient.ListByServer method.
type BackupsClientListByServerOptions struct {
	// placeholder for future optional parameters
}

// CapabilitiesListResult - location capability
type CapabilitiesListResult struct {
	// READ-ONLY; Link to retrieve next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; A list of supported capabilities.
	Value []*CapabilityProperties `json:"value,omitempty" azure:"ro"`
}

// CapabilityProperties - Location capabilities.
type CapabilityProperties struct {
	// READ-ONLY; A list of supported flexible server editions.
	SupportedFlexibleServerEditions []*ServerEditionCapability `json:"supportedFlexibleServerEditions,omitempty" azure:"ro"`

	// READ-ONLY; supported geo backup regions
	SupportedGeoBackupRegions []*string `json:"supportedGeoBackupRegions,omitempty" azure:"ro"`

	// READ-ONLY; Supported high availability mode
	SupportedHAMode []*string `json:"supportedHAMode,omitempty" azure:"ro"`

	// READ-ONLY; zone name
	Zone *string `json:"zone,omitempty" azure:"ro"`
}

// CheckNameAvailabilityClientExecuteOptions contains the optional parameters for the CheckNameAvailabilityClient.Execute
// method.
type CheckNameAvailabilityClientExecuteOptions struct {
	// placeholder for future optional parameters
}

// CheckVirtualNetworkSubnetUsageClientExecuteOptions contains the optional parameters for the CheckVirtualNetworkSubnetUsageClient.Execute
// method.
type CheckVirtualNetworkSubnetUsageClientExecuteOptions struct {
	// placeholder for future optional parameters
}

// CloudError - An error response from the Batch service.
type CloudError struct {
	// The resource management error response.
	Error *ErrorResponse `json:"error,omitempty"`
}

// Configuration - Represents a Configuration.
type Configuration struct {
	// The properties of a configuration.
	Properties *ConfigurationProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The system metadata relating to this resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ConfigurationForBatchUpdate - Represents a Configuration.
type ConfigurationForBatchUpdate struct {
	// Name of the configuration.
	Name *string `json:"name,omitempty"`

	// The properties can be updated for a configuration.
	Properties *ConfigurationForBatchUpdateProperties `json:"properties,omitempty"`
}

// ConfigurationForBatchUpdateProperties - The properties can be updated for a configuration.
type ConfigurationForBatchUpdateProperties struct {
	// Source of the configuration.
	Source *string `json:"source,omitempty"`

	// Value of the configuration.
	Value *string `json:"value,omitempty"`
}

// ConfigurationListForBatchUpdate - A list of server configurations to update.
type ConfigurationListForBatchUpdate struct {
	// The list of server configurations.
	Value []*ConfigurationForBatchUpdate `json:"value,omitempty"`
}

// ConfigurationListResult - A list of server configurations.
type ConfigurationListResult struct {
	// The link used to get the next page of operations.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of server configurations.
	Value []*Configuration `json:"value,omitempty"`
}

// ConfigurationProperties - The properties of a configuration.
type ConfigurationProperties struct {
	// Source of the configuration.
	Source *ConfigurationSource `json:"source,omitempty"`

	// Value of the configuration.
	Value *string `json:"value,omitempty"`

	// READ-ONLY; Allowed values of the configuration.
	AllowedValues *string `json:"allowedValues,omitempty" azure:"ro"`

	// READ-ONLY; Data type of the configuration.
	DataType *string `json:"dataType,omitempty" azure:"ro"`

	// READ-ONLY; Default value of the configuration.
	DefaultValue *string `json:"defaultValue,omitempty" azure:"ro"`

	// READ-ONLY; Description of the configuration.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; If is the configuration pending restart or not.
	IsConfigPendingRestart *IsConfigPendingRestart `json:"isConfigPendingRestart,omitempty" azure:"ro"`

	// READ-ONLY; If is the configuration dynamic.
	IsDynamicConfig *IsDynamicConfig `json:"isDynamicConfig,omitempty" azure:"ro"`

	// READ-ONLY; If is the configuration read only.
	IsReadOnly *IsReadOnly `json:"isReadOnly,omitempty" azure:"ro"`
}

// ConfigurationsClientBeginBatchUpdateOptions contains the optional parameters for the ConfigurationsClient.BeginBatchUpdate
// method.
type ConfigurationsClientBeginBatchUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ConfigurationsClientBeginUpdateOptions contains the optional parameters for the ConfigurationsClient.BeginUpdate method.
type ConfigurationsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ConfigurationsClientGetOptions contains the optional parameters for the ConfigurationsClient.Get method.
type ConfigurationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationsClientListByServerOptions contains the optional parameters for the ConfigurationsClient.ListByServer method.
type ConfigurationsClientListByServerOptions struct {
	// placeholder for future optional parameters
}

// DataEncryption - The date encryption for cmk.
type DataEncryption struct {
	// Geo backup key uri as key vault can't cross region, need cmk in same region as geo backup
	GeoBackupKeyURI *string `json:"geoBackupKeyUri,omitempty"`

	// Geo backup user identity resource id as identity can't cross region, need identity in same region as geo backup
	GeoBackupUserAssignedIdentityID *string `json:"geoBackupUserAssignedIdentityId,omitempty"`

	// Primary key uri
	PrimaryKeyURI *string `json:"primaryKeyUri,omitempty"`

	// Primary user identity resource id
	PrimaryUserAssignedIdentityID *string `json:"primaryUserAssignedIdentityId,omitempty"`

	// The key type, AzureKeyVault for enable cmk, SystemManaged for disable cmk.
	Type *DataEncryptionType `json:"type,omitempty"`
}

// Database - Represents a Database.
type Database struct {
	// The properties of a database.
	Properties *DatabaseProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The system metadata relating to this resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// DatabaseListResult - A List of databases.
type DatabaseListResult struct {
	// The link used to get the next page of operations.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of databases housed in a server
	Value []*Database `json:"value,omitempty"`
}

// DatabaseProperties - The properties of a database.
type DatabaseProperties struct {
	// The charset of the database.
	Charset *string `json:"charset,omitempty"`

	// The collation of the database.
	Collation *string `json:"collation,omitempty"`
}

// DatabasesClientBeginCreateOrUpdateOptions contains the optional parameters for the DatabasesClient.BeginCreateOrUpdate
// method.
type DatabasesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DatabasesClientBeginDeleteOptions contains the optional parameters for the DatabasesClient.BeginDelete method.
type DatabasesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DatabasesClientGetOptions contains the optional parameters for the DatabasesClient.Get method.
type DatabasesClientGetOptions struct {
	// placeholder for future optional parameters
}

// DatabasesClientListByServerOptions contains the optional parameters for the DatabasesClient.ListByServer method.
type DatabasesClientListByServerOptions struct {
	// placeholder for future optional parameters
}

// DelegatedSubnetUsage - Delegated subnet usage data.
type DelegatedSubnetUsage struct {
	// READ-ONLY; name of the subnet
	SubnetName *string `json:"subnetName,omitempty" azure:"ro"`

	// READ-ONLY; Number of used delegated subnets
	Usage *int64 `json:"usage,omitempty" azure:"ro"`
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info interface{} `json:"info,omitempty" azure:"ro"`

	// READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.)
type ErrorResponse struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo `json:"additionalInfo,omitempty" azure:"ro"`

	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error details.
	Details []*ErrorResponse `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The error target.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// FirewallRule - Represents a server firewall rule.
type FirewallRule struct {
	// REQUIRED; The properties of a firewall rule.
	Properties *FirewallRuleProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The system metadata relating to this resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// FirewallRuleListResult - A list of firewall rules.
type FirewallRuleListResult struct {
	// The link used to get the next page of operations.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of firewall rules in a server.
	Value []*FirewallRule `json:"value,omitempty"`
}

// FirewallRuleProperties - The properties of a server firewall rule.
type FirewallRuleProperties struct {
	// REQUIRED; The end IP address of the server firewall rule. Must be IPv4 format.
	EndIPAddress *string `json:"endIpAddress,omitempty"`

	// REQUIRED; The start IP address of the server firewall rule. Must be IPv4 format.
	StartIPAddress *string `json:"startIpAddress,omitempty"`
}

// FirewallRulesClientBeginCreateOrUpdateOptions contains the optional parameters for the FirewallRulesClient.BeginCreateOrUpdate
// method.
type FirewallRulesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// FirewallRulesClientBeginDeleteOptions contains the optional parameters for the FirewallRulesClient.BeginDelete method.
type FirewallRulesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// FirewallRulesClientGetOptions contains the optional parameters for the FirewallRulesClient.Get method.
type FirewallRulesClientGetOptions struct {
	// placeholder for future optional parameters
}

// FirewallRulesClientListByServerOptions contains the optional parameters for the FirewallRulesClient.ListByServer method.
type FirewallRulesClientListByServerOptions struct {
	// placeholder for future optional parameters
}

// GetPrivateDNSZoneSuffixClientExecuteOptions contains the optional parameters for the GetPrivateDNSZoneSuffixClient.Execute
// method.
type GetPrivateDNSZoneSuffixClientExecuteOptions struct {
	// placeholder for future optional parameters
}

// GetPrivateDNSZoneSuffixResponse - The response of get private dns zone suffix.
type GetPrivateDNSZoneSuffixResponse struct {
	// Represents the private DNS zone suffix.
	PrivateDNSZoneSuffix *string `json:"privateDnsZoneSuffix,omitempty"`
}

// HighAvailability - Network related properties of a server
type HighAvailability struct {
	// High availability mode for a server.
	Mode *HighAvailabilityMode `json:"mode,omitempty"`

	// Availability zone of the standby server.
	StandbyAvailabilityZone *string `json:"standbyAvailabilityZone,omitempty"`

	// READ-ONLY; The state of server high availability.
	State *HighAvailabilityState `json:"state,omitempty" azure:"ro"`
}

// Identity - Properties to configure Identity for Bring your Own Keys
type Identity struct {
	// Type of managed service identity.
	Type *string `json:"type,omitempty"`

	// Metadata of user assigned identity.
	UserAssignedIdentities map[string]interface{} `json:"userAssignedIdentities,omitempty"`

	// READ-ONLY; ObjectId from the KeyVault
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`

	// READ-ONLY; TenantId from the KeyVault
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// LocationBasedCapabilitiesClientListOptions contains the optional parameters for the LocationBasedCapabilitiesClient.List
// method.
type LocationBasedCapabilitiesClientListOptions struct {
	// placeholder for future optional parameters
}

// MaintenanceWindow - Maintenance window of a server.
type MaintenanceWindow struct {
	// indicates whether custom window is enabled or disabled
	CustomWindow *string `json:"customWindow,omitempty"`

	// day of week for maintenance window
	DayOfWeek *int32 `json:"dayOfWeek,omitempty"`

	// start hour for maintenance window
	StartHour *int32 `json:"startHour,omitempty"`

	// start minute for maintenance window
	StartMinute *int32 `json:"startMinute,omitempty"`
}

// NameAvailability - Represents a resource name availability.
type NameAvailability struct {
	// Error Message.
	Message *string `json:"message,omitempty"`

	// Indicates whether the resource name is available.
	NameAvailable *bool `json:"nameAvailable,omitempty"`

	// Reason for name being unavailable.
	Reason *string `json:"reason,omitempty"`
}

// NameAvailabilityRequest - Request from client to check resource name availability.
type NameAvailabilityRequest struct {
	// REQUIRED; Resource name to verify.
	Name *string `json:"name,omitempty"`

	// Resource type used for verification.
	Type *string `json:"type,omitempty"`
}

// Network related properties of a server
type Network struct {
	// Delegated subnet resource id used to setup vnet for a server.
	DelegatedSubnetResourceID *string `json:"delegatedSubnetResourceId,omitempty"`

	// Private DNS zone resource id.
	PrivateDNSZoneResourceID *string `json:"privateDnsZoneResourceId,omitempty"`

	// READ-ONLY; Whether or not public network access is allowed for this server. Value is 'Disabled' when server has VNet integration.
	PublicNetworkAccess *EnableStatusEnum `json:"publicNetworkAccess,omitempty" azure:"ro"`
}

// Operation - REST API operation definition.
type Operation struct {
	// The localized display information for this particular operation or action.
	Display *OperationDisplay `json:"display,omitempty"`

	// The name of the operation being performed on this particular object.
	Name *string `json:"name,omitempty"`

	// The intended executor of the operation.
	Origin *string `json:"origin,omitempty"`

	// Additional descriptions for the operation.
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// OperationDisplay - Display metadata associated with the operation.
type OperationDisplay struct {
	// Operation description.
	Description *string `json:"description,omitempty"`

	// Localized friendly name for the operation.
	Operation *string `json:"operation,omitempty"`

	// Operation resource provider name.
	Provider *string `json:"provider,omitempty"`

	// Resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
}

// OperationListResult - A list of resource provider operations.
type OperationListResult struct {
	// URL client should use to fetch the next page (per server side paging).
	NextLink *string `json:"nextLink,omitempty"`

	// Collection of available operation details
	Value []*Operation `json:"value,omitempty"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ReplicasClientListByServerOptions contains the optional parameters for the ReplicasClient.ListByServer method.
type ReplicasClientListByServerOptions struct {
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

// SKU - Billing information related properties of a server.
type SKU struct {
	// REQUIRED; The name of the sku, e.g. StandardD32sv3.
	Name *string `json:"name,omitempty"`

	// REQUIRED; The tier of the particular SKU, e.g. GeneralPurpose.
	Tier *SKUTier `json:"tier,omitempty"`
}

// SKUCapability - Sku capability
type SKUCapability struct {
	// READ-ONLY; vCore name
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; supported IOPS
	SupportedIops *int64 `json:"supportedIops,omitempty" azure:"ro"`

	// READ-ONLY; supported memory per vCore in MB
	SupportedMemoryPerVCoreMB *int64 `json:"supportedMemoryPerVCoreMB,omitempty" azure:"ro"`

	// READ-ONLY; supported vCores
	VCores *int64 `json:"vCores,omitempty" azure:"ro"`
}

// Server - Represents a server.
type Server struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// The cmk identity for the server.
	Identity *Identity `json:"identity,omitempty"`

	// Properties of the server.
	Properties *ServerProperties `json:"properties,omitempty"`

	// The SKU (pricing tier) of the server.
	SKU *SKU `json:"sku,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The system metadata relating to this resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ServerBackup - Server backup properties
type ServerBackup struct {
	// The properties of a server backup.
	Properties *ServerBackupProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The system metadata relating to this resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ServerBackupListResult - A list of server backups.
type ServerBackupListResult struct {
	// The link used to get the next page of operations.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of backups of a server.
	Value []*ServerBackup `json:"value,omitempty"`
}

// ServerBackupProperties - The properties of a server backup.
type ServerBackupProperties struct {
	// Backup type.
	BackupType *string `json:"backupType,omitempty"`

	// Backup completed time (ISO8601 format).
	CompletedTime *time.Time `json:"completedTime,omitempty"`

	// Backup source
	Source *string `json:"source,omitempty"`
}

// ServerEditionCapability - Server edition capabilities.
type ServerEditionCapability struct {
	// READ-ONLY; Server edition name
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; A list of supported server versions.
	SupportedServerVersions []*ServerVersionCapability `json:"supportedServerVersions,omitempty" azure:"ro"`

	// READ-ONLY; A list of supported storage editions
	SupportedStorageEditions []*StorageEditionCapability `json:"supportedStorageEditions,omitempty" azure:"ro"`
}

// ServerForUpdate - Parameters allowed to update for a server.
type ServerForUpdate struct {
	// The cmk identity for the server.
	Identity *Identity `json:"identity,omitempty"`

	// The properties that can be updated for a server.
	Properties *ServerPropertiesForUpdate `json:"properties,omitempty"`

	// The SKU (pricing tier) of the server.
	SKU *SKU `json:"sku,omitempty"`

	// Application-specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`
}

// ServerListResult - A list of servers.
type ServerListResult struct {
	// The link used to get the next page of operations.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of servers
	Value []*Server `json:"value,omitempty"`
}

// ServerProperties - The properties of a server.
type ServerProperties struct {
	// The administrator's login name of a server. Can only be specified when the server is being created (and is required for
	// creation).
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	// The password of the administrator login (required for server creation).
	AdministratorLoginPassword *string `json:"administratorLoginPassword,omitempty"`

	// availability Zone information of the server.
	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	// Backup related properties of a server.
	Backup *Backup `json:"backup,omitempty"`

	// The mode to create a new MySQL server.
	CreateMode *CreateMode `json:"createMode,omitempty"`

	// The Data Encryption for CMK.
	DataEncryption *DataEncryption `json:"dataEncryption,omitempty"`

	// High availability related properties of a server.
	HighAvailability *HighAvailability `json:"highAvailability,omitempty"`

	// Maintenance window of a server.
	MaintenanceWindow *MaintenanceWindow `json:"maintenanceWindow,omitempty"`

	// Network related properties of a server.
	Network *Network `json:"network,omitempty"`

	// The replication role.
	ReplicationRole *ReplicationRole `json:"replicationRole,omitempty"`

	// Restore point creation time (ISO8601 format), specifying the time to restore from.
	RestorePointInTime *time.Time `json:"restorePointInTime,omitempty"`

	// The source MySQL server id.
	SourceServerResourceID *string `json:"sourceServerResourceId,omitempty"`

	// Storage related properties of a server.
	Storage *Storage `json:"storage,omitempty"`

	// Server version.
	Version *ServerVersion `json:"version,omitempty"`

	// READ-ONLY; The fully qualified domain name of a server.
	FullyQualifiedDomainName *string `json:"fullyQualifiedDomainName,omitempty" azure:"ro"`

	// READ-ONLY; The maximum number of replicas that a primary server can have.
	ReplicaCapacity *int32 `json:"replicaCapacity,omitempty" azure:"ro"`

	// READ-ONLY; The state of a server.
	State *ServerState `json:"state,omitempty" azure:"ro"`
}

// ServerPropertiesForUpdate - The properties that can be updated for a server.
type ServerPropertiesForUpdate struct {
	// The password of the administrator login.
	AdministratorLoginPassword *string `json:"administratorLoginPassword,omitempty"`

	// Backup related properties of a server.
	Backup *Backup `json:"backup,omitempty"`

	// The Data Encryption for CMK.
	DataEncryption *DataEncryption `json:"dataEncryption,omitempty"`

	// High availability related properties of a server.
	HighAvailability *HighAvailability `json:"highAvailability,omitempty"`

	// Maintenance window of a server.
	MaintenanceWindow *MaintenanceWindow `json:"maintenanceWindow,omitempty"`

	// The replication role of the server.
	ReplicationRole *ReplicationRole `json:"replicationRole,omitempty"`

	// Storage related properties of a server.
	Storage *Storage `json:"storage,omitempty"`
}

// ServerRestartParameter - Server restart parameters.
type ServerRestartParameter struct {
	// The maximum allowed failover time in seconds.
	MaxFailoverSeconds *int32 `json:"maxFailoverSeconds,omitempty"`

	// Whether or not failover to standby server when restarting a server with high availability enabled.
	RestartWithFailover *EnableStatusEnum `json:"restartWithFailover,omitempty"`
}

// ServerVersionCapability - Server version capabilities.
type ServerVersionCapability struct {
	// READ-ONLY; server version
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; A list of supported Skus
	SupportedSKUs []*SKUCapability `json:"supportedSkus,omitempty" azure:"ro"`
}

// ServersClientBeginCreateOptions contains the optional parameters for the ServersClient.BeginCreate method.
type ServersClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServersClientBeginDeleteOptions contains the optional parameters for the ServersClient.BeginDelete method.
type ServersClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServersClientBeginFailoverOptions contains the optional parameters for the ServersClient.BeginFailover method.
type ServersClientBeginFailoverOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServersClientBeginRestartOptions contains the optional parameters for the ServersClient.BeginRestart method.
type ServersClientBeginRestartOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServersClientBeginStartOptions contains the optional parameters for the ServersClient.BeginStart method.
type ServersClientBeginStartOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServersClientBeginStopOptions contains the optional parameters for the ServersClient.BeginStop method.
type ServersClientBeginStopOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServersClientBeginUpdateOptions contains the optional parameters for the ServersClient.BeginUpdate method.
type ServersClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServersClientGetOptions contains the optional parameters for the ServersClient.Get method.
type ServersClientGetOptions struct {
	// placeholder for future optional parameters
}

// ServersClientListByResourceGroupOptions contains the optional parameters for the ServersClient.ListByResourceGroup method.
type ServersClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ServersClientListOptions contains the optional parameters for the ServersClient.List method.
type ServersClientListOptions struct {
	// placeholder for future optional parameters
}

// Storage Profile properties of a server
type Storage struct {
	// Enable Storage Auto Grow or not.
	AutoGrow *EnableStatusEnum `json:"autoGrow,omitempty"`

	// Storage IOPS for a server.
	Iops *int32 `json:"iops,omitempty"`

	// Max storage size allowed for a server.
	StorageSizeGB *int32 `json:"storageSizeGB,omitempty"`

	// READ-ONLY; The sku name of the server storage.
	StorageSKU *string `json:"storageSku,omitempty" azure:"ro"`
}

// StorageEditionCapability - storage edition capability
type StorageEditionCapability struct {
	// READ-ONLY; Maximum backup retention days
	MaxBackupRetentionDays *int64 `json:"maxBackupRetentionDays,omitempty" azure:"ro"`

	// READ-ONLY; The maximum supported storage size.
	MaxStorageSize *int64 `json:"maxStorageSize,omitempty" azure:"ro"`

	// READ-ONLY; Minimal backup retention days
	MinBackupRetentionDays *int64 `json:"minBackupRetentionDays,omitempty" azure:"ro"`

	// READ-ONLY; The minimal supported storage size.
	MinStorageSize *int64 `json:"minStorageSize,omitempty" azure:"ro"`

	// READ-ONLY; storage edition name
	Name *string `json:"name,omitempty" azure:"ro"`
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

// UserAssignedIdentity - Metadata of user assigned identity.
type UserAssignedIdentity struct {
	// READ-ONLY; Client Id of user assigned identity
	ClientID *string `json:"clientId,omitempty" azure:"ro"`

	// READ-ONLY; Principal Id of user assigned identity
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`
}

// VirtualNetworkSubnetUsageParameter - Virtual network subnet usage parameter
type VirtualNetworkSubnetUsageParameter struct {
	// Virtual network resource id.
	VirtualNetworkResourceID *string `json:"virtualNetworkResourceId,omitempty"`
}

// VirtualNetworkSubnetUsageResult - Virtual network subnet usage data.
type VirtualNetworkSubnetUsageResult struct {
	// READ-ONLY; A list of delegated subnet usage
	DelegatedSubnetsUsage []*DelegatedSubnetUsage `json:"delegatedSubnetsUsage,omitempty" azure:"ro"`
}
