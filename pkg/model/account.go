// Package model provides ...
package model

import "time"

// use snake_case as column name

// Account 博客账户
type Account struct {
	Username  string     `gorm:"column:username;primaryKey;size:64" bson:"username"`                          // 用户名
	Password  string     `gorm:"column:password;not null;size:64" bson:"password"`                            // 密码
	Email     string     `gorm:"column:email;not null;size:64" bson:"email"`                                  // 邮件地址
	PhoneN    string     `gorm:"column:phone_n;not null;size:64" bson:"phone_n"`                              // 手机号
	Address   string     `gorm:"column:address;not null;size:128" bson:"address"`                             // 地址信息
	LogoutAt  *time.Time `gorm:"type:datetime;column:logout_at" bson:"logout_at"`                             // 登出时间
	LoginIP   string     `gorm:"column:login_ip;not null;size:64" bson:"login_ip"`                            // 最近登录IP
	LoginUA   string     `gorm:"column:login_ua;not null;size:64" bson:"login_ua"`                            // 最近登录IP
	LoginAt   time.Time  `gorm:"type:datetime;column:login_at;default:current_timestamp" bson:"login_at"`     // 最近登录时间
	CreatedAt time.Time  `gorm:"type:datetime;column:created_at;default:current_timestamp" bson:"created_at"` // 创建时间
}
