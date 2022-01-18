//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/FailoverGroupGet.json
func ExampleFailoverGroupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewFailoverGroupsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<failover-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.FailoverGroupsClientGetResult)
}

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/FailoverGroupCreateOrUpdate.json
func ExampleFailoverGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewFailoverGroupsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<failover-group-name>",
		armsql.FailoverGroup{
			Properties: &armsql.FailoverGroupProperties{
				Databases: []*string{
					to.StringPtr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server/databases/testdb-1"),
					to.StringPtr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server/databases/testdb-2")},
				PartnerServers: []*armsql.PartnerInfo{
					{
						ID: to.StringPtr("<id>"),
					}},
				ReadOnlyEndpoint: &armsql.FailoverGroupReadOnlyEndpoint{
					FailoverPolicy: armsql.ReadOnlyEndpointFailoverPolicy("Disabled").ToPtr(),
				},
				ReadWriteEndpoint: &armsql.FailoverGroupReadWriteEndpoint{
					FailoverPolicy:                         armsql.ReadWriteEndpointFailoverPolicy("Automatic").ToPtr(),
					FailoverWithDataLossGracePeriodMinutes: to.Int32Ptr(480),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.FailoverGroupsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/FailoverGroupDelete.json
func ExampleFailoverGroupsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewFailoverGroupsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<failover-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/FailoverGroupUpdate.json
func ExampleFailoverGroupsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewFailoverGroupsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<failover-group-name>",
		armsql.FailoverGroupUpdate{
			Properties: &armsql.FailoverGroupUpdateProperties{
				Databases: []*string{
					to.StringPtr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server/databases/testdb-1")},
				ReadWriteEndpoint: &armsql.FailoverGroupReadWriteEndpoint{
					FailoverPolicy:                         armsql.ReadWriteEndpointFailoverPolicy("Automatic").ToPtr(),
					FailoverWithDataLossGracePeriodMinutes: to.Int32Ptr(120),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.FailoverGroupsClientUpdateResult)
}

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/FailoverGroupList.json
func ExampleFailoverGroupsClient_ListByServer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewFailoverGroupsClient("<subscription-id>", cred, nil)
	pager := client.ListByServer("<resource-group-name>",
		"<server-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/FailoverGroupFailover.json
func ExampleFailoverGroupsClient_BeginFailover() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewFailoverGroupsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginFailover(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<failover-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.FailoverGroupsClientFailoverResult)
}

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/FailoverGroupForceFailoverAllowDataLoss.json
func ExampleFailoverGroupsClient_BeginForceFailoverAllowDataLoss() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewFailoverGroupsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginForceFailoverAllowDataLoss(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<failover-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.FailoverGroupsClientForceFailoverAllowDataLossResult)
}