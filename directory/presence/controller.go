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

package presence

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nymtech/nym/validator/nym/directory/models"
	"net/http"
)

type Config struct {
	Sanitizer Sanitizer
	Service IService
}

// controller is the presence controller
type controller struct {
	service        IService
	sanitizer      Sanitizer
}

type Controller interface {
	RegisterRoutes(router *gin.Engine)
}

func New(cfg Config) Controller {
	return &controller{
		service: cfg.Service,
		sanitizer: cfg.Sanitizer,
	}
}

func (controller *controller) RegisterRoutes(router *gin.Engine) {
	router.POST("/api/presence/mix", controller.RegisterMixPresence)
	router.POST("/api/presence/gateway", controller.RegisterGatewayPresence)
	router.DELETE("/api/presence/:id", controller.UnregisterPresence)
	router.PATCH("/api/presence/reputation/:id", controller.ChangeReputation)
	router.GET("/api/presence/topology", controller.GetTopology)
}

// RegisterMixPresence ...
// @Summary Lets a mixnode tell the directory server it's coming online
// @Description On Nym nodes startup they register their presence indicating they should be alive and get added to the set of active nodes in the topology.
// @ID registerMixPresence
// @Accept  json
// @Produce  json
// @Tags presence
// @Param   object      body   models.MixRegistrationInfo     true  "object"
// @Success 201
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/presence/mix [post]
func (controller *controller) RegisterMixPresence(ctx *gin.Context) {
	var presence models.MixRegistrationInfo
	if err := ctx.ShouldBindJSON(&presence); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	controller.sanitizer.Sanitize(&presence)
	controller.service.AddMixRegistrationPresence(presence)
	ctx.JSON(http.StatusOK, gin.H{"ok": true})
}

// RegisterGatewayPresence ...
// @Summary Lets a gateway tell the directory server it's coming online
// @Description On Nym nodes startup they register their presence indicating they should be alive and get added to the set of active nodes in the topology.
// @ID registerGatewayPresence
// @Accept  json
// @Produce  json
// @Tags presence
// @Param   object      body   models.GatewayRegistrationInfo     true  "object"
// @Success 201
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/presence/gateway [post]
func (controller *controller) RegisterGatewayPresence(ctx *gin.Context) {
	var presence models.GatewayRegistrationInfo
	if err := ctx.ShouldBindJSON(&presence); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("register gate presence: %+v\n", presence)}

// UnregisterPresence ...
// @Summary Unregister presence of node.
// @Description Messages sent by a node on powering down to indicate it's going offline so that it should get removed from active topology.
// @ID unregisterPresence
// @Accept  json
// @Produce  json
// @Tags presence
// @Param id path string true "Node Identity"
// @Success 200
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/presence/{id} [delete]
func (controller *controller) UnregisterPresence(ctx *gin.Context) {
	id := ctx.Param("id")

	fmt.Printf("unregister presence for: %+v\n", id)
}

// ChangeReputation ...
// @Summary Change reputation of a node
// @Description Changes reputation of given node to some specified value
// @ID changeReputation
// @Accept  json
// @Produce  json
// @Tags presence
// @Param id path string true "Node Identity"
// @Param reputation query integer true "New Reputation"
// @Success 200
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/presence/reputation/{id} [patch]
func (controller *controller) ChangeReputation(ctx *gin.Context) {
	id := ctx.Param("id")
	newRep := ctx.Request.URL.Query().Get("reputation")


	fmt.Printf("unregister presence for: %+v to %+v\n", id, newRep)
}

// GetTopology ...
// @Summary Lists which Nym mixnodes and gateways on the network alongside their reputation.
// @Description On Nym nodes startup they register their presence indicating they should be alive. This method provides a list of nodes which have done so.
// @ID getTopology
// @Produce  json
// @Tags presence
// @Success 200 {object} models.Topology
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/presence/topology [get]
func (controller *controller) GetTopology(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, controller.service.GetTopology())
}