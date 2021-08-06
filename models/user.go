package models

type User struct {
	Id       uint   `json:"id" binding:"required" gorm:"autoIncrement;primaryKey"`
	Name     string `json:"name" binding:"required" gorm:"not null;size:30"`
	Nickname string `json:"nickname" gorm:"not null;size:10"`
	Password string `json:"password" gorm:"not null;size:30" binding:"required"`
	Token    string `json:"token"`
	Coin     int    `json:"coin" gorm:"default:0"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
