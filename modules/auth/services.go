package auth

import (
	"example.com/config"
	"example.com/lib"
	"example.com/lib/jwt"
	"example.com/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	Config *config.Config
	Db     *gorm.DB
}

func (service *AuthService) Register(registerDto *RegisterDto) (result lib.H, e *lib.ServiceError) {
	user := models.UserModel{
		Email:  registerDto.Email,
		Name:   registerDto.Name,
		Status: registerDto.Name,
	}

	err := user.SetPassword(registerDto.Password)

	if err != nil {
		e = lib.Error(HashingPassErr)
		return
	}

	if created := service.Db.Create(&user); created.Error != nil {
		e = lib.Error(EmailExistsErr)
		return
	}

	result = user.Transform()
	return
}

func (service AuthService) Login(loginDto *LoginDto, jwtService *jwt.JwtService) (result lib.H, e *lib.ServiceError) {
	user := models.UserModel{}

	records := service.Db.Find(&user, &models.UserModel{Email: loginDto.Email})

	if records.RowsAffected <= 0 {
		e = lib.Error(UserNotFoundErr)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password)); err != nil {
		e = lib.Error(WrongPasswordErr)
		return
	}

	accessToken, aerr := jwtService.SignToken(user.UUID, jwt.AccessToken)
	refreshToken, rerr := jwtService.SignToken(user.UUID, jwt.RefreshToken)

	if aerr != nil || rerr != nil {
		e = lib.Error(TokenGenerateErr)
	}

	result = user.Transform()
	result["access_token"] = accessToken
	result["refresh_token"] = refreshToken

	return
}
