// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by aliasgen. DO NOT EDIT.

// Package asset aliases all exported identifiers in package
// "cloud.google.com/go/asset/apiv1p2beta1/assetpb".
//
// Deprecated: Please use types in: cloud.google.com/go/asset/apiv1p2beta1/assetpb
package asset

import (
	src "cloud.google.com/go/asset/apiv1p2beta1/assetpb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/asset/apiv1p2beta1/assetpb
const (
	ContentType_CONTENT_TYPE_UNSPECIFIED = src.ContentType_CONTENT_TYPE_UNSPECIFIED
	ContentType_IAM_POLICY               = src.ContentType_IAM_POLICY
	ContentType_RESOURCE                 = src.ContentType_RESOURCE
)

// Deprecated: Please use vars in: cloud.google.com/go/asset/apiv1p2beta1/assetpb
var (
	ContentType_name                                      = src.ContentType_name
	ContentType_value                                     = src.ContentType_value
	File_google_cloud_asset_v1p2beta1_asset_service_proto = src.File_google_cloud_asset_v1p2beta1_asset_service_proto
	File_google_cloud_asset_v1p2beta1_assets_proto        = src.File_google_cloud_asset_v1p2beta1_assets_proto
)

// Cloud asset. This includes all Google Cloud Platform resources, Cloud IAM
// policies, and other non-GCP assets.
//
// Deprecated: Please use types in: cloud.google.com/go/asset/apiv1p2beta1/assetpb
type Asset = src.Asset

// AssetServiceClient is the client API for AssetService service. For
// semantics around ctx use and closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/asset/apiv1p2beta1/assetpb
type AssetServiceClient = src.AssetServiceClient

// AssetServiceServer is the server API for AssetService service.
//
// Deprecated: Please use types in: cloud.google.com/go/asset/apiv1p2beta1/assetpb
type AssetServiceServer = src.AssetServiceServer

// Asset content type.
//
// Deprecated: Please use types in: cloud.google.com/go/asset/apiv1p2beta1/assetpb
type ContentType = src.ContentType

// Create asset feed request.
//
// Deprecated: Please use types in: cloud.google.com/go/asset/apiv1p2beta1/assetpb
type CreateFeedRequest = src.CreateFeedRequest
type DeleteFeedRequest = src.DeleteFeedRequest

// An asset feed used to export asset updates to a destinations. An asset feed
// filter controls what updates are exported. The asset feed must be created
// within a project, organization, or folder. Supported destinations are: Cloud
// Pub/Sub topics.
//
// Deprecated: Please use types in: cloud.google.com/go/asset/apiv1p2beta1/assetpb
type Feed = src.Feed

// Output configuration for asset feed destination.
//
// Deprecated: Please use types in: cloud.google.com/go/asset/apiv1p2beta1/assetpb
type FeedOutputConfig = src.FeedOutputConfig
type FeedOutputConfig_PubsubDestination = src.FeedOutputConfig_PubsubDestination

// A Cloud Storage location.
//
// Deprecated: Please use types in: cloud.google.com/go/asset/apiv1p2beta1/assetpb
type GcsDestination = src.GcsDestination
type GcsDestination_Uri = src.GcsDestination_Uri

// Get asset feed request.
//
// Deprecated: Please use types in: cloud.google.com/go/asset/apiv1p2beta1/assetpb
type GetFeedRequest = src.GetFeedRequest

// List asset feeds request.
//
// Deprecated: Please use types in: cloud.google.com/go/asset/apiv1p2beta1/assetpb
type ListFeedsRequest = src.ListFeedsRequest
type ListFeedsResponse = src.ListFeedsResponse

// Output configuration for export assets destination.
//
// Deprecated: Please use types in: cloud.google.com/go/asset/apiv1p2beta1/assetpb
type OutputConfig = src.OutputConfig
type OutputConfig_GcsDestination = src.OutputConfig_GcsDestination

// A Cloud Pubsub destination.
//
// Deprecated: Please use types in: cloud.google.com/go/asset/apiv1p2beta1/assetpb
type PubsubDestination = src.PubsubDestination

// Representation of a cloud resource.
//
// Deprecated: Please use types in: cloud.google.com/go/asset/apiv1p2beta1/assetpb
type Resource = src.Resource

// Temporal asset. In addition to the asset, the temporal asset includes the
// status of the asset and valid from and to time of it.
//
// Deprecated: Please use types in: cloud.google.com/go/asset/apiv1p2beta1/assetpb
type TemporalAsset = src.TemporalAsset

// A time window of (start_time, end_time].
//
// Deprecated: Please use types in: cloud.google.com/go/asset/apiv1p2beta1/assetpb
type TimeWindow = src.TimeWindow

// UnimplementedAssetServiceServer can be embedded to have forward compatible
// implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/asset/apiv1p2beta1/assetpb
type UnimplementedAssetServiceServer = src.UnimplementedAssetServiceServer

// Update asset feed request.
//
// Deprecated: Please use types in: cloud.google.com/go/asset/apiv1p2beta1/assetpb
type UpdateFeedRequest = src.UpdateFeedRequest

// Deprecated: Please use funcs in: cloud.google.com/go/asset/apiv1p2beta1/assetpb
func NewAssetServiceClient(cc grpc.ClientConnInterface) AssetServiceClient {
	return src.NewAssetServiceClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/asset/apiv1p2beta1/assetpb
func RegisterAssetServiceServer(s *grpc.Server, srv AssetServiceServer) {
	src.RegisterAssetServiceServer(s, srv)
}
