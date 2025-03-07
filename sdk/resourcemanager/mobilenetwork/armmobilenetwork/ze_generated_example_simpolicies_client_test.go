//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SimPolicyDelete.json
func ExampleSimPoliciesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewSimPoliciesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"rg1",
		"testMobileNetwork",
		"testPolicy",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SimPolicyGet.json
func ExampleSimPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewSimPoliciesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"rg1",
		"testMobileNetwork",
		"testPolicy",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SimPolicyCreate.json
func ExampleSimPoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewSimPoliciesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"testMobileNetwork",
		"testPolicy",
		armmobilenetwork.SimPolicy{
			Location: to.Ptr("eastus"),
			Properties: &armmobilenetwork.SimPolicyPropertiesFormat{
				DefaultSlice: &armmobilenetwork.SliceResourceID{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
				},
				RegistrationTimer: to.Ptr[int32](3240),
				SliceConfigurations: []*armmobilenetwork.SliceConfiguration{
					{
						DataNetworkConfigurations: []*armmobilenetwork.DataNetworkConfiguration{
							{
								FiveQi:                              to.Ptr[int32](9),
								AdditionalAllowedSessionTypes:       []*armmobilenetwork.PduSessionType{},
								AllocationAndRetentionPriorityLevel: to.Ptr[int32](9),
								AllowedServices: []*armmobilenetwork.ServiceResourceID{
									{
										ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/services/testService"),
									}},
								DataNetwork: &armmobilenetwork.DataNetworkResourceID{
									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/dataNetworks/testdataNetwork"),
								},
								DefaultSessionType:      to.Ptr(armmobilenetwork.PduSessionTypeIPv4),
								PreemptionCapability:    to.Ptr(armmobilenetwork.PreemptionCapabilityNotPreempt),
								PreemptionVulnerability: to.Ptr(armmobilenetwork.PreemptionVulnerabilityPreemptable),
								SessionAmbr: &armmobilenetwork.Ambr{
									Downlink: to.Ptr("1 Gbps"),
									Uplink:   to.Ptr("500 Mbps"),
								},
							}},
						DefaultDataNetwork: &armmobilenetwork.DataNetworkResourceID{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/dataNetworks/testdataNetwork"),
						},
						Slice: &armmobilenetwork.SliceResourceID{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
						},
					}},
				UeAmbr: &armmobilenetwork.Ambr{
					Downlink: to.Ptr("1 Gbps"),
					Uplink:   to.Ptr("500 Mbps"),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SimPolicyUpdateTags.json
func ExampleSimPoliciesClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewSimPoliciesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.UpdateTags(ctx,
		"rg1",
		"testMobileNetwork",
		"testPolicy",
		armmobilenetwork.TagsObject{
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SimPolicyListByMobileNetwork.json
func ExampleSimPoliciesClient_NewListByMobileNetworkPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewSimPoliciesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByMobileNetworkPager("testResourceGroupName",
		"testMobileNetwork",
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
