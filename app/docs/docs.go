// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/2xx": {
            "get": {
                "description": "Get an OK response [status: 200] for testing purposes.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Response from service"
                ],
                "summary": "Get OK response.",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/3xx": {
            "get": {
                "description": "Make redirect [status: 301] to \"/2xx\" for testing purposes.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Response from service"
                ],
                "summary": "Get OK response.",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/4xx": {
            "get": {
                "description": "Get an error response [status: 400] for testing purposes.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Response from service"
                ],
                "summary": "Get BadRequest response.",
                "responses": {
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/5xx": {
            "get": {
                "description": "Get an error response [status: 500] for testing purposes.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Response from service"
                ],
                "summary": "Get InternalServerError response",
                "responses": {
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/v1/liveness": {
            "get": {
                "description": "This is LivenessProbe for K8S.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Service State"
                ],
                "summary": "Return service LivenessProbe.",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/liveness-change": {
            "get": {
                "description": "Change LivenessProbe service flag for check K8S reaction (expect pod restart).",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Service State"
                ],
                "summary": "Change LivenessProbe: Success/Failure.",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/oomkill": {
            "get": {
                "description": "Gradual increase in memory consumption for OOM.",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Service State"
                ],
                "summary": "Increase memory consumption.",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/readness": {
            "get": {
                "description": "This is ReadnessProbe for K8S.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Service State"
                ],
                "summary": "Return service ReadnessProbe.",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/readness-change": {
            "get": {
                "description": "Change ReadnessProbe service flag for check K8S reaction (expect traffic interrupt, check by req logs: \"replica_id\").",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Service State"
                ],
                "summary": "Change ReadnessProbe: Success/Failure.",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/throttling": {
            "get": {
                "description": "Increase cpu consumption for throttling with duration 1 minute.",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Service State"
                ],
                "summary": "Increase cpu consumption.",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
