package dto

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserSignup struct {
	Name     string `json:"name,omitempty" binding:"required"`
	Address  string `json:"address"`
	Username string `json:"username,omitempty" binding:"required"`
	Email    string `gorm:"unique" json:"email" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}

type UserUpdate struct {
	Name     string `json:"name,omitempty" binding:"required"`
	Address  string `json:"address"`
	Username string `json:"username,omitempty" binding:"required"`
	Email    string `gorm:"unique" json:"email" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}

type RetrieveUserInfo struct {
	Name     string `json:"name,omitempty" binding:"required"`
	Address  string `json:"address"`
	Username string `json:"username,omitempty" binding:"required"`
	Email    string `gorm:"unique" json:"email" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}
