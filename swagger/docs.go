// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://vaaani.live",
        "contact": {
            "name": "API Support",
            "url": "https://www.vaaani.live",
            "email": "support@vaaani.live"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/audio/dub": {
            "post": {
                "description": "Places a dub request on the machine, if the task is successfully created then you will get the uuid, that you can use to poll and get the task status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Create DubRequest"
                ],
                "summary": "Create Dub Request",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/send_video.AudioRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/send_video.Response"
                        }
                    }
                }
            }
        },
        "/poll/:uuid": {
            "get": {
                "description": "get the status of ongoing task request of processing the audio.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Poll Get"
                ],
                "summary": "poll ongoing task status.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uuid",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/TaskPool.TaskStatus"
                        }
                    }
                }
            }
        },
        "/profile/add_money": {
            "post": {
                "description": "Adds a new transaction record to the transaction table, and you will get the final credits of user's account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Create AddMoney"
                ],
                "summary": "Adds a new transaction record to the transaction table",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/add_money.AddReqModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/profile/create": {
            "post": {
                "description": "Adds a new user where the primary key will be the gmail id and profile id will be given in response",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Create Profile"
                ],
                "summary": "create a profile for the user using google auth",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/create_profile.CreateReqModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/profile/deduct_money": {
            "post": {
                "description": "-ve credits are not yet handled",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Delete CutMoneyService"
                ],
                "summary": "deduct x amount of money from user's account",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/TaskPool.DeductRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/profile/txn_history": {
            "get": {
                "description": "lazy loading can be added, to implement that pagination on BE will be required",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get TxnHistory"
                ],
                "summary": "get a list of all the txns that the user did",
                "parameters": [
                    {
                        "type": "string",
                        "description": "rajatn@gmail.com",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "TaskPool.DeductRequest": {
            "type": "object",
            "properties": {
                "cost": {
                    "type": "number"
                },
                "email_id": {
                    "type": "string"
                },
                "euid": {
                    "type": "string"
                },
                "subtitle": {
                    "type": "string"
                },
                "video": {
                    "type": "string"
                }
            }
        },
        "TaskPool.TaskStatus": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "There was an error"
                },
                "links": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "object",
                        "additionalProperties": {
                            "type": "string"
                        }
                    }
                },
                "processing_complete": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "add_money.AddReqModel": {
            "type": "object",
            "properties": {
                "credits": {
                    "type": "integer",
                    "example": 1000
                },
                "email": {
                    "type": "string",
                    "example": "rajatn@gmail.com"
                }
            }
        },
        "create_profile.CreateReqModel": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "rajatn@gmail.com"
                },
                "initial_credits": {
                    "description": "default value",
                    "type": "integer",
                    "example": 2100
                },
                "name": {
                    "type": "string",
                    "example": "Rajat Kumar"
                },
                "phone": {
                    "type": "integer",
                    "example": 8010201921
                }
            }
        },
        "send_video.AudioRequest": {
            "type": "object",
            "required": [
                "audio_length",
                "email",
                "file_link",
                "languages"
            ],
            "properties": {
                "audio_length": {
                    "type": "number",
                    "example": 10
                },
                "email": {
                    "type": "string",
                    "example": "rajatn@gmail.com"
                },
                "file_link": {
                    "type": "string",
                    "example": "https://www.emaple.com/file"
                },
                "languages": {
                    "description": "TODO: Add support for dubbing multiple languages at the same time",
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "hindi|telugu|marathi|bengali|tamil"
                    ]
                }
            }
        },
        "send_video.Response": {
            "type": "object",
            "properties": {
                "euid": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "api.vaaani.live",
	BasePath:         "",
	Schemes:          []string{"http"},
	Title:            "SocketBoys/Backned APIs",
	Description:      "Testing Swagger APIs.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
