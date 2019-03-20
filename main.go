package main

import (
	calc_html "github.com/fforootd/calc/html"

	"github.com/fforootd/calc/config"
)

func main() {
	conf := config.ReadConfigs(
		"./config/head.yaml",
		"./config/polyamine.yaml",
		"./config/tail.yaml",
		"./config/app.yaml",
	)
	calc_html.NewServer(conf).ListenAndServe()
}
