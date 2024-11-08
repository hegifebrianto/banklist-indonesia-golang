// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://www.terbitdigitalagency.com",
            "email": "hegy.febrianto@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/banks": {
            "get": {
                "description": "banks",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "application/json"
                ],
                "tags": [
                    "banks",
                    "banks"
                ],
                "summary": "banks",
                "responses": {
                    "200": {
                        "description": "List of Banks with total count",
                        "schema": {
                            "$ref": "#/definitions/model.BanksDataResponse"
                        }
                    },
                    "206": {
                        "description": "Error lain",
                        "schema": {
                            "$ref": "#/definitions/model.FailResponse"
                        }
                    },
                    "207": {
                        "description": "Gagal Login",
                        "schema": {
                            "$ref": "#/definitions/model.FailResponse"
                        }
                    },
                    "209": {
                        "description": "Invalid JSON input",
                        "schema": {
                            "$ref": "#/definitions/model.FailResponse"
                        }
                    },
                    "210": {
                        "description": "Error Middleware Internal",
                        "schema": {
                            "$ref": "#/definitions/model.FailResponse"
                        }
                    },
                    "214": {
                        "description": "Data Pengguna Tidak Ditemukan",
                        "schema": {
                            "$ref": "#/definitions/model.FailResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Bank": {
            "type": "object",
            "properties": {
                "bankname": {
                    "type": "string",
                    "example": "BRI"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "model.BanksDataResponse": {
            "type": "object",
            "properties": {
                "banks": {
                    "description": "List of provinces",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Bank"
                    }
                },
                "total_count": {
                    "description": "Total number of provinces",
                    "type": "integer"
                }
            }
        },
        "model.FailResponse": {
            "type": "object",
            "required": [
                "msg"
            ],
            "properties": {
                "msg": {
                    "type": "string",
                    "example": "Pesan kesalahan"
                }
            }
        }
    },
    "tags": [
        {
            "description": "API yang digunakan untuk mendapatkan data banks",
            "name": "banks"
        }
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "API for the get of Data sub district & district in Indonesia",
	Description:      "This is a sample API documentation.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
