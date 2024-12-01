package option

var Opt struct {
	ConfigPath    string        `long:"config" short:"f" description:"配置文件路径"`
	Version       Version       `command:"version" long:"version" short:"v" description:"版本信息"`
	RunService    RunService    `command:"run" namespace:"group2" env-namespace:"group2" description:"api启动host"`
	InitTable     InitTable     `command:"inittable"  description:"初始化数据库"`
	InitSuperUser InitSuperUser `command:"inituser" description:"初始化超级管理员、重置超级管理员密码"`
}
