package dto

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserSignup struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Bio      string `json:"bio" binding:"-"`
	Role     string `json:"role" binding:"-"`
}

type UserUpdate struct {
	Email string `json:"email" binding:"-"`
	Bio   string `json:"bio" binding:"-"`
	Role  string `json:"role" binding:"-"`
}

type RetrieveUserInfo struct {
	Username string `uri:"username" json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
}
