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

// @base| xx系统管理 | v0.0.1
// @desc| 系统管理 2024年 12月 10日
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
