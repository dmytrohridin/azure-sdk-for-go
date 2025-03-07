//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_ListByReplicationProtectionContainers.json
func ExampleReplicationProtectedItemsClient_NewListByReplicationProtectionContainersPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("vault1",
		"resourceGroupPS1",
		"c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByReplicationProtectionContainersPager("cloud1",
		"cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_Get.json
func ExampleReplicationProtectedItemsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("vault1",
		"resourceGroupPS1",
		"c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"cloud1",
		"cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
		"f8491e4f-817a-40dd-a90c-af773978c75b",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_Create.json
func ExampleReplicationProtectedItemsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("vault1",
		"resourceGroupPS1",
		"c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"cloud1",
		"cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
		"f8491e4f-817a-40dd-a90c-af773978c75b",
		armrecoveryservicessiterecovery.EnableProtectionInput{
			Properties: &armrecoveryservicessiterecovery.EnableProtectionInputProperties{
				PolicyID:          to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/protectionprofile1"),
				ProtectableItemID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationProtectionContainers/cloud_6d224fc6-f326-5d35-96de-fbf51efb3179/replicationProtectableItems/f8491e4f-817a-40dd-a90c-af773978c75b"),
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.HyperVReplicaAzureEnableProtectionInput{
					InstanceType: to.Ptr("HyperVReplicaAzure"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_Purge.json
func ExampleReplicationProtectedItemsClient_BeginPurge() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("vault1",
		"resourceGroupPS1",
		"c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginPurge(ctx,
		"cloud1",
		"cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
		"c0c14913-3d7a-48ea-9531-cc99e0e686e6",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_Update.json
func ExampleReplicationProtectedItemsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("vault1",
		"resourceGroupPS1",
		"c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"cloud1",
		"cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
		"f8491e4f-817a-40dd-a90c-af773978c75b",
		armrecoveryservicessiterecovery.UpdateReplicationProtectedItemInput{
			Properties: &armrecoveryservicessiterecovery.UpdateReplicationProtectedItemInputProperties{
				LicenseType: to.Ptr(armrecoveryservicessiterecovery.LicenseTypeWindowsServer),
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.HyperVReplicaAzureUpdateReplicationProtectedItemInput{
					InstanceType: to.Ptr("HyperVReplicaAzure"),
				},
				RecoveryAzureVMName:            to.Ptr("vm1"),
				RecoveryAzureVMSize:            to.Ptr("Basic_A0"),
				SelectedRecoveryAzureNetworkID: to.Ptr("/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/siterecoveryProd1/providers/Microsoft.Network/virtualNetworks/vnetavrai"),
				VMNics: []*armrecoveryservicessiterecovery.VMNicInputDetails{
					{
						IPConfigs: []*armrecoveryservicessiterecovery.IPConfigInputDetails{
							{
								IPConfigName:            to.Ptr("ipconfig1"),
								IsPrimary:               to.Ptr(true),
								RecoveryStaticIPAddress: to.Ptr("10.0.2.46"),
								RecoverySubnetName:      to.Ptr("subnet1"),
							}},
						NicID:         to.Ptr("TWljcm9zb2Z0OkY4NDkxRTRGLTgxN0EtNDBERC1BOTBDLUFGNzczOTc4Qzc1Qlw3NjAwMzMxRS03NDk4LTQ0QTQtQjdDNy0xQjY1NkJDREQ1MkQ="),
						SelectionType: to.Ptr("SelectedByUser"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_AddDisks.json
func ExampleReplicationProtectedItemsClient_BeginAddDisks() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("vault1",
		"resourceGroupPS1",
		"c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginAddDisks(ctx,
		"cloud1",
		"cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
		"f8491e4f-817a-40dd-a90c-af773978c75b",
		armrecoveryservicessiterecovery.AddDisksInput{
			Properties: &armrecoveryservicessiterecovery.AddDisksInputProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2AAddDisksInput{
					InstanceType: to.Ptr("A2A"),
					VMDisks: []*armrecoveryservicessiterecovery.A2AVMDiskInputDetails{
						{
							DiskURI:                             to.Ptr("https://vmstorage.blob.core.windows.net/vhds/datadisk1.vhd"),
							PrimaryStagingAzureStorageAccountID: to.Ptr("/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourcegroups/primaryResource/providers/Microsoft.Storage/storageAccounts/vmcachestorage"),
							RecoveryAzureStorageAccountID:       to.Ptr("/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourcegroups/recoveryResource/providers/Microsoft.Storage/storageAccounts/recoverystorage"),
						}},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_ApplyRecoveryPoint.json
func ExampleReplicationProtectedItemsClient_BeginApplyRecoveryPoint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("vault1",
		"resourceGroupPS1",
		"c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginApplyRecoveryPoint(ctx,
		"cloud1",
		"cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
		"f8491e4f-817a-40dd-a90c-af773978c75b",
		armrecoveryservicessiterecovery.ApplyRecoveryPointInput{
			Properties: &armrecoveryservicessiterecovery.ApplyRecoveryPointInputProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.HyperVReplicaAzureApplyRecoveryPointInput{
					InstanceType: to.Ptr("HyperVReplicaAzure"),
				},
				RecoveryPointID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationProtectionContainers/cloud_6d224fc6-f326-5d35-96de-fbf51efb3179/replicationProtectedItems/f8491e4f-817a-40dd-a90c-af773978c75b/recoveryPoints/e4d05fe9-5dfd-47be-b50b-aad306b2802d"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_FailoverCancel.json
func ExampleReplicationProtectedItemsClient_BeginFailoverCancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("vault1",
		"resourceGroupPS1",
		"c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginFailoverCancel(ctx,
		"cloud1",
		"cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
		"f8491e4f-817a-40dd-a90c-af773978c75b",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_FailoverCommit.json
func ExampleReplicationProtectedItemsClient_BeginFailoverCommit() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("vault1",
		"resourceGroupPS1",
		"c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginFailoverCommit(ctx,
		"cloud1",
		"cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
		"f8491e4f-817a-40dd-a90c-af773978c75b",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_PlannedFailover.json
func ExampleReplicationProtectedItemsClient_BeginPlannedFailover() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("vault1",
		"resourceGroupPS1",
		"c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginPlannedFailover(ctx,
		"cloud1",
		"cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
		"f8491e4f-817a-40dd-a90c-af773978c75b",
		armrecoveryservicessiterecovery.PlannedFailoverInput{
			Properties: &armrecoveryservicessiterecovery.PlannedFailoverInputProperties{
				FailoverDirection: to.Ptr("PrimaryToRecovery"),
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.HyperVReplicaAzurePlannedFailoverProviderInput{
					InstanceType: to.Ptr("HyperVReplicaAzure"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_Delete.json
func ExampleReplicationProtectedItemsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("vault1",
		"resourceGroupPS1",
		"c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"cloud1",
		"cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
		"c0c14913-3d7a-48ea-9531-cc99e0e686e6",
		armrecoveryservicessiterecovery.DisableProtectionInput{
			Properties: &armrecoveryservicessiterecovery.DisableProtectionInputProperties{
				ReplicationProviderInput: &armrecoveryservicessiterecovery.DisableProtectionProviderSpecificInput{
					InstanceType: to.Ptr("DisableProtectionProviderSpecificInput"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_RemoveDisks.json
func ExampleReplicationProtectedItemsClient_BeginRemoveDisks() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("vault1",
		"resourceGroupPS1",
		"c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginRemoveDisks(ctx,
		"cloud1",
		"cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
		"f8491e4f-817a-40dd-a90c-af773978c75b",
		armrecoveryservicessiterecovery.RemoveDisksInput{
			Properties: &armrecoveryservicessiterecovery.RemoveDisksInputProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2ARemoveDisksInput{
					InstanceType: to.Ptr("A2A"),
					VMDisksUris: []*string{
						to.Ptr("https://vmstorage.blob.core.windows.net/vhds/datadisk1.vhd")},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_RepairReplication.json
func ExampleReplicationProtectedItemsClient_BeginRepairReplication() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("vault1",
		"resourceGroupPS1",
		"c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginRepairReplication(ctx,
		"cloud1",
		"cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
		"f8491e4f-817a-40dd-a90c-af773978c75b",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_Reprotect.json
func ExampleReplicationProtectedItemsClient_BeginReprotect() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("vault1",
		"resourceGroupPS1",
		"c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginReprotect(ctx,
		"cloud1",
		"cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
		"f8491e4f-817a-40dd-a90c-af773978c75b",
		armrecoveryservicessiterecovery.ReverseReplicationInput{
			Properties: &armrecoveryservicessiterecovery.ReverseReplicationInputProperties{
				FailoverDirection: to.Ptr("PrimaryToRecovery"),
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.HyperVReplicaAzureReprotectInput{
					InstanceType: to.Ptr("HyperVReplicaAzure"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_ResolveHealthErrors.json
func ExampleReplicationProtectedItemsClient_BeginResolveHealthErrors() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("vault1",
		"resourceGroupPS1",
		"c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginResolveHealthErrors(ctx,
		"cloud1",
		"cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
		"f8491e4f-817a-40dd-a90c-af773978c75b",
		armrecoveryservicessiterecovery.ResolveHealthInput{
			Properties: &armrecoveryservicessiterecovery.ResolveHealthInputProperties{
				HealthErrors: []*armrecoveryservicessiterecovery.ResolveHealthError{
					{
						HealthErrorID: to.Ptr("3:8020"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_SwitchProvider.json
func ExampleReplicationProtectedItemsClient_BeginSwitchProvider() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("vault1",
		"resourceGroupPS1",
		"c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginSwitchProvider(ctx,
		"cloud1",
		"cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
		"f8491e4f-817a-40dd-a90c-af773978c75b",
		armrecoveryservicessiterecovery.SwitchProviderInput{
			Properties: &armrecoveryservicessiterecovery.SwitchProviderInputProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.InMageAzureV2SwitchProviderInput{
					InstanceType:      to.Ptr("InMageAzureV2"),
					TargetApplianceID: to.Ptr("5efaa202-e958-435e-8171-706bf735fcc4"),
					TargetFabricID:    to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud2"),
					TargetVaultID:     to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault2"),
				},
				TargetInstanceType: to.Ptr("InMageRcm"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_TestFailover.json
func ExampleReplicationProtectedItemsClient_BeginTestFailover() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("vault1",
		"resourceGroupPS1",
		"c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginTestFailover(ctx,
		"cloud1",
		"cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
		"f8491e4f-817a-40dd-a90c-af773978c75b",
		armrecoveryservicessiterecovery.TestFailoverInput{
			Properties: &armrecoveryservicessiterecovery.TestFailoverInputProperties{
				FailoverDirection: to.Ptr("PrimaryToRecovery"),
				NetworkID:         to.Ptr("/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/siterecoveryProd1/providers/Microsoft.Network/virtualNetworks/vnetavrai"),
				NetworkType:       to.Ptr("VmNetworkAsInput"),
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.HyperVReplicaAzureTestFailoverInput{
					InstanceType: to.Ptr("HyperVReplicaAzure"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_TestFailoverCleanup.json
func ExampleReplicationProtectedItemsClient_BeginTestFailoverCleanup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("vault1",
		"resourceGroupPS1",
		"c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginTestFailoverCleanup(ctx,
		"cloud1",
		"cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
		"f8491e4f-817a-40dd-a90c-af773978c75b",
		armrecoveryservicessiterecovery.TestFailoverCleanupInput{
			Properties: &armrecoveryservicessiterecovery.TestFailoverCleanupInputProperties{
				Comments: to.Ptr("Test Failover Cleanup"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_UnplannedFailover.json
func ExampleReplicationProtectedItemsClient_BeginUnplannedFailover() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("vault1",
		"resourceGroupPS1",
		"c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUnplannedFailover(ctx,
		"cloud1",
		"cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
		"f8491e4f-817a-40dd-a90c-af773978c75b",
		armrecoveryservicessiterecovery.UnplannedFailoverInput{
			Properties: &armrecoveryservicessiterecovery.UnplannedFailoverInputProperties{
				FailoverDirection: to.Ptr("PrimaryToRecovery"),
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.HyperVReplicaAzureUnplannedFailoverInput{
					InstanceType: to.Ptr("HyperVReplicaAzure"),
				},
				SourceSiteOperations: to.Ptr("NotRequired"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_UpdateAppliance.json
func ExampleReplicationProtectedItemsClient_BeginUpdateAppliance() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("Ayan-0106-SA-Vault",
		"Ayan-0106-SA-RG",
		"7c943c1b-5122-4097-90c8-861411bdd574", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdateAppliance(ctx,
		"Ayan-0106-SA-Vaultreplicationfabric",
		"Ayan-0106-SA-Vaultreplicationcontainer",
		"idclab-vcen67_50158124-8857-3e08-0893-2ddf8ebb8c1f",
		armrecoveryservicessiterecovery.UpdateApplianceForReplicationProtectedItemInput{
			Properties: &armrecoveryservicessiterecovery.UpdateApplianceForReplicationProtectedItemInputProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.InMageRcmUpdateApplianceForReplicationProtectedItemInput{
					InstanceType:   to.Ptr("InMageRcm"),
					RunAsAccountID: to.Ptr(""),
				},
				TargetApplianceID: to.Ptr(""),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_UpdateMobilityService.json
func ExampleReplicationProtectedItemsClient_BeginUpdateMobilityService() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("WCUSVault",
		"wcusValidations",
		"b364ed8d-4279-4bf8-8fd1-56f8fa0ae05c", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdateMobilityService(ctx,
		"WIN-JKKJ31QI8U2",
		"cloud_c6780228-83bd-4f3e-a70e-cb46b7da33a0",
		"79dd20ab-2b40-11e7-9791-0050568f387e",
		armrecoveryservicessiterecovery.UpdateMobilityServiceRequest{
			Properties: &armrecoveryservicessiterecovery.UpdateMobilityServiceRequestProperties{
				RunAsAccountID: to.Ptr("2"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_List.json
func ExampleReplicationProtectedItemsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("vault1",
		"resourceGroupPS1",
		"c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager(&armrecoveryservicessiterecovery.ReplicationProtectedItemsClientListOptions{SkipToken: nil,
		Filter: nil,
	})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
