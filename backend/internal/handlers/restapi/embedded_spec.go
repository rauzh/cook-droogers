// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Swagger-Cook-Droogers",
    "version": "1.0.0"
  },
  "host": "0.0.0.0:13337",
  "basePath": "/api",
  "paths": {
    "/fetch-stats": {
      "post": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "tags": [
          "manager"
        ],
        "summary": "Fetch statistics for manager's artists",
        "operationId": "fetchStats",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/heartbeat": {
      "get": {
        "summary": "Check health",
        "responses": {
          "200": {
            "description": "Success"
          }
        }
      }
    },
    "/managers": {
      "get": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "admin"
        ],
        "summary": "Get list of managers",
        "operationId": "getManagers",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/ManagerDTO"
              }
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      },
      "post": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "admin"
        ],
        "summary": "Create manager",
        "operationId": "addManager",
        "parameters": [
          {
            "type": "string",
            "description": "ID пользователя",
            "name": "user_id",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "Manager successfully created",
            "schema": {
              "$ref": "#/definitions/ManagerDTO"
            }
          },
          "403": {
            "description": "Manager alredy exists"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/publish": {
      "post": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "tags": [
          "artist"
        ],
        "summary": "Create publish request",
        "operationId": "publishReq",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "description": "ID релиза",
            "name": "release_id",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "format": "date",
            "description": "Желаемая дата публикации",
            "name": "date",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "Request successfully created"
          },
          "400": {
            "description": "Invalid request fields"
          },
          "403": {
            "description": "Invalid user type"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/register": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "tags": [
          "guest"
        ],
        "summary": "Create new user",
        "operationId": "register",
        "parameters": [
          {
            "type": "string",
            "description": "Имя пользователя",
            "name": "username",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "format": "email",
            "description": "Email пользователя",
            "name": "email",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "Пароль пользователя",
            "name": "password",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "User successfully created"
          },
          "403": {
            "description": "User alredy exists"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/releases": {
      "get": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "tags": [
          "artist"
        ],
        "summary": "Get releases",
        "operationId": "getRelease",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/ReleaseDTO"
              }
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      },
      "post": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "tags": [
          "artist"
        ],
        "summary": "Upload release",
        "operationId": "addRelease",
        "parameters": [
          {
            "type": "string",
            "description": "Название релиза",
            "name": "title",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "format": "date",
            "description": "Дата написания релиза",
            "name": "date",
            "in": "query",
            "required": true
          },
          {
            "description": "Треки данного релиза",
            "name": "tracks",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/TrackDTO"
              }
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Request successfully created"
          },
          "400": {
            "description": "Invalid request fields"
          },
          "403": {
            "description": "Invalid user type"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/requests": {
      "get": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "tags": [
          "non-member",
          "manager",
          "artist"
        ],
        "summary": "Get requests",
        "operationId": "getRequests",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/RequestDTO"
              }
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/requests/{req_id}": {
      "get": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "tags": [
          "non-member",
          "manager",
          "artist"
        ],
        "summary": "Get specified request",
        "operationId": "getRequest",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "description": "ID заявки",
            "name": "req_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/PublishRequestDTO"
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/requests/{req_id}/accept": {
      "post": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "tags": [
          "manager"
        ],
        "summary": "Accept specified request",
        "operationId": "acceptRequest",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "description": "ID заявки",
            "name": "req_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/requests/{req_id}/decline": {
      "post": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "tags": [
          "manager"
        ],
        "summary": "Decline specified request",
        "operationId": "declineRequest",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "description": "ID заявки",
            "name": "req_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/sign-contract": {
      "post": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "tags": [
          "non-member"
        ],
        "summary": "Create sign request",
        "operationId": "signContract",
        "parameters": [
          {
            "type": "string",
            "description": "Псевдоним",
            "name": "nickname",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "Request successfully created"
          },
          "400": {
            "description": "Invalid request fields"
          },
          "403": {
            "description": "Invalid user type"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/stats": {
      "get": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "tags": [
          "artist",
          "manager"
        ],
        "summary": "Get stats",
        "operationId": "getStats",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/users": {
      "get": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "admin"
        ],
        "summary": "Get list of users",
        "operationId": "getUsers",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/UserDTO"
              }
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    }
  },
  "definitions": {
    "Category": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "ManagerDTO": {
      "type": "object",
      "properties": {
        "artists": {
          "type": "array",
          "items": {
            "type": "integer",
            "format": "uint64"
          }
        },
        "manager_id": {
          "type": "integer",
          "format": "uint64"
        },
        "user_id": {
          "type": "integer",
          "format": "uint64"
        }
      }
    },
    "Pet": {
      "type": "object",
      "required": [
        "name",
        "photoUrls"
      ],
      "properties": {
        "category": {
          "$ref": "#/definitions/Category"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string",
          "example": "doggie"
        },
        "photoUrls": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "status": {
          "description": "pet status in the store",
          "type": "string",
          "enum": [
            "available",
            "pending",
            "sold"
          ]
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Tag"
          }
        }
      }
    },
    "PublishRequestDTO": {
      "type": "object",
      "properties": {
        "base": {
          "$ref": "#/definitions/RequestDTO"
        },
        "description": {
          "type": "string"
        },
        "expected_date": {
          "type": "string",
          "format": "date"
        },
        "grade": {
          "type": "integer"
        },
        "release_id": {
          "type": "integer",
          "format": "uint64"
        }
      }
    },
    "RegUserDTO": {
      "type": "object",
      "required": [
        "name",
        "email",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string",
          "format": "email"
        },
        "name": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "ReleaseDTO": {
      "type": "object",
      "properties": {
        "artist_id": {
          "type": "integer",
          "format": "uint64"
        },
        "date_creation": {
          "type": "string",
          "format": "date"
        },
        "release_id": {
          "type": "integer",
          "format": "uint64"
        },
        "status": {
          "type": "string",
          "enum": [
            "Unpublished",
            "Published"
          ]
        },
        "title": {
          "type": "string"
        },
        "tracks": {
          "type": "array",
          "items": {
            "type": "integer",
            "format": "uint64"
          }
        }
      }
    },
    "RequestDTO": {
      "type": "object",
      "properties": {
        "applier_id": {
          "type": "integer",
          "format": "uint64"
        },
        "date": {
          "type": "string",
          "format": "date"
        },
        "manager_id": {
          "type": "integer",
          "format": "uint64"
        },
        "request_id": {
          "type": "integer",
          "format": "uint64"
        },
        "status": {
          "type": "string",
          "enum": [
            "New",
            "Processing",
            "On approval",
            "Closed"
          ]
        },
        "type": {
          "type": "string",
          "enum": [
            "Publish",
            "Sign"
          ]
        }
      }
    },
    "SignRequestDTO": {
      "type": "object",
      "properties": {
        "base": {
          "$ref": "#/definitions/RequestDTO"
        },
        "description": {
          "type": "string"
        },
        "nickname": {
          "type": "string"
        }
      }
    },
    "StatsDTO": {
      "type": "object",
      "properties": {
        "date": {
          "type": "string",
          "format": "date"
        },
        "likes": {
          "type": "integer",
          "format": "uint64"
        },
        "stat_id": {
          "type": "integer",
          "format": "uint64"
        },
        "streams": {
          "type": "integer",
          "format": "uint64"
        },
        "track_id": {
          "type": "integer",
          "format": "uint64"
        }
      }
    },
    "Tag": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "TrackDTO": {
      "type": "object",
      "required": [
        "title",
        "duration",
        "genre",
        "type"
      ],
      "properties": {
        "artists": {
          "type": "array",
          "items": {
            "type": "integer",
            "format": "uint64"
          }
        },
        "duration": {
          "type": "integer",
          "format": "uint64"
        },
        "genre": {
          "type": "string"
        },
        "title": {
          "type": "string"
        },
        "track_id": {
          "type": "integer",
          "format": "uint64"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "UserDTO": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string",
          "format": "email"
        },
        "name": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "type": {
          "type": "integer",
          "enum": [
            0,
            1,
            2,
            3
          ]
        },
        "user_id": {
          "type": "integer",
          "format": "uint64"
        }
      }
    }
  },
  "responses": {
    "UnauthorizedError": {
      "description": "Authentication information is missing or invalid",
      "headers": {
        "WWW_Authenticate": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "basicAuth": {
      "type": "basic"
    }
  },
  "tags": [
    {
      "name": "artist"
    },
    {
      "name": "manager"
    },
    {
      "name": "non-member"
    },
    {
      "name": "guest"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Swagger-Cook-Droogers",
    "version": "1.0.0"
  },
  "host": "0.0.0.0:13337",
  "basePath": "/api",
  "paths": {
    "/fetch-stats": {
      "post": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "tags": [
          "manager"
        ],
        "summary": "Fetch statistics for manager's artists",
        "operationId": "fetchStats",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/heartbeat": {
      "get": {
        "summary": "Check health",
        "responses": {
          "200": {
            "description": "Success"
          }
        }
      }
    },
    "/managers": {
      "get": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "admin"
        ],
        "summary": "Get list of managers",
        "operationId": "getManagers",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/ManagerDTO"
              }
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      },
      "post": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "admin"
        ],
        "summary": "Create manager",
        "operationId": "addManager",
        "parameters": [
          {
            "type": "string",
            "description": "ID пользователя",
            "name": "user_id",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "Manager successfully created",
            "schema": {
              "$ref": "#/definitions/ManagerDTO"
            }
          },
          "403": {
            "description": "Manager alredy exists"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/publish": {
      "post": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "tags": [
          "artist"
        ],
        "summary": "Create publish request",
        "operationId": "publishReq",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "description": "ID релиза",
            "name": "release_id",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "format": "date",
            "description": "Желаемая дата публикации",
            "name": "date",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "Request successfully created"
          },
          "400": {
            "description": "Invalid request fields"
          },
          "403": {
            "description": "Invalid user type"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/register": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "tags": [
          "guest"
        ],
        "summary": "Create new user",
        "operationId": "register",
        "parameters": [
          {
            "type": "string",
            "description": "Имя пользователя",
            "name": "username",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "format": "email",
            "description": "Email пользователя",
            "name": "email",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "Пароль пользователя",
            "name": "password",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "User successfully created"
          },
          "403": {
            "description": "User alredy exists"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/releases": {
      "get": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "tags": [
          "artist"
        ],
        "summary": "Get releases",
        "operationId": "getRelease",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/ReleaseDTO"
              }
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      },
      "post": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "tags": [
          "artist"
        ],
        "summary": "Upload release",
        "operationId": "addRelease",
        "parameters": [
          {
            "type": "string",
            "description": "Название релиза",
            "name": "title",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "format": "date",
            "description": "Дата написания релиза",
            "name": "date",
            "in": "query",
            "required": true
          },
          {
            "description": "Треки данного релиза",
            "name": "tracks",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/TrackDTO"
              }
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Request successfully created"
          },
          "400": {
            "description": "Invalid request fields"
          },
          "403": {
            "description": "Invalid user type"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/requests": {
      "get": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "tags": [
          "non-member",
          "manager",
          "artist"
        ],
        "summary": "Get requests",
        "operationId": "getRequests",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/RequestDTO"
              }
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/requests/{req_id}": {
      "get": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "tags": [
          "non-member",
          "manager",
          "artist"
        ],
        "summary": "Get specified request",
        "operationId": "getRequest",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "description": "ID заявки",
            "name": "req_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/PublishRequestDTO"
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/requests/{req_id}/accept": {
      "post": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "tags": [
          "manager"
        ],
        "summary": "Accept specified request",
        "operationId": "acceptRequest",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "description": "ID заявки",
            "name": "req_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/requests/{req_id}/decline": {
      "post": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "tags": [
          "manager"
        ],
        "summary": "Decline specified request",
        "operationId": "declineRequest",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "description": "ID заявки",
            "name": "req_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/sign-contract": {
      "post": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "tags": [
          "non-member"
        ],
        "summary": "Create sign request",
        "operationId": "signContract",
        "parameters": [
          {
            "type": "string",
            "description": "Псевдоним",
            "name": "nickname",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "Request successfully created"
          },
          "400": {
            "description": "Invalid request fields"
          },
          "403": {
            "description": "Invalid user type"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/stats": {
      "get": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "tags": [
          "artist",
          "manager"
        ],
        "summary": "Get stats",
        "operationId": "getStats",
        "responses": {
          "200": {
            "description": "Success"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/users": {
      "get": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "admin"
        ],
        "summary": "Get list of users",
        "operationId": "getUsers",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/UserDTO"
              }
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    }
  },
  "definitions": {
    "Category": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "ManagerDTO": {
      "type": "object",
      "properties": {
        "artists": {
          "type": "array",
          "items": {
            "type": "integer",
            "format": "uint64"
          }
        },
        "manager_id": {
          "type": "integer",
          "format": "uint64"
        },
        "user_id": {
          "type": "integer",
          "format": "uint64"
        }
      }
    },
    "Pet": {
      "type": "object",
      "required": [
        "name",
        "photoUrls"
      ],
      "properties": {
        "category": {
          "$ref": "#/definitions/Category"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string",
          "example": "doggie"
        },
        "photoUrls": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "status": {
          "description": "pet status in the store",
          "type": "string",
          "enum": [
            "available",
            "pending",
            "sold"
          ]
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Tag"
          }
        }
      }
    },
    "PublishRequestDTO": {
      "type": "object",
      "properties": {
        "base": {
          "$ref": "#/definitions/RequestDTO"
        },
        "description": {
          "type": "string"
        },
        "expected_date": {
          "type": "string",
          "format": "date"
        },
        "grade": {
          "type": "integer"
        },
        "release_id": {
          "type": "integer",
          "format": "uint64"
        }
      }
    },
    "RegUserDTO": {
      "type": "object",
      "required": [
        "name",
        "email",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string",
          "format": "email"
        },
        "name": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "ReleaseDTO": {
      "type": "object",
      "properties": {
        "artist_id": {
          "type": "integer",
          "format": "uint64"
        },
        "date_creation": {
          "type": "string",
          "format": "date"
        },
        "release_id": {
          "type": "integer",
          "format": "uint64"
        },
        "status": {
          "type": "string",
          "enum": [
            "Unpublished",
            "Published"
          ]
        },
        "title": {
          "type": "string"
        },
        "tracks": {
          "type": "array",
          "items": {
            "type": "integer",
            "format": "uint64"
          }
        }
      }
    },
    "RequestDTO": {
      "type": "object",
      "properties": {
        "applier_id": {
          "type": "integer",
          "format": "uint64"
        },
        "date": {
          "type": "string",
          "format": "date"
        },
        "manager_id": {
          "type": "integer",
          "format": "uint64"
        },
        "request_id": {
          "type": "integer",
          "format": "uint64"
        },
        "status": {
          "type": "string",
          "enum": [
            "New",
            "Processing",
            "On approval",
            "Closed"
          ]
        },
        "type": {
          "type": "string",
          "enum": [
            "Publish",
            "Sign"
          ]
        }
      }
    },
    "SignRequestDTO": {
      "type": "object",
      "properties": {
        "base": {
          "$ref": "#/definitions/RequestDTO"
        },
        "description": {
          "type": "string"
        },
        "nickname": {
          "type": "string"
        }
      }
    },
    "StatsDTO": {
      "type": "object",
      "properties": {
        "date": {
          "type": "string",
          "format": "date"
        },
        "likes": {
          "type": "integer",
          "format": "uint64"
        },
        "stat_id": {
          "type": "integer",
          "format": "uint64"
        },
        "streams": {
          "type": "integer",
          "format": "uint64"
        },
        "track_id": {
          "type": "integer",
          "format": "uint64"
        }
      }
    },
    "Tag": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "TrackDTO": {
      "type": "object",
      "required": [
        "title",
        "duration",
        "genre",
        "type"
      ],
      "properties": {
        "artists": {
          "type": "array",
          "items": {
            "type": "integer",
            "format": "uint64"
          }
        },
        "duration": {
          "type": "integer",
          "format": "uint64"
        },
        "genre": {
          "type": "string"
        },
        "title": {
          "type": "string"
        },
        "track_id": {
          "type": "integer",
          "format": "uint64"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "UserDTO": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string",
          "format": "email"
        },
        "name": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "type": {
          "type": "integer",
          "enum": [
            0,
            1,
            2,
            3
          ]
        },
        "user_id": {
          "type": "integer",
          "format": "uint64"
        }
      }
    }
  },
  "responses": {
    "UnauthorizedError": {
      "description": "Authentication information is missing or invalid",
      "headers": {
        "WWW_Authenticate": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "basicAuth": {
      "type": "basic"
    }
  },
  "tags": [
    {
      "name": "artist"
    },
    {
      "name": "manager"
    },
    {
      "name": "non-member"
    },
    {
      "name": "guest"
    }
  ]
}`))
}
