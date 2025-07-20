package option

import (
	"bufio"
	"go-service/internal/bootstrap"
	"go-service/internal/config"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/dangweiwu/microkit/yamlconfig"
)

/*
*
启动api
*/
type RunService struct {
	ApiHost string `long:"apihost" description:"api启动host"`
}

func (this *RunService) Execute(args []string) error {
	var c config.Config

	loadEnvFile()

	yamlconfig.MustLoad(Opt.ConfigPath, &c)

	boot := bootstrap.NewBootStrap()
	err := boot.Init(c)
	if err != nil {
		return err
	}
	boot.Run()
	return nil
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
