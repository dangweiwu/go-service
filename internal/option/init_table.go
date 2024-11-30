package option

type InitTable struct {
}

func (*InitTable) Usage() string {
	return `//迁移数据库结构，但不会删除未使用的列。`
}

func (this *InitTable) Execute(args []string) error {
	return nil
}
