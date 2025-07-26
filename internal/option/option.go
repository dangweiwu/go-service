package option

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

var Opt struct {
	ConfigPath    string        `long:"config" short:"f" description:"配置文件路径"`
	EnvFile       string        `long:"envfile" short:"e" description:"环境变量文件路径"`
	Version       Version       `command:"version" description:"版本信息"`
	RunService    RunService    `command:"run" description:"api启动host"`
	InitTable     InitTable     `command:"inittable"  description:"初始化数据库"`
	InitSuperUser InitSuperUser `command:"inituser" description:"初始化超级管理员、重置超级管理员密码"`
}

func loadEnvFile() {
	if len(Opt.EnvFile) <= 0 {
		return
	}
	// try to open the environment file
	f, err := os.Open(Opt.EnvFile)
	if err != nil {
		log.Printf("Fail to open environment file, %s, err: %v \n", Opt.EnvFile, err)
		return
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	for {
		// for each line
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		// if line starts with '#', it is a comment line, ignore it
		line = strings.TrimSpace(line)
		if len(line) > 0 && line[0] == '#' {
			continue
		}
		// if environment variable is exported with "export"
		if strings.HasPrefix(line, "export") && len(line) > len("export") && unicode.IsSpace(rune(line[len("export")])) {
			line = strings.TrimSpace(line[len("export"):])
		}
		// split the environment variable with "="
		pos := strings.Index(line, "=")
		if pos != -1 {
			k := strings.TrimSpace(line[0:pos])
			v := strings.TrimSpace(line[pos+1:])
			// if key and value are not empty, put it into the environment
			if len(k) > 0 && len(v) > 0 {
				os.Setenv(k, v)
			}
		}
	}
}
