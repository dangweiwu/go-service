package authmodel

import "go-service/internal/service/apiservice/pkg/model"

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
	Memo     string   `json:"memo" gorm:"type:text;comment:备注" binding:"max=300" doc:"|d 备注"`
}

func (AuthPo) TableName() string {
	return AUTH_TABLE
}

type AuthForm struct {
	model.Model `swaggerignore:"true"`
	Name        string `json:"name" gorm:"size:100;comment:名称" binding:"max=100,required"`                         //权限名称
	Code        string `json:"code" gorm:"size:100;not null;unique;comment:权限ID"  binding:"required,max=100"`      //权限编码 该编码与前端auth控制对应
	OrderNum    int    `json:"order_num" gorm:"default:0;comment:排序" binding:"" `                                  //排序
	Api         string `json:"api" gorm:"size:200;default:'';comment:接口" binding:"max=200"`                        //api URL
	Method      string `json:"method" gorm:"size:10;default:'';comment:请求方式" binding:"max=50" `                    //GET | POST | PUT | DELETE
	Kind        string `json:"kind" gorm:"type:enum('1','0');default:'0';comment:0 api 1 菜单" binding:"oneof=0 1" ` //类型 0 api | 1 菜单
	ParentId    int    `json:"parent_id" gorm:"default:0;not null;comment:上级菜单ID" `                                //父类ID
	Memo        string `json:"memo" gorm:"type:text;comment:备注" binding:"max=300"`                                 //备注
}

func (AuthForm) TableName() string {
	return AUTH_TABLE
}

type AuthUpdateForm struct {
	ID       int64  `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	Name     string `json:"name" gorm:"size:100;comment:名称" binding:"max=100,required"`                        //权限名称
	OrderNum int    `json:"order_num" gorm:"default:0;comment:排序" binding:""`                                  //排序
	Api      string `json:"api" gorm:"size:200;default:'';comment:接口" binding:"max=200"`                       //API
	Method   string `json:"method" gorm:"size:10;default:'';comment:请求方式" binding:"max=50"`                    //GET POST PUT DELETE
	Kind     string `json:"kind" gorm:"type:enum('1','0');default:'0';comment:0 api 1 菜单" binding:"oneof=0 1"` //类型 0 api | 1 菜单
	ParentId int    `json:"parent_id" gorm:"default:0;not null;comment:上级菜单ID" `                               //父类ID
	Memo     string `json:"memo" gorm:"type:text;comment:备注" binding:"max=300"`                                //备注
}

func (AuthUpdateForm) TableName() string {
	return AUTH_TABLE
}

type AuthVo struct {
	ID       int64    `json:"id" gorm:"primaryKey"`
	Name     string   `json:"name" gorm:"size:100;comment:名称" binding:"max=100,required"`                        //权限名称
	Code     string   `json:"code" gorm:"size:100;not null;unique;comment:权限ID"  binding:"required,max=100"`     //权限编码
	OrderNum int      `json:"order_num" gorm:"default:0;comment:排序" binding:"" `                                 //排序
	Api      string   `json:"api" gorm:"size:200;default:'';comment:接口" binding:"max=200"`                       //API
	Method   string   `json:"method" gorm:"size:10;default:'';comment:请求方式" binding:"max=50"`                    // GET POST PUT DELETE
	Kind     string   `json:"kind" gorm:"type:enum('1','0');default:'0';comment:0 api 1 菜单" binding:"oneof=0 1"` //类型 0 api | 1 菜单
	ParentId int      `json:"parent_id" gorm:"default:0;not null;comment:上级菜单ID" `                               //父级ID
	Children []AuthVo `json:"children" gorm:"foreignkey:ParentId"`                                               //子集
	Memo     string   `json:"memo" gorm:"type:text;comment:备注" binding:"max=300"`                                //备注
}

func (AuthVo) TableName() string {
	return AUTH_TABLE
}
