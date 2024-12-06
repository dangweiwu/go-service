package adminmodel

import "go-service/internal/apiservice/pkg/model"

const ADMIN_TABLE = "admin"

type AdminPo struct {
	model.Model
	Account      string `json:"account" gorm:"type:varchar(50);unique;comment:账号" binding:"required"`                     //账号
	Phone        string `json:"phone" gorm:"type:varchar(50);comment:电话" binding:"max=11"`                                //手机号 只进行11位校验
	Name         string `json:"name" gorm:"size:100;not null;default:'';comment:名称" binding:"max=100"`                    //名称
	Status       string `json:"status" gorm:"type:enum('0','1');default:'1';comment:'0:无效|1有效'" binding:"oneof=0 1"`      //状态 1 close 0 open
	Password     string `json:"password" gorm:"size:100;not null;comment:密码" binding:"max=100,required"`                  //密码
	Memo         string `json:"memo" gorm:"type:text;comment:备注" binding:"max=300"`                                       //备注
	Email        string `json:"email" gorm:"type:varchar(100);default:'';comment:邮件" binding:"omitempty,email"`           //email 非空则进行email有效性校验
	IsSuperAdmin string `json:"is_super_admin" gorm:"type:enum('1','0');default:'0';comment:是否超级管理员" binding:"oneof=0 1"` // 是否超级管理员 1:是 0:否
	Role         string `json:"role" gorm:"size:100;not null;index;comment:角色"`                                           //角色
}

func (AdminPo) TableName() string {
	return ADMIN_TABLE
}

// @doc | adminmodel.AdminForm
type AdminForm struct {
	model.Model
	Account      string `json:"account" gorm:"type:varchar(50);unique;comment:账号" binding:"required" doc:"|d 账号 |e admin "`
	Phone        string `json:"phone" gorm:"type:varchar(50);comment:电话" binding:"max=11" doc:"|d 手机号 |e 123456789"`
	Name         string `json:"name" gorm:"size:100;not null;default:'';comment:名称" binding:"max=100" doc:"|d 姓名 |e 张三"`
	Status       string `json:"status" gorm:"type:enum('0','1');default:'1';comment:'0:无效|1有效'" binding:"oneof=0 1" doc:"|d 状态 |e 1 |c 0 无效 1 有效 |t string"` //状态 0无效|1有效 0时会把当前在线人员踢下线
	Password     string `json:"password" gorm:"size:100;not null;comment:密码" binding:"max=100,required" doc:"|d 密码 |e 123456"`
	Memo         string `json:"memo" gorm:"type:text;comment:备注" binding:"max=300" doc:"|d 备注"`
	Email        string `json:"email" gorm:"type:varchar(100);default:'';comment:邮件" binding:"omitempty,email" doc:"|d Email"`
	IsSuperAdmin string `json:"is_super_admin" gorm:"type:enum('1','0');default:'0';comment:是否超级管理员" binding:"oneof=0 1" doc:"|d 是否超级管理员 |e 1 |c 0:否 1:是 |t string"` // 是否超级管理员 1:是 0:否
	Role         string `json:"role" gorm:"size:100;not null;index;comment:角色" binding:"max=100" doc:"|d 角色ID"`                                                      //角色

}

func (AdminForm) TableName() string {
	return ADMIN_TABLE
}

// get用 不带password
// @doc | adminmodel.AdminVo
type AdminVo struct {
	model.Model
	Account      string `json:"account" gorm:"type:varchar(50);unique;comment:账号" binding:"required" doc:"|d 账号 |e admin "`
	Phone        string `json:"phone" gorm:"type:varchar(50);comment:电话" binding:"max=11"  doc:"|d 手机号 |e 123456789"`
	Name         string `json:"name" gorm:"size:100;not null;default:'';comment:名称" binding:"max=100" doc:"|d 姓名 |e 张三"`
	Status       string `json:"status" gorm:"type:enum('0','1');default:'1';comment:'0:无效|1有效'" binding:"oneof=0 1" doc:"|d 状态 |e 1 |c 0 无效 1 有效 |t string"` //状态 0无效|1有效 0时会把当前在线人员踢下线
	Memo         string `json:"memo" gorm:"type:text;comment:备注" binding:"max=300" doc:"|d 备注"`
	Email        string `json:"email" gorm:"type:varchar(100);default:'';comment:邮件" binding:"omitempty,email" doc:"|d Email"`
	IsSuperAdmin string `json:"is_super_admin" gorm:"type:enum('1','0');default:'0';comment:是否超级管理员" binding:"oneof=0 1" doc:"|d 是否超级管理员 |e 1 |c 0:否 1:是 |t string"`
	Role         string `json:"role" gorm:"size:100;not null;index;comment:角色" binding:"max=100" doc:"|d 角色ID"`
}

var AdminViewField = []string{"id", "created_at", "updated_at", "account", "phone", "name", "status", "memo", "email", "is_super_admin", "role"}

func (AdminVo) TableName() string {
	return ADMIN_TABLE
}

// update用
// @doc | adminmodel.AdminUpdateForm
type AdminUpdateForm struct {
	model.Model
	Phone        string `json:"phone" gorm:"type:varchar(50);comment:电话" binding:"max=11" doc:"|d 手机号 |e 123456789"`
	Name         string `json:"name" gorm:"size:100;not null;default:'';comment:名称" binding:"max=100" doc:"|d 姓名 |e 张三"`
	Status       string `json:"status" gorm:"type:enum('0','1');default:'1';comment:'0无效|1有效'" binding:"oneof=0 1" doc:"|d 状态 |e 1 |c 0 无效 1 有效 |t string"` //状态 0无效|1有效 0时会把当前在线人员踢下线
	Memo         string `json:"memo" gorm:"type:text;comment:备注" binding:"max=300" doc:"|d 备注"`
	Email        string `json:"email" gorm:"type:varchar(100);default:'';comment:邮件" binding:"omitempty,email" doc:"|d Email"`                                       //Email
	IsSuperAdmin string `json:"is_super_admin" gorm:"type:enum('1','0');default:'0';comment:是否超级管理员" binding:"oneof=0 1" doc:"|d 是否超级管理员 |e 1 |c 0:否 1:是 |t string"` //是否超级管理员
	Role         string `json:"role" gorm:"size:100;not null;index;comment:角色" binding:"max=100" doc:"|d 角色ID"`                                                      //角色
}

var AdminUpdateField = []string{"id", "updated_at", "phone", "name", "status", "memo", "email", "is_super_admin", "role"}

func (AdminUpdateForm) TableName() string {
	return ADMIN_TABLE
}
