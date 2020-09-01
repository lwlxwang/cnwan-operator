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

package servregistry

// This file contains functions that perform operations on namespaces,
// such as create/update/delete.
// These functions belong to a ServiceRegistryBroker, defined in
// broker.go

// ManageNs takes data of a namespace and performs the necessary steps
// to reflect that data to the service registry.
//
// For example: create a namespace in service registry or update it
// properly.
func (b *Broker) ManageNs(nsData *Namespace) (regNs *Namespace, err error) {
	// nsData: data of the namespace in Kubernetes (latest state)
	// regNs: data of the namespace currently in the service registry

	if b.Reg == nil {
		return nil, ErrServRegNotProvided
	}

	// -- Validate
	if nsData == nil {
		return nil, ErrNsNotProvided
	}

	if len(nsData.Name) == 0 {
		return nil, ErrNsNameNotProvided
	}

	// -- Init
	b.lock.Lock()
	defer b.lock.Unlock()
	if nsData.Metadata == nil {
		nsData.Metadata = map[string]string{}
	}
	nsData.Metadata[b.opKey] = b.opVal
	l := b.log.WithName("ManageNs").WithValues("ns-name", nsData.Name)

	// -- Do stuff
	l.V(1).Info("going to load namespace from service registry")

	regNs, err = b.Reg.GetNs(nsData.Name)
	if err != nil {
		if err != ErrNotFound {
			l.Error(err, "error occurred while getting namespace from service registry")
			return
		}

		// If you're here, it means that the namespace does not exist.
		// Let's create it.
		l.V(1).Info("namespace does not exist in service registry, going to create it")
		regNs, err = b.Reg.CreateNs(nsData)
		if err != nil {
			l.Error(err, "error occurred while creating namespace in service registry")
			return
		}

		l.V(0).Info("namespace created correctly")
		regNs = nsData
	}

	if by, exists := regNs.Metadata[b.opKey]; by != b.opVal || !exists {
		// If the namespace is not owned (as in, managed by) us, then it's
		// better not to touch it.
		l.V(0).Info("namespace is not owned by the operator and thus will not be updated")
		return
	}

	if !b.deepEqualMetadata(nsData.Metadata, regNs.Metadata) {
		l.V(1).Info("namespace metadata need to be updated")
		regNs, err = b.Reg.UpdateNs(nsData)
		if err != nil {
			l.Error(err, "error while trying to update namespace in service registry")
			return nil, err
		}
	}

	return
}