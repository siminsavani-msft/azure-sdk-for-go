//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ListJobsByAgent.json
func ExampleJobsClient_NewListByAgentPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobsClient().NewListByAgentPager("group1", "server1", "agent1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.JobListResult = armsql.JobListResult{
		// 	Value: []*armsql.Job{
		// 		{
		// 			Name: to.Ptr("job1"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/jobAccounts/jobs"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/jobs/job1"),
		// 			Properties: &armsql.JobProperties{
		// 				Description: to.Ptr("my favourite job"),
		// 				Schedule: &armsql.JobSchedule{
		// 					Type: to.Ptr(armsql.JobScheduleTypeRecurring),
		// 					Enabled: to.Ptr(true),
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-09-24T23:59:59Z"); return t}()),
		// 					Interval: to.Ptr("PT5M"),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-09-24T18:30:01Z"); return t}()),
		// 				},
		// 				Version: to.Ptr[int32](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("job3"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/jobAccounts/jobs"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/jobs/job3"),
		// 			Properties: &armsql.JobProperties{
		// 				Description: to.Ptr("this job will be scheduled once"),
		// 				Schedule: &armsql.JobSchedule{
		// 					Type: to.Ptr(armsql.JobScheduleTypeOnce),
		// 					Enabled: to.Ptr(true),
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-27T06:00:00Z"); return t}()),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-07-01T03:45:00Z"); return t}()),
		// 				},
		// 				Version: to.Ptr[int32](1),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("job2"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/jobAccounts/jobs"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/jobs/job2"),
		// 			Properties: &armsql.JobProperties{
		// 				Description: to.Ptr("this job will never be automatically scheduled"),
		// 				Schedule: &armsql.JobSchedule{
		// 					Type: to.Ptr(armsql.JobScheduleTypeOnce),
		// 					Enabled: to.Ptr(false),
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-27T06:00:00Z"); return t}()),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-07-01T03:45:00Z"); return t}()),
		// 				},
		// 				Version: to.Ptr[int32](1),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetJob.json
func ExampleJobsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobsClient().Get(ctx, "group1", "server1", "agent1", "job1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Job = armsql.Job{
	// 	Name: to.Ptr("job1"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/jobAccounts/jobs"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/jobs/job1"),
	// 	Properties: &armsql.JobProperties{
	// 		Description: to.Ptr("my favourite job"),
	// 		Schedule: &armsql.JobSchedule{
	// 			Type: to.Ptr(armsql.JobScheduleTypeOnce),
	// 			Enabled: to.Ptr(true),
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-09-24T23:59:59Z"); return t}()),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-09-24T18:30:01Z"); return t}()),
	// 		},
	// 		Version: to.Ptr[int32](0),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateJobMax.json
func ExampleJobsClient_CreateOrUpdate_createAJobWithAllPropertiesSpecified() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobsClient().CreateOrUpdate(ctx, "group1", "server1", "agent1", "job1", armsql.Job{
		Properties: &armsql.JobProperties{
			Description: to.Ptr("my favourite job"),
			Schedule: &armsql.JobSchedule{
				Type:      to.Ptr(armsql.JobScheduleTypeRecurring),
				Enabled:   to.Ptr(true),
				EndTime:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-09-24T23:59:59Z"); return t }()),
				Interval:  to.Ptr("PT5M"),
				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-09-24T18:30:01Z"); return t }()),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Job = armsql.Job{
	// 	Name: to.Ptr("job1"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/jobAccounts/jobs"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/jobs/job1"),
	// 	Properties: &armsql.JobProperties{
	// 		Description: to.Ptr("my favourite job"),
	// 		Schedule: &armsql.JobSchedule{
	// 			Type: to.Ptr(armsql.JobScheduleTypeRecurring),
	// 			Enabled: to.Ptr(true),
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-09-24T23:59:59Z"); return t}()),
	// 			Interval: to.Ptr("PT5M"),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-09-24T18:30:01Z"); return t}()),
	// 		},
	// 		Version: to.Ptr[int32](0),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateJobMin.json
func ExampleJobsClient_CreateOrUpdate_createAJobWithDefaultProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobsClient().CreateOrUpdate(ctx, "group1", "server1", "agent1", "job1", armsql.Job{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Job = armsql.Job{
	// 	Name: to.Ptr("job1"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/jobAccounts/jobs"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/jobs/job1"),
	// 	Properties: &armsql.JobProperties{
	// 		Description: to.Ptr(""),
	// 		Schedule: &armsql.JobSchedule{
	// 			Type: to.Ptr(armsql.JobScheduleTypeOnce),
	// 			Enabled: to.Ptr(false),
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "9999-12-31T11:59:59Z"); return t}()),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T00:00:00Z"); return t}()),
	// 		},
	// 		Version: to.Ptr[int32](0),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DeleteJob.json
func ExampleJobsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewJobsClient().Delete(ctx, "group1", "server1", "agent1", "job1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
