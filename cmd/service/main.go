package main

import (
	flags "github.com/jessevdk/go-flags"
	"go-service/internal/option"
	"os"
)

var (
	GitCommit = "None"
	Version   = "None"
	GitBranch = "None"
	BuildTS   = "None"
)

// @title           管理系统
// @version         1.0.0
// @description     通用框架
// @description.markdown api.md

// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @tag.name  1-系统-我的
// @tag.description 登录、退出、token刷新、我的详情、信息修改、密码修改。
// @tag.name  2-系统-用户管理
// @tag.description 用户添加、修改、删除、重置密码、角色配置。
// @tag.name  3-系统-权限管理
// @tag.description 权限添加、修改、删除。正式发版前，此页面删除。
// @tag.name  4-系统-角色管理
// @tag.description 角色添加、修改、删除、权限配置。
func main() {
	option.Versionstr = Version
	option.GitCommit = GitCommit
	option.GitBranch = GitBranch
	option.BuildTS = BuildTS

	p := flags.NewParser(&option.Opt, flags.Default)
	p.ShortDescription = "v1.0 server"
	p.LongDescription = `v1.0 Server`

	if _, err := p.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}
}
