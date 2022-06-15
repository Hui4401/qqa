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

	if err := r.Run(":8000"); err != nil {
		panic("run server error")
	}
}
