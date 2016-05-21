// httpserver
package main

import (
	"httpserver/house"

	"github.com/codegangsta/martini"
)

func main() {
	m := martini.Classic()
	//house.QueryPositonNew("新龙城")

	m.Get("/house/:position/list", house.List)
	m.Get("/house/:position/list/new", house.ListNew)
	m.Get("/house/:position/list/changed", house.ListChanged)
	m.Get("/house/:position/list/:date", house.ListHistory)

	m.Get("/", func() (int, string) {
		return 200, "hello, world"
	})

	m.Run()
}
