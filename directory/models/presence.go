// Copyright 2020 Nym Technologies SA
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

package models

// NodeInfo comes from a node telling us it's alive
type NodeInfo struct {
	MixHost     string `json:"mixHost" binding:"required"`
	IdentityKey string `json:"identityKey" binding:"required" gorm:"primaryKey;unique"`
	SphinxKey   string `json:"sphinxKey" binding:"required"`
	Version     string `json:"version" binding:"required"`
	Location    string `json:"location"`
	// ideally it would also involve a signature, but it's fine for time being
}

type MixRegistrationInfo struct {
	NodeInfo
	Layer uint `json:"layer" binding:"required"`
}

type RegisteredMix struct {
	MixRegistrationInfo
	RegistrationTime int64 `json:"registrationTime" gorm:"autoCreateTime:nano"`
	Reputation       int64 `json:"reputation"`
}

type GatewayRegistrationInfo struct {
	NodeInfo
	ClientsHost string `json:"clientsHost" binding:"required"`
}

type RegisteredGateway struct {
	GatewayRegistrationInfo
	RegistrationTime int64 `json:"registrationTime" gorm:"autoCreateTime:nano"`
	Reputation       int64 `json:"reputation"`
}

type Topology struct {
	MixNodes []RegisteredMix     `json:"mixNodes" binding:"required"`
	Gateways []RegisteredGateway `json:"gateways" binding:"required"`
}
