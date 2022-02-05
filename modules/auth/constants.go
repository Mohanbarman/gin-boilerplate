package auth

import (
	"net/http"

	"example.com/lib"
)

const (
	EmailExistsErr = iota + 1
	EmailNotFoundErr
	UserNotFoundErr
	WrongPasswordErr
	TokenGenerateErr
	HashingPassErr
	ResetPasswordLinkExpErr
	SomethingWentWrongErr
)

var HttpErrors = map[int]*lib.HttpResponseStruct{
	EmailExistsErr: lib.HttpResponse(http.StatusBadRequest).Errors(lib.H{
		"password": "Email already exists",
	}),
	UserNotFoundErr: lib.HttpResponse(http.StatusNotFound).Message("User not found"),
	EmailNotFoundErr: lib.HttpResponse(http.StatusNotFound).Errors(lib.H{
		"email": "User not found",
	}),
	WrongPasswordErr: lib.HttpResponse(http.StatusBadRequest).Errors(lib.H{
		"password": "Invalid credentials",
	}),
	TokenGenerateErr:        lib.HttpResponse(http.StatusInternalServerError).Message("Failed to create user please contact customer care"),
	ResetPasswordLinkExpErr: lib.HttpResponse(http.StatusBadRequest).Message("Reset password link expired"),
	SomethingWentWrongErr:   lib.HttpResponse(http.StatusInternalServerError).Message("Something went wrong"),
}
