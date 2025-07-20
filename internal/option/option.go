package option

var Opt struct {
	ConfigPath    string        `long:"config" short:"f" description:"配置文件路径"`
	EnvFile       string        `long:"envfile" short:"e" description:"环境变量文件路径"`
	Version       Version       `command:"version" description:"版本信息"`
	RunService    RunService    `command:"run" description:"api启动host"`
	InitTable     InitTable     `command:"inittable"  description:"初始化数据库"`
	InitSuperUser InitSuperUser `command:"inituser" description:"初始化超级管理员、重置超级管理员密码"`
}
