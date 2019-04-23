// Copyright 2019 Smart-Edge.com, Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package clients

import (
	"context"

	"github.com/pkg/errors"
	"github.com/smartedgemec/controller-ce/grpc"
	"github.com/smartedgemec/controller-ce/pb"
)

// ApplicationDeploymentServiceClient wraps the PB client.
type ApplicationDeploymentServiceClient struct {
	PBCli pb.ApplicationDeploymentServiceClient
}

// NewApplicationDeploymentServiceClient creates a new client.
func NewApplicationDeploymentServiceClient(
	conn *grpc.ClientConn,
) *ApplicationDeploymentServiceClient {
	return &ApplicationDeploymentServiceClient{
		conn.NewApplicationDeploymentServiceClient(),
	}
}

// DeployContainer deploys a container application.
func (c *ApplicationDeploymentServiceClient) DeployContainer(
	ctx context.Context,
	app *pb.Application,
) error {
	_, err := c.PBCli.DeployContainer(
		ctx,
		app)

	if err != nil {
		return errors.Wrap(err, "error deploying application")
	}

	return nil
}

// DeployVM deploys a VM application.
func (c *ApplicationDeploymentServiceClient) DeployVM(
	ctx context.Context,
	app *pb.Application,
) error {
	_, err := c.PBCli.DeployVM(
		ctx,
		app)

	if err != nil {
		return errors.Wrap(err, "error deploying application")
	}

	return nil
}

// GetStatus retrieves an application's status.
func (c *ApplicationDeploymentServiceClient) GetStatus(
	ctx context.Context,
	id string,
) (*pb.LifecycleStatus, error) {
	status, err := c.PBCli.GetStatus(
		ctx,
		&pb.ApplicationID{Id: id})

	if err != nil {
		return nil, errors.Wrap(err, "error retrieving application")
	}

	return status, nil
}

// Redeploy redeploys an application.
func (c *ApplicationDeploymentServiceClient) Redeploy(
	ctx context.Context,
	app *pb.Application,
) error {
	_, err := c.PBCli.Redeploy(
		ctx,
		app)

	if err != nil {
		return errors.Wrap(err, "error redeploying application")
	}

	return nil
}

// Undeploy undeploys an application.
func (c *ApplicationDeploymentServiceClient) Undeploy(
	ctx context.Context,
	id string,
) error {
	_, err := c.PBCli.Undeploy(
		ctx,
		&pb.ApplicationID{
			Id: id,
		})

	if err != nil {
		return errors.Wrap(err, "error removing application")
	}

	return nil
}