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
        "/create-user": {
            "post": {
                "description": "Create a Sub-User/Member User Account",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Accounts"
                ],
                "summary": "Create a Sub-User/Member User Account",
                "operationId": "create-user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email of the new user",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password of the new user",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Optional Role ID string for new user. Defaults to 'member' if not specified",
                        "name": "roleID",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Optional true or false value to set new user 2FA preference",
                        "name": "has2FA",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.loginSampleResponseError500"
                        }
                    }
                }
            }
        },
        "/images/profile-pic/{imageName}": {
            "get": {
                "description": "Get User Profile Picture",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Profiles"
                ],
                "summary": "Get User Profile Picture",
                "operationId": "get-profile-pic",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Picture Filename",
                        "name": "imageName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Login to FamTrust (Supports 2FA by Email)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Authentication"
                ],
                "summary": "Login to FamTrust (Supports 2FA by Email)",
                "operationId": "login",
                "parameters": [
                    {
                        "description": "User Credentials",
                        "name": "Credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.loginRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "User 2FA Code",
                        "name": "2FACode",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.loginSampleResponse200"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handlers.loginSampleResponseError401"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.loginSampleResponseError500"
                        }
                    }
                }
            }
        },
        "/profile": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve User Profile Details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Profiles"
                ],
                "summary": "Retrieve User Profile Details",
                "operationId": "profile",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.profileSampleResponse200"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handlers.loginSampleResponseError401"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.loginSampleResponseError500"
                        }
                    }
                }
            }
        },
        "/profile/create": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create New User Profile",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Profiles"
                ],
                "summary": "Create New User Profile",
                "operationId": "create-profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User's first name",
                        "name": "firstName",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User's last name",
                        "name": "lastName",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User's biography",
                        "name": "bio",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "User's National Identification Number",
                        "name": "nin",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "User's Bank Verification Number",
                        "name": "bvn",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "User's default group ID",
                        "name": "defaultGroupID",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "User's profile picture",
                        "name": "profilePicture",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/profile/update": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update User Profile",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Profiles"
                ],
                "summary": "Update User Profile",
                "operationId": "update-profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User's first name",
                        "name": "firstName",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "User's last name",
                        "name": "lastName",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "User's biography",
                        "name": "bio",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "User's National Identification Number",
                        "name": "nin",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "User's Bank Verification Number",
                        "name": "bvn",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "User's default group ID",
                        "name": "defaultGroupID",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "User's profile picture",
                        "name": "profilePicture",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/signup": {
            "post": {
                "description": "Create an Admin/Main User Account",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Accounts"
                ],
                "summary": "Create an Admin/Main User Account",
                "operationId": "signup",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email of the new user",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password of the new user",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Optional true or false value to set new user 2FA preference",
                        "name": "has2FA",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.loginSampleResponseError500"
                        }
                    }
                }
            }
        },
        "/users/by/default-group": {
            "get": {
                "description": "Get Users by Group ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Retrieve-Users"
                ],
                "summary": "Get Users by Group ID",
                "operationId": "users-by-groupID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Group ID to filter Users By",
                        "name": "groupID",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.loginSampleResponseError500"
                        }
                    }
                }
            }
        },
        "/validate": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Validate User Login Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Authentication"
                ],
                "summary": "Validate User Login Token",
                "operationId": "validate",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.validateSampleResponse200"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handlers.loginSampleResponseError401"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.loginSampleResponseError500"
                        }
                    }
                }
            }
        },
        "/verify-bvn": {
            "get": {
                "description": "Verify User Signup BVN - Currently this is any 10 digit positive integer not currently in use by another user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Verifications"
                ],
                "summary": "Verify User Signup BVN",
                "operationId": "verify-bvn",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "BVN",
                        "name": "bvn",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/verify-email": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Send User-Email Verification Token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Verifications"
                ],
                "summary": "Send User-Email Verification Token",
                "operationId": "send-verify-token",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/verify-email/verify": {
            "get": {
                "description": "Verify User Email Address via Token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Verifications"
                ],
                "summary": "Verify User Email Address via Token",
                "operationId": "verify-email-token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email verification Token",
                        "name": "code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/verify-nin": {
            "get": {
                "description": "Verify User Signup NIN - Currently this is any 10 digit positive integer not currently in use by another user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Verifications"
                ],
                "summary": "Verify User Signup NIN",
                "operationId": "verify-nin",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "NIN",
                        "name": "nin",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.loginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "handlers.loginSampleResponse200": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "User Logged in successfully"
                },
                "status": {
                    "type": "string",
                    "example": "success"
                },
                "statusCode": {
                    "type": "integer",
                    "example": 200
                },
                "token": {
                    "type": "string",
                    "example": "b6d4a7e1d2d841a1afe874a2a5c15d8b"
                }
            }
        },
        "handlers.loginSampleResponseError401": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Invalid Credentials"
                },
                "status": {
                    "type": "string",
                    "example": "error"
                },
                "statusCode": {
                    "type": "integer",
                    "example": 401
                }
            }
        },
        "handlers.loginSampleResponseError500": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "An error occured"
                },
                "status": {
                    "type": "string",
                    "example": "error"
                },
                "statusCode": {
                    "type": "integer",
                    "example": 500
                }
            }
        },
        "handlers.profileSampleResponse200": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Request successful"
                },
                "profile": {
                    "type": "object",
                    "properties": {
                        "bio": {
                            "type": "string",
                            "example": "The best FamTrust user of all time."
                        },
                        "bvn": {
                            "type": "integer",
                            "example": 35473783473
                        },
                        "createdAt": {
                            "type": "string",
                            "example": "2024-07-22T14:30:00Z"
                        },
                        "firstName": {
                            "type": "string",
                            "example": "Famtrust"
                        },
                        "id": {
                            "type": "string",
                            "example": "a5c9f82e-6b7a-4a53-a81c-82b1e2f453a6"
                        },
                        "lastName": {
                            "type": "string",
                            "example": "Guru"
                        },
                        "nin": {
                            "type": "integer",
                            "example": 35473745433
                        },
                        "profilePictureUrl": {
                            "type": "string",
                            "example": "https://image.famtrust.biz/dkkjieikdjfoej.jpg"
                        },
                        "updatedAt": {
                            "type": "string",
                            "example": "2024-07-22T14:30:00Z"
                        },
                        "userID": {
                            "type": "string",
                            "example": "d38f91b2-dc3b-4f9d-aeb4-7b95c91e9d08"
                        }
                    }
                },
                "status": {
                    "type": "string",
                    "example": "success"
                },
                "statusCode": {
                    "type": "integer",
                    "example": 200
                },
                "token": {
                    "type": "string",
                    "example": "b6d4a7e1d2d841a1afe874a2a5c15d8b"
                }
            }
        },
        "handlers.validateSampleResponse200": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "user": {
                            "$ref": "#/definitions/handlers.validateSampleResponse200User"
                        }
                    }
                },
                "message": {
                    "type": "string",
                    "example": "Request successful"
                },
                "status": {
                    "type": "string",
                    "example": "success"
                },
                "statusCode": {
                    "type": "integer",
                    "example": 200
                },
                "token": {
                    "type": "string",
                    "example": "b6d4a7e1d2d841a1afe874a2a5c15d8b"
                }
            }
        },
        "handlers.validateSampleResponse200User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "user@example.com"
                },
                "has2FA": {
                    "type": "boolean",
                    "example": true
                },
                "id": {
                    "type": "string",
                    "example": "d38f91b2-dc3b-4f9d-aeb4-7b95c91e9d08"
                },
                "isFreezed": {
                    "type": "boolean",
                    "example": true
                },
                "isVerified": {
                    "type": "boolean",
                    "example": true
                },
                "lastLogin": {
                    "type": "string",
                    "example": "2024-07-22T14:30:00Z"
                },
                "role": {
                    "$ref": "#/definitions/handlers.validateSampleResponseRole"
                }
            }
        },
        "handlers.validateSampleResponseRole": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "admin"
                },
                "permissions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "canTransact",
                        " canWithdraw"
                    ]
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1/",
	Schemes:          []string{},
	Title:            "FamTrust API Backend - Auth",
	Description:      "This is the Authentication and Authorization Micro-service for the FamTrust Web App.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
