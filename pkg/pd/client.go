// Copyright 2016 DeepFabric, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package pd

import (
	"github.com/deepfabric/elasticell/pkg/storage"
)

// Client pd client
type Client struct {
}

// AllocID ask pd for a uniq id
func (c *Client) AllocID() (int64, error) {
	return 0, nil
}

// GetClusterID get cluster id from pd
func (c *Client) GetClusterID() (int64, error) {
	return 0, nil
}

// IsClusterBootstrapped ask pd, the cluster is bootstrapped.
func (c *Client) IsClusterBootstrapped() (bool, error) {
	return false, nil
}

// BootstrapCluster tell pd to bootstart cluster.
func (c *Client) BootstrapCluster(cell *storage.Cell) error {
	return nil
}

// TellPDStoreStarted tell pd the store on this node is started.
func (c *Client) TellPDStoreStarted(store *storage.Store) error {
	return nil
}
