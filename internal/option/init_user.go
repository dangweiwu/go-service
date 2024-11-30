package option

type InitSuperUser struct {
	Password string `long:"password" description:"超级管理员设置密码"`
}

func (this *InitSuperUser) Usage() string {
	return `//设置超级管理员密码`
}

func (this *InitSuperUser) Execute(args []string) error {
	return nil
}
