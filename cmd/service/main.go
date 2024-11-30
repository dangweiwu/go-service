package main

import (
	flags "github.com/jessevdk/go-flags"
	"go-service/internal/option"
	"os"
)

func main() {
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