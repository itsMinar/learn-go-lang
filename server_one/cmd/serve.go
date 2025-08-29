package cmd

import (
	"my_server/config"
	"my_server/rest"
)

func Serve() {
	cnf := config.GetConfig()

	rest.Start(cnf)
}
