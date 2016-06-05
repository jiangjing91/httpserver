// httpserver
package main

import (
	"fmt"
	"httpserver/house"
	"os"

	"httpserver/dbaccess"

	"github.com/codegangsta/martini"
	"github.com/widuu/goini"
)

func main() {

	var configFile string
	if len(os.Args) > 1 {
		fmt.Println(os.Args)
		configFile = os.Args[1]

		if len(configFile) == 0 {
			fmt.Println("need a config file")
			return
		}
	} else {
		configFile = "config.ini"
	}

	fmt.Println(configFile)

	config := goini.SetConfig(configFile)

	env := config.GetValue("ENV", "env")
	port := config.GetValue(env, "port")
	DbUsername := config.GetValue(env, "dbusername")
	DbPassword := config.GetValue(env, "dbpasswd")
	DbName := config.GetValue(env, "dbname")

	db := dbaccess.DbConnect(DbUsername, DbPassword, DbName)

	m := martini.Classic()

	if env == "ONLINE" {
		martini.Env = martini.Prod
	}

	m.Map(db)

	m.Get("/house/:position/list", house.List)
	m.Get("/house/:position/list/new", house.ListNew)
	m.Get("/house/:position/list/changed", house.ListChanged)
	m.Get("/house/:position/list/:date", house.ListHistory)

	m.Get("/", func() (int, string) {
		return 200, "hello, world"
	})
	m.Use(martini.Static("web"))

	m.RunOnAddr(":" + port)
}
