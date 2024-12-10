package authmodel

import "go-service/internal/apiservice/pkg/model"

const AUTH_TABLE = "auth"

type AuthPo struct {
	model.Model
	Name     string   `json:"name" gorm:"size:100;comment:名称" binding:"max=100"`
	Code     string   `json:"code" gorm:"size:100;not null;unique;comment:权限ID" binding:"required,max=100"`
	OrderNum int      `json:"order_num" gorm:"default:0;comment:排序" binding:""`
	Api      string   `json:"api" gorm:"size:200;default:'';comment:接口" binding:"max=200"`
	Method   string   `json:"method" gorm:"size:10;default:'';comment:请求方式" binding:"max=50"`
	Kind     string   `json:"kind" gorm:"type:enum('1','0');default:'0';comment:0 api 1 菜单" binding:"oneof=0 1"` //0 api 1 菜单
	ParentId int      `json:"parent_id" gorm:"default:0;not null;comment:上级菜单ID"`
	Children []AuthPo `json:"children" gorm:"foreignkey:ParentId"`
}

func (AuthPo) TableName() string {
	return AUTH_TABLE
}

// @doc | authmodel.AuthForm
type AuthForm struct {
	model.Model
	Name     string `json:"name" gorm:"size:100;comment:名称" binding:"max=100,required" doc:"|d 权限名称"`
	Code     string `json:"code" gorm:"size:100;not null;unique;comment:权限ID"  binding:"required,max=100" doc:"|d 权限ID"`
	OrderNum int    `json:"order_num" gorm:"default:0;comment:排序" binding:""  doc:"|d 排序"`
	Api      string `json:"api" gorm:"size:200;default:'';comment:接口" binding:"max=200" doc:"|d API"`
	Method   string `json:"method" gorm:"size:10;default:'';comment:请求方式" binding:"max=50"  doc:"|d 方法 |e GET"`                           // GET POST PUT DELETE
	Kind     string `json:"kind" gorm:"type:enum('1','0');default:'0';comment:0 api 1 菜单" binding:"oneof=0 1"  doc:"|d 类型 |c 0:api 1:菜单"` //0 api 1 菜单
	ParentId int    `json:"parent_id" gorm:"default:0;not null;comment:上级菜单ID"  doc:"|d 上级ID"`                                            //父类ID
}

func (AuthForm) TableName() string {
	return AUTH_TABLE
}

// @doc| authmodel.AuthUpdateForm
type AuthUpdateForm struct {
	model.Model
	Name     string `json:"name" gorm:"size:100;comment:名称" binding:"max=100,required" doc:"|d 权限名称"`
	OrderNum int    `json:"order_num" gorm:"default:0;comment:排序" binding:""  doc:"|d 排序"`
	Api      string `json:"api" gorm:"size:200;default:'';comment:接口" binding:"max=200" doc:"|d API"`
	Method   string `json:"method" gorm:"size:10;default:'';comment:请求方式" binding:"max=50"  doc:"|d 方法 |e GET"`                           // GET POST PUT DELETE
	Kind     string `json:"kind" gorm:"type:enum('1','0');default:'0';comment:0 api 1 菜单" binding:"oneof=0 1"  doc:"|d 类型 |c 0:api 1:菜单"` //0 api 1 菜单
	ParentId int    `json:"parent_id" gorm:"default:0;not null;comment:上级菜单ID"  doc:"|d 上级ID"`                                            //父类ID
}

func (AuthUpdateForm) TableName() string {
	return AUTH_TABLE
}

// @doc| authmodel.AuthVo
type AuthVo struct {
	ID       int64    `json:"id" gorm:"primaryKey"`
	Name     string   `json:"name" gorm:"size:100;comment:名称" binding:"max=100,required" doc:"|d 权限名称"`
	Code     string   `json:"code" gorm:"size:100;not null;unique;comment:权限ID"  binding:"required,max=100" doc:"|d 权限ID"`
	OrderNum int      `json:"order_num" gorm:"default:0;comment:排序" binding:""  doc:"|d 排序"`
	Api      string   `json:"api" gorm:"size:200;default:'';comment:接口" binding:"max=200" doc:"|d API"`
	Method   string   `json:"method" gorm:"size:10;default:'';comment:请求方式" binding:"max=50"  doc:"|d 方法 |e GET"`                           // GET POST PUT DELETE
	Kind     string   `json:"kind" gorm:"type:enum('1','0');default:'0';comment:0 api 1 菜单" binding:"oneof=0 1"  doc:"|d 类型 |c 0:api 1:菜单"` //0 api 1 菜单
	ParentId int      `json:"parent_id" gorm:"default:0;not null;comment:上级菜单ID"  doc:"|d 上级ID"`
	Children []AuthVo `json:"children" gorm:"foreignkey:ParentId" doc:"|d 子集 |t []self "`
}

func (AuthVo) TableName() string {
	return AUTH_TABLE
}
