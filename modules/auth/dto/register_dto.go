package auth

type RegisterDto struct {
	Email    string `json:"email" form:"email" binding:"required,email,max=100"`
	Password string `json:"password" form:"password" binding:"required,max=100,min=8"`
	Name     string `json:"name" form:"name" binding:"required,max=100"`
	Status   string `json:"status" form:"status" binding:"required,max=150"`
}
