// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-10-20 12:56:20.2726975 +0100 BST m=+0.174689501

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "https://github.com/nymtech/nym-validator/license"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/healthcheck": {
            "get": {
                "description": "Returns a 200 if the directory server is available. Good route to use for automated monitoring.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "healthcheck"
                ],
                "summary": "Lets the directory server tell the world it's alive.",
                "operationId": "healthCheck",
                "responses": {
                    "200": {}
                }
            }
        },
        "/api/mixmining": {
            "post": {
                "description": "Nym network monitor sends packets through the system and checks if they make it. The network monitor then hits this method to report whether the node was up at a given time.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mixmining"
                ],
                "summary": "Lets the network monitor create a new uptime status for a mix",
                "operationId": "addMixStatus",
                "parameters": [
                    {
                        "description": "object",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.MixStatus"
                        }
                    }
                ],
                "responses": {
                    "201": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/mixmining/batch": {
            "post": {
                "description": "Nym network monitor sends packets through the system and checks if they make it. The network monitor then hits this method to report whether nodes were up at a given time.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mixmining"
                ],
                "summary": "Lets the network monitor create a new uptime status for multiple mixes",
                "operationId": "batchCreateMixStatus",
                "parameters": [
                    {
                        "description": "object",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BatchMixStatus"
                        }
                    }
                ],
                "responses": {
                    "201": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/mixmining/fullreport": {
            "get": {
                "description": "Provides summary uptime statistics for last 5 minutes, day, week, and month",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mixmining"
                ],
                "summary": "Retrieves a summary report of historical mix status",
                "operationId": "batchGetMixStatusReport",
                "responses": {
                    "200": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/mixmining/node/{pubkey}/history": {
            "get": {
                "description": "Lists all mixnode statuses for a given node pubkey",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mixmining"
                ],
                "summary": "Lists mixnode activity",
                "operationId": "listMixStatuses",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Mixnode Pubkey",
                        "name": "pubkey",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.MixStatus"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/mixmining/node/{pubkey}/report": {
            "get": {
                "description": "Provides summary uptime statistics for last 5 minutes, day, week, and month",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mixmining"
                ],
                "summary": "Retrieves a summary report of historical mix status",
                "operationId": "getMixStatusReport",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Mixnode Pubkey",
                        "name": "pubkey",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/presence/gateway": {
            "post": {
                "description": "On Nym nodes startup they register their presence indicating they should be alive and get added to the set of active nodes in the topology.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "presence"
                ],
                "summary": "Lets a gateway tell the directory server it's coming online",
                "operationId": "registerGatewayPresence",
                "parameters": [
                    {
                        "description": "object",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.GatewayRegistrationInfo"
                        }
                    }
                ],
                "responses": {
                    "201": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/presence/mix": {
            "post": {
                "description": "On Nym nodes startup they register their presence indicating they should be alive and get added to the set of active nodes in the topology.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "presence"
                ],
                "summary": "Lets a mixnode tell the directory server it's coming online",
                "operationId": "registerMixPresence",
                "parameters": [
                    {
                        "description": "object",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.MixRegistrationInfo"
                        }
                    }
                ],
                "responses": {
                    "201": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/presence/reputation/{id}": {
            "patch": {
                "description": "Changes reputation of given node to some specified value",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "presence"
                ],
                "summary": "Change reputation of a node",
                "operationId": "changeReputation",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Node Identity",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "New Reputation",
                        "name": "reputation",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/presence/topology": {
            "get": {
                "description": "On Nym nodes startup they register their presence indicating they should be alive. This method provides a list of nodes which have done so.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "presence"
                ],
                "summary": "Lists which Nym mixnodes and gateways on the network alongside their reputation.",
                "operationId": "getTopology",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Topology"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/api/presence/{id}": {
            "delete": {
                "description": "Messages sent by a node on powering down to indicate it's going offline so that it should get removed from active topology.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "presence"
                ],
                "summary": "Unregister presence of node.",
                "operationId": "unregisterPresence",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Node Identity",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.BatchMixStatus": {
            "type": "object",
            "required": [
                "status"
            ],
            "properties": {
                "status": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.MixStatus"
                    }
                }
            }
        },
        "models.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "models.GatewayRegistrationInfo": {
            "type": "object",
            "required": [
                "clientsHost",
                "identityKey",
                "mixHost",
                "sphinxKey",
                "version"
            ],
            "properties": {
                "clientsHost": {
                    "type": "string"
                },
                "identityKey": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "mixHost": {
                    "type": "string"
                },
                "sphinxKey": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "models.MixRegistrationInfo": {
            "type": "object",
            "required": [
                "identityKey",
                "layer",
                "mixHost",
                "sphinxKey",
                "version"
            ],
            "properties": {
                "identityKey": {
                    "type": "string"
                },
                "layer": {
                    "type": "integer"
                },
                "location": {
                    "type": "string"
                },
                "mixHost": {
                    "type": "string"
                },
                "sphinxKey": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "models.MixStatus": {
            "type": "object",
            "required": [
                "ipVersion",
                "pubKey",
                "up"
            ],
            "properties": {
                "ipVersion": {
                    "type": "string"
                },
                "pubKey": {
                    "type": "string"
                },
                "up": {
                    "type": "boolean"
                }
            }
        },
        "models.RegisteredGateway": {
            "type": "object",
            "required": [
                "clientsHost",
                "identityKey",
                "mixHost",
                "sphinxKey",
                "version"
            ],
            "properties": {
                "clientsHost": {
                    "type": "string"
                },
                "identityKey": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "mixHost": {
                    "type": "string"
                },
                "registrationTime": {
                    "type": "integer"
                },
                "reputation": {
                    "type": "integer"
                },
                "sphinxKey": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "models.RegisteredMix": {
            "type": "object",
            "required": [
                "identityKey",
                "layer",
                "mixHost",
                "sphinxKey",
                "version"
            ],
            "properties": {
                "identityKey": {
                    "type": "string"
                },
                "layer": {
                    "type": "integer"
                },
                "location": {
                    "type": "string"
                },
                "mixHost": {
                    "type": "string"
                },
                "registrationTime": {
                    "type": "integer"
                },
                "reputation": {
                    "type": "integer"
                },
                "sphinxKey": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "models.Topology": {
            "type": "object",
            "required": [
                "gateways",
                "mixNodes"
            ],
            "properties": {
                "gateways": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.RegisteredGateway"
                    }
                },
                "mixNodes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.RegisteredMix"
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.9.0-dev",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Nym Directory API",
	Description: "A directory API allowing Nym nodes and clients to connect to each other.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
