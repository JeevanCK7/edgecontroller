// Copyright 2019 Smart-Edge.com, Inc. All rights reserved.
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

package node

import (
	"context"

	cce "github.com/smartedgemec/controller-ce"
	"github.com/smartedgemec/controller-ce/grpc"
	gclients "github.com/smartedgemec/controller-ce/grpc/clients"
)

// ClientConn wraps a Node and provides a Connect() method to create wrapped
// gRPC clients.
// TODO does this need an interface?
type ClientConn struct {
	Node *cce.Node

	conn *grpc.ClientConn

	AppDeploySvcCli *gclients.ApplicationDeploymentServiceClient
	AppLifeSvcCli   *gclients.ApplicationLifecycleServiceClient
	AppPolicySvcCli *gclients.ApplicationPolicyServiceClient
	VNFDeploySvcCli *gclients.VNFDeploymentServiceClient
	VNFLifeSvcCli   *gclients.VNFLifecycleServiceClient
}

// Connect connects to a node via grpc.Dial.
func (cc *ClientConn) Connect(ctx context.Context) error {
	var err error
	if cc.conn, err = grpc.Dial(ctx, cc.Node.GRPCTarget); err != nil {
		return err
	}

	cc.AppDeploySvcCli = gclients.NewApplicationDeploymentServiceClient(cc.conn)
	cc.AppLifeSvcCli = gclients.NewApplicationLifecycleServiceClient(cc.conn)
	cc.AppPolicySvcCli = gclients.NewApplicationPolicyServiceClient(cc.conn)
	cc.VNFDeploySvcCli = gclients.NewVNFDeploymentServiceClient(cc.conn)
	cc.VNFLifeSvcCli = gclients.NewVNFLifecycleServiceClient(cc.conn)

	return nil
}