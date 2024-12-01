package option

import "fmt"

var (
	Versionstr = "None"
	GitCommit  = "None"
	GitBranch  = "None"
	BuildTS    = "None"
)

type Version struct{}

func (this Version) Execute(args []string) error {
	const padding = 20 // 设置字段宽度

	fmt.Printf("%-*s %s\n", padding, "Version:", Versionstr)
	fmt.Printf("%-*s %s\n", padding, "Git Commit:", GitCommit)
	fmt.Printf("%-*s %s\n", padding, "Git Branch:", GitBranch)
	fmt.Printf("%-*s %s\n", padding, "Build Time:", BuildTS)
	return nil
}
