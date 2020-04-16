package main

import (
	calc_html "github.com/ms-uzh/calc/html"

	"github.com/ms-uzh/calc/config"
)

func main() {
	conf := config.ReadConfigs(
		"./config/head.yaml",
		"./config/polyamine.yaml",
		"./config/tail.yaml",
		"./config/spider.yaml",
		"./config/app.yaml",
	)
	calc_html.NewServer(conf).ListenAndServe()
}
