package rolemodel

import "go-service/internal/apiservice/pkg/model"

const _rolename = "role"

type RolePo struct {
	model.Model
	Code     string             `json:"code" gorm:"size:100;not null;unique;comment:角色ID" binding:"required,max=100"` //编码
	Name     string             `json:"name" gorm:"size:100;comment:角色名称" binding:"max=100" `                         //名称
	OrderNum int                `json:"order_num" gorm:"default:0;comment:排序"`                                        //排序
	Status   string             `json:"status" gorm:"type:enum('0','1');default:'1';comment:状态"`                      //状态 0 禁用 | 1启用
	Memo     string             `json:"memo" gorm:"type:text;comment:备注" binding:"max=300" `                          //备注
	Auth     model.List[string] `json:"auth" gorm:"type:json;comment:角色code列表" swaggertype:"array,string"`            //权限编码列表 eg [auth1,...]
}

func (RolePo) TableName() string {
	return _rolename
}

var ViewFields = []string{"id", "code", "name", "order_num", "status", "memo", "auth"}

type RoleForm struct {
	model.Model `swaggerignore:"true"`
	Code        string `json:"code" gorm:"size:100;not null;unique;comment:角色ID" binding:"required,max=100"` //编码 全服唯一，禁止重复
	Name        string `json:"name" gorm:"size:100;comment:角色名称" binding:"max=100"`                          //名称
	OrderNum    int    `json:"order_num" gorm:"default:0;comment:排序"`                                        //排序 建议6位编码12顶级菜单34当前菜单56接口编码
	Status      string `json:"status" gorm:"type:enum('0','1');default:'1';comment:状态"`                      //状态 1 启动 | 0禁用
	Memo        string `json:"memo" gorm:"type:text;comment:备注" binding:"max=300"`                           //备注
}

func (RoleForm) TableName() string {
	return _rolename
}

type RoleUpdate struct {
	model.Model `swaggerignore:"true"`
	Name        string `json:"name" gorm:"size:100;comment:角色名称" binding:"max=100"`                             //名称
	OrderNum    int    `json:"order_num" gorm:"default:0;comment:排序"`                                           // 排序 建议6位编码12顶级菜单34当前菜单56接口编码
	Status      string `json:"status" gorm:"type:enum('0','1');default:'1';comment:状态" doc:"|d 状态 |c 0:禁用1:启用"` //状态 1 启动 | 0禁用
	Memo        string `json:"memo" gorm:"type:text;comment:备注" binding:"max=300" `                             //备注
}

func (RoleUpdate) TableName() string {
	return _rolename
}

type RoleAuthForm struct {
	ID   int64              `json:"id" swaggerignore:"true" gorm:"primaryKey"`
	Auth model.List[string] `json:"auth" gorm:"type:json;comment:角色code列表"  swaggertype:"array,string"` //权限编码列表 eg [auth1,auth2...]
}

func (RoleAuthForm) TableName() string {
	return _rolename
}
