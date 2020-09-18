// Copyright © 2020 Cisco
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
//
// All rights reserved.

package servicedirectory

import (
	"context"

	sd "cloud.google.com/go/servicedirectory/apiv1beta1"
	gax "github.com/googleapis/gax-go"
	sdpb "google.golang.org/genproto/googleapis/cloud/servicedirectory/v1beta1"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

type fakeRegClient struct {
}

func getFakeHandler() *servDir {
	return &servDir{
		project: "project",
		region:  "us",
		context: context.Background(),
		client:  &fakeRegClient{},
		log:     zap.New(zap.UseDevMode(true)),
		timeout: 2,
	}
}

func (f *fakeRegClient) GetNamespace(ctx context.Context, req *sdpb.GetNamespaceRequest, opts ...gax.CallOption) (*sdpb.Namespace, error) {
	// TODO
	return nil, nil
}

func (f *fakeRegClient) ListNamespaces(ctx context.Context, req *sdpb.ListNamespacesRequest, opts ...gax.CallOption) *sd.NamespaceIterator {
	// TODO
	return nil
}

func (f *fakeRegClient) CreateNamespace(ctx context.Context, req *sdpb.CreateNamespaceRequest, opts ...gax.CallOption) (*sdpb.Namespace, error) {
	// TODO
	return nil, nil
}

func (f *fakeRegClient) UpdateNamespace(ctx context.Context, req *sdpb.UpdateNamespaceRequest, opts ...gax.CallOption) (*sdpb.Namespace, error) {
	// TODO
	return nil, nil
}

func (f *fakeRegClient) DeleteNamespace(ctx context.Context, req *sdpb.DeleteNamespaceRequest, opts ...gax.CallOption) error {
	// TODO
	return nil
}

func (f *fakeRegClient) GetService(ctx context.Context, req *sdpb.GetServiceRequest, opts ...gax.CallOption) (*sdpb.Service, error) {
	// TODO
	return nil, nil
}

func (f *fakeRegClient) CreateService(ctx context.Context, req *sdpb.CreateServiceRequest, opts ...gax.CallOption) (*sdpb.Service, error) {
	// TODO
	return nil, nil
}

func (f *fakeRegClient) UpdateService(ctx context.Context, req *sdpb.UpdateServiceRequest, opts ...gax.CallOption) (*sdpb.Service, error) {
	// TODO
	return nil, nil
}

func (f *fakeRegClient) DeleteService(ctx context.Context, req *sdpb.DeleteServiceRequest, opts ...gax.CallOption) error {
	// TODO
	return nil
}

func (f *fakeRegClient) GetEndpoint(ctx context.Context, req *sdpb.GetEndpointRequest, opts ...gax.CallOption) (*sdpb.Endpoint, error) {
	// TODO
	return nil, nil
}

func (f *fakeRegClient) CreateEndpoint(ctx context.Context, req *sdpb.CreateEndpointRequest, opts ...gax.CallOption) (*sdpb.Endpoint, error) {
	// TODO
	return nil, nil
}

func (f *fakeRegClient) UpdateEndpoint(ctx context.Context, req *sdpb.UpdateEndpointRequest, opts ...gax.CallOption) (*sdpb.Endpoint, error) {
	// TODO
	return nil, nil
}

func (f *fakeRegClient) DeleteEndpoint(ctx context.Context, req *sdpb.DeleteEndpointRequest, opts ...gax.CallOption) error {
	// TODO
	return nil
}

func (f *fakeRegClient) Close() error { return nil }

func (f *fakeRegClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return nil, nil
}

func (f *fakeRegClient) ListEndpoints(ctx context.Context, req *sdpb.ListEndpointsRequest, opts ...gax.CallOption) *sd.EndpointIterator {
	return nil
}

func (f *fakeRegClient) ListServices(ctx context.Context, req *sdpb.ListServicesRequest, opts ...gax.CallOption) *sd.ServiceIterator {
	return nil
}
func (f *fakeRegClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return nil, nil
}
func (f *fakeRegClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	return nil, nil
}