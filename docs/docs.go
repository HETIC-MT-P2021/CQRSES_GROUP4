// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

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
        "contact": {
            "name": "HETIC-MT5"
        },
        "license": {
            "name": "MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/articles": {
            "get": {
                "description": "Get an array of article struct",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articles"
                ],
                "summary": "Get all articles from elastic search",
                "responses": {
                    "200": {
                        "description": "GET /articles",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articles"
                ],
                "summary": "Create article in elastic search",
                "parameters": [
                    {
                        "description": "Add article",
                        "name": "article",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/article.Request"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "created",
                        "schema": {
                            "$ref": "#/definitions/pkg.HTTPStatus"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.HTTPError"
                        }
                    }
                }
            }
        },
        "/articles/{aggregate_article_id}": {
            "get": {
                "description": "Get article struct",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articles"
                ],
                "summary": "Get an article from elastic search",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Article ID",
                        "name": "aggregate_article_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/database.Article"
                        }
                    },
                    "404": {
                        "description": "Article Not found",
                        "schema": {
                            "$ref": "#/definitions/pkg.HTTPError"
                        }
                    }
                }
            }
        },
        "/articles/{aggregate_article_id}}": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articles"
                ],
                "summary": "Update article in elastic search",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Article ID",
                        "name": "aggregate_article_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update article",
                        "name": "article",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/article.Request"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "updated",
                        "schema": {
                            "$ref": "#/definitions/pkg.HTTPStatus"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.HTTPError"
                        }
                    }
                }
            }
        },
        "/fixtures/event-store": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "fixtures"
                ],
                "summary": "Create event in elastic search",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/pkg.HTTPStatus"
                        }
                    },
                    "500": {
                        "description": "Not Created",
                        "schema": {
                            "$ref": "#/definitions/pkg.HTTPStatus"
                        }
                    }
                }
            }
        },
        "/fixtures/read-model": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "fixtures"
                ],
                "summary": "Create read-model in elastic search",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/pkg.HTTPStatus"
                        }
                    },
                    "500": {
                        "description": "Not Created",
                        "schema": {
                            "$ref": "#/definitions/pkg.HTTPStatus"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Using JWT auth (look headers for token)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Connect user to app",
                "parameters": [
                    {
                        "description": "Account to login",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.requestLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Empty",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.HTTPError"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Using JWT auth",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Create new account",
                "parameters": [
                    {
                        "description": "Add account",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.requestRegister"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Status",
                        "schema": {
                            "$ref": "#/definitions/pkg.HTTPStatus"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "article.Request": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "auth.requestLogin": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "auth.requestRegister": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "database.Article": {
            "type": "object",
            "properties": {
                "aggregate_article_id": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "pkg.HTTPError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "pkg.HTTPStatus": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        }
    },
    "x-extension-openapi": {
        "example": "value on a json format"
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
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "CQRS Event sourcing  documentation API",
	Description: "This is the api documentation for the CQRS Event sourcing implementation for HETIC",
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
