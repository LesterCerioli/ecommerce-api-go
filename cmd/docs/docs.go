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
		"/auth/token": {
			"post": {
				"summary": "Generate Token",
				"description": "Generates a new JWT token for API access",
                "tags": ["Auth"],
                "consumes": ["application/json"],
                "produces": ["application/json"],
				"parameters": [
					{
						"in": "body",
                        "name": "credentials",
                        "required": true,
						"schema": {
							"type": "object",
                            "required": ["client_id", "secret"],
							"properties": {
								"client_id": { "type": "string" },
                                "secret": { "type": "string" }
							}
						}
					}
				],
				"responses": {
					"201": {
						"description": "Token generated successfully",
						"schema": {
							"type": "object",
							"proprerties": {
								"token": { "type": "string" },
                                "client_id": { "type": "string" },
                                "expires_at": { "type": "string" }
							}
						}
					},
					"400": { "description": "Invalid request" },
                    "401": { "description": "Invalid credentials" }
				}
			}
		},
		"/auth/token/validate": {
			"post": {
				"summary": "Validate Token",
                "description": "Validates a JWT token",
                "tags": ["Auth"],
                "consumes": ["application/json"],
                "produces": ["application/json"],
				"parameters": [
					{
						"in": "body",
                        "name": "token",
                        "required": true,
						"schema": {
							"type": "object",
							"proprties": {
								"token": { "type": "string" }
							}
						}
					}
				],
				"responses": {
					"200": {
						"description": "Token is valid",
						"schema": {
							"type": "object",
							"properties": {
								"message": { "type": "string" },
                                "status": { "type": "string" }
							}
						}
					},
					"401": { "description": "Invalid token" }
				}

			}
		}
	},
	"securityDefinitions": {
		"BearerAuth": {
			"description": "Type \"Bearer\" followed by a space and the JWT token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
			
		}
	}
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:7000",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "E-commerce API",
	Description:      "API for E-commerce",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
