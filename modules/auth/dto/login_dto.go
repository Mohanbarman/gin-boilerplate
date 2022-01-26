package auth

type LoginDto struct {
	Email    string `json:"email" form:"email" binding:"email" validators:"email"`
	Password string `json:"password" form:"password" binding:"required"`
}
