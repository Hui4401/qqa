package main

import (
	"qa/conf"
	"qa/cron"
	"qa/routes"
)

func main() {

	conf.Init()

	cron.StartSchedule()

	r := routes.NewRouter()

	r.Run(":8000")
}
