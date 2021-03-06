package servicefabricapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/servicefabric/mgmt/2017-07-01-preview/servicefabric"
	"github.com/Azure/go-autorest/autorest"
)

// ClustersClientAPI contains the set of methods on the ClustersClient type.
type ClustersClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, clusterName string, APIVersion string, subscriptionID string, parameters servicefabric.Cluster) (result servicefabric.ClustersCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string, APIVersion string, subscriptionID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string, APIVersion string, subscriptionID string) (result servicefabric.Cluster, err error)
	List(ctx context.Context, APIVersion string, subscriptionID string) (result servicefabric.ClusterListResult, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, APIVersion string, subscriptionID string) (result servicefabric.ClusterListResult, err error)
	Update(ctx context.Context, resourceGroupName string, clusterName string, APIVersion string, subscriptionID string, parameters servicefabric.ClusterUpdateParameters) (result servicefabric.ClustersUpdateFuture, err error)
}

var _ ClustersClientAPI = (*servicefabric.ClustersClient)(nil)

// ClusterVersionsClientAPI contains the set of methods on the ClusterVersionsClient type.
type ClusterVersionsClientAPI interface {
	Get(ctx context.Context, location string, APIVersion string, subscriptionID string, clusterVersion string) (result servicefabric.ClusterCodeVersionsListResult, err error)
	GetByEnvironment(ctx context.Context, location string, environment string, APIVersion string, subscriptionID string, clusterVersion string) (result servicefabric.ClusterCodeVersionsListResult, err error)
	List(ctx context.Context, location string, APIVersion string, subscriptionID string) (result servicefabric.ClusterCodeVersionsListResult, err error)
	ListByEnvironment(ctx context.Context, location string, environment string, APIVersion string, subscriptionID string) (result servicefabric.ClusterCodeVersionsListResult, err error)
}

var _ ClusterVersionsClientAPI = (*servicefabric.ClusterVersionsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context, APIVersion string) (result servicefabric.OperationListResultPage, err error)
}

var _ OperationsClientAPI = (*servicefabric.OperationsClient)(nil)
