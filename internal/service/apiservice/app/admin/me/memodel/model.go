package memodel

import (
	"go-service/internal/service/apiservice/pkg/model"
	"strconv"
)

// redis login id
func GetAdminRedisLoginId(id int) string {
	return "lgn:" + strconv.Itoa(id)
}

// my info
// @doc | memodel.MeInfo
type MeInfo struct {
	Account      string `json:"account" gorm:"type:varchar(50);unique;comment:账号" binding:"required"`                     //账号
	Phone        string `json:"phone" gorm:"type:varchar(50);unique;comment:电话" binding:"max=11"`                         //电话
	Name         string `json:"name" gorm:"size:100;not null;default:'';comment:名称" binding:"max=100"`                    //姓名
	Memo         string `json:"memo" gorm:"type:text;comment:备注" binding:"max=300" `                                      //备注
	Email        string `json:"email" gorm:"type:varchar(100);default:'';comment:邮件" binding:"omitempty,email"`           //Email
	IsSuperAdmin string `json:"is_super_admin" gorm:"type:enum('1','0');default:'0';comment:是否超级管理员" binding:"oneof=0 1"` // 是否超级管理员 0不是 1是
	Role         string `json:"role" gorm:"size:100;not null;index;comment:角色"`                                           //角色代码
}

type MeInfoVo struct {
	Account      string `json:"account" gorm:"type:varchar(50);unique;comment:账号" binding:"required" doc:"|d 账号"`
	Phone        string `json:"phone" gorm:"type:varchar(50);unique;comment:电话" binding:"max=11" doc:"|d 电话"`
	Name         string `json:"name" gorm:"size:100;not null;default:'';comment:名称" binding:"max=100" doc:"|d 姓名"`
	Memo         string `json:"memo" gorm:"type:text;comment:备注" binding:"max=300" doc:"|d 备注"`
	Email        string `json:"email" gorm:"type:varchar(100);default:'';comment:邮件" binding:"omitempty,email" doc:"|d Email"`
	IsSuperAdmin string `json:"is_super_admin" gorm:"type:enum('1','0');default:'0';comment:是否超级管理员" binding:"oneof=0 1" doc:"|d 是否超级管理员 |c 0不是 1是"`
	Role         string `json:"role" gorm:"size:100;not null;index;comment:角色" doc:"|d 角色ID"` //角色代码
	RoleName     string `json:"role_name"`
}

var MeViewField = []string{"account", "phone", "name", "memo", "email", "is_super_admin", "role"}

func (MeInfo) TableName() string {
	return "admin"
}

// my update
type MeForm struct {
	model.Model `swaggerignore:"true"`
	Phone       string `json:"phone" gorm:"type:varchar(50);unique;comment:电话" binding:"max=11"`                //电话
	Name        string `json:"name" gorm:"size:100;not null;default:'';comment:名称" binding:"max=100"`           //姓名
	Memo        string `json:"memo" gorm:"type:text;comment:备注" binding:"max=300" `                             //备注
	Email       string `json:"email" gorm:"type:varchar(100);default:'';comment:邮件" binding:"omitempty,email" ` //email
}

var MeUpdateField = []string{"id", "updated_at", "email", "phone", "name", "memo"}

func (MeForm) TableName() string {
	return "admin"
}

// log form
type LoginForm struct {
	Account  string `json:"account" binding:"required"`  //账号
	Password string `json:"password" binding:"required"` //密码
}

// 登陆返回token
type LogRep struct {
	AccessToken  string `json:"access_token"`  //access token
	RefreshToken string `json:"refresh_token"` //刷新 token
	TokenExp     int64  `json:"token_exp"`     //有效期时间戳，秒
}

// 修改密码
type PasswordForm struct {
	Password    string `json:"password" binding:"required"`     //原始密码
	NewPassword string `json:"new_password" binding:"required"` //新密码
}
