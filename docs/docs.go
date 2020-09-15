// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-09-15 23:15:10.538651 +0700 WIB m=+0.155404925

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
            "name": "Nuryanto",
            "url": "https://www.linkedin.com/in/nuryanto-1b2721156/",
            "email": "nuryantofattih@gmail.com"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api.v1/capster/auth/change_password": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Change Password",
                "parameters": [
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "OS",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "Version",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "req param #changes are possible to adjust the form of the registration form from frontend",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ResetPasswd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tool.ResponseModel"
                        }
                    }
                }
            }
        },
        "/api.v1/capster/auth/forgot": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Forgot Password",
                "parameters": [
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "OS",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "Version",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "req param #changes are possible to adjust the form of the registration form from frontend",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.ForgotForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tool.ResponseModel"
                        }
                    }
                }
            }
        },
        "/api.v1/capster/auth/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "OS",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "Version",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "req param #changes are possible to adjust the form of the registration form from frontend",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.LoginForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tool.ResponseModel"
                        }
                    }
                }
            }
        },
        "/api.v1/capster/auth/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "OS",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "Version",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "req param #changes are possible to adjust the form of the registration form from frontend",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.RegisterForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tool.ResponseModel"
                        }
                    }
                }
            }
        },
        "/api.v1/capster/auth/verify": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Verify",
                "parameters": [
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "OS",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "Version",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "req param #changes are possible to adjust the form of the registration form from frontend",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.VerifyForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tool.ResponseModel"
                        }
                    }
                }
            }
        },
        "/api.v1/capster/order": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "GetList Order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "OS",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "Version",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "PerPage",
                        "name": "perpage",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Search",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "InitSearch",
                        "name": "initsearch",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "SortField",
                        "name": "sortfield",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseModelList"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Add Order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "OS",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "Version",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "req param #changes are possible to adjust the form of the registration form from frontend",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.OrderPost"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tool.ResponseModel"
                        }
                    }
                }
            }
        },
        "/api.v1/capster/order/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "GetById",
                "parameters": [
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "OS",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "Version",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tool.ResponseModel"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Rubah Profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "OS",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "Version",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "req param #changes are possible to adjust the form of the registration form from frontend",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.OrderPost"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tool.ResponseModel"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Delete Order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "OS",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "Version",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tool.ResponseModel"
                        }
                    }
                }
            }
        },
        "/api.v1/capster/paket": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Paket"
                ],
                "summary": "GetList Paket",
                "parameters": [
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "OS",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "Version",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "PerPage",
                        "name": "perpage",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Search",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "InitSearch",
                        "name": "initsearch",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "SortField",
                        "name": "sortfield",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseModelList"
                        }
                    }
                }
            }
        },
        "/api.v1/capster/paket/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Paket"
                ],
                "summary": "GetById",
                "parameters": [
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "OS",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "Version",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tool.ResponseModel"
                        }
                    }
                }
            }
        },
        "/api.v1/capster/user": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "GetList User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "OS",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "Version",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "PerPage",
                        "name": "perpage",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Search",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "InitSearch",
                        "name": "initsearch",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "SortField",
                        "name": "sortfield",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseModelList"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Add User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "OS",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "Version",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "req param #changes are possible to adjust the form of the registration form from frontend",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.AddUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tool.ResponseModel"
                        }
                    }
                }
            }
        },
        "/api.v1/capster/user/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "GetById",
                "parameters": [
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "OS",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "Version",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tool.ResponseModel"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Rubah Profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "OS",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "Version",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "req param #changes are possible to adjust the form of the registration form from frontend",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.UpdateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tool.ResponseModel"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "OS",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "Version",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tool.ResponseModel"
                        }
                    }
                }
            }
        },
        "/api.v1/fileupload": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Upload file",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "FileUpload"
                ],
                "summary": "File Upload",
                "parameters": [
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "OS",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OS Device",
                        "name": "Version",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "account image",
                        "name": "upload_file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "path images",
                        "name": "path",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tool.ResponseModel"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AddUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "is_admin": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "telp": {
                    "type": "string"
                }
            }
        },
        "models.ForgotForm": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                }
            }
        },
        "models.LoginForm": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "pwd": {
                    "type": "string"
                }
            }
        },
        "models.OrderDPost": {
            "type": "object",
            "properties": {
                "durasi_end": {
                    "type": "integer"
                },
                "durasi_start": {
                    "type": "integer"
                },
                "paket_id": {
                    "type": "integer"
                },
                "paket_name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "models.OrderPost": {
            "type": "object",
            "properties": {
                "customer_name": {
                    "type": "string"
                },
                "order_date": {
                    "description": "BarberID     int          ` + "`" + `json:\"barber_id\" valid:\"Required\"` + "`" + `\nCapsterID    int          ` + "`" + `json:\"capster_id,omitempty\"` + "`" + `",
                    "type": "string"
                },
                "paket_ids": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.OrderDPost"
                    }
                },
                "telp": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.RegisterForm": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "confirm_pwd": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "pwd": {
                    "type": "string"
                },
                "user_type": {
                    "type": "string"
                }
            }
        },
        "models.ResetPasswd": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "confirm_pwd": {
                    "type": "string"
                },
                "pwd": {
                    "type": "string"
                }
            }
        },
        "models.ResponseModelList": {
            "type": "object",
            "properties": {
                "all_column": {
                    "type": "string"
                },
                "data": {
                    "type": "object"
                },
                "define_column": {
                    "type": "string"
                },
                "define_size": {
                    "type": "string"
                },
                "last_page": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "page": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "models.UpdateUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "telp": {
                    "type": "string"
                }
            }
        },
        "models.VerifyForm": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "verify_code": {
                    "type": "string"
                }
            }
        },
        "tool.ResponseModel": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "message": {
                    "description": "Code int         ` + "`" + `json:\"code\"` + "`" + `",
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Starter",
	Description: "Backend REST API for golang nuryanto2121",
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
