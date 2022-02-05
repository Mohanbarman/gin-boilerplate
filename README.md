<h1 align="center"> Gin Boilerplate </h1> <br>
<p align="center"> Build apis with gin faster with this template</p>

# Features

- [Validation errors](#validation-errors)
- [Authentication middleware](#authentication-middleware)
- [Http response builder](#http-response-builder)
- [Request body validation with validation middleware](#request-body-validation-with-validation-middleware)
## Validation errors
```
POST /api/register HTTP/1.1
Content-Type: application/json

{"email": "john.com", "password": "secure"}
```
Response
```json
{
    "code": 400,
    "success": false,
    "errors": {
        "email": [
            "invalid email"
        ]
    }
}
```

## Authentication middleware
```go
// building authentication middleware
authMiddleware := middlewares.AuthMiddleware{
    Jwt: &jwtService,
    DB:  db,
}

// initializing route with authentication middleware
router.GET("/me", authMiddleware.Validate(jwt.AccessToken), authCtrl.GetMe())

// getting user in controller
user := c.MustGet("user").(*models.UserModel)
```

## Http response builder
```go
lib.HttpResponse(200).Message("User registered successfully").Send(c)
```
Response
```json
{
     "code": "200",
     "success": true,
     "message": "User registered successfully"
}
```

## Request body validation with validation middleware
```go
// dto
type RegisterDto struct {
	Email    string `json:"email" form:"email" binding:"required,email,max=100"`
	Password string `json:"password" form:"password" binding:"required,max=100,min=8"`
	Name     string `json:"name" form:"name" binding:"required,max=100"`
	Status   string `json:"status" form:"status" binding:"required,max=150"`
}

// adding validation middleware in route
router.POST("/login", middlewares.Validate(&LoginDto{}), authCtrl.Login())

// retreiving the dto struct in controller
dto := c.MustGet("data").(*LoginDto)
```
