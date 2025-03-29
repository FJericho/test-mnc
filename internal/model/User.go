package model

type User struct {
    ID       uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    Fullname string `json:"fullname"`
    Username string `json:"username"`
    Password string `json:"password"`
    Balance int `json:"balance" valid:"range(0|100000000)"`
    History []History `json:"history"`
    Payment []Payment `json:"payment"`
}

type Login struct {
    Username string `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
    Password string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
}